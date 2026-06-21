package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

type Ranker interface {
	Ranking() []string
}

func (l *League) MatchResult(team1 string, team1Score int, team2 string, team2Score int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}
	if _, ok := l.Teams[team2]; !ok {
		return
	}
	if team1Score == team2Score {
		return
	}
	if team1Score > team2Score {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l *League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})

	return names
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	results := ranker.Ranking()
	for _, v := range results {
		io.WriteString(writer, v)
		writer.Write([]byte("\n"))
	}
}

func main() {
	l := League{
		Name: "NBA",
		Teams: map[string]Team{
			"Lakers": {
				Name:    "Lakers",
				Players: []string{"LeBron James", "Anthony Davis", "Russell Westbrook", "Carmelo Anthony", "Dwight Howard"},
			},
			"Warriors": {
				Name:    "Warriors",
				Players: []string{"Stephen Curry", "Klay Thompson", "Draymond Green", "Andrew Wiggins", "James Wiseman"},
			},
			"Nets": {
				Name:    "Nets",
				Players: []string{"Kevin Durant", "Kyrie Irving", "James Harden", "Blake Griffin", "Joe Harris"},
			},
			"Bucks": {
				Name:    "Bucks",
				Players: []string{"Giannis Antetokounmpo", "Khris Middleton", "Jrue Holiday", "Brook Lopez", "Bobby Portis"},
			},
			"Suns": {
				Name:    "Suns",
				Players: []string{"Chris Paul", "Devin Booker", "Deandre Ayton", "Mikal Bridges", "Jae Crowder"},
			},
		},
		Wins: map[string]int{},
	}

	l.MatchResult("Lakers", 113, "Warriors", 110)
	l.MatchResult("Nets", 120, "Bucks", 115)
	l.MatchResult("Suns", 105, "Lakers", 102)
	l.MatchResult("Warriors", 130, "Nets", 125)
	l.MatchResult("Bucks", 110, "Suns", 108)
	l.MatchResult("Lakers", 118, "Nets", 112)
	l.MatchResult("Warriors", 122, "Suns", 119)
	l.MatchResult("Bucks", 117, "Lakers", 114)
	l.MatchResult("Nets", 121, "Suns", 118)
	l.MatchResult("Warriors", 128, "Bucks", 124)

	RankPrinter(&l, os.Stdout)
}
