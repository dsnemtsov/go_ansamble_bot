package util

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"path"
)

func SongTextMessage(songName string) string {
	songName += ".txt"
	filePath := path.Join("resources", "text", songName)
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Print(err)
	}
	return string(file)
}

func AudioMessage(audioPackage string, audioName string) *tele.Audio {
	audioName += ".mp3"
	filePath := path.Join("resources", "audio", audioPackage, audioName)

	audio := &tele.Audio{
		File: tele.FromDisk(filePath),
	}

	return audio
}
