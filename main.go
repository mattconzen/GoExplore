package main

import "fmt"
import "math/rand"
import "time"

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
  println("Hello! Welcome to GoExplore.");
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
  //Print the count of cards in the Deck
  fmt.Printf("Card count: %d\n",len(deck.cards))

  //Print each card in the deck.
  for _,card := range deck.cards {
    fmt.Printf("%s|%d\n", card.color, card.value)
  }
}

func TestStructs() {
  // Write each color in the Colors collection.
  for _,element := range Colors {
    println (element);
  }

  newCard := Card{ color: "white", value: 1 }

  println(newCard.color)
}
