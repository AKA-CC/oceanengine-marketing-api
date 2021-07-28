package oauth

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model/oauth"
)

// UserInfo 获取授权User信息
func UserInfo(clt *core.SdkClient, accessToken string) (*oauth.UserInfoResponse, error) {
	var resp oauth.UserInfoResponse
	err := clt.Get("2/user/info", nil, &resp, accessToken)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
