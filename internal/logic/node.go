package logic

import (
	"Blueberry/internal/model"
	"Blueberry/pkg/k8s"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
)

const (
	PatchReplaceKey   = "$patch"
	PatchReplaceValue = "replace"
)

func GetNodeList() ([]*model.Node, error) {
	nodeList := make([]*model.Node, 0)
	ctx := context.TODO()
	k8sNodeList, err := k8s.Client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range k8sNodeList.Items {
		n := &model.Node{}
		n.FillWithK8sNode(&item)
		nodeList = append(nodeList, n)
	}
	return nodeList, nil
}

func GetNodeDetail(nodeName string) (*model.Node, error) {
	ctx := context.TODO()
	node := &model.Node{}
	k8sNode, err := k8s.Client.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	node.FillWithK8sNodeDetail(k8sNode)
	return node, nil
}

func UpdateNodeLabels(labelsUpdate *model.NodeLabelsUpdate) error {
	labelMap := make(map[string]string)
	for _, label := range labelsUpdate.Labels {
		labelMap[label.Key] = label.Value
	}
	labelMap[PatchReplaceKey] = PatchReplaceValue
	patchData := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": labelMap,
		},
	}
	patchDataBytes, err := json.Marshal(&patchData)
	if err != nil {
		return err
	}
	_, err = k8s.Client.CoreV1().Nodes().Patch(
		context.TODO(),
		labelsUpdate.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func UpdateNodeTaints(taintsUpdate *model.NodeTaintUpdate) error {

	patchData := map[string]interface{}{
		"spec": map[string]interface{}{
			"taints": taintsUpdate.Taints,
		},
	}
	patchDataBytes, err := json.Marshal(&patchData)
	if err != nil {
		return err
	}
	_, err = k8s.Client.CoreV1().Nodes().Patch(
		context.TODO(),
		taintsUpdate.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}
