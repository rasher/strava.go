/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package strava
import (
	"time"
)
// DetailedAthlete struct for DetailedAthlete
type DetailedAthlete struct {
	// The unique identifier of the athlete
	Id int32 `json:"id,omitempty"`
	// Resource state, indicates level of detail. Possible values: 1 -> \"meta\", 2 -> \"summary\", 3 -> \"detail\"
	ResourceState int32 `json:"resource_state,omitempty"`
	// The athlete's first name.
	Firstname string `json:"firstname,omitempty"`
	// The athlete's last name.
	Lastname string `json:"lastname,omitempty"`
	// URL to a 62x62 pixel profile picture.
	ProfileMedium string `json:"profile_medium,omitempty"`
	// URL to a 124x124 pixel profile picture.
	Profile string `json:"profile,omitempty"`
	// The athlete's city.
	City string `json:"city,omitempty"`
	// The athlete's state or geographical region.
	State string `json:"state,omitempty"`
	// The athlete's country.
	Country string `json:"country,omitempty"`
	// The athlete's sex.
	Sex string `json:"sex,omitempty"`
	// Deprecated.  Use summit field instead. Whether the athlete has any Summit subscription.
	Premium bool `json:"premium,omitempty"`
	// Whether the athlete has any Summit subscription.
	Summit bool `json:"summit,omitempty"`
	// The time at which the athlete was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The time at which the athlete was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The athlete's follower count.
	FollowerCount int32 `json:"follower_count,omitempty"`
	// The athlete's friend count.
	FriendCount int32 `json:"friend_count,omitempty"`
	// The athlete's preferred unit system.
	MeasurementPreference string `json:"measurement_preference,omitempty"`
	// The athlete's FTP (Functional Threshold Power).
	Ftp int32 `json:"ftp,omitempty"`
	// The athlete's weight.
	Weight float32 `json:"weight,omitempty"`
	// The athlete's clubs.
	Clubs []SummaryClub `json:"clubs,omitempty"`
	// The athlete's bikes.
	Bikes []SummaryGear `json:"bikes,omitempty"`
	// The athlete's shoes.
	Shoes []SummaryGear `json:"shoes,omitempty"`
}
