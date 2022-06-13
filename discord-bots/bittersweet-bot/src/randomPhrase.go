package main

import (
	"math/rand"
	"time"
)

func getRandomPhrase(p []string) string {
	rand.Seed(time.Now().Unix()) 
	return p[rand.Intn(len(p))]
}