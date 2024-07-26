package model

import corev1 "k8s.io/api/core/v1"

const (
	ConditionTrueStatus = "True"
)
const (
	InternalIPType = "InternalIP"
	ExternalIPType = "ExternalIP"
)

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

type NodeLabelsUpdate struct {
	Name   string        `json:"name"`
	Labels []ListMapItem `json:"labels"`
}

type NodeTaintUpdate struct {
	Name   string         `json:"name"`
	Taints []corev1.Taint `json:"taints"`
}

func (n *Node) GetIP(node *corev1.Node) {
	n.ExternalIp = "<none>"
	n.InternalIp = "<none>"
	var eFlag, iFlag bool

	for _, address := range node.Status.Addresses {
		if address.Type == ExternalIPType && !eFlag {
			n.ExternalIp = address.Address
			eFlag = true
		} else if address.Type == InternalIPType && !iFlag {
			n.InternalIp = address.Address
			iFlag = true
		}
	}
}

func (n *Node) GetStatus(node *corev1.Node) {
	for _, condition := range node.Status.Conditions {
		if condition.Status == ConditionTrueStatus {
			n.Status = string(condition.Type)
		}
	}
}

func (n *Node) GetLabels(node *corev1.Node) {
	labels := make([]ListMapItem, 0)
	for k, v := range node.Labels {
		labels = append(labels, ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	n.Labels = labels
}

func (n *Node) FillWithK8sNodeDetail(node *corev1.Node) {
	n.FillWithK8sNode(node)
	n.GetLabels(node)
	n.Taints = node.Spec.Taints
}

func (n *Node) FillWithK8sNode(node *corev1.Node) {
	n.Name = node.Name
	n.Age = node.CreationTimestamp.Unix()
	n.Version = node.Status.NodeInfo.KubeletVersion
	n.KernelVersion = node.Status.NodeInfo.KernelVersion
	n.ContainerRuntime = node.Status.NodeInfo.ContainerRuntimeVersion
	n.OSImage = node.Status.NodeInfo.OSImage
	n.GetStatus(node)
	n.GetIP(node)
}
