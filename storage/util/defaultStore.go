//go:build !wasm
// +build !wasm

/*
 * Flow Emulator
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"github.com/onflow/flow-emulator/storage"
	"github.com/onflow/flow-emulator/storage/redis"
	"github.com/onflow/flow-emulator/storage/sqlite"
)

func CreateDefaultStorage() (storage.Store, error) {
	return sqlite.New(sqlite.InMemory)
}

func NewSqliteStorage(url string) (storage.Store, error) {
	return sqlite.New(url)
}

func NewRedisStorage(url string) (storage.Store, error) {
	return redis.New(url)
}
