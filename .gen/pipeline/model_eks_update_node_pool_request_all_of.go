/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// EksUpdateNodePoolRequestAllOf struct for EksUpdateNodePoolRequestAllOf
type EksUpdateNodePoolRequestAllOf struct {
	Autoscaling NodePoolAutoScaling `json:"autoscaling,omitempty"`
	// Size of the EBS volume in GBs of the nodes in the pool.
	VolumeSize int32 `json:"volumeSize,omitempty"`
	// The instance type to use for your node pool.
	InstanceType string `json:"instanceType,omitempty"`
	// The instance AMI to use for your node pool.
	Image string `json:"image,omitempty"`
	// The upper limit price for the requested spot instance. If this field is empty or 0 on-demand instances are used instead of spot instances.
	SpotPrice string `json:"spotPrice,omitempty"`
	Options BaseUpdateNodePoolOptions `json:"options,omitempty"`
}
