package main

import (
	"flag"
	"log"
	"opsoa_plug/config/global"
)

type Plug struct {
	ID uint32
}

func main() {
	// 初始化
	flag.IntVar(&global.Port, "port", 0, "source file")
	flag.StringVar(&global.Key, "key", "", "source file")
	//
	if global.Port != 0 && global.Key != "" {
		//
	} else {
		log.Fatal(global.SystemFlagFail)
	}
}
