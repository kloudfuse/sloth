// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	slothv1 "github.com/kloudfuse/sloth/pkg/kubernetes/api/sloth/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrometheusServiceLevels implements PrometheusServiceLevelInterface
type FakePrometheusServiceLevels struct {
	Fake *FakeSlothV1
	ns   string
}

var prometheusservicelevelsResource = schema.GroupVersionResource{Group: "sloth.slok.dev", Version: "v1", Resource: "prometheusservicelevels"}

var prometheusservicelevelsKind = schema.GroupVersionKind{Group: "sloth.slok.dev", Version: "v1", Kind: "PrometheusServiceLevel"}

// Get takes name of the prometheusServiceLevel, and returns the corresponding prometheusServiceLevel object, and an error if there is any.
func (c *FakePrometheusServiceLevels) Get(ctx context.Context, name string, options v1.GetOptions) (result *slothv1.PrometheusServiceLevel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(prometheusservicelevelsResource, c.ns, name), &slothv1.PrometheusServiceLevel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slothv1.PrometheusServiceLevel), err
}

// List takes label and field selectors, and returns the list of PrometheusServiceLevels that match those selectors.
func (c *FakePrometheusServiceLevels) List(ctx context.Context, opts v1.ListOptions) (result *slothv1.PrometheusServiceLevelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(prometheusservicelevelsResource, prometheusservicelevelsKind, c.ns, opts), &slothv1.PrometheusServiceLevelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &slothv1.PrometheusServiceLevelList{ListMeta: obj.(*slothv1.PrometheusServiceLevelList).ListMeta}
	for _, item := range obj.(*slothv1.PrometheusServiceLevelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prometheusServiceLevels.
func (c *FakePrometheusServiceLevels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(prometheusservicelevelsResource, c.ns, opts))

}

// Create takes the representation of a prometheusServiceLevel and creates it.  Returns the server's representation of the prometheusServiceLevel, and an error, if there is any.
func (c *FakePrometheusServiceLevels) Create(ctx context.Context, prometheusServiceLevel *slothv1.PrometheusServiceLevel, opts v1.CreateOptions) (result *slothv1.PrometheusServiceLevel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(prometheusservicelevelsResource, c.ns, prometheusServiceLevel), &slothv1.PrometheusServiceLevel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slothv1.PrometheusServiceLevel), err
}

// Update takes the representation of a prometheusServiceLevel and updates it. Returns the server's representation of the prometheusServiceLevel, and an error, if there is any.
func (c *FakePrometheusServiceLevels) Update(ctx context.Context, prometheusServiceLevel *slothv1.PrometheusServiceLevel, opts v1.UpdateOptions) (result *slothv1.PrometheusServiceLevel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(prometheusservicelevelsResource, c.ns, prometheusServiceLevel), &slothv1.PrometheusServiceLevel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slothv1.PrometheusServiceLevel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrometheusServiceLevels) UpdateStatus(ctx context.Context, prometheusServiceLevel *slothv1.PrometheusServiceLevel, opts v1.UpdateOptions) (*slothv1.PrometheusServiceLevel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(prometheusservicelevelsResource, "status", c.ns, prometheusServiceLevel), &slothv1.PrometheusServiceLevel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slothv1.PrometheusServiceLevel), err
}

// Delete takes name of the prometheusServiceLevel and deletes it. Returns an error if one occurs.
func (c *FakePrometheusServiceLevels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(prometheusservicelevelsResource, c.ns, name, opts), &slothv1.PrometheusServiceLevel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrometheusServiceLevels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(prometheusservicelevelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &slothv1.PrometheusServiceLevelList{})
	return err
}

// Patch applies the patch and returns the patched prometheusServiceLevel.
func (c *FakePrometheusServiceLevels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *slothv1.PrometheusServiceLevel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(prometheusservicelevelsResource, c.ns, name, pt, data, subresources...), &slothv1.PrometheusServiceLevel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slothv1.PrometheusServiceLevel), err
}
