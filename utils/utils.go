package utils

import (
	"fmt"
	"os"
	"time"
)

const (
	ReleaseEnv = "onl"
)

// GetDate 返回当前时间
func GetDate() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}

// GetRunTime 获取当前系统环境
func GetRunTime() string {
	RunTime := os.Getenv("RUN_TIME")
	if RunTime == "" {
		fmt.Println("No RUN_TIME Can't start")
	}
	return RunTime
}

func IsRelease() bool {
	return GetRunTime() == ReleaseEnv
}

// Either 返回一个存在的字符串
func Either(list ...string) string {
	for _, v := range list {
		if v != "" {
			return v
		}
	}
	return ""
}
