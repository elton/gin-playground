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
	"github.com/gin-gonic/gin"
)

// Routers all routers for the application.
func Routers(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// v1.GET("/user/:name", Path)
		// v1.GET("/user/:name/*action", Path)
		// curl -i http://localhost:8080/api/v1/welcome\?firstname=Jane\&lastname=Doe
		v1.GET("/welcome", QueryString)
		// curl -v -X POST \
		//   http://localhost:8080/api/v1/login \
		//   -H 'content-type: application/json' \
		//   -d '{ "user": "manu","password":"123" }'
		v1.POST("/login", Login)
		v1.GET("/user/:pname/:id", URIBind)
	}

}
