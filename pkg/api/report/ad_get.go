package report

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model/report"
)

func AdGet(clt *core.SdkClient, accessToken string, req *report.AdGetRequest) (*report.AdGetResponse, error) {
	var resp report.AdGetResponse
	err := clt.Get("2/report/ad/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
