package controller

type RespCode int64

// custom response code definitions added here
const (
	CodeSuccess RespCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:         "Success",
	CodeInvalidParam:    "Invalid parameter",
	CodeUserExist:       "User already exists",
	CodeUserNotExist:    "User does not exist",
	CodeInvalidPassword: "Invalid password",
	CodeServerBusy:      "Server is busy",

	CodeInvalidToken: "Invalid authentication",
	CodeNeedLogin:    "Login required",
}

func (code RespCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
