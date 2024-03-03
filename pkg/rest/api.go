package rest

import (
	"net/http"
)

func (c *Client) GetVmList(request GetVmListRequest, params Params) (*GetVmListResponse, *VmwareFusionError) {
	return SendRequest[GetVmListRequest, GetVmListResponse](
		c, VMWareAPI_GetVmList, http.MethodGet, request, params)
}

func (c *Client) GetVm(request GetVmSettingInformationRequest, params Params) (*GetVmSettingInformationResponse, *VmwareFusionError) {
	return SendRequest[GetVmSettingInformationRequest, GetVmSettingInformationResponse](
		c, VMWareAPI_GetVm, http.MethodGet, request, params)
}
