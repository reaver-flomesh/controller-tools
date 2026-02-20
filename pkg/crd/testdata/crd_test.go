/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cronjob

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/yaml"
)

var _ = Describe("CronJob CRD", func() {
	It("should be successfully applied", func(ctx SpecContext) {
		data, err := os.ReadFile("testdata.kubebuilder.io_cronjobs.yaml")
		Expect(err).NotTo(HaveOccurred())

		crd := &apiextensionsv1.CustomResourceDefinition{}
		err = yaml.UnmarshalStrict(data, crd)
		Expect(err).NotTo(HaveOccurred())

		err = k8sClient.Create(ctx, crd)
		Expect(err).NotTo(HaveOccurred())
	})
})
