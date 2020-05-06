// +build !ignore_autogenerated

//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certs) DeepCopyInto(out *Certs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certs.
func (in *Certs) DeepCopy() *Certs {
	if in == nil {
		return nil
	}
	out := new(Certs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collectd) DeepCopyInto(out *Collectd) {
	*out = *in
	in.RouterResource.DeepCopyInto(&out.RouterResource)
	in.Resource.DeepCopyInto(&out.Resource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collectd.
func (in *Collectd) DeepCopy() *Collectd {
	if in == nil {
		return nil
	}
	out := new(Collectd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Exporter) DeepCopyInto(out *Exporter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Exporter.
func (in *Exporter) DeepCopy() *Exporter {
	if in == nil {
		return nil
	}
	out := new(Exporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Exporter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterList) DeepCopyInto(out *ExporterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Exporter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterList.
func (in *ExporterList) DeepCopy() *ExporterList {
	if in == nil {
		return nil
	}
	out := new(ExporterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExporterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterSpec) DeepCopyInto(out *ExporterSpec) {
	*out = *in
	out.Certs = in.Certs
	in.Collectd.DeepCopyInto(&out.Collectd)
	in.NodeExporter.DeepCopyInto(&out.NodeExporter)
	in.KubeStateMetrics.DeepCopyInto(&out.KubeStateMetrics)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterSpec.
func (in *ExporterSpec) DeepCopy() *ExporterSpec {
	if in == nil {
		return nil
	}
	out := new(ExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterStatus) DeepCopyInto(out *ExporterStatus) {
	*out = *in
	in.Collectd.DeepCopyInto(&out.Collectd)
	in.NodeExporter.DeepCopyInto(&out.NodeExporter)
	in.KubeState.DeepCopyInto(&out.KubeState)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterStatus.
func (in *ExporterStatus) DeepCopy() *ExporterStatus {
	if in == nil {
		return nil
	}
	out := new(ExporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeStateMetrics) DeepCopyInto(out *KubeStateMetrics) {
	*out = *in
	in.RouterResource.DeepCopyInto(&out.RouterResource)
	in.Resource.DeepCopyInto(&out.Resource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeStateMetrics.
func (in *KubeStateMetrics) DeepCopy() *KubeStateMetrics {
	if in == nil {
		return nil
	}
	out := new(KubeStateMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeExporter) DeepCopyInto(out *NodeExporter) {
	*out = *in
	in.RouterResource.DeepCopyInto(&out.RouterResource)
	in.Resource.DeepCopyInto(&out.Resource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeExporter.
func (in *NodeExporter) DeepCopy() *NodeExporter {
	if in == nil {
		return nil
	}
	out := new(NodeExporter)
	in.DeepCopyInto(out)
	return out
}
