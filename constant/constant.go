package constant

const (
	Success = 200
	// 数据库操作失败以100开头
	InsertTccDataErrCode = 100
	NotFoundErrCode      = 404

	//
	RequestTypeTry     = 0
	RequestTypeConfirm = 1
	RequestTypeCancel  = 2

	// 当前请求处理的状态：0：待处理（默认），1：提交成功，2：提交失败（需要继续提交），3：回滚成功，4：回滚失败（需要继续回滚），5：人工干预
	RequestInfoStatus_0 = 0
	RequestInfoStatus_1 = 1
	RequestInfoStatus_2 = 2
	RequestInfoStatus_3 = 3
	RequestInfoStatus_4 = 4
	RequestInfoStatus_5 = 4

	RetryTimes  = 5 // 重复多少次，不再重试，需要进行人工干预
	SendSuccess = 1 // 已发送邮件

)
