// Package mysql assemble all functions required to perform SQL queries
package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Lord-Y/versions-api/commons"
	"github.com/Lord-Y/versions-api/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rs/zerolog/log"
	"github.com/syyongx/php2go"
)

// InitDB permit to initialiaze or migrate databases
func InitDB() {
	log.Debug().Msg("starting db initialization/migration")
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
		log.Fatal().Err(err).Msgf("could not ping DB: %v", err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Could not start sql migration with error msg: %v", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		sqlDIR,
		commons.SqlDriver,
		driver,
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("Migration failed: %v", err)
		return
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msgf("An error occurred while syncing the database with error msg: %v", err)
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
		log.Error().Err(err).Msg("Failed to connect to DB")
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO versions(`workload`, `platform`, `environment`, `version`, `changelog_url`, `raw`, `status`) VALUES(?,?,?,?,?,?,?)")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	res, err := stmt.Exec(
		php2go.Addslashes(d.Workload),
		php2go.Addslashes(d.Platform),
		php2go.Addslashes(d.Environment),
		php2go.Addslashes(d.Version),
		php2go.Addslashes(d.ChangelogURL),
		php2go.Addslashes(d.Raw),
		php2go.Addslashes(strings.ToLower(d.Status)),
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
		log.Error().Err(err).Msg("Failed to connect to DB")
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE versions SET status = ? WHERE versions_id = ?")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	_, err = stmt.Exec(
		php2go.Addslashes(strings.ToLower(d.Status)),
		d.VersionId,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()
	return nil
}

// ReadEnvironment permit to get data into sql instance
func ReadEnvironment(d models.ReadEnvironment) (z []map[string]interface{}, err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT *, (SELECT count(version) FROM versions WHERE workload = ? AND platform = ? AND environment = ?) total FROM versions WHERE workload = ? AND platform = ? AND environment = ? ORDER BY date DESC LIMIT ?, ?")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(
		php2go.Addslashes(d.Workload),
		php2go.Addslashes(d.Platform),
		php2go.Addslashes(d.Environment),
		php2go.Addslashes(d.Workload),
		php2go.Addslashes(d.Platform),
		php2go.Addslashes(d.Environment),
		d.StartLimit,
		d.EndLimit,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	m := make([]map[string]interface{}, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		sub := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = php2go.Stripslashes(string(col))
			}
			sub[columns[i]] = value
		}
		m = append(m, sub)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return m, nil
}

// ReadPlatform permit to get data into sql instance
func ReadPlatform(d models.ReadPlatform) (z []map[string]interface{}, err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT *, (SELECT count(version) FROM versions WHERE workload = ? AND platform = ?) total FROM versions WHERE workload = ? AND platform = ? ORDER BY date DESC LIMIT ?, ?")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(
		php2go.Addslashes(d.Workload),
		php2go.Addslashes(d.Platform),
		php2go.Addslashes(d.Workload),
		php2go.Addslashes(d.Platform),
		d.StartLimit,
		d.EndLimit,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	m := make([]map[string]interface{}, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		sub := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = php2go.Stripslashes(string(col))
			}
			sub[columns[i]] = value
		}
		m = append(m, sub)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return m, nil
}

// ReadHome permit to get data into sql instance
func ReadHome() (z []map[string]interface{}, err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM versions ORDER BY date DESC LIMIT 50")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil && err != sql.ErrNoRows {
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	m := make([]map[string]interface{}, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		sub := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = php2go.Stripslashes(string(col))
			}
			sub[columns[i]] = value
		}
		m = append(m, sub)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return m, nil
}

// ReadDistinctWorkloads permit to get data into sql instance
func ReadDistinctWorkloads() (z []map[string]interface{}, err error) {
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

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	m := make([]map[string]interface{}, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		sub := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = php2go.Stripslashes(string(col))
			}
			sub[columns[i]] = value
		}
		m = append(m, sub)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return m, nil
}

// Raw permit to get data from raw column instance
func Raw(d models.Raw) (z map[string]interface{}, err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT raw FROM versions WHERE workload = ? AND environment = ? AND version = ? LIMIT 1")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(
		php2go.Addslashes(d.Workload),
		php2go.Addslashes(d.Environment),
		php2go.Addslashes(d.Version),
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	m := make(map[string]interface{}, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = php2go.Stripslashes(string(col))
			}
			m[columns[i]] = value
		}
	}
	if err = rows.Err(); err != nil {
		return
	}
	return m, nil
}

// RawById permit to get data from raw by version_id column instance
func RawById(d models.RawById) (z map[string]interface{}, err error) {
	db, err := sql.Open(
		commons.SqlDriver,
		commons.BuildDSN(),
	)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM versions WHERE versions_id = ? LIMIT 1")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(
		d.VersionID,
	)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	m := make(map[string]interface{}, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = php2go.Stripslashes(string(col))
			}
			m[columns[i]] = value
		}
	}
	if err = rows.Err(); err != nil {
		return
	}
	return m, nil
}

func ReadForUnitTesting() (z map[string]string) {
	db, err := sql.Open("mysql", commons.BuildDSN())
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to DB")
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT versions_id, workload, platform, environment FROM versions LIMIT 1")
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil && err != sql.ErrNoRows {
		log.Error().Err(err).Msgf("Error occured from DB func - QueryRows %v - stmt %v", rows, stmt)
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Error().Err(err).Msg("Error occured from DB func")
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	m := make(map[string]string)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Error().Err(err).Msg("Error occured from DB func")
			return
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			m[columns[i]] = value
		}
	}
	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error occured from DB func")
		return
	}
	return m
}
