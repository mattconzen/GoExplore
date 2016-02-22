package main

import "fmt"
import "math/rand"
import "time"

type Card struct {
  Color string
  Value int
}

type CardCollection struct {
  Cards []Card
}

type P1 struct {
  Cards CardCollection
}

type P2 struct {
  Cards CardCollection
}

var Colors = []string{"White", "Red", "Blue", "Green", "Yellow"}

func main() {
  println("Hello! Welcome to GoExplore.");
  rand.Seed(time.Now().UTC().UnixNano())

  var deck CardCollection

  deck = InitializeDeck(deck)
  deck = Shuffle(deck)

  TestDeckSetup(deck)
}

func InitializeDeck(deck CardCollection) (CardCollection) {
  for _,color := range Colors {
    for i := 1; i <= 10; i++ {
      deck.Cards = append(deck.Cards, Card{Color: color, Value: i})
    }

    //Add betting cards as cards with a value of zero for now.
    deck.Cards = append(deck.Cards, Card{Color: color, Value: 0})
    deck.Cards = append(deck.Cards, Card{Color: color, Value: 0})
    deck.Cards = append(deck.Cards, Card{Color: color, Value: 0})
  }
  return deck
}

func Shuffle(cards CardCollection) (CardCollection) {
  for i := range cards.Cards {
    j := rand.Intn(i+1)
    cards.Cards[i], cards.Cards[j] = cards.Cards[j], cards.Cards[i]
  }
  return cards
}

func TestDeckSetup(deck CardCollection) {
  //Print the count of cards in the Deck
  fmt.Printf("Deck count: %d\n",len(deck.Cards))

  //Print each card in the deck.
  for _,card := range deck.Cards {
    fmt.Printf("%s|%d\n", card.Color, card.Value)
  }
}

func TestStructs() {
  // Write each color in the Colors collection.
  for _,element := range Colors {
    println (element);
  }

  newCard := Card{ Color: "white", Value: 1 }

  println(newCard.Color)
}
