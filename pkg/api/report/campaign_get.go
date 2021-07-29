package report

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model/report"
)

func CampaignGet(clt *core.SdkClient, accessToken string, req *report.CampaignGetRequest) (*report.CampaignGetResponse, error) {
	var resp report.CampaignGetResponse
	err := clt.Get("2/report/campaign/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
