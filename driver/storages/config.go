package storages

import (
	"os"
	"strconv"
)

func LoadConfig() Config {

	redisDB := getEnv("KELPI_REDIS_DB", "0")
	redisDBint, _ := strconv.Atoi(redisDB)

	ret := Config{
		RootDomain:    getEnv("KELPI_ROOT_DOMAIN", "com."),
		RedisHost:     getEnv("KELPI_REDIS_HOST", "localhost:6379"),
		RedisPassword: getEnv("KELPI_REDIS_PASSWORD", "qwerty"),
		RedisDB:       redisDBint,
		SoaMname:      getEnv("KELPI_SOA_MNAME	", "kelpi.example.com."),
		SoaRname:      getEnv("KELPI_SOA_RNAME	", "hostmaster.kelpi.example.com."),
		SoaSerial:     getEnv("KELPI_SOA_SERIAL", "1"),
		SoaRefresh:    getEnv("KELPI_SOA_REFRESH", "3600"),
		SoaRetry:      getEnv("KELPI_SOA_RETRY	", "600"),
		SoaExpire:     getEnv("KELPI_SOA_EXPIRE", "86400"),
		SoaMinimum:    getEnv("KELPI_SOA_MINIMUM", "1"),
		SoaTTL:        getEnv("KELPI_SOA_TTL", "86400"),
	}

	return ret
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
