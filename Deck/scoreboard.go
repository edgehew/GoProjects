package main

import "fmt"

type Scoreboard struct {
	round 	int
	score map[string]int
}

func makeScoreboard(player []Player) Scoreboard {
	var s Scoreboard
	s.round = 1
	s.score = make(map[string]int)
	for _, p := range player {
		s.addPlayer(p.name)
	}
	return s
}

func (s* Scoreboard) addPlayer(name string) {
	s.score[name] = 0
}

func (s* Scoreboard) addToScore(player *Player, amount int) {
	name := player.name
	oldScore := s.score[name]
	s.score[name] += amount
	fmt.Printf("%v score changed from %v to %v\n", name, oldScore, s.score[name])
}