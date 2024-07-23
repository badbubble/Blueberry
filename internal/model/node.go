package model

import corev1 "k8s.io/api/core/v1"

type Node struct {
	Name             string         `json:"name"`
	Status           string         `json:"status"`
	Age              int64          `json:"age"`
	InternalIp       string         `json:"internalIp"`
	ExternalIp       string         `json:"externalIp"`
	Version          string         `json:"version"`
	OSImage          string         `json:"osImage"`
	KernelVersion    string         `json:"kernelVersion"`
	ContainerRuntime string         `json:"containerRuntime"`
	Labels           []ListMapItem  `json:"labels"`
	Taints           []corev1.Taint `json:"taints"`
}

func (n *Node) FillWithK8sNode(node *corev1.Node) {
	n.Name = node.Name
	n.Age = node.CreationTimestamp.Unix()
	n.Version = node.Status.NodeInfo.KubeletVersion
	n.KernelVersion = node.Status.NodeInfo.KernelVersion
	n.ContainerRuntime = node.Status.NodeInfo.ContainerRuntimeVersion
	n.OSImage = node.Status.NodeInfo.OSImage

}
