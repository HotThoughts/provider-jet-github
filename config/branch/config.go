package branch

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures provider
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_branch", func(r *config.Resource) {
		r.ShortGroup = "branch"
		r.ExternalName = config.IdentifierFromProvider

		r.References["repository"] = config.Reference{
			Type: "github.com/HotThoughts/provider-jet-github/apis/repository/v1alpha1.Repository",
		}
	})
}
