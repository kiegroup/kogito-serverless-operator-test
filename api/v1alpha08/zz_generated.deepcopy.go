//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha08

import (
	"github.com/serverlessworkflow/sdk-go/v2/model"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildPlatformConfig) DeepCopyInto(out *BuildPlatformConfig) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.BuildStrategyOptions != nil {
		in, out := &in.BuildStrategyOptions, &out.BuildStrategyOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Registry = in.Registry
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildPlatformConfig.
func (in *BuildPlatformConfig) DeepCopy() *BuildPlatformConfig {
	if in == nil {
		return nil
	}
	out := new(BuildPlatformConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildPlatformSpec) DeepCopyInto(out *BuildPlatformSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildPlatformSpec.
func (in *BuildPlatformSpec) DeepCopy() *BuildPlatformSpec {
	if in == nil {
		return nil
	}
	out := new(BuildPlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildTemplate) DeepCopyInto(out *BuildTemplate) {
	*out = *in
	out.Timeout = in.Timeout
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildTemplate.
func (in *BuildTemplate) DeepCopy() *BuildTemplate {
	if in == nil {
		return nil
	}
	out := new(BuildTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapWorkflowResource) DeepCopyInto(out *ConfigMapWorkflowResource) {
	*out = *in
	out.ConfigMap = in.ConfigMap
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapWorkflowResource.
func (in *ConfigMapWorkflowResource) DeepCopy() *ConfigMapWorkflowResource {
	if in == nil {
		return nil
	}
	out := new(ConfigMapWorkflowResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSpec) DeepCopyInto(out *ContainerSpec) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ContainerPort, len(*in))
		copy(*out, *in)
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ResizePolicy != nil {
		in, out := &in.ResizePolicy, &out.ResizePolicy
		*out = make([]v1.ContainerResizePolicy, len(*in))
		copy(*out, *in)
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeDevices != nil {
		in, out := &in.VolumeDevices, &out.VolumeDevices
		*out = make([]v1.VolumeDevice, len(*in))
		copy(*out, *in)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.StartupProbe != nil {
		in, out := &in.StartupProbe, &out.StartupProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Lifecycle != nil {
		in, out := &in.Lifecycle, &out.Lifecycle
		*out = new(v1.Lifecycle)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSpec.
func (in *ContainerSpec) DeepCopy() *ContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevModePlatformSpec) DeepCopyInto(out *DevModePlatformSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevModePlatformSpec.
func (in *DevModePlatformSpec) DeepCopy() *DevModePlatformSpec {
	if in == nil {
		return nil
	}
	out := new(DevModePlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flow) DeepCopyInto(out *Flow) {
	*out = *in
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(model.Start)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DataInputSchema != nil {
		in, out := &in.DataInputSchema, &out.DataInputSchema
		*out = new(model.DataInputSchema)
		**out = **in
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(model.Secrets, len(*in))
		copy(*out, *in)
	}
	if in.Constants != nil {
		in, out := &in.Constants, &out.Constants
		*out = new(model.Constants)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(model.Timeouts)
		(*in).DeepCopyInto(*out)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make(model.Errors, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(model.Metadata, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = make(model.Auths, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.States != nil {
		in, out := &in.States, &out.States
		*out = make([]model.State, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make(model.Events, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Functions != nil {
		in, out := &in.Functions, &out.Functions
		*out = make(model.Functions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = make(model.Retries, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flow.
func (in *Flow) DeepCopy() *Flow {
	if in == nil {
		return nil
	}
	out := new(Flow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceOptions) DeepCopyInto(out *PersistenceOptions) {
	*out = *in
	if in.PostgreSql != nil {
		in, out := &in.PostgreSql, &out.PostgreSql
		*out = new(PersistencePostgreSql)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceOptions.
func (in *PersistenceOptions) DeepCopy() *PersistenceOptions {
	if in == nil {
		return nil
	}
	out := new(PersistenceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistencePostgreSql) DeepCopyInto(out *PersistencePostgreSql) {
	*out = *in
	out.SecretRef = in.SecretRef
	in.ServiceRef.DeepCopyInto(&out.ServiceRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistencePostgreSql.
func (in *PersistencePostgreSql) DeepCopy() *PersistencePostgreSql {
	if in == nil {
		return nil
	}
	out := new(PersistencePostgreSql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpec) DeepCopyInto(out *PodSpec) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AutomountServiceAccountToken != nil {
		in, out := &in.AutomountServiceAccountToken, &out.AutomountServiceAccountToken
		*out = new(bool)
		**out = **in
	}
	if in.ShareProcessNamespace != nil {
		in, out := &in.ShareProcessNamespace, &out.ShareProcessNamespace
		*out = new(bool)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]v1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int32)
		**out = **in
	}
	if in.DNSConfig != nil {
		in, out := &in.DNSConfig, &out.DNSConfig
		*out = new(v1.PodDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessGates != nil {
		in, out := &in.ReadinessGates, &out.ReadinessGates
		*out = make([]v1.PodReadinessGate, len(*in))
		copy(*out, *in)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.EnableServiceLinks != nil {
		in, out := &in.EnableServiceLinks, &out.EnableServiceLinks
		*out = new(bool)
		**out = **in
	}
	if in.PreemptionPolicy != nil {
		in, out := &in.PreemptionPolicy, &out.PreemptionPolicy
		*out = new(v1.PreemptionPolicy)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SetHostnameAsFQDN != nil {
		in, out := &in.SetHostnameAsFQDN, &out.SetHostnameAsFQDN
		*out = new(bool)
		**out = **in
	}
	if in.OS != nil {
		in, out := &in.OS, &out.OS
		*out = new(v1.PodOS)
		**out = **in
	}
	if in.HostUsers != nil {
		in, out := &in.HostUsers, &out.HostUsers
		*out = new(bool)
		**out = **in
	}
	if in.SchedulingGates != nil {
		in, out := &in.SchedulingGates, &out.SchedulingGates
		*out = make([]v1.PodSchedulingGate, len(*in))
		copy(*out, *in)
	}
	if in.ResourceClaims != nil {
		in, out := &in.ResourceClaims, &out.ResourceClaims
		*out = make([]v1.PodResourceClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpec.
func (in *PodSpec) DeepCopy() *PodSpec {
	if in == nil {
		return nil
	}
	out := new(PodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTemplateSpec) DeepCopyInto(out *PodTemplateSpec) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTemplateSpec.
func (in *PodTemplateSpec) DeepCopy() *PodTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(PodTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSqlSecretOptions) DeepCopyInto(out *PostgreSqlSecretOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSqlSecretOptions.
func (in *PostgreSqlSecretOptions) DeepCopy() *PostgreSqlSecretOptions {
	if in == nil {
		return nil
	}
	out := new(PostgreSqlSecretOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSqlServiceOptions) DeepCopyInto(out *PostgreSqlServiceOptions) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSqlServiceOptions.
func (in *PostgreSqlServiceOptions) DeepCopy() *PostgreSqlServiceOptions {
	if in == nil {
		return nil
	}
	out := new(PostgreSqlServiceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpec) DeepCopyInto(out *RegistrySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpec.
func (in *RegistrySpec) DeepCopy() *RegistrySpec {
	if in == nil {
		return nil
	}
	out := new(RegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceOptions)
		(*in).DeepCopyInto(*out)
	}
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesPlatformSpec) DeepCopyInto(out *ServicesPlatformSpec) {
	*out = *in
	if in.DataIndex != nil {
		in, out := &in.DataIndex, &out.DataIndex
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesPlatformSpec.
func (in *ServicesPlatformSpec) DeepCopy() *ServicesPlatformSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesPlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlow) DeepCopyInto(out *SonataFlow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlow.
func (in *SonataFlow) DeepCopy() *SonataFlow {
	if in == nil {
		return nil
	}
	out := new(SonataFlow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SonataFlow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowBuild) DeepCopyInto(out *SonataFlowBuild) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowBuild.
func (in *SonataFlowBuild) DeepCopy() *SonataFlowBuild {
	if in == nil {
		return nil
	}
	out := new(SonataFlowBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SonataFlowBuild) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowBuildList) DeepCopyInto(out *SonataFlowBuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SonataFlowBuild, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowBuildList.
func (in *SonataFlowBuildList) DeepCopy() *SonataFlowBuildList {
	if in == nil {
		return nil
	}
	out := new(SonataFlowBuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SonataFlowBuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowBuildSpec) DeepCopyInto(out *SonataFlowBuildSpec) {
	*out = *in
	in.BuildTemplate.DeepCopyInto(&out.BuildTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowBuildSpec.
func (in *SonataFlowBuildSpec) DeepCopy() *SonataFlowBuildSpec {
	if in == nil {
		return nil
	}
	out := new(SonataFlowBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowBuildStatus) DeepCopyInto(out *SonataFlowBuildStatus) {
	*out = *in
	in.InnerBuild.DeepCopyInto(&out.InnerBuild)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowBuildStatus.
func (in *SonataFlowBuildStatus) DeepCopy() *SonataFlowBuildStatus {
	if in == nil {
		return nil
	}
	out := new(SonataFlowBuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowList) DeepCopyInto(out *SonataFlowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SonataFlow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowList.
func (in *SonataFlowList) DeepCopy() *SonataFlowList {
	if in == nil {
		return nil
	}
	out := new(SonataFlowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SonataFlowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowPlatform) DeepCopyInto(out *SonataFlowPlatform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowPlatform.
func (in *SonataFlowPlatform) DeepCopy() *SonataFlowPlatform {
	if in == nil {
		return nil
	}
	out := new(SonataFlowPlatform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SonataFlowPlatform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowPlatformList) DeepCopyInto(out *SonataFlowPlatformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SonataFlowPlatform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowPlatformList.
func (in *SonataFlowPlatformList) DeepCopy() *SonataFlowPlatformList {
	if in == nil {
		return nil
	}
	out := new(SonataFlowPlatformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SonataFlowPlatformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowPlatformSpec) DeepCopyInto(out *SonataFlowPlatformSpec) {
	*out = *in
	in.Build.DeepCopyInto(&out.Build)
	out.DevMode = in.DevMode
	in.Services.DeepCopyInto(&out.Services)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowPlatformSpec.
func (in *SonataFlowPlatformSpec) DeepCopy() *SonataFlowPlatformSpec {
	if in == nil {
		return nil
	}
	out := new(SonataFlowPlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowPlatformStatus) DeepCopyInto(out *SonataFlowPlatformStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowPlatformStatus.
func (in *SonataFlowPlatformStatus) DeepCopy() *SonataFlowPlatformStatus {
	if in == nil {
		return nil
	}
	out := new(SonataFlowPlatformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowSpec) DeepCopyInto(out *SonataFlowSpec) {
	*out = *in
	in.Flow.DeepCopyInto(&out.Flow)
	in.Resources.DeepCopyInto(&out.Resources)
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowSpec.
func (in *SonataFlowSpec) DeepCopy() *SonataFlowSpec {
	if in == nil {
		return nil
	}
	out := new(SonataFlowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SonataFlowStatus) DeepCopyInto(out *SonataFlowStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.Address.DeepCopyInto(&out.Address)
	in.LastTimeRecoverAttempt.DeepCopyInto(&out.LastTimeRecoverAttempt)
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SonataFlowStatus.
func (in *SonataFlowStatus) DeepCopy() *SonataFlowStatus {
	if in == nil {
		return nil
	}
	out := new(SonataFlowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowResources) DeepCopyInto(out *WorkflowResources) {
	*out = *in
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]ConfigMapWorkflowResource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowResources.
func (in *WorkflowResources) DeepCopy() *WorkflowResources {
	if in == nil {
		return nil
	}
	out := new(WorkflowResources)
	in.DeepCopyInto(out)
	return out
}
