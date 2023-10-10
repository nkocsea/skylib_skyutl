package skyutl

import (
	"context"

	"suntech.com.vn/skylib/skylog.git/skylog"
)

//CheckDeadline function
func CheckDeadline(ctx context.Context, funcName string, ret ...interface{}) interface{} {
	if ctx.Err() == context.DeadlineExceeded {
		skylog.Infof("%v function is DeadlineExceeded", funcName)
		return ret
	}
	return nil
}
