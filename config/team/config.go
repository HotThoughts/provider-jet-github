package team

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("github_team", func(r *config.Resource) {
        r.ShortGroup = "team"
				r.ExternalName = config.IdentifierFromProvider
    })
}
