// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DdosProtectionConfigInitParameters struct {

	// Google Cloud Armor offers the following options to help protect systems against DDoS attacks:
	DdosProtection *string `json:"ddosProtection,omitempty" tf:"ddos_protection,omitempty"`
}

type DdosProtectionConfigObservation struct {

	// Google Cloud Armor offers the following options to help protect systems against DDoS attacks:
	DdosProtection *string `json:"ddosProtection,omitempty" tf:"ddos_protection,omitempty"`
}

type DdosProtectionConfigParameters struct {

	// Google Cloud Armor offers the following options to help protect systems against DDoS attacks:
	// +kubebuilder:validation:Optional
	DdosProtection *string `json:"ddosProtection" tf:"ddos_protection,omitempty"`
}

type RegionSecurityPolicyInitParameters struct {

	// Configuration for Google Cloud Armor DDOS Proctection Config.
	// Structure is documented below.
	DdosProtectionConfig *DdosProtectionConfigInitParameters `json:"ddosProtectionConfig,omitempty" tf:"ddos_protection_config,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type indicates the intended use of the security policy.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies.
	// A user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits.
	// Rules may then specify matching values for these fields.
	// Structure is documented below.
	UserDefinedFields []UserDefinedFieldsInitParameters `json:"userDefinedFields,omitempty" tf:"user_defined_fields,omitempty"`
}

type RegionSecurityPolicyObservation struct {

	// Configuration for Google Cloud Armor DDOS Proctection Config.
	// Structure is documented below.
	DdosProtectionConfig *DdosProtectionConfigObservation `json:"ddosProtectionConfig,omitempty" tf:"ddos_protection_config,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Fingerprint of this resource. This field is used internally during
	// updates of this resource.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/securityPolicies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier for the resource. This identifier is defined by the server.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created Region Security Policy should reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Server-defined URL for the resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Server-defined URL for this resource with the resource id.
	SelfLinkWithPolicyID *string `json:"selfLinkWithPolicyId,omitempty" tf:"self_link_with_policy_id,omitempty"`

	// The type indicates the intended use of the security policy.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies.
	// A user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits.
	// Rules may then specify matching values for these fields.
	// Structure is documented below.
	UserDefinedFields []UserDefinedFieldsObservation `json:"userDefinedFields,omitempty" tf:"user_defined_fields,omitempty"`
}

type RegionSecurityPolicyParameters struct {

	// Configuration for Google Cloud Armor DDOS Proctection Config.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DdosProtectionConfig *DdosProtectionConfigParameters `json:"ddosProtectionConfig,omitempty" tf:"ddos_protection_config,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created Region Security Policy should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The type indicates the intended use of the security policy.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies.
	// A user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits.
	// Rules may then specify matching values for these fields.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	UserDefinedFields []UserDefinedFieldsParameters `json:"userDefinedFields,omitempty" tf:"user_defined_fields,omitempty"`
}

type UserDefinedFieldsInitParameters struct {

	// The base relative to which 'offset' is measured. Possible values are:
	Base *string `json:"base,omitempty" tf:"base,omitempty"`

	// If specified, apply this mask (bitwise AND) to the field to ignore bits before matching.
	// Encoded as a hexadecimal number (starting with "0x").
	// The last byte of the field (in network byte order) corresponds to the least significant byte of the mask.
	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`

	// The name of this field. Must be unique within the policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Offset of the first byte of the field (in network byte order) relative to 'base'.
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Size of the field in bytes. Valid values: 1-4.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type UserDefinedFieldsObservation struct {

	// The base relative to which 'offset' is measured. Possible values are:
	Base *string `json:"base,omitempty" tf:"base,omitempty"`

	// If specified, apply this mask (bitwise AND) to the field to ignore bits before matching.
	// Encoded as a hexadecimal number (starting with "0x").
	// The last byte of the field (in network byte order) corresponds to the least significant byte of the mask.
	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`

	// The name of this field. Must be unique within the policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Offset of the first byte of the field (in network byte order) relative to 'base'.
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Size of the field in bytes. Valid values: 1-4.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type UserDefinedFieldsParameters struct {

	// The base relative to which 'offset' is measured. Possible values are:
	// +kubebuilder:validation:Optional
	Base *string `json:"base" tf:"base,omitempty"`

	// If specified, apply this mask (bitwise AND) to the field to ignore bits before matching.
	// Encoded as a hexadecimal number (starting with "0x").
	// The last byte of the field (in network byte order) corresponds to the least significant byte of the mask.
	// +kubebuilder:validation:Optional
	Mask *string `json:"mask,omitempty" tf:"mask,omitempty"`

	// The name of this field. Must be unique within the policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Offset of the first byte of the field (in network byte order) relative to 'base'.
	// +kubebuilder:validation:Optional
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Size of the field in bytes. Valid values: 1-4.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

// RegionSecurityPolicySpec defines the desired state of RegionSecurityPolicy
type RegionSecurityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionSecurityPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RegionSecurityPolicyInitParameters `json:"initProvider,omitempty"`
}

// RegionSecurityPolicyStatus defines the observed state of RegionSecurityPolicy.
type RegionSecurityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionSecurityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegionSecurityPolicy is the Schema for the RegionSecurityPolicys API. Represents a Region Cloud Armor Security Policy resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp-beta}
type RegionSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionSecurityPolicySpec   `json:"spec"`
	Status            RegionSecurityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionSecurityPolicyList contains a list of RegionSecurityPolicys
type RegionSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionSecurityPolicy `json:"items"`
}

// Repository type metadata.
var (
	RegionSecurityPolicy_Kind             = "RegionSecurityPolicy"
	RegionSecurityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionSecurityPolicy_Kind}.String()
	RegionSecurityPolicy_KindAPIVersion   = RegionSecurityPolicy_Kind + "." + CRDGroupVersion.String()
	RegionSecurityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RegionSecurityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionSecurityPolicy{}, &RegionSecurityPolicyList{})
}
