package conf

type Config struct {
	Db  Db  `json:"db"`
	JiG JiG `json:"ji_g"`
}

type Db struct {
	DbString string `json:"db_string"`
	DbName   string `json:"db_name"`
}
type JiG struct {
	AppKey       string `json:"app_key"`       // 911af8bb6b5798c72d42ce20
	MasterSecret string `json:"master_secret"` // c130ecb2403fec2514f879ff
	AuthString   string `json:"auth_string"`
}

func NewConfig() Config {
	return Config{
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
