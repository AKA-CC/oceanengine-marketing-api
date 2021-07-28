package oauth

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model/oauth"
)

// AccessToken 获取Access Token
func AccessToken(clt *core.SdkClient, authCode string) (*oauth.AccessTokenResponse, error) {
	req := &oauth.AccessTokenRequest{
		AppId:     clt.AppId,
		Secret:    clt.Secret,
		GrantType: "auth_code",
		AuthCode:  authCode,
	}

	var resp oauth.AccessTokenResponse
	err := clt.Post("oauth2/access_token/", req, &resp, "")
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
