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
)

// checks if the LogStatusSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogStatusSummary{}

// LogStatusSummary Summary of log status
type LogStatusSummary struct {
	// Number of recipients
	Recipients *int64 `json:"Recipients,omitempty"`
	// Number of emails
	EmailTotal *int64 `json:"EmailTotal,omitempty"`
	// Number of SMS
	SmsTotal *int64 `json:"SmsTotal,omitempty"`
	// Number of delivered messages
	Delivered *int64 `json:"Delivered,omitempty"`
	// Number of bounced messages
	Bounced *int64 `json:"Bounced,omitempty"`
	// Number of messages in progress
	InProgress *int64 `json:"InProgress,omitempty"`
	// Number of opened messages
	Opened *int64 `json:"Opened,omitempty"`
	// Number of clicked messages
	Clicked *int64 `json:"Clicked,omitempty"`
	// Number of unsubscribed messages
	Unsubscribed *int64 `json:"Unsubscribed,omitempty"`
	// Number of complaint messages
	Complaints *int64 `json:"Complaints,omitempty"`
	// Number of inbound messages
	Inbound *int64 `json:"Inbound,omitempty"`
	// Number of manually cancelled messages
	ManualCancel *int64 `json:"ManualCancel,omitempty"`
	// Number of messages flagged with 'Not Delivered'
	NotDelivered *int64 `json:"NotDelivered,omitempty"`
}

// NewLogStatusSummary instantiates a new LogStatusSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStatusSummary() *LogStatusSummary {
	this := LogStatusSummary{}
	return &this
}

// NewLogStatusSummaryWithDefaults instantiates a new LogStatusSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStatusSummaryWithDefaults() *LogStatusSummary {
	this := LogStatusSummary{}
	return &this
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *LogStatusSummary) GetRecipients() int64 {
	if o == nil || IsNil(o.Recipients) {
		var ret int64
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetRecipientsOk() (*int64, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *LogStatusSummary) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given int64 and assigns it to the Recipients field.
func (o *LogStatusSummary) SetRecipients(v int64) {
	o.Recipients = &v
}

// GetEmailTotal returns the EmailTotal field value if set, zero value otherwise.
func (o *LogStatusSummary) GetEmailTotal() int64 {
	if o == nil || IsNil(o.EmailTotal) {
		var ret int64
		return ret
	}
	return *o.EmailTotal
}

// GetEmailTotalOk returns a tuple with the EmailTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetEmailTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.EmailTotal) {
		return nil, false
	}
	return o.EmailTotal, true
}

// HasEmailTotal returns a boolean if a field has been set.
func (o *LogStatusSummary) HasEmailTotal() bool {
	if o != nil && !IsNil(o.EmailTotal) {
		return true
	}

	return false
}

// SetEmailTotal gets a reference to the given int64 and assigns it to the EmailTotal field.
func (o *LogStatusSummary) SetEmailTotal(v int64) {
	o.EmailTotal = &v
}

// GetSmsTotal returns the SmsTotal field value if set, zero value otherwise.
func (o *LogStatusSummary) GetSmsTotal() int64 {
	if o == nil || IsNil(o.SmsTotal) {
		var ret int64
		return ret
	}
	return *o.SmsTotal
}

// GetSmsTotalOk returns a tuple with the SmsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetSmsTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.SmsTotal) {
		return nil, false
	}
	return o.SmsTotal, true
}

// HasSmsTotal returns a boolean if a field has been set.
func (o *LogStatusSummary) HasSmsTotal() bool {
	if o != nil && !IsNil(o.SmsTotal) {
		return true
	}

	return false
}

// SetSmsTotal gets a reference to the given int64 and assigns it to the SmsTotal field.
func (o *LogStatusSummary) SetSmsTotal(v int64) {
	o.SmsTotal = &v
}

// GetDelivered returns the Delivered field value if set, zero value otherwise.
func (o *LogStatusSummary) GetDelivered() int64 {
	if o == nil || IsNil(o.Delivered) {
		var ret int64
		return ret
	}
	return *o.Delivered
}

// GetDeliveredOk returns a tuple with the Delivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetDeliveredOk() (*int64, bool) {
	if o == nil || IsNil(o.Delivered) {
		return nil, false
	}
	return o.Delivered, true
}

// HasDelivered returns a boolean if a field has been set.
func (o *LogStatusSummary) HasDelivered() bool {
	if o != nil && !IsNil(o.Delivered) {
		return true
	}

	return false
}

// SetDelivered gets a reference to the given int64 and assigns it to the Delivered field.
func (o *LogStatusSummary) SetDelivered(v int64) {
	o.Delivered = &v
}

// GetBounced returns the Bounced field value if set, zero value otherwise.
func (o *LogStatusSummary) GetBounced() int64 {
	if o == nil || IsNil(o.Bounced) {
		var ret int64
		return ret
	}
	return *o.Bounced
}

// GetBouncedOk returns a tuple with the Bounced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetBouncedOk() (*int64, bool) {
	if o == nil || IsNil(o.Bounced) {
		return nil, false
	}
	return o.Bounced, true
}

// HasBounced returns a boolean if a field has been set.
func (o *LogStatusSummary) HasBounced() bool {
	if o != nil && !IsNil(o.Bounced) {
		return true
	}

	return false
}

// SetBounced gets a reference to the given int64 and assigns it to the Bounced field.
func (o *LogStatusSummary) SetBounced(v int64) {
	o.Bounced = &v
}

// GetInProgress returns the InProgress field value if set, zero value otherwise.
func (o *LogStatusSummary) GetInProgress() int64 {
	if o == nil || IsNil(o.InProgress) {
		var ret int64
		return ret
	}
	return *o.InProgress
}

// GetInProgressOk returns a tuple with the InProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetInProgressOk() (*int64, bool) {
	if o == nil || IsNil(o.InProgress) {
		return nil, false
	}
	return o.InProgress, true
}

// HasInProgress returns a boolean if a field has been set.
func (o *LogStatusSummary) HasInProgress() bool {
	if o != nil && !IsNil(o.InProgress) {
		return true
	}

	return false
}

// SetInProgress gets a reference to the given int64 and assigns it to the InProgress field.
func (o *LogStatusSummary) SetInProgress(v int64) {
	o.InProgress = &v
}

// GetOpened returns the Opened field value if set, zero value otherwise.
func (o *LogStatusSummary) GetOpened() int64 {
	if o == nil || IsNil(o.Opened) {
		var ret int64
		return ret
	}
	return *o.Opened
}

// GetOpenedOk returns a tuple with the Opened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetOpenedOk() (*int64, bool) {
	if o == nil || IsNil(o.Opened) {
		return nil, false
	}
	return o.Opened, true
}

// HasOpened returns a boolean if a field has been set.
func (o *LogStatusSummary) HasOpened() bool {
	if o != nil && !IsNil(o.Opened) {
		return true
	}

	return false
}

// SetOpened gets a reference to the given int64 and assigns it to the Opened field.
func (o *LogStatusSummary) SetOpened(v int64) {
	o.Opened = &v
}

// GetClicked returns the Clicked field value if set, zero value otherwise.
func (o *LogStatusSummary) GetClicked() int64 {
	if o == nil || IsNil(o.Clicked) {
		var ret int64
		return ret
	}
	return *o.Clicked
}

// GetClickedOk returns a tuple with the Clicked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetClickedOk() (*int64, bool) {
	if o == nil || IsNil(o.Clicked) {
		return nil, false
	}
	return o.Clicked, true
}

// HasClicked returns a boolean if a field has been set.
func (o *LogStatusSummary) HasClicked() bool {
	if o != nil && !IsNil(o.Clicked) {
		return true
	}

	return false
}

// SetClicked gets a reference to the given int64 and assigns it to the Clicked field.
func (o *LogStatusSummary) SetClicked(v int64) {
	o.Clicked = &v
}

// GetUnsubscribed returns the Unsubscribed field value if set, zero value otherwise.
func (o *LogStatusSummary) GetUnsubscribed() int64 {
	if o == nil || IsNil(o.Unsubscribed) {
		var ret int64
		return ret
	}
	return *o.Unsubscribed
}

// GetUnsubscribedOk returns a tuple with the Unsubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetUnsubscribedOk() (*int64, bool) {
	if o == nil || IsNil(o.Unsubscribed) {
		return nil, false
	}
	return o.Unsubscribed, true
}

// HasUnsubscribed returns a boolean if a field has been set.
func (o *LogStatusSummary) HasUnsubscribed() bool {
	if o != nil && !IsNil(o.Unsubscribed) {
		return true
	}

	return false
}

// SetUnsubscribed gets a reference to the given int64 and assigns it to the Unsubscribed field.
func (o *LogStatusSummary) SetUnsubscribed(v int64) {
	o.Unsubscribed = &v
}

// GetComplaints returns the Complaints field value if set, zero value otherwise.
func (o *LogStatusSummary) GetComplaints() int64 {
	if o == nil || IsNil(o.Complaints) {
		var ret int64
		return ret
	}
	return *o.Complaints
}

// GetComplaintsOk returns a tuple with the Complaints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetComplaintsOk() (*int64, bool) {
	if o == nil || IsNil(o.Complaints) {
		return nil, false
	}
	return o.Complaints, true
}

// HasComplaints returns a boolean if a field has been set.
func (o *LogStatusSummary) HasComplaints() bool {
	if o != nil && !IsNil(o.Complaints) {
		return true
	}

	return false
}

// SetComplaints gets a reference to the given int64 and assigns it to the Complaints field.
func (o *LogStatusSummary) SetComplaints(v int64) {
	o.Complaints = &v
}

// GetInbound returns the Inbound field value if set, zero value otherwise.
func (o *LogStatusSummary) GetInbound() int64 {
	if o == nil || IsNil(o.Inbound) {
		var ret int64
		return ret
	}
	return *o.Inbound
}

// GetInboundOk returns a tuple with the Inbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetInboundOk() (*int64, bool) {
	if o == nil || IsNil(o.Inbound) {
		return nil, false
	}
	return o.Inbound, true
}

// HasInbound returns a boolean if a field has been set.
func (o *LogStatusSummary) HasInbound() bool {
	if o != nil && !IsNil(o.Inbound) {
		return true
	}

	return false
}

// SetInbound gets a reference to the given int64 and assigns it to the Inbound field.
func (o *LogStatusSummary) SetInbound(v int64) {
	o.Inbound = &v
}

// GetManualCancel returns the ManualCancel field value if set, zero value otherwise.
func (o *LogStatusSummary) GetManualCancel() int64 {
	if o == nil || IsNil(o.ManualCancel) {
		var ret int64
		return ret
	}
	return *o.ManualCancel
}

// GetManualCancelOk returns a tuple with the ManualCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetManualCancelOk() (*int64, bool) {
	if o == nil || IsNil(o.ManualCancel) {
		return nil, false
	}
	return o.ManualCancel, true
}

// HasManualCancel returns a boolean if a field has been set.
func (o *LogStatusSummary) HasManualCancel() bool {
	if o != nil && !IsNil(o.ManualCancel) {
		return true
	}

	return false
}

// SetManualCancel gets a reference to the given int64 and assigns it to the ManualCancel field.
func (o *LogStatusSummary) SetManualCancel(v int64) {
	o.ManualCancel = &v
}

// GetNotDelivered returns the NotDelivered field value if set, zero value otherwise.
func (o *LogStatusSummary) GetNotDelivered() int64 {
	if o == nil || IsNil(o.NotDelivered) {
		var ret int64
		return ret
	}
	return *o.NotDelivered
}

// GetNotDeliveredOk returns a tuple with the NotDelivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStatusSummary) GetNotDeliveredOk() (*int64, bool) {
	if o == nil || IsNil(o.NotDelivered) {
		return nil, false
	}
	return o.NotDelivered, true
}

// HasNotDelivered returns a boolean if a field has been set.
func (o *LogStatusSummary) HasNotDelivered() bool {
	if o != nil && !IsNil(o.NotDelivered) {
		return true
	}

	return false
}

// SetNotDelivered gets a reference to the given int64 and assigns it to the NotDelivered field.
func (o *LogStatusSummary) SetNotDelivered(v int64) {
	o.NotDelivered = &v
}

func (o LogStatusSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogStatusSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recipients) {
		toSerialize["Recipients"] = o.Recipients
	}
	if !IsNil(o.EmailTotal) {
		toSerialize["EmailTotal"] = o.EmailTotal
	}
	if !IsNil(o.SmsTotal) {
		toSerialize["SmsTotal"] = o.SmsTotal
	}
	if !IsNil(o.Delivered) {
		toSerialize["Delivered"] = o.Delivered
	}
	if !IsNil(o.Bounced) {
		toSerialize["Bounced"] = o.Bounced
	}
	if !IsNil(o.InProgress) {
		toSerialize["InProgress"] = o.InProgress
	}
	if !IsNil(o.Opened) {
		toSerialize["Opened"] = o.Opened
	}
	if !IsNil(o.Clicked) {
		toSerialize["Clicked"] = o.Clicked
	}
	if !IsNil(o.Unsubscribed) {
		toSerialize["Unsubscribed"] = o.Unsubscribed
	}
	if !IsNil(o.Complaints) {
		toSerialize["Complaints"] = o.Complaints
	}
	if !IsNil(o.Inbound) {
		toSerialize["Inbound"] = o.Inbound
	}
	if !IsNil(o.ManualCancel) {
		toSerialize["ManualCancel"] = o.ManualCancel
	}
	if !IsNil(o.NotDelivered) {
		toSerialize["NotDelivered"] = o.NotDelivered
	}
	return toSerialize, nil
}

type NullableLogStatusSummary struct {
	value *LogStatusSummary
	isSet bool
}

func (v NullableLogStatusSummary) Get() *LogStatusSummary {
	return v.value
}

func (v *NullableLogStatusSummary) Set(val *LogStatusSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStatusSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStatusSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStatusSummary(val *LogStatusSummary) *NullableLogStatusSummary {
	return &NullableLogStatusSummary{value: val, isSet: true}
}

func (v NullableLogStatusSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStatusSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


