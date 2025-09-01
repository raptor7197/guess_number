package main

import (
"fmt"
//"bufio"
//"os"
//"strings"
//"strconv"
"math/rand"
"time"
)
 func main() {
fmt.Println("Welcome to the number guessing game")
	fmt.Println("Rules:")
	fmt.Println("- The computer will pick a number between 1 and 100.")
	fmt.Println("- Choose a difficulty level to determine how many attempts you get.")
	fmt.Println("- After each guess, you’ll be told if the number is higher or lower.")
	fmt.Println("- You win if you guess the number before running out of chances.")
	fmt.Println()

rand.Seed(time.Now().UnixNano())
target:=rand.Intn(100) +1

var level string 
attempts:=0
fmt.Println("choose difficulty : easy, medium ,hard")
fmt.Scanln(&level)

switch level {
case "easy":
attempts = 10
case "medium":
attempts=5
case "hard":
attempts=3
default:
fmt.Println("dumbass setting to medium")
attempts=5
}
var guess int 
for i:=0 ;i<=attempts;i++ {
fmt.Printf("Attempt %d of %d enter your guess ",i,attempts)
fmt.Scanln(&guess)

if guess==target {
fmt.Println("HORAAYYYYYY")
} else if guess < target {
fmt.Println("too low mate")
} else  { 
   {fmt.Println("too fuckin high ")}
//{
//fmt.Println("better funckin luck next time")
//}
}}
fmt.Println("ran out of attempts the number was %v ",target)

}
