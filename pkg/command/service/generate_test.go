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

package service

import (
	"testing"

	"gotest.tools/assert"
	k8sfake "k8s.io/client-go/kubernetes/fake"
	"knative.dev/kperf/pkg"
	"knative.dev/kperf/pkg/testutil"
)

func TestNewServiceGenerateCommand(t *testing.T) {
	t.Run("incompleted args for service generate", func(t *testing.T) {
		client := k8sfake.NewSimpleClientset()
		p := &pkg.PerfParams{
			ClientSet: client,
		}
		cmd := NewServiceGenerateCommand(p)

		_, err := testutil.ExecuteCommand(cmd)
		assert.ErrorContains(t, err, "required flag(s) \"batch\", \"interval\", \"minScale\" not set")

		_, err = testutil.ExecuteCommand(cmd, "-b", "1")
		assert.ErrorContains(t, err, "required flag(s) \"interval\", \"minScale\" not set")

		_, err = testutil.ExecuteCommand(cmd, "-b", "1", "-i", "1")
		assert.ErrorContains(t, err, "required flag(s) \"minScale\" not set")

		_, err = testutil.ExecuteCommand(cmd, "-b", "1", "-i", "1", "--minScale", "1")
		assert.ErrorContains(t, err, "both ns and nsPrefix are empty")
	})
}
