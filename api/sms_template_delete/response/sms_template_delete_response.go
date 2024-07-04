package sms_template_delete_response

type SmsTemplateDeleteResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsTemplateDeleteData `json:"data"`
}
type SmsTemplateDeleteData struct {
	// 是否成功 0表示成功
	Code int64 `json:"code"`
	// 说明
	Message string `json:"message"`
}
