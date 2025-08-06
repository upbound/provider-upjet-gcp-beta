// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	projectserviceidentity "github.com/upbound/provider-gcp-beta/internal/controller/cluster/cloudplatform/projectserviceidentity"
	serviceaccount "github.com/upbound/provider-gcp-beta/internal/controller/cluster/cloudplatform/serviceaccount"
	healthcheck "github.com/upbound/provider-gcp-beta/internal/controller/cluster/compute/healthcheck"
	regionbackendservice "github.com/upbound/provider-gcp-beta/internal/controller/cluster/compute/regionbackendservice"
	regionsecuritypolicy "github.com/upbound/provider-gcp-beta/internal/controller/cluster/compute/regionsecuritypolicy"
	cluster "github.com/upbound/provider-gcp-beta/internal/controller/cluster/container/cluster"
	nodepool "github.com/upbound/provider-gcp-beta/internal/controller/cluster/container/nodepool"
	servertlspolicy "github.com/upbound/provider-gcp-beta/internal/controller/cluster/networksecurity/servertlspolicy"
	providerconfig "github.com/upbound/provider-gcp-beta/internal/controller/cluster/providerconfig"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectserviceidentity.Setup,
		serviceaccount.Setup,
		healthcheck.Setup,
		regionbackendservice.Setup,
		regionsecuritypolicy.Setup,
		cluster.Setup,
		nodepool.Setup,
		servertlspolicy.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
