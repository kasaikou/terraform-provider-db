package provider

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kasaikou/terraform-provider-db/client"
)

func (p *databaseProvider) Schema(_ context.Context, _ provider.SchemaRequest, res *provider.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"driver": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(sql.Drivers()...),
				},
			},
			"data_source": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

type databaseProviderConfig struct {
	Driver     types.String `tfsdk:"driver"`
	DataSource types.String `tfsdk:"data_source"`
}

func (p *databaseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, res *provider.ConfigureResponse) {

	config := databaseProviderConfig{}
	diagnostics := req.Config.Get(ctx, &config)
	if res.Diagnostics.Append(diagnostics...); res.Diagnostics.HasError() {
		return
	}

	client, err := client.New(config.Driver.ValueString(), config.DataSource.ValueString())
	if err != nil {
		res.Diagnostics.AddError(
			"Cannot make DB connection",
			fmt.Sprintf("Failed to connection with driver and data_source (error=\"%s\").", err.Error()),
		)
	}

	if res.Diagnostics.HasError() {
		return
	}

	res.DataSourceData = client
	res.ResourceData = client
}
