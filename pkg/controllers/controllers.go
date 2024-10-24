// Copyright (c) 2024 Red Hat, Inc.

package controllers

import ctrl "sigs.k8s.io/controller-runtime"

// +kubebuilder:rbac:groups="",resources=configmaps;secrets,verbs=get;list;watch;create;patch;delete
// +kubebuilder:rbac:groups="",resources=events,verbs=create
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update
// +kubebuilder:rbac:groups=nim.opendatahub.io,resources=odhnimapps,verbs=get;list;watch;create;patch;delete

const (
	Finalizer_NimAppCleanup = "nim.opendatahub.io/cleanup_finalizer"
	Label_NimApp            = "nim.opendatahub.io/nim-app"
)

// ControllerOptions is encapsulating the global options for use with all controllers
type ControllerOptions struct {
	Manager ctrl.Manager
}

// controllerSetups is used for registering controllers for loading
var controllerSetups []func(ControllerOptions) error

// SetupControllers is used for setting up all registered controllers with the global options
func SetupControllers(opts ControllerOptions) error {
	for _, ctrlSetup := range controllerSetups {
		if err := ctrlSetup(opts); err != nil {
			return err
		}
	}
	return nil
}
