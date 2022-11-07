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

// MovingStream struct for MovingStream
type MovingStream struct {
	// The number of data points in this stream
	OriginalSize *int32 `json:"original_size,omitempty"`
	// The level of detail (sampling) in which this stream was returned
	Resolution *string `json:"resolution,omitempty"`
	// The base series used in the case the stream was downsampled
	SeriesType *string `json:"series_type,omitempty"`
	// The sequence of moving values for this stream, as boolean values
	Data []bool `json:"data,omitempty"`
}

// NewMovingStream instantiates a new MovingStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovingStream() *MovingStream {
	this := MovingStream{}
	return &this
}

// NewMovingStreamWithDefaults instantiates a new MovingStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovingStreamWithDefaults() *MovingStream {
	this := MovingStream{}
	return &this
}

// GetOriginalSize returns the OriginalSize field value if set, zero value otherwise.
func (o *MovingStream) GetOriginalSize() int32 {
	if o == nil || isNil(o.OriginalSize) {
		var ret int32
		return ret
	}
	return *o.OriginalSize
}

// GetOriginalSizeOk returns a tuple with the OriginalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovingStream) GetOriginalSizeOk() (*int32, bool) {
	if o == nil || isNil(o.OriginalSize) {
    return nil, false
	}
	return o.OriginalSize, true
}

// HasOriginalSize returns a boolean if a field has been set.
func (o *MovingStream) HasOriginalSize() bool {
	if o != nil && !isNil(o.OriginalSize) {
		return true
	}

	return false
}

// SetOriginalSize gets a reference to the given int32 and assigns it to the OriginalSize field.
func (o *MovingStream) SetOriginalSize(v int32) {
	o.OriginalSize = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *MovingStream) GetResolution() string {
	if o == nil || isNil(o.Resolution) {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovingStream) GetResolutionOk() (*string, bool) {
	if o == nil || isNil(o.Resolution) {
    return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *MovingStream) HasResolution() bool {
	if o != nil && !isNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *MovingStream) SetResolution(v string) {
	o.Resolution = &v
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *MovingStream) GetSeriesType() string {
	if o == nil || isNil(o.SeriesType) {
		var ret string
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovingStream) GetSeriesTypeOk() (*string, bool) {
	if o == nil || isNil(o.SeriesType) {
    return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *MovingStream) HasSeriesType() bool {
	if o != nil && !isNil(o.SeriesType) {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given string and assigns it to the SeriesType field.
func (o *MovingStream) SetSeriesType(v string) {
	o.SeriesType = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MovingStream) GetData() []bool {
	if o == nil || isNil(o.Data) {
		var ret []bool
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovingStream) GetDataOk() ([]bool, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MovingStream) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []bool and assigns it to the Data field.
func (o *MovingStream) SetData(v []bool) {
	o.Data = v
}

func (o MovingStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OriginalSize) {
		toSerialize["original_size"] = o.OriginalSize
	}
	if !isNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !isNil(o.SeriesType) {
		toSerialize["series_type"] = o.SeriesType
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMovingStream struct {
	value *MovingStream
	isSet bool
}

func (v NullableMovingStream) Get() *MovingStream {
	return v.value
}

func (v *NullableMovingStream) Set(val *MovingStream) {
	v.value = val
	v.isSet = true
}

func (v NullableMovingStream) IsSet() bool {
	return v.isSet
}

func (v *NullableMovingStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovingStream(val *MovingStream) *NullableMovingStream {
	return &NullableMovingStream{value: val, isSet: true}
}

func (v NullableMovingStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovingStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


