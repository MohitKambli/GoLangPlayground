package main

import (
	"fmt"
	"strings"
	"sort"
	"math"
)

/*func sayGreeting(name string) {
	fmt.Printf("Good Morning %v\n", name)
}

func sayBye(name string) {
	fmt.Printf("Good Bye %v\n", name)
}

func cycleNames(name []string, f func(string)) {
	for _, value := range name {
		f(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}*/

/*func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")
	initials := []string{}
	for _, value := range names {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}*/

// Pass By Value
// Non-Pointer Values
/*func updateName(name string) string {
	name = "wedge"
	return name
}*/

// Pointer Wrapper Values
/*func updateMenu(menu map[string]float64) {
	menu["coffee"] = 2.99
}*/

// Pass By Reference (Pointers)
/*func updateName(name *string) {
	*name = "wedge"
}*/

// For structs, Pointers are automatically dereferenced
/*type bill struct {
	name  string
	items map[string]float64
	tip   float64
}*/

// Make new Bills
/*func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}*/

// Format the Bill
/*func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0
	// List Items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+": ", value)
		total += value
	}

	// Tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip: ", b.tip)

	// Total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total: ", total+b.tip)

	return fs
}*/

// Update Tip
/*func (b *bill) updateTip(tip float64) {
	b.tip = tip
}*/

// Add An Item to Bill
/*func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}*/

// Save Bill
/*func (b *bill) save() {
	data := [] byte {b.format()}
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}*/

// Creating a Bill
/*func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Create a new bill name: ")
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}*/

// Fetching Input from User
/*func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}*/

/*func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
		case "a":
			name, _ := getInput("Item Name: ", reader)
			price, _ := getInput("Item Price: ", reader)
			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("The price must be a number")
				promptOptions(b)
			}
			b.addItem(name, p)
			fmt.Println("Item Added - "name, price)
			promptOptions(b)
		case "t":
			tip, _ := getInput("Tip Amount: ", reader)
			t, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The price must be a number")
				promptOptions(b)
			}
			b.updateTip(t)
			fmt.Println("Tip Added - "tip)
			promptOptions(b)
		case "s":
			b.save()
			fmt.Println("You chose to save the bill", b.name)
		default:
			fmt.Println("You chose an invalid option!")
	}
}*/






/*
// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}
*/








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

	// Strings Package 
	//greeting := "Hello there friends!"
	// (Atleast, the following methods of the strings package don't modify the original string)
	/*fmt.Println(strings.Contains(greeting, "Hello"))
	// Doesn't alter original string
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))
	// Original Value Remains Unchanged
	fmt.Println("Original String: ", greeting) */

	// Sort Package
	/*ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	// Alters the orignal slice/array
	fmt.Println(ages)
	// Searches in the sorted slice/array and not in the original array
	validIntIndex := sort.SearchInts(ages, 60)
	fmt.Println(validIntIndex)
	invalidIntIndex := sort.SearchInts(ages, 100)
	// Doesn't return -1, but instead returns an index which is out of bounds, specifically n
	fmt.Println(invalidIntIndex)
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)
	validStringIndex := sort.SearchStrings(names, "mario")
	fmt.Println(validStringIndex)*/

	// Loops
	/*x := 0
	for x < 5 {
		fmt.Println("Value of X: ", x)
		x++
	}*/
	/*for i := 0; i < 5; i++ {
		fmt.Println("Value of i: ", i)
	}*/
	//names := []string{"mario", "luigi", "yoshi", "peach"}
	/*for i := 0; i < len(names); i++ {
		fmt.Println("Name: ", names[i])
	}*/
	/*for index, value := range names {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}*/
	// Underscore can be used for the option that is not required
	// Option can either be the index or the actual value
	// If we don't use a particular option, then we would get an error, hence use underscore
	/*for _, value := range names {
		fmt.Printf("Value: %v\n", value)
	}*/
	/*for _, value := range names {
		fmt.Printf("Value: %v\n", value)
		// Following line doesn't make any changes to the original slice/array
		value = "new string"
	}
	fmt.Println(names)*/

	// Booleans & Conditionals
	/*age := 45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)
	if age < 50 {
		fmt.Println("Age is less than 50")
	} else if age > 50 {
		fmt.Println("Age is greater than 50")
	} else {
		fmt.Println("Age is equal to 50")
	}*/
	/*names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at Index: ", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at Index: ", index)
			break
		}
		fmt.Printf("Value at index %v : %v\n", index, value)
	}*/

	// Functions
	/*sayGreeting("mario")
	sayBye("luigi")
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)
	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1, ", ", a2)
	fmt.Printf("Circle 1 : %0.3f\nCircle 2 : %0.3f", a1, a2)*/

	// More on Functions
	/*firstName1, secondName1 := getInitials("tifa lockhart")
	fmt.Println(firstName1, ", ", secondName1)
	firstName2, secondName2 := getInitials("cloud strife")
	fmt.Println(firstName2, ", ", secondName2)
	firstName3, secondName3 := getInitials("barret")
	fmt.Println(firstName3, ", ", secondName3)*/

	// Maps
	/*menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])
	for key, value := range menu {
		fmt.Println(key, " - ", value)
	}
	phonebook := map[int]string{
		12345: "mario",
		23456: "luigi",
		34567: "peach",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[23456])
	// Updating value of one existing key
	phonebook[34567] = "bowser"
	for key, value := range phonebook {
		fmt.Println(key, " - ", value)
	}*/

	// Pass By Value
	// Non-Pointer Values
	/*name := "tifa"
	name = updateName(name)
	fmt.Println(name)*/

	// Pointer Wrapper Values
	/*menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)*/

	// Pointers
	/*name := "tifa"
	mName := &name
	fmt.Println("Memory Address: ", &mName)
	fmt.Println("Value at Memory Address: ", *mName)
	fmt.Println(name)
	updateName(mName)
	fmt.Println(name)*/

	// Bill Initialization (Struct)
	/*myBill := createBill()
	promptOptions(myBill)*/

	/*
	// Interface
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
 	*/
}
