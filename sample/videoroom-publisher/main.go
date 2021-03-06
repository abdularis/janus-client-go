package main

import (
	"github.com/abdularis/janus-client-go/janus"
	"github.com/abdularis/janus-client-go/janus/videoroom"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	janus.SetVerboseRequestResponse(true)
	var roomID int64 = 7555683579550993055

	gateway, err := janus.Connect("ws://localhost:8188")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	_, _ = gateway.Info()

	session, err := gateway.Create()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	handle, err := session.Attach(videoroom.PackageName)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	go func() {
		t := time.NewTicker(time.Second * 50)
		for ; ; <-t.C {
			_ = session.KeepAlive()
		}
	}()

	exists, err := videoroom.Exists(handle, roomID)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	if !exists {
		config := videoroom.CreateRoomConfig{}.
			RoomID(roomID).
			MaxPublishers(100)
		err = videoroom.CreateRoom(handle, config)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
	}

	//publisher := videoroom.NewPublisher(context.Background(), handle, roomID)
	//
	//wrtc := NewLocalWebRTCAgent("./sample/sample-video-scenery.ogg", "./sample/sample-video-scenery.ivf")
	//wrtc.Start(publisher, roomID)
}
