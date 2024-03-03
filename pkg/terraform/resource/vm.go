package resource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leryn1122/terraform-provider-vmware-fusion/pkg/rest"
)

var (
	_ resource.Resource              = &VirtualMachineResource{}
	_ resource.ResourceWithConfigure = &VirtualMachineResource{}
)

type VirtualMachineResource struct {
	client *rest.Client
}

type virtualMachineEntity struct {
	VmId types.String `tfsdk:"vm_id"`
}

func (vm VirtualMachineResource) Metadata(_ context.Context, _ resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = ResourcePrefix + "vm"
}

func (vm VirtualMachineResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vm_id": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func (vm VirtualMachineResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
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

func (vm VirtualMachineResource) Read(_ context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (vm VirtualMachineResource) Update(_ context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}

func (vm VirtualMachineResource) Delete(_ context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}

func (vm VirtualMachineResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}
	vm.client = request.ProviderData.(*rest.Client)
}
