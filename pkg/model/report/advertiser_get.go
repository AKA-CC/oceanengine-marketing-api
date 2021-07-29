package report

import (
	"encoding/json"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/enum"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model"
	"net/url"
	"strconv"
	"time"
)

type AdvertiserGetRequest struct {
	AdvertiserId    uint64                   `json:"advertiser_id,omitempty"`
	StartDate       time.Time                `json:"start_date,omitempty"`       // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate         time.Time                `json:"end_date,omitempty"`         // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	Fields          []string                 `json:"fields,omitempty"`           // 指定需要的指标名称
	GroupBy         []enum.StatGroupBy       `json:"group_by,omitempty"`         // 分组条件默认为 STAT_GROUP_BY_FIELD_STAT_TIME
	TimeGranularity enum.StatTimeGranularity `json:"time_granularity,omitempty"` // 时间粒度, 默认值: STAT_TIME_GRANULARITY_DAILY
	Page            int                      `json:"page,omitempty"`             // 页码；默认值: 1
	PageSize        int                      `json:"page_size,omitempty"`        // 页面大小，即每页展示的数据量；默认值: 20；取值范围: 1-1000
}

type AdvertiserGetResponse struct {
	model.BaseResponse
	Data *AdvertiserGetResponseData `json:"data,omitempty"`
}

type AdvertiserGetResponseData struct {
	List     []AdvertiserGetResponseItem `json:"list,omitempty"`
	PageInfo *model.PageInfo             `json:"page_info,omitempty"`
}

type AdvertiserGetResponseItem struct {
	AdvertiserId uint64 `json:"advertiser_id,omitempty"` // 广告主ID
	BaseResponseData
}

func (r AdvertiserGetRequest) Encode() string {
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
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}
