package pkg

type ExecMini struct {
	ID      uint32 `json:"id" gorm:"comment:执行id"`
	Uuid    string `json:"uuid" gorm:"comment:独占客户端id,允许的最大公众执行数 0不限,分片区域"`
	RunType uint32 `json:"run_type" gorm:"comment:运行状态: 1:未运行, 2:分配中, 3:执行中"`
}

// Cache job_{group}_{exec_type}_{id}
// exec_type 1:每天固定时间, 2:每小时固定时间, 3:每分种固定时间, 4:每周固定时间, 5:每月固定时间, 7:间隔n秒
type Exec struct {
	Content
	ID          uint32    `json:"id" gorm:"comment:执行id"`
	Cron        string    `json:"cron" gorm:"comment:Cron表达式"`
	Uuid        string    `json:"uuid" gorm:"comment:独占客户端id,允许的最大公众执行数 0不限,分片区域"`
	Group       uint32    `json:"group" gorm:"comment:执行客户端群组id"`
	RunType     uint32    `json:"run_type" gorm:"comment:运行状态: 1:未运行, 2:分配中, 3:执行中"`
	Mode        uint32    `json:"mode" gorm:"comment:执行方式: 1:独占执行, 2:公众执行, 3:分片执行, 4:执行后移除此任务"`
	Type        uint32    `json:"type" gorm:"comment:执行类型 1:Shell, 2:Bean, 3:Http, 4: Protobuf, 5: WebSocket"`
	UpdateTime  uint64    `json:"update_time" gorm:"comment:后台修改本配置的时间"`
	MaxTime     uint32    `json:"max_time" gorm:"comment:最长执行时间"`
	MaxRefresh  uint32    `json:"max_refresh" gorm:"default:1;comment:最大重试次数"`
	StartTime   uint32    `json:"sun_time" gorm:"default:0;comment:开始执行时间"`
	RepeatTimes uint32    `json:"repeat_times" gorm:"default:0;comment:当前重试次数"`
	LogType     uint32    `json:"log_type" gorm:"default:1;comment:原始日志类型 1:丢弃, 2:储存到文件, 3:储存到数据库"`
	TaskStart   uint64    `json:"task_start" gorm:"default:0;comment:任务开始执行时间"`
	TaskStop    uint64    `json:"task_stop" gorm:"default:0;comment:任务停止不再执行时间"`
	LogList     []LogInfo `json:"log_list" gorm:"comment:原始日志类型 日志详情"`
}

type LogInfo struct {
	Title string `json:"title" gorm:"comment:日志标题"`
	Value string `json:"value" gorm:"comment:日志详情"`
	Type  uint32 `json:"type" gorm:"comment:日志状态 1:初始化, 2:开始, 3:执行中, 4:执行成功, 5:报错且到达最大次数, 6:超时, 7:执行错误"`
	Time  uint64 `json:"time" gorm:"comment:日志时间"`
}

// Cache Task Info
type EtcdExec struct {
	Content
	ID         uint32 `json:"id" gorm:"comment:执行id"`
	Cron       string `json:"cron" gorm:"comment:Cron表达式"`
	Uuid       string `json:"uuid" gorm:"comment:独占客户端id,允许的最大公众执行数 0不限,分片区域"`
	Group      uint32 `json:"group" gorm:"comment:执行客户端群组id"`
	RunType    uint32 `json:"run_type" gorm:"comment:运行状态: 1:未运行, 2:分配中, 3:执行中"`
	Mode       uint32 `json:"mode" gorm:"comment:执行方式: 1:独占执行, 2:公众执行, 3:分片执行, 4:执行后移除此任务"`
	Type       uint32 `json:"type" gorm:"comment:执行类型 1:Shell, 2:Bean, 3:Http, 4: Protobuf, 5: WebSocket"`
	TaskStart  uint64 `json:"task_start" gorm:"default:0;comment:任务开始执行时间"`
	TaskStop   uint64 `json:"task_stop" gorm:"default:0;comment:任务停止不再执行时间"`
	UpdateTime int64  `json:"update_time" gorm:"comment:后台修改本配置的时间"`
	MaxTime    uint32 `json:"max_time" gorm:"comment:最长执行时间"`
	MaxRefresh uint32 `json:"max_refresh" gorm:"default:1;comment:最大重试次数"`
	LogType    uint32 `json:"log_type" gorm:"default:1;comment:原始日志类型 1:丢弃, 2:储存到文件, 3:储存到数据库"`
}

type Content struct {
	Plug struct {
		Type      uint32   `json:"type" gorm:"comment: 1:执行二进制统一插件, 2:执行go统一插件"`
		Program   string   `json:"program" gorm:"comment:go、php等程序所在目录"`
		User      string   `json:"user" gorm:"comment:执行用户名 (Linux专用)"`
		Path      string   `json:"path" gorm:"comment:执行路径"`
		File      string   `json:"file" gorm:"comment:执行文件"`
		Parameter string   `json:"parameter" gorm:"comment:参数"`
		EnvList   []string `json:"env_list" gorm:"comment:执行环境列表"`
	} `json:"plug" gorm:"comment:执行插件命令"`
	Shell struct {
		User    string   `json:"user" gorm:"comment:执行用户名 (Linux专用)"`
		Path    string   `json:"path" gorm:"comment:执行路径"`
		Command string   `json:"command" gorm:"comment:执行命令"`
		EnvList []string `json:"env_list" gorm:"comment:执行环境列表"`
	} `json:"shell" gorm:"comment:执行shell命令"`
	Http struct {
		Url  string            `json:"url" gorm:"comment:访问http地址"`
		Type string            `json:"type" gorm:"comment:请求类型 POST, GET, Head"`
		Head map[string]string `json:"head" gorm:"comment:请求头"`
		Form string            `json:"form" gorm:"comment:提交的表单信息 a=1&b=2"`
	}
	WebSocket struct {
		Host string              `json:"host" gorm:"comment:服务端地址"`
		Text string              `json:"form" gorm:"comment:提交的信息"`
		Head map[string][]string `json:"head" gorm:"comment:设置头部信息"`
	}
	Protobuf struct {
		Type   string `json:"type" gorm:"comment:1:Http2, 2:tpc"`
		Host   string `json:"host" gorm:"comment:服务端ip/域名:端口"`
		Struct string `json:"struct" gorm:"comment:提交的数据JSON形式"`
	}
}
