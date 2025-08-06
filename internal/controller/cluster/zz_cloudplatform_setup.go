// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	projectserviceidentity "github.com/upbound/provider-gcp-beta/internal/controller/cluster/cloudplatform/projectserviceidentity"
	serviceaccount "github.com/upbound/provider-gcp-beta/internal/controller/cluster/cloudplatform/serviceaccount"
)

// Setup_cloudplatform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectserviceidentity.Setup,
		serviceaccount.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudplatform creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectserviceidentity.SetupGated,
		serviceaccount.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
