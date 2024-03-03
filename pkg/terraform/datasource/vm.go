package dataresource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leryn1122/terraform-provider-vmware-fusion/pkg/rest"
)

var (
	_ datasource.DataSource              = &VirtualMachineDatasource{}
	_ datasource.DataSourceWithConfigure = &VirtualMachineDatasource{}
)

type VirtualMachineDatasource struct {
	client *rest.Client
}

type virtualMachineEntity struct {
	VmId types.String `tfsdk:"vm_id"`
}

func (vm VirtualMachineDatasource) Metadata(_ context.Context, _ datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = DataResourcePrefix + "vm"
}

func (vm VirtualMachineDatasource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vm_id": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func (vm VirtualMachineDatasource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var vmEntity virtualMachineEntity
	request.Config.Get(ctx, &vmEntity)

	vmId := vmEntity.VmId.ValueString()

	resp, err := vm.client.GetVm(rest.GetVmSettingInformationRequest{}, map[string]interface{}{
		"id": vmId,
	})
	if err != nil {
		response.Diagnostics.AddError(
			err.Message,
			err.Message,
		)
	}

	vmEntity.VmId = types.StringValue(resp.Id)

	diags := response.State.Set(ctx, &vmEntity)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (vm VirtualMachineDatasource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	vm.client = request.ProviderData.(*rest.Client)
}
