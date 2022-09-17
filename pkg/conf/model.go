package conf

type system struct {
	Listen        []string `validate:"required"`
	Debug         bool
	SessionSecret string
}

type database struct {
	Type        string `validate:"required"`
	Host        string
	Port        int
	Username    string
	Password    string
	Database    string
	Charset     string
	DBFile      string
	TablePrefix string
}

type redis struct {
	Network  string
	Server   string
	Password string
	DB       int
}

type cors struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowCredentials bool
	MaxAge           int
	AllowHeaders     []string
	ExposeHeaders    []string
}
