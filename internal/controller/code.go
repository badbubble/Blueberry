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
	CodeCreateConfigMapError
	CodeUpdateConfigMapError
	DeleteConfigMapError
	CodeGetConfigMapListError
	CodeGetConfigMapDetailError
	CodeCreateSecretError
	CodeDeleteSecretError
	CodeGetSecretDetailError
	CodeGetSecretListError
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:                 "Success",
	CodeServerBusy:              "The server is busy. Please try again later.",
	CodeParameterError:          "The parameter you sent is not correct.",
	CodeInvalidParameter:        "Some key parameters have not been set.",
	CodeCreatePodError:          "Fail to create a Pod",
	CodeGetPodDetailError:       "Fail to get the pod's information.",
	CodeGetPodListError:         "Fail to get a list of pods",
	CodeGetNodeDetailError:      "Fail to get the node's information.",
	CodeGetNodeListError:        "Fail to get a list of nodes.",
	CodeUpdateNodeLabelsError:   "Fail to update labels of the node.",
	CodeCreateConfigMapError:    "Fail to create a ConfigMap.",
	CodeUpdateConfigMapError:    "Fail to update a ConfigMap.",
	DeleteConfigMapError:        "Fail to delete a ConfigMap",
	CodeGetConfigMapListError:   "Fail to get a list of ConfigMaps",
	CodeGetConfigMapDetailError: "Fail to get the detail of a ConfigMap",
	CodeCreateSecretError:       "Fail to create a Secret.",
	CodeDeleteSecretError:       "Fail to Delete a Secret.",
	CodeGetSecretDetailError:    "Fail to Get secret detail",
	CodeGetSecretListError:      "Fail to Get a list of secrets.",
}

func (r RespCode) Msg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
