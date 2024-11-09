package bizerrors

import "git.dustess.com/shared-golib/duserr"

var (
	ErrHeaderNotFound          = duserr.NewBizError(20034004, "业务模块不存在！")
	ErrHeaderFieldKeyRepeat    = duserr.NewBizError(20034005, "fieldKey重复！")
	ErrHeaderFieldKeyNotFound  = duserr.NewBizError(20034006, "fieldKey不存在！")
	ErrHeaderFieldKeyEmpty     = duserr.NewBizError(20034007, "fieldKey不能为空！")
	ErrHeaderFieldNameEmpty    = duserr.NewBizError(20034008, "fieldName不能为空！")
	ErrHeaderBizModelEmpty     = duserr.NewBizError(20034009, "bizModel不能为空！")
	ErrHeaderAccountEmpty      = duserr.NewBizError(20034010, "account不能为空！")
	ErrHeaderItemsEmpty        = duserr.NewBizError(20034011, "items不能为空！")
	ErrHeaderAuthEmpty         = duserr.NewBizError(20034012, "非法account|user！")
	ErrHeaderAccountModel      = duserr.NewBizError(20034013, "未开启租户表头！")
	ErrHeaderBizModelRepeat    = duserr.NewBizError(20034014, "业务模块重复！")
	ErrHeaderFieldPermNotFound = duserr.NewBizError(20034015, "表头field_perm不存在！")
	ErrHeaderFieldTypeInvalid  = duserr.NewBizError(20034016, "field_type只能为system、custom")
	ErrHeaderParmaInvalid      = duserr.NewBizError(20034017, "非法的参数")
)
