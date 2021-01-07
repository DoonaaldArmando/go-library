package message

import (
	"time"

	utils "github.com/DoonaaldArmando/go-library/init/utils"

	"github.com/google/uuid"
)

// IMessage  someting more  //
type IMessage interface {
	NewMessage(
		userAgent string,
		user string,
		contentType string,
		eventCode string,
		correlationID string,
		source string,
		messageType string,
		data interface{},
	) Message
}

// Message //
type Message struct {
	ID     string
	Audit  audit
	Header header
	Data   interface{}
}

// Audit //
type audit struct {
	Time      string
	IP        string
	UserAgent string
	User      string
	Source    string
	Version   string
}

// Header //
type header struct {
	MessageType   string
	ContentType   string
	CorrelationID string
	EventCode     string
}

// NewMessage //
func (Message) NewMessage(
	userAgent string,
	user string,
	contentType string,
	eventCode string,
	correlationID string,
	source string,
	messageType string,
	data interface{},
) Message {
	var ip utils.GetIP = utils.IP{}
	ipv4, err := ip.GetIP()

	if err != nil {
		panic(err)
	}

	return Message{
		ID: uuid.New().String(),
		Audit: audit{
			Time:      time.Now().String(),
			IP:        ipv4,
			UserAgent: userAgent,
			User:      user,
			Source:    source,
			Version:   "TODO",
		},
		Header: header{
			MessageType:   messageType,
			ContentType:   contentType,
			CorrelationID: correlationID,
			EventCode:     eventCode,
		},
		Data: data,
	}
}
