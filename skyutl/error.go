package skyutl

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lingdor/stackerror"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/nkocsea/skylib_skylog/skylog"
)

var (
	//Unauthenticated error
	Unauthenticated = status.Error(codes.Unauthenticated, "SYS.MSG.UNAUTHENTICATED_ERROR")
	//BadRequest error
	BadRequest = status.Error(codes.InvalidArgument, "SYS.MSG.INVALID_ARGUMENT_ERROR")
	//Internal error
	Internal = status.Error(codes.Internal, "SYS.MSG.INTERNAL_ERROR")
	//NeedLogin error
	NeedLogin = status.Error(codes.Unauthenticated, "SYS.MSG.NEED_LOGIN_ERROR")
	//CallRESTAPIError error
	CallRESTAPIError = status.Error(codes.InvalidArgument, "SYS.MSG.INTERNAL_ERROR")
)

// CustomError function
func CustomError(code codes.Code, message string, field string, trace interface{}) error {
	fullMsg, _ := json.Marshal(map[string]interface{}{
		"field":   field,
		"message": message,
	})
	st := status.New(code, string(fullMsg))
	if field != "" {
		var js []byte
		if trace != nil {
			js, _ = json.Marshal(trace)
		}

		v := &errdetails.BadRequest_FieldViolation{
			Field:       field,
			Description: message + string(js),
		}
		br := &errdetails.BadRequest{}
		br.FieldViolations = append(br.FieldViolations, v)
		st, err := st.WithDetails(br)
		if err != nil {
			return err
		}

		return st.Err()
	}

	return st.Err()
}

// Error400 function
func Error400(message string, field string, trace interface{}) error {
	return CustomError(codes.InvalidArgument, message, field, trace)
}

// Error500 function
func Error500(err error) error {
	return status.Error(codes.Internal, err.Error())
}

// ExistedError return existed error with field name
func ExistedError(field string) error {
	return Error400(fmt.Sprintf("SYS.MSG.%v_IS_EXISTED", strings.ToUpper(field)), field, nil)
}

// DuplicatedError return duplicated error with field name
func DuplicatedError(field string) error {
	return Error400(fmt.Sprintf("SYS.MSG.%v_IS_DUPLICATED", strings.ToUpper(field)), field, nil)
}

func CatchError() {
	if err := recover(); err != nil {
		e := stackerror.New("=======CatchError=======")
		skylog.DetailError("CatchError has error: ", err)
		fmt.Println(e.Error())
	}
}

func TxCatchError(tx *sql.Tx) {
	if err := recover(); err != nil {
		e := stackerror.New("=======CatchError=======")
		skylog.DetailError("CatchError has error: ", err)
		fmt.Println(e.Error())
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
