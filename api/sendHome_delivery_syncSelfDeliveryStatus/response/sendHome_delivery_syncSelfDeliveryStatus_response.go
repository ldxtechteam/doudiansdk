package sendHome_delivery_syncSelfDeliveryStatus_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeDeliverySyncSelfDeliveryStatusResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeDeliverySyncSelfDeliveryStatusData `json:"data"`
}
type SendHomeDeliverySyncSelfDeliveryStatusData struct {
}
