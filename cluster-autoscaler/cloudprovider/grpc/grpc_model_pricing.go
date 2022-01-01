/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpccloudprovider

import (
	"context"
	"log"
	"time"

	apiv1 "k8s.io/api/core/v1"
	errors "k8s.io/autoscaler/cluster-autoscaler/utils/errors"
)

// GrpcPriceModel implements PriceModel interface for GCE.
type GrpcPriceModel struct {
	name    string
	manager *GrpcManager
}

// NodePrice returns a price of running the given node for a given period of time.
// All prices are in USD.
func (model *GrpcPriceModel) NodePrice(node *apiv1.Node, startTime time.Time, endTime time.Time) (float64, error) {
	manager := model.manager

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	pricingModelService, err := manager.GetPricingModelServiceClient()

	if err == nil {
		r, err := pricingModelService.NodePrice(ctx, &NodePriceRequest{ProviderID: manager.GetCloudProviderID(), Node: toJSON(node), StartTime: startTime.Unix(), EndTime: endTime.Unix()})

		if err != nil {
			log.Printf("Could not get PriceModel::NodePrice for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return 0, err
		} else if rerr := r.GetError(); rerr != nil {
			log.Printf("Cloud provider:%s call PriceModel::NodePrice got error: %v", manager.GetCloudProviderID(), rerr)
			return 0, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return r.GetPrice(), nil
	}

	return 0, err
}

// PodPrice returns a theoretical minimum price of running a pod for a given
// period of time on a perfectly matching machine.
func (model *GrpcPriceModel) PodPrice(pod *apiv1.Pod, startTime time.Time, endTime time.Time) (float64, error) {
	manager := model.manager

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	pricingModelService, err := manager.GetPricingModelServiceClient()

	if err == nil {
		r, err := pricingModelService.PodPrice(ctx, &PodPriceRequest{ProviderID: manager.GetCloudProviderID(), Pod: toJSON(pod), StartTime: startTime.Unix(), EndTime: endTime.Unix()})

		if err != nil {
			log.Printf("Could not get PriceModel::PodPrice for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return 0, err
		} else if rerr := r.GetError(); rerr != nil {
			log.Printf("Cloud provider:%s call PriceModel::PodPrice got error: %v", manager.GetCloudProviderID(), rerr)
			return 0, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return r.GetPrice(), nil
	}

	return 0, err
}
