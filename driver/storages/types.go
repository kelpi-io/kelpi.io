package storages

type Config struct {
	RootDomain    string
	RedisHost     string
	RedisPassword string
	RedisDB       int
	SoaMname      string
	SoaRname      string
	SoaSerial     string
	SoaRefresh    string
	SoaRetry      string
	SoaExpire     string
	SoaMinimum    string
	SoaTTL        string
}
