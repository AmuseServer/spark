package bizerrors

import (
	"git.dustess.com/shared-golib/duserr"
)

// 2003000 ~ 20030050
var (
	ErrProfileNotFound = duserr.NewBizError(20030000, "no profile found")
	ErrProfileGetFail  = duserr.NewBizError(20030001, "get profile fail")
	ErrProfileKeyEmpty = duserr.NewBizError(20030002, "profile key 不能为空")
	ErrProfileDBFail   = duserr.NewBizError(20030003, "数据库操作失败")
)
