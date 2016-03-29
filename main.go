package main

import "fmt"
import "math/rand"
import "time"
import "github.com/fatih/color"
import tl "github.com/JoelOtter/termloop"
import "os"

func (p *Player) Play(i int) {
  p.board.cards = append(p.board.cards, p.hand.cards[i])
  p.hand.cards = append(p.hand.cards[:i], p.hand.cards[i+1:]...)
}

func (p *Player) Draw() {
  p.hand.cards = append(p.hand.cards, deck.cards[:1]...)
  deck.cards =  deck.cards[1:]
}

var Colors = []string{"White", "Red", "Blue", "Green", "Yellow"}

var P1_HAND_Y_COORDINATE = 5
var P2_HAND_Y_COORDINATE = 15

var p1 Player
var p2 Player

var g *tl.Game

var deck CardCollection

type DrawableCard struct {
  text *tl.Text
  card Card
}

func (d *DrawableCard) Draw(s *tl.Screen) {
  d.text.Draw(s)
}

func (d *DrawableCard) Tick(ev tl.Event) {
  // Do stuff on each tick??
  if ev.Type == tl.EventKey {
    switch ev.Key {
      case tl.KeyEsc:
        os.Exit(3)
      case tl.KeyEnter:
        p2.Draw()
        DrawHand(p2.hand.cards, P2_HAND_Y_COORDINATE, g)
    }
  }
}

func main() {
  println("Hello! Welcome to GoExplore.\n\n");
  rand.Seed(time.Now().UTC().UnixNano())

  deck = InitializeDeck(deck)
  deck = Shuffle(deck)

  p1.hand = CardCollection{cards: deck.cards[:8]}
  deck.cards = deck.cards[8:]

  p2.hand = CardCollection{cards: deck.cards[:8]}
  deck.cards = deck.cards[8:]

  fmt.Printf("\nP1: ")
  PrintCards(p1.hand)
  fmt.Printf("\nP2: ")
  PrintCards(p2.hand)

  //Initiate the Game Screen
  g = tl.NewGame()

  DrawHand(p1.hand.cards, P1_HAND_Y_COORDINATE, g)
  DrawHand(p2.hand.cards, P2_HAND_Y_COORDINATE, g)

  g.Start()
  // for {
  //   //Simulate P1's first turn by playing a card at random
  //   //Slice a card off of P1 and add it to the Board
  //   fmt.Printf("\nPlayer 1 plays the 3rd card:")
  //   p1.Play(3)
  //   p1.Draw(&deck)
  //
  //   fmt.Printf("\nP1: ")
  //   PrintCards(p1.hand)
  //   PrintCards(p1.board)
  //
  //   //Ask for Player Input (P2) to play a card from P2's hand
  //   fmt.Printf("\nPlayer 2's turn, enter a card index: ")
  //   var cardIndex int
  //
  //   fmt.Scanf("%d", &cardIndex)
  //
  //   p2.Play(cardIndex)
  //   p2.Draw(&deck)
  //
  //   fmt.Printf("\nP2: ")
  //   PrintCards(p2.hand)
  //   PrintCards(p2.board)
  //
  //   if len(deck.cards) == 0 {
  //     fmt.Printf("Game over!")
  //     os.Exit(2)
  //   }
  // }
}

func DrawHand(cards []Card, Y_COORDINATE int, g *tl.Game) {
  // For each Card in P1's hand...
  i := 0

  for _,card := range cards {
    cardValue := fmt.Sprintf("%d", card.value)

    g.Screen().AddEntity(&DrawableCard{
      text: tl.NewText(i,Y_COORDINATE, cardValue, GetTLColorFromString(card.color), tl.ColorBlack),
      card: card,
    })

    i = i+4
  }

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

func GetTLColorFromString(color string) (tl.Attr) {
  switch color {
    case "Yellow":
      return 4
    case "Green":
      return 3
    case "Blue":
      return 5
    case "Red":
      return 2
    case "White":
      return 8
    default:
      return 0
    }
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
