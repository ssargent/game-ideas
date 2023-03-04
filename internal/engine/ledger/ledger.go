package ledger

import (
	"errors"
	"fmt"
	"time"

	"github.com/asdine/storm/q"
	"github.com/asdine/storm/v3"
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

func CreateLedger(e *engine.Engine, l *Ledger) (*Ledger, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("error creating ledger id %w", err)
	}

	l.ID = id

	if err := e.DB.Save(l); err != nil {
		return nil, fmt.Errorf("error creating ledger %w", err)
	}

	return l, nil
}

func GetLedger(e *engine.Engine, l *Ledger) (*Ledger, error) {
	var ledgers []Ledger
	if err := e.DB.Select(
		q.Eq("PlayerID", l.PlayerID),
		q.Eq("Name", l.Name),
		q.Eq("Type", l.Type),
		q.Eq("Currency", l.Currency),
	).Find(&ledgers); err != nil {
		return nil, err
	}

	// we should have only one
	if len(ledgers) > 1 {
		return nil, errors.New("database error, multiple ledgers found")
	}

	return &ledgers[0], nil
}

func GetOrCreateLedger(e *engine.Engine, l *Ledger) (*Ledger, error) {
	foundLedger, err := GetLedger(e, l)
	if err != nil && err != storm.ErrNotFound {
		return nil, fmt.Errorf("error getting ledger %w", err)
	}

	if foundLedger != nil {
		return foundLedger, nil
	}

	if err == storm.ErrNotFound {
		createdledger, err := CreateLedger(e, l)
		if err != nil {
			return nil, fmt.Errorf("error creating ledger %w", err)
		}

		return createdledger, nil
	}

	return nil, nil
}

func Init(e *engine.Engine) {
	e.DB.Init(&Ledger{})
	e.DB.Init(&Entry{})
}
