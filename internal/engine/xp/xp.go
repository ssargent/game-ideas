package xp

import (
	"strings"

	"github.com/google/uuid"
	"github.com/ssargent/game-ideas/internal/engine"
	"github.com/ssargent/game-ideas/internal/engine/messaging"
)

type ExperienceService interface {
	messaging.MessageService
	Total(player uuid.UUID) (int, error)
}

type experienceService struct {
	game *engine.Engine
}

// Accept implements ExperienceService
func (e *experienceService) Accept(g *messaging.GameMessage) bool {
	if strings.HasPrefix(g.Type, "kill-creature::") ||
		strings.HasPrefix(g.Type, "attack-creature::") ||
		strings.HasSuffix(g.Type, "block-attack::") {
		return true
	}

	return false
}

// Name implements ExperienceService
func (e *experienceService) Name() string {
	return "xp"
}

// Receive implements ExperienceService
func (e *experienceService) Receive(msg *messaging.GameMessage) error {
	panic("unimplemented")
}

// Total implements ExperienceService
func (e *experienceService) Total(player uuid.UUID) (int, error) {
	panic("unimplemented")
}

func NewExperienceService(game *engine.Engine) ExperienceService {
	return &experienceService{game: game}
}
