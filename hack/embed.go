// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package hack

import _ "embed"

// MainTemplate is populated with provider main program template.
//
//go:embed main.go.tmpl
var MainTemplate string
