// Copyright 2016 CoreOS, Inc.
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

// Package integration is used for testing zetcd against a full etcd server.
//
// Build `go test -tags zkdocker` to run tests against a zookeeper instance
// over docker.
package integration

// force import so glide picks it up
import (
	_ "go.etcd.io/etcd/tests/v3/integration"
)
