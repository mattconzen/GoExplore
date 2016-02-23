package main

import "fmt"
import "math/rand"
import "time"
import "github.com/fatih/color"
import "os"

type Card struct {
  color string
  value int
}

type CardCollection struct {
  cards []Card
}

type Player struct {
  hand CardCollection
  board CardCollection
}

func (p *Player) Play(i int) {
  p.board.cards = append(p.board.cards, p.hand.cards[i])
  p.hand.cards = append(p.hand.cards[:i], p.hand.cards[i+1:]...)
}

func (p *Player) Draw(deck *CardCollection) {
  p.hand.cards = append(p.hand.cards, deck.cards[:1]...)
  deck.cards =  deck.cards[1:]
}

var Colors = []string{"White", "Red", "Blue", "Green", "Yellow"}

func main() {
  println("Hello! Welcome to GoExplore.\n\n");
  rand.Seed(time.Now().UTC().UnixNano())

  var deck CardCollection

  deck = InitializeDeck(deck)
  deck = Shuffle(deck)

  var p1 Player
  var p2 Player

  p1.hand = CardCollection{cards: deck.cards[:8]}
  deck.cards = deck.cards[8:]

  p2.hand = CardCollection{cards: deck.cards[:8]}
  deck.cards = deck.cards[8:]

  fmt.Printf("\nP1: ")
  PrintCards(p1.hand)
  fmt.Printf("\nP2: ")
  PrintCards(p2.hand)

  for {
    //Simulate P1's first turn by playing a card at random
    //Slice a card off of P1 and add it to the Board
    fmt.Printf("\nPlayer 1 plays the 3rd card:")
    p1.Play(3)
    p1.Draw(&deck)

    fmt.Printf("\nP1: ")
    PrintCards(p1.hand)
    PrintCards(p1.board)

    //Ask for Player Input (P2) to play a card from P2's hand
    fmt.Printf("\nPlayer 2's turn, enter a card index: ")
    var cardIndex int

    fmt.Scanf("%d", &cardIndex)

    p2.Play(cardIndex)
    p2.Draw(&deck)

    fmt.Printf("\nP2: ")
    PrintCards(p2.hand)
    PrintCards(p2.board)

    if len(deck.cards) == 0 {
      fmt.Printf("Game over!")
      os.Exit(2)
    }
  }
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
        color.Set(color.FgYellow, color.Bold)
      case "Green":
        color.Set(color.FgGreen, color.Bold)
      case "Blue":
        color.Set(color.FgBlue, color.Bold)
      case "Red":
        color.Set(color.FgRed, color.Bold)
      case "White":
        color.Set(color.FgWhite, color.Bold)
    }
    fmt.Printf("%d ", card.value)
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
