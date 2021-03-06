/*
Copyright 2018 The Knative Authors

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

package fake

import (
	v1alpha1 "github.com/knative/observability/pkg/apis/sink/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLogSinks implements LogSinkInterface
type FakeLogSinks struct {
	Fake *FakeObservabilityV1alpha1
	ns   string
}

var logsinksResource = schema.GroupVersionResource{Group: "observability.knative.dev", Version: "v1alpha1", Resource: "logsinks"}

var logsinksKind = schema.GroupVersionKind{Group: "observability.knative.dev", Version: "v1alpha1", Kind: "LogSink"}

// Get takes name of the logSink, and returns the corresponding logSink object, and an error if there is any.
func (c *FakeLogSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.LogSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logsinksResource, c.ns, name), &v1alpha1.LogSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogSink), err
}

// List takes label and field selectors, and returns the list of LogSinks that match those selectors.
func (c *FakeLogSinks) List(opts v1.ListOptions) (result *v1alpha1.LogSinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logsinksResource, logsinksKind, c.ns, opts), &v1alpha1.LogSinkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogSinkList{ListMeta: obj.(*v1alpha1.LogSinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogSinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logSinks.
func (c *FakeLogSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logsinksResource, c.ns, opts))

}

// Create takes the representation of a logSink and creates it.  Returns the server's representation of the logSink, and an error, if there is any.
func (c *FakeLogSinks) Create(logSink *v1alpha1.LogSink) (result *v1alpha1.LogSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logsinksResource, c.ns, logSink), &v1alpha1.LogSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogSink), err
}

// Update takes the representation of a logSink and updates it. Returns the server's representation of the logSink, and an error, if there is any.
func (c *FakeLogSinks) Update(logSink *v1alpha1.LogSink) (result *v1alpha1.LogSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logsinksResource, c.ns, logSink), &v1alpha1.LogSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogSink), err
}

// Delete takes name of the logSink and deletes it. Returns an error if one occurs.
func (c *FakeLogSinks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logsinksResource, c.ns, name), &v1alpha1.LogSink{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logsinksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogSinkList{})
	return err
}

// Patch applies the patch and returns the patched logSink.
func (c *FakeLogSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logsinksResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogSink), err
}
