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

// DetailedGear struct for DetailedGear
type DetailedGear struct {
	// The gear's unique identifier.
	Id *string `json:"id,omitempty"`
	// Resource state, indicates level of detail. Possible values: 2 -> \"summary\", 3 -> \"detail\"
	ResourceState *int32 `json:"resource_state,omitempty"`
	// Whether this gear's is the owner's default one.
	Primary *bool `json:"primary,omitempty"`
	// The gear's name.
	Name *string `json:"name,omitempty"`
	// The distance logged with this gear.
	Distance *float32 `json:"distance,omitempty"`
	// The gear's brand name.
	BrandName *string `json:"brand_name,omitempty"`
	// The gear's model name.
	ModelName *string `json:"model_name,omitempty"`
	// The gear's frame type (bike only).
	FrameType *int32 `json:"frame_type,omitempty"`
	// The gear's description.
	Description *string `json:"description,omitempty"`
}

// NewDetailedGear instantiates a new DetailedGear object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedGear() *DetailedGear {
	this := DetailedGear{}
	return &this
}

// NewDetailedGearWithDefaults instantiates a new DetailedGear object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedGearWithDefaults() *DetailedGear {
	this := DetailedGear{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DetailedGear) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DetailedGear) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DetailedGear) SetId(v string) {
	o.Id = &v
}

// GetResourceState returns the ResourceState field value if set, zero value otherwise.
func (o *DetailedGear) GetResourceState() int32 {
	if o == nil || isNil(o.ResourceState) {
		var ret int32
		return ret
	}
	return *o.ResourceState
}

// GetResourceStateOk returns a tuple with the ResourceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetResourceStateOk() (*int32, bool) {
	if o == nil || isNil(o.ResourceState) {
    return nil, false
	}
	return o.ResourceState, true
}

// HasResourceState returns a boolean if a field has been set.
func (o *DetailedGear) HasResourceState() bool {
	if o != nil && !isNil(o.ResourceState) {
		return true
	}

	return false
}

// SetResourceState gets a reference to the given int32 and assigns it to the ResourceState field.
func (o *DetailedGear) SetResourceState(v int32) {
	o.ResourceState = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *DetailedGear) GetPrimary() bool {
	if o == nil || isNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetPrimaryOk() (*bool, bool) {
	if o == nil || isNil(o.Primary) {
    return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *DetailedGear) HasPrimary() bool {
	if o != nil && !isNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *DetailedGear) SetPrimary(v bool) {
	o.Primary = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DetailedGear) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DetailedGear) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DetailedGear) SetName(v string) {
	o.Name = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *DetailedGear) GetDistance() float32 {
	if o == nil || isNil(o.Distance) {
		var ret float32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetDistanceOk() (*float32, bool) {
	if o == nil || isNil(o.Distance) {
    return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *DetailedGear) HasDistance() bool {
	if o != nil && !isNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float32 and assigns it to the Distance field.
func (o *DetailedGear) SetDistance(v float32) {
	o.Distance = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *DetailedGear) GetBrandName() string {
	if o == nil || isNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetBrandNameOk() (*string, bool) {
	if o == nil || isNil(o.BrandName) {
    return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *DetailedGear) HasBrandName() bool {
	if o != nil && !isNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *DetailedGear) SetBrandName(v string) {
	o.BrandName = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *DetailedGear) GetModelName() string {
	if o == nil || isNil(o.ModelName) {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetModelNameOk() (*string, bool) {
	if o == nil || isNil(o.ModelName) {
    return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *DetailedGear) HasModelName() bool {
	if o != nil && !isNil(o.ModelName) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *DetailedGear) SetModelName(v string) {
	o.ModelName = &v
}

// GetFrameType returns the FrameType field value if set, zero value otherwise.
func (o *DetailedGear) GetFrameType() int32 {
	if o == nil || isNil(o.FrameType) {
		var ret int32
		return ret
	}
	return *o.FrameType
}

// GetFrameTypeOk returns a tuple with the FrameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetFrameTypeOk() (*int32, bool) {
	if o == nil || isNil(o.FrameType) {
    return nil, false
	}
	return o.FrameType, true
}

// HasFrameType returns a boolean if a field has been set.
func (o *DetailedGear) HasFrameType() bool {
	if o != nil && !isNil(o.FrameType) {
		return true
	}

	return false
}

// SetFrameType gets a reference to the given int32 and assigns it to the FrameType field.
func (o *DetailedGear) SetFrameType(v int32) {
	o.FrameType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DetailedGear) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedGear) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DetailedGear) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DetailedGear) SetDescription(v string) {
	o.Description = &v
}

func (o DetailedGear) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ResourceState) {
		toSerialize["resource_state"] = o.ResourceState
	}
	if !isNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !isNil(o.BrandName) {
		toSerialize["brand_name"] = o.BrandName
	}
	if !isNil(o.ModelName) {
		toSerialize["model_name"] = o.ModelName
	}
	if !isNil(o.FrameType) {
		toSerialize["frame_type"] = o.FrameType
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableDetailedGear struct {
	value *DetailedGear
	isSet bool
}

func (v NullableDetailedGear) Get() *DetailedGear {
	return v.value
}

func (v *NullableDetailedGear) Set(val *DetailedGear) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedGear) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedGear) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedGear(val *DetailedGear) *NullableDetailedGear {
	return &NullableDetailedGear{value: val, isSet: true}
}

func (v NullableDetailedGear) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedGear) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


