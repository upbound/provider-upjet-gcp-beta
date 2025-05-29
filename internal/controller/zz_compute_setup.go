// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	healthcheck "github.com/upbound/provider-gcp-beta/internal/controller/compute/healthcheck"
	regionbackendservice "github.com/upbound/provider-gcp-beta/internal/controller/compute/regionbackendservice"
	regionsecuritypolicy "github.com/upbound/provider-gcp-beta/internal/controller/compute/regionsecuritypolicy"
)

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		healthcheck.Setup,
		regionbackendservice.Setup,
		regionsecuritypolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
