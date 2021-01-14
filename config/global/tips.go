package global

// 系统描述
const (
	SystemFlagFail = "无法检测到正确的传入参数,请确认执行方式是否正确"
)

// 逻辑判断
const (
	LogicNoArrayGather   = "变量不是切片、数组、集合"
	LogicNoStructPointer = "变量不是结构指针"
)

// Task                 任务相关
const (
	TaskTypeInit          = "任务初始化"
	TaskNotExist          = "任务不存在"
	TaskTypeStart         = "任务开始执行"
	TaskTypeInitValue     = "任务开始载入,并进行环境初始化"
	TaskMatchingFailed    = "任务匹配失败"
	TaskExecutionFailed   = "任务执行失败"
	TaskMissionCompleted  = "任务执行完毕"
	TaskReachedMaxRetries = "任务已经到了最大重试次数"
	TaskReachedMaxTime    = "任务已经到了最大执行时长"
	TaskModifySkipForNow  = "发现任务更新，跳过本轮执行"
)
