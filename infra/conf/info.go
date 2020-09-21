package conf

type Log struct {
	level string
	path  string
}

type Sql struct {
	host  string
	port  string
	db    string
	user  string
	paswd string
}
type Net struct {
	host string
	port string
}

type Viper struct {
	log Log
	sql Sql
	net Net
}
