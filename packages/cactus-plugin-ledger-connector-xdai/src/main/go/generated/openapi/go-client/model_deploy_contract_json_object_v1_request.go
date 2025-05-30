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

// checks if the DeployContractJsonObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployContractJsonObjectV1Request{}

// DeployContractJsonObjectV1Request struct for DeployContractJsonObjectV1Request
type DeployContractJsonObjectV1Request struct {
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	Gas *float32 `json:"gas,omitempty"`
	GasPrice *string `json:"gasPrice,omitempty"`
	// The amount of milliseconds to wait for a transaction receipt with theaddress of the contract(which indicates successful deployment) beforegiving up and crashing.
	TimeoutMs *float32 `json:"timeoutMs,omitempty"`
	ContractJSON ContractJSON `json:"contractJSON"`
	// The list of arguments to pass in to the constructor of the contract being deployed.
	ConstructorArgs []interface{} `json:"constructorArgs,omitempty"`
}

// NewDeployContractJsonObjectV1Request instantiates a new DeployContractJsonObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployContractJsonObjectV1Request(web3SigningCredential Web3SigningCredential, contractJSON ContractJSON) *DeployContractJsonObjectV1Request {
	this := DeployContractJsonObjectV1Request{}
	this.Web3SigningCredential = web3SigningCredential
	var timeoutMs float32 = 60000
	this.TimeoutMs = &timeoutMs
	this.ContractJSON = contractJSON
	return &this
}

// NewDeployContractJsonObjectV1RequestWithDefaults instantiates a new DeployContractJsonObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployContractJsonObjectV1RequestWithDefaults() *DeployContractJsonObjectV1Request {
	this := DeployContractJsonObjectV1Request{}
	var timeoutMs float32 = 60000
	this.TimeoutMs = &timeoutMs
	return &this
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *DeployContractJsonObjectV1Request) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *DeployContractJsonObjectV1Request) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *DeployContractJsonObjectV1Request) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *DeployContractJsonObjectV1Request) GetGas() float32 {
	if o == nil || IsNil(o.Gas) {
		var ret float32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractJsonObjectV1Request) GetGasOk() (*float32, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *DeployContractJsonObjectV1Request) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given float32 and assigns it to the Gas field.
func (o *DeployContractJsonObjectV1Request) SetGas(v float32) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *DeployContractJsonObjectV1Request) GetGasPrice() string {
	if o == nil || IsNil(o.GasPrice) {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractJsonObjectV1Request) GetGasPriceOk() (*string, bool) {
	if o == nil || IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *DeployContractJsonObjectV1Request) HasGasPrice() bool {
	if o != nil && !IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *DeployContractJsonObjectV1Request) SetGasPrice(v string) {
	o.GasPrice = &v
}

// GetTimeoutMs returns the TimeoutMs field value if set, zero value otherwise.
func (o *DeployContractJsonObjectV1Request) GetTimeoutMs() float32 {
	if o == nil || IsNil(o.TimeoutMs) {
		var ret float32
		return ret
	}
	return *o.TimeoutMs
}

// GetTimeoutMsOk returns a tuple with the TimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractJsonObjectV1Request) GetTimeoutMsOk() (*float32, bool) {
	if o == nil || IsNil(o.TimeoutMs) {
		return nil, false
	}
	return o.TimeoutMs, true
}

// HasTimeoutMs returns a boolean if a field has been set.
func (o *DeployContractJsonObjectV1Request) HasTimeoutMs() bool {
	if o != nil && !IsNil(o.TimeoutMs) {
		return true
	}

	return false
}

// SetTimeoutMs gets a reference to the given float32 and assigns it to the TimeoutMs field.
func (o *DeployContractJsonObjectV1Request) SetTimeoutMs(v float32) {
	o.TimeoutMs = &v
}

// GetContractJSON returns the ContractJSON field value
func (o *DeployContractJsonObjectV1Request) GetContractJSON() ContractJSON {
	if o == nil {
		var ret ContractJSON
		return ret
	}

	return o.ContractJSON
}

// GetContractJSONOk returns a tuple with the ContractJSON field value
// and a boolean to check if the value has been set.
func (o *DeployContractJsonObjectV1Request) GetContractJSONOk() (*ContractJSON, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractJSON, true
}

// SetContractJSON sets field value
func (o *DeployContractJsonObjectV1Request) SetContractJSON(v ContractJSON) {
	o.ContractJSON = v
}

// GetConstructorArgs returns the ConstructorArgs field value if set, zero value otherwise.
func (o *DeployContractJsonObjectV1Request) GetConstructorArgs() []interface{} {
	if o == nil || IsNil(o.ConstructorArgs) {
		var ret []interface{}
		return ret
	}
	return o.ConstructorArgs
}

// GetConstructorArgsOk returns a tuple with the ConstructorArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractJsonObjectV1Request) GetConstructorArgsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ConstructorArgs) {
		return nil, false
	}
	return o.ConstructorArgs, true
}

// HasConstructorArgs returns a boolean if a field has been set.
func (o *DeployContractJsonObjectV1Request) HasConstructorArgs() bool {
	if o != nil && !IsNil(o.ConstructorArgs) {
		return true
	}

	return false
}

// SetConstructorArgs gets a reference to the given []interface{} and assigns it to the ConstructorArgs field.
func (o *DeployContractJsonObjectV1Request) SetConstructorArgs(v []interface{}) {
	o.ConstructorArgs = v
}

func (o DeployContractJsonObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployContractJsonObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.GasPrice) {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if !IsNil(o.TimeoutMs) {
		toSerialize["timeoutMs"] = o.TimeoutMs
	}
	toSerialize["contractJSON"] = o.ContractJSON
	if !IsNil(o.ConstructorArgs) {
		toSerialize["constructorArgs"] = o.ConstructorArgs
	}
	return toSerialize, nil
}

type NullableDeployContractJsonObjectV1Request struct {
	value *DeployContractJsonObjectV1Request
	isSet bool
}

func (v NullableDeployContractJsonObjectV1Request) Get() *DeployContractJsonObjectV1Request {
	return v.value
}

func (v *NullableDeployContractJsonObjectV1Request) Set(val *DeployContractJsonObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployContractJsonObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployContractJsonObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployContractJsonObjectV1Request(val *DeployContractJsonObjectV1Request) *NullableDeployContractJsonObjectV1Request {
	return &NullableDeployContractJsonObjectV1Request{value: val, isSet: true}
}

func (v NullableDeployContractJsonObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployContractJsonObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


