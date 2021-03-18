/*
Copyright 2020 kubeflow.org.

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

package utils

import (
	v1beta1api "github.com/kubeflow/kfserving/pkg/apis/serving/v1beta1"
)

// IsMMSPredictor enables MMS predictor when MultiModelServer field in ComponentExtensions
// is set or predictor config sets MMS to true and storage uri is not set
func IsMMSPredictor(predictor *v1beta1api.PredictorSpec, isvcConfig *v1beta1api.InferenceServicesConfig) bool {
	mmsExtension := predictor.GetExtensions().MultiModelServer
	isStorageURISet := predictor.GetImplementation().GetStorageUri() == nil
	if mmsExtension != nil {
		return *mmsExtension && isStorageURISet
	}
	return predictor.GetImplementation().IsMMS(isvcConfig) && isStorageURISet
}
