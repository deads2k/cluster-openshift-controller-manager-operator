// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	operatorv1 "github.com/openshift/api/operator/v1"
	applyconfigurationsoperatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeAPIServers implements KubeAPIServerInterface
type FakeKubeAPIServers struct {
	Fake *FakeOperatorV1
}

var kubeapiserversResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "kubeapiservers"}

var kubeapiserversKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "KubeAPIServer"}

// Get takes name of the kubeAPIServer, and returns the corresponding kubeAPIServer object, and an error if there is any.
func (c *FakeKubeAPIServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.KubeAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubeapiserversResource, name), &operatorv1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeAPIServer), err
}

// List takes label and field selectors, and returns the list of KubeAPIServers that match those selectors.
func (c *FakeKubeAPIServers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.KubeAPIServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubeapiserversResource, kubeapiserversKind, opts), &operatorv1.KubeAPIServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.KubeAPIServerList{ListMeta: obj.(*operatorv1.KubeAPIServerList).ListMeta}
	for _, item := range obj.(*operatorv1.KubeAPIServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeAPIServers.
func (c *FakeKubeAPIServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubeapiserversResource, opts))
}

// Create takes the representation of a kubeAPIServer and creates it.  Returns the server's representation of the kubeAPIServer, and an error, if there is any.
func (c *FakeKubeAPIServers) Create(ctx context.Context, kubeAPIServer *operatorv1.KubeAPIServer, opts v1.CreateOptions) (result *operatorv1.KubeAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubeapiserversResource, kubeAPIServer), &operatorv1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeAPIServer), err
}

// Update takes the representation of a kubeAPIServer and updates it. Returns the server's representation of the kubeAPIServer, and an error, if there is any.
func (c *FakeKubeAPIServers) Update(ctx context.Context, kubeAPIServer *operatorv1.KubeAPIServer, opts v1.UpdateOptions) (result *operatorv1.KubeAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubeapiserversResource, kubeAPIServer), &operatorv1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeAPIServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeAPIServers) UpdateStatus(ctx context.Context, kubeAPIServer *operatorv1.KubeAPIServer, opts v1.UpdateOptions) (*operatorv1.KubeAPIServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubeapiserversResource, "status", kubeAPIServer), &operatorv1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeAPIServer), err
}

// Delete takes name of the kubeAPIServer and deletes it. Returns an error if one occurs.
func (c *FakeKubeAPIServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubeapiserversResource, name, opts), &operatorv1.KubeAPIServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeAPIServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubeapiserversResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.KubeAPIServerList{})
	return err
}

// Patch applies the patch and returns the patched kubeAPIServer.
func (c *FakeKubeAPIServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.KubeAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeapiserversResource, name, pt, data, subresources...), &operatorv1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeAPIServer), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied kubeAPIServer.
func (c *FakeKubeAPIServers) Apply(ctx context.Context, kubeAPIServer *applyconfigurationsoperatorv1.KubeAPIServerApplyConfiguration, opts v1.ApplyOptions) (result *operatorv1.KubeAPIServer, err error) {
	if kubeAPIServer == nil {
		return nil, fmt.Errorf("kubeAPIServer provided to Apply must not be nil")
	}
	data, err := json.Marshal(kubeAPIServer)
	if err != nil {
		return nil, err
	}
	name := kubeAPIServer.Name
	if name == nil {
		return nil, fmt.Errorf("kubeAPIServer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeapiserversResource, *name, types.ApplyPatchType, data), &operatorv1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeAPIServer), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeKubeAPIServers) ApplyStatus(ctx context.Context, kubeAPIServer *applyconfigurationsoperatorv1.KubeAPIServerApplyConfiguration, opts v1.ApplyOptions) (result *operatorv1.KubeAPIServer, err error) {
	if kubeAPIServer == nil {
		return nil, fmt.Errorf("kubeAPIServer provided to Apply must not be nil")
	}
	data, err := json.Marshal(kubeAPIServer)
	if err != nil {
		return nil, err
	}
	name := kubeAPIServer.Name
	if name == nil {
		return nil, fmt.Errorf("kubeAPIServer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeapiserversResource, *name, types.ApplyPatchType, data, "status"), &operatorv1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeAPIServer), err
}
