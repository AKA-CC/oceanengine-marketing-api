package oauth

import (
	"fmt"
	"github.com/AKA-CC/oceanengine-marketing-api/pkg/core"
	"net/url"
)

// GrantUrl 授权URL
func GrantUrl(clt *core.SdkClient, redirectUrl string, state string) string {
	values := &url.Values{}
	values.Set("app_id", clt.AppId)
	if state != "" {
		values.Set("state", state)
	}
	values.Set("redirect_url", redirectUrl)
	return fmt.Sprintf("https://ad.oceanengine.com/openapi/audit/oauth.html?%s", values.Encode())
}
