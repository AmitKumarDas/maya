/*
Copyright 2019 The OpenEBS Authors

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
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=kubeassert

// KubeAssert contains the desired assertions of one or
// more resources
type KubeAssert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeAssertSpec   `json:"spec"`
	Status KubeAssertStatus `json:"status"`
}

// KubeAssertSpec provides the specifications of a
// KubeAssert
type KubeAssertSpec struct {
	// Describe is the detailed description of what
	// this assertion is supposed to achieve
	Describe string `json:"desc"`

	// Checks represents the list of checks to
	// be executed to evaluate the overall
	// assertion
	Checks []Check `json:"checks"`
}

// Check represents a verification
type Check struct {
	// It describes this check
	It string `json:"it"`

	// Kind of the resource to be verified
	Kind string `json:"kind"`

	// Name of the resource to be verified
	Name string `json:"name"`

	// Namespace of the resource to be verified
	Namespace string `json:"namespace"`

	// APIVersion of the resource
	APIVersion string `json:"apiVersion"`

	// LabelSelector to filter the resource(s)
	// that needs to be verified
	LabelSelector string `json:"labelSelector"`

	// AnnotationSelector to filter the resource(s)
	// that needs to be verified
	AnnotationSelector string `json:"annotationSelector"`

	// Expect contains the rules of this
	// check
	Expect Expect `json:"expect"`
}

// Expect contains the rules to evaluate
// a check
type Expect struct {
	// Match consists of a list of expected
	// matches
	Match []string `json:"match"`

	// Options represent the tunables to use
	// while evaluating these matches
	Options Options `json:"options"`
}

// Options represent the tunables that may be
// used while evaluating the matches
type Options struct {
	// Number of seconds before expectation is initiated.
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`

	// Number of seconds after which the handler times out.
	TimeoutSeconds int32 `json:"timeoutSeconds"`

	// How often (in seconds) to perform the check.
	PeriodSeconds int32 `json:"periodSeconds"`

	// Minimum consecutive successes for the probe to be considered
	// successful after having failed.
	SuccessThreshold int32 `json:"successThreshold"`

	// Minimum consecutive failures for the probe to be considered
	// failed after having succeeded.
	FailureThreshold int32 `json:"failureThreshold"`
}

// KubeAssertStatus represents the current state of KubeAssert
type KubeAssertStatus struct {
	Phase string `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=kubeasserts

// KubeAssertList is a list of KubeAsserts
type KubeAssertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KubeAssert `json:"items"`
}
