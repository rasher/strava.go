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

// CadenceStream struct for CadenceStream
type CadenceStream struct {
	// The number of data points in this stream
	OriginalSize *int32 `json:"original_size,omitempty"`
	// The level of detail (sampling) in which this stream was returned
	Resolution *string `json:"resolution,omitempty"`
	// The base series used in the case the stream was downsampled
	SeriesType *string `json:"series_type,omitempty"`
	// The sequence of cadence values for this stream, in rotations per minute
	Data []int32 `json:"data,omitempty"`
}

// NewCadenceStream instantiates a new CadenceStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCadenceStream() *CadenceStream {
	this := CadenceStream{}
	return &this
}

// NewCadenceStreamWithDefaults instantiates a new CadenceStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCadenceStreamWithDefaults() *CadenceStream {
	this := CadenceStream{}
	return &this
}

// GetOriginalSize returns the OriginalSize field value if set, zero value otherwise.
func (o *CadenceStream) GetOriginalSize() int32 {
	if o == nil || o.OriginalSize == nil {
		var ret int32
		return ret
	}
	return *o.OriginalSize
}

// GetOriginalSizeOk returns a tuple with the OriginalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CadenceStream) GetOriginalSizeOk() (*int32, bool) {
	if o == nil || o.OriginalSize == nil {
		return nil, false
	}
	return o.OriginalSize, true
}

// HasOriginalSize returns a boolean if a field has been set.
func (o *CadenceStream) HasOriginalSize() bool {
	if o != nil && o.OriginalSize != nil {
		return true
	}

	return false
}

// SetOriginalSize gets a reference to the given int32 and assigns it to the OriginalSize field.
func (o *CadenceStream) SetOriginalSize(v int32) {
	o.OriginalSize = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *CadenceStream) GetResolution() string {
	if o == nil || o.Resolution == nil {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CadenceStream) GetResolutionOk() (*string, bool) {
	if o == nil || o.Resolution == nil {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *CadenceStream) HasResolution() bool {
	if o != nil && o.Resolution != nil {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *CadenceStream) SetResolution(v string) {
	o.Resolution = &v
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *CadenceStream) GetSeriesType() string {
	if o == nil || o.SeriesType == nil {
		var ret string
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CadenceStream) GetSeriesTypeOk() (*string, bool) {
	if o == nil || o.SeriesType == nil {
		return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *CadenceStream) HasSeriesType() bool {
	if o != nil && o.SeriesType != nil {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given string and assigns it to the SeriesType field.
func (o *CadenceStream) SetSeriesType(v string) {
	o.SeriesType = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CadenceStream) GetData() []int32 {
	if o == nil || o.Data == nil {
		var ret []int32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CadenceStream) GetDataOk() ([]int32, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CadenceStream) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []int32 and assigns it to the Data field.
func (o *CadenceStream) SetData(v []int32) {
	o.Data = v
}

func (o CadenceStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OriginalSize != nil {
		toSerialize["original_size"] = o.OriginalSize
	}
	if o.Resolution != nil {
		toSerialize["resolution"] = o.Resolution
	}
	if o.SeriesType != nil {
		toSerialize["series_type"] = o.SeriesType
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCadenceStream struct {
	value *CadenceStream
	isSet bool
}

func (v NullableCadenceStream) Get() *CadenceStream {
	return v.value
}

func (v *NullableCadenceStream) Set(val *CadenceStream) {
	v.value = val
	v.isSet = true
}

func (v NullableCadenceStream) IsSet() bool {
	return v.isSet
}

func (v *NullableCadenceStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCadenceStream(val *CadenceStream) *NullableCadenceStream {
	return &NullableCadenceStream{value: val, isSet: true}
}

func (v NullableCadenceStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCadenceStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


