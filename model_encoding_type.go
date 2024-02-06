/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package elasticemail

import (
	"encoding/json"
	"fmt"
)

// EncodingType Encoding type for the email headers
type EncodingType string

// List of EncodingType
const (
	ENCODINGTYPE_USER_PROVIDED EncodingType = "UserProvided"
	ENCODINGTYPE_NONE EncodingType = "None"
	ENCODINGTYPE_RAW7BIT EncodingType = "Raw7bit"
	ENCODINGTYPE_RAW8BIT EncodingType = "Raw8bit"
	ENCODINGTYPE_QUOTED_PRINTABLE EncodingType = "QuotedPrintable"
	ENCODINGTYPE_BASE64 EncodingType = "Base64"
	ENCODINGTYPE_UUE EncodingType = "Uue"
)

// All allowed values of EncodingType enum
var AllowedEncodingTypeEnumValues = []EncodingType{
	"UserProvided",
	"None",
	"Raw7bit",
	"Raw8bit",
	"QuotedPrintable",
	"Base64",
	"Uue",
}

func (v *EncodingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EncodingType(value)
	for _, existing := range AllowedEncodingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EncodingType", value)
}

// NewEncodingTypeFromValue returns a pointer to a valid EncodingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEncodingTypeFromValue(v string) (*EncodingType, error) {
	ev := EncodingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EncodingType: valid values are %v", v, AllowedEncodingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EncodingType) IsValid() bool {
	for _, existing := range AllowedEncodingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EncodingType value
func (v EncodingType) Ptr() *EncodingType {
	return &v
}

type NullableEncodingType struct {
	value *EncodingType
	isSet bool
}

func (v NullableEncodingType) Get() *EncodingType {
	return v.value
}

func (v *NullableEncodingType) Set(val *EncodingType) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodingType) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodingType(val *EncodingType) *NullableEncodingType {
	return &NullableEncodingType{value: val, isSet: true}
}

func (v NullableEncodingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
