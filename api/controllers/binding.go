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

	"github.com/elton/gin-playground/api/models"
	"github.com/gin-gonic/gin"
)

// Login user log in and binding to the user struct.
func Login(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in."})
}

// URIBind binding uri to model.
func URIBind(ctx *gin.Context) {
	person := models.Person{}
	if err := ctx.ShouldBindUri(&person); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	ctx.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}
