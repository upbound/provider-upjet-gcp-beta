// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package compute

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_compute_region_security_policy", func(r *config.Resource) {
		r.MarkAsRequired("region")
	})

	p.AddResourceConfigurator("google_compute_region_backend_service", func(r *config.Resource) {
		r.MarkAsRequired("region")
		r.References["health_checks"] = config.Reference{
			TerraformName: "google_compute_health_check",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
