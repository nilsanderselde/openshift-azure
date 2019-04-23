package fakerp

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/openshift-azure/pkg/cluster/kubeclient"
	"github.com/openshift/openshift-azure/test/clients/azure"
	"github.com/openshift/openshift-azure/test/sanity"
	"github.com/openshift/openshift-azure/test/util/log"
)

type target struct {
	DiscoveredLabels map[string]string `json:"discoveredLabels"`
	Labels           map[string]string `json:"labels"`
	ScrapeURL        string            `json:"scrapeURL"`
	LastError        string            `json:"lastError"`
	LastScrape       string            `json:"lastScrape"`
	Health           string            `json:"health"`
}

type targetsResponse struct {
	Status string `json:"status"`
	Data   struct {
		ActiveTargets  []target `json:"activeTargets"`
		DroppedTargets []target `json:"droppedTargets"`
	} `json:"data"`
}

var _ = Describe("Prometheus E2E tests [Fake]", func() {
	var (
		azurecli *azure.Client
	)

	BeforeEach(func() {
		var err error
		azurecli, err = azure.NewClientFromEnvironment(context.Background(), log.GetTestLogger(), false)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should register all the necessary prometheus targets", func() {
		token, err := sanity.Checker.Client.Admin.GetServiceAccountToken("openshift-monitoring", "prometheus-k8s")
		Expect(err).NotTo(HaveOccurred())

		route, err := sanity.Checker.Client.Admin.RouteV1.Routes("openshift-monitoring").Get("prometheus-k8s", meta_v1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())

		req, err := http.NewRequest(http.MethodGet, "https://"+route.Spec.Host+"/api/v1/targets", nil)
		Expect(err).NotTo(HaveOccurred())
		req.Header.Add("Authorization", "Bearer "+string(token))

		cli := &http.Client{
			Transport: &kubeclient.RetryingRoundTripper{
				Log: log.GetTestLogger(),
				RoundTripper: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
				Retries:    5,
				GetTimeout: 30 * time.Second,
			},
		}

		resp, err := cli.Do(req)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		var tr targetsResponse
		d := json.NewDecoder(resp.Body)
		err = d.Decode(&tr)
		Expect(err).NotTo(HaveOccurred())

		healthyTargets := map[string]int{}
		for _, t := range tr.Data.ActiveTargets {
			if t.Health == "up" {
				healthyTargets[t.Labels["job"]]++
			}
		}

		cs, err := azurecli.OpenShiftManagedClusters.Get(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())

		nodes, masters := int(*cs.Properties.MasterPoolProfile.Count), int(*cs.Properties.MasterPoolProfile.Count)
		for _, app := range cs.Properties.AgentPoolProfiles {
			nodes += int(*app.Count)
		}

		Expect(healthyTargets).To(Equal(map[string]int{
			"alertmanager-main":           3,
			"apiserver":                   masters,
			"canary":                      2,
			"etcd-metrics":                3,
			"router-stats":                3,
			"cluster-monitoring-operator": 1,
			"kube-controllers":            masters,
			"kube-state-metrics":          2,
			"kubelet":                     nodes * 2,
			"node-exporter":               nodes,
			"prometheus-k8s":              2,
			"prometheus-operator":         1,
		}))
	})
})
