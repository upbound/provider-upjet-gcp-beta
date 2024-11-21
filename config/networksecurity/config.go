// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package networksecurity

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_network_security_server_tls_policy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "location")
		r.ShortGroup = "networksecurity"
		r.Kind = "ServerTLSPolicy"
	})
}
