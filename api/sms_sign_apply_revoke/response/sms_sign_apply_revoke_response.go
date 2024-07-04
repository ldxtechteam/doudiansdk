package sms_sign_apply_revoke_response

type SmsSignApplyRevokeResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsSignApplyRevokeData `json:"data"`
}
type SmsSignApplyRevokeData struct {
	// 是否成功 0表示成功
	Code int64 `json:"code"`
	// 说明
	Message string `json:"message"`
}
