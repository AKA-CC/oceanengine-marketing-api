package advertiser

import (
	"encoding/json"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/enum"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model"
	"net/url"
)

type InfoRequest struct {
	AdvertiserIds []uint64 `json:"advertiser_ids,omitempty"` // 广告主ID集合
	Fields        []string `json:"fields,omitempty"`
}

type InfoResponse struct {
	model.BaseResponse
	Data []InfoResponseData `json:"data,omitempty"`
}

type InfoResponseData struct {
	Id                      uint64                `json:"id,omitempty"`                        // 广告主ID
	Name                    string                `json:"name,omitempty"`                      // 账户名
	Role                    enum.AdvertiserRole   `json:"role,omitempty"`                      // 角色
	Status                  enum.AdvertiserStatus `json:"status,omitempty"`                    // 状态
	Address                 string                `json:"address,omitempty"`                   // 地址
	LicenseUrl              string                `json:"license_url,omitempty"`               // 执照预览地址(链接默认1小时内有效)
	LicenseNo               string                `json:"license_no,omitempty"`                // 执照编号
	LicenseProvince         string                `json:"license_province,omitempty"`          // 执照省份
	LicenseCity             string                `json:"license_city,omitempty"`              // 执照城市
	Company                 string                `json:"company,omitempty"`                   // 公司名
	Brand                   string                `json:"brand,omitempty"`                     // 经营类别
	PromotionArea           string                `json:"promotion_area,omitempty"`            // 运营区域
	PromotionCenterProvince string                `json:"promotion_center_province,omitempty"` // 运营省份
	PromotionCenterCity     string                `json:"promotion_center_city,omitempty"`     // 运营城市
	FirstIndustryName       string                `json:"first_industry_name,omitempty"`       // 一级行业名
	SecondIndustryName      string                `json:"second_industry_name,omitempty"`      // 二级行业名
	Reason                  string                `json:"reason,omitempty"`                    // 审核拒绝原因
	CreateTime              string                `json:"create_time,omitempty"`               // 创建时间
}

func (r InfoRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIds)
	fieldsBytes, _ := json.Marshal(r.Fields)
	values := &url.Values{}
	values.Set("advertiser_ids", string(idsBytes))
	values.Set("fields", string(fieldsBytes))
	return values.Encode()
}