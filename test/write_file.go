/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// WriteFile writes a file during testing.
func WriteFile(t *testing.T, filename string, format string, args ...interface{}) {
	t.Helper()

	WriteFileWithPem(t, filename, 0644, format, args...)
}

// WriteFileWithPerm writes a file with specific permissions during testing.
func WriteFileWithPem(t *testing.T, filename string, perm os.FileMode, format string, args ...interface{}) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		t.Fatal(err)
	}

	if err := ioutil.WriteFile(filename, []byte(fmt.Sprintf(format, args...)), perm); err != nil {
		t.Fatal(err)
	}
}
