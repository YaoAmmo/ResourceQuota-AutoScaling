package v1alpha1

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"my.demo/ResourceQuota-AutoScaling/pkg/api/v1alpha1"
)

type ResourcesQuotaAutoScalingGetter interface {
	ResourcesQuotaAutoScaling(namespace string) ResourcesQuotaAutoScalingInterface
}

type ResourcesQuotaAutoScalingInterface interface {
	Create(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.CreateOptions) (*v1alpha1.ResourcequotaAutoscaling, error)
	Update(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.UpdateOptions) (*v1alpha1.ResourcequotaAutoscaling, error)
	Delete(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.DeleteOptions) error
	Get(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.GetOptions) (*v1alpha1.ResourcequotaAutoscaling, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.ResourcequotaAutoscalingList, error)
	DeleteCollection(ctx context.Context, listOpts metav1.ListOptions, deleteOpts metav1.DeleteOptions) error
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ct context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1alpha1.ResourcequotaAutoscaling, error)
}

type ResourcesQuotaAutoScaling struct {
	client rest.Interface
	ns     string
}

func (s *ResourcesQuotaAutoScaling) Create(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.CreateOptions) (result *v1alpha1.ResourcequotaAutoscaling, err error) {
	result = &v1alpha1.ResourcequotaAutoscaling{}
	err = s.client.Post().
		Namespace(s.ns).
		Name("resourcesQuotaAutoScaling").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourcesQuotaAutoScaling).
		Do(ctx).
		Into(result)
	return
}

func (s *ResourcesQuotaAutoScaling) Update(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.UpdateOptions) (*v1alpha1.ResourcequotaAutoscaling, error) {

}
