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

// TimeStreamAllOf struct for TimeStreamAllOf
type TimeStreamAllOf struct {
	// The sequence of time values for this stream, in seconds
	Data []int32 `json:"data,omitempty"`
}

// NewTimeStreamAllOf instantiates a new TimeStreamAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeStreamAllOf() *TimeStreamAllOf {
	this := TimeStreamAllOf{}
	return &this
}

// NewTimeStreamAllOfWithDefaults instantiates a new TimeStreamAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeStreamAllOfWithDefaults() *TimeStreamAllOf {
	this := TimeStreamAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TimeStreamAllOf) GetData() []int32 {
	if o == nil || o.Data == nil {
		var ret []int32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeStreamAllOf) GetDataOk() ([]int32, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TimeStreamAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []int32 and assigns it to the Data field.
func (o *TimeStreamAllOf) SetData(v []int32) {
	o.Data = v
}

func (o TimeStreamAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTimeStreamAllOf struct {
	value *TimeStreamAllOf
	isSet bool
}

func (v NullableTimeStreamAllOf) Get() *TimeStreamAllOf {
	return v.value
}

func (v *NullableTimeStreamAllOf) Set(val *TimeStreamAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeStreamAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeStreamAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeStreamAllOf(val *TimeStreamAllOf) *NullableTimeStreamAllOf {
	return &NullableTimeStreamAllOf{value: val, isSet: true}
}

func (v NullableTimeStreamAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeStreamAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


