package mongo

import (
	"fmt"
	"time"

	"github.com/Goldwin/ies-pik-cms/pkg/attendance/dto"
	"github.com/Goldwin/ies-pik-cms/pkg/attendance/entities"
	"github.com/samber/lo"
)

const (
	AttendanceCollection              = "attendances"
	AttendanceSummaryCollection       = "attendance_summaries"
	EventScheduleCollection           = "event_schedules"
	EventCollection                   = "events"
	PersonCollection                  = "persons"
	HouseholdCollection               = "households"
	PersonHouseholdCollection         = "person_households"
	PersonAttendanceSummaryCollection = "person_attendance_summary"
	LabelsCollection                  = "labels"
)

type PersonAttendanceSummaryModel struct {
	ID                   string          `bson:"_id"`
	PersonID             string          `bson:"personId"`
	ScheduleID           string          `bson:"scheduleId"`
	FirstEventAttendance AttendanceModel `bson:"firstAttendance"`
	LastEventAttendance  AttendanceModel `bson:"lastAttendance"`
}

func (m *PersonAttendanceSummaryModel) ToEntity() *entities.PersonAttendanceSummary {
	return &entities.PersonAttendanceSummary{
		PersonID:             m.PersonID,
		ScheduleID:           m.ScheduleID,
		FirstEventAttendance: m.FirstEventAttendance.ToEntity(),
		LastEventAttendance:  m.LastEventAttendance.ToEntity(),
	}
}

func toPersonAttendanceSummaryModel(e *entities.PersonAttendanceSummary) PersonAttendanceSummaryModel {
	return PersonAttendanceSummaryModel{
		ID:                   e.ID(),
		PersonID:             e.PersonID,
		ScheduleID:           e.ScheduleID,
		FirstEventAttendance: toAttendanceModel(e.FirstEventAttendance),
		LastEventAttendance:  toAttendanceModel(e.LastEventAttendance),
	}
}

type EventScheduleModel struct {
	ID             string                       `bson:"_id"`
	Name           string                       `bson:"name"`
	TimezoneOffset int                          `bson:"timezoneOffset"`
	Type           string                       `bson:"type"`
	Activities     []EventScheduleActivityModel `bson:"activities"`
	Date           time.Time                    `bson:"date"`
	Days           []time.Weekday               `bson:"days"`
	StartDate      time.Time                    `bson:"startDate"`
	EndDate        time.Time                    `bson:"endDate"`
	StartTime      string                       `bson:"startTime"`
	EndTime        string                       `bson:"endTime"`
}

func toEventScheduleModel(e *entities.EventSchedule) EventScheduleModel {
	return EventScheduleModel{
		ID:             e.ID,
		Name:           e.Name,
		TimezoneOffset: e.TimezoneOffset,
		Type:           string(e.Type),
		Activities: lo.Map(e.Activities, func(e *entities.EventScheduleActivity, _ int) EventScheduleActivityModel {
			return toEventScheduleActivityModel(e)
		}),
		Date:      e.Date,
		Days:      e.Days,
		StartDate: e.StartDate,
		EndDate:   e.EndDate,
		StartTime: e.StartTime.String(),
		EndTime:   e.EndTime.String(),
	}
}

func (e *EventScheduleModel) ToEventSchedule() *entities.EventSchedule {
	var startTime, endTime entities.HourMinute
	startTime.SetFromStringOrZero(e.StartTime)
	endTime.SetFromStringOrMaxValue(e.EndTime)
	return &entities.EventSchedule{
		ID:             e.ID,
		Name:           e.Name,
		TimezoneOffset: e.TimezoneOffset,
		Type:           entities.EventScheduleType(e.Type),
		Activities: lo.Map(e.Activities, func(e EventScheduleActivityModel, _ int) *entities.EventScheduleActivity {
			return e.ToEventScheduleActivity()
		}),
		OneTimeEventSchedule: entities.OneTimeEventSchedule{
			Date: e.Date,
		},
		WeeklyEventSchedule: entities.WeeklyEventSchedule{
			Days: e.Days,
		},
		DailyEventSchedule: entities.DailyEventSchedule{
			StartDate: e.StartDate,
			EndDate:   e.EndDate,
		},
		StartTime: startTime,
		EndTime:   endTime,
	}
}

type EventScheduleActivityModel struct {
	ID     string               `bson:"_id"`
	Name   string               `bson:"name"`
	Hour   int                  `bson:"hour"`
	Minute int                  `bson:"minute"`
	Labels []ActivityLabelModel `bson:"labels"`
}

func toEventScheduleActivityModel(e *entities.EventScheduleActivity) EventScheduleActivityModel {
	return EventScheduleActivityModel{
		ID:     e.ID,
		Name:   e.Name,
		Hour:   e.Hour,
		Minute: e.Minute,
		Labels: lo.Map(e.Labels, func(e *entities.ActivityLabel, _ int) ActivityLabelModel {
			return FromActivityLabelEntity(e)
		}),
	}
}

func (e *EventScheduleActivityModel) ToEventScheduleActivity() *entities.EventScheduleActivity {
	return &entities.EventScheduleActivity{
		ID:     e.ID,
		Name:   e.Name,
		Hour:   e.Hour,
		Minute: e.Minute,
		Labels: lo.Map(e.Labels, func(model ActivityLabelModel, _ int) *entities.ActivityLabel {
			return model.ToEntity()
		}),
	}
}

type EventModel struct {
	ID              string               `bson:"_id"`
	Name            string               `bson:"name"`
	ScheduleID      string               `bson:"scheduleId"`
	EventActivities []EventActivityModel `bson:"eventActivities"`
	StartDate       time.Time            `bson:"startDate"`
	EndDate         time.Time            `bson:"endDate"`
}

func (e *EventModel) ToEvent() *entities.Event {
	return &entities.Event{
		ID:         e.ID,
		ScheduleID: e.ScheduleID,
		Name:       e.Name,
		EventActivities: lo.Map(e.EventActivities, func(e EventActivityModel, _ int) *entities.EventActivity {
			return e.ToEventActivity()
		}),
		StartDate: e.StartDate,
		EndDate:   e.EndDate,
	}
}

func (e *EventModel) ToDTO() dto.EventDTO {
	return dto.EventDTO{
		ID:         e.ID,
		ScheduleID: e.ScheduleID,
		Name:       e.Name,
		Activities: lo.Map(e.EventActivities, func(e EventActivityModel, _ int) dto.EventActivityDTO {
			return dto.EventActivityDTO{
				ID:   e.ID,
				Name: e.Name,
				Time: e.Time,
			}
		}),
		StartDate: e.StartDate,
		EndDate:   e.EndDate,
	}
}

func toEventModel(e *entities.Event) EventModel {
	return EventModel{
		ID:         e.ID,
		Name:       e.Name,
		ScheduleID: e.ScheduleID,
		EventActivities: lo.Map(e.EventActivities, func(e *entities.EventActivity, _ int) EventActivityModel {
			return toEventActivityModel(e)
		}),
		StartDate: e.StartDate,
		EndDate:   e.EndDate,
	}
}

type EventActivityModel struct {
	ID     string               `bson:"_id"`
	Name   string               `bson:"name"`
	Time   time.Time            `bson:"time"`
	Labels []ActivityLabelModel `bson:"labels"`
}

func (e *EventActivityModel) ToEventActivity() *entities.EventActivity {
	return &entities.EventActivity{
		ID:   e.ID,
		Name: e.Name,
		Time: e.Time,
		Labels: lo.Map(e.Labels, func(model ActivityLabelModel, _ int) *entities.ActivityLabel {
			return model.ToEntity()
		}),
	}
}

func toEventActivityModel(e *entities.EventActivity) EventActivityModel {
	return EventActivityModel{
		ID:   e.ID,
		Name: e.Name,
		Time: e.Time,
		Labels: lo.Map(e.Labels, func(e *entities.ActivityLabel, _ int) ActivityLabelModel {
			return FromActivityLabelEntity(e)
		}),
	}
}

type PersonModel struct {
	ID                string `bson:"_id"`
	PersonID          string `bson:"personId"`
	FirstName         string `bson:"firstName"`
	MiddleName        string `bson:"middleName"`
	LastName          string `bson:"lastName"`
	ProfilePictureUrl string `bson:"profilePictureUrl"`
	Birthday          string `bson:"birthday"`
}

func (p *PersonModel) ToEntity() *entities.Person {
	id := p.ID
	if id == "" {
		id = p.PersonID
	}
	return &entities.Person{
		PersonID:          id,
		FirstName:         p.FirstName,
		MiddleName:        p.MiddleName,
		LastName:          p.LastName,
		ProfilePictureUrl: p.ProfilePictureUrl,
	}
}

func (p *PersonModel) ToDTO() dto.PersonDTO {
	id := p.ID
	if id == "" {
		id = p.PersonID
	}
	var year, month, day int
	fmt.Sscanf(p.Birthday, "%d-%d-%d", &year, &month, &day)

	birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return dto.PersonDTO{
		ID:                id,
		FirstName:         p.FirstName,
		MiddleName:        p.MiddleName,
		LastName:          p.LastName,
		ProfilePictureUrl: p.ProfilePictureUrl,
		Age:               int(time.Now().Sub(birthday).Hours() / 24 / 365),
	}
}

type HouseholdModel struct {
	ID               string        `bson:"_id"`
	Name             string        `bson:"name"`
	HouseholdHead    PersonModel   `bson:"householdHead"`
	PictureUrl       string        `bson:"pictureUrl"`
	HouseholdMembers []PersonModel `bson:"householdMembers"`
}

type AttendanceModel struct {
	ID            string             `bson:"_id"`
	Event         EventModel         `bson:"event"`
	EventActivity EventActivityModel `bson:"eventActivity"`

	Attendee    PersonModel `bson:"attendee"`
	CheckedInBy PersonModel `bson:"checkedInBy"`

	SecurityCode   string    `bson:"securityCode"`
	SecurityNumber int       `bson:"securityNumber"`
	CheckinTime    time.Time `bson:"checkinTime"`

	Type      string `bson:"type"`
	FirstTime bool   `bson:"firstTime"`
}

func (a *AttendanceModel) ToEntity() *entities.Attendance {
	return &entities.Attendance{
		ID:             a.ID,
		Event:          a.Event.ToEvent(),
		EventActivity:  a.EventActivity.ToEventActivity(),
		Attendee:       a.Attendee.ToEntity(),
		CheckedInBy:    a.CheckedInBy.ToEntity(),
		SecurityCode:   a.SecurityCode,
		SecurityNumber: a.SecurityNumber,
		CheckinTime:    a.CheckinTime,
		Type:           entities.AttendanceType(a.Type),
		FirstTime:      a.FirstTime,
	}
}

func toAttendanceModel(e *entities.Attendance) AttendanceModel {
	return AttendanceModel{
		ID:            e.ID,
		Event:         toEventModel(e.Event),
		EventActivity: toEventActivityModel(e.EventActivity),
		Attendee: PersonModel{
			ID:                e.Attendee.PersonID,
			PersonID:          e.Attendee.PersonID,
			FirstName:         e.Attendee.FirstName,
			MiddleName:        e.Attendee.MiddleName,
			LastName:          e.Attendee.LastName,
			ProfilePictureUrl: e.Attendee.ProfilePictureUrl,
		},
		CheckedInBy: PersonModel{
			ID:                e.CheckedInBy.PersonID,
			PersonID:          e.CheckedInBy.PersonID,
			FirstName:         e.CheckedInBy.FirstName,
			MiddleName:        e.CheckedInBy.MiddleName,
			LastName:          e.CheckedInBy.LastName,
			ProfilePictureUrl: e.CheckedInBy.ProfilePictureUrl,
		},
		SecurityCode:   e.SecurityCode,
		SecurityNumber: e.SecurityNumber,
		CheckinTime:    e.CheckinTime,
		Type:           string(e.Type),
		FirstTime:      e.FirstTime,
	}
}

func (e *AttendanceModel) ToAttendance() *entities.Attendance {
	return &entities.Attendance{
		ID:             e.ID,
		Event:          e.Event.ToEvent(),
		EventActivity:  e.EventActivity.ToEventActivity(),
		Attendee:       &entities.Person{PersonID: e.Attendee.PersonID, FirstName: e.Attendee.FirstName, MiddleName: e.Attendee.MiddleName, LastName: e.Attendee.LastName, ProfilePictureUrl: e.Attendee.ProfilePictureUrl},
		SecurityCode:   e.SecurityCode,
		SecurityNumber: e.SecurityNumber,
		CheckinTime:    e.CheckinTime,
		Type:           entities.AttendanceType(e.Type),
	}
}

type ActivityAttendanceSummaryModel struct {
	ID          string         `bson:"_id"`
	Name        string         `bson:"name"`
	Total       int            `bson:"total"`
	TotalByType map[string]int `bson:"totalByType"`
}

func (e *ActivityAttendanceSummaryModel) ToDTO() dto.ActivityAttendanceSummaryDTO {
	return dto.ActivityAttendanceSummaryDTO{
		Name:        e.Name,
		Total:       e.Total,
		TotalByType: e.TotalByType,
	}
}

type EventAttendanceSummaryModel struct {
	ID              string    `bson:"_id"`
	ScheduleID      string    `bson:"scheduleId"`
	Date            time.Time `bson:"date"`
	TotalCheckedIn  int       `bson:"totalCheckedIn"`
	TotalCheckedOut int       `bson:"totalCheckedOut"`
	TotalFirstTimer int       `bson:"totalFirstTimer"`
	Total           int       `bson:"total"`

	TotalByType        map[string]int                   `bson:"totalByType"`
	AcitivitiesSummary []ActivityAttendanceSummaryModel `bson:"activitiesSummary"`
	LastUpdated        time.Time                        `bson:"lastUpdated"`
	NextUpdate         time.Time                        `bson:"nextUpdate"`
}

func (e *EventAttendanceSummaryModel) ToDTO() dto.EventAttendanceSummaryDTO {
	return dto.EventAttendanceSummaryDTO{
		TotalCheckedIn:  e.TotalCheckedIn,
		TotalCheckedOut: e.TotalCheckedOut,
		TotalFirstTimer: e.TotalFirstTimer,
		Total:           e.Total,
		TotalByType:     e.TotalByType,
		AcitivitiesSummary: lo.Map(e.AcitivitiesSummary, func(ee ActivityAttendanceSummaryModel, _ int) dto.ActivityAttendanceSummaryDTO {
			return ee.ToDTO()
		}),
		Date: e.Date,
		ID:   e.ID,
	}
}

type ActivityLabelModel struct {
	LabelID         string   `bson:"labelId"`
	LabelName       string   `bson:"labelName"`
	Type            string   `bson:"type"`
	AttendanceTypes []string `bson:"attendanceType"`
	Quantity        int      `bson:"quantity"`
}

func (model ActivityLabelModel) ToEntity() *entities.ActivityLabel {
	return &entities.ActivityLabel{
		LabelID:   model.LabelID,
		LabelName: model.LabelName,
		Type:      entities.LabelType(model.Type),
		AttendanceTypes: lo.Map(model.AttendanceTypes, func(attendanceType string, _ int) entities.AttendanceType {
			return entities.AttendanceType(attendanceType)
		}),
		Quantity: model.Quantity,
	}
}

func FromActivityLabelEntity(e *entities.ActivityLabel) ActivityLabelModel {
	return ActivityLabelModel{
		LabelID:   e.LabelID,
		LabelName: e.LabelName,
		Type:      string(e.Type),
		AttendanceTypes: lo.Map(e.AttendanceTypes, func(attendanceType entities.AttendanceType, _ int) string {
			return string(attendanceType)
		}),
		Quantity: e.Quantity,
	}
}

func (model ActivityLabelModel) ToDTO() dto.ActivityLabelDTO {
	return dto.ActivityLabelDTO{
		LabelID:         model.LabelID,
		LabelName:       model.LabelName,
		Type:            model.Type,
		AttendanceTypes: model.AttendanceTypes,
		Quantity:        model.Quantity,
	}
}

type LabelModel struct {
	ID          string           `bson:"_id"`
	Name        string           `bson:"name"`
	Type        string           `bson:"type"`
	Orientation string           `bson:"orientation"`
	PaperSize   []float64        `bson:"paperSize"`
	Margin      []float64        `bson:"margin"`
	Objects     []map[string]any `bson:"objects"`
}

func (model LabelModel) ToDTO() dto.LabelDTO {
	return dto.LabelDTO{
		ID:          model.ID,
		Name:        model.Name,
		Type:        model.Type,
		Orientation: model.Orientation,
		PaperSize:   model.PaperSize,
		Margin:      model.Margin,
		Objects:     model.Objects,
	}
}

func (model LabelModel) ToEntity() *entities.Label {
	return &entities.Label{
		ID:          model.ID,
		Name:        model.Name,
		Type:        entities.LabelType(model.Type),
		Orientation: model.Orientation,
		PaperSize:   model.PaperSize,
		Margin:      model.Margin,
		Objects:     model.Objects,
	}
}
