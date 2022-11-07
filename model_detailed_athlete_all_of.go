/*
Strava API v3

The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package strava

import (
	"encoding/json"
)

// DetailedAthleteAllOf struct for DetailedAthleteAllOf
type DetailedAthleteAllOf struct {
	// The athlete's follower count.
	FollowerCount *int32 `json:"follower_count,omitempty"`
	// The athlete's friend count.
	FriendCount *int32 `json:"friend_count,omitempty"`
	// The athlete's preferred unit system.
	MeasurementPreference *string `json:"measurement_preference,omitempty"`
	// The athlete's FTP (Functional Threshold Power).
	Ftp *int32 `json:"ftp,omitempty"`
	// The athlete's weight.
	Weight *float32 `json:"weight,omitempty"`
	// The athlete's clubs.
	Clubs []SummaryClub `json:"clubs,omitempty"`
	// The athlete's bikes.
	Bikes []SummaryGear `json:"bikes,omitempty"`
	// The athlete's shoes.
	Shoes []SummaryGear `json:"shoes,omitempty"`
}

// NewDetailedAthleteAllOf instantiates a new DetailedAthleteAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedAthleteAllOf() *DetailedAthleteAllOf {
	this := DetailedAthleteAllOf{}
	return &this
}

// NewDetailedAthleteAllOfWithDefaults instantiates a new DetailedAthleteAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedAthleteAllOfWithDefaults() *DetailedAthleteAllOf {
	this := DetailedAthleteAllOf{}
	return &this
}

// GetFollowerCount returns the FollowerCount field value if set, zero value otherwise.
func (o *DetailedAthleteAllOf) GetFollowerCount() int32 {
	if o == nil || isNil(o.FollowerCount) {
		var ret int32
		return ret
	}
	return *o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedAthleteAllOf) GetFollowerCountOk() (*int32, bool) {
	if o == nil || isNil(o.FollowerCount) {
    return nil, false
	}
	return o.FollowerCount, true
}

// HasFollowerCount returns a boolean if a field has been set.
func (o *DetailedAthleteAllOf) HasFollowerCount() bool {
	if o != nil && !isNil(o.FollowerCount) {
		return true
	}

	return false
}

// SetFollowerCount gets a reference to the given int32 and assigns it to the FollowerCount field.
func (o *DetailedAthleteAllOf) SetFollowerCount(v int32) {
	o.FollowerCount = &v
}

// GetFriendCount returns the FriendCount field value if set, zero value otherwise.
func (o *DetailedAthleteAllOf) GetFriendCount() int32 {
	if o == nil || isNil(o.FriendCount) {
		var ret int32
		return ret
	}
	return *o.FriendCount
}

// GetFriendCountOk returns a tuple with the FriendCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedAthleteAllOf) GetFriendCountOk() (*int32, bool) {
	if o == nil || isNil(o.FriendCount) {
    return nil, false
	}
	return o.FriendCount, true
}

// HasFriendCount returns a boolean if a field has been set.
func (o *DetailedAthleteAllOf) HasFriendCount() bool {
	if o != nil && !isNil(o.FriendCount) {
		return true
	}

	return false
}

// SetFriendCount gets a reference to the given int32 and assigns it to the FriendCount field.
func (o *DetailedAthleteAllOf) SetFriendCount(v int32) {
	o.FriendCount = &v
}

// GetMeasurementPreference returns the MeasurementPreference field value if set, zero value otherwise.
func (o *DetailedAthleteAllOf) GetMeasurementPreference() string {
	if o == nil || isNil(o.MeasurementPreference) {
		var ret string
		return ret
	}
	return *o.MeasurementPreference
}

// GetMeasurementPreferenceOk returns a tuple with the MeasurementPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedAthleteAllOf) GetMeasurementPreferenceOk() (*string, bool) {
	if o == nil || isNil(o.MeasurementPreference) {
    return nil, false
	}
	return o.MeasurementPreference, true
}

// HasMeasurementPreference returns a boolean if a field has been set.
func (o *DetailedAthleteAllOf) HasMeasurementPreference() bool {
	if o != nil && !isNil(o.MeasurementPreference) {
		return true
	}

	return false
}

// SetMeasurementPreference gets a reference to the given string and assigns it to the MeasurementPreference field.
func (o *DetailedAthleteAllOf) SetMeasurementPreference(v string) {
	o.MeasurementPreference = &v
}

// GetFtp returns the Ftp field value if set, zero value otherwise.
func (o *DetailedAthleteAllOf) GetFtp() int32 {
	if o == nil || isNil(o.Ftp) {
		var ret int32
		return ret
	}
	return *o.Ftp
}

// GetFtpOk returns a tuple with the Ftp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedAthleteAllOf) GetFtpOk() (*int32, bool) {
	if o == nil || isNil(o.Ftp) {
    return nil, false
	}
	return o.Ftp, true
}

// HasFtp returns a boolean if a field has been set.
func (o *DetailedAthleteAllOf) HasFtp() bool {
	if o != nil && !isNil(o.Ftp) {
		return true
	}

	return false
}

// SetFtp gets a reference to the given int32 and assigns it to the Ftp field.
func (o *DetailedAthleteAllOf) SetFtp(v int32) {
	o.Ftp = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *DetailedAthleteAllOf) GetWeight() float32 {
	if o == nil || isNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedAthleteAllOf) GetWeightOk() (*float32, bool) {
	if o == nil || isNil(o.Weight) {
    return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *DetailedAthleteAllOf) HasWeight() bool {
	if o != nil && !isNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *DetailedAthleteAllOf) SetWeight(v float32) {
	o.Weight = &v
}

// GetClubs returns the Clubs field value if set, zero value otherwise.
func (o *DetailedAthleteAllOf) GetClubs() []SummaryClub {
	if o == nil || isNil(o.Clubs) {
		var ret []SummaryClub
		return ret
	}
	return o.Clubs
}

// GetClubsOk returns a tuple with the Clubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedAthleteAllOf) GetClubsOk() ([]SummaryClub, bool) {
	if o == nil || isNil(o.Clubs) {
    return nil, false
	}
	return o.Clubs, true
}

// HasClubs returns a boolean if a field has been set.
func (o *DetailedAthleteAllOf) HasClubs() bool {
	if o != nil && !isNil(o.Clubs) {
		return true
	}

	return false
}

// SetClubs gets a reference to the given []SummaryClub and assigns it to the Clubs field.
func (o *DetailedAthleteAllOf) SetClubs(v []SummaryClub) {
	o.Clubs = v
}

// GetBikes returns the Bikes field value if set, zero value otherwise.
func (o *DetailedAthleteAllOf) GetBikes() []SummaryGear {
	if o == nil || isNil(o.Bikes) {
		var ret []SummaryGear
		return ret
	}
	return o.Bikes
}

// GetBikesOk returns a tuple with the Bikes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedAthleteAllOf) GetBikesOk() ([]SummaryGear, bool) {
	if o == nil || isNil(o.Bikes) {
    return nil, false
	}
	return o.Bikes, true
}

// HasBikes returns a boolean if a field has been set.
func (o *DetailedAthleteAllOf) HasBikes() bool {
	if o != nil && !isNil(o.Bikes) {
		return true
	}

	return false
}

// SetBikes gets a reference to the given []SummaryGear and assigns it to the Bikes field.
func (o *DetailedAthleteAllOf) SetBikes(v []SummaryGear) {
	o.Bikes = v
}

// GetShoes returns the Shoes field value if set, zero value otherwise.
func (o *DetailedAthleteAllOf) GetShoes() []SummaryGear {
	if o == nil || isNil(o.Shoes) {
		var ret []SummaryGear
		return ret
	}
	return o.Shoes
}

// GetShoesOk returns a tuple with the Shoes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedAthleteAllOf) GetShoesOk() ([]SummaryGear, bool) {
	if o == nil || isNil(o.Shoes) {
    return nil, false
	}
	return o.Shoes, true
}

// HasShoes returns a boolean if a field has been set.
func (o *DetailedAthleteAllOf) HasShoes() bool {
	if o != nil && !isNil(o.Shoes) {
		return true
	}

	return false
}

// SetShoes gets a reference to the given []SummaryGear and assigns it to the Shoes field.
func (o *DetailedAthleteAllOf) SetShoes(v []SummaryGear) {
	o.Shoes = v
}

func (o DetailedAthleteAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FollowerCount) {
		toSerialize["follower_count"] = o.FollowerCount
	}
	if !isNil(o.FriendCount) {
		toSerialize["friend_count"] = o.FriendCount
	}
	if !isNil(o.MeasurementPreference) {
		toSerialize["measurement_preference"] = o.MeasurementPreference
	}
	if !isNil(o.Ftp) {
		toSerialize["ftp"] = o.Ftp
	}
	if !isNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !isNil(o.Clubs) {
		toSerialize["clubs"] = o.Clubs
	}
	if !isNil(o.Bikes) {
		toSerialize["bikes"] = o.Bikes
	}
	if !isNil(o.Shoes) {
		toSerialize["shoes"] = o.Shoes
	}
	return json.Marshal(toSerialize)
}

type NullableDetailedAthleteAllOf struct {
	value *DetailedAthleteAllOf
	isSet bool
}

func (v NullableDetailedAthleteAllOf) Get() *DetailedAthleteAllOf {
	return v.value
}

func (v *NullableDetailedAthleteAllOf) Set(val *DetailedAthleteAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedAthleteAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedAthleteAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedAthleteAllOf(val *DetailedAthleteAllOf) *NullableDetailedAthleteAllOf {
	return &NullableDetailedAthleteAllOf{value: val, isSet: true}
}

func (v NullableDetailedAthleteAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedAthleteAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


