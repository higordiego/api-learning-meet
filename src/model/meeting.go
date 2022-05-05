package model

import (
	"fmt"
	"golang-meeting/src/database"
)

type Meeting struct {
	Resource      string `json:"resource"`
	UserId        int    `json:"user_id"`
	Topic         string `json:"topic"`
	ApplicationId int    `json:"application_id"`
	Attemps       int    `json:"attempts"`
	Sent          string `json:"sent"`
	Received      string `json:"received"`
}

func (m *Meeting) Created() error {

	sql, err := database.Connect()

	if err != nil {
		return err
	}

	insertMeeting, err := sql.Prepare(`insert into meetings(resource, user_id, topic, application_id, attempts, sent, received) values(?, ?, ?, ?, ?, ?, ?);`)

	if err != nil {
		return err
	}

	if _, err := insertMeeting.Exec(m.Resource, m.UserId, m.Topic, m.ApplicationId, m.Attemps, m.Sent, m.Received); err != nil {

		return err
	}

	return sql.Close()
}
