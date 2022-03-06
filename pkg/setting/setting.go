package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// 1.  编写与配置项保持一致的结构体（App、Server、Database）
// 2. 使用 MapTo 将配置项映射到结构体上 具体结构体的类型是怎么样子的需要看你自己的使用场景而定
// 3. 对一些需特殊设置的配置项进行再赋值

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	PrefixUrl      string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string
	FontSavePath   string

	ExportSavePath string
	QrCodeSavePath string
	LogSavePath    string
	LogSaveName    string
	LogFileExt     string
	TimeFormat     string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

// 开始执行 主要的文档依据就是 go-ini 这个包
func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second // 对于time类型的需要 * 一个量 完成类型转化
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}
