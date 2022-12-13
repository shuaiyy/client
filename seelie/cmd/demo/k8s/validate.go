package k8s

import (
	"fmt"
	"strings"

	metav1validation "k8s.io/apimachinery/pkg/apis/meta/v1/validation"
	"k8s.io/apimachinery/pkg/util/validation"
)

// LabelValue if valid
func LabelValue(v string) error {
	if ss := validation.IsValidLabelValue(v); ss != nil {
		return fmt.Errorf("invalid label value: %s, msg: %s", v, strings.Join(ss, " "))
	}
	return nil
}

// ValidLabels ...
func ValidLabels(l map[string]string) error {
	errorList := metav1validation.ValidateLabels(l, nil)
	if len(errorList) > 0 {
		return errorList.ToAggregate()
	}
	return nil
}
