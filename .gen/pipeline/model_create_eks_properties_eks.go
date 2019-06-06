/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.15.4
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type CreateEksPropertiesEks struct {
	Version string `json:"version,omitempty"`
	NodePools map[string]NodePoolsAmazon `json:"nodePools"`
	Vpc EksVpc `json:"vpc,omitempty"`
	// Id of the RouteTable of the VPC to be used by subnets. This is used only when subnets are created into existing VPC.
	RouteTableId string `json:"routeTableId,omitempty"`
	Subnets []EksSubnet `json:"subnets,omitempty"`
}