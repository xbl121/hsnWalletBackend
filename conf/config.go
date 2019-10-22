package conf

import "time"

type Config struct {
	Service Service
	Db  Db
	JiG JiG
}

type Db struct {
	DbString string
	DbName   string
}
type JiG struct {
	AppKey       string       // 911af8bb6b5798c72d42ce20
	MasterSecret string // c130ecb2403fec2514f879ff
	AuthString   string
}

type Service struct {
	IP           string
	Port         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}
func NewConfig() Config {
	return Config{
		Service:Service{
			IP:           "127.0.0.1",
			Port:         "7780",
			WriteTimeout: 15,
			ReadTimeout:  15,
		},
		Db: Db{
			DbString: "127.0.0.1:27017/walletBackend",
			DbName:   "walletBackend",
		},
		JiG: JiG{
			AppKey:       "911af8bb6b5798c72d42ce20",
			MasterSecret: "c130ecb2403fec2514f879ff",
			AuthString:   "911af8bb6b5798c72d42ce20:c130ecb2403fec2514f879ff",
		},
	}
}
