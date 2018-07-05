package config

import (
	"crypto/rsa"
	"crypto/x509"

	"github.com/satori/uuid"
	"k8s.io/client-go/tools/clientcmd/api/v1"
)

type Config struct {
	Version int

	ImageOffer     string
	ImagePublisher string
	ImageSKU       string
	ImageVersion   string

	// CAs
	EtcdCaKey            *rsa.PrivateKey
	EtcdCaCert           *x509.Certificate
	CaKey                *rsa.PrivateKey
	CaCert               *x509.Certificate
	FrontProxyCaKey      *rsa.PrivateKey
	FrontProxyCaCert     *x509.Certificate
	ServiceSigningCaKey  *rsa.PrivateKey
	ServiceSigningCaCert *x509.Certificate
	ServiceCatalogCaKey  *rsa.PrivateKey
	ServiceCatalogCaCert *x509.Certificate

	// container images for pods
	MasterEtcdImage            string
	MasterAPIImage             string
	MasterControllersImage     string
	BootstrapAutoapproverImage string
	ServiceCatalogImage        string
	TunnelImage                string
	SyncImage                  string
	TemplateServiceBrokerImage string

	// etcd certificates
	EtcdServerKey  *rsa.PrivateKey
	EtcdServerCert *x509.Certificate
	EtcdPeerKey    *rsa.PrivateKey
	EtcdPeerCert   *x509.Certificate
	EtcdClientKey  *rsa.PrivateKey
	EtcdClientCert *x509.Certificate

	// control plane certificates
	MasterServerKey          *rsa.PrivateKey
	MasterServerCert         *x509.Certificate
	TunnelKey                *rsa.PrivateKey
	TunnelCert               *x509.Certificate
	AdminKey                 *rsa.PrivateKey
	AdminCert                *x509.Certificate
	AggregatorFrontProxyKey  *rsa.PrivateKey
	AggregatorFrontProxyCert *x509.Certificate
	MasterKubeletClientKey   *rsa.PrivateKey
	MasterKubeletClientCert  *x509.Certificate
	MasterProxyClientKey     *rsa.PrivateKey
	MasterProxyClientCert    *x509.Certificate
	OpenShiftMasterKey       *rsa.PrivateKey
	OpenShiftMasterCert      *x509.Certificate

	ServiceCatalogServerKey     *rsa.PrivateKey
	ServiceCatalogServerCert    *x509.Certificate
	ServiceCatalogAPIClientKey  *rsa.PrivateKey
	ServiceCatalogAPIClientCert *x509.Certificate
	BootstrapAutoapproverKey    *rsa.PrivateKey
	BootstrapAutoapproverCert   *x509.Certificate

	// master-config configurables
	ImageConfigFormat string

	// misc control plane configurables
	ServiceAccountKey *rsa.PrivateKey
	SessionSecretAuth []byte
	SessionSecretEnc  []byte
	HtPasswd          []byte

	// kubeconfigs
	AdminKubeconfig                 *v1.Config
	MasterKubeconfig                *v1.Config
	ServiceCatalogAPIKubeconfig     *v1.Config
	BootstrapAutoapproverKubeconfig *v1.Config

	// nodes
	SSHKey                  *rsa.PrivateKey
	NodeBootstrapKey        *rsa.PrivateKey
	NodeBootstrapCert       *x509.Certificate
	NodeBootstrapKubeconfig *v1.Config

	// needed by import
	RegistryStorageAccount         string
	RegistryHTTPSecret             []byte
	AlertManagerProxySessionSecret []byte
	AlertsProxySessionSecret       []byte
	PrometheusProxySessionSecret   []byte
	ServiceCatalogClusterID        uuid.UUID
	RegistryKey                    *rsa.PrivateKey
	RegistryCert                   *x509.Certificate
	RouterKey                      *rsa.PrivateKey
	RouterCert                     *x509.Certificate

	// for development
	ImageResourceGroup string
	ImageResourceName  string

	TunnelHostname string
}
