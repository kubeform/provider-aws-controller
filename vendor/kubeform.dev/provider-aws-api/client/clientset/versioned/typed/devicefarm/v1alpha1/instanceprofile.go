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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/devicefarm/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InstanceProfilesGetter has a method to return a InstanceProfileInterface.
// A group's client should implement this interface.
type InstanceProfilesGetter interface {
	InstanceProfiles(namespace string) InstanceProfileInterface
}

// InstanceProfileInterface has methods to work with InstanceProfile resources.
type InstanceProfileInterface interface {
	Create(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.CreateOptions) (*v1alpha1.InstanceProfile, error)
	Update(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.UpdateOptions) (*v1alpha1.InstanceProfile, error)
	UpdateStatus(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.UpdateOptions) (*v1alpha1.InstanceProfile, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.InstanceProfile, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.InstanceProfileList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstanceProfile, err error)
	InstanceProfileExpansion
}

// instanceProfiles implements InstanceProfileInterface
type instanceProfiles struct {
	client rest.Interface
	ns     string
}

// newInstanceProfiles returns a InstanceProfiles
func newInstanceProfiles(c *DevicefarmV1alpha1Client, namespace string) *instanceProfiles {
	return &instanceProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the instanceProfile, and returns the corresponding instanceProfile object, and an error if there is any.
func (c *instanceProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstanceProfile, err error) {
	result = &v1alpha1.InstanceProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("instanceprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InstanceProfiles that match those selectors.
func (c *instanceProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstanceProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.InstanceProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("instanceprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested instanceProfiles.
func (c *instanceProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("instanceprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a instanceProfile and creates it.  Returns the server's representation of the instanceProfile, and an error, if there is any.
func (c *instanceProfiles) Create(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.CreateOptions) (result *v1alpha1.InstanceProfile, err error) {
	result = &v1alpha1.InstanceProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("instanceprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(instanceProfile).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a instanceProfile and updates it. Returns the server's representation of the instanceProfile, and an error, if there is any.
func (c *instanceProfiles) Update(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.UpdateOptions) (result *v1alpha1.InstanceProfile, err error) {
	result = &v1alpha1.InstanceProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("instanceprofiles").
		Name(instanceProfile.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(instanceProfile).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *instanceProfiles) UpdateStatus(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.UpdateOptions) (result *v1alpha1.InstanceProfile, err error) {
	result = &v1alpha1.InstanceProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("instanceprofiles").
		Name(instanceProfile.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(instanceProfile).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the instanceProfile and deletes it. Returns an error if one occurs.
func (c *instanceProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("instanceprofiles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *instanceProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("instanceprofiles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched instanceProfile.
func (c *instanceProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstanceProfile, err error) {
	result = &v1alpha1.InstanceProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("instanceprofiles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
