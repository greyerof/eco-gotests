//go:build !ignore_autogenerated

// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ComplianceMap) DeepCopyInto(out *ComplianceMap) {
	{
		in := &in
		*out = make(ComplianceMap, len(*in))
		for key, val := range *in {
			var outVal *CompliancePerClusterStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(CompliancePerClusterStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceMap.
func (in ComplianceMap) DeepCopy() ComplianceMap {
	if in == nil {
		return nil
	}
	out := new(ComplianceMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompliancePerClusterStatus) DeepCopyInto(out *CompliancePerClusterStatus) {
	*out = *in
	if in.AggregatePolicyStatus != nil {
		in, out := &in.AggregatePolicyStatus, &out.AggregatePolicyStatus
		*out = make(map[string]*ConfigurationPolicyStatus, len(*in))
		for key, val := range *in {
			var outVal *ConfigurationPolicyStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(ConfigurationPolicyStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompliancePerClusterStatus.
func (in *CompliancePerClusterStatus) DeepCopy() *CompliancePerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CompliancePerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationPolicy) DeepCopyInto(out *ConfigurationPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ConfigurationPolicySpec)
		(*in).DeepCopyInto(*out)
	}
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationPolicy.
func (in *ConfigurationPolicy) DeepCopy() *ConfigurationPolicy {
	if in == nil {
		return nil
	}
	out := new(ConfigurationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationPolicyList) DeepCopyInto(out *ConfigurationPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigurationPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationPolicyList.
func (in *ConfigurationPolicyList) DeepCopy() *ConfigurationPolicyList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationPolicySpec) DeepCopyInto(out *ConfigurationPolicySpec) {
	*out = *in
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	if in.ObjectTemplates != nil {
		in, out := &in.ObjectTemplates, &out.ObjectTemplates
		*out = make([]*ObjectTemplate, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ObjectTemplate)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.EvaluationInterval = in.EvaluationInterval
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationPolicySpec.
func (in *ConfigurationPolicySpec) DeepCopy() *ConfigurationPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationPolicyStatus) DeepCopyInto(out *ConfigurationPolicyStatus) {
	*out = *in
	if in.CompliancyDetails != nil {
		in, out := &in.CompliancyDetails, &out.CompliancyDetails
		*out = make([]TemplateStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]RelatedObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationPolicyStatus.
func (in *ConfigurationPolicyStatus) DeepCopy() *ConfigurationPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluationInterval) DeepCopyInto(out *EvaluationInterval) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluationInterval.
func (in *EvaluationInterval) DeepCopy() *EvaluationInterval {
	if in == nil {
		return nil
	}
	out := new(EvaluationInterval)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMetadata) DeepCopyInto(out *ObjectMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMetadata.
func (in *ObjectMetadata) DeepCopy() *ObjectMetadata {
	if in == nil {
		return nil
	}
	out := new(ObjectMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectProperties) DeepCopyInto(out *ObjectProperties) {
	*out = *in
	if in.CreatedByPolicy != nil {
		in, out := &in.CreatedByPolicy, &out.CreatedByPolicy
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectProperties.
func (in *ObjectProperties) DeepCopy() *ObjectProperties {
	if in == nil {
		return nil
	}
	out := new(ObjectProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectResource) DeepCopyInto(out *ObjectResource) {
	*out = *in
	out.Metadata = in.Metadata
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectResource.
func (in *ObjectResource) DeepCopy() *ObjectResource {
	if in == nil {
		return nil
	}
	out := new(ObjectResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTemplate) DeepCopyInto(out *ObjectTemplate) {
	*out = *in
	in.ObjectDefinition.DeepCopyInto(&out.ObjectDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTemplate.
func (in *ObjectTemplate) DeepCopy() *ObjectTemplate {
	if in == nil {
		return nil
	}
	out := new(ObjectTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelatedObject) DeepCopyInto(out *RelatedObject) {
	*out = *in
	out.Object = in.Object
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ObjectProperties)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelatedObject.
func (in *RelatedObject) DeepCopy() *RelatedObject {
	if in == nil {
		return nil
	}
	out := new(RelatedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]NonEmptyString, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]NonEmptyString, len(*in))
		copy(*out, *in)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = new([]metav1.LabelSelectorRequirement)
		if **in != nil {
			in, out := *in, *out
			*out = make([]metav1.LabelSelectorRequirement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStatus) DeepCopyInto(out *TemplateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Validity.DeepCopyInto(&out.Validity)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStatus.
func (in *TemplateStatus) DeepCopy() *TemplateStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Validity) DeepCopyInto(out *Validity) {
	*out = *in
	if in.Valid != nil {
		in, out := &in.Valid, &out.Valid
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Validity.
func (in *Validity) DeepCopy() *Validity {
	if in == nil {
		return nil
	}
	out := new(Validity)
	in.DeepCopyInto(out)
	return out
}
