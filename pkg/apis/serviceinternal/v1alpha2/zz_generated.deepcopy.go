// +build !ignore_autogenerated

/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationReference) DeepCopyInto(out *ApplicationReference) {
	*out = *in
	in.Reference.DeepCopyInto(&out.Reference)
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]intstr.IntOrString, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationReference.
func (in *ApplicationReference) DeepCopy() *ApplicationReference {
	if in == nil {
		return nil
	}
	out := new(ApplicationReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingProjection) DeepCopyInto(out *ServiceBindingProjection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingProjection.
func (in *ServiceBindingProjection) DeepCopy() *ServiceBindingProjection {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingProjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingProjection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingProjectionList) DeepCopyInto(out *ServiceBindingProjectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBindingProjection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingProjectionList.
func (in *ServiceBindingProjectionList) DeepCopy() *ServiceBindingProjectionList {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingProjectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingProjectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingProjectionSpec) DeepCopyInto(out *ServiceBindingProjectionSpec) {
	*out = *in
	out.Binding = in.Binding
	in.Application.DeepCopyInto(&out.Application)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingProjectionSpec.
func (in *ServiceBindingProjectionSpec) DeepCopy() *ServiceBindingProjectionSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingProjectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingProjectionStatus) DeepCopyInto(out *ServiceBindingProjectionStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingProjectionStatus.
func (in *ServiceBindingProjectionStatus) DeepCopy() *ServiceBindingProjectionStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingProjectionStatus)
	in.DeepCopyInto(out)
	return out
}
