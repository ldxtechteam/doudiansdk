package open_binaryupload_response

type OpenBinaryuploadResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OpenBinaryuploadData `json:"data"`
}
type OpenBinaryuploadData struct {
	// 上传二进制文件的uri
	StoreUri string `json:"store_uri"`
}
