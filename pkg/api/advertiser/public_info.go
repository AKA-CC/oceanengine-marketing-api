package advertiser

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model/advertiser"
)

func PublicInfo(clt *core.SdkClient, accessToken string, req *advertiser.PublicInfoRequest) (*advertiser.PublicInfoResponse, error) {
	var resp advertiser.PublicInfoResponse
	err := clt.Get("2/advertiser/public_info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
