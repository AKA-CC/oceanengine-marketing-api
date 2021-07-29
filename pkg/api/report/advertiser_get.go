package report

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model/report"
)

// AdvertiserGet 广告主数据
func AdvertiserGet(clt *core.SdkClient, accessToken string, req *report.AdvertiserGetRequest) (*report.AdvertiserGetResponse, error) {
	var resp report.AdvertiserGetResponse
	err := clt.Get("2/report/advertiser/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
