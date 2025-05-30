/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the GetDiscoveryResultsResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDiscoveryResultsResponseV1{}

// GetDiscoveryResultsResponseV1 Response for GetDiscoveryResults endpoint.
type GetDiscoveryResultsResponseV1 struct {
	// Definitions of all MSP defined on the channel
	Msps map[string]GetDiscoveryResultsResponseV1MspsValue `json:"msps"`
	// List of orderers defined on the channel
	Orderers map[string]GetDiscoveryResultsResponseV1OrderersValue `json:"orderers"`
	// List of peers organized by owner MSP.
	PeersByMSP map[string]GetDiscoveryResultsResponseV1PeersByMSPValue `json:"peersByMSP"`
	Timestamp float32 `json:"timestamp"`
}

// NewGetDiscoveryResultsResponseV1 instantiates a new GetDiscoveryResultsResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDiscoveryResultsResponseV1(msps map[string]GetDiscoveryResultsResponseV1MspsValue, orderers map[string]GetDiscoveryResultsResponseV1OrderersValue, peersByMSP map[string]GetDiscoveryResultsResponseV1PeersByMSPValue, timestamp float32) *GetDiscoveryResultsResponseV1 {
	this := GetDiscoveryResultsResponseV1{}
	this.Msps = msps
	this.Orderers = orderers
	this.PeersByMSP = peersByMSP
	this.Timestamp = timestamp
	return &this
}

// NewGetDiscoveryResultsResponseV1WithDefaults instantiates a new GetDiscoveryResultsResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDiscoveryResultsResponseV1WithDefaults() *GetDiscoveryResultsResponseV1 {
	this := GetDiscoveryResultsResponseV1{}
	return &this
}

// GetMsps returns the Msps field value
func (o *GetDiscoveryResultsResponseV1) GetMsps() map[string]GetDiscoveryResultsResponseV1MspsValue {
	if o == nil {
		var ret map[string]GetDiscoveryResultsResponseV1MspsValue
		return ret
	}

	return o.Msps
}

// GetMspsOk returns a tuple with the Msps field value
// and a boolean to check if the value has been set.
func (o *GetDiscoveryResultsResponseV1) GetMspsOk() (*map[string]GetDiscoveryResultsResponseV1MspsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msps, true
}

// SetMsps sets field value
func (o *GetDiscoveryResultsResponseV1) SetMsps(v map[string]GetDiscoveryResultsResponseV1MspsValue) {
	o.Msps = v
}

// GetOrderers returns the Orderers field value
func (o *GetDiscoveryResultsResponseV1) GetOrderers() map[string]GetDiscoveryResultsResponseV1OrderersValue {
	if o == nil {
		var ret map[string]GetDiscoveryResultsResponseV1OrderersValue
		return ret
	}

	return o.Orderers
}

// GetOrderersOk returns a tuple with the Orderers field value
// and a boolean to check if the value has been set.
func (o *GetDiscoveryResultsResponseV1) GetOrderersOk() (*map[string]GetDiscoveryResultsResponseV1OrderersValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orderers, true
}

// SetOrderers sets field value
func (o *GetDiscoveryResultsResponseV1) SetOrderers(v map[string]GetDiscoveryResultsResponseV1OrderersValue) {
	o.Orderers = v
}

// GetPeersByMSP returns the PeersByMSP field value
func (o *GetDiscoveryResultsResponseV1) GetPeersByMSP() map[string]GetDiscoveryResultsResponseV1PeersByMSPValue {
	if o == nil {
		var ret map[string]GetDiscoveryResultsResponseV1PeersByMSPValue
		return ret
	}

	return o.PeersByMSP
}

// GetPeersByMSPOk returns a tuple with the PeersByMSP field value
// and a boolean to check if the value has been set.
func (o *GetDiscoveryResultsResponseV1) GetPeersByMSPOk() (*map[string]GetDiscoveryResultsResponseV1PeersByMSPValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeersByMSP, true
}

// SetPeersByMSP sets field value
func (o *GetDiscoveryResultsResponseV1) SetPeersByMSP(v map[string]GetDiscoveryResultsResponseV1PeersByMSPValue) {
	o.PeersByMSP = v
}

// GetTimestamp returns the Timestamp field value
func (o *GetDiscoveryResultsResponseV1) GetTimestamp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetDiscoveryResultsResponseV1) GetTimestampOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetDiscoveryResultsResponseV1) SetTimestamp(v float32) {
	o.Timestamp = v
}

func (o GetDiscoveryResultsResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDiscoveryResultsResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msps"] = o.Msps
	toSerialize["orderers"] = o.Orderers
	toSerialize["peersByMSP"] = o.PeersByMSP
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableGetDiscoveryResultsResponseV1 struct {
	value *GetDiscoveryResultsResponseV1
	isSet bool
}

func (v NullableGetDiscoveryResultsResponseV1) Get() *GetDiscoveryResultsResponseV1 {
	return v.value
}

func (v *NullableGetDiscoveryResultsResponseV1) Set(val *GetDiscoveryResultsResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDiscoveryResultsResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDiscoveryResultsResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDiscoveryResultsResponseV1(val *GetDiscoveryResultsResponseV1) *NullableGetDiscoveryResultsResponseV1 {
	return &NullableGetDiscoveryResultsResponseV1{value: val, isSet: true}
}

func (v NullableGetDiscoveryResultsResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDiscoveryResultsResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


