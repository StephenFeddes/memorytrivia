package repository

import (
	"log"

	"github.com/StephenFeddes/memorytrivia/apps/backend/services/account/internal/infrastructure/database"
)

type PostgresLobbyRepository struct {
	db *database.PostgresDB
}

func NewPostgresLobbyRepository(db *database.PostgresDB) *PostgresLobbyRepository {
	return &PostgresLobbyRepository{db: db}
}

func (r *PostgresLobbyRepository) CreateLobby(size uint32, ownerID uint32) error {
	tx, err := r.db.Client.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO member (id) VALUES ($1)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(ownerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	var lobbyID uint32
	stmt, err = tx.Prepare("INSERT INTO lobby (size, owner_id) VALUES ($1, $2) RETURNING id")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(size, ownerID).Scan(&lobbyID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	stmt, err = tx.Prepare("INSERT INTO lobby_member (lobby_id, member_id) VALUES ($1, $2)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(lobbyID, ownerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
