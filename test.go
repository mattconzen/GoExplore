package main 

func TestStructs() {
  // Write each color in the Colors collection.
  for _,element := range Colors {
    println (element);
  }

  newCard := Card{ color: "white", value: 1 }

  println(newCard.color)
}
