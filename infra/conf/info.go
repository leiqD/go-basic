package conf

type LogInfo struct {
	Level          string
	Path           string
	MaxSize        int
	MaxBackupNum   int
	MackupDuration int
}
type SqlParam struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

type DataStoreInfo struct {
	Host   string
	Port   string
	Db     string
	User   string
	Paswd  string
	Net    string
	Params SqlParam
}
type NetInfo struct {
	Host string
	Port string
}

type Viper struct {
	Log       LogInfo
	DataStore DataStoreInfo
	Net       NetInfo
}
