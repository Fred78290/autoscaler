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

import "time"

// GrpcConfig is handles config.
type GrpcConfig struct {
	Address    string `json:"address"`
	Identifier string `json:"secret"`
	Timeout    int    `json:"timeout"`
}

// GetAddress returns grpc server address
func (t *GrpcConfig) GetAddress() string {
	return t.Address
}

// GetIdentifier returns the cloud provider ID
func (t *GrpcConfig) GetIdentifier() string {
	return t.Identifier
}

// GetTimeout returns the timeout nanoseconds based
func (t *GrpcConfig) GetTimeout() time.Duration {
	return time.Duration(t.Timeout) * time.Second
}
