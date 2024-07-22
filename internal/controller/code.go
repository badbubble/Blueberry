package controller

type RespCode int64

const (
	CodeSuccess RespCode = 1000 + iota
	CodeServerBusy
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:    "Success",
	CodeServerBusy: "The server is busy. Please try again later.",
}

func (r RespCode) Msg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
