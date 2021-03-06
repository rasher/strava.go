# Lap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique identifier of this lap | [optional] 
**Activity** | [**MetaActivity**](MetaActivity.md) |  | [optional] 
**Athlete** | [**MetaAthlete**](MetaAthlete.md) |  | [optional] 
**AverageCadence** | **float32** | The lap&#39;s average cadence | [optional] 
**AverageSpeed** | **float32** | The lap&#39;s average speed | [optional] 
**Distance** | **float32** | The lap&#39;s distance, in meters | [optional] 
**ElapsedTime** | **int32** | The lap&#39;s elapsed time, in seconds | [optional] 
**StartIndex** | **int32** | The start index of this effort in its activity&#39;s stream | [optional] 
**EndIndex** | **int32** | The end index of this effort in its activity&#39;s stream | [optional] 
**LapIndex** | **int32** | The index of this lap in the activity it belongs to | [optional] 
**MaxSpeed** | **float32** | The maximum speed of this lat, in meters per second | [optional] 
**MovingTime** | **int32** | The lap&#39;s moving time, in seconds | [optional] 
**Name** | **string** | The name of the lap | [optional] 
**PaceZone** | **int32** | The athlete&#39;s pace zone during this lap | [optional] 
**Split** | **int32** |  | [optional] 
**StartDate** | [**time.Time**](time.Time.md) | The time at which the lap was started. | [optional] 
**StartDateLocal** | [**time.Time**](time.Time.md) | The time at which the lap was started in the local timezone. | [optional] 
**TotalElevationGain** | **float32** | The elevation gain of this lap, in meters | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


