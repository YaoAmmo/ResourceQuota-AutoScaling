package v1alpha1

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ResourcesQuotaAutoScalingGetter
}

type ResourcesQuotaAUtoScalingClient struct {
	restClient rest.Interface
}
