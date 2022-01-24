package model

type Zap struct {
	Level         string //级别
	Format        string //输出
	Prefix        string //日志前缀
	Director      string //日志文件
	ShowLine      bool   //显示行
	EncodeLevel   string //编码级别
	StackTraceKey string //栈名
	LogInConsole  bool   //输出控制台
}
