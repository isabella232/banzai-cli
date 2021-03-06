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
// BackupOptions struct for BackupOptions
type BackupOptions struct {
	IncludedNamespaces []string `json:"includedNamespaces,omitempty"`
	IncludedResources []string `json:"includedResources,omitempty"`
	ExcludedNamespaces []string `json:"excludedNamespaces,omitempty"`
	ExcludedResources []string `json:"excludedResources,omitempty"`
	SnapshotVolumes bool `json:"snapshotVolumes,omitempty"`
	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`
}
