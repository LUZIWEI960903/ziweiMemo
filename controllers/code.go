package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
	CodeOverdueToken
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:         "Success!!",
	CodeInvalidParam:    "Require params failed!!",
	CodeUserExist:       "User exist!!",
	CodeUserNotExist:    "User not exist!!",
	CodeInvalidPassword: "Username or password error!!",
	CodeServerBusy:      "Server busy!!",
	CodeNeedLogin:       "Please login in first!!",
	CodeInvalidToken:    "Invalid token!!",
	CodeOverdueToken:    "Token overdue!!",
}

func (code ResCode) Msg() string {
	msg, ok := CodeMsgMap[code]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
