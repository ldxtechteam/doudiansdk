package sendHome_store_updateStoreState_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeStoreUpdateStoreStateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeStoreUpdateStoreStateData `json:"data"`
}
type SendHomeStoreUpdateStoreStateData struct {
}
