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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudfront/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KeyGroupsGetter has a method to return a KeyGroupInterface.
// A group's client should implement this interface.
type KeyGroupsGetter interface {
	KeyGroups(namespace string) KeyGroupInterface
}

// KeyGroupInterface has methods to work with KeyGroup resources.
type KeyGroupInterface interface {
	Create(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.CreateOptions) (*v1alpha1.KeyGroup, error)
	Update(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.UpdateOptions) (*v1alpha1.KeyGroup, error)
	UpdateStatus(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.UpdateOptions) (*v1alpha1.KeyGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KeyGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KeyGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeyGroup, err error)
	KeyGroupExpansion
}

// keyGroups implements KeyGroupInterface
type keyGroups struct {
	client rest.Interface
	ns     string
}

// newKeyGroups returns a KeyGroups
func newKeyGroups(c *CloudfrontV1alpha1Client, namespace string) *keyGroups {
	return &keyGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the keyGroup, and returns the corresponding keyGroup object, and an error if there is any.
func (c *keyGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KeyGroup, err error) {
	result = &v1alpha1.KeyGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keygroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KeyGroups that match those selectors.
func (c *keyGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeyGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KeyGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested keyGroups.
func (c *keyGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("keygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a keyGroup and creates it.  Returns the server's representation of the keyGroup, and an error, if there is any.
func (c *keyGroups) Create(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.CreateOptions) (result *v1alpha1.KeyGroup, err error) {
	result = &v1alpha1.KeyGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("keygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keyGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a keyGroup and updates it. Returns the server's representation of the keyGroup, and an error, if there is any.
func (c *keyGroups) Update(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.UpdateOptions) (result *v1alpha1.KeyGroup, err error) {
	result = &v1alpha1.KeyGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keygroups").
		Name(keyGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keyGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *keyGroups) UpdateStatus(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.UpdateOptions) (result *v1alpha1.KeyGroup, err error) {
	result = &v1alpha1.KeyGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keygroups").
		Name(keyGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keyGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the keyGroup and deletes it. Returns an error if one occurs.
func (c *keyGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keygroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *keyGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keygroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched keyGroup.
func (c *keyGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeyGroup, err error) {
	result = &v1alpha1.KeyGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("keygroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}