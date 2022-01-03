// Package postgres assemble all functions required to perform SQL queries
package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Lord-Y/versions/commons"
	customLogger "github.com/Lord-Y/versions/logger"
	"github.com/Lord-Y/versions/models"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

func init() {
	customLogger.SetLoggerLogLevel()
}

// InitDB permit to initialiaze or migrate databases
func InitDB() {
	log.Debug().Msg("starting db initialization/migration")
	fileDir, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("Not able to get current directory")
	}
	log.Debug().Msgf("Use db sql driver %s", commons.SqlDriver)
	sqlDIR := fmt.Sprintf("file://%s%s", fileDir, "/sql/postgres")
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to connect to DB")
		return
	}
	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msgf("could not ping DB: %s", err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Could not start sql migration with error msg: %s", err.Error())
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		sqlDIR,
		commons.SqlDriver,
		driver,
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("Migration failed: %s", err.Error())
		return
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msgf("An error occurred while syncing the database with error msg: %s", err.Error())
		return
	}
	defer db.Close()
	log.Info().Msg("Database migrated successfully")
}

// Ping permit to ping sql instance
func Ping() (b bool) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while connecting to DB")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msg("Error occured while pinging DB")
		return
	}
	return true
}

// Create permit to insert data into sql instance
func Create(d models.Create) (z int64, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	tx, err := db.Begin(ctx)
	if err != nil {
		return
	}
	//golangci-lint fail on this check while the transaction error is checked
	defer tx.Rollback(ctx) //nolint

	err = tx.QueryRow(
		ctx,
		"INSERT INTO versions(workload, platform, environment, version, changelog_url, raw, status) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING versions_id",
		d.Workload,
		d.Platform,
		d.Environment,
		d.Version,
		d.ChangelogURL,
		d.Raw,
		strings.ToLower(d.Status),
	).Scan(
		&z,
	)
	if err = tx.Commit(ctx); err != nil {
		return
	}
	return
}

// UpdateStatus permit to insert data into sql instance
func UpdateStatus(d models.UpdateStatus) (err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	tx, err := db.Begin(ctx)
	if err != nil {
		return
	}
	//golangci-lint fail on this check while the transaction error is checked
	defer tx.Rollback(ctx) //nolint

	_, err = tx.Exec(
		ctx,
		"UPDATE versions SET status = $1 WHERE versions_id = $2",
		strings.ToLower(d.Status),
		d.VersionId,
	)
	if err = tx.Commit(ctx); err != nil {
		return
	}
	return
}

// ReadEnvironment permit to get data into sql instance
func ReadEnvironment(d models.ReadEnvironment) (z []models.DBReadCommon, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(
		ctx,
		"SELECT *, (SELECT count(version) FROM versions WHERE workload = $1 AND platform = $2 AND environment = $3) total FROM versions WHERE workload = $1 AND platform = $2 AND environment = $3 ORDER BY date DESC OFFSET $4 LIMIT $5",
		d.Workload,
		d.Platform,
		d.Environment,
		d.StartLimit,
		d.EndLimit,
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var x models.DBReadCommon
		if err = rows.Scan(
			&x.Versions_id,
			&x.Workload,
			&x.Platform,
			&x.Environment,
			&x.Version,
			&x.Changelog_url,
			&x.Raw,
			&x.Status,
			&x.Date,
			&x.Total,
		); err != nil {
			return
		}
		z = append(z, x)
	}
	return z, nil
}

// ReadEnvironmentLatest permit to get latest version with status deployed or completed
func ReadEnvironmentLatest(d models.ReadEnvironmentLatest) (z models.DbVersion, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	var query string
	if d.Whatever {
		query = "SELECT version FROM versions WHERE workload = $1 AND platform = $2 AND environment = $3 ORDER BY date DESC LIMIT 1"
	} else {
		query = "SELECT version FROM versions WHERE workload = $1 AND platform = $2 AND environment = $3 AND status IN ('completed', 'deployed') ORDER BY date DESC LIMIT 1"
	}

	err = db.QueryRow(
		ctx,
		query,
		d.Workload,
		d.Platform,
		d.Environment,
	).Scan(
		&z.Version,
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	return z, nil
}

// ReadPlatform permit to get data into sql instance
func ReadPlatform(d models.ReadPlatform) (z []models.DBReadCommon, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(
		ctx,
		"SELECT *, (SELECT count(version) FROM versions WHERE workload = $1 AND platform = $2) total FROM versions WHERE workload = $1 AND platform = $2 ORDER BY date DESC OFFSET $3 LIMIT $4",
		d.Workload,
		d.Platform,
		d.StartLimit,
		d.EndLimit,
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var x models.DBReadCommon
		if err = rows.Scan(
			&x.Versions_id,
			&x.Workload,
			&x.Platform,
			&x.Environment,
			&x.Version,
			&x.Changelog_url,
			&x.Raw,
			&x.Status,
			&x.Date,
			&x.Total,
		); err != nil {
			return
		}
		z = append(z, x)
	}
	return z, nil
}

// ReadHome permit to get data into sql instance
func ReadHome() (z []models.DBCommons, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(
		ctx,
		"SELECT * FROM versions ORDER BY date DESC LIMIT 25",
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var x models.DBCommons
		if err = rows.Scan(
			&x.Versions_id,
			&x.Workload,
			&x.Platform,
			&x.Environment,
			&x.Version,
			&x.Changelog_url,
			&x.Raw,
			&x.Status,
			&x.Date,
		); err != nil {
			return
		}
		z = append(z, x)
	}
	return z, nil
}

// ReadDistinctWorkloads permit to get data into sql instance
func ReadDistinctWorkloads() (z []models.DBReadDistinctWorkloads, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(
		ctx,
		"SELECT DISTINCT workload,platform,environment FROM versions ORDER BY workload,platform,environment ASC",
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var x models.DBReadDistinctWorkloads
		if err = rows.Scan(
			&x.Workload,
			&x.Platform,
			&x.Environment,
		); err != nil {
			return
		}
		z = append(z, x)
	}
	return z, nil
}

// Raw permit to get data from raw column instance
func Raw(d models.Raw) (z models.DBRaw, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	err = db.QueryRow(
		ctx,
		"SELECT raw FROM versions WHERE workload = $1 AND environment = $2 AND version = $3 LIMIT 1",
		d.Workload,
		d.Environment,
		d.Version,
	).Scan(
		&z.Raw,
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	return z, nil
}

// RawById permit to get data from raw by version_id column instance
func RawById(d models.RawById) (z models.DBCommons, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	err = db.QueryRow(
		ctx,
		"SELECT * FROM versions WHERE versions_id = $1 LIMIT 1",
		d.VersionID,
	).Scan(
		&z.Versions_id,
		&z.Workload,
		&z.Platform,
		&z.Environment,
		&z.Version,
		&z.Changelog_url,
		&z.Raw,
		&z.Status,
		&z.Date,
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	return z, nil
}

func ReadForUnitTesting(status string) (z models.DBCommons, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	err = db.QueryRow(
		ctx,
		"SELECT * FROM versions WHERE status = $1 LIMIT 1",
		status,
	).Scan(
		&z.Versions_id,
		&z.Workload,
		&z.Platform,
		&z.Environment,
		&z.Version,
		&z.Changelog_url,
		&z.Raw,
		&z.Status,
		&z.Date,
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	return z, nil
}

// GetLastXDaysDeployments permit to get data into sql instance
func GetLastXDaysDeployments() (z []models.DBGetLastXDaysDeployments, err error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, commons.BuildDSN())
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(
		ctx,
		"SELECT COUNT(versions_id) total,workload,platform,environment,status,TO_DATE(to_char(date,'YYYY-MM-DD'),'YYYY-MM-DD') date FROM versions WHERE TO_DATE(to_char(date,'YYYY-MM-DD'),'YYYY-MM-DD') >= (NOW() - INTERVAL '10 DAY') GROUP BY status,workload,platform,environment,TO_DATE(to_char(date,'YYYY-MM-DD'),'YYYY-MM-DD')",
	)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var x models.DBGetLastXDaysDeployments
		if err = rows.Scan(
			&x.Total,
			&x.Workload,
			&x.Platform,
			&x.Environment,
			&x.Status,
			&x.Date,
		); err != nil {
			return
		}
		z = append(z, x)
	}
	return z, nil
}
