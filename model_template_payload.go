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
	"bytes"
	"fmt"
)

// checks if the TemplatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplatePayload{}

// TemplatePayload New template object
type TemplatePayload struct {
	// Template name
	Name string `json:"Name"`
	// Default subject of email.
	Subject *string `json:"Subject,omitempty"`
	// Email content of this template
	Body []BodyPart `json:"Body,omitempty"`
	TemplateScope *TemplateScope `json:"TemplateScope,omitempty"`
}

type _TemplatePayload TemplatePayload

// NewTemplatePayload instantiates a new TemplatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatePayload(name string) *TemplatePayload {
	this := TemplatePayload{}
	this.Name = name
	var templateScope TemplateScope = TEMPLATESCOPE_PERSONAL
	this.TemplateScope = &templateScope
	return &this
}

// NewTemplatePayloadWithDefaults instantiates a new TemplatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePayloadWithDefaults() *TemplatePayload {
	this := TemplatePayload{}
	var templateScope TemplateScope = TEMPLATESCOPE_PERSONAL
	this.TemplateScope = &templateScope
	return &this
}

// GetName returns the Name field value
func (o *TemplatePayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplatePayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplatePayload) SetName(v string) {
	o.Name = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TemplatePayload) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePayload) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TemplatePayload) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TemplatePayload) SetSubject(v string) {
	o.Subject = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *TemplatePayload) GetBody() []BodyPart {
	if o == nil || IsNil(o.Body) {
		var ret []BodyPart
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePayload) GetBodyOk() ([]BodyPart, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *TemplatePayload) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given []BodyPart and assigns it to the Body field.
func (o *TemplatePayload) SetBody(v []BodyPart) {
	o.Body = v
}

// GetTemplateScope returns the TemplateScope field value if set, zero value otherwise.
func (o *TemplatePayload) GetTemplateScope() TemplateScope {
	if o == nil || IsNil(o.TemplateScope) {
		var ret TemplateScope
		return ret
	}
	return *o.TemplateScope
}

// GetTemplateScopeOk returns a tuple with the TemplateScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePayload) GetTemplateScopeOk() (*TemplateScope, bool) {
	if o == nil || IsNil(o.TemplateScope) {
		return nil, false
	}
	return o.TemplateScope, true
}

// HasTemplateScope returns a boolean if a field has been set.
func (o *TemplatePayload) HasTemplateScope() bool {
	if o != nil && !IsNil(o.TemplateScope) {
		return true
	}

	return false
}

// SetTemplateScope gets a reference to the given TemplateScope and assigns it to the TemplateScope field.
func (o *TemplatePayload) SetTemplateScope(v TemplateScope) {
	o.TemplateScope = &v
}

func (o TemplatePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	if !IsNil(o.Subject) {
		toSerialize["Subject"] = o.Subject
	}
	if !IsNil(o.Body) {
		toSerialize["Body"] = o.Body
	}
	if !IsNil(o.TemplateScope) {
		toSerialize["TemplateScope"] = o.TemplateScope
	}
	return toSerialize, nil
}

func (o *TemplatePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Name",
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

	varTemplatePayload := _TemplatePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplatePayload)

	if err != nil {
		return err
	}

	*o = TemplatePayload(varTemplatePayload)

	return err
}

type NullableTemplatePayload struct {
	value *TemplatePayload
	isSet bool
}

func (v NullableTemplatePayload) Get() *TemplatePayload {
	return v.value
}

func (v *NullableTemplatePayload) Set(val *TemplatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatePayload(val *TemplatePayload) *NullableTemplatePayload {
	return &NullableTemplatePayload{value: val, isSet: true}
}

func (v NullableTemplatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


