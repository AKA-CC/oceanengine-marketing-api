package oauth

import (
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/model/oauth"
)

// RefreshToken 刷新Refresh Token
func RefreshToken(clt *core.SdkClient, refreshToken string) (*oauth.AccessTokenResponse, error) {
	req := &oauth.AccessTokenRequest{
		AppId:        clt.AppId,
		Secret:       clt.Secret,
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	}
	var resp oauth.AccessTokenResponse
	err := clt.Post("oauth2/refresh_token/", req, &resp, "")
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
