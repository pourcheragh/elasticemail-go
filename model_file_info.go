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
	"time"
)

// checks if the FileInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileInfo{}

// FileInfo File information
type FileInfo struct {
	// Name of your file including extension.
	FileName *string `json:"FileName,omitempty"`
	// Size of your attachment (in bytes).
	Size NullableInt32 `json:"Size,omitempty"`
	// Date of creation in YYYY-MM-DDThh:ii:ss format
	DateAdded *time.Time `json:"DateAdded,omitempty"`
	// Date when the file will be deleted from your Account.
	ExpirationDate NullableTime `json:"ExpirationDate,omitempty"`
	// Content type of the file.
	ContentType *string `json:"ContentType,omitempty"`
}

// NewFileInfo instantiates a new FileInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileInfo() *FileInfo {
	this := FileInfo{}
	return &this
}

// NewFileInfoWithDefaults instantiates a new FileInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileInfoWithDefaults() *FileInfo {
	this := FileInfo{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *FileInfo) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInfo) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *FileInfo) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *FileInfo) SetFileName(v string) {
	o.FileName = &v
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileInfo) GetSize() int32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInfo) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *FileInfo) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *FileInfo) SetSize(v int32) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *FileInfo) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *FileInfo) UnsetSize() {
	o.Size.Unset()
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *FileInfo) GetDateAdded() time.Time {
	if o == nil || IsNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInfo) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *FileInfo) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *FileInfo) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileInfo) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInfo) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FileInfo) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableTime and assigns it to the ExpirationDate field.
func (o *FileInfo) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *FileInfo) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *FileInfo) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *FileInfo) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInfo) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *FileInfo) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *FileInfo) SetContentType(v string) {
	o.ContentType = &v
}

func (o FileInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileName) {
		toSerialize["FileName"] = o.FileName
	}
	if o.Size.IsSet() {
		toSerialize["Size"] = o.Size.Get()
	}
	if !IsNil(o.DateAdded) {
		toSerialize["DateAdded"] = o.DateAdded
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["ExpirationDate"] = o.ExpirationDate.Get()
	}
	if !IsNil(o.ContentType) {
		toSerialize["ContentType"] = o.ContentType
	}
	return toSerialize, nil
}

type NullableFileInfo struct {
	value *FileInfo
	isSet bool
}

func (v NullableFileInfo) Get() *FileInfo {
	return v.value
}

func (v *NullableFileInfo) Set(val *FileInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFileInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFileInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileInfo(val *FileInfo) *NullableFileInfo {
	return &NullableFileInfo{value: val, isSet: true}
}

func (v NullableFileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


