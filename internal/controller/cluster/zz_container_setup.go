// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "github.com/upbound/provider-gcp-beta/internal/controller/cluster/container/cluster"
	nodepool "github.com/upbound/provider-gcp-beta/internal/controller/cluster/container/nodepool"
)

// Setup_container creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_container(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		nodepool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_container creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_container(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.SetupGated,
		nodepool.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_container registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_container(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		cluster.SetupWebhookWithManager,
		nodepool.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
