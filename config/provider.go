/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/HotThoughts/provider-jet-github/config/branch"
	"github.com/HotThoughts/provider-jet-github/config/repository"
	"github.com/HotThoughts/provider-jet-github/config/team"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourcePrefix = "github"
	modulePath     = "github.com/HotThoughts/provider-jet-github"
)

//go:embed schema.json
var providerSchema string

// IncludedResources lists all resource patterns included in small set release.
var IncludedResources = []string{
	// Branch
	"github_branch$",
	// Repository
	"github_repository$",
	// Team
	"github_team$",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList(IncludedResources),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		repository.Configure,
		branch.Configure,
		team.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
