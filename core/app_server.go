package core

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/onedss/EasyGoLib/utils"
	"github.com/onedss/onedarwin/rtsp"
	"github.com/onedss/service"
	"log"
	"os"
)

func StartApp() {
	log.SetPrefix("[EasyDarwin] ")
	log.SetFlags(log.LstdFlags)
	if utils.Debug {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	}
	sec := utils.Conf().Section("service")
	svcConfig := &service.Config{
		Name:        sec.Key("name").MustString("EasyDarwin_Service"),
		DisplayName: sec.Key("display_name").MustString("EasyDarwin_Service"),
		Description: sec.Key("description").MustString("EasyDarwin_Service"),
	}

	httpPort := utils.Conf().Section("http").Key("port").MustInt(10008)
	rtspServer := rtsp.GetServer()
	p := &application{
		httpPort:   httpPort,
		rtspPort:   rtspServer.TCPPort,
		rtspServer: rtspServer,
	}
	var s, err = service.New(p, svcConfig)
	if err != nil {
		log.Println(err)
		utils.PauseExit()
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "install" || os.Args[1] == "stop" {
			figure.NewFigure("EasyDarwin", "", false).Print()
		}
		log.Println(svcConfig.Name, os.Args[1], "...")
		if err = service.Control(s, os.Args[1]); err != nil {
			log.Println(err)
			utils.PauseExit()
		}
		log.Println(svcConfig.Name, os.Args[1], "ok")
		return
	}
	figure.NewFigure("EasyDarwin", "", false).Print()
	if err = s.Run(); err != nil {
		log.Println(err)
		utils.PauseExit()
	}
}
