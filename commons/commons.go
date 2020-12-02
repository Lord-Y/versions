package commons

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	SqlDriver = strings.TrimSpace(os.Getenv("SQL_DRIVER"))
)

// BuildDSN stand to create string url to connect to the sql instance
func BuildDSN() string {
	return strings.TrimSpace(os.Getenv("DB_URI"))
}

func RedisEnabled() (b bool) {
	b, _ = strconv.ParseBool(strings.TrimSpace(os.Getenv("REDIS_ENABLED")))
	return
}

func GetRedisURI() string {
	return strings.TrimSpace(os.Getenv("REDIS_URI"))
}

// RandomInt function is explicit
func RandomInt(min int, max int) int {
	if min < 1 {
		min = 1
	}
	if max < 1 {
		max = 1
	}
	if min == 1 && max == 1 {
		return 1
	}
	if min > max {
		return 1
	}
	// rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(max-min)
}

// GetPagination function is explicit
func GetPagination(page int, start int, end int, rangeLimit int) (startLimit int, EndLimit int) {
	if page <= 0 {
		page = 1
	}
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 1
	}
	if page == 1 {
		start = 0
	}
	if rangeLimit < 1 {
		rangeLimit = 1
	}
	switch page {
	case 1:
		return start, end
	default:
		return rangeLimit * (page - 1), rangeLimit
	}
}

// GetMD5HashWithSum function is explicit
func GetMD5HashWithSum(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// RandomValueFromArray function is explicit
func RandomValueFromArray(array []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]string, len(array))
	for i, v := range r.Perm(len(array)) {
		n[i] = array[v]
	}
	return n[0]
}
