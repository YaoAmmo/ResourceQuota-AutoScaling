package v1alpha1

import "k8s.io/client-go/rest"

type ResourcesQuotaAutoScalingController interface {
	RESTClient() rest.Interface
	RQAScaling
}