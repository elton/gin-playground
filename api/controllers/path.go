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

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Path Parameters in path
func Path(ctx *gin.Context) {

	name := ctx.Param("name")
	action := ctx.Param("action")
	if action != "" {
		message := name + " is " + action
		ctx.String(http.StatusOK, message)
		return
	}
	ctx.String(http.StatusOK, "Hello %s", name)

}

// QueryString query string parameters
func QueryString(ctx *gin.Context) {
	firstname := ctx.DefaultQuery("firstname", "Guest")
	lastname := ctx.Query("lastname")

	ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
