package metrics

import (
	"github.com/Lord-Y/versions/commons"
	"github.com/Lord-Y/versions/models"
	"github.com/Lord-Y/versions/mysql"
	"github.com/Lord-Y/versions/postgres"
	"github.com/rs/zerolog/log"
)

func GetLastXDaysDeployments() (z []models.DBGetLastXDaysDeployments) {
	var err error

	var result []models.DBGetLastXDaysDeployments
	if commons.SqlDriver == "mysql" {
		result, err = mysql.GetLastXDaysDeployments()
	} else {
		result, err = postgres.GetLastXDaysDeployments()
	}
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing database query")
		return
	}
	return result
}
