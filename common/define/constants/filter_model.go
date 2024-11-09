package constants

const (
	AllCusToFollowUp = 2  // 全部联系人待跟进
	AllCusRecycling  = 3  // 全部联系人即将回收
	AllCusInFollowUp = 4  // 全部联系人跟进中
	AllCus           = 5  // 全部联系人
	MyCusToFollowUp  = 6  // 我的联系人待跟进
	MyCusRecycling   = 7  // 我的联系人即将回收
	MyCusFocus       = 8  // 我的联系人重点关注
	MyCusInFollowUp  = 9  // 我的联系人跟进中
	MyCus            = 10 // 我的联系人

	AllCorpFollowUp   = 11 // 全部企业待跟进
	AllCorpRecycling  = 12 // 全部企业即将回收
	AllCorpInFollowUp = 13 // 全部企业跟进中
	AllCorp           = 14 // 全部企业
	MyCorpToFollowUp  = 15 // 我的企业待跟进
	MyCorpRecycling   = 16 // 我的企业即将回收
	MyCorpInFollowUp  = 17 // 我的企业跟进中
	MyCorp            = 18 // 我的企业
)

var (
	CorpFilterMap = map[int]bool{
		AllCorpFollowUp:   true,
		AllCorpRecycling:  true,
		AllCorpInFollowUp: true,
		AllCorp:           true,
		MyCorpToFollowUp:  true,
		MyCorpRecycling:   true,
		MyCorpInFollowUp:  true,
		MyCorp:            true,
	}
)
