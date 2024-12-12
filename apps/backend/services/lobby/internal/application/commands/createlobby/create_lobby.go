package createlobby

import (
	"context"
	"errors"
)

type DatabaseLobbyCreator interface {
	CreateLobby(size uint32, ownerID uint32) error
}

type CreateLobby struct {
	db DatabaseLobbyCreator
}

func NewCreateLobby(db DatabaseLobbyCreator) *CreateLobby {
	return &CreateLobby{db: db}
}

func (c *CreateLobby) Execute(ctx context.Context, size uint32, ownerID uint32) (string, error) {
	// Validate input
	if size <= 1 {
		errMsg := "size must be greater than 1"
		return errMsg, errors.New(errMsg)
	} else if size > 4 {
		errMsg := "size must be less than or equal to 4"
		return errMsg, errors.New(errMsg)
	}

	// Create lobby
	err := c.db.CreateLobby(size, ownerID); if err != nil {
		return "", err
	}

	// Return success message
	return "lobby created", nil
}
