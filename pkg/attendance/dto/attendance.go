package dto

import (
	"time"

	"github.com/Goldwin/ies-pik-cms/pkg/attendance/entities"
)

type AttendeeDTO struct {
	PersonID          string `json:"personId"`
	FirstName         string `json:"firstName"`
	MiddleName        string `json:"middleName"`
	LastName          string `json:"lastName"`
	ProfilePictureURL string `json:"profilePictureUrl"`
}

type EventAttendanceDTO struct {
	ID             string           `json:"id"`
	Event          EventDTO         `json:"event"`
	Activity       EventActivityDTO `json:"activity"`
	Person         AttendeeDTO      `json:"person"`
	SecurityCode   string           `json:"securityCode"`
	SecurityNumber int              `json:"securityNumber"`
	CheckinTime    time.Time        `json:"checkinTime"`
	AttendanceType string           `json:"attendanceType"`
}

func FromAttendanceEntities(result *entities.Attendance) EventAttendanceDTO {
	return EventAttendanceDTO{
		ID:             result.ID,
		Event:          FromEventEntities(result.Event),
		Activity:       EventActivityDTO{ID: result.EventActivity.ID, Name: result.EventActivity.Name, Time: result.EventActivity.Time},
		Person:         AttendeeDTO{PersonID: result.PersonID, FirstName: result.FirstName, MiddleName: result.MiddleName, LastName: result.LastName, ProfilePictureURL: result.ProfilePictureUrl},
		SecurityCode:   result.SecurityCode,
		SecurityNumber: result.SecurityNumber,
		CheckinTime:    result.CheckinTime,
		AttendanceType: string(result.Type),
	}
}
