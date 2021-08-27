package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceCA provides information to configure an operator to manage the service cert controllers
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type ServiceCA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	//spec holds user settable values for configuration
	// +kubebuilder:validation:Required
	// +required
	Spec ServiceCASpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status ServiceCAStatus `json:"status"`
}

type ServiceCASpec struct {
	OperatorSpec `json:",inline"`
}

type ServiceCAStatus struct {
	OperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceCAList is a collection of items
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type ServiceCAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items contains the items
	Items []ServiceCA `json:"items"`
}
