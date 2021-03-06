package flanneltpr

import (
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomObject represents the Flannel TPR's custom object. It holds the
// specifications of the resource the Flannel operator is interested in.
type CustomObject struct {
	apismetav1.TypeMeta   `json:",inline"`
	apismetav1.ObjectMeta `json:"metadata,omitempty"`

	Spec Spec `json:"spec" yaml:"spec"`
}
