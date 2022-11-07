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

// SummaryAthleteAllOf struct for SummaryAthleteAllOf
type SummaryAthleteAllOf struct {
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

// NewSummaryAthleteAllOf instantiates a new SummaryAthleteAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryAthleteAllOf() *SummaryAthleteAllOf {
	this := SummaryAthleteAllOf{}
	return &this
}

// NewSummaryAthleteAllOfWithDefaults instantiates a new SummaryAthleteAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryAthleteAllOfWithDefaults() *SummaryAthleteAllOf {
	this := SummaryAthleteAllOf{}
	return &this
}

// GetResourceState returns the ResourceState field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetResourceState() int32 {
	if o == nil || isNil(o.ResourceState) {
		var ret int32
		return ret
	}
	return *o.ResourceState
}

// GetResourceStateOk returns a tuple with the ResourceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetResourceStateOk() (*int32, bool) {
	if o == nil || isNil(o.ResourceState) {
    return nil, false
	}
	return o.ResourceState, true
}

// HasResourceState returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasResourceState() bool {
	if o != nil && !isNil(o.ResourceState) {
		return true
	}

	return false
}

// SetResourceState gets a reference to the given int32 and assigns it to the ResourceState field.
func (o *SummaryAthleteAllOf) SetResourceState(v int32) {
	o.ResourceState = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetFirstname() string {
	if o == nil || isNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetFirstnameOk() (*string, bool) {
	if o == nil || isNil(o.Firstname) {
    return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasFirstname() bool {
	if o != nil && !isNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *SummaryAthleteAllOf) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetLastname() string {
	if o == nil || isNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetLastnameOk() (*string, bool) {
	if o == nil || isNil(o.Lastname) {
    return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasLastname() bool {
	if o != nil && !isNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *SummaryAthleteAllOf) SetLastname(v string) {
	o.Lastname = &v
}

// GetProfileMedium returns the ProfileMedium field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetProfileMedium() string {
	if o == nil || isNil(o.ProfileMedium) {
		var ret string
		return ret
	}
	return *o.ProfileMedium
}

// GetProfileMediumOk returns a tuple with the ProfileMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetProfileMediumOk() (*string, bool) {
	if o == nil || isNil(o.ProfileMedium) {
    return nil, false
	}
	return o.ProfileMedium, true
}

// HasProfileMedium returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasProfileMedium() bool {
	if o != nil && !isNil(o.ProfileMedium) {
		return true
	}

	return false
}

// SetProfileMedium gets a reference to the given string and assigns it to the ProfileMedium field.
func (o *SummaryAthleteAllOf) SetProfileMedium(v string) {
	o.ProfileMedium = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetProfile() string {
	if o == nil || isNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetProfileOk() (*string, bool) {
	if o == nil || isNil(o.Profile) {
    return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasProfile() bool {
	if o != nil && !isNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *SummaryAthleteAllOf) SetProfile(v string) {
	o.Profile = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
    return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SummaryAthleteAllOf) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SummaryAthleteAllOf) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *SummaryAthleteAllOf) SetCountry(v string) {
	o.Country = &v
}

// GetSex returns the Sex field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetSex() string {
	if o == nil || isNil(o.Sex) {
		var ret string
		return ret
	}
	return *o.Sex
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetSexOk() (*string, bool) {
	if o == nil || isNil(o.Sex) {
    return nil, false
	}
	return o.Sex, true
}

// HasSex returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasSex() bool {
	if o != nil && !isNil(o.Sex) {
		return true
	}

	return false
}

// SetSex gets a reference to the given string and assigns it to the Sex field.
func (o *SummaryAthleteAllOf) SetSex(v string) {
	o.Sex = &v
}

// GetPremium returns the Premium field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetPremium() bool {
	if o == nil || isNil(o.Premium) {
		var ret bool
		return ret
	}
	return *o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetPremiumOk() (*bool, bool) {
	if o == nil || isNil(o.Premium) {
    return nil, false
	}
	return o.Premium, true
}

// HasPremium returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasPremium() bool {
	if o != nil && !isNil(o.Premium) {
		return true
	}

	return false
}

// SetPremium gets a reference to the given bool and assigns it to the Premium field.
func (o *SummaryAthleteAllOf) SetPremium(v bool) {
	o.Premium = &v
}

// GetSummit returns the Summit field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetSummit() bool {
	if o == nil || isNil(o.Summit) {
		var ret bool
		return ret
	}
	return *o.Summit
}

// GetSummitOk returns a tuple with the Summit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetSummitOk() (*bool, bool) {
	if o == nil || isNil(o.Summit) {
    return nil, false
	}
	return o.Summit, true
}

// HasSummit returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasSummit() bool {
	if o != nil && !isNil(o.Summit) {
		return true
	}

	return false
}

// SetSummit gets a reference to the given bool and assigns it to the Summit field.
func (o *SummaryAthleteAllOf) SetSummit(v bool) {
	o.Summit = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SummaryAthleteAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SummaryAthleteAllOf) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryAthleteAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SummaryAthleteAllOf) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SummaryAthleteAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o SummaryAthleteAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResourceState) {
		toSerialize["resource_state"] = o.ResourceState
	}
	if !isNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !isNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !isNil(o.ProfileMedium) {
		toSerialize["profile_medium"] = o.ProfileMedium
	}
	if !isNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.Sex) {
		toSerialize["sex"] = o.Sex
	}
	if !isNil(o.Premium) {
		toSerialize["premium"] = o.Premium
	}
	if !isNil(o.Summit) {
		toSerialize["summit"] = o.Summit
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSummaryAthleteAllOf struct {
	value *SummaryAthleteAllOf
	isSet bool
}

func (v NullableSummaryAthleteAllOf) Get() *SummaryAthleteAllOf {
	return v.value
}

func (v *NullableSummaryAthleteAllOf) Set(val *SummaryAthleteAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryAthleteAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryAthleteAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryAthleteAllOf(val *SummaryAthleteAllOf) *NullableSummaryAthleteAllOf {
	return &NullableSummaryAthleteAllOf{value: val, isSet: true}
}

func (v NullableSummaryAthleteAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryAthleteAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


