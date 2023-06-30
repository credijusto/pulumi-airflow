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

package airflow

import (
	"fmt"
	"github.com/Hellthrashers/pulumi-airflow/provider/pkg/version"
	airflow "github.com/drfaust92/terraform-provider-airflow/shim"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"path/filepath"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "airflow"
	// modules:
	mainMod = "index" // the airflow module
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
	p := shimv2.NewProvider(airflow.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "airflow",
		Publisher:         "Marco Antonio Ojeda De Pablo",
		PluginDownloadURL: "github://api.github.com/Hellthrashers/pulumi-airflow",
		Description:       "A Pulumi package for creating and managing airflow cloud resources.",
		Keywords:          []string{"pulumi", "airflow", "apache/airflow", "dags"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/Hellthrashers/pulumi-airflow",
		Repository:        "https://github.com/Hellthrashers/pulumi-airflow",
		GitHubOrg:         "drfaust92",
		Config: map[string]*tfbridge.SchemaInfo{
			"base_endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"AIRFLOW_BASE_ENDPOINT",
					},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"airflow_connection": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Connection")},
			"airflow_dag":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Dag")},
			"airflow_pool":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Pool")},
			"airflow_variable":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Variable")},
			"airflow_role":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Role")},
			"airflow_user":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
			"airflow_dag_run":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DagRun")},
		},

		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "pulumi-airflow",
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
				fmt.Sprintf("github.com/Hellthrashers/pulumi-%[1]s/sdk/", mainPkg),
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

	prov.SetAutonaming(255, "-")

	return prov
}
