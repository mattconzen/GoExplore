package main

import "fmt"
import "math/rand"
import "time"
import "github.com/fatih/color"

type Card struct {
  color string
  value int
}

type CardCollection struct {
  cards []Card
}

type Player struct {
  hand CardCollection
}

var Colors = []string{"White", "Red", "Blue", "Green", "Yellow"}

func main() {
  println("Hello! Welcome to GoExplore.\n\n");
  rand.Seed(time.Now().UTC().UnixNano())

  var deck CardCollection

  deck = InitializeDeck(deck)
  deck = Shuffle(deck)

  PrintCards(deck)

  var p1 Player
  var p2 Player

  p1.hand = CardCollection{cards: deck.cards[:8]}
  p2.hand = CardCollection{cards: deck.cards[8:16]}

  PrintCards(p1.hand)
  PrintCards(p2.hand)

  //TODO: Simulate P1's first turn by playing a card at random

  //TODO: Ask for Player Input (P2) to play a card from P2's hand

  //TODO: Add game loop: do until...len(deck.cards) = 0
}

func InitializeDeck(deck CardCollection) (CardCollection) {
  for _,color := range Colors {
    for i := 1; i <= 10; i++ {
      deck.cards = append(deck.cards, Card{color: color, value: i})
    }

    //Add betting cards as cards with a value of zero for now.
    deck.cards = append(deck.cards, Card{color: color, value: 0})
    deck.cards = append(deck.cards, Card{color: color, value: 0})
    deck.cards = append(deck.cards, Card{color: color, value: 0})
  }
  return deck
}

func Shuffle(cardCol CardCollection) (CardCollection) {
  for i := range cardCol.cards {
    j := rand.Intn(i+1)
    cardCol.cards[i], cardCol.cards[j] = cardCol.cards[j], cardCol.cards[i]
  }
  return cardCol
}

func PrintCards(deck CardCollection) {
  //Print each card in the passed in collection.
  for _,card := range deck.cards {
    switch card.color {
      case "Yellow":
        color.Set(color.FgYellow).Add(color.Bold)
      case "Green":
        color.Set(color.FgGreen).Add(color.Bold)
      case "Blue":
        color.Set(color.FgBlue).Add(color.Bold)
      case "Red":
        color.Set(color.FgRed).Add(color.Bold)
      case "White":
        color.Set(color.FgWhite).Add(color.Bold)
    }
    fmt.Printf("%s|%d\n", card.color, card.value)
    color.Unset()
  }

  //Print the count of cards in the Deck
  fmt.Printf("Card count: %d\n\n",len(deck.cards))
}

func TestStructs() {
  // Write each color in the Colors collection.
  for _,element := range Colors {
    println (element);
  }

  newCard := Card{ color: "white", value: 1 }

  println(newCard.color)
}
