package advertiser

import (
	"encoding/json"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model"
	"net/url"
)

type PublicInfoRequest struct {
	AdvertiserIds []uint64 `json:"advertiser_ids,omitempty"` // 广告主ID集合
}

type PublicInfoResponse struct {
	model.BaseResponse
	Data []PublicInfoResponseData `json:"data,omitempty"`
}

type PublicInfoResponseData struct {
	Id                      uint64                `json:"id,omitempty"`                        // 广告主ID
	Name                    string                `json:"name,omitempty"`                      // 账户名
	Company                 string                `json:"company,omitempty"`                   // 公司名
	FirstIndustryName       string                `json:"first_industry_name,omitempty"`       // 一级行业名
	SecondIndustryName      string                `json:"second_industry_name,omitempty"`      // 二级行业名
}

func (r PublicInfoRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIds)
	values := &url.Values{}
	values.Set("advertiser_ids", string(idsBytes))
	return values.Encode()
}