package skyutl

import (
	"context"

	"github.com/nkocsea/skylib_skylog/skylog"
)

//CheckDeadline function
func CheckDeadline(ctx context.Context, funcName string, ret ...interface{}) interface{} {
	if ctx.Err() == context.DeadlineExceeded {
		skylog.Infof("%v function is DeadlineExceeded", funcName)
		return ret
	}
	return nil
}
