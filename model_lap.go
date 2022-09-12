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

// Lap struct for Lap
type Lap struct {
	// The unique identifier of this lap
	Id *int64 `json:"id,omitempty"`
	Activity *MetaActivity `json:"activity,omitempty"`
	Athlete *MetaAthlete `json:"athlete,omitempty"`
	// The lap's average cadence
	AverageCadence *float32 `json:"average_cadence,omitempty"`
	// The lap's average speed
	AverageSpeed *float32 `json:"average_speed,omitempty"`
	// The lap's distance, in meters
	Distance *float32 `json:"distance,omitempty"`
	// The lap's elapsed time, in seconds
	ElapsedTime *int32 `json:"elapsed_time,omitempty"`
	// The start index of this effort in its activity's stream
	StartIndex *int32 `json:"start_index,omitempty"`
	// The end index of this effort in its activity's stream
	EndIndex *int32 `json:"end_index,omitempty"`
	// The index of this lap in the activity it belongs to
	LapIndex *int32 `json:"lap_index,omitempty"`
	// The maximum speed of this lat, in meters per second
	MaxSpeed *float32 `json:"max_speed,omitempty"`
	// The lap's moving time, in seconds
	MovingTime *int32 `json:"moving_time,omitempty"`
	// The name of the lap
	Name *string `json:"name,omitempty"`
	// The athlete's pace zone during this lap
	PaceZone *int32 `json:"pace_zone,omitempty"`
	Split *int32 `json:"split,omitempty"`
	// The time at which the lap was started.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The time at which the lap was started in the local timezone.
	StartDateLocal *time.Time `json:"start_date_local,omitempty"`
	// The elevation gain of this lap, in meters
	TotalElevationGain *float32 `json:"total_elevation_gain,omitempty"`
}

// NewLap instantiates a new Lap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLap() *Lap {
	this := Lap{}
	return &this
}

// NewLapWithDefaults instantiates a new Lap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLapWithDefaults() *Lap {
	this := Lap{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Lap) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Lap) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Lap) SetId(v int64) {
	o.Id = &v
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *Lap) GetActivity() MetaActivity {
	if o == nil || o.Activity == nil {
		var ret MetaActivity
		return ret
	}
	return *o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetActivityOk() (*MetaActivity, bool) {
	if o == nil || o.Activity == nil {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *Lap) HasActivity() bool {
	if o != nil && o.Activity != nil {
		return true
	}

	return false
}

// SetActivity gets a reference to the given MetaActivity and assigns it to the Activity field.
func (o *Lap) SetActivity(v MetaActivity) {
	o.Activity = &v
}

// GetAthlete returns the Athlete field value if set, zero value otherwise.
func (o *Lap) GetAthlete() MetaAthlete {
	if o == nil || o.Athlete == nil {
		var ret MetaAthlete
		return ret
	}
	return *o.Athlete
}

// GetAthleteOk returns a tuple with the Athlete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetAthleteOk() (*MetaAthlete, bool) {
	if o == nil || o.Athlete == nil {
		return nil, false
	}
	return o.Athlete, true
}

// HasAthlete returns a boolean if a field has been set.
func (o *Lap) HasAthlete() bool {
	if o != nil && o.Athlete != nil {
		return true
	}

	return false
}

// SetAthlete gets a reference to the given MetaAthlete and assigns it to the Athlete field.
func (o *Lap) SetAthlete(v MetaAthlete) {
	o.Athlete = &v
}

// GetAverageCadence returns the AverageCadence field value if set, zero value otherwise.
func (o *Lap) GetAverageCadence() float32 {
	if o == nil || o.AverageCadence == nil {
		var ret float32
		return ret
	}
	return *o.AverageCadence
}

// GetAverageCadenceOk returns a tuple with the AverageCadence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetAverageCadenceOk() (*float32, bool) {
	if o == nil || o.AverageCadence == nil {
		return nil, false
	}
	return o.AverageCadence, true
}

// HasAverageCadence returns a boolean if a field has been set.
func (o *Lap) HasAverageCadence() bool {
	if o != nil && o.AverageCadence != nil {
		return true
	}

	return false
}

// SetAverageCadence gets a reference to the given float32 and assigns it to the AverageCadence field.
func (o *Lap) SetAverageCadence(v float32) {
	o.AverageCadence = &v
}

// GetAverageSpeed returns the AverageSpeed field value if set, zero value otherwise.
func (o *Lap) GetAverageSpeed() float32 {
	if o == nil || o.AverageSpeed == nil {
		var ret float32
		return ret
	}
	return *o.AverageSpeed
}

// GetAverageSpeedOk returns a tuple with the AverageSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetAverageSpeedOk() (*float32, bool) {
	if o == nil || o.AverageSpeed == nil {
		return nil, false
	}
	return o.AverageSpeed, true
}

// HasAverageSpeed returns a boolean if a field has been set.
func (o *Lap) HasAverageSpeed() bool {
	if o != nil && o.AverageSpeed != nil {
		return true
	}

	return false
}

// SetAverageSpeed gets a reference to the given float32 and assigns it to the AverageSpeed field.
func (o *Lap) SetAverageSpeed(v float32) {
	o.AverageSpeed = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *Lap) GetDistance() float32 {
	if o == nil || o.Distance == nil {
		var ret float32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetDistanceOk() (*float32, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *Lap) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float32 and assigns it to the Distance field.
func (o *Lap) SetDistance(v float32) {
	o.Distance = &v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *Lap) GetElapsedTime() int32 {
	if o == nil || o.ElapsedTime == nil {
		var ret int32
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetElapsedTimeOk() (*int32, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *Lap) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int32 and assigns it to the ElapsedTime field.
func (o *Lap) SetElapsedTime(v int32) {
	o.ElapsedTime = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *Lap) GetStartIndex() int32 {
	if o == nil || o.StartIndex == nil {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetStartIndexOk() (*int32, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *Lap) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *Lap) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetEndIndex returns the EndIndex field value if set, zero value otherwise.
func (o *Lap) GetEndIndex() int32 {
	if o == nil || o.EndIndex == nil {
		var ret int32
		return ret
	}
	return *o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetEndIndexOk() (*int32, bool) {
	if o == nil || o.EndIndex == nil {
		return nil, false
	}
	return o.EndIndex, true
}

// HasEndIndex returns a boolean if a field has been set.
func (o *Lap) HasEndIndex() bool {
	if o != nil && o.EndIndex != nil {
		return true
	}

	return false
}

// SetEndIndex gets a reference to the given int32 and assigns it to the EndIndex field.
func (o *Lap) SetEndIndex(v int32) {
	o.EndIndex = &v
}

// GetLapIndex returns the LapIndex field value if set, zero value otherwise.
func (o *Lap) GetLapIndex() int32 {
	if o == nil || o.LapIndex == nil {
		var ret int32
		return ret
	}
	return *o.LapIndex
}

// GetLapIndexOk returns a tuple with the LapIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetLapIndexOk() (*int32, bool) {
	if o == nil || o.LapIndex == nil {
		return nil, false
	}
	return o.LapIndex, true
}

// HasLapIndex returns a boolean if a field has been set.
func (o *Lap) HasLapIndex() bool {
	if o != nil && o.LapIndex != nil {
		return true
	}

	return false
}

// SetLapIndex gets a reference to the given int32 and assigns it to the LapIndex field.
func (o *Lap) SetLapIndex(v int32) {
	o.LapIndex = &v
}

// GetMaxSpeed returns the MaxSpeed field value if set, zero value otherwise.
func (o *Lap) GetMaxSpeed() float32 {
	if o == nil || o.MaxSpeed == nil {
		var ret float32
		return ret
	}
	return *o.MaxSpeed
}

// GetMaxSpeedOk returns a tuple with the MaxSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetMaxSpeedOk() (*float32, bool) {
	if o == nil || o.MaxSpeed == nil {
		return nil, false
	}
	return o.MaxSpeed, true
}

// HasMaxSpeed returns a boolean if a field has been set.
func (o *Lap) HasMaxSpeed() bool {
	if o != nil && o.MaxSpeed != nil {
		return true
	}

	return false
}

// SetMaxSpeed gets a reference to the given float32 and assigns it to the MaxSpeed field.
func (o *Lap) SetMaxSpeed(v float32) {
	o.MaxSpeed = &v
}

// GetMovingTime returns the MovingTime field value if set, zero value otherwise.
func (o *Lap) GetMovingTime() int32 {
	if o == nil || o.MovingTime == nil {
		var ret int32
		return ret
	}
	return *o.MovingTime
}

// GetMovingTimeOk returns a tuple with the MovingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetMovingTimeOk() (*int32, bool) {
	if o == nil || o.MovingTime == nil {
		return nil, false
	}
	return o.MovingTime, true
}

// HasMovingTime returns a boolean if a field has been set.
func (o *Lap) HasMovingTime() bool {
	if o != nil && o.MovingTime != nil {
		return true
	}

	return false
}

// SetMovingTime gets a reference to the given int32 and assigns it to the MovingTime field.
func (o *Lap) SetMovingTime(v int32) {
	o.MovingTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Lap) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Lap) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Lap) SetName(v string) {
	o.Name = &v
}

// GetPaceZone returns the PaceZone field value if set, zero value otherwise.
func (o *Lap) GetPaceZone() int32 {
	if o == nil || o.PaceZone == nil {
		var ret int32
		return ret
	}
	return *o.PaceZone
}

// GetPaceZoneOk returns a tuple with the PaceZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetPaceZoneOk() (*int32, bool) {
	if o == nil || o.PaceZone == nil {
		return nil, false
	}
	return o.PaceZone, true
}

// HasPaceZone returns a boolean if a field has been set.
func (o *Lap) HasPaceZone() bool {
	if o != nil && o.PaceZone != nil {
		return true
	}

	return false
}

// SetPaceZone gets a reference to the given int32 and assigns it to the PaceZone field.
func (o *Lap) SetPaceZone(v int32) {
	o.PaceZone = &v
}

// GetSplit returns the Split field value if set, zero value otherwise.
func (o *Lap) GetSplit() int32 {
	if o == nil || o.Split == nil {
		var ret int32
		return ret
	}
	return *o.Split
}

// GetSplitOk returns a tuple with the Split field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetSplitOk() (*int32, bool) {
	if o == nil || o.Split == nil {
		return nil, false
	}
	return o.Split, true
}

// HasSplit returns a boolean if a field has been set.
func (o *Lap) HasSplit() bool {
	if o != nil && o.Split != nil {
		return true
	}

	return false
}

// SetSplit gets a reference to the given int32 and assigns it to the Split field.
func (o *Lap) SetSplit(v int32) {
	o.Split = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Lap) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Lap) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Lap) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStartDateLocal returns the StartDateLocal field value if set, zero value otherwise.
func (o *Lap) GetStartDateLocal() time.Time {
	if o == nil || o.StartDateLocal == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateLocal
}

// GetStartDateLocalOk returns a tuple with the StartDateLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetStartDateLocalOk() (*time.Time, bool) {
	if o == nil || o.StartDateLocal == nil {
		return nil, false
	}
	return o.StartDateLocal, true
}

// HasStartDateLocal returns a boolean if a field has been set.
func (o *Lap) HasStartDateLocal() bool {
	if o != nil && o.StartDateLocal != nil {
		return true
	}

	return false
}

// SetStartDateLocal gets a reference to the given time.Time and assigns it to the StartDateLocal field.
func (o *Lap) SetStartDateLocal(v time.Time) {
	o.StartDateLocal = &v
}

// GetTotalElevationGain returns the TotalElevationGain field value if set, zero value otherwise.
func (o *Lap) GetTotalElevationGain() float32 {
	if o == nil || o.TotalElevationGain == nil {
		var ret float32
		return ret
	}
	return *o.TotalElevationGain
}

// GetTotalElevationGainOk returns a tuple with the TotalElevationGain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lap) GetTotalElevationGainOk() (*float32, bool) {
	if o == nil || o.TotalElevationGain == nil {
		return nil, false
	}
	return o.TotalElevationGain, true
}

// HasTotalElevationGain returns a boolean if a field has been set.
func (o *Lap) HasTotalElevationGain() bool {
	if o != nil && o.TotalElevationGain != nil {
		return true
	}

	return false
}

// SetTotalElevationGain gets a reference to the given float32 and assigns it to the TotalElevationGain field.
func (o *Lap) SetTotalElevationGain(v float32) {
	o.TotalElevationGain = &v
}

func (o Lap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Activity != nil {
		toSerialize["activity"] = o.Activity
	}
	if o.Athlete != nil {
		toSerialize["athlete"] = o.Athlete
	}
	if o.AverageCadence != nil {
		toSerialize["average_cadence"] = o.AverageCadence
	}
	if o.AverageSpeed != nil {
		toSerialize["average_speed"] = o.AverageSpeed
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.ElapsedTime != nil {
		toSerialize["elapsed_time"] = o.ElapsedTime
	}
	if o.StartIndex != nil {
		toSerialize["start_index"] = o.StartIndex
	}
	if o.EndIndex != nil {
		toSerialize["end_index"] = o.EndIndex
	}
	if o.LapIndex != nil {
		toSerialize["lap_index"] = o.LapIndex
	}
	if o.MaxSpeed != nil {
		toSerialize["max_speed"] = o.MaxSpeed
	}
	if o.MovingTime != nil {
		toSerialize["moving_time"] = o.MovingTime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PaceZone != nil {
		toSerialize["pace_zone"] = o.PaceZone
	}
	if o.Split != nil {
		toSerialize["split"] = o.Split
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.StartDateLocal != nil {
		toSerialize["start_date_local"] = o.StartDateLocal
	}
	if o.TotalElevationGain != nil {
		toSerialize["total_elevation_gain"] = o.TotalElevationGain
	}
	return json.Marshal(toSerialize)
}

type NullableLap struct {
	value *Lap
	isSet bool
}

func (v NullableLap) Get() *Lap {
	return v.value
}

func (v *NullableLap) Set(val *Lap) {
	v.value = val
	v.isSet = true
}

func (v NullableLap) IsSet() bool {
	return v.isSet
}

func (v *NullableLap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLap(val *Lap) *NullableLap {
	return &NullableLap{value: val, isSet: true}
}

func (v NullableLap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


