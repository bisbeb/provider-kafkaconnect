package controller

import (
    ctrl "sigs.k8s.io/controller-runtime"

    "github.com/crossplane/crossplane-runtime/pkg/controller"

    "github.com/crossplane/provider-kafkaconnect/internal/controller/connector"
    "github.com/crossplane/provider-kafkaconnect/internal/controller/config"
)

// Setup creates all KafkaConnect controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
    for _, setup := range []func(ctrl.Manager, controller.Options) error{
        config.Setup,
        connector.Setup,
    } {
        if err := setup(mgr, o); err != nil {
            return err
        }
    }
    return nil
}
