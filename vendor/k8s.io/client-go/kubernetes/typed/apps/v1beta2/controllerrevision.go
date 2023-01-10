/*
Copyright The Kubernetes Authors.

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

package v1beta2

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	logicalcluster "github.com/kcp-dev/logicalcluster/v2"
	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	appsv1beta2 "k8s.io/client-go/applyconfigurations/apps/v1beta2"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// ControllerRevisionsGetter has a method to return a ControllerRevisionInterface.
// A group's client should implement this interface.
type ControllerRevisionsGetter interface {
	ControllerRevisions(namespace string) ControllerRevisionInterface
}

// ControllerRevisionInterface has methods to work with ControllerRevision resources.
type ControllerRevisionInterface interface {
	Create(ctx context.Context, controllerRevision *v1beta2.ControllerRevision, opts v1.CreateOptions) (*v1beta2.ControllerRevision, error)
	Update(ctx context.Context, controllerRevision *v1beta2.ControllerRevision, opts v1.UpdateOptions) (*v1beta2.ControllerRevision, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.ControllerRevision, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.ControllerRevisionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.ControllerRevision, err error)
	Apply(ctx context.Context, controllerRevision *appsv1beta2.ControllerRevisionApplyConfiguration, opts v1.ApplyOptions) (result *v1beta2.ControllerRevision, err error)
	ControllerRevisionExpansion
}

// controllerRevisions implements ControllerRevisionInterface
type controllerRevisions struct {
	client  rest.Interface
	cluster logicalcluster.Name
	ns      string
}

// newControllerRevisions returns a ControllerRevisions
func newControllerRevisions(c *AppsV1beta2Client, namespace string) *controllerRevisions {
	return &controllerRevisions{
		client:  c.RESTClient(),
		cluster: c.cluster,
		ns:      namespace,
	}
}

// Get takes name of the controllerRevision, and returns the corresponding controllerRevision object, and an error if there is any.
func (c *controllerRevisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.ControllerRevision, err error) {
	result = &v1beta2.ControllerRevision{}
	err = c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ControllerRevisions that match those selectors.
func (c *controllerRevisions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.ControllerRevisionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.ControllerRevisionList{}
	err = c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested controllerRevisions.
func (c *controllerRevisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a controllerRevision and creates it.  Returns the server's representation of the controllerRevision, and an error, if there is any.
func (c *controllerRevisions) Create(ctx context.Context, controllerRevision *v1beta2.ControllerRevision, opts v1.CreateOptions) (result *v1beta2.ControllerRevision, err error) {
	result = &v1beta2.ControllerRevision{}
	err = c.client.Post().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(controllerRevision).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a controllerRevision and updates it. Returns the server's representation of the controllerRevision, and an error, if there is any.
func (c *controllerRevisions) Update(ctx context.Context, controllerRevision *v1beta2.ControllerRevision, opts v1.UpdateOptions) (result *v1beta2.ControllerRevision, err error) {
	result = &v1beta2.ControllerRevision{}
	err = c.client.Put().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		Name(controllerRevision.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(controllerRevision).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the controllerRevision and deletes it. Returns an error if one occurs.
func (c *controllerRevisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *controllerRevisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched controllerRevision.
func (c *controllerRevisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.ControllerRevision, err error) {
	result = &v1beta2.ControllerRevision{}
	err = c.client.Patch(pt).
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied controllerRevision.
func (c *controllerRevisions) Apply(ctx context.Context, controllerRevision *appsv1beta2.ControllerRevisionApplyConfiguration, opts v1.ApplyOptions) (result *v1beta2.ControllerRevision, err error) {
	if controllerRevision == nil {
		return nil, fmt.Errorf("controllerRevision provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(controllerRevision)
	if err != nil {
		return nil, err
	}
	name := controllerRevision.Name
	if name == nil {
		return nil, fmt.Errorf("controllerRevision.Name must be provided to Apply")
	}
	result = &v1beta2.ControllerRevision{}
	err = c.client.Patch(types.ApplyPatchType).
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("controllerrevisions").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
