package registry

type Registration struct {
	ServiceName ServiceName //服务名称
	ServiceURL  string      //URL
}
type ServiceName string

const (
	LogService     = ServiceName("LogService") //存在的服务形成常量
	GradingService = ServiceName("GradingService")
)
