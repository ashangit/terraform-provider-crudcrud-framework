package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-scaffolding-framework/internal/crudcrud"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &CrudcrudResource{}
var _ resource.ResourceWithImportState = &CrudcrudResource{}

func NewCrudcrudResource() resource.Resource {
	return &CrudcrudResource{}
}

// ExampleResource defines the resource implementation.
type CrudcrudResource struct {
	client *crudcrud.CrudcrudClient
}

// ExampleResourceModel describes the resource data model.
type CrudcrudResourceModel struct {
	name  types.String `tfsdk:"name"`
	color types.String `tfsdk:"color"`
	age   types.Int64  `tfsdk:"age"`
	Id    types.String `tfsdk:"id"`
}

func (r *CrudcrudResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_unicorn"
}

func (r *CrudcrudResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Resource Crudcrud",

		Attributes: map[string]tfsdk.Attribute{
			"name": {
				MarkdownDescription: "Example configurable attribute",
				Required:            true,
				Type:                types.StringType,
			},
			"color": {
				MarkdownDescription: "Example configurable attribute",
				Required:            true,
				Type:                types.StringType,
			},
			"age": {
				MarkdownDescription: "Example configurable attribute",
				Required:            true,
				Type:                types.Int64Type,
			},
			"id": {
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Identifier",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
		},
	}, nil
}

func (r *CrudcrudResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*crudcrud.CrudcrudClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *crudcrud.CrudcrudClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *CrudcrudResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CrudcrudResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	unicorn := crudcrud.Unicorn{Name: data.name.Value, Age: int(data.age.Value), Color: data.color.Value}

	if err := r.client.Create(&unicorn); err != nil {
		resp.Diagnostics.AddError("Failed to create Unicorn", "detailed failed")
		return
	}

	fmt.Printf("Create: %v\n", unicorn)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a unicorn resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrudcrudResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CrudcrudResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	unicorn, err := r.client.Get(data.Id.Value)
	if err != nil {
		resp.Diagnostics.AddError("Failed to read Unicorn", "detailed failed")
		return
	}

	fmt.Printf("Read: %v\n", unicorn)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "Read a unicorn resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrudcrudResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CrudcrudResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	unicorn := crudcrud.Unicorn{Id: data.Id.Value, Name: data.name.Value, Age: int(data.age.Value), Color: data.color.Value}

	if err := r.client.Update(unicorn); err != nil {
		resp.Diagnostics.AddError("Failed to update Unicorn", "detailed failed")
		return
	}

	fmt.Printf("Read: %v\n", data)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "Update a unicorn resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrudcrudResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CrudcrudResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.Delete(data.Id.Value); err != nil {
		resp.Diagnostics.AddError("Failed to delete Unicorn", "detailed failed")
		return
	}

	fmt.Printf("Read: %v\n", data)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "Delete a unicorn resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrudcrudResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
