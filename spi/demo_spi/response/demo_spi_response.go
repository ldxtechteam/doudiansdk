package demo_spi_response

import doudian_sdk "github.com/ldxtechteam/doudiansdk/core"

type DemoSpiResponse struct {
	doudian_sdk.BaseDoudianOpSpiResponse
	Data *DemoSpiData `json:"data"`
}

func (d *DemoSpiResponse) GetData() interface{} {
	return d.Data
}

type DemoSpiData struct {
	Data1 string
}
