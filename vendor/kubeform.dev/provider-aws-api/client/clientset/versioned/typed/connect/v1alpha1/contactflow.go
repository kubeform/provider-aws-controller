/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/connect/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ContactFlowsGetter has a method to return a ContactFlowInterface.
// A group's client should implement this interface.
type ContactFlowsGetter interface {
	ContactFlows(namespace string) ContactFlowInterface
}

// ContactFlowInterface has methods to work with ContactFlow resources.
type ContactFlowInterface interface {
	Create(ctx context.Context, contactFlow *v1alpha1.ContactFlow, opts v1.CreateOptions) (*v1alpha1.ContactFlow, error)
	Update(ctx context.Context, contactFlow *v1alpha1.ContactFlow, opts v1.UpdateOptions) (*v1alpha1.ContactFlow, error)
	UpdateStatus(ctx context.Context, contactFlow *v1alpha1.ContactFlow, opts v1.UpdateOptions) (*v1alpha1.ContactFlow, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ContactFlow, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ContactFlowList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ContactFlow, err error)
	ContactFlowExpansion
}

// contactFlows implements ContactFlowInterface
type contactFlows struct {
	client rest.Interface
	ns     string
}

// newContactFlows returns a ContactFlows
func newContactFlows(c *ConnectV1alpha1Client, namespace string) *contactFlows {
	return &contactFlows{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the contactFlow, and returns the corresponding contactFlow object, and an error if there is any.
func (c *contactFlows) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ContactFlow, err error) {
	result = &v1alpha1.ContactFlow{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("contactflows").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ContactFlows that match those selectors.
func (c *contactFlows) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ContactFlowList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ContactFlowList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("contactflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested contactFlows.
func (c *contactFlows) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("contactflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a contactFlow and creates it.  Returns the server's representation of the contactFlow, and an error, if there is any.
func (c *contactFlows) Create(ctx context.Context, contactFlow *v1alpha1.ContactFlow, opts v1.CreateOptions) (result *v1alpha1.ContactFlow, err error) {
	result = &v1alpha1.ContactFlow{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("contactflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(contactFlow).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a contactFlow and updates it. Returns the server's representation of the contactFlow, and an error, if there is any.
func (c *contactFlows) Update(ctx context.Context, contactFlow *v1alpha1.ContactFlow, opts v1.UpdateOptions) (result *v1alpha1.ContactFlow, err error) {
	result = &v1alpha1.ContactFlow{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("contactflows").
		Name(contactFlow.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(contactFlow).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *contactFlows) UpdateStatus(ctx context.Context, contactFlow *v1alpha1.ContactFlow, opts v1.UpdateOptions) (result *v1alpha1.ContactFlow, err error) {
	result = &v1alpha1.ContactFlow{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("contactflows").
		Name(contactFlow.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(contactFlow).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the contactFlow and deletes it. Returns an error if one occurs.
func (c *contactFlows) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("contactflows").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *contactFlows) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("contactflows").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched contactFlow.
func (c *contactFlows) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ContactFlow, err error) {
	result = &v1alpha1.ContactFlow{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("contactflows").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}