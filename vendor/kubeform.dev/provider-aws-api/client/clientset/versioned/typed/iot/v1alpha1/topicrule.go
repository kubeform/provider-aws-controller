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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/iot/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TopicRulesGetter has a method to return a TopicRuleInterface.
// A group's client should implement this interface.
type TopicRulesGetter interface {
	TopicRules(namespace string) TopicRuleInterface
}

// TopicRuleInterface has methods to work with TopicRule resources.
type TopicRuleInterface interface {
	Create(ctx context.Context, topicRule *v1alpha1.TopicRule, opts v1.CreateOptions) (*v1alpha1.TopicRule, error)
	Update(ctx context.Context, topicRule *v1alpha1.TopicRule, opts v1.UpdateOptions) (*v1alpha1.TopicRule, error)
	UpdateStatus(ctx context.Context, topicRule *v1alpha1.TopicRule, opts v1.UpdateOptions) (*v1alpha1.TopicRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TopicRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TopicRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TopicRule, err error)
	TopicRuleExpansion
}

// topicRules implements TopicRuleInterface
type topicRules struct {
	client rest.Interface
	ns     string
}

// newTopicRules returns a TopicRules
func newTopicRules(c *IotV1alpha1Client, namespace string) *topicRules {
	return &topicRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the topicRule, and returns the corresponding topicRule object, and an error if there is any.
func (c *topicRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TopicRule, err error) {
	result = &v1alpha1.TopicRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("topicrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TopicRules that match those selectors.
func (c *topicRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TopicRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TopicRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("topicrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested topicRules.
func (c *topicRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("topicrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a topicRule and creates it.  Returns the server's representation of the topicRule, and an error, if there is any.
func (c *topicRules) Create(ctx context.Context, topicRule *v1alpha1.TopicRule, opts v1.CreateOptions) (result *v1alpha1.TopicRule, err error) {
	result = &v1alpha1.TopicRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("topicrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(topicRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a topicRule and updates it. Returns the server's representation of the topicRule, and an error, if there is any.
func (c *topicRules) Update(ctx context.Context, topicRule *v1alpha1.TopicRule, opts v1.UpdateOptions) (result *v1alpha1.TopicRule, err error) {
	result = &v1alpha1.TopicRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("topicrules").
		Name(topicRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(topicRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *topicRules) UpdateStatus(ctx context.Context, topicRule *v1alpha1.TopicRule, opts v1.UpdateOptions) (result *v1alpha1.TopicRule, err error) {
	result = &v1alpha1.TopicRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("topicrules").
		Name(topicRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(topicRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the topicRule and deletes it. Returns an error if one occurs.
func (c *topicRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("topicrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *topicRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("topicrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched topicRule.
func (c *topicRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TopicRule, err error) {
	result = &v1alpha1.TopicRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("topicrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
