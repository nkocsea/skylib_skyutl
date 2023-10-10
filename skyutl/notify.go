package skyutl

import (
	"context"
)

func SendNotify(ctx context.Context, urlOrServiceAddr string) error {
	accessToken, _, _ := GetLoginAccessToken(ctx)
	_, err := RestPost(urlOrServiceAddr, "/core/notification/v1/send", map[string]interface{}{}, accessToken)
	if err != nil {
		return err
	}

	return nil
}
