package v1alpha1

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"my.demo/ResourceQuota-AutoScaling/pkg/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

type ResourcesQuotaAutoScalingGetter interface {
	ResourcesQuotaAutoScaling(namespace string) ResourcesQuotaAutoScalingInterface
}

//实现ResourcesQuotaAutoScaling的运行方法
type ResourcesQuotaAutoScalingInterface interface {
	Create(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.CreateOptions) (*v1alpha1.ResourcequotaAutoscaling, error)
	Update(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.UpdateOptions) (*v1alpha1.ResourcequotaAutoscaling, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.ResourcequotaAutoscaling, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.ResourcequotaAutoscalingList, error)
	DeleteCollection(ctx context.Context, listOpts metav1.ListOptions, deleteOpts metav1.DeleteOptions) error
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1alpha1.ResourcequotaAutoscaling, error)
}

type resourcesQuotaAutoScaling struct {
	client rest.Interface
	ns     string
}

func newResourcesQuotaAutoScaling(c *ResourcesQuotaAUtoScalingClient, namespace string) *resourcesQuotaAutoScaling {
	return &resourcesQuotaAutoScaling{
		client: c.RESTClient,
		ns:     namespace,
	}
}


func (s *resourcesQuotaAutoScaling) Create(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.CreateOptions) (result *v1alpha1.ResourcequotaAutoscaling, err error) {
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

func (s *resourcesQuotaAutoScaling) Update(ctx context.Context, resourcesQuotaAutoScaling *v1alpha1.ResourcequotaAutoscaling, opts metav1.UpdateOptions) (result *v1alpha1.ResourcequotaAutoscaling, err error) {
	result = &v1alpha1.ResourcequotaAutoscaling{}
	err = s.client.Put().
		Namespace(s.ns).
		Resource("resourcesQuotaAutoScaling").
		Name(resourcesQuotaAutoScaling.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourcesQuotaAutoScaling).
		Do(ctx).
		Into(result)
	return
}

func (s *resourcesQuotaAutoScaling) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error  {
	return s.client.Delete().
		Namespace(s.ns).
		Resource("resourcesQuotaAutoScaling").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (s *resourcesQuotaAutoScaling) Get(ctx context.Context, name string, opts metav1.GetOptions) (result *v1alpha1.ResourcequotaAutoscaling, err error)  {
	result = &v1alpha1.ResourcequotaAutoscaling{}
	err = s.client.Get().
		Namespace(s.ns).
		Resource("resourcesQuotaAutoScaling").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (s *resourcesQuotaAutoScaling) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.ResourcequotaAutoscalingList, err error)  {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourcequotaAutoscalingList{}
	err = s.client.Get().
		Namespace(s.ns).
		Resource("resourcesQuotaAutoScaling").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (s *resourcesQuotaAutoScaling) DeleteCollection(ctx context.Context, listOpts metav1.ListOptions, deleteOpts metav1.DeleteOptions) error  {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return s.client.Delete().
		Namespace(s.ns).
		Resource("resourcesQuotaAutoScaling").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&deleteOpts).
		Do(ctx).
		Error()
}

func (s *resourcesQuotaAutoScaling) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)  {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return s.client.Get().
		Namespace(s.ns).
		Resource("resourcesQuotaAutoScaling").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (s *resourcesQuotaAutoScaling) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.ResourcequotaAutoscaling, err error)  {
	result = &v1alpha1.ResourcequotaAutoscaling{}
	err = s.client.Patch(pt).
		Namespace(s.ns).
		Resource("resourcesQuotaAutoScaling").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}