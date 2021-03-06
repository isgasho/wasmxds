// +build !ignore_autogenerated

// Copyright 2020 Tetrate
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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasmExtension) DeepCopyInto(out *WasmExtension) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmExtension.
func (in *WasmExtension) DeepCopy() *WasmExtension {
	if in == nil {
		return nil
	}
	out := new(WasmExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WasmExtension) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasmExtensionConfigValue) DeepCopyInto(out *WasmExtensionConfigValue) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(WasmExtensionConfigValueRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmExtensionConfigValue.
func (in *WasmExtensionConfigValue) DeepCopy() *WasmExtensionConfigValue {
	if in == nil {
		return nil
	}
	out := new(WasmExtensionConfigValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasmExtensionConfigValueRef) DeepCopyInto(out *WasmExtensionConfigValueRef) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(WasmExtensionConfigValueRefAttribute)
		**out = **in
	}
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(WasmExtensionConfigValueRefAttribute)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmExtensionConfigValueRef.
func (in *WasmExtensionConfigValueRef) DeepCopy() *WasmExtensionConfigValueRef {
	if in == nil {
		return nil
	}
	out := new(WasmExtensionConfigValueRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasmExtensionConfigValueRefAttribute) DeepCopyInto(out *WasmExtensionConfigValueRefAttribute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmExtensionConfigValueRefAttribute.
func (in *WasmExtensionConfigValueRefAttribute) DeepCopy() *WasmExtensionConfigValueRefAttribute {
	if in == nil {
		return nil
	}
	out := new(WasmExtensionConfigValueRefAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasmExtensionList) DeepCopyInto(out *WasmExtensionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WasmExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmExtensionList.
func (in *WasmExtensionList) DeepCopy() *WasmExtensionList {
	if in == nil {
		return nil
	}
	out := new(WasmExtensionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WasmExtensionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasmExtensionSpec) DeepCopyInto(out *WasmExtensionSpec) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
	if in.VMConfiguration != nil {
		in, out := &in.VMConfiguration, &out.VMConfiguration
		*out = new(WasmExtensionConfigValue)
		(*in).DeepCopyInto(*out)
	}
	if in.PluginConfiguration != nil {
		in, out := &in.PluginConfiguration, &out.PluginConfiguration
		*out = new(WasmExtensionConfigValue)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmExtensionSpec.
func (in *WasmExtensionSpec) DeepCopy() *WasmExtensionSpec {
	if in == nil {
		return nil
	}
	out := new(WasmExtensionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasmExtensionSpecImage) DeepCopyInto(out *WasmExtensionSpecImage) {
	*out = *in
	if in.Sha256 != nil {
		in, out := &in.Sha256, &out.Sha256
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmExtensionSpecImage.
func (in *WasmExtensionSpecImage) DeepCopy() *WasmExtensionSpecImage {
	if in == nil {
		return nil
	}
	out := new(WasmExtensionSpecImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WasmExtensionStatus) DeepCopyInto(out *WasmExtensionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmExtensionStatus.
func (in *WasmExtensionStatus) DeepCopy() *WasmExtensionStatus {
	if in == nil {
		return nil
	}
	out := new(WasmExtensionStatus)
	in.DeepCopyInto(out)
	return out
}
