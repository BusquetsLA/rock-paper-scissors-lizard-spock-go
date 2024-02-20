package rpslv

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors and lizard
	PAPER    = 1 // beats rock and spock
	SCISSORS = 2 // beats paper and lizard
	LIZARD   = 3 // beats spock and paper
	SPOCK    = 4 // beats scissors and rock
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

var winMsg = []string{
	"All right, you win.",
	"You've been lucky this time.",
	"I'll be back.",
}
var loseMsg = []string{
	"It's not that easy to beat a machine.",
	"Better luck next time.",
	"I just had luck.",
}
var tieMsg = []string{
	"This is awkward.",
	"Ok, let's untie this.",
	"Hmmm.",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano()) // Generate random value
	computerValue := rand.Intn(5)    // Assign random value to computer
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "ROCK"
	case PAPER:
		computerChoice = "PAPER"
	case SCISSORS:
		computerChoice = "SCISSORS"
	case LIZARD:
		computerChoice = "LIZARD"
	case SPOCK:
		computerChoice = "SPOCK"
	}

	msgInt := rand.Intn(3) // 0 to 2
	msg := ""

	switch {
	case playerValue == computerValue:
		roundResult = "It's a draw"
		msg = tieMsg[msgInt]
	case (playerValue == 0 && computerValue == 2) || (playerValue == 0 && computerValue == 3) || (playerValue == 1 && computerValue == 0) || (playerValue == 1 && computerValue == 4) || (playerValue == 2 && computerValue == 1) || (playerValue == 2 && computerValue == 3) || (playerValue == 3 && computerValue == 4) || (playerValue == 3 && computerValue == 1) || (playerValue == 4 && computerValue == 0) || (playerValue == 4 && computerValue == 2):
		roundResult = "Player wins!"
		msg = winMsg[msgInt]
	default:
		roundResult = "Computer wins!"
		msg = loseMsg[msgInt]
	}

	result := Round{
		Message:        msg,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}

	return result
}
