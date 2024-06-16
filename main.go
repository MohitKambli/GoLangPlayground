package main

import "fmt"

func main() {
	// Println Function
	/*fmt.Println("Hello World")*/

	// Print Function
	/*fmt.Print("Rengoku ")
	fmt.Print("Mitsuri")*/

	// Printf Function
	/*name := "Mohit"
	age := 26
	fmt.Printf("My name is %v and my age is %v\n", name, age)
	fmt.Printf("My name is %q and my age is %q\n", name, age)
	fmt.Printf("Age is of type %T\n", age)
	fmt.Printf("You scored %.2f points in the game\n", 92.65)
	fmt.Printf("You scored %.1f points in the game\n", 92.66)*/

	// Sprintf Function
	/*name := "Mohit"
	age := 26
	var str = fmt.Sprintf("My name is %v and my age is %v\n", name, age)
	fmt.Println("Saved String : ", str)*/

	// Strings
	/*var nameOne string = "Mohit"
	var nameTwo = "Kambli"
	var nameThree string
	fmt.Println("Hello " + nameOne + " " + nameTwo + "!")
	nameOne = "Ameya"
	nameThree = "Kale"
	nameFour := "Pratik"
	fmt.Println("Hello " + nameOne + " " + nameThree + "!" + nameFour)*/

	// Integers
	/*var ageOne int = 25
	var ageTwo = 30
	var ageThree int
	ageFour := 50
	fmt.Println(ageOne, ageTwo, ageThree, ageFour)*/

	// Floats
	/*var scoreOne float32 = 9.5
	var scoreTwo = 8.3
	var scoreThree float64
	scoreFour := 9.9
	fmt.Println(scoreOne, scoreTwo, scoreThree, scoreFour)*/

	// Arrays
	/*//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))
	names := [4]string{"abc", "pqr", "lmn", "xyz"}
	fmt.Println(names, len(names))*/

	// Slices
	/*scores := []int{10, 20, 30}
	scores = append(scores, 40)
	fmt.Println(scores, len(scores))*/

	// Slice Ranges
	/*rangeOne := scores[1:3]
	fmt.Println(rangeOne, len(rangeOne))
	rangeTwo := scores[2:]
	fmt.Println(rangeTwo, len(rangeTwo))
	rangeThree := scores[:3]
	fmt.Println(rangeThree, len(rangeThree))
	rangeThree = append(rangeThree, 40)
	fmt.Println(rangeThree, len(rangeThree))*/
}
