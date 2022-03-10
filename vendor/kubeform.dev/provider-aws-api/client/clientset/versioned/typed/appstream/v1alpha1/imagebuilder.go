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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/appstream/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImageBuildersGetter has a method to return a ImageBuilderInterface.
// A group's client should implement this interface.
type ImageBuildersGetter interface {
	ImageBuilders(namespace string) ImageBuilderInterface
}

// ImageBuilderInterface has methods to work with ImageBuilder resources.
type ImageBuilderInterface interface {
	Create(ctx context.Context, imageBuilder *v1alpha1.ImageBuilder, opts v1.CreateOptions) (*v1alpha1.ImageBuilder, error)
	Update(ctx context.Context, imageBuilder *v1alpha1.ImageBuilder, opts v1.UpdateOptions) (*v1alpha1.ImageBuilder, error)
	UpdateStatus(ctx context.Context, imageBuilder *v1alpha1.ImageBuilder, opts v1.UpdateOptions) (*v1alpha1.ImageBuilder, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ImageBuilder, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ImageBuilderList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageBuilder, err error)
	ImageBuilderExpansion
}

// imageBuilders implements ImageBuilderInterface
type imageBuilders struct {
	client rest.Interface
	ns     string
}

// newImageBuilders returns a ImageBuilders
func newImageBuilders(c *AppstreamV1alpha1Client, namespace string) *imageBuilders {
	return &imageBuilders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the imageBuilder, and returns the corresponding imageBuilder object, and an error if there is any.
func (c *imageBuilders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImageBuilder, err error) {
	result = &v1alpha1.ImageBuilder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagebuilders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ImageBuilders that match those selectors.
func (c *imageBuilders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImageBuilderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ImageBuilderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagebuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested imageBuilders.
func (c *imageBuilders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("imagebuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a imageBuilder and creates it.  Returns the server's representation of the imageBuilder, and an error, if there is any.
func (c *imageBuilders) Create(ctx context.Context, imageBuilder *v1alpha1.ImageBuilder, opts v1.CreateOptions) (result *v1alpha1.ImageBuilder, err error) {
	result = &v1alpha1.ImageBuilder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("imagebuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imageBuilder).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a imageBuilder and updates it. Returns the server's representation of the imageBuilder, and an error, if there is any.
func (c *imageBuilders) Update(ctx context.Context, imageBuilder *v1alpha1.ImageBuilder, opts v1.UpdateOptions) (result *v1alpha1.ImageBuilder, err error) {
	result = &v1alpha1.ImageBuilder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagebuilders").
		Name(imageBuilder.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imageBuilder).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *imageBuilders) UpdateStatus(ctx context.Context, imageBuilder *v1alpha1.ImageBuilder, opts v1.UpdateOptions) (result *v1alpha1.ImageBuilder, err error) {
	result = &v1alpha1.ImageBuilder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagebuilders").
		Name(imageBuilder.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imageBuilder).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the imageBuilder and deletes it. Returns an error if one occurs.
func (c *imageBuilders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagebuilders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *imageBuilders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagebuilders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched imageBuilder.
func (c *imageBuilders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageBuilder, err error) {
	result = &v1alpha1.ImageBuilder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("imagebuilders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}