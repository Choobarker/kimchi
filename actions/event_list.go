package actions

import "github.com/bwmarrin/discordgo"

// EventList represents the event list action.
type EventList struct{}

// NewEventList creates a new event list action.
func NewEventList() (Action, error) {
	return EventList{}, nil
}

// Process executes the event list action.
func (action EventList) Process(s *discordgo.Session, m *discordgo.MessageCreate) error {
	s.ChannelMessageSend(m.ChannelID, "Here are your events")
	return nil
}
