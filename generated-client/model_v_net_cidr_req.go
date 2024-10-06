/*
Azure IPAM

 Azure IPAM is a lightweight solution developed on top of the Azure platform designed to help Azure customers manage their enterprise IP Address space easily and effectively. 

API version: 3.4.0
Contact: ipam@microsoft.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VNetCIDRReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VNetCIDRReq{}

// VNetCIDRReq DOCSTRING
type VNetCIDRReq struct {
	Space string `json:"space"`
	Blocks []string `json:"blocks"`
	Size int32 `json:"size"`
	ReverseSearch NullableBool `json:"reverse_search,omitempty"`
	SmallestCidr NullableBool `json:"smallest_cidr,omitempty"`
}

type _VNetCIDRReq VNetCIDRReq

// NewVNetCIDRReq instantiates a new VNetCIDRReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVNetCIDRReq(space string, blocks []string, size int32) *VNetCIDRReq {
	this := VNetCIDRReq{}
	this.Space = space
	this.Blocks = blocks
	this.Size = size
	return &this
}

// NewVNetCIDRReqWithDefaults instantiates a new VNetCIDRReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVNetCIDRReqWithDefaults() *VNetCIDRReq {
	this := VNetCIDRReq{}
	return &this
}

// GetSpace returns the Space field value
func (o *VNetCIDRReq) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *VNetCIDRReq) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *VNetCIDRReq) SetSpace(v string) {
	o.Space = v
}

// GetBlocks returns the Blocks field value
func (o *VNetCIDRReq) GetBlocks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value
// and a boolean to check if the value has been set.
func (o *VNetCIDRReq) GetBlocksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blocks, true
}

// SetBlocks sets field value
func (o *VNetCIDRReq) SetBlocks(v []string) {
	o.Blocks = v
}

// GetSize returns the Size field value
func (o *VNetCIDRReq) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *VNetCIDRReq) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *VNetCIDRReq) SetSize(v int32) {
	o.Size = v
}

// GetReverseSearch returns the ReverseSearch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VNetCIDRReq) GetReverseSearch() bool {
	if o == nil || IsNil(o.ReverseSearch.Get()) {
		var ret bool
		return ret
	}
	return *o.ReverseSearch.Get()
}

// GetReverseSearchOk returns a tuple with the ReverseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VNetCIDRReq) GetReverseSearchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReverseSearch.Get(), o.ReverseSearch.IsSet()
}

// HasReverseSearch returns a boolean if a field has been set.
func (o *VNetCIDRReq) HasReverseSearch() bool {
	if o != nil && o.ReverseSearch.IsSet() {
		return true
	}

	return false
}

// SetReverseSearch gets a reference to the given NullableBool and assigns it to the ReverseSearch field.
func (o *VNetCIDRReq) SetReverseSearch(v bool) {
	o.ReverseSearch.Set(&v)
}
// SetReverseSearchNil sets the value for ReverseSearch to be an explicit nil
func (o *VNetCIDRReq) SetReverseSearchNil() {
	o.ReverseSearch.Set(nil)
}

// UnsetReverseSearch ensures that no value is present for ReverseSearch, not even an explicit nil
func (o *VNetCIDRReq) UnsetReverseSearch() {
	o.ReverseSearch.Unset()
}

// GetSmallestCidr returns the SmallestCidr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VNetCIDRReq) GetSmallestCidr() bool {
	if o == nil || IsNil(o.SmallestCidr.Get()) {
		var ret bool
		return ret
	}
	return *o.SmallestCidr.Get()
}

// GetSmallestCidrOk returns a tuple with the SmallestCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VNetCIDRReq) GetSmallestCidrOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmallestCidr.Get(), o.SmallestCidr.IsSet()
}

// HasSmallestCidr returns a boolean if a field has been set.
func (o *VNetCIDRReq) HasSmallestCidr() bool {
	if o != nil && o.SmallestCidr.IsSet() {
		return true
	}

	return false
}

// SetSmallestCidr gets a reference to the given NullableBool and assigns it to the SmallestCidr field.
func (o *VNetCIDRReq) SetSmallestCidr(v bool) {
	o.SmallestCidr.Set(&v)
}
// SetSmallestCidrNil sets the value for SmallestCidr to be an explicit nil
func (o *VNetCIDRReq) SetSmallestCidrNil() {
	o.SmallestCidr.Set(nil)
}

// UnsetSmallestCidr ensures that no value is present for SmallestCidr, not even an explicit nil
func (o *VNetCIDRReq) UnsetSmallestCidr() {
	o.SmallestCidr.Unset()
}

func (o VNetCIDRReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VNetCIDRReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["space"] = o.Space
	toSerialize["blocks"] = o.Blocks
	toSerialize["size"] = o.Size
	if o.ReverseSearch.IsSet() {
		toSerialize["reverse_search"] = o.ReverseSearch.Get()
	}
	if o.SmallestCidr.IsSet() {
		toSerialize["smallest_cidr"] = o.SmallestCidr.Get()
	}
	return toSerialize, nil
}

func (o *VNetCIDRReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"space",
		"blocks",
		"size",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVNetCIDRReq := _VNetCIDRReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVNetCIDRReq)

	if err != nil {
		return err
	}

	*o = VNetCIDRReq(varVNetCIDRReq)

	return err
}

type NullableVNetCIDRReq struct {
	value *VNetCIDRReq
	isSet bool
}

func (v NullableVNetCIDRReq) Get() *VNetCIDRReq {
	return v.value
}

func (v *NullableVNetCIDRReq) Set(val *VNetCIDRReq) {
	v.value = val
	v.isSet = true
}

func (v NullableVNetCIDRReq) IsSet() bool {
	return v.isSet
}

func (v *NullableVNetCIDRReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVNetCIDRReq(val *VNetCIDRReq) *NullableVNetCIDRReq {
	return &NullableVNetCIDRReq{value: val, isSet: true}
}

func (v NullableVNetCIDRReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVNetCIDRReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


