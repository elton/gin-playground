// Copyright 2020 Elton Zheng
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

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elton/gin-playground/api/controllers"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestPingRoute(t *testing.T) {
	r := gin.Default()
	controllers.Routers(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/welcome?firstname=Jane&lastname=Doe", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
