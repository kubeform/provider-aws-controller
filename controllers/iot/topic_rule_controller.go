// Code generated by Kubeform. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/go-logr/logr"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	auditlib "go.bytebuilders.dev/audit/lib"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	meta_util "kmodules.xyz/client-go/meta"
	iotv1alpha1 "kubeform.dev/provider-aws-api/apis/iot/v1alpha1"
	"kubeform.dev/provider-aws-controller/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// TopicRuleReconciler reconciles a TopicRule object
type TopicRuleReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	Gvk              schema.GroupVersionKind // GVK of the Resource
	Provider         *tfschema.Provider      // returns a *schema.Provider from the provider package
	Resource         *tfschema.Resource      // returns *schema.Resource
	TypeName         string                  // resource type
	WatchOnlyDefault bool
}

// +kubebuilder:rbac:groups=iot.aws.kubeform.com,resources=topicrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=iot.aws.kubeform.com,resources=topicrules/status,verbs=get;update;patch

func (r *TopicRuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("topicrule", req.NamespacedName)

	if r.WatchOnlyDefault && req.Namespace != v1.NamespaceDefault {
		log.Info("Only default namespace is supported for Kubeform Community, Please upgrade to Kubeform Enterprise to use any namespace.")
		return ctrl.Result{}, nil
	}
	var unstructuredObj unstructured.Unstructured
	unstructuredObj.SetGroupVersionKind(r.Gvk)

	if err := r.Get(ctx, req.NamespacedName, &unstructuredObj); err != nil {
		log.Error(err, "unable to fetch TopicRule")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	rClient := r.Client
	provider := r.Provider
	res := r.Resource
	gv := r.Gvk.GroupVersion()
	tName := r.TypeName
	jsonit := controllers.GetJSONItr(iotv1alpha1.GetEncoder(), iotv1alpha1.GetDecoder())
	err := controllers.StartProcess(rClient, provider, ctx, res, gv, &unstructuredObj, tName, jsonit)
	return ctrl.Result{}, err
}

func (r *TopicRuleReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, auditor *auditlib.EventPublisher) error {
	if auditor != nil {
		if err := auditor.SetupWithManager(ctx, mgr, &iotv1alpha1.TopicRule{}); err != nil {
			klog.Error(err, "unable to set up auditor", iotv1alpha1.TopicRule{}.APIVersion, iotv1alpha1.TopicRule{}.Kind)
			return err
		}
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&iotv1alpha1.TopicRule{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return !meta_util.MustAlreadyReconciled(e.Object)
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return (e.ObjectNew.(metav1.Object)).GetDeletionTimestamp() != nil || !meta_util.MustAlreadyReconciled(e.ObjectNew)
			},
		}).
		Owns(&v1.Secret{}).
		Complete(r)
}
