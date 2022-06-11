package util

import (
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
