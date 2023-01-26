// +groupName=crd.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ContainerSpec struct {
	Image string `json:"image,omitempty"`
	Port  int32  `json:"port,omitempty"`
}

// SakiibBhaiSpec defines the desired state of SakiibBhaiCRD
type SakiibBhaiSpec struct {
	Name      string        `json:"name,omitempty"`
	Replicas  *int32        `json:"replicas"`
	Container ContainerSpec `json:"container,container"`
}

// SakiibBhaiStatus defines the observed state of SakiibBhaiCRD
type SakiibBhaiStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SakiibBhai is the Schema for the sakiibBhai API
type SakiibBhai struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SakiibBhaiSpec   `json:"spec"`
	Status SakiibBhaiStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SakiibBhaiList contains a list of SakiibBhai
type SakiibBhaiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SakiibBhai `json:"items"`
}
