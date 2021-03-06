package v1alpha1

import (
	"github.com/pulumi/pulumi-kubernetes-operator/pkg/apis/pulumi/shared"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Stack is the Schema for the stacks API.
// Deprecated: Note Stacks from pulumi.com/v1alpha1 is deprecated in favor of pulumi.com/v1.
// It is completely backward compatible. Users are strongly encouraged to switch to pulumi.com/v1.
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=stacks,scope=Namespaced
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   shared.StackSpec   `json:"spec,omitempty"`
	Status shared.StackStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackList contains a list of Stack
// Deprecated: Note Stack:ost from pulumi.com/v1alpha1 is deprecated in favor of pulumi.com/v1.
// It is completely backward compatible. Users are strongly encouraged to switch to pulumi.com/v1.
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
