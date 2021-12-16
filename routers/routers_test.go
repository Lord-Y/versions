package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/Lord-Y/versions-api/cache"
	"github.com/Lord-Y/versions-api/commons"
	"github.com/Lord-Y/versions-api/models"
	"github.com/Lord-Y/versions-api/mysql"
	"github.com/Lord-Y/versions-api/postgres"
	"github.com/alecthomas/assert"
	"github.com/icrowley/fake"
	"github.com/rs/zerolog/log"
)

func performRequest(r http.Handler, headers map[string]string, method string, url string, payload string) (z *httptest.ResponseRecorder, err error) {
	var (
		req *http.Request
	)

	switch method {
	case "GET":
		req, err = http.NewRequest(method, url, nil)
	case "POST":
		req, err = http.NewRequest(method, url, strings.NewReader(payload))
	default:
		if payload == "" {
			req, err = http.NewRequest(method, url, nil)
		} else {
			req, err = http.NewRequest(method, url, strings.NewReader(payload))
		}
	}
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while initalising http request")
		return nil, err
	}
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w, nil
}

func TestPing(t *testing.T) {
	if commons.SqlDriver == "mysql" {
		if !mysql.Ping() {
			log.Error().Msg("Fail to ping database instance")
			t.Fail()
		}
	} else {
		if !postgres.Ping() {
			log.Error().Msg("Fail to ping database instance")
			t.Fail()
		}
	}
}

func TestHealth(t *testing.T) {
	headers := make(map[string]string)

	assert := assert.New(t)
	router := SetupRouter()
	w, err := performRequest(router, headers, "GET", "/api/v1/versions/health", "")
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing http request")
		t.Fail()
		return
	}

	assert.Contains(w.Body.String(), "OK", "Fail to get /health body")
	assert.Equal(200, w.Code, "Fail to get /health")
}

func TestHealthz(t *testing.T) {
	headers := make(map[string]string)

	assert := assert.New(t)
	router := SetupRouter()
	w, err := performRequest(router, headers, "GET", "/api/v1/versions/healthz", "")
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing http request")
		t.Fail()
		return
	}

	assert.Contains(w.Body.String(), "status", "Fail to get /healthz body")
	assert.Equal(200, w.Code, "Fail to get /healthz")
}

func createFakeVerion() (s string) {
	return fmt.Sprintf("%d.%d.%d", commons.RandomInt(1, 5), commons.RandomInt(1, 25), commons.RandomInt(1, 20))
}
func TestCreate(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	type result struct {
		VersionId int64
	}
	statusList := []string{
		"deployed",
		"failed",
	}

	assert := assert.New(t)
	router := SetupRouter()
	payload := fmt.Sprintf("workload=%s", fake.CharactersN(3))
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			payload += fmt.Sprintf("&platform=%s", fake.CharactersN(10))
			payload += fmt.Sprintf("&environment=%s", fake.CharactersN(10))
		} else {
			payload = fmt.Sprintf("workload=%s", fake.CharactersN(3))
			payload += fmt.Sprintf("&platform=%s", fake.CharactersN(10))
			payload += fmt.Sprintf("&environment=%s", fake.CharactersN(10))
		}
		payload += fmt.Sprintf("&version=%s", createFakeVerion())
		payload += fmt.Sprintf("&changelogURL=http://www.%s/changelog", strings.ToLower(fake.DomainName()))
		payload += "&raw={'a': 'b'}"
		payload += "&status=ongoing"

		w, err := performRequest(router, headers, "POST", "/api/v1/versions/create", payload)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing http request")
			t.Fail()
			return
		}

		assert.Contains(w.Body.String(), "versionId", fmt.Sprintf("Fail to create new version during test number %d with payload %s", i, payload))
		assert.Equal(201, w.Code, fmt.Sprintf("Fail to get right http status code during test number %d", i))
		r := regexp.MustCompile(`&version=.*\..*\..*&`)
		payload = r.ReplaceAllString(payload, "&")

		log.Info().Msg("Let's update status")
		if w.Code == 201 {
			r := result{}
			err = json.Unmarshal(w.Body.Bytes(), &r)
			if err != nil {
				log.Error().Err(err).Msg("Error occured while unmarshalling data")
				t.Fail()
				return
			}
			payloadStatus := fmt.Sprintf("versionId=%d&status=%s", r.VersionId, commons.RandomValueFromArray(statusList))

			wu, err := performRequest(router, headers, "POST", "/api/v1/versions/update/status", payloadStatus)
			if err != nil {
				log.Error().Err(err).Msg("Error occured while performing http request")
				t.Fail()
				return
			}
			assert.Equal(200, wu.Code, fmt.Sprintf("Fail to update deployment versionId %d during test number %d", r.VersionId, i))
		}
	}
	if commons.RedisEnabled() {
		cache.RedisFlushDB(commons.GetRedisURI())
	}
}

func TestReadEnvironment_200(t *testing.T) {
	headers := make(map[string]string)
	var (
		result models.DBReadForUnitTesting
		err    error
	)

	assert := assert.New(t)
	router := SetupRouter()
	if commons.SqlDriver == "mysql" {
		result, err = mysql.ReadForUnitTesting("deployed")
		if err != nil {
			log.Error().Msg("Result from DB is empty")
			t.Fail()
			return
		}
	} else {
		result, err = postgres.ReadForUnitTesting("deployed")
		if err != nil {
			log.Error().Msg("Result from DB is empty")
			t.Fail()
			return
		}
	}
	if result != (models.DBReadForUnitTesting{}) && result.Workload != "" && result.Platform != "" {
		url := fmt.Sprintf("/api/v1/versions/read/environment?workload=%s&platform=%s&environment=%s", result.Workload, result.Platform, result.Environment)
		w, err := performRequest(router, headers, "GET", url, "")
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing http request")
			t.Fail()
			return
		}
		assert.Equal(200, w.Code, "Fail to get expected status code")
		assert.Contains(w.Body.String(), "workload", "Fail to get expected content")

	}
}

func TestReadEnvironment_400(t *testing.T) {
	headers := make(map[string]string)

	assert := assert.New(t)
	router := SetupRouter()
	w, err := performRequest(router, headers, "GET", "/api/v1/versions/read/environment?workload=plop", "")
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing http request")
		t.Fail()
		return
	}

	assert.Equal(400, w.Code, "Fail to get expected status code")
}

func TestReadEnvironment_404(t *testing.T) {
	headers := make(map[string]string)

	assert := assert.New(t)
	router := SetupRouter()
	w, err := performRequest(router, headers, "GET", "/api/v1/versions/read/environment?workload=plop&platform=platform&environment=environment", "")
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing http request")
		t.Fail()
		return
	}

	assert.Equal(404, w.Code, "Fail to get expected status code")
}

func TestReadEnvironmentLatest_200(t *testing.T) {
	var (
		headers map[string]string
		result  models.DBReadForUnitTesting
		err     error
	)
	assert := assert.New(t)
	router := SetupRouter()
	if commons.SqlDriver == "mysql" {
		result, err = mysql.ReadForUnitTesting("deployed")
		if err != nil {
			log.Error().Msg("Result from DB is empty")
			t.Fail()
			return
		}
	} else {
		result, err = postgres.ReadForUnitTesting("deployed")
		if err != nil {
			log.Error().Msg("Result from DB is empty")
			t.Fail()
			return
		}
	}
	if result != (models.DBReadForUnitTesting{}) && result.Workload != "" && result.Platform != "" {
		url := fmt.Sprintf("/api/v1/versions/read/environment/latest?workload=%s&platform=%s&environment=%s", result.Workload, result.Platform, result.Environment)
		w, err := performRequest(router, headers, "GET", url, "")
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing http request")
			t.Fail()
			return
		}
		assert.Equal(200, w.Code, "Fail to get expected status code")
		assert.Contains(w.Body.String(), "version", "Fail to get expected content")

	}
}

func TestReadEnvironmentLatestWhatever_200(t *testing.T) {
	var (
		headers map[string]string
		result  models.DBReadForUnitTesting
		err     error
	)
	assert := assert.New(t)
	router := SetupRouter()
	if commons.SqlDriver == "mysql" {
		result, err = mysql.ReadForUnitTesting("failed")
		if err != nil {
			log.Error().Msg("Result from DB is empty")
			t.Fail()
			return
		}
	} else {
		result, err = postgres.ReadForUnitTesting("failed")
		if err != nil {
			log.Error().Msg("Result from DB is empty")
			t.Fail()
			return
		}
	}
	if result != (models.DBReadForUnitTesting{}) && result.Workload != "" && result.Platform != "" {
		url := fmt.Sprintf("/api/v1/versions/read/environment/latest/whatever?workload=%s&platform=%s&environment=%s", result.Workload, result.Platform, result.Environment)
		w, err := performRequest(router, headers, "GET", url, "")
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing http request")
			t.Fail()
			return
		}
		assert.Equal(200, w.Code, "Fail to get expected status code")
		assert.Contains(w.Body.String(), "version", "Fail to get expected content")

	}
}

func TestReadEnvironmentLatest_404(t *testing.T) {
	headers := make(map[string]string)

	assert := assert.New(t)
	router := SetupRouter()
	w, err := performRequest(router, headers, "GET", "/api/v1/versions/read/environment?workload=plop&platform=platform&environment=environment", "")
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing http request")
		t.Fail()
		return
	}

	assert.Equal(404, w.Code, "Fail to get expected status code")
}

func TestReadPlatform_200(t *testing.T) {
	headers := make(map[string]string)
	var (
		result models.DBReadForUnitTesting
		err    error
	)

	assert := assert.New(t)
	router := SetupRouter()
	if commons.SqlDriver == "mysql" {
		result, err = mysql.ReadForUnitTesting("deployed")
		if err != nil {
			log.Error().Msg("Result from DB is empty")
			t.Fail()
			return
		}
	} else {
		result, err = postgres.ReadForUnitTesting("deployed")
		if err != nil {
			log.Error().Msg("Result from DB is empty")
			t.Fail()
			return
		}
	}
	if result != (models.DBReadForUnitTesting{}) && result.Workload != "" && result.Platform != "" {
		url := fmt.Sprintf("/api/v1/versions/read/platform?workload=%s&platform=%s", result.Workload, result.Platform)
		w, err := performRequest(router, headers, "GET", url, "")
		if err != nil {
			log.Error().Err(err).Msg("Error occured while performing http request")
			t.Fail()
			return
		}
		assert.Equal(200, w.Code, "Fail to get expected status code")
		assert.Contains(w.Body.String(), "workload", "Fail to get expected content")

	}
}

func TestReadPlatform_400(t *testing.T) {
	headers := make(map[string]string)

	assert := assert.New(t)
	router := SetupRouter()
	w, err := performRequest(router, headers, "GET", "/api/v1/versions/read/platform?workload=plop", "")
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing http request")
		t.Fail()
		return
	}

	assert.Equal(400, w.Code, "Fail to get expected status code")
}

func TestReadPlatform_404(t *testing.T) {
	headers := make(map[string]string)

	assert := assert.New(t)
	router := SetupRouter()
	w, err := performRequest(router, headers, "GET", "/api/v1/versions/read/platform?workload=plop&platform=platform", "")
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing http request")
		t.Fail()
		return
	}

	assert.Equal(404, w.Code, "Fail to get expected status code")
}

func TestReadDistinctWorkloads(t *testing.T) {
	headers := make(map[string]string)

	assert := assert.New(t)
	router := SetupRouter()
	w, err := performRequest(router, headers, "GET", "/api/v1/versions/read/distinct/workloads", "")
	if err != nil {
		log.Error().Err(err).Msg("Error occured while performing http request")
		t.Fail()
		return
	}
	assert.Equal(200, w.Code, "Fail to get expected status code")
	assert.Contains(w.Body.String(), "workload", "Fail to get expected content")
}
