package configs

// 服务配置 防止变量污染故用函数组织
func GetServerConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)

	serverConfig["HOST"] = "127.0.0.1"            //监听地址
	serverConfig["PORT"] = "9090"               //监听端口
	serverConfig["VIEWS_PATTERN"] = "static/*" //模板路径pattern
	serverConfig["ENV"] = "debug"               // 环境模式 release/debug/test
	return
}