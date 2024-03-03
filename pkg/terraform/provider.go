package terraform

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/leryn1122/terraform-provider-vmware-fusion/pkg/rest"
	customdatasource "github.com/leryn1122/terraform-provider-vmware-fusion/pkg/terraform/datasource"
	customresource "github.com/leryn1122/terraform-provider-vmware-fusion/pkg/terraform/resource"
	"os"
)

const (
	ProviderKey_Username = "username"
	ProviderKey_Password = "password"
	ProviderKey_URL      = "url"
	ProviderKey_Insecure = "insecure"
	ProviderKey_Debug    = "debug"
)

var (
	_ provider.Provider = &vmwareFusionProvider{}
)

type vmwareFusionProvider struct {
}

type vmwareFusionProviderEntity struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	BaseURL  types.String `tfsdk:"url"`
	Insecure types.Bool   `tfsdk:"insecure"`
}

// NewProvider Construct a new VMWare Fusion provider instance.
func NewProvider() provider.Provider {
	return &vmwareFusionProvider{}
}

func (p *vmwareFusionProvider) Metadata(_ context.Context, _ provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "vmware-fusion"
	response.Version = "0.1.0"
}

func (p *vmwareFusionProvider) Schema(_ context.Context, _ provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "Provider for VMWare Fusion API RESTful operations",
		Attributes: map[string]schema.Attribute{
			ProviderKey_Username: schema.StringAttribute{
				Description: "The username of VMWare Fusion API RESTful operations",
				Required:    true,
				Sensitive:   true,
			},
			ProviderKey_Password: schema.StringAttribute{
				Description: "The password of VMWare Fusion RESTful API operations",
				Required:    true,
				Sensitive:   true,
			},
			ProviderKey_URL: schema.StringAttribute{
				Description: "The base URL of VMWare Fusion RESTful API operations",
				Required:    true,
			},
			ProviderKey_Insecure: schema.BoolAttribute{
				Description: "True if using HTTPS towards VMWare Fusion RESTful API operations",
				Optional:    true,
			},
		},
	}
}

func (p *vmwareFusionProvider) Configure(ctx context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configured provider: terraform-provider-vmware-fusion")

	var config vmwareFusionProviderEntity
	diags := request.Config.Get(ctx, &config)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if config.Username.IsUnknown() {
		response.Diagnostics.AddAttributeError(
			path.Root(ProviderKey_Username),
			"",
			"",
		)
	}

	if config.Password.IsUnknown() {
		response.Diagnostics.AddAttributeError(
			path.Root(ProviderKey_Password),
			"",
			"",
		)
	}

	if config.BaseURL.IsUnknown() {
		response.Diagnostics.AddAttributeError(
			path.Root(ProviderKey_URL),
			"",
			"",
		)
	}

	if response.Diagnostics.HasError() {
		return
	}

	username := os.Getenv("VMFS_USERNAME")
	password := os.Getenv("VMFS_PASSWORD")
	url := os.Getenv("VMFS_URL")

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}
	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}
	if !config.BaseURL.IsNull() {
		url = config.BaseURL.ValueString()
	}

	if response.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, ProviderKey_Username, username)
	ctx = tflog.SetField(ctx, ProviderKey_Password, password)
	ctx = tflog.SetField(ctx, ProviderKey_URL, url)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, ProviderKey_Password)

	client, err := rest.NewClient(username, password, url)
	if err != nil {
		response.Diagnostics.AddError(err.Error(), err.Error())
		return
	}

	response.DataSourceData = client
	response.ResourceData = client
}

func (p *vmwareFusionProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		func() datasource.DataSource {
			return customdatasource.VirtualMachineDatasource{}
		},
	}
}

func (p *vmwareFusionProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		func() resource.Resource {
			return customresource.VirtualMachineResource{}
		},
	}
}
