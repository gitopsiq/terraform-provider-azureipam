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


// ResponseGetSpacesApiSpacesGet struct for ResponseGetSpacesApiSpacesGet
type ResponseGetSpacesApiSpacesGet struct {
	ArrayOfSpace *[]Space
	ArrayOfSpaceBasic *[]SpaceBasic
	ArrayOfSpaceBasicUtil *[]SpaceBasicUtil
	ArrayOfSpaceExpand *[]SpaceExpand
	ArrayOfSpaceExpandUtil *[]SpaceExpandUtil
	ArrayOfSpaceUtil *[]SpaceUtil
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResponseGetSpacesApiSpacesGet) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfSpace
	err = json.Unmarshal(data, &dst.ArrayOfSpace);
	if err == nil {
		jsonArrayOfSpace, _ := json.Marshal(dst.ArrayOfSpace)
		if string(jsonArrayOfSpace) == "{}" { // empty struct
			dst.ArrayOfSpace = nil
		} else {
			return nil // data stored in dst.ArrayOfSpace, return on the first match
		}
	} else {
		dst.ArrayOfSpace = nil
	}

	// try to unmarshal JSON data into ArrayOfSpaceBasic
	err = json.Unmarshal(data, &dst.ArrayOfSpaceBasic);
	if err == nil {
		jsonArrayOfSpaceBasic, _ := json.Marshal(dst.ArrayOfSpaceBasic)
		if string(jsonArrayOfSpaceBasic) == "{}" { // empty struct
			dst.ArrayOfSpaceBasic = nil
		} else {
			return nil // data stored in dst.ArrayOfSpaceBasic, return on the first match
		}
	} else {
		dst.ArrayOfSpaceBasic = nil
	}

	// try to unmarshal JSON data into ArrayOfSpaceBasicUtil
	err = json.Unmarshal(data, &dst.ArrayOfSpaceBasicUtil);
	if err == nil {
		jsonArrayOfSpaceBasicUtil, _ := json.Marshal(dst.ArrayOfSpaceBasicUtil)
		if string(jsonArrayOfSpaceBasicUtil) == "{}" { // empty struct
			dst.ArrayOfSpaceBasicUtil = nil
		} else {
			return nil // data stored in dst.ArrayOfSpaceBasicUtil, return on the first match
		}
	} else {
		dst.ArrayOfSpaceBasicUtil = nil
	}

	// try to unmarshal JSON data into ArrayOfSpaceExpand
	err = json.Unmarshal(data, &dst.ArrayOfSpaceExpand);
	if err == nil {
		jsonArrayOfSpaceExpand, _ := json.Marshal(dst.ArrayOfSpaceExpand)
		if string(jsonArrayOfSpaceExpand) == "{}" { // empty struct
			dst.ArrayOfSpaceExpand = nil
		} else {
			return nil // data stored in dst.ArrayOfSpaceExpand, return on the first match
		}
	} else {
		dst.ArrayOfSpaceExpand = nil
	}

	// try to unmarshal JSON data into ArrayOfSpaceExpandUtil
	err = json.Unmarshal(data, &dst.ArrayOfSpaceExpandUtil);
	if err == nil {
		jsonArrayOfSpaceExpandUtil, _ := json.Marshal(dst.ArrayOfSpaceExpandUtil)
		if string(jsonArrayOfSpaceExpandUtil) == "{}" { // empty struct
			dst.ArrayOfSpaceExpandUtil = nil
		} else {
			return nil // data stored in dst.ArrayOfSpaceExpandUtil, return on the first match
		}
	} else {
		dst.ArrayOfSpaceExpandUtil = nil
	}

	// try to unmarshal JSON data into ArrayOfSpaceUtil
	err = json.Unmarshal(data, &dst.ArrayOfSpaceUtil);
	if err == nil {
		jsonArrayOfSpaceUtil, _ := json.Marshal(dst.ArrayOfSpaceUtil)
		if string(jsonArrayOfSpaceUtil) == "{}" { // empty struct
			dst.ArrayOfSpaceUtil = nil
		} else {
			return nil // data stored in dst.ArrayOfSpaceUtil, return on the first match
		}
	} else {
		dst.ArrayOfSpaceUtil = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResponseGetSpacesApiSpacesGet)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResponseGetSpacesApiSpacesGet) MarshalJSON() ([]byte, error) {
	if src.ArrayOfSpace != nil {
		return json.Marshal(&src.ArrayOfSpace)
	}

	if src.ArrayOfSpaceBasic != nil {
		return json.Marshal(&src.ArrayOfSpaceBasic)
	}

	if src.ArrayOfSpaceBasicUtil != nil {
		return json.Marshal(&src.ArrayOfSpaceBasicUtil)
	}

	if src.ArrayOfSpaceExpand != nil {
		return json.Marshal(&src.ArrayOfSpaceExpand)
	}

	if src.ArrayOfSpaceExpandUtil != nil {
		return json.Marshal(&src.ArrayOfSpaceExpandUtil)
	}

	if src.ArrayOfSpaceUtil != nil {
		return json.Marshal(&src.ArrayOfSpaceUtil)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableResponseGetSpacesApiSpacesGet struct {
	value *ResponseGetSpacesApiSpacesGet
	isSet bool
}

func (v NullableResponseGetSpacesApiSpacesGet) Get() *ResponseGetSpacesApiSpacesGet {
	return v.value
}

func (v *NullableResponseGetSpacesApiSpacesGet) Set(val *ResponseGetSpacesApiSpacesGet) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseGetSpacesApiSpacesGet) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseGetSpacesApiSpacesGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseGetSpacesApiSpacesGet(val *ResponseGetSpacesApiSpacesGet) *NullableResponseGetSpacesApiSpacesGet {
	return &NullableResponseGetSpacesApiSpacesGet{value: val, isSet: true}
}

func (v NullableResponseGetSpacesApiSpacesGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseGetSpacesApiSpacesGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


