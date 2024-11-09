package bizerrors

import "git.dustess.com/shared-golib/duserr"

// 20030050 ~ 20030100
var (
	ErrFilterGroupNotFound   = duserr.NewBizError(20030050, "筛选组不存在")
	ErrFilterGroupNameEmpty  = duserr.NewBizError(20030051, "filter group name 不能为空")
	ErrFilterGroupIDEmpty    = duserr.NewBizError(20030052, "filter group id 不能为空")
	ErrFilterGroupNameRepeat = duserr.NewBizError(20030053, "filter group name 重复")
	ErrFilterGroupMaxLimit   = duserr.NewBizError(20030054, "filter group 最多允许20个")
	ErrFilterGroupDBFail     = duserr.NewBizError(20030055, "数据库操作失败")
)

var ErrFilterParmaErr = duserr.NewBizError(20030056, "%s")

func NewErrFilterBizUnknown(msg string) *duserr.BizError {
	return ErrFilterParmaErr.Params(msg)
}
