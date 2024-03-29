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

// TimedZoneRange A union type representing the time spent in a given zone.
type TimedZoneRange struct {
	// The minimum value in the range.
	Min *int32 `json:"min,omitempty"`
	// The maximum value in the range.
	Max *int32 `json:"max,omitempty"`
	// The number of seconds spent in this zone
	Time *int32 `json:"time,omitempty"`
}

// NewTimedZoneRange instantiates a new TimedZoneRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimedZoneRange() *TimedZoneRange {
	this := TimedZoneRange{}
	return &this
}

// NewTimedZoneRangeWithDefaults instantiates a new TimedZoneRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimedZoneRangeWithDefaults() *TimedZoneRange {
	this := TimedZoneRange{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *TimedZoneRange) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimedZoneRange) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *TimedZoneRange) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *TimedZoneRange) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *TimedZoneRange) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimedZoneRange) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *TimedZoneRange) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *TimedZoneRange) SetMax(v int32) {
	o.Max = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TimedZoneRange) GetTime() int32 {
	if o == nil || o.Time == nil {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimedZoneRange) GetTimeOk() (*int32, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TimedZoneRange) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *TimedZoneRange) SetTime(v int32) {
	o.Time = &v
}

func (o TimedZoneRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableTimedZoneRange struct {
	value *TimedZoneRange
	isSet bool
}

func (v NullableTimedZoneRange) Get() *TimedZoneRange {
	return v.value
}

func (v *NullableTimedZoneRange) Set(val *TimedZoneRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTimedZoneRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTimedZoneRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimedZoneRange(val *TimedZoneRange) *NullableTimedZoneRange {
	return &NullableTimedZoneRange{value: val, isSet: true}
}

func (v NullableTimedZoneRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimedZoneRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


