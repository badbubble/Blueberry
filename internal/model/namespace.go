package model

type Namespace struct {
	Name              string `json:"name"`
	CreationTimestamp int64  `json:"CreationTimestamp"`
	Status            string `json:"status"`
}
