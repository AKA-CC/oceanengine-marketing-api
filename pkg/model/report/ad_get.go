package report

import (
	"encoding/json"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/enum"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model"
	"net/url"
	"strconv"
	"time"
)

type AdGetRequest struct {
	AdvertiserId    uint64                   `json:"advertiser_id,omitempty"`
	StartDate       time.Time                `json:"start_date,omitempty"`       // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate         time.Time                `json:"end_date,omitempty"`         // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	Fields          []string                 `json:"fields,omitempty"`           // 指定需要的指标名称
	GroupBy         []enum.StatGroupBy       `json:"group_by,omitempty"`         // 分组条件默认为 STAT_GROUP_BY_FIELD_STAT_TIME
	TimeGranularity enum.StatTimeGranularity `json:"time_granularity,omitempty"` // 时间粒度, 默认值: STAT_TIME_GRANULARITY_DAILY
	Filtering       *AdGetFiltering          `json:"filtering,omitempty"`        // 筛选字段
	OrderField      string                   `json:"order_field,omitempty"`      // 排序指标字段，不传默认填空
	OrderType       enum.OrderType           `json:"order_type,omitempty"`       // 排序类型；默认值: DESC；允许值: ASC, DESC，不传入order_field时不保证顺序
	Page            int                      `json:"page,omitempty"`             // 页码；默认值: 1
	PageSize        int                      `json:"page_size,omitempty"`        // 页面大小，即每页展示的数据量；默认值: 20；取值范围: 1-1000
}

type AdGetFiltering struct {
	CampaignIds  []uint64            `json:"campaign_ids,omitempty"`  // 广告组 id 列表
	AdName       string              `json:"ad_name,omitempty"`       // 广告名
	AdIds        []uint64            `json:"ad_ids,omitempty"`        // 广告计划 id 列表
	Pricings     []enum.AdPricing    `json:"pricings,omitempty"`      // 出价方式列表
	LandingTypes []enum.LandingType  `json:"landing_types,omitempty"` // 推广目的列表
	Status       enum.CampaignStatus `json:"status,omitempty"`        // 广告组状态
}

type AdGetResponse struct {
	model.BaseResponse
	Data *AdGetResponseData `json:"data,omitempty"`
}

type AdGetResponseData struct {
	List     []AdGetResponseItem `json:"list,omitempty"`
	PageInfo *model.PageInfo     `json:"page_info,omitempty"`
}

type AdGetResponseItem struct {
	CampaignId   uint64 `json:"campaign_id,omitempty"`   // 广告组id，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	CampaignName string `json:"campaign_name,omitempty"` // 广告组名称，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	AdId         uint64 `json:"ad_id,omitempty"`         // 计划id，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	AdName       string `json:"ad_name,omitempty"`       // 计划名称，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	BaseResponseData
}

func (r AdGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserId, 10))
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	if r.Fields != nil {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	if r.GroupBy != nil {
		groupBy, _ := json.Marshal(r.GroupBy)
		values.Set("group_by", string(groupBy))
	}
	if r.TimeGranularity != "" {
		values.Set("time_granularity", string(r.TimeGranularity))
	}
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}
