package actions

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Roll represents the roll action.
type Roll struct{}

// NewRoll creates a new say action.
func NewRoll() (Action, error) {
	return Roll{}, nil
}

// Process executes the roll action.
func (action Roll) Process(s *discordgo.Session, m *discordgo.MessageCreate) error {
	digits := 1
	min := 1
	
	if len(m.Content) == 18 {
	  digits = 2
	}

	start := 16
	fin := start + digits
	maxStr := m.Content[start:fin]

	if maxInt, err := strconv.Atoi(maxStr); err == nil {
		output := 0

		if maxInt <= 1 {
			s.ChannelMessageSend(m.ChannelID, "You somehow roll a 0...")
		}	else {
			output = rand.Intn(maxInt) + min
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%d", output))
		}
	}

	return nil
}
