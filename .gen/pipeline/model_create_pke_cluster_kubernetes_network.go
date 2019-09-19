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

type CreatePkeClusterKubernetesNetwork struct {
	ServiceCIDR string `json:"serviceCIDR,omitempty"`
	PodCIDR string `json:"podCIDR,omitempty"`
	Provider string `json:"provider,omitempty"`
	ProviderConfig map[string]interface{} `json:"providerConfig,omitempty"`
}
