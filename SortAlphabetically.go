func main() {

 	fruits := "banana,apple,orange,grape"
 	fruitList := strings.Split(fruits, ",")
 	for i, fruit := range fruitList {
 		fruitList[i] = strings.TrimSpace(fruit)
 	}
 	sort.Strings(fruitList)

 	for i, fruit := range fruitList {
 		fruitList[i] = strings.ToUpper(fruit[:1]) + strings.ToLower(fruit[1:])
 		fmt.Println(fruitList[i])
 	}

 }
