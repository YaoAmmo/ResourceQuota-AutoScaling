/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ResourcequotaAutoscalingSpec defines the desired state of ResourcequotaAutoscaling
type ResourcequotaAutoscalingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ResourcequotaAutoscaling. Edit resourcequotaautoscaling_types.go to remove/update
	ScaleTargetRef *ScaleTargetRef `json:"scaleTargetRef,omitempty"`
	MaxScaleCount int              `json:"maxScaleCount,omitempty"`
	ScaleResources *ScaleResources `json:"scaleResources,omitempty"`
}

// ResourcequotaAutoscalingStatus defines the observed state of ResourcequotaAutoscaling
type ResourcequotaAutoscalingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Message string                 `json:"message,omitempty"`
	Status bool                    `json:"status,omitempty"`
	BeforeIncrease *BeforeIncrease `json:"beforeIncrease,omitempty"`
	AfterIncrease *AfterIncrease   `json:"afterIncrease,omitempty"`

}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ResourcequotaAutoscaling is the Schema for the resourcequotaautoscalings API
type ResourcequotaAutoscaling struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourcequotaAutoscalingSpec   `json:"spec,omitempty"`
	Status ResourcequotaAutoscalingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ResourcequotaAutoscalingList contains a list of ResourcequotaAutoscaling
type ResourcequotaAutoscalingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcequotaAutoscaling `json:"items"`
}

type ScaleTargetRef struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type ScaleResources struct {
	CPU uint64 `json:"cpu"`
	Memory uint64 `json:"memory"`
}
type AfterIncrease struct {
	CPU uint64 `json:"cpu,omitempty"`
	Memory uint64 `json:"memory,omitempty"`
}

type BeforeIncrease struct {
	CPU uint64 `json:"cpu,omitempty"`
	Memory uint64 `json:"memory,omitempty"`
	TimeStamp time.Duration `json:"timeStamp,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ResourcequotaAutoscaling{}, &ResourcequotaAutoscalingList{})
}


