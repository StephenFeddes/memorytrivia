package repository

import (
	"log"

	"github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/database"
	pb "github.com/StephenFeddes/memorytrivia/apps/backend/services/lobby/internal/infrastructure/grpc/protobufs"
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

func (r *PostgresLobbyRepository) GetLobby(id uint32) (*pb.Lobby, error) {
	lobby := &pb.Lobby{}
	err := r.db.Client.QueryRow(
		"SELECT id, size, owner_id FROM lobby WHERE id = $1", id,
	).Scan(&lobby.Id, &lobby.Size, &lobby.OwnerId)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Client.Query("SELECT member_id FROM lobby_member WHERE lobby_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var memberID uint32
		err = rows.Scan(&memberID)
		if err != nil {
			return nil, err
		}
		lobby.MemberIds = append(lobby.MemberIds, memberID)
	}

	return lobby, nil
}
