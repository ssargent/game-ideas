package ledger

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ssargent/game-ideas/internal/engine"
)

type Ledger struct {
	ID       uuid.UUID `storm:"index"`
	PlayerID uuid.UUID `storm:"index"`
	Name     string    `storm:"index"`
	Type     string    `storm:"index"`
	Currency string
}

type Entry struct {
	ID        uuid.UUID
	LedgerID  uuid.UUID `storm:"index"`
	Action    string
	Amount    int64
	CreatedAt int64 `storm:"index"`
}

func (l *Ledger) Credit(e *engine.Engine, amount int64) error {
	le := Entry{
		ID:        uuid.New(),
		LedgerID:  l.ID,
		Action:    "CREDIT",
		Amount:    amount,
		CreatedAt: time.Now().Unix(),
	}

	if err := e.DB.Save(&le); err != nil {
		return fmt.Errorf("error crediting ledger %w", err)
	}

	return nil
}
func (l *Ledger) Debit(e *engine.Engine, amount int64) error {
	le := Entry{
		ID:        uuid.New(),
		LedgerID:  l.ID,
		Action:    "DEBIT",
		Amount:    amount,
		CreatedAt: time.Now().Unix(),
	}

	if err := e.DB.Save(&le); err != nil {
		return fmt.Errorf("error debiting ledger %w", err)
	}

	return nil
}
