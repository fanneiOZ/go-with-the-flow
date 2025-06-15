package donation

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"shareddomain/entity"
	"sharedinfra/db"
	"sharedinfra/db/postgres"
	"time"
)

type campaignDbState struct {
	id                   int
	version              uint
	title, description   string
	createdAt, updatedAt time.Time
}

type CampaignRepo struct {
	dbContext *postgres.Context
}

func NewCampaignRepo(pgContext *postgres.Context) *CampaignRepo {
	return &CampaignRepo{pgContext}
}

func (repo *CampaignRepo) FindById(id int) (*Campaign, error) {
	conn, err := repo.dbContext.Acquire()
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(
		context.Background(),
		"SELECT * FROM donation.campaign WHERE id = $1",
		id,
	)

	var dbState campaignDbState
	err = row.Scan(
		&dbState.id,
		&dbState.version,
		&dbState.title,
		&dbState.description,
		&dbState.createdAt,
		&dbState.updatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return NewCampaign(
			Campaign{
				id:          dbState.id,
				Title:       dbState.title,
				Description: dbState.description,
				Active:      true,
				createdAt:   dbState.createdAt,
				updatedAt:   dbState.updatedAt,

				Version: entity.NewVersion(dbState.version),
			},
		),
		nil
}

func (repo *CampaignRepo) Save(campaign *Campaign) (*Campaign, error) {
	conn, err := repo.dbContext.Acquire()
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	isNewEntity := campaign.id == 0

	if isNewEntity {
		return repo.insert(conn, campaign)
	}

	return repo.update(conn, campaign)
}

func (repo *CampaignRepo) insert(conn *pgxpool.Conn, campaign *Campaign) (*Campaign, error) {
	result, err := conn.Query(
		context.Background(),
		`
			INSERT INTO donation.campaign (title, description, version)
			VALUES ($1, $2, $3)
			RETURNING id, version, title, description, created_at, updated_at
			`,
		campaign.Title,
		campaign.Description,
		campaign.EntityVersion(),
	)

	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, db.ErrNoReturningResult
	}

	var dbState campaignDbState
	err = result.Scan(&dbState.id, &dbState.version, &dbState.title, &dbState.description, &dbState.createdAt, &dbState.updatedAt)
	if err != nil {
		return nil, err
	}

	return NewCampaign(
			Campaign{
				id:          dbState.id,
				Title:       dbState.title,
				Description: dbState.description,
				Active:      true,
				createdAt:   dbState.createdAt,
				updatedAt:   dbState.updatedAt,

				Version: entity.NewVersion(dbState.version),
			},
		),
		nil
}

func (repo *CampaignRepo) update(conn *pgxpool.Conn, campaign *Campaign) (*Campaign, error) {
	result, err := conn.Query(
		context.Background(),
		`
		UPDATE donation.campaign
		SET title = $3, description = $4, updated_at = current_timestamp, version = $5
		WHERE id = $1 AND version = $2
		RETURNING updated_at, version
		`,
		campaign.Id(),
		campaign.EntityVersion(),
		campaign.Title,
		campaign.Description,
		campaign.Version.Next().EntityVersion(),
	)

	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, db.ErrEntityVersionConflicted
	}

	var updatedAt time.Time
	var nextVersion uint
	_ = result.Scan(&updatedAt, &nextVersion)
	log.Print(nextVersion)

	return NewCampaign(
			Campaign{
				id:          campaign.id,
				Title:       campaign.Title,
				Description: campaign.Description,
				Active:      campaign.Active,
				createdAt:   campaign.createdAt,
				updatedAt:   updatedAt,
				Version:     entity.NewVersion(nextVersion),
			},
		),
		nil
}
