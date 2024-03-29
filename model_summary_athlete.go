/*
Strava API v3

The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package strava

import (
	"encoding/json"
	"time"
)

// SummaryAthlete struct for SummaryAthlete
type SummaryAthlete struct {
	// The unique identifier of the athlete
	Id *int64 `json:"id,omitempty"`
	// Resource state, indicates level of detail. Possible values: 1 -> \"meta\", 2 -> \"summary\", 3 -> \"detail\"
	ResourceState *int32 `json:"resource_state,omitempty"`
	// The athlete's first name.
	Firstname *string `json:"firstname,omitempty"`
	// The athlete's last name.
	Lastname *string `json:"lastname,omitempty"`
	// URL to a 62x62 pixel profile picture.
	ProfileMedium *string `json:"profile_medium,omitempty"`
	// URL to a 124x124 pixel profile picture.
	Profile *string `json:"profile,omitempty"`
	// The athlete's city.
	City *string `json:"city,omitempty"`
	// The athlete's state or geographical region.
	State *string `json:"state,omitempty"`
	// The athlete's country.
	Country *string `json:"country,omitempty"`
	// The athlete's sex.
	Sex *string `json:"sex,omitempty"`
	// Deprecated.  Use summit field instead. Whether the athlete has any Summit subscription.
	Premium *bool `json:"premium,omitempty"`
	// Whether the athlete has any Summit subscription.
	Summit *bool `json:"summit,omitempty"`
	// The time at which the athlete was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time at which the athlete was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewSummaryAthlete instantiates a new SummaryAthlete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryAthlete() *SummaryAthlete {
	this := SummaryAthlete{}
	return &this
}

// NewSummaryAthleteWithDefaults instantiates a new SummaryAthlete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryAthleteWithDefaults() *SummaryAthlete {
	this := SummaryAthlete{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SummaryAthlete) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SummaryAthlete) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SummaryAthlete) SetId(v int64) {
	o.Id = &v
}

// GetResourceState returns the ResourceState field value if set, zero value otherwise.
func (o *SummaryAthlete) GetResourceState() int32 {
	if o == nil || o.ResourceState == nil {
		var ret int32
		return ret
	}
	return *o.ResourceState
}

// GetResourceStateOk returns a tuple with the ResourceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetResourceStateOk() (*int32, bool) {
	if o == nil || o.ResourceState == nil {
		return nil, false
	}
	return o.ResourceState, true
}

// HasResourceState returns a boolean if a field has been set.
func (o *SummaryAthlete) HasResourceState() bool {
	if o != nil && o.ResourceState != nil {
		return true
	}

	return false
}

// SetResourceState gets a reference to the given int32 and assigns it to the ResourceState field.
func (o *SummaryAthlete) SetResourceState(v int32) {
	o.ResourceState = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *SummaryAthlete) GetFirstname() string {
	if o == nil || o.Firstname == nil {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetFirstnameOk() (*string, bool) {
	if o == nil || o.Firstname == nil {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *SummaryAthlete) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *SummaryAthlete) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *SummaryAthlete) GetLastname() string {
	if o == nil || o.Lastname == nil {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetLastnameOk() (*string, bool) {
	if o == nil || o.Lastname == nil {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *SummaryAthlete) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *SummaryAthlete) SetLastname(v string) {
	o.Lastname = &v
}

// GetProfileMedium returns the ProfileMedium field value if set, zero value otherwise.
func (o *SummaryAthlete) GetProfileMedium() string {
	if o == nil || o.ProfileMedium == nil {
		var ret string
		return ret
	}
	return *o.ProfileMedium
}

// GetProfileMediumOk returns a tuple with the ProfileMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetProfileMediumOk() (*string, bool) {
	if o == nil || o.ProfileMedium == nil {
		return nil, false
	}
	return o.ProfileMedium, true
}

// HasProfileMedium returns a boolean if a field has been set.
func (o *SummaryAthlete) HasProfileMedium() bool {
	if o != nil && o.ProfileMedium != nil {
		return true
	}

	return false
}

// SetProfileMedium gets a reference to the given string and assigns it to the ProfileMedium field.
func (o *SummaryAthlete) SetProfileMedium(v string) {
	o.ProfileMedium = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SummaryAthlete) GetProfile() string {
	if o == nil || o.Profile == nil {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetProfileOk() (*string, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SummaryAthlete) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *SummaryAthlete) SetProfile(v string) {
	o.Profile = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SummaryAthlete) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SummaryAthlete) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SummaryAthlete) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SummaryAthlete) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SummaryAthlete) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SummaryAthlete) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *SummaryAthlete) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *SummaryAthlete) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *SummaryAthlete) SetCountry(v string) {
	o.Country = &v
}

// GetSex returns the Sex field value if set, zero value otherwise.
func (o *SummaryAthlete) GetSex() string {
	if o == nil || o.Sex == nil {
		var ret string
		return ret
	}
	return *o.Sex
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetSexOk() (*string, bool) {
	if o == nil || o.Sex == nil {
		return nil, false
	}
	return o.Sex, true
}

// HasSex returns a boolean if a field has been set.
func (o *SummaryAthlete) HasSex() bool {
	if o != nil && o.Sex != nil {
		return true
	}

	return false
}

// SetSex gets a reference to the given string and assigns it to the Sex field.
func (o *SummaryAthlete) SetSex(v string) {
	o.Sex = &v
}

// GetPremium returns the Premium field value if set, zero value otherwise.
func (o *SummaryAthlete) GetPremium() bool {
	if o == nil || o.Premium == nil {
		var ret bool
		return ret
	}
	return *o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetPremiumOk() (*bool, bool) {
	if o == nil || o.Premium == nil {
		return nil, false
	}
	return o.Premium, true
}

// HasPremium returns a boolean if a field has been set.
func (o *SummaryAthlete) HasPremium() bool {
	if o != nil && o.Premium != nil {
		return true
	}

	return false
}

// SetPremium gets a reference to the given bool and assigns it to the Premium field.
func (o *SummaryAthlete) SetPremium(v bool) {
	o.Premium = &v
}

// GetSummit returns the Summit field value if set, zero value otherwise.
func (o *SummaryAthlete) GetSummit() bool {
	if o == nil || o.Summit == nil {
		var ret bool
		return ret
	}
	return *o.Summit
}

// GetSummitOk returns a tuple with the Summit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetSummitOk() (*bool, bool) {
	if o == nil || o.Summit == nil {
		return nil, false
	}
	return o.Summit, true
}

// HasSummit returns a boolean if a field has been set.
func (o *SummaryAthlete) HasSummit() bool {
	if o != nil && o.Summit != nil {
		return true
	}

	return false
}

// SetSummit gets a reference to the given bool and assigns it to the Summit field.
func (o *SummaryAthlete) SetSummit(v bool) {
	o.Summit = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SummaryAthlete) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SummaryAthlete) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SummaryAthlete) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SummaryAthlete) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthlete) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SummaryAthlete) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SummaryAthlete) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o SummaryAthlete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ResourceState != nil {
		toSerialize["resource_state"] = o.ResourceState
	}
	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}
	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}
	if o.ProfileMedium != nil {
		toSerialize["profile_medium"] = o.ProfileMedium
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Sex != nil {
		toSerialize["sex"] = o.Sex
	}
	if o.Premium != nil {
		toSerialize["premium"] = o.Premium
	}
	if o.Summit != nil {
		toSerialize["summit"] = o.Summit
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSummaryAthlete struct {
	value *SummaryAthlete
	isSet bool
}

func (v NullableSummaryAthlete) Get() *SummaryAthlete {
	return v.value
}

func (v *NullableSummaryAthlete) Set(val *SummaryAthlete) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryAthlete) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryAthlete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryAthlete(val *SummaryAthlete) *NullableSummaryAthlete {
	return &NullableSummaryAthlete{value: val, isSet: true}
}

func (v NullableSummaryAthlete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryAthlete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


