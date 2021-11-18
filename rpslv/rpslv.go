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
	rand.Seed(time.Now().UnixNano()) // genera el valor aleatorio
	computerValue := rand.Intn(5)    // se lo asigna a la compu
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		// break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		// break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		// break
	case LIZARD:
		computerChoice = "Computer chose LIZARD"
		// break
	case SPOCK:
		computerChoice = "Computer chose SPOCK"
		// break
	default:
	}

	// la generacion de mensajes
	msgInt := rand.Intn(3) // 0 a 2
	msg := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		msg = tieMsg[msgInt]
	} else if playerValue == 0 && computerValue == 2 {
		roundResult = "Player wins!" // rock beats scissors
		msg = winMsg[msgInt]
	} else if playerValue == 0 && computerValue == 3 {
		roundResult = "Player wins!" // rock beats lizard
		msg = winMsg[msgInt]
	} else if playerValue == 1 && computerValue == 0 {
		roundResult = "Player wins!" // paper beats rock
		msg = winMsg[msgInt]
	} else if playerValue == 1 && computerValue == 4 {
		roundResult = "Player wins!" // paper beats spock
		msg = winMsg[msgInt]
	} else if playerValue == 2 && computerValue == 1 {
		roundResult = "Player wins!" // scissors beats paper
		msg = winMsg[msgInt]
	} else if playerValue == 2 && computerValue == 3 {
		roundResult = "Player wins!" // scissors beats lizard
		msg = winMsg[msgInt]
	} else if playerValue == 3 && computerValue == 4 {
		roundResult = "Player wins!" // lizard beats spock
		msg = winMsg[msgInt]
	} else if playerValue == 3 && computerValue == 1 {
		roundResult = "Player wins!" // lizard beats paper
		msg = winMsg[msgInt]
	} else if playerValue == 4 && computerValue == 0 {
		roundResult = "Player wins!" // spock beats rock
		msg = winMsg[msgInt]
	} else if playerValue == 4 && computerValue == 2 {
		roundResult = "Player wins!" // spock beats scissors
		msg = winMsg[msgInt]
	} else {
		roundResult = "Computer wins!"
		msg = loseMsg[msgInt]
	}

	var result Round
	result.Message = msg
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult

	return result
}
