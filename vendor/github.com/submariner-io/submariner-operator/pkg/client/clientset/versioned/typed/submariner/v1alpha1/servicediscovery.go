/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/submariner-io/submariner-operator/api/submariner/v1alpha1"
	scheme "github.com/submariner-io/submariner-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceDiscoveriesGetter has a method to return a ServiceDiscoveryInterface.
// A group's client should implement this interface.
type ServiceDiscoveriesGetter interface {
	ServiceDiscoveries(namespace string) ServiceDiscoveryInterface
}

// ServiceDiscoveryInterface has methods to work with ServiceDiscovery resources.
type ServiceDiscoveryInterface interface {
	Create(ctx context.Context, serviceDiscovery *v1alpha1.ServiceDiscovery, opts v1.CreateOptions) (*v1alpha1.ServiceDiscovery, error)
	Update(ctx context.Context, serviceDiscovery *v1alpha1.ServiceDiscovery, opts v1.UpdateOptions) (*v1alpha1.ServiceDiscovery, error)
	UpdateStatus(ctx context.Context, serviceDiscovery *v1alpha1.ServiceDiscovery, opts v1.UpdateOptions) (*v1alpha1.ServiceDiscovery, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceDiscovery, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceDiscoveryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceDiscovery, err error)
	ServiceDiscoveryExpansion
}

// serviceDiscoveries implements ServiceDiscoveryInterface
type serviceDiscoveries struct {
	client rest.Interface
	ns     string
}

// newServiceDiscoveries returns a ServiceDiscoveries
func newServiceDiscoveries(c *SubmarinerV1alpha1Client, namespace string) *serviceDiscoveries {
	return &serviceDiscoveries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceDiscovery, and returns the corresponding serviceDiscovery object, and an error if there is any.
func (c *serviceDiscoveries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceDiscovery, err error) {
	result = &v1alpha1.ServiceDiscovery{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceDiscoveries that match those selectors.
func (c *serviceDiscoveries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceDiscoveryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceDiscoveryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceDiscoveries.
func (c *serviceDiscoveries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceDiscovery and creates it.  Returns the server's representation of the serviceDiscovery, and an error, if there is any.
func (c *serviceDiscoveries) Create(ctx context.Context, serviceDiscovery *v1alpha1.ServiceDiscovery, opts v1.CreateOptions) (result *v1alpha1.ServiceDiscovery, err error) {
	result = &v1alpha1.ServiceDiscovery{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicediscoveries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceDiscovery).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceDiscovery and updates it. Returns the server's representation of the serviceDiscovery, and an error, if there is any.
func (c *serviceDiscoveries) Update(ctx context.Context, serviceDiscovery *v1alpha1.ServiceDiscovery, opts v1.UpdateOptions) (result *v1alpha1.ServiceDiscovery, err error) {
	result = &v1alpha1.ServiceDiscovery{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicediscoveries").
		Name(serviceDiscovery.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceDiscovery).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceDiscoveries) UpdateStatus(ctx context.Context, serviceDiscovery *v1alpha1.ServiceDiscovery, opts v1.UpdateOptions) (result *v1alpha1.ServiceDiscovery, err error) {
	result = &v1alpha1.ServiceDiscovery{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicediscoveries").
		Name(serviceDiscovery.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceDiscovery).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceDiscovery and deletes it. Returns an error if one occurs.
func (c *serviceDiscoveries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicediscoveries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceDiscoveries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicediscoveries").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceDiscovery.
func (c *serviceDiscoveries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceDiscovery, err error) {
	result = &v1alpha1.ServiceDiscovery{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicediscoveries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
