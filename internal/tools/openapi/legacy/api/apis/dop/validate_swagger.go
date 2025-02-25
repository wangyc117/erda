// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dop

import (
	"net/http"

	"github.com/erda-project/erda/internal/tools/openapi/legacy/api/apis"
)

var ValidateSwagger = apis.ApiSpec{
	Path:        "/api/apim/validate-swagger",
	BackendPath: "/api/apim/validate-swagger",
	Host:        APIMAddr,
	Scheme:      "http",
	Method:      http.MethodPost,
	CheckLogin:  false,
	CheckToken:  false,
	//RequestType:  apistructs.APIValidateSwaggerRequest{},
	ResponseType: nil,
	Doc:          "validate swagger",
}
