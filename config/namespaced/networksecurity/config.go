// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package networksecurity

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_network_security_server_tls_policy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "location")
		r.ShortGroup = "networksecurity"
		r.Kind = "ServerTLSPolicy"
	})
}
