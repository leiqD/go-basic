package conf

type Log struct {
	Level          string
	Path           string
	MaxSize        int
	MaxBackupNum   int
	MackupDuration int
}

type Sql struct {
	Host  string
	Port  string
	Db    string
	User  string
	Paswd string
}
type Net struct {
	Host string
	Port string
}

type Viper struct {
	Log Log
	Sql Sql
	Net Net
}
