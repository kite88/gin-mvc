package conf

const (
	AppName  string = "gin_mvc"
	HttpPort string = "1996"
	RunMode  string = "debug" //debug test release
	//mysql
	DbHost        string = "localhost"
	DbPort        string = "3306"
	DbName        string = "gin_mvc"
	DbUser        string = "root"
	DbPassword    string = "root"
	DbTablePrefix string = "go_"
	//redis
	RedisAddr     string = "localhost:6379"
	RedisPassword string = ""
	RedisDb       int    = 1 //默认数据库
	RedisPrefix   string = AppName + "."
	//JWT
	JwtSecret string = "ICGF7AWJaeSMuhnbZVMpsDjIAypSwq6b"
	//token超时时间(秒)
	TokenExp int64 = 60 * 60 * 24 * 7
	//时间格式
	TimeFormatLayout = "2006-01-02 15:04:05"
)
