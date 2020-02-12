/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	_ "github.com/apache/camel-k/pkg/builder/kaniko"
	_ "github.com/apache/camel-k/pkg/builder/s2i"
	"github.com/apache/camel-k/pkg/cmd"

	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// List of addons that we want to include
	_ "github.com/apache/camel-k/addons/master"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	ctx, cancel := context.WithCancel(context.Background())

	// Cancel ctx as soon as main returns
	defer cancel()

	rootCmd, err := cmd.NewKamelCommand(ctx)
	exitOnError(err)

	err = rootCmd.Execute()
	exitOnError(err)
}

func exitOnError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
