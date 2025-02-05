package main

import (
	"log"
	"time"

	"github.com/robfig/cron"

	"github.com/RoggXD/gin-blog/models"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	_, _ = c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	_, _ = c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}