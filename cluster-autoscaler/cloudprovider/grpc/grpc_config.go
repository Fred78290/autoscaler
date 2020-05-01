package grpccloudprovider

import "time"

// GrpcConfig is handles config.
type GrpcConfig struct {
	Address              string         `json:"address"`
	Identifier           string         `json:"secret"`
	Timeout              int            `json:"timeout"`
	KubeAdmConfiguration *KubeAdmConfig `json:"config"`
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
