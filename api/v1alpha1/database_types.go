// Copyright (c) 2021 Aiven, Helsinki, Finland. https://aiven.io/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Format="^[a-zA-Z0-9_-]*$"
	// Project to link the database to
	Project string `json:"project"`

	// +kubebuilder:validation:MaxLength=63
	// PostgreSQL service to link the database to
	ServiceName string `json:"serviceName"`

	// +kubebuilder:validation:MaxLength=128
	// Default string sort order (LC_COLLATE) of the database. Default value: en_US.UTF-8
	LcCollate string `json:"lcCollate,omitempty"`

	// +kubebuilder:validation:MaxLength=128
	// Default character classification (LC_CTYPE) of the database. Default value: en_US.UTF-8
	LcType string `json:"lcCtype,omitempty"`

	// It is a Kubernetes side deletion protections, which prevents the database
	// from being deleted by Kubernetes. It is recommended to enable this for any production
	// databases containing critical data.
	TerminationProtection bool `json:"terminationProtection,omitempty"`
}

// DatabaseStatus defines the observed state of Database
type DatabaseStatus struct {
	DatabaseSpec `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Database is the Schema for the databases API
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseSpec   `json:"spec,omitempty"`
	Status DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Database
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
