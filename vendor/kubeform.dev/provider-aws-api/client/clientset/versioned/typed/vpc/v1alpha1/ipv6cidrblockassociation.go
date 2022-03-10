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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/vpc/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// Ipv6CIDRBlockAssociationsGetter has a method to return a Ipv6CIDRBlockAssociationInterface.
// A group's client should implement this interface.
type Ipv6CIDRBlockAssociationsGetter interface {
	Ipv6CIDRBlockAssociations(namespace string) Ipv6CIDRBlockAssociationInterface
}

// Ipv6CIDRBlockAssociationInterface has methods to work with Ipv6CIDRBlockAssociation resources.
type Ipv6CIDRBlockAssociationInterface interface {
	Create(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.CreateOptions) (*v1alpha1.Ipv6CIDRBlockAssociation, error)
	Update(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.UpdateOptions) (*v1alpha1.Ipv6CIDRBlockAssociation, error)
	UpdateStatus(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.UpdateOptions) (*v1alpha1.Ipv6CIDRBlockAssociation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Ipv6CIDRBlockAssociation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.Ipv6CIDRBlockAssociationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error)
	Ipv6CIDRBlockAssociationExpansion
}

// ipv6CIDRBlockAssociations implements Ipv6CIDRBlockAssociationInterface
type ipv6CIDRBlockAssociations struct {
	client rest.Interface
	ns     string
}

// newIpv6CIDRBlockAssociations returns a Ipv6CIDRBlockAssociations
func newIpv6CIDRBlockAssociations(c *VpcV1alpha1Client, namespace string) *ipv6CIDRBlockAssociations {
	return &ipv6CIDRBlockAssociations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ipv6CIDRBlockAssociation, and returns the corresponding ipv6CIDRBlockAssociation object, and an error if there is any.
func (c *ipv6CIDRBlockAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	result = &v1alpha1.Ipv6CIDRBlockAssociation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ipv6CIDRBlockAssociations that match those selectors.
func (c *ipv6CIDRBlockAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Ipv6CIDRBlockAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Ipv6CIDRBlockAssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ipv6CIDRBlockAssociations.
func (c *ipv6CIDRBlockAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ipv6CIDRBlockAssociation and creates it.  Returns the server's representation of the ipv6CIDRBlockAssociation, and an error, if there is any.
func (c *ipv6CIDRBlockAssociations) Create(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.CreateOptions) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	result = &v1alpha1.Ipv6CIDRBlockAssociation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ipv6CIDRBlockAssociation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ipv6CIDRBlockAssociation and updates it. Returns the server's representation of the ipv6CIDRBlockAssociation, and an error, if there is any.
func (c *ipv6CIDRBlockAssociations) Update(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.UpdateOptions) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	result = &v1alpha1.Ipv6CIDRBlockAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		Name(ipv6CIDRBlockAssociation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ipv6CIDRBlockAssociation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *ipv6CIDRBlockAssociations) UpdateStatus(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.UpdateOptions) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	result = &v1alpha1.Ipv6CIDRBlockAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		Name(ipv6CIDRBlockAssociation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ipv6CIDRBlockAssociation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ipv6CIDRBlockAssociation and deletes it. Returns an error if one occurs.
func (c *ipv6CIDRBlockAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ipv6CIDRBlockAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ipv6CIDRBlockAssociation.
func (c *ipv6CIDRBlockAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	result = &v1alpha1.Ipv6CIDRBlockAssociation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ipv6cidrblockassociations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}