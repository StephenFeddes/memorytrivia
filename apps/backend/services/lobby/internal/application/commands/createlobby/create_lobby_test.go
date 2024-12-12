package createlobby

import (
	"context"
	"fmt"
	"testing"
)

type mockDatabaseLobbyCreator struct {
	shouldReturnError bool
}

// CreateLobby implements DatabaseLobbyCreator.
func (m *mockDatabaseLobbyCreator) CreateLobby(size uint32, ownerID uint32) error {
	if m.shouldReturnError {
		return fmt.Errorf("error creating lobby")
	}
	return nil
}

func TestCreateLobby(t *testing.T) {
	t.Run("should create a lobby", func(t *testing.T) {
		// Arrange
		createLobby := NewCreateLobby(&mockDatabaseLobbyCreator{shouldReturnError: false})

		// Act
		status, err := createLobby.Execute(context.TODO(), 2, 1)

		// Assert
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		} else if status != "lobby created" {
			t.Fatalf("expected status: %q, got: %q", "lobby created", status)
		}
	})

	t.Run("should fail if size is less than 2", func(t *testing.T) {
		// Arrange
		createLobby := NewCreateLobby(&mockDatabaseLobbyCreator{shouldReturnError: false})

		// Act
		_, err := createLobby.Execute(context.TODO(), 1, 1)

		// Assert
		expectedMessage := "size must be greater than 1"
		if err == nil || err.Error() != expectedMessage {
			t.Fatalf("expected error message: %q, got: %q", expectedMessage, err)
		}
	})

	t.Run("should fail if size is less than 2", func(t *testing.T) {
		// Arrange
		createLobby := NewCreateLobby(&mockDatabaseLobbyCreator{shouldReturnError: false})

		// Act
		_, err := createLobby.Execute(context.TODO(), 1, 1)

		// Assert
		expectedMessage := "size must be greater than 1"
		if err == nil || err.Error() != expectedMessage {
			t.Fatalf("expected error message: %q, got: %q", expectedMessage, err)
		}
	})

	t.Run("should fail if database returns error when creating the lobby", func(t *testing.T) {
		// Arrange
		createLobby := NewCreateLobby(&mockDatabaseLobbyCreator{shouldReturnError: true})

		// Act
		_, err := createLobby.Execute(context.TODO(), 2, 1)
	
		// Assert
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
