//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by deepcopy-gen. DO NOT EDIT.

package rsyslog

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingRule) DeepCopyInto(out *LoggingRule) {
	*out = *in
	if in.ProgramNames != nil {
		in, out := &in.ProgramNames, &out.ProgramNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingRule.
func (in *LoggingRule) DeepCopy() *LoggingRule {
	if in == nil {
		return nil
	}
	out := new(LoggingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RsyslogRelpConfig) DeepCopyInto(out *RsyslogRelpConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.LoggingRules != nil {
		in, out := &in.LoggingRules, &out.LoggingRules
		*out = make([]LoggingRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RebindInterval != nil {
		in, out := &in.RebindInterval, &out.RebindInterval
		*out = new(int)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ResumeRetryCount != nil {
		in, out := &in.ResumeRetryCount, &out.ResumeRetryCount
		*out = new(int)
		**out = **in
	}
	if in.ReportSuspensionContinuation != nil {
		in, out := &in.ReportSuspensionContinuation, &out.ReportSuspensionContinuation
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RsyslogRelpConfig.
func (in *RsyslogRelpConfig) DeepCopy() *RsyslogRelpConfig {
	if in == nil {
		return nil
	}
	out := new(RsyslogRelpConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RsyslogRelpConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
	if in.SecretReferenceName != nil {
		in, out := &in.SecretReferenceName, &out.SecretReferenceName
		*out = new(string)
		**out = **in
	}
	if in.PermittedPeer != nil {
		in, out := &in.PermittedPeer, &out.PermittedPeer
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AuthMode != nil {
		in, out := &in.AuthMode, &out.AuthMode
		*out = new(AuthMode)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}
