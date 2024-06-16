package storages

import (
	"os"
	"strconv"
)

var ConfigDriver = Config{}

func (cfg *Config) LoadConfig() {

	redisDB := getEnv("KELPI_REDIS_DB", "0")
	redisDBint, _ := strconv.Atoi(redisDB)

	cfg.RootDomain = getEnv("KELPI_ROOT_DOMAIN", "com.")
	cfg.RedisHost = getEnv("KELPI_REDIS_HOST", "localhost:6379")
	cfg.RedisPassword = getEnv("KELPI_REDIS_PASSWORD", "qwerty")
	cfg.RedisDB = redisDBint
	cfg.SoaMname = getEnv("KELPI_SOA_MNAME", "kelpi.example.com.")
	cfg.SoaRname = getEnv("KELPI_SOA_RNAME", "hostmaster.kelpi.example.com.")
	cfg.SoaSerial = getEnv("KELPI_SOA_SERIAL", "1")
	cfg.SoaRefresh = getEnv("KELPI_SOA_REFRESH", "3600")
	cfg.SoaRetry = getEnv("KELPI_SOA_RETRY", "600")
	cfg.SoaExpire = getEnv("KELPI_SOA_EXPIRE", "86400")
	cfg.SoaMinimum = getEnv("KELPI_SOA_MINIMUM", "1")
	cfg.SoaTTL = getEnv("KELPI_SOA_TTL", "86400")

}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
