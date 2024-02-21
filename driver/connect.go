package driver

// ConnectConfig 连接配置
type ConnectConfig struct {
	Protocol string //协议，一般都是tcp
	Host     string // 链接地址
	Port     int    // 端口
	DbName   string // 数据库名称
	User     string // 用户名
	Password string // 密码
	Charset  string // 字符集
}
