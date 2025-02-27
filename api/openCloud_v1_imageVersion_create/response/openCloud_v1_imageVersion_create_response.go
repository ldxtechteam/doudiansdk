package openCloud_v1_imageVersion_create_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type OpenCloudV1ImageVersionCreateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OpenCloudV1ImageVersionCreateData `json:"data"`
}
type OpenCloudV1ImageVersionCreateData struct {
	// 镜像版本id
	VersionId string `json:"version_id"`
}
