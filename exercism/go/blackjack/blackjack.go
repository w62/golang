package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := 0
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "queen":
		value = 10
	case "king":
		value = 10
	default:
		value = 0
	}
	return value
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	aces := 0

	result := ""

	if card1 == "ace" {
		aces ++
	}

	if card2 == "ace" {
		aces ++
	}

	if dealerCard == "ace" {
		aces ++
	}

	if aces >= 2 {
		result = "P"
		return result
	}

	hand := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)

	switch (true) {
	case hand == 21:
		if dealer < 10 {
			result = "W"
		} else {
			result = "S"
		}
	case hand >= 17 && hand <= 20:
		result = "S"

	case hand >= 12 && hand <= 16:
		  if dealer >= 7 {
			  return "H"
		  } else {
			  return "S"
		  }

	  case hand <= 11:
		  result = "H"
	  default:
		  result = "S"
	  }

	  return result





	panic("Please implement the FirstTurn function")
}
