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

// AccessLevel the model 'AccessLevel'
type AccessLevel string

// List of AccessLevel
const (
	ACCESSLEVEL_NONE AccessLevel = "None"
	ACCESSLEVEL_VIEW_ACCOUNT AccessLevel = "ViewAccount"
	ACCESSLEVEL_VIEW_CONTACTS AccessLevel = "ViewContacts"
	ACCESSLEVEL_VIEW_FORMS AccessLevel = "ViewForms"
	ACCESSLEVEL_VIEW_TEMPLATES AccessLevel = "ViewTemplates"
	ACCESSLEVEL_VIEW_CAMPAIGNS AccessLevel = "ViewCampaigns"
	ACCESSLEVEL_VIEW_CHANNELS AccessLevel = "ViewChannels"
	ACCESSLEVEL_VIEW_AUTOMATIONS AccessLevel = "ViewAutomations"
	ACCESSLEVEL_VIEW_SURVEYS AccessLevel = "ViewSurveys"
	ACCESSLEVEL_VIEW_SETTINGS AccessLevel = "ViewSettings"
	ACCESSLEVEL_VIEW_BILLING AccessLevel = "ViewBilling"
	ACCESSLEVEL_VIEW_SUB_ACCOUNTS AccessLevel = "ViewSubAccounts"
	ACCESSLEVEL_VIEW_USERS AccessLevel = "ViewUsers"
	ACCESSLEVEL_VIEW_FILES AccessLevel = "ViewFiles"
	ACCESSLEVEL_VIEW_REPORTS AccessLevel = "ViewReports"
	ACCESSLEVEL_MODIFY_ACCOUNT AccessLevel = "ModifyAccount"
	ACCESSLEVEL_MODIFY_CONTACTS AccessLevel = "ModifyContacts"
	ACCESSLEVEL_MODIFY_FORMS AccessLevel = "ModifyForms"
	ACCESSLEVEL_MODIFY_TEMPLATES AccessLevel = "ModifyTemplates"
	ACCESSLEVEL_MODIFY_CAMPAIGNS AccessLevel = "ModifyCampaigns"
	ACCESSLEVEL_MODIFY_CHANNELS AccessLevel = "ModifyChannels"
	ACCESSLEVEL_MODIFY_AUTOMATIONS AccessLevel = "ModifyAutomations"
	ACCESSLEVEL_MODIFY_SURVEYS AccessLevel = "ModifySurveys"
	ACCESSLEVEL_MODIFY_FILES AccessLevel = "ModifyFiles"
	ACCESSLEVEL_EXPORT AccessLevel = "Export"
	ACCESSLEVEL_SEND_SMTP AccessLevel = "SendSmtp"
	ACCESSLEVEL_SEND_SMS AccessLevel = "SendSMS"
	ACCESSLEVEL_MODIFY_SETTINGS AccessLevel = "ModifySettings"
	ACCESSLEVEL_MODIFY_BILLING AccessLevel = "ModifyBilling"
	ACCESSLEVEL_MODIFY_PROFILE AccessLevel = "ModifyProfile"
	ACCESSLEVEL_MODIFY_SUB_ACCOUNTS AccessLevel = "ModifySubAccounts"
	ACCESSLEVEL_MODIFY_USERS AccessLevel = "ModifyUsers"
	ACCESSLEVEL_SECURITY AccessLevel = "Security"
	ACCESSLEVEL_MODIFY_LANGUAGE AccessLevel = "ModifyLanguage"
	ACCESSLEVEL_VIEW_SUPPORT AccessLevel = "ViewSupport"
	ACCESSLEVEL_SEND_HTTP AccessLevel = "SendHttp"
	ACCESSLEVEL_MODIFY2_FA_EMAIL AccessLevel = "Modify2FAEmail"
	ACCESSLEVEL_MODIFY_SUPPORT AccessLevel = "ModifySupport"
	ACCESSLEVEL_VIEW_CUSTOM_FIELDS AccessLevel = "ViewCustomFields"
	ACCESSLEVEL_MODIFY_CUSTOM_FIELDS AccessLevel = "ModifyCustomFields"
	ACCESSLEVEL_MODIFY_WEB_NOTIFICATIONS AccessLevel = "ModifyWebNotifications"
	ACCESSLEVEL_EXTENDED_LOGS AccessLevel = "ExtendedLogs"
	ACCESSLEVEL_VERIFY_EMAILS AccessLevel = "VerifyEmails"
	ACCESSLEVEL_MODIFY2_FA_SMS AccessLevel = "Modify2FASms"
	ACCESSLEVEL_DISABLE_CONTACTS_STORE AccessLevel = "DisableContactsStore"
	ACCESSLEVEL_MODIFY_LANDING_PAGES AccessLevel = "ModifyLandingPages"
	ACCESSLEVEL_VIEW_LANDING_PAGES AccessLevel = "ViewLandingPages"
	ACCESSLEVEL_MODIFY_SUPPRESSIONS AccessLevel = "ModifySuppressions"
	ACCESSLEVEL_VIEW_SUPPRESSIONS AccessLevel = "ViewSuppressions"
	ACCESSLEVEL_VIEW_DRAG_DROP_EDITOR AccessLevel = "ViewDragDropEditor"
	ACCESSLEVEL_VIEW_TEMPLATE_EDITOR AccessLevel = "ViewTemplateEditor"
)

// All allowed values of AccessLevel enum
var AllowedAccessLevelEnumValues = []AccessLevel{
	"None",
	"ViewAccount",
	"ViewContacts",
	"ViewForms",
	"ViewTemplates",
	"ViewCampaigns",
	"ViewChannels",
	"ViewAutomations",
	"ViewSurveys",
	"ViewSettings",
	"ViewBilling",
	"ViewSubAccounts",
	"ViewUsers",
	"ViewFiles",
	"ViewReports",
	"ModifyAccount",
	"ModifyContacts",
	"ModifyForms",
	"ModifyTemplates",
	"ModifyCampaigns",
	"ModifyChannels",
	"ModifyAutomations",
	"ModifySurveys",
	"ModifyFiles",
	"Export",
	"SendSmtp",
	"SendSMS",
	"ModifySettings",
	"ModifyBilling",
	"ModifyProfile",
	"ModifySubAccounts",
	"ModifyUsers",
	"Security",
	"ModifyLanguage",
	"ViewSupport",
	"SendHttp",
	"Modify2FAEmail",
	"ModifySupport",
	"ViewCustomFields",
	"ModifyCustomFields",
	"ModifyWebNotifications",
	"ExtendedLogs",
	"VerifyEmails",
	"Modify2FASms",
	"DisableContactsStore",
	"ModifyLandingPages",
	"ViewLandingPages",
	"ModifySuppressions",
	"ViewSuppressions",
	"ViewDragDropEditor",
	"ViewTemplateEditor",
}

func (v *AccessLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessLevel(value)
	for _, existing := range AllowedAccessLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessLevel", value)
}

// NewAccessLevelFromValue returns a pointer to a valid AccessLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessLevelFromValue(v string) (*AccessLevel, error) {
	ev := AccessLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessLevel: valid values are %v", v, AllowedAccessLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessLevel) IsValid() bool {
	for _, existing := range AllowedAccessLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessLevel value
func (v AccessLevel) Ptr() *AccessLevel {
	return &v
}

type NullableAccessLevel struct {
	value *AccessLevel
	isSet bool
}

func (v NullableAccessLevel) Get() *AccessLevel {
	return v.value
}

func (v *NullableAccessLevel) Set(val *AccessLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessLevel(val *AccessLevel) *NullableAccessLevel {
	return &NullableAccessLevel{value: val, isSet: true}
}

func (v NullableAccessLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

