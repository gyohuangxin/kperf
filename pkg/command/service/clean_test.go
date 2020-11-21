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
	"context"
	"testing"

	"gotest.tools/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sfake "k8s.io/client-go/kubernetes/fake"
	"knative.dev/kperf/pkg"
	"knative.dev/kperf/pkg/testutil"
)

func TestNewServiceCleanCommand(t *testing.T) {
	t.Run("incompleted or wrong args for service clean", func(t *testing.T) {
		client := k8sfake.NewSimpleClientset()
		p := &pkg.PerfParams{
			ClientSet: client,
		}
		cmd := NewServiceCleanCommand(p)

		_, err := testutil.ExecuteCommand(cmd)
		assert.ErrorContains(t, err, "both ns and nsPrefix are empty")

		_, err = testutil.ExecuteCommand(cmd, "--nsPrefix", "test-kperf-", "--nsRange", "2,1")
		assert.ErrorContains(t, err, "failed to parse namespace range 2,1")
	})

	t.Run("clean service as expected", func(t *testing.T) {
		client := k8sfake.NewSimpleClientset()
		p := &pkg.PerfParams{
			ClientSet: client,
		}

		nsSpec1 := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "test-kperf-1"}}
		_, err := p.ClientSet.CoreV1().Namespaces().Create(context.TODO(), nsSpec1, metav1.CreateOptions{})
		if err != nil {
			return
		}

		cmd := NewServiceCleanCommand(p)
		_, err = testutil.ExecuteCommand(cmd, "--ns", "test-kperf-1")
		assert.NilError(t, err)

		nsSpec2 := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "test-kperf-2"}}
		_, err = p.ClientSet.CoreV1().Namespaces().Create(context.TODO(), nsSpec2, metav1.CreateOptions{})
		if err != nil {
			return
		}

		cmd = NewServiceCleanCommand(p)
		_, err = testutil.ExecuteCommand(cmd, "--nsPrefix", "test-kperf-", "--nsRange", "1,2")
		assert.NilError(t, err)

		cmd = NewServiceCleanCommand(p)
		_, err = testutil.ExecuteCommand(cmd, "--nsPrefix", "test-kperf-", "--nsRange", "3,4")
		assert.ErrorContains(t, err, "no namespace found with prefix test-kperf-")
	})
}
