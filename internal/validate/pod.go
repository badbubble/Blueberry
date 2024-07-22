package validate

import (
	"Blueberry/internal/model"
	"errors"
)



func PodCreate(pod *model.Pod) error {
	if pod.Base.Name == "" {
		return errors.New("no pod's name ")
	}
	if len(pod.Containers) == 0 {
		return errors.New("pod should at least have one container")
	}
	return nil
}
