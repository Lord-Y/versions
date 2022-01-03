// Package mysql assemble all functions required to perform SQL queries
package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Lord-Y/versions/commons"
	customLogger "github.com/Lord-Y/versions/logger"
	"github.com/Lord-Y/versions/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rs/zerolog/log"
)

func init() {
	customLogger.SetLoggerLogLevel()
}

// InitDB permit to initialiaze or migrate databases
func InitDB() {
	log.Debug().Msg("Starting db initialization/migration")
	fileDir, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("Not able to get current directory")
	}
	log.Debug().Msgf("Use db sql driver %s", commons.SqlDriver)
	sqlDIR := fmt.Sprintf("file://%s%s", fileDir, "/sql/mysql")
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
	driver, err := mysql.WithInstance(db, &mysql.Config{})
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
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO versions(`workload`, `platform`, `environment`, `version`, `changelog_url`, `raw`, `status`) VALUES(?,?,?,?,?,?,?)")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	res, err := stmt.Exec(
		d.Workload,
		d.Platform,
		d.Environment,
		d.Version,
		d.ChangelogURL,
		d.Raw,
		strings.ToLower(d.Status),
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return
	}
	defer stmt.Close()
	return lastInsertId, nil
}

// UpdateStatus permit to insert data into sql instance
func UpdateStatus(d models.UpdateStatus) (err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE versions SET status = ? WHERE versions_id = ?")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	_, err = stmt.Exec(
		strings.ToLower(d.Status),
		d.VersionId,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()
	return nil
}

// ReadEnvironment permit to get data into sql instance
func ReadEnvironment(d models.ReadEnvironment) (z []models.DBReadCommon, err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT *, (SELECT COUNT(version) FROM versions WHERE workload = ? AND platform = ? AND environment = ?) total FROM versions WHERE workload = ? AND platform = ? AND environment = ? ORDER BY date DESC LIMIT ?, ?")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(
		d.Workload,
		d.Platform,
		d.Environment,
		d.Workload,
		d.Platform,
		d.Environment,
		d.StartLimit,
		d.EndLimit,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}

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
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	var query string
	if d.Whatever {
		query = "SELECT version FROM versions WHERE workload = ? AND platform = ? AND environment = ? ORDER BY date DESC LIMIT 1"
	} else {
		query = "SELECT version FROM versions WHERE workload = ? AND platform = ? AND environment = ? AND status IN ('completed', 'deployed') ORDER BY date DESC LIMIT 1"
	}

	err = db.QueryRowContext(
		ctx,
		query,
		d.Workload,
		d.Platform,
		d.Environment,
	).Scan(
		&z.Version,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return z, nil
}

// ReadPlatform permit to get data into sql instance
func ReadPlatform(d models.ReadPlatform) (z []models.DBReadCommon, err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT *, (SELECT COUNT(version) FROM versions WHERE workload = ? AND platform = ?) total FROM versions WHERE workload = ? AND platform = ? ORDER BY date DESC LIMIT ?, ?")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(
		d.Workload,
		d.Platform,
		d.Workload,
		d.Platform,
		d.StartLimit,
		d.EndLimit,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}

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
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM versions ORDER BY date DESC LIMIT 25")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil && err != sql.ErrNoRows {
		return
	}

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
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT DISTINCT workload,platform,environment FROM versions ORDER BY workload,platform,environment ASC")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil && err != sql.ErrNoRows {
		return
	}

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
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	err = db.QueryRowContext(
		ctx,
		"SELECT raw FROM versions WHERE workload = ? AND environment = ? AND version = ? LIMIT 1",
		d.Workload,
		d.Environment,
		d.Version,
	).Scan(
		&z.Raw,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return z, nil
}

// RawById permit to get data from raw by version_id column instance
func RawById(d models.RawById) (z models.DBCommons, err error) {
	ctx := context.Background()
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	err = db.QueryRowContext(
		ctx,
		"SELECT * FROM versions WHERE versions_id = ? LIMIT 1",
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
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return z, nil
}

func ReadForUnitTesting(status string) (z models.DBCommons, err error) {
	ctx := context.Background()
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	err = db.QueryRowContext(
		ctx,
		"SELECT * FROM versions WHERE status = ? LIMIT 1",
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
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return z, nil
}

// GetLastXDaysDeployments permit to get data into sql instance
func GetLastXDaysDeployments() (z []models.DBGetLastXDaysDeployments, err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	// SELECT COUNT(versions_id) total,workload,platform,environment,status,DATE_FORMAT(date,'%Y-%m-%d 00:00:00') date FROM versions WHERE DATE_FORMAT(date,'%Y-%m-%d 00:00:00') >= (DATE_FORMAT(NOW(),'%Y-%m-%d 00:00:00') - INTERVAL 7 DAY) GROUP BY status,workload,platform,environment,date;
	// I cannot use previous query as go will throw on error while parsing sql date_format
	// and using a golang function with time.Parse is not helping either
	stmt, err := db.Prepare("SELECT COUNT(versions_id) total,workload,platform,environment,status,date FROM versions WHERE DATE_FORMAT(date,'%Y-%m-%d') >= (DATE_FORMAT(NOW(),'%Y-%m-%d') - INTERVAL 10 DAY) GROUP BY status,workload,platform,environment,date")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil && err != sql.ErrNoRows {
		return
	}

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

// func df(d time.Time) *time.Time {
// 	strDate := time.Time(time.Unix(int64(d.Unix()), 0)).Format("2021-12-16")
// 	log.Info().Msgf("dateeeeeeeeee %s strdate %s", d, strDate)
// 	layout := "2021-12-16"
// 	t, err := time.Parse(layout, strDate)
// 	if err != nil {
// 		log.Error().Err(err).Msg("fuck")
// 	}
// 	return &t
// }
