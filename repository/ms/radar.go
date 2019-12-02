package ms

import (
	"github.com/fly365com/rabbit/message/activityLog"
	"ticket/entity"
)

type RadarMsRepo struct {
}

func NewRadarMs() *RadarMsRepo {
	return &RadarMsRepo{}
}

func (rdrMs *RadarMsRepo) LogCreateTicketActivity(ticket *entity.Ticket) {
	logActivity := activityLog.LogActivity{
		Identifier: ticket.Email,
		Action:     activityLog.LOG_ACTIVITY_CREATE_TICKET,
		OldData:    "",
		NewData:    ticket,
		System:     activityLog.SYSTEM_TYPE_TICKET,
	}

	go activityLog.LoggerActivity(logActivity)
}
