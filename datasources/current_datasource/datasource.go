package current_datasource

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kasaikou/terraform-provider-db/client"
)

type currentDataSource struct {
	client *client.DatabaseClient
}

func New() datasource.DataSource { return &currentDataSource{} }

func (d *currentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, res *datasource.MetadataResponse) {
	*res = datasource.MetadataResponse{
		TypeName: req.ProviderTypeName + "_current",
	}
}

func (d *currentDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, res *datasource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user": schema.StringAttribute{
				Computed: true,
			},
			"dbname": schema.StringAttribute{
				Computed: true,
			},
			"version": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *currentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, res *datasource.ConfigureResponse) {

	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.DatabaseClient)
	if !ok {
		res.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.DatabaseClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
	}

	if res.Diagnostics.HasError() {
		return
	}

	d.client = client
}

type currentDataSourceResponse struct {
	User    types.String `tfsdk:"user"`
	DBName  types.String `tfsdk:"dbname"`
	Version types.String `tfsdk:"version"`
}

func (d *currentDataSource) Read(ctx context.Context, req datasource.ReadRequest, res *datasource.ReadResponse) {

	current, err := d.client.Current(ctx)
	if err != nil {
		res.Diagnostics.AddError(
			"Failed to Get current status",
			err.Error(),
		)
	}

	result := currentDataSourceResponse{
		User:    types.StringValue(current.User),
		DBName:  types.StringValue(current.DBName),
		Version: types.StringValue(current.Version),
	}

	diagnostics := res.State.Set(ctx, &result)
	res.Diagnostics.Append(diagnostics...)
}
