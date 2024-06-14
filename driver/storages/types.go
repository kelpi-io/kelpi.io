package storages

type Config struct {
	RootDomain    string
	RedisHost     string
	RedisPassword string
	RedisDB       int64
	SoaMname      string
	SoaRname      string
	SoaSerial     string
	SoaRefresh    string
	SoaRetry      string
	SoaExpire     string
	SoaMinimum    string
	SoaTTL        string
}
