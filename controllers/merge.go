/*

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

package controllers

import (
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
)

func mergeResources(source interface{}, target interface{}) error {
	resource, err := json.Marshal(target)
	if err != nil {
		return err
	}
	patch, err := json.Marshal(source)
	if err != nil {
		return err
	}
	result, err := strategicpatch.StrategicMergePatch(resource, patch, target)
	if err != nil {
		return err
	}
	err = json.Unmarshal(result, target)
	if err != nil {
		return err
	}
	return nil
}
