package ms

import (
	"github.com/fly365com/rabbit/message/activityLog"
	"ticket/entity"
)

type RadarMs struct {
}

func NewRadarMs() *RadarMs {
	return &RadarMs{}
}

func (rdrMs *RadarMs) LogCreateTicketActivity(ticket *entity.Ticket) {
	logActivity := activityLog.LogActivity{
		Identifier: ticket.Email,
		Action:     activityLog.LOG_ACTIVITY_CREATE_TICKET,
		OldData:    "",
		NewData:    ticket,
		System:     activityLog.SYSTEM_TYPE_TICKET,
	}

	go activityLog.LoggerActivity(logActivity)
}
