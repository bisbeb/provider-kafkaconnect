package config

import (
    ctrl "sigs.k8s.io/controller-runtime"

    "github.com/crossplane/crossplane-runtime/pkg/controller"
    "github.com/crossplane/crossplane-runtime/pkg/event"
    "github.com/crossplane/crossplane-runtime/pkg/reconciler/providerconfig"
    "github.com/crossplane/crossplane-runtime/pkg/resource"

    "github.com/crossplane/provider-kafkaconnect/apis/v1alpha1"
)

// Setup adds a controller that reconciles ProviderConfigs.
func Setup(mgr ctrl.Manager, o controller.Options) error {
    name := providerconfig.ControllerName(v1alpha1.ProviderConfigGroupKind)

    of := resource.ProviderConfigKinds{
        Config:    v1alpha1.ProviderConfigGroupVersionKind,
        UsageList: v1alpha1.ProviderConfigUsageListGroupVersionKind,
    }

    r := providerconfig.NewReconciler(mgr, of,
        providerconfig.WithLogger(o.Logger.WithValues("controller", name)),
        providerconfig.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
    )

    return ctrl.NewControllerManagedBy(mgr).
        Named(name).
        WithOptions(o.ForControllerRuntime()).
        For(&v1alpha1.ProviderConfig{}).
        Watches(&v1alpha1.ProviderConfigUsage{}, &resource.EnqueueRequestForProviderConfig{}).
        Complete(r)
}
