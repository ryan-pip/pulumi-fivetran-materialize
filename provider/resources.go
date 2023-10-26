// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fivetran

import (
	"fmt"
	"path/filepath"

	"github.com/MaterializeInc/pulumi-fivetran/provider/pkg/version"
	"github.com/fivetran/terraform-provider-fivetran/fivetran"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "fivetran"
	// modules:
	mainMod = "index" // the fivetran module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(fivetran.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "fivetran",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Fivetran",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumi",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/MaterializeInc",
		Description:       "A Pulumi package for creating and managing fivetran cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "fivetran", "category/utility"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/MaterializeInc/pulumi-fivetran",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "fivetran",
		Config: map[string]*tfbridge.SchemaInfo{
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FIVETRAN_API_KEY"},
				},
			},
			"api_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FIVETRAN_API_SECRET"},
				},
			},
			"api_url": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FIVETRAN_API_URL"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"fivetran_connector": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Connector"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"config": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"last_synced_changes__utc_": {
									Name: "last_synced_changes_utc",
								},
							},
						},
					},
				},
			},
			"fivetran_connector_schedule":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConnectorSchedule")},
			"fivetran_connector_schema_config": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConnectorSchemaConfig")},
			"fivetran_dbt_project":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DbtProject")},
			"fivetran_dbt_transformation": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DbtTransformation"),
			},
			"fivetran_destination": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Destination")},
			"fivetran_group":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Group")},
			"fivetran_group_users": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "GroupUsers")},
			"fivetran_user":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
			"fivetran_webhook":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Webhook")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"fivetran_user":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
			"fivetran_users":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUsers")},
			"fivetran_group":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGroup")},
			"fivetran_groups":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGroups")},
			"fivetran_group_connectors":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGroupConnectors")},
			"fivetran_group_users":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGroupUsers")},
			"fivetran_destination":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDestination")},
			"fivetran_connectors_metadata": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getConnectorsMetadata")},
			"fivetran_connector": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getConnector"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"config": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"last_synced_changes__utc_": {
									Name: "last_synced_changes_utc",
								},
							},
						},
					},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/MaterializeInc/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource tokens,
	// and apply auto aliasing for full backwards compatibility.
	// For more information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("fivetran_", mainMod,
		tokens.MakeStandard(mainPkg)))
	// prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
