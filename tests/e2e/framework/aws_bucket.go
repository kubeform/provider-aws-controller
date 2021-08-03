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

package framework

import (
	"context"
	"time"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kubeform.dev/provider-aws-api/apis/s3/v1alpha1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

func (i *Invocation) Bucket(name, secretName string) *v1alpha1.Bucket {
	acl := "private"

	return &v1alpha1.Bucket{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.BucketSpec{
			ProviderRef: corev1.LocalObjectReference{
				Name: secretName,
			},
			Resource: v1alpha1.BucketSpecResource{
				Acl:    &acl,
				Bucket: &name,
				Tags: &map[string]string{
					"env": "dev",
				},
			},
		},
	}
}

func (f *Framework) CreateBucket(obj *v1alpha1.Bucket) error {
	_, err := f.kfClient.S3V1alpha1().Buckets(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func (f *Framework) UpdateBucket(obj *v1alpha1.Bucket) error {
	obj, err := f.kfClient.S3V1alpha1().Buckets(obj.Namespace).Get(context.TODO(), obj.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	updateBucket := *obj.Spec.Resource.Bucket + "-update"
	obj.Spec.Resource.Bucket = &updateBucket

	_, err = f.kfClient.S3V1alpha1().Buckets(obj.Namespace).Update(context.TODO(), obj, metav1.UpdateOptions{})
	return err
}

func (f *Framework) DeleteBucket(meta metav1.ObjectMeta) error {
	return f.kfClient.S3V1alpha1().Buckets(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInBackground())
}

func (f *Framework) EventuallyBucketRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			s3Bucket, err := f.kfClient.S3V1alpha1().Buckets(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return s3Bucket.Status.Phase == status.CurrentStatus
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyBucketDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kfClient.S3V1alpha1().Buckets(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
