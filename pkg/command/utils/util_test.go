// Copyright © 2020 The Knative Authors
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

package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestGenerateCSVFile(t *testing.T) {
	t.Run("generate CSV file successfully", func(t *testing.T) {
		path := "/tmp/generate-csv-test.csv"
		rows := [][]string{{"1", "2", "3"}, {"1", "2", "3"}, {"1", "2", "3"}}
		err := GenerateCSVFile(path, rows)
		assert.NilError(t, err)
	})

	t.Run("return error if path is not available", func(t *testing.T) {
		path := "/tmp"
		rows := [][]string{{"1", "2", "3"}, {"1", "2", "3"}, {"1", "2", "3"}}
		err := GenerateCSVFile(path, rows)
		assert.ErrorContains(t, err, "failed to create csv file open /tmp: is a directory", err)
	})
}

func TestGenerateHTMLFile(t *testing.T) {
	t.Run("generate HTML file successfully", func(t *testing.T) {
		sourceCSV := "test.csv"
		targetHTML := "/tmp/test.html"
		err := GenerateHTMLFile(sourceCSV, targetHTML)
		assert.NilError(t, err)
	})
}
