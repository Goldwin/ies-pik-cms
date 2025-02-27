package controllers

import (
	"github.com/Goldwin/ies-pik-cms/internal/middleware"
	"github.com/Goldwin/ies-pik-cms/pkg/attendance"
	"github.com/Goldwin/ies-pik-cms/pkg/attendance/dto"
	"github.com/Goldwin/ies-pik-cms/pkg/attendance/queries"
	"github.com/Goldwin/ies-pik-cms/pkg/common/out"
	"github.com/gin-gonic/gin"
)

type attendanceController struct {
	attendanceComponent attendance.AttendanceComponent
	middlewareComponent middleware.MiddlewareComponent
}

func InitializeAttendanceController(r *gin.Engine, middleware middleware.MiddlewareComponent, attendanceComponent attendance.AttendanceComponent) {
	attendanceController := attendanceController{
		attendanceComponent: attendanceComponent,
		middlewareComponent: middleware,
	}

	rg := r.Group("attendance")
	rg.POST("schedules", middleware.Auth("ATTENDANCE_CREATE_SCHEDULE"), attendanceController.createSchedule)
	rg.GET("schedules", middleware.Auth("ATTENDANCE_VIEW_SCHEDULE"), attendanceController.listEventSchedules)

	scheduleURL := "schedules/:scheduleID"
	rg.GET(scheduleURL, middleware.Auth("ATTENDANCE_VIEW_SCHEDULE"), attendanceController.getEventSchedule)
	rg.PUT(scheduleURL, middleware.Auth("ATTENDANCE_UPDATE_SCHEDULE"), attendanceController.updateEventSchedule)
	rg.DELETE(scheduleURL, middleware.Auth("ATTENDANCE_DELETE_SCHEDULE"), attendanceController.archiveEventSchedule)

	rg.POST("schedules/:scheduleID/create-next-event", middleware.Auth("ATTENDANCE_CREATE_EVENT"), attendanceController.createNextEvent)

	activitiesUrl := scheduleURL + "/activities"
	activityUrl := activitiesUrl + "/:activityID"
	rg.POST(activitiesUrl, middleware.Auth("ATTENDANCE_CREATE_SCHEDULE_ACTIVITY"), attendanceController.createEventScheduleActivity)
	rg.PUT(activityUrl, middleware.Auth("ATTENDANCE_UPDATE_SCHEDULE_ACTIVITY"), attendanceController.updateEventScheduleActivity)
	rg.DELETE(activityUrl, middleware.Auth("ATTENDANCE_DELETE_SCHEDULE_ACTIVITY"), attendanceController.removeEventScheduleActivity)

	rg.GET("schedules/:scheduleID/events", middleware.Auth("ATTENDANCE_VIEW_EVENTS"), attendanceController.listEventsBySchedule)
	rg.GET("schedules/:scheduleID/events/:eventID", middleware.Auth("ATTENDANCE_VIEW_EVENTS"), attendanceController.getEventBySchedule)

	rg.GET("schedules/:scheduleID/events/:eventID/attendees", middleware.Auth("ATTENDANCE_VIEW_EVENT_ATTENDANCE"), attendanceController.listEventAttendance)
	rg.POST("schedules/:scheduleID/events/:eventID/checkin", middleware.Auth("ATTENDANCE_CREATE_EVENT_ATTENDANCE"), attendanceController.checkIn)

	rg.GET("schedules/:scheduleID/events/:eventID/summary", middleware.Auth("ATTENDANCE_VIEW_EVENT_SUMMARY"), attendanceController.getSummary)

	rg.GET("schedules/:scheduleID/stats", middleware.Auth("ATTENDANCE_VIEW_SCHEDULE_STATS"), attendanceController.getEventScheduleStats)

	rg.GET("labels", middleware.Auth("ATTENDANCE_VIEW_LABELS"), attendanceController.listLabels)
}

func (a *attendanceController) listEventSchedules(c *gin.Context) {
	var query queries.ListEventScheduleFilter

	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := &outputDecorator[queries.ListEventScheduleResult]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		},
		successFunc: func(result queries.ListEventScheduleResult) {
			c.JSON(200, result)
		},
	}
	a.attendanceComponent.ListEventSchedules(c, query, output).Wait()
}

func (a *attendanceController) createSchedule(c *gin.Context) {
	var schedule dto.EventScheduleDTO
	output := &outputDecorator[dto.EventScheduleDTO]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result dto.EventScheduleDTO) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}
	err := c.ShouldBindJSON(&schedule)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	a.attendanceComponent.CreateEventSchedule(c, schedule, output).Wait()
}

func (a *attendanceController) getEventSchedule(c *gin.Context) {
	id := c.Param("scheduleID")
	output := &outputDecorator[queries.GetEventScheduleResult]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(err.Code, gin.H{
				"error": err,
			})
		},
		successFunc: func(result queries.GetEventScheduleResult) {
			c.JSON(200, result)
		},
	}
	a.attendanceComponent.GetEventSchedule(c, queries.GetEventScheduleFilter{
		ScheduleID: id,
	}, output).Wait()
}

func (a *attendanceController) updateEventSchedule(c *gin.Context) {
	var schedule dto.EventScheduleDTO
	output := &outputDecorator[dto.EventScheduleDTO]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result dto.EventScheduleDTO) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}
	err := c.ShouldBindJSON(&schedule)
	schedule.ID = c.Param("scheduleID")
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	a.attendanceComponent.UpdateEventSchedule(c, schedule, output).Wait()
}

func (a *attendanceController) createNextEvent(c *gin.Context) {
	scheduleID := c.Param("scheduleID")

	output := &outputDecorator[[]dto.EventDTO]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result []dto.EventDTO) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}

	a.attendanceComponent.CreateNextEvent(c, scheduleID, output).Wait()
}

func (a *attendanceController) archiveEventSchedule(c *gin.Context) {
	//TODO fill this
}

func (a *attendanceController) getEventScheduleStats(c *gin.Context) {
	var input queries.GetEventScheduleStatsFilter
	input.ScheduleID = c.Param("scheduleID")

	output := &outputDecorator[queries.GetEventScheduleStatsResult]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result queries.GetEventScheduleStatsResult) {
			c.JSON(200, result)
		},
	}

	a.attendanceComponent.GetEventScheduleStats(c, input, output).Wait()
}

func (a *attendanceController) listEventsBySchedule(c *gin.Context) {
	var input queries.ListEventByScheduleFilter
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	input.ScheduleID = c.Param("scheduleID")

	err = input.Validate()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := &outputDecorator[queries.ListEventByScheduleResult]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, err)
		},
		successFunc: func(result queries.ListEventByScheduleResult) {
			c.JSON(200, result)
		},
	}
	a.attendanceComponent.ListEventsBySchedule(c, input, output).Wait()
}

func (a *attendanceController) getEventBySchedule(c *gin.Context) {
	var input queries.GetEventFilter

	input.EventID = c.Param("eventID")
	input.ScheduleID = c.Param("scheduleID")

	output := &outputDecorator[queries.GetEventResult]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(err.Code, err)
		},
		successFunc: func(result queries.GetEventResult) {
			c.JSON(200, result)
		},
	}

	a.attendanceComponent.GetEvent(c, input, output).Wait()
}

func (a *attendanceController) listEventAttendance(c *gin.Context) {
	var input queries.ListEventAttendanceFilter
	err := c.ShouldBindQuery(&input)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	input.EventID = c.Param("eventID")

	output := &outputDecorator[queries.ListEventAttendanceResult]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, err)
		},
		successFunc: func(result queries.ListEventAttendanceResult) {
			c.JSON(200, result)
		},
	}

	a.attendanceComponent.ListEventAttendance(c, input, output).Wait()
}

func (a *attendanceController) createEventScheduleActivity(c *gin.Context) {
	var data dto.EventScheduleActivityDTO
	err := c.ShouldBindJSON(&data)
	data.ScheduleID = c.Param("scheduleID")

	if err != nil {
		c.JSON(400, gin.H{
			"error": gin.H{
				"code":    "400",
				"message": err.Error(),
			},
		})
		return
	}

	output := &outputDecorator[dto.EventScheduleDTO]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result dto.EventScheduleDTO) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}

	a.attendanceComponent.AddEventScheduleActivity(c, data, output).Wait()
}

func (a *attendanceController) updateEventScheduleActivity(c *gin.Context) {
	var data dto.EventScheduleActivityDTO
	err := c.ShouldBindJSON(&data)
	data.ScheduleID = c.Param("scheduleID")
	data.ID = c.Param("activityID")

	if err != nil {
		c.JSON(400, gin.H{
			"error": gin.H{
				"code":    "400",
				"message": err.Error(),
			},
		})
		return
	}

	output := &outputDecorator[dto.EventScheduleDTO]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result dto.EventScheduleDTO) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}

	a.attendanceComponent.UpdateEventScheduleActivity(c, data, output).Wait()
}

func (a *attendanceController) removeEventScheduleActivity(c *gin.Context) {
	var data dto.EventScheduleActivityDTO
	data.ScheduleID = c.Param("scheduleID")
	data.ID = c.Param("activityID")

	output := &outputDecorator[dto.EventScheduleDTO]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result dto.EventScheduleDTO) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}

	a.attendanceComponent.RemoveEventScheduleActivity(c, data, output).Wait()
}

func (a *attendanceController) checkIn(c *gin.Context) {
	var data dto.HouseholdCheckinDTO
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"error": gin.H{
				"code":    "400",
				"message": err.Error(),
			},
		})
	}

	output := &outputDecorator[[]dto.EventAttendanceDTO]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result []dto.EventAttendanceDTO) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}

	a.attendanceComponent.HouseholdCheckin(c, data, output).Wait()
}

func (a *attendanceController) getSummary(c *gin.Context) {
	var filter queries.GetEventAttendanceSummaryFilter
	filter.EventID = c.Param("eventID")

	output := &outputDecorator[queries.GetEventAttendanceSummaryResult]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, err)
		},
		successFunc: func(result queries.GetEventAttendanceSummaryResult) {
			c.JSON(200, result)
		},
	}
	a.attendanceComponent.GetEventAttendanceSummary(c, filter, output).Wait()
}

func (a *attendanceController) listLabels(c *gin.Context) {
	var filter queries.ListLabelsFilter

	err := c.ShouldBindQuery(&filter)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := &outputDecorator[queries.ListLabelsResult]{
		output: nil,
		errFunction: func(err out.AppErrorDetail) {
			c.JSON(400, err)
		},
		successFunc: func(result queries.ListLabelsResult) {
			c.JSON(200, result)
		},
	}
	a.attendanceComponent.ListLabels(c, filter, output).Wait()
}
