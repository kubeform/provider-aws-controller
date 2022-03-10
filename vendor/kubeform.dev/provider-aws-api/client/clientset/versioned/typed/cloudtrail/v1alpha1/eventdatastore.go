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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudtrail/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EventDataStoresGetter has a method to return a EventDataStoreInterface.
// A group's client should implement this interface.
type EventDataStoresGetter interface {
	EventDataStores(namespace string) EventDataStoreInterface
}

// EventDataStoreInterface has methods to work with EventDataStore resources.
type EventDataStoreInterface interface {
	Create(ctx context.Context, eventDataStore *v1alpha1.EventDataStore, opts v1.CreateOptions) (*v1alpha1.EventDataStore, error)
	Update(ctx context.Context, eventDataStore *v1alpha1.EventDataStore, opts v1.UpdateOptions) (*v1alpha1.EventDataStore, error)
	UpdateStatus(ctx context.Context, eventDataStore *v1alpha1.EventDataStore, opts v1.UpdateOptions) (*v1alpha1.EventDataStore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EventDataStore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EventDataStoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventDataStore, err error)
	EventDataStoreExpansion
}

// eventDataStores implements EventDataStoreInterface
type eventDataStores struct {
	client rest.Interface
	ns     string
}

// newEventDataStores returns a EventDataStores
func newEventDataStores(c *CloudtrailV1alpha1Client, namespace string) *eventDataStores {
	return &eventDataStores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventDataStore, and returns the corresponding eventDataStore object, and an error if there is any.
func (c *eventDataStores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventDataStore, err error) {
	result = &v1alpha1.EventDataStore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventdatastores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventDataStores that match those selectors.
func (c *eventDataStores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventDataStoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EventDataStoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventdatastores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventDataStores.
func (c *eventDataStores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventdatastores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a eventDataStore and creates it.  Returns the server's representation of the eventDataStore, and an error, if there is any.
func (c *eventDataStores) Create(ctx context.Context, eventDataStore *v1alpha1.EventDataStore, opts v1.CreateOptions) (result *v1alpha1.EventDataStore, err error) {
	result = &v1alpha1.EventDataStore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventdatastores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventDataStore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a eventDataStore and updates it. Returns the server's representation of the eventDataStore, and an error, if there is any.
func (c *eventDataStores) Update(ctx context.Context, eventDataStore *v1alpha1.EventDataStore, opts v1.UpdateOptions) (result *v1alpha1.EventDataStore, err error) {
	result = &v1alpha1.EventDataStore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventdatastores").
		Name(eventDataStore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventDataStore).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *eventDataStores) UpdateStatus(ctx context.Context, eventDataStore *v1alpha1.EventDataStore, opts v1.UpdateOptions) (result *v1alpha1.EventDataStore, err error) {
	result = &v1alpha1.EventDataStore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventdatastores").
		Name(eventDataStore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventDataStore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the eventDataStore and deletes it. Returns an error if one occurs.
func (c *eventDataStores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventdatastores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventDataStores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventdatastores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched eventDataStore.
func (c *eventDataStores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventDataStore, err error) {
	result = &v1alpha1.EventDataStore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventdatastores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
