package skyutl

import (
	"context"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

type LoginInfo struct {
	UserId       int64
	CompanyId    int64
	BranchId     int64
	DepartmentId int64
	PartnerId    int64
	PartnerName    string
	DiffHour    float64
	DeviceId     int64
	AccountType  int32
}

// GetLoginAccessToken function return accessToken, suffix token from context and jwt manager
func GetLoginAccessToken(ctx context.Context) (string, string, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return "", "", Unauthenticated
	}

	auths := md["authorization"]
	if len(auths) == 0 {
		return "", "", Unauthenticated
	}

	prefix := "Bearer "
	splitter := "|||"
	suffix := ""
	accessToken := strings.Replace(auths[0], prefix, "", 1)
	index := strings.Index(accessToken, splitter)
	if index >= 0 {
		suffix = accessToken[len(splitter)+index:]
		accessToken = accessToken[:index]
	}
	return accessToken, suffix, nil
}

func getTokenSuffix(accessToken string) (string, string, error) {
	splitter := "|||"
	suffix := ""
	index := strings.Index(accessToken, splitter)
	if index >= 0 {
		suffix = accessToken[len(splitter)+index:]
		accessToken = accessToken[:index]
	}
	return accessToken, suffix, nil
}

// GetUserClaims function return user id from context and jwt manager
func GetUserClaims(ctx context.Context) (*jwt.MapClaims, error) {
	accessToken, _, err := GetLoginAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	return GetUserClaimsFromToken(accessToken)
}

// ToAccount convert from user claims to Account
func ToAccount(claims *jwt.MapClaims) Account {
	userID, _ := ToInt64((*claims)["userId"])
	partnerId, _ := ToInt64((*claims)["partnerId"])
	partnerName, _ := ToString((*claims)["partnerName"])
	diffHour, _ := ToFloat64((*claims)["diffHour"])
	username, _ := ToString((*claims)["username"])
	fullName, _ := ToString((*claims)["fullName"])
	deviceId, _ := ToInt64((*claims)["deviceId"])
	accountType, _ := ToInt32((*claims)["accountType"])

	return Account{
		Id:          userID,
		PartnerId:   partnerId,
		PartnerName:   partnerName,
		DiffHour:   diffHour,
		Username:    &username,
		FullName:    fullName,
		DeviceId:    deviceId,
		AccountType: accountType,
	}
}

// GetAccountInfo function return Account from context and jwt manager
func GetAccountInfo(ctx context.Context) (Account, error) {
	userClaims, err := GetUserClaims(ctx)

	if err != nil {
		return Account{}, err
	}

	return ToAccount(userClaims), nil
}

// GetAccountInfoFromToken function return Account from token
func GetAccountInfoFromToken(token string) (Account, error) {
	userClaims, err := GetUserClaimsFromToken(token)
	if err != nil {
		return Account{}, err
	}

	return ToAccount(userClaims), nil
}

// GetLoginInfo function return userID, companyID, branchID, departmentID
func GetLoginInfo(ctx context.Context) (int64, int64, int64, int64, error) {
	accessToken, suffix, err := GetLoginAccessToken(ctx)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	var companyID, branchID, departmentID int64
	if len(suffix) > 0 {
		parts := strings.Split(suffix, "|")
		if len(parts) >= 3 {
			companyID, _ = ToInt64(parts[0])
			branchID, _ = ToInt64(parts[1])
			departmentID, _ = ToInt64(parts[2])
		}
	}

	userClaims, err := GetUserClaimsFromToken(accessToken)
	if err != nil {
		return 0, 0, 0, 0, Unauthenticated
	}
	userID, err := ToInt64((*userClaims)["userId"])
	if err != nil {
		return 0, 0, 0, 0, Unauthenticated
	}

	return userID, companyID, branchID, departmentID, nil
}

// GetLoginInfoV2 function return struct { UserID, CompanyID, BranchID, DepartmentID}
func GetLoginInfoV2(ctx context.Context) (LoginInfo, error) {
	accessToken, suffix, err := GetLoginAccessToken(ctx)
	if err != nil {
		return LoginInfo{}, err
	}

	var companyID, branchID, departmentID int64
	var diffHour float64
	if len(suffix) > 0 {
		parts := strings.Split(suffix, "|")
		if len(parts) >= 3 {
			companyID, _ = ToInt64(parts[0])
			branchID, _ = ToInt64(parts[1])
			departmentID, _ = ToInt64(parts[2])

			if len(parts) >= 4 {
				diffHour, _ = ToFloat64(parts[3])
			}
		}
	}

	userClaims, err := GetUserClaimsFromToken(accessToken)
	if err != nil {
		return LoginInfo{}, Unauthenticated
	}
	userID, err := ToInt64((*userClaims)["userId"])
	if err != nil {
		return LoginInfo{}, Unauthenticated
	}
	partnerId, _ := ToInt64((*userClaims)["partnerId"])
	partnerName, _ := ToString((*userClaims)["partnerName"])
	deviceId, _ := ToInt64((*userClaims)["deviceId"])
	accountType, _ := ToInt32((*userClaims)["accountType"])
	

	return LoginInfo{
		UserId:       userID,
		CompanyId:    companyID,
		BranchId:     branchID,
		DepartmentId: departmentID,
		PartnerId:    partnerId,
		PartnerName:    partnerName,
		DiffHour:    diffHour,
		DeviceId:     deviceId,
		AccountType:  accountType,
	}, nil
}

func DecodeToken(accessToken string) (LoginInfo, error) {
	_, suffix, err := getTokenSuffix(accessToken)
	
	if err != nil {
		return LoginInfo{}, err
	}

	var companyID, branchID, departmentID int64
	var diffHour float64
	if len(suffix) > 0 {
		parts := strings.Split(suffix, "|")
		if len(parts) >= 3 {
			companyID, _ = ToInt64(parts[0])
			branchID, _ = ToInt64(parts[1])
			departmentID, _ = ToInt64(parts[2])

			if len(parts) >= 4 {
				diffHour, _ = ToFloat64(parts[3])
			}
		}
	}

	userClaims, err := GetUserClaimsFromToken(accessToken)
	if err != nil {
		return LoginInfo{}, Unauthenticated
	}
	userID, err := ToInt64((*userClaims)["userId"])
	if err != nil {
		return LoginInfo{}, Unauthenticated
	}
	partnerId, _ := ToInt64((*userClaims)["partnerId"])
	partnerName, _ := ToString((*userClaims)["partnerName"])
	deviceId, _ := ToInt64((*userClaims)["deviceId"])
	accountType, _ := ToInt32((*userClaims)["accountType"])

	return LoginInfo{
		UserId:       userID,
		CompanyId:    companyID,
		BranchId:     branchID,
		DepartmentId: departmentID,
		PartnerId:    partnerId,
		PartnerName:    partnerName,
		DiffHour:    diffHour,
		DeviceId:     deviceId,
		AccountType:  accountType,
	}, nil
}

// GetUserID function return user id from context
func GetUserID(ctx context.Context) (int64, error) {
	userClaims, err := GetUserClaims(ctx)

	if err != nil {
		return 0, err
	}

	userID, _ := ToInt64((*userClaims)["userId"])
	return userID, nil
}

// GetUserIDFromToken function return user id from token
func GetUserIDFromToken(token string) (int64, error) {
	userClaims, err := GetUserClaimsFromToken(token)
	if err != nil {
		return 0, err
	}

	return ToInt64((*userClaims)["userId"])
}

// GetUserClaimsFromToken function return user claims from token
func GetUserClaimsFromToken(token string) (*jwt.MapClaims, error) {
	if len(token) == 0 {
		return nil, BadRequest
	}

	return JwtManagerInstance.Verify(token)
}
