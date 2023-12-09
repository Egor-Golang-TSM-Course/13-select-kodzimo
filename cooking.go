package main

import (
	"log"
	"time"
)

type Chief struct {
	Name        string
	CookingTime time.Duration
}

func (c Chief) cooking() {
	log.Printf("Я - повар %s. Начинаю готовку", c.Name)
	time.Sleep(c.CookingTime * time.Second)
	log.Printf("Повар %s --- готовку закончил", c.Name)
}
