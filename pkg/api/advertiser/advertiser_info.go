package advertiser

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model/advertiser"
)

// Info 广告主信息
func Info(clt *core.SdkClient, accessToken string, req *advertiser.InfoRequest) (*advertiser.InfoResponse, error) {
	var resp advertiser.InfoResponse
	err := clt.Get("2/advertiser/info", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
