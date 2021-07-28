package oauth

import (
	"encoding/json"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model"
)

// AccessTokenRequest 获取Access Token请求数据结构
type AccessTokenRequest struct {
	AppId        string `json:"app_id,omitempty"`        // 开发者申请的应用APP_ID
	Secret       string `json:"secret,omitempty"`        // 开发者应用的私钥Secret
	GrantType    string `json:"grant_type,omitempty"`    // 授权类型
	AuthCode     string `json:"auth_code,omitempty"`     // 授权码
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新token
}

// AccessTokenResponse 获取Access Token返回数据结构
type AccessTokenResponse struct {
	model.BaseResponse
	Data AccessTokenResponseData `json:"data,omitempty"`
}

// AccessTokenResponseData Access Token具体数据
type AccessTokenResponseData struct {
	AccessToken           string   `json:"access_token,omitempty"`             // 用于验证权限的token
	ExpiresIn             int      `json:"expires_in,omitempty"`               // access_token剩余有效时间,单位(秒)
	RefreshToken          string   `json:"refresh_token,omitempty"`            // 刷新access_token
	RefreshTokenExpiresIn uint64   `json:"refresh_token_expires_in,omitempty"` // refresh_token剩余有效时间,单位(秒)
	AdvertiserIds         []uint64 `json:"advertiser_ids,omitempty"`           //授权的账户id列表
}

func (r AccessTokenRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}