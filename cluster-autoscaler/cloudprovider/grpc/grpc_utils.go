package grpccloudprovider

import (
	"encoding/json"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
	//	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
)

// NodeFromJSON deserialize a string to apiv1.Node
func NodeFromJSON(s string) (*apiv1.Node, error) {
	data := &apiv1.Node{}

	err := json.Unmarshal([]byte(s), &data)

	return data, err
}

// PodFromJSON deserialize a string to apiv1.Pod
func PodFromJSON(s string) (*apiv1.Pod, error) {
	data := &apiv1.Pod{}

	err := json.Unmarshal([]byte(s), &data)

	return data, err
}

func nodesAsJSON(nodes []*apiv1.Node) []string {
	if nodes == nil {
		return nil
	}

	r := make([]string, len(nodes))

	for i, node := range nodes {
		b, _ := json.Marshal(node)
		r[i] = string(b)
	}

	return r
}

func fromJSON(v string) struct{} {
	var r struct{}

	json.Unmarshal([]byte(v), &r)

	return r
}

func toJSON(v interface{}) string {
	if v == nil {
		return ""
	}

	b, _ := json.Marshal(v)

	return string(b)
}

// Transform grpccloudprovier.NodeInfo to schedulercache.NodeInfo
func newNodeInfo(nodeInfo *NodeInfo) *schedulerframework.NodeInfo {
	schedulerNodeInfo := schedulerframework.NewNodeInfo()
	node, err := NodeFromJSON(nodeInfo.GetNode())

	if err == nil {
		schedulerNodeInfo.SetNode(node)
	}

	for _, s := range nodeInfo.GetPods() {
		pod, err := PodFromJSON(s)

		if err == nil {
			schedulerNodeInfo.AddPod(pod)
		}
	}

	return schedulerNodeInfo
}

func toResourceLimiter(rl *cloudprovider.ResourceLimiter) *ResourceLimiter {
	resourceLimiter := &ResourceLimiter{}

	for _, resourceName := range rl.GetResources() {
		if rl.HasMinLimitSet(resourceName) {
			if resourceLimiter.MinLimits == nil {
				resourceLimiter.MinLimits = make(map[string]int64)
			}

			resourceLimiter.MinLimits[resourceName] = rl.GetMin(resourceName)
		}

		if rl.HasMaxLimitSet(resourceName) {
			if resourceLimiter.MaxLimits == nil {
				resourceLimiter.MaxLimits = make(map[string]int64)
			}

			resourceLimiter.MaxLimits[resourceName] = rl.GetMax(resourceName)
		}
	}

	return resourceLimiter
}
