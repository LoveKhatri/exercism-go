package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "ace":
		return 11
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10

	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var parsedCard1 int = ParseCard(card1)
	var parsedCard2 int = ParseCard(card2)
	var parsedDealerCard int = ParseCard(dealerCard)

	switch {
	case parsedCard1 == 11 && parsedCard2 == 11:
		return "P"
	case parsedCard1+parsedCard2 == 21:
		switch {
		case parsedDealerCard != 11 && parsedDealerCard != 10:
			return "W"
		default:
			return "S"
		}
	case parsedCard1+parsedCard2 >= 17 && parsedCard1+parsedCard2 <= 20:
		return "S"
	case parsedCard1+parsedCard2 >= 12 && parsedCard1+parsedCard2 <= 16:
		switch {
		case parsedDealerCard >= 7:
			return "H"
		default:
			return "S"
		}
	case parsedCard1+parsedCard2 <= 11:
		return "H"
	default:
		return ""
	}
}
