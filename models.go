package main 

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