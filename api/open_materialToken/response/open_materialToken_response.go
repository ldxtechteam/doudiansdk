package open_materialToken_response

type OpenMaterialTokenResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OpenMaterialTokenData `json:"data"`
}
type OpenMaterialTokenData struct {
	// 获取下载地址query
	AuthQuery string `json:"auth_query"`
}
