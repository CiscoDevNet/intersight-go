/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2025062323
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// NetworkDiscoveredNeighborResponse - The response body of a HTTP GET request for the 'network.DiscoveredNeighbor' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'network.DiscoveredNeighbor' resources.
type NetworkDiscoveredNeighborResponse struct {
	MoAggregateTransform          *MoAggregateTransform
	MoDocumentCount               *MoDocumentCount
	MoTagSummary                  *MoTagSummary
	NetworkDiscoveredNeighborList *NetworkDiscoveredNeighborList
}

// MoAggregateTransformAsNetworkDiscoveredNeighborResponse is a convenience function that returns MoAggregateTransform wrapped in NetworkDiscoveredNeighborResponse
func MoAggregateTransformAsNetworkDiscoveredNeighborResponse(v *MoAggregateTransform) NetworkDiscoveredNeighborResponse {
	return NetworkDiscoveredNeighborResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNetworkDiscoveredNeighborResponse is a convenience function that returns MoDocumentCount wrapped in NetworkDiscoveredNeighborResponse
func MoDocumentCountAsNetworkDiscoveredNeighborResponse(v *MoDocumentCount) NetworkDiscoveredNeighborResponse {
	return NetworkDiscoveredNeighborResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNetworkDiscoveredNeighborResponse is a convenience function that returns MoTagSummary wrapped in NetworkDiscoveredNeighborResponse
func MoTagSummaryAsNetworkDiscoveredNeighborResponse(v *MoTagSummary) NetworkDiscoveredNeighborResponse {
	return NetworkDiscoveredNeighborResponse{
		MoTagSummary: v,
	}
}

// NetworkDiscoveredNeighborListAsNetworkDiscoveredNeighborResponse is a convenience function that returns NetworkDiscoveredNeighborList wrapped in NetworkDiscoveredNeighborResponse
func NetworkDiscoveredNeighborListAsNetworkDiscoveredNeighborResponse(v *NetworkDiscoveredNeighborList) NetworkDiscoveredNeighborResponse {
	return NetworkDiscoveredNeighborResponse{
		NetworkDiscoveredNeighborList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NetworkDiscoveredNeighborResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("failed to unmarshal NetworkDiscoveredNeighborResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("failed to unmarshal NetworkDiscoveredNeighborResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("failed to unmarshal NetworkDiscoveredNeighborResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'network.DiscoveredNeighbor.List'
	if jsonDict["ObjectType"] == "network.DiscoveredNeighbor.List" {
		// try to unmarshal JSON data into NetworkDiscoveredNeighborList
		err = json.Unmarshal(data, &dst.NetworkDiscoveredNeighborList)
		if err == nil {
			return nil // data stored in dst.NetworkDiscoveredNeighborList, return on the first match
		} else {
			dst.NetworkDiscoveredNeighborList = nil
			return fmt.Errorf("failed to unmarshal NetworkDiscoveredNeighborResponse as NetworkDiscoveredNeighborList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NetworkDiscoveredNeighborResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NetworkDiscoveredNeighborList != nil {
		return json.Marshal(&src.NetworkDiscoveredNeighborList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NetworkDiscoveredNeighborResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	if obj.NetworkDiscoveredNeighborList != nil {
		return obj.NetworkDiscoveredNeighborList
	}

	// all schemas are nil
	return nil
}

type NullableNetworkDiscoveredNeighborResponse struct {
	value *NetworkDiscoveredNeighborResponse
	isSet bool
}

func (v NullableNetworkDiscoveredNeighborResponse) Get() *NetworkDiscoveredNeighborResponse {
	return v.value
}

func (v *NullableNetworkDiscoveredNeighborResponse) Set(val *NetworkDiscoveredNeighborResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDiscoveredNeighborResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDiscoveredNeighborResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDiscoveredNeighborResponse(val *NetworkDiscoveredNeighborResponse) *NullableNetworkDiscoveredNeighborResponse {
	return &NullableNetworkDiscoveredNeighborResponse{value: val, isSet: true}
}

func (v NullableNetworkDiscoveredNeighborResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDiscoveredNeighborResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
