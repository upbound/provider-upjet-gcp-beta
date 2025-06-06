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

type ServiceAccountInitParameters struct {

	// If set to true, skip service account creation if a service account with the same email already exists.
	CreateIgnoreAlreadyExists *bool `json:"createIgnoreAlreadyExists,omitempty" tf:"create_ignore_already_exists,omitempty"`

	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether a service account is disabled or not. Defaults to false. This field has no effect during creation.
	// Must be set after creation to disable a service account.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ServiceAccountObservation struct {

	// If set to true, skip service account creation if a service account with the same email already exists.
	CreateIgnoreAlreadyExists *bool `json:"createIgnoreAlreadyExists,omitempty" tf:"create_ignore_already_exists,omitempty"`

	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether a service account is disabled or not. Defaults to false. This field has no effect during creation.
	// Must be set after creation to disable a service account.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The e-mail address of the service account. This value
	// should be referenced from any google_iam_policy data sources
	// that would grant the service account privileges.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// an identifier for the resource with format projects/{{project}}/serviceAccounts/{{email}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Identity of the service account in the form serviceAccount:{email}. This value is often used to refer to the service account in order to grant IAM permissions.
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// The fully-qualified name of the service account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The unique id of the service account.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type ServiceAccountParameters struct {

	// If set to true, skip service account creation if a service account with the same email already exists.
	// +kubebuilder:validation:Optional
	CreateIgnoreAlreadyExists *bool `json:"createIgnoreAlreadyExists,omitempty" tf:"create_ignore_already_exists,omitempty"`

	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether a service account is disabled or not. Defaults to false. This field has no effect during creation.
	// Must be set after creation to disable a service account.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The display name for the service account.
	// Can be updated without creating a new resource.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// ServiceAccountSpec defines the desired state of ServiceAccount
type ServiceAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountParameters `json:"forProvider"`
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
	InitProvider ServiceAccountInitParameters `json:"initProvider,omitempty"`
}

// ServiceAccountStatus defines the observed state of ServiceAccount.
type ServiceAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceAccount is the Schema for the ServiceAccounts API. Allows management of a Google Cloud Platform service account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp-beta}
type ServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountSpec   `json:"spec"`
	Status            ServiceAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountList contains a list of ServiceAccounts
type ServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccount `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccount_Kind             = "ServiceAccount"
	ServiceAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccount_Kind}.String()
	ServiceAccount_KindAPIVersion   = ServiceAccount_Kind + "." + CRDGroupVersion.String()
	ServiceAccount_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccount{}, &ServiceAccountList{})
}
