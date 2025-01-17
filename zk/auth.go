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

package zk

import (
	"github.com/sleepymole/zetcd"
)

func NewAuth(addrs []string) zetcd.AuthFunc {
	return func(zka zetcd.AuthConn) (zetcd.Session, error) {
		return newSession(addrs, zka)
	}
}

func NewZK() zetcd.ZKFunc {
	return func(s zetcd.Session) (zetcd.ZK, error) {
		zk, err := newZK(s)
		return zetcd.NewZKLog(zk), err
	}
}
