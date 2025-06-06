/*
Hyperledger Cactus Plugin - Connector Xdai

Can perform basic tasks on a Xdai ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-xdai

import (
	"encoding/json"
)

// checks if the DeployContractV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployContractV1Response{}

// DeployContractV1Response struct for DeployContractV1Response
type DeployContractV1Response struct {
	TransactionReceipt Web3TransactionReceipt `json:"transactionReceipt"`
}

// NewDeployContractV1Response instantiates a new DeployContractV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployContractV1Response(transactionReceipt Web3TransactionReceipt) *DeployContractV1Response {
	this := DeployContractV1Response{}
	this.TransactionReceipt = transactionReceipt
	return &this
}

// NewDeployContractV1ResponseWithDefaults instantiates a new DeployContractV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployContractV1ResponseWithDefaults() *DeployContractV1Response {
	this := DeployContractV1Response{}
	return &this
}

// GetTransactionReceipt returns the TransactionReceipt field value
func (o *DeployContractV1Response) GetTransactionReceipt() Web3TransactionReceipt {
	if o == nil {
		var ret Web3TransactionReceipt
		return ret
	}

	return o.TransactionReceipt
}

// GetTransactionReceiptOk returns a tuple with the TransactionReceipt field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Response) GetTransactionReceiptOk() (*Web3TransactionReceipt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionReceipt, true
}

// SetTransactionReceipt sets field value
func (o *DeployContractV1Response) SetTransactionReceipt(v Web3TransactionReceipt) {
	o.TransactionReceipt = v
}

func (o DeployContractV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployContractV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionReceipt"] = o.TransactionReceipt
	return toSerialize, nil
}

type NullableDeployContractV1Response struct {
	value *DeployContractV1Response
	isSet bool
}

func (v NullableDeployContractV1Response) Get() *DeployContractV1Response {
	return v.value
}

func (v *NullableDeployContractV1Response) Set(val *DeployContractV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployContractV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployContractV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployContractV1Response(val *DeployContractV1Response) *NullableDeployContractV1Response {
	return &NullableDeployContractV1Response{value: val, isSet: true}
}

func (v NullableDeployContractV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployContractV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


