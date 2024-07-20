package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/kasaikou/terraform-provider-db/datasources/current_datasource"
)

type databaseProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &databaseProvider{
			version: version,
		}
	}
}

func (p *databaseProvider) Metadata(_ context.Context, _ provider.MetadataRequest, res *provider.MetadataResponse) {
	*res = provider.MetadataResponse{
		TypeName: "db",
		Version:  p.version,
	}
}

func (p *databaseProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		current_datasource.New,
	}
}

func (p *databaseProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
