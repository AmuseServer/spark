package constants

import (
	"strconv"
	"strings"
)

type Switch string

const (
	CompanyManagement Switch = "CompanyManagement" // 大B开关
	FriendManagement  Switch = "FriendManagement"  // 好友管理开关
)

func (s Switch) String() string {
	return string(s)
}

const FeatureFlagPrefix = "FeatureFlags-"

func (s Switch) FeatureFlagCode() int32 {
	if !strings.HasPrefix(s.String(), FeatureFlagPrefix) {
		return -1
	}
	result := strings.Split(s.String(), "-")
	if len(result) != 2 {
		return -1
	}
	codeStr := result[1]
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		return -1
	}
	return int32(code)
}
