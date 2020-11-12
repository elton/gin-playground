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
	"io"
	"os"
	"time"

	"github.com/elton/gin-playground/api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	date := time.Now().Format("20060102")
	// Logging to a file.
	f, err := os.Create("gin-" + date + ".log")
	if err != nil {
		panic(err)
	}
	// close f on exit and check for its returned error
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	// Force log's color
	gin.ForceConsoleColor()
	// write the log to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	controllers.Routers(r)

	r.Run()
}
