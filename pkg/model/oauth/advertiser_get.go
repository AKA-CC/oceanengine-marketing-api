package oauth

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/enum"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model"
	"net/url"
)

type AdvertiserGetRequest struct {
	AccessToken string `json:"access_token,omitempty"` // 授权access_token
	AppId       string `json:"app_id,omitempty"`       // 开发者申请的应用APP_ID
	Secret      string `json:"secret,omitempty"`       // 开发者应用的私钥Secret
}

type AdvertiserGetResponse struct {
	model.BaseResponse
	Data *AdvertiserGetResponseData `json:"data,omitempty"`
}

type AdvertiserGetResponseData struct {
	List []Advertiser `json:"list,omitempty"`
}

type Advertiser struct {
	AdvertiserID   uint64           `json:"advertiser_id,omitempty"`   // 账号id
	AdvertiserName string           `json:"advertiser_name,omitempty"` // 账号名称
	AdvertiserRole uint             `json:"advertiser_role,omitempty"` // 账号角色，1-普通广告主，2-账号管家，3-一级代理商，4-二级代理商
	IsValid        bool             `json:"is_valid,omitempty"`        // 授权有效性，允许值：true/false；false表示对应的user在客户中心/一站式平台代理商平台变更了对此账号的权限,需要到对应平台进行调整过来；
	AccountRole    enum.AccountRole `json:"account_role,omitempty"`    // 新版授权账号角色
}

func (r AdvertiserGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("app_id", r.AppId)
	values.Set("secret", r.Secret)
	values.Set("access_token", r.AccessToken)
	return values.Encode()
}