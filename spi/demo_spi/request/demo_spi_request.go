package demo_spi_request

import (
	doudian_sdk "github.com/ldxtechteam/doudiansdk/core"
	demo_spi_response "github.com/ldxtechteam/doudiansdk/spi/demo_spi/response"
	"github.com/ldxtechteam/doudiansdk/utils"
)

type DemoSpiRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param    *DemoSpiParam
	Response *demo_spi_response.DemoSpiResponse
}

type DemoSpiParam struct {
	Arg1 string `json:"arg_1"`
}

func (b *DemoSpiRequest) GetParamJsonObject() interface{} {
	return b.Param
}

func (b *DemoSpiRequest) GetResponseObject() interface{} {
	return b.Response
}

func (b *DemoSpiRequest) Execute() (interface{}, error) {
	return b.GetClient().Request(b)
}

func (b *DemoSpiRequest) ResponseJson() (string, error) {
	responseObj, err := b.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}

func New() *DemoSpiRequest {
	request := new(DemoSpiRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(DemoSpiParam)
	response := new(demo_spi_response.DemoSpiResponse)
	response.Data = new(demo_spi_response.DemoSpiData)
	request.Response = response
	return request
}
