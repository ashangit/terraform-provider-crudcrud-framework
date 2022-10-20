package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-scaffolding-framework/internal/crudcrud"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &CrudcrudProvider{}
var _ provider.ProviderWithMetadata = &CrudcrudProvider{}

type CrudcrudProvider struct {
	version string
}

type CrudcrudProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *CrudcrudProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "crudcrud"
	resp.Version = p.version
}

func (p *CrudcrudProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"endpoint": {
				MarkdownDescription: "Crudcrud provider endpoint",
				Required:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func (p *CrudcrudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CrudcrudProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := crudcrud.CrudcrudClient{Endpoint: data.Endpoint.Value}
	//resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CrudcrudProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCrudcrudResource,
	}
}

func (p *CrudcrudProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		//CrudcrudDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CrudcrudProvider{
			version: version,
		}
	}
}
