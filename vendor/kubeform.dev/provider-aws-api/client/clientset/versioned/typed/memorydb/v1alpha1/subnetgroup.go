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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/memorydb/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SubnetGroupsGetter has a method to return a SubnetGroupInterface.
// A group's client should implement this interface.
type SubnetGroupsGetter interface {
	SubnetGroups(namespace string) SubnetGroupInterface
}

// SubnetGroupInterface has methods to work with SubnetGroup resources.
type SubnetGroupInterface interface {
	Create(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.CreateOptions) (*v1alpha1.SubnetGroup, error)
	Update(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.UpdateOptions) (*v1alpha1.SubnetGroup, error)
	UpdateStatus(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.UpdateOptions) (*v1alpha1.SubnetGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SubnetGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SubnetGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SubnetGroup, err error)
	SubnetGroupExpansion
}

// subnetGroups implements SubnetGroupInterface
type subnetGroups struct {
	client rest.Interface
	ns     string
}

// newSubnetGroups returns a SubnetGroups
func newSubnetGroups(c *MemorydbV1alpha1Client, namespace string) *subnetGroups {
	return &subnetGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the subnetGroup, and returns the corresponding subnetGroup object, and an error if there is any.
func (c *subnetGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SubnetGroup, err error) {
	result = &v1alpha1.SubnetGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SubnetGroups that match those selectors.
func (c *subnetGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubnetGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SubnetGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested subnetGroups.
func (c *subnetGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("subnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a subnetGroup and creates it.  Returns the server's representation of the subnetGroup, and an error, if there is any.
func (c *subnetGroups) Create(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.CreateOptions) (result *v1alpha1.SubnetGroup, err error) {
	result = &v1alpha1.SubnetGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("subnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnetGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a subnetGroup and updates it. Returns the server's representation of the subnetGroup, and an error, if there is any.
func (c *subnetGroups) Update(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.UpdateOptions) (result *v1alpha1.SubnetGroup, err error) {
	result = &v1alpha1.SubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnetgroups").
		Name(subnetGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnetGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *subnetGroups) UpdateStatus(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.UpdateOptions) (result *v1alpha1.SubnetGroup, err error) {
	result = &v1alpha1.SubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnetgroups").
		Name(subnetGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnetGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the subnetGroup and deletes it. Returns an error if one occurs.
func (c *subnetGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnetgroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *subnetGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnetgroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched subnetGroup.
func (c *subnetGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SubnetGroup, err error) {
	result = &v1alpha1.SubnetGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("subnetgroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
