package bizerrors

import "git.dustess.com/shared-golib/duserr"

// 20030100 ~ 20030110

var (
	ErrDataTemplateDBFail = duserr.NewBizError(20030101, "数据库操作失败")
)
