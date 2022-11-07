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

// Split struct for Split
type Split struct {
	// The average speed of this split, in meters per second
	AverageSpeed *float32 `json:"average_speed,omitempty"`
	// The distance of this split, in meters
	Distance *float32 `json:"distance,omitempty"`
	// The elapsed time of this split, in seconds
	ElapsedTime *int32 `json:"elapsed_time,omitempty"`
	// The elevation difference of this split, in meters
	ElevationDifference *float32 `json:"elevation_difference,omitempty"`
	// The pacing zone of this split
	PaceZone *int32 `json:"pace_zone,omitempty"`
	// The moving time of this split, in seconds
	MovingTime *int32 `json:"moving_time,omitempty"`
	// N/A
	Split *int32 `json:"split,omitempty"`
}

// NewSplit instantiates a new Split object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplit() *Split {
	this := Split{}
	return &this
}

// NewSplitWithDefaults instantiates a new Split object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitWithDefaults() *Split {
	this := Split{}
	return &this
}

// GetAverageSpeed returns the AverageSpeed field value if set, zero value otherwise.
func (o *Split) GetAverageSpeed() float32 {
	if o == nil || o.AverageSpeed == nil {
		var ret float32
		return ret
	}
	return *o.AverageSpeed
}

// GetAverageSpeedOk returns a tuple with the AverageSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetAverageSpeedOk() (*float32, bool) {
	if o == nil || o.AverageSpeed == nil {
		return nil, false
	}
	return o.AverageSpeed, true
}

// HasAverageSpeed returns a boolean if a field has been set.
func (o *Split) HasAverageSpeed() bool {
	if o != nil && o.AverageSpeed != nil {
		return true
	}

	return false
}

// SetAverageSpeed gets a reference to the given float32 and assigns it to the AverageSpeed field.
func (o *Split) SetAverageSpeed(v float32) {
	o.AverageSpeed = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *Split) GetDistance() float32 {
	if o == nil || o.Distance == nil {
		var ret float32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetDistanceOk() (*float32, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *Split) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float32 and assigns it to the Distance field.
func (o *Split) SetDistance(v float32) {
	o.Distance = &v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *Split) GetElapsedTime() int32 {
	if o == nil || o.ElapsedTime == nil {
		var ret int32
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetElapsedTimeOk() (*int32, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *Split) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int32 and assigns it to the ElapsedTime field.
func (o *Split) SetElapsedTime(v int32) {
	o.ElapsedTime = &v
}

// GetElevationDifference returns the ElevationDifference field value if set, zero value otherwise.
func (o *Split) GetElevationDifference() float32 {
	if o == nil || o.ElevationDifference == nil {
		var ret float32
		return ret
	}
	return *o.ElevationDifference
}

// GetElevationDifferenceOk returns a tuple with the ElevationDifference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetElevationDifferenceOk() (*float32, bool) {
	if o == nil || o.ElevationDifference == nil {
		return nil, false
	}
	return o.ElevationDifference, true
}

// HasElevationDifference returns a boolean if a field has been set.
func (o *Split) HasElevationDifference() bool {
	if o != nil && o.ElevationDifference != nil {
		return true
	}

	return false
}

// SetElevationDifference gets a reference to the given float32 and assigns it to the ElevationDifference field.
func (o *Split) SetElevationDifference(v float32) {
	o.ElevationDifference = &v
}

// GetPaceZone returns the PaceZone field value if set, zero value otherwise.
func (o *Split) GetPaceZone() int32 {
	if o == nil || o.PaceZone == nil {
		var ret int32
		return ret
	}
	return *o.PaceZone
}

// GetPaceZoneOk returns a tuple with the PaceZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetPaceZoneOk() (*int32, bool) {
	if o == nil || o.PaceZone == nil {
		return nil, false
	}
	return o.PaceZone, true
}

// HasPaceZone returns a boolean if a field has been set.
func (o *Split) HasPaceZone() bool {
	if o != nil && o.PaceZone != nil {
		return true
	}

	return false
}

// SetPaceZone gets a reference to the given int32 and assigns it to the PaceZone field.
func (o *Split) SetPaceZone(v int32) {
	o.PaceZone = &v
}

// GetMovingTime returns the MovingTime field value if set, zero value otherwise.
func (o *Split) GetMovingTime() int32 {
	if o == nil || o.MovingTime == nil {
		var ret int32
		return ret
	}
	return *o.MovingTime
}

// GetMovingTimeOk returns a tuple with the MovingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetMovingTimeOk() (*int32, bool) {
	if o == nil || o.MovingTime == nil {
		return nil, false
	}
	return o.MovingTime, true
}

// HasMovingTime returns a boolean if a field has been set.
func (o *Split) HasMovingTime() bool {
	if o != nil && o.MovingTime != nil {
		return true
	}

	return false
}

// SetMovingTime gets a reference to the given int32 and assigns it to the MovingTime field.
func (o *Split) SetMovingTime(v int32) {
	o.MovingTime = &v
}

// GetSplit returns the Split field value if set, zero value otherwise.
func (o *Split) GetSplit() int32 {
	if o == nil || o.Split == nil {
		var ret int32
		return ret
	}
	return *o.Split
}

// GetSplitOk returns a tuple with the Split field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetSplitOk() (*int32, bool) {
	if o == nil || o.Split == nil {
		return nil, false
	}
	return o.Split, true
}

// HasSplit returns a boolean if a field has been set.
func (o *Split) HasSplit() bool {
	if o != nil && o.Split != nil {
		return true
	}

	return false
}

// SetSplit gets a reference to the given int32 and assigns it to the Split field.
func (o *Split) SetSplit(v int32) {
	o.Split = &v
}

func (o Split) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AverageSpeed != nil {
		toSerialize["average_speed"] = o.AverageSpeed
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.ElapsedTime != nil {
		toSerialize["elapsed_time"] = o.ElapsedTime
	}
	if o.ElevationDifference != nil {
		toSerialize["elevation_difference"] = o.ElevationDifference
	}
	if o.PaceZone != nil {
		toSerialize["pace_zone"] = o.PaceZone
	}
	if o.MovingTime != nil {
		toSerialize["moving_time"] = o.MovingTime
	}
	if o.Split != nil {
		toSerialize["split"] = o.Split
	}
	return json.Marshal(toSerialize)
}

type NullableSplit struct {
	value *Split
	isSet bool
}

func (v NullableSplit) Get() *Split {
	return v.value
}

func (v *NullableSplit) Set(val *Split) {
	v.value = val
	v.isSet = true
}

func (v NullableSplit) IsSet() bool {
	return v.isSet
}

func (v *NullableSplit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplit(val *Split) *NullableSplit {
	return &NullableSplit{value: val, isSet: true}
}

func (v NullableSplit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


