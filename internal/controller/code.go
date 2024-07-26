package controller

type RespCode int64

const (
	CodeSuccess RespCode = 1000 + iota
	CodeServerBusy
	CodeParameterError
	CodeInvalidParameter
	CodeCreatePodError
	CodeGetPodDetailError
	CodeGetPodListError
	CodeGetNodeDetailError
	CodeGetNodeListError
	CodeUpdateNodeLabelsError
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:               "Success",
	CodeServerBusy:            "The server is busy. Please try again later.",
	CodeParameterError:        "The parameter you sent is not correct.",
	CodeInvalidParameter:      "Some key parameters have not been set.",
	CodeCreatePodError:        "Fail to create a Pod",
	CodeGetPodDetailError:     "Fail to get the pod's information.",
	CodeGetPodListError:       "Fail to get a list of pods",
	CodeGetNodeDetailError:    "Fail to get the node's information.",
	CodeGetNodeListError:      "Fail to get a list of nodes.",
	CodeUpdateNodeLabelsError: "Fail to update labels of the node.",
}

func (r RespCode) Msg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
