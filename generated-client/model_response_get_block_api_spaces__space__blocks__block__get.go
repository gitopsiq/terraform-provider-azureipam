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
	"fmt"
)


// ResponseGetBlockApiSpacesSpaceBlocksBlockGet struct for ResponseGetBlockApiSpacesSpaceBlocksBlockGet
type ResponseGetBlockApiSpacesSpaceBlocksBlockGet struct {
	Block *Block
	BlockBasic *BlockBasic
	BlockBasicUtil *BlockBasicUtil
	BlockExpand *BlockExpand
	BlockExpandUtil *BlockExpandUtil
	BlockUtil *BlockUtil
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResponseGetBlockApiSpacesSpaceBlocksBlockGet) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Block
	err = json.Unmarshal(data, &dst.Block);
	if err == nil {
		jsonBlock, _ := json.Marshal(dst.Block)
		if string(jsonBlock) == "{}" { // empty struct
			dst.Block = nil
		} else {
			return nil // data stored in dst.Block, return on the first match
		}
	} else {
		dst.Block = nil
	}

	// try to unmarshal JSON data into BlockBasic
	err = json.Unmarshal(data, &dst.BlockBasic);
	if err == nil {
		jsonBlockBasic, _ := json.Marshal(dst.BlockBasic)
		if string(jsonBlockBasic) == "{}" { // empty struct
			dst.BlockBasic = nil
		} else {
			return nil // data stored in dst.BlockBasic, return on the first match
		}
	} else {
		dst.BlockBasic = nil
	}

	// try to unmarshal JSON data into BlockBasicUtil
	err = json.Unmarshal(data, &dst.BlockBasicUtil);
	if err == nil {
		jsonBlockBasicUtil, _ := json.Marshal(dst.BlockBasicUtil)
		if string(jsonBlockBasicUtil) == "{}" { // empty struct
			dst.BlockBasicUtil = nil
		} else {
			return nil // data stored in dst.BlockBasicUtil, return on the first match
		}
	} else {
		dst.BlockBasicUtil = nil
	}

	// try to unmarshal JSON data into BlockExpand
	err = json.Unmarshal(data, &dst.BlockExpand);
	if err == nil {
		jsonBlockExpand, _ := json.Marshal(dst.BlockExpand)
		if string(jsonBlockExpand) == "{}" { // empty struct
			dst.BlockExpand = nil
		} else {
			return nil // data stored in dst.BlockExpand, return on the first match
		}
	} else {
		dst.BlockExpand = nil
	}

	// try to unmarshal JSON data into BlockExpandUtil
	err = json.Unmarshal(data, &dst.BlockExpandUtil);
	if err == nil {
		jsonBlockExpandUtil, _ := json.Marshal(dst.BlockExpandUtil)
		if string(jsonBlockExpandUtil) == "{}" { // empty struct
			dst.BlockExpandUtil = nil
		} else {
			return nil // data stored in dst.BlockExpandUtil, return on the first match
		}
	} else {
		dst.BlockExpandUtil = nil
	}

	// try to unmarshal JSON data into BlockUtil
	err = json.Unmarshal(data, &dst.BlockUtil);
	if err == nil {
		jsonBlockUtil, _ := json.Marshal(dst.BlockUtil)
		if string(jsonBlockUtil) == "{}" { // empty struct
			dst.BlockUtil = nil
		} else {
			return nil // data stored in dst.BlockUtil, return on the first match
		}
	} else {
		dst.BlockUtil = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResponseGetBlockApiSpacesSpaceBlocksBlockGet)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResponseGetBlockApiSpacesSpaceBlocksBlockGet) MarshalJSON() ([]byte, error) {
	if src.Block != nil {
		return json.Marshal(&src.Block)
	}

	if src.BlockBasic != nil {
		return json.Marshal(&src.BlockBasic)
	}

	if src.BlockBasicUtil != nil {
		return json.Marshal(&src.BlockBasicUtil)
	}

	if src.BlockExpand != nil {
		return json.Marshal(&src.BlockExpand)
	}

	if src.BlockExpandUtil != nil {
		return json.Marshal(&src.BlockExpandUtil)
	}

	if src.BlockUtil != nil {
		return json.Marshal(&src.BlockUtil)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet struct {
	value *ResponseGetBlockApiSpacesSpaceBlocksBlockGet
	isSet bool
}

func (v NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet) Get() *ResponseGetBlockApiSpacesSpaceBlocksBlockGet {
	return v.value
}

func (v *NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet) Set(val *ResponseGetBlockApiSpacesSpaceBlocksBlockGet) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseGetBlockApiSpacesSpaceBlocksBlockGet(val *ResponseGetBlockApiSpacesSpaceBlocksBlockGet) *NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet {
	return &NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet{value: val, isSet: true}
}

func (v NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseGetBlockApiSpacesSpaceBlocksBlockGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


