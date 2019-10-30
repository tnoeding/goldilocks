// Copyright 2019 FairwindsOps Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package summary

import (
	"testing"

	"github.com/fairwindsops/goldilocks/pkg/kube"
	"github.com/fairwindsops/goldilocks/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	kubeClientVPA := kube.GetMockVPAClient()
	kubeClient := kube.GetMockClient()

	var summary Summary

	got, err := Run(kubeClient, kubeClientVPA, utils.VpaLabels, "true")
	assert.NoError(t, err)

	assert.EqualValues(t, got, summary)
}

// func TestRunSummary(t *testing.T) {
// constructSummary should be refactored to pass a client in.
// This test is commented out for now and
// Will be refactored/renamed after
// Some work is done on refactoring functions for better scope
// 	kubeClientVPA := kube.GetMockVPAClient()

// 	updateMode := v1beta2.UpdateModeOff
// 	var testVPA = &v1beta2.VerticalPodAutoscaler{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      "test-vpa",
// 			Labels:    utils.VpaLabels,
// 			Namespace: "testing",
// 		},
// 		Spec: v1beta2.VerticalPodAutoscalerSpec{
// 			TargetRef: &autoscaling.CrossVersionObjectReference{
// 				APIVersion: "extensions/v1beta1",
// 				Kind:       "Deployment",
// 				Name:       "test-vpa",
// 			},
// 			UpdatePolicy: &v1beta2.PodUpdatePolicy{
// 				UpdateMode: &updateMode,
// 			},
// 		},
// 	}
// 	var testVPANoLabels = &v1beta2.VerticalPodAutoscaler{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      "test-vpa",
// 			Namespace: "testing",
// 		},
// 		Spec: v1beta2.VerticalPodAutoscalerSpec{
// 			TargetRef: &autoscaling.CrossVersionObjectReference{
// 				APIVersion: "extensions/v1beta1",
// 				Kind:       "Deployment",
// 				Name:       "test-vpa",
// 			},
// 			UpdatePolicy: &v1beta2.PodUpdatePolicy{
// 				UpdateMode: &updateMode,
// 			},
// 		},
// 	}

// 	_, errOk := kubeClientVPA.Client.AutoscalingV1beta2().VerticalPodAutoscalers("testing").Create(testVPA)
// 	assert.NoError(t, errOk)

// 	_, errOk2 := kubeClientVPA.Client.AutoscalingV1beta2().VerticalPodAutoscalers("testing").Create(testVPANoLabels)
// 	assert.NoError(t, errOk2)

// 	var summary = Summary{
// 		Namespaces: []string{
// 			"testing",
// 		},
// 	}

// 	got, err := Run(kubeClientVPA, utils.VpaLabels, "true")
// 	assert.NoError(t, err)

// 	assert.EqualValues(t, summary, got)
// }