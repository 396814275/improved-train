package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局变量，用来保存程序的所有配置信息
var Conf = new(MultipleConfig)

type MultipleConfig struct {
	*Appconfig   `mapstructure:"app"`
	*Logconfig   `mapstructure:"log"`
	*Mysqlconfig `mapstructure:"mysql"`
	*Redisconfig `mapstructure:"redis"`
}

type Appconfig struct {
	Name      string `mapstructure:"name"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Gin_mode  string `mapstructure:"gin_mode"`
}
type Logconfig struct {
	Level      string `mapstructure:"level"`
	Mode       string `mapstructure:"mode"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type Mysqlconfig struct {
	Host         string `mapstructure:"sql_host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
}
type Redisconfig struct {
	Host     string `mapstructure:"redis_host"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	Poolsize int    `mapstructure:"poolsize"`
}

func Init(filePath string) (err error) {
	//方式1;直接指定配置文件路径（相对路劲或绝对路径）
	//viper.SetConfigFile("config.json")
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./conf/")
	//方式2：制定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	//配置文件名不需要带后缀
	//配置文件位置可配置多个
	//viper.SetConfigName("config") //制定配置文件名（不带后缀）
	//viper.AddConfigPath("./conf/")
	//viper.AddConfigPath(".")
	//基本是配合远程配置中心使用的,告诉viper当前的数据使用什么格式去解析
	//viper.SetConfigType("yaml")

	viper.SetConfigFile(filePath)

	//读取配置文件
	err = viper.ReadInConfig()
	//读取配置信息失败
	if err != nil {
		//读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed,error:%v\n", err)
	}
	//把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal failed,err:", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件修改了：%s\n", e.Name)
		if err := viper.Unmarshal(&Conf); err != nil {
			fmt.Println("viper.Unmarshal failed,err:", err)
		}
	})

	return
}
