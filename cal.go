//Name:aungelia barletta
//Date edited: 2/13/20
//Description: a calculator 
//cite:




package main

import "fmt"

func main(){
  fmt.Print("Enter a number:")
  var number1 float64
  fmt.Scanf("%f", &number1)

  fmt.Print("Enter a number:")
  var number2 float64
  fmt.Scanf("%f", &number2)

  
  output := number1 + number2
  
  output2 := number1 - number2
  
  output3 := number1 * number2
  
  output4 := number1 / number2

  fmt.Print("Addition: ")

  fmt.Println(output)
  fmt.Print("Subtraction: ")
  fmt.Println(output2)
  fmt.Print("Multiplication: ")
  fmt.Println(output3)
  fmt.Print("Divition: ")
  fmt.Println(output4)


}
