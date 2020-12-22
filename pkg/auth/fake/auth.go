/*
Copyright (c) 2018 Bitnami

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

package fake

import (
	authUtils "github.wdf.sap.corp/kubernetes/hub/pkg/auth"
)

type Auth struct {
	ForbiddenActions []authUtils.Action
}

func (f *Auth) Validate() error {
	return nil
}

func (f *Auth) GetForbiddenActions(namespace, action, manifest string) ([]authUtils.Action, error) {
	return f.ForbiddenActions, nil
}
