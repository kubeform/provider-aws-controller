/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	awsbucket "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"
	"kubeform.dev/provider-aws-controller/tests/e2e/framework"
)

var _ = Describe("Test", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
		if !framework.RunTest(framework.BUCKET, whichController) {
			Skip(fmt.Sprintf("`%s` test is applied only when whichController flag is either `all` or `%s` but got `%s`", framework.BUCKET, framework.BUCKET, whichController))
		}
	})
	Describe("AWS", func() {
		Context("BucketController", func() {
			var (
				providerRef *core.Secret
				secretName  string
				bucketName  string
				bucket      *awsbucket.Bucket
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				bucketName = f.GetRandomName("")
				providerRef = f.AwsProviderRef(secretName)
				bucket = f.Bucket(bucketName, secretName)
			})

			AfterEach(func() {
				By("Deleting Bucket")
				err = f.DeleteBucket(bucket.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting Bucket")
				f.EventuallyBucketDeleted(bucket.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})

			It("should create and delete bucket successfully", func() {
				By("Creating AwsProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Bucket")
				err = f.CreateBucket(bucket)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for Running Bucket")
				f.EventuallyBucketRunning(bucket.ObjectMeta).Should(BeTrue())
			})

			It("should create, update and delete bucket successfully", func() {
				By("Creating AwsProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Bucket")
				err = f.CreateBucket(bucket)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for Running Bucket")
				f.EventuallyBucketRunning(bucket.ObjectMeta).Should(BeTrue())

				By("Updating Bucket")
				err = f.UpdateBucket(bucket)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Bucket")
				f.EventuallyBucketRunning(bucket.ObjectMeta).Should(BeTrue())
			})
		})
	})
})
