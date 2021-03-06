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

// FakeMetricSinks implements MetricSinkInterface
type FakeMetricSinks struct {
	Fake *FakeObservabilityV1alpha1
	ns   string
}

var metricsinksResource = schema.GroupVersionResource{Group: "observability.knative.dev", Version: "v1alpha1", Resource: "metricsinks"}

var metricsinksKind = schema.GroupVersionKind{Group: "observability.knative.dev", Version: "v1alpha1", Kind: "MetricSink"}

// Get takes name of the metricSink, and returns the corresponding metricSink object, and an error if there is any.
func (c *FakeMetricSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.MetricSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(metricsinksResource, c.ns, name), &v1alpha1.MetricSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricSink), err
}

// List takes label and field selectors, and returns the list of MetricSinks that match those selectors.
func (c *FakeMetricSinks) List(opts v1.ListOptions) (result *v1alpha1.MetricSinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(metricsinksResource, metricsinksKind, c.ns, opts), &v1alpha1.MetricSinkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MetricSinkList{ListMeta: obj.(*v1alpha1.MetricSinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.MetricSinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested metricSinks.
func (c *FakeMetricSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(metricsinksResource, c.ns, opts))

}

// Create takes the representation of a metricSink and creates it.  Returns the server's representation of the metricSink, and an error, if there is any.
func (c *FakeMetricSinks) Create(metricSink *v1alpha1.MetricSink) (result *v1alpha1.MetricSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(metricsinksResource, c.ns, metricSink), &v1alpha1.MetricSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricSink), err
}

// Update takes the representation of a metricSink and updates it. Returns the server's representation of the metricSink, and an error, if there is any.
func (c *FakeMetricSinks) Update(metricSink *v1alpha1.MetricSink) (result *v1alpha1.MetricSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(metricsinksResource, c.ns, metricSink), &v1alpha1.MetricSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricSink), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMetricSinks) UpdateStatus(metricSink *v1alpha1.MetricSink) (*v1alpha1.MetricSink, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(metricsinksResource, "status", c.ns, metricSink), &v1alpha1.MetricSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricSink), err
}

// Delete takes name of the metricSink and deletes it. Returns an error if one occurs.
func (c *FakeMetricSinks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(metricsinksResource, c.ns, name), &v1alpha1.MetricSink{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMetricSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(metricsinksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MetricSinkList{})
	return err
}

// Patch applies the patch and returns the patched metricSink.
func (c *FakeMetricSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MetricSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(metricsinksResource, c.ns, name, pt, data, subresources...), &v1alpha1.MetricSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricSink), err
}
