package controllers

type resCode int64

const (
	codeSuccess resCode = 1000 + iota
	codeInvalidParam
	codeUserExist
	codeInvalidPassword
	codeServerBusy

	codeNeedLogin
	codeInvalidToken
)

var codeMsgMap = map[resCode]string{
	codeSuccess:         "Success!!",
	codeInvalidParam:    "Require params failed!!",
	codeUserExist:       "User exist!!",
	codeInvalidPassword: "Username or password error!!",
	codeServerBusy:      "Server busy!!",
	codeNeedLogin:       "Please login in first!!",
	codeInvalidToken:    "Invalid token!!",
}

func (code resCode) msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[codeServerBusy]
	}
	return msg
}
