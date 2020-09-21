/*
Copyright 2020 The Knative Authors

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-rabbitmq/pkg/apis/sources/v1alpha1"
)

// FakeRabbitmqSources implements RabbitmqSourceInterface
type FakeRabbitmqSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var rabbitmqsourcesResource = schema.GroupVersionResource{Group: "sources.knative.dev", Version: "v1alpha1", Resource: "rabbitmqsources"}

var rabbitmqsourcesKind = schema.GroupVersionKind{Group: "sources.knative.dev", Version: "v1alpha1", Kind: "RabbitmqSource"}

// Get takes name of the rabbitmqSource, and returns the corresponding rabbitmqSource object, and an error if there is any.
func (c *FakeRabbitmqSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RabbitmqSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rabbitmqsourcesResource, c.ns, name), &v1alpha1.RabbitmqSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqSource), err
}

// List takes label and field selectors, and returns the list of RabbitmqSources that match those selectors.
func (c *FakeRabbitmqSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RabbitmqSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rabbitmqsourcesResource, rabbitmqsourcesKind, c.ns, opts), &v1alpha1.RabbitmqSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RabbitmqSourceList{ListMeta: obj.(*v1alpha1.RabbitmqSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.RabbitmqSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rabbitmqSources.
func (c *FakeRabbitmqSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rabbitmqsourcesResource, c.ns, opts))

}

// Create takes the representation of a rabbitmqSource and creates it.  Returns the server's representation of the rabbitmqSource, and an error, if there is any.
func (c *FakeRabbitmqSources) Create(ctx context.Context, rabbitmqSource *v1alpha1.RabbitmqSource, opts v1.CreateOptions) (result *v1alpha1.RabbitmqSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rabbitmqsourcesResource, c.ns, rabbitmqSource), &v1alpha1.RabbitmqSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqSource), err
}

// Update takes the representation of a rabbitmqSource and updates it. Returns the server's representation of the rabbitmqSource, and an error, if there is any.
func (c *FakeRabbitmqSources) Update(ctx context.Context, rabbitmqSource *v1alpha1.RabbitmqSource, opts v1.UpdateOptions) (result *v1alpha1.RabbitmqSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rabbitmqsourcesResource, c.ns, rabbitmqSource), &v1alpha1.RabbitmqSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRabbitmqSources) UpdateStatus(ctx context.Context, rabbitmqSource *v1alpha1.RabbitmqSource, opts v1.UpdateOptions) (*v1alpha1.RabbitmqSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rabbitmqsourcesResource, "status", c.ns, rabbitmqSource), &v1alpha1.RabbitmqSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqSource), err
}

// Delete takes name of the rabbitmqSource and deletes it. Returns an error if one occurs.
func (c *FakeRabbitmqSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rabbitmqsourcesResource, c.ns, name), &v1alpha1.RabbitmqSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRabbitmqSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rabbitmqsourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RabbitmqSourceList{})
	return err
}

// Patch applies the patch and returns the patched rabbitmqSource.
func (c *FakeRabbitmqSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RabbitmqSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rabbitmqsourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RabbitmqSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqSource), err
}