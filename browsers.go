package browsers

import "context"

// Intent 操作参数
type Intent struct {
	Context context.Context
	Action  string
	URL     string
	Args    []string
}

// Browser 浏览器代理
type Browser interface {
	Accept(i *Intent) bool
	Open(i *Intent) (Task, error)
}

// Registration 浏览器注册信息
type Registration struct {
	Enabled  bool
	Priority int
	Browser  Browser
}

// Registry 浏览器注册接口
type Registry interface {
	Registration() *Registration
}

// Service 浏览器服务
type Service interface {
	Open(i *Intent) (Task, error)
	Start(i *Intent) (Task, error)
	Run(i *Intent) error
}

// Task 表示启动浏览器的任务
type Task interface {
	Start() error
	Run() error
	Join() error
}
