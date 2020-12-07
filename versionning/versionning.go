// Package versionning assemble all functions required all http endpoints
package versionning

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Lord-Y/versions-api/cache"
	"github.com/Lord-Y/versions-api/commons"
	"github.com/Lord-Y/versions-api/models"
	"github.com/Lord-Y/versions-api/mysql"
	"github.com/Lord-Y/versions-api/postgres"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var (
	cacheExpire = time.Duration(86400 * 30)
)

// Create permit to insert new deployment in DB
func Create(c *gin.Context) {
	var (
		d      models.Create
		err    error
		result int64
	)
	if err = c.ShouldBind(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if commons.SqlDriver == "mysql" {
		result, err = mysql.Create(d)
	} else {
		result, err = postgres.Create(d)
	}
	log.Info().Msgf("result %v", result)

	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing db query")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	} else {
		if commons.RedisEnabled() {
			cache.RedisDeleteKeysHasPrefix(commons.GetRedisURI(), []string{
				"w_p_",
			})
		}
		c.JSON(http.StatusCreated, gin.H{"versionId": result})
	}
}

// UpdateStatus permit to update status of deployment in DB
func UpdateStatus(c *gin.Context) {
	var (
		d   models.UpdateStatus
		err error
	)
	if err = c.ShouldBind(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if commons.SqlDriver == "mysql" {
		err = mysql.UpdateStatus(d)
	} else {
		err = postgres.UpdateStatus(d)
	}

	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing db query")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	} else {
		if commons.RedisEnabled() {
			cache.RedisDeleteKeysHasPrefix(commons.GetRedisURI(), []string{
				"w_p_",
			})
		}
		c.JSON(http.StatusOK, "OK")
	}
}

// ReadEnvironment permit to get new deployment in DB
func ReadEnvironment(c *gin.Context) {
	var (
		d   models.ReadEnvironment
		err error
	)

	if err = c.ShouldBind(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d.StartLimit, d.EndLimit = commons.GetPagination(d.Page, 0, d.RangeLimit, d.RangeLimit)

	if commons.RedisEnabled() {
		keyName := fmt.Sprintf("w_p_e_%x", commons.GetMD5HashWithSum(fmt.Sprintf("%v", d)))
		result, err := cache.RedisGet(commons.GetRedisURI(), keyName)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while retrieving data from cache")
		}
		if len(result) > 0 {
			var a interface{}
			err = json.Unmarshal([]byte(result), &a)
			if err != nil {
				log.Error().Err(err).Msg("Error occured while unmarshalling data")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(http.StatusOK, a)
		} else {
			result := make([]map[string]interface{}, 0)
			if commons.SqlDriver == "mysql" {
				result, err = mysql.ReadEnvironment(d)
			} else {
				result, err = postgres.ReadEnvironment(d)
			}
			if err != nil {
				log.Error().Err(err).Msg("Error occured while performing db query")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			if len(result) == 0 {
				c.AbortWithStatus(404)
			} else {
				b, err := json.Marshal(result)
				if err != nil {
					log.Error().Err(err).Msg("Error occured while marshalling data")
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				} else {
					cache.RedisSet(commons.GetRedisURI(), keyName, string(b), cacheExpire)
					c.JSON(http.StatusOK, result)
				}
			}
		}
	} else {
		result := make([]map[string]interface{}, 0)
		if commons.SqlDriver == "mysql" {
			result, err = mysql.ReadEnvironment(d)
		} else {
			result, err = postgres.ReadEnvironment(d)
		}
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing db query")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if len(result) == 0 {
			c.AbortWithStatus(404)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// ReadPlatform permit to get new deployment in DB
func ReadPlatform(c *gin.Context) {
	var (
		d   models.ReadPlatform
		err error
	)

	if err = c.ShouldBind(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d.StartLimit, d.EndLimit = commons.GetPagination(d.Page, 0, d.RangeLimit, d.RangeLimit)

	if commons.RedisEnabled() {
		keyName := fmt.Sprintf("w_p_%x", commons.GetMD5HashWithSum(fmt.Sprintf("%v", d)))
		result, err := cache.RedisGet(commons.GetRedisURI(), keyName)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while retrieving data from cache")
		}
		if len(result) > 0 {
			var a interface{}
			err = json.Unmarshal([]byte(result), &a)
			if err != nil {
				log.Error().Err(err).Msg("Error occured while unmarshalling data")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(http.StatusOK, a)
		} else {
			result := make([]map[string]interface{}, 0)
			if commons.SqlDriver == "mysql" {
				result, err = mysql.ReadPlatform(d)
			} else {
				result, err = postgres.ReadPlatform(d)
			}
			if err != nil {
				log.Error().Err(err).Msg("Error occured while performing db query")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			if len(result) == 0 {
				c.AbortWithStatus(404)
			} else {
				b, err := json.Marshal(result)
				if err != nil {
					log.Error().Err(err).Msg("Error occured while marshalling data")
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				} else {
					cache.RedisSet(commons.GetRedisURI(), keyName, string(b), cacheExpire)
					c.JSON(http.StatusOK, result)
				}
			}
		}
	} else {
		result := make([]map[string]interface{}, 0)
		if commons.SqlDriver == "mysql" {
			result, err = mysql.ReadPlatform(d)
		} else {
			result, err = postgres.ReadPlatform(d)
		}
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing db query")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if len(result) == 0 {
			c.AbortWithStatus(404)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// ReadHome permit to get last nth deployment in DB
func ReadHome(c *gin.Context) {
	var (
		err error
	)
	if commons.RedisEnabled() {
		keyName := fmt.Sprintf("w_p_e_home")
		result, err := cache.RedisGet(commons.GetRedisURI(), keyName)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while retrieving data from cache")
		}
		if len(result) > 0 {
			var a interface{}
			err = json.Unmarshal([]byte(result), &a)
			if err != nil {
				log.Error().Err(err).Msg("Error occured while unmarshalling data")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(http.StatusOK, a)
		} else {
			result := make([]map[string]interface{}, 0)
			if commons.SqlDriver == "mysql" {
				result, err = mysql.ReadHome()
			} else {
				result, err = postgres.ReadHome()
			}
			if err != nil {
				log.Error().Err(err).Msg("Error occured while performing db query")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			if len(result) == 0 {
				c.AbortWithStatus(204)
			} else {
				b, err := json.Marshal(result)
				if err != nil {
					log.Error().Err(err).Msg("Error occured while marshalling data")
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				} else {
					cache.RedisSet(commons.GetRedisURI(), keyName, string(b), cacheExpire)
					c.JSON(http.StatusOK, result)
				}
			}
		}
	} else {
		result := make([]map[string]interface{}, 0)
		if commons.SqlDriver == "mysql" {
			result, err = mysql.ReadHome()
		} else {
			result, err = postgres.ReadHome()
		}
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing db query")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if len(result) == 0 {
			c.AbortWithStatus(204)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetDistinctWorkload permit to get disctinct workload from DB
func ReadDistinctWorkloads(c *gin.Context) {
	var (
		err error
	)
	if commons.RedisEnabled() {
		keyName := fmt.Sprintf("w_p_e_distinct_workload")
		result, err := cache.RedisGet(commons.GetRedisURI(), keyName)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while retrieving data from cache")
		}
		if len(result) > 0 {
			var a interface{}
			err = json.Unmarshal([]byte(result), &a)
			if err != nil {
				log.Error().Err(err).Msg("Error occured while unmarshalling data")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(http.StatusOK, a)
		} else {
			result := make([]map[string]interface{}, 0)
			if commons.SqlDriver == "mysql" {
				result, err = mysql.ReadDistinctWorkloads()
			} else {
				result, err = postgres.ReadDistinctWorkloads()
			}
			if err != nil {
				log.Error().Err(err).Msg("Error occured while performing db query")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			if len(result) == 0 {
				c.AbortWithStatus(204)
			} else {
				b, err := json.Marshal(result)
				if err != nil {
					log.Error().Err(err).Msg("Error occured while marshalling data")
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				} else {
					cache.RedisSet(commons.GetRedisURI(), keyName, string(b), cacheExpire)
					c.JSON(http.StatusOK, result)
				}
			}
		}
	} else {
		result := make([]map[string]interface{}, 0)
		if commons.SqlDriver == "mysql" {
			result, err = mysql.ReadDistinctWorkloads()
		} else {
			result, err = postgres.ReadDistinctWorkloads()
		}
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing db query")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if len(result) == 0 {
			c.AbortWithStatus(204)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// Raw permit to get raw data from DB
func Raw(c *gin.Context) {
	var (
		d   models.Raw
		err error
	)
	if err = c.ShouldBind(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if commons.RedisEnabled() {
		keyName := fmt.Sprintf("w_p_e_%x", commons.GetMD5HashWithSum(fmt.Sprintf("raw_%v", d)))
		result, err := cache.RedisGet(commons.GetRedisURI(), keyName)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while retrieving data from cache")
		}
		if len(result) > 0 {
			var a interface{}
			err = json.Unmarshal([]byte(result), &a)
			if err != nil {
				log.Error().Err(err).Msg("Error occured while unmarshalling data")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(http.StatusOK, a)
		} else {
			result := make(map[string]interface{}, 0)
			if commons.SqlDriver == "mysql" {
				result, err = mysql.Raw(d)
			} else {
				result, err = postgres.Raw(d)
			}
			if err != nil {
				log.Error().Err(err).Msg("Error occured while performing db query")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			if len(result) == 0 {
				c.AbortWithStatus(404)
			} else {
				b, err := json.Marshal(result)
				if err != nil {
					log.Error().Err(err).Msg("Error occured while marshalling data")
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				} else {
					cache.RedisSet(commons.GetRedisURI(), keyName, string(b), cacheExpire)
					c.JSON(http.StatusOK, result)
				}
			}
		}
	} else {
		result := make(map[string]interface{}, 0)
		if commons.SqlDriver == "mysql" {
			result, err = mysql.Raw(d)
		} else {
			result, err = postgres.Raw(d)
		}
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing db query")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if len(result) == 0 {
			c.AbortWithStatus(404)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// RawById permit to get raw by id data from DB
func RawById(c *gin.Context) {
	var (
		d   models.RawById
		err error
	)
	if err = c.ShouldBind(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if commons.RedisEnabled() {
		keyName := fmt.Sprintf("w_p_e_%x", commons.GetMD5HashWithSum(fmt.Sprintf("raw_%v", d)))
		result, err := cache.RedisGet(commons.GetRedisURI(), keyName)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while retrieving data from cache")
		}
		if len(result) > 0 {
			var a interface{}
			err = json.Unmarshal([]byte(result), &a)
			if err != nil {
				log.Error().Err(err).Msg("Error occured while unmarshalling data")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(http.StatusOK, a)
		} else {
			result := make(map[string]interface{}, 0)
			if commons.SqlDriver == "mysql" {
				result, err = mysql.RawById(d)
			} else {
				result, err = postgres.RawById(d)
			}
			if err != nil {
				log.Error().Err(err).Msg("Error occured while performing db query")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			if len(result) == 0 {
				c.AbortWithStatus(404)
			} else {
				b, err := json.Marshal(result)
				if err != nil {
					log.Error().Err(err).Msg("Error occured while marshalling data")
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				} else {
					cache.RedisSet(commons.GetRedisURI(), keyName, string(b), cacheExpire)
					c.JSON(http.StatusOK, result)
				}
			}
		}
	} else {
		result := make(map[string]interface{}, 0)
		if commons.SqlDriver == "mysql" {
			result, err = mysql.RawById(d)
		} else {
			result, err = postgres.RawById(d)
		}
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing db query")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if len(result) == 0 {
			c.AbortWithStatus(404)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
