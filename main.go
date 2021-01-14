package main

import (
	"errors"
	"flag"
	"opsoa_plug/config/global"
	"opsoa_plug/pkg"
)

type Plug struct {
	ID uint32
}

// @Tags Start
// @Summary 开始执行任务
// @Security Start
// @Success error
func (p Plug) Start(value string) error {
	// 初始化
	err := errors.New(global.SystemFlagFail)
	flag.IntVar(&global.ID, "id", 0, "source file")
	flag.IntVar(&global.Port, "port", 0, "source file")
	flag.StringVar(&global.Key, "key", "", "source file")
	if global.Port != 0 && global.Key != "" {
		// 判断是否链接到客户端
		_, err = pkg.AddLabel(2, global.TaskTypeStart, value)
	}
	return err
}

func (p Plug) Set(label string, value string) error {
	_, err := pkg.AddLabel(4, label, value)
	return err
}
