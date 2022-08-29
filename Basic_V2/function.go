package main

import "fmt"

// * easy
// todo important
// ! dont forget

// function tanpa return
func sayHello(firstname string, lastname string, nomor int){
	fmt.Printf("hello %s %s %d \n",firstname,lastname,nomor)
}

// fuction dengan return
func getHello(name string) string {
	return "hay "+name
}

// todo Function multiple values
func getFullName(firstname string, lastname string)(string, string){
	return firstname, lastname
}

//todo Named return value
func getCompleteName()(firstname,midname,lastname string){
	firstname = "anthon"
	midname = "broklyn"
	lastname = "berg"
	return
}

//todo variadic function
//todo mirip seperti args** kalo di python
// bisa juga begini func rumus(nama string, value ...int)
func sumAll(numbers ...int){
	for _,number := range numbers{
		fmt.Printf("%d, ", number)
	}
	fmt.Printf("ini indeks ke %d \n",numbers[1])
}

//todo function value: memungkinkan kita untuk menyimpan function
//todo didalam variable persegi := getPersegi
func getGoodBye(name string) string {
	return "goodbye "+ name
}


//todo Function as parameter
//! jangan lupa harus di inget inget
func sayHelloWithFilter(name string, filter func(string) string ){
	nameFiltered := filter(name)
	fmt.Println("hello",nameFiltered)
}

//! bikin fuction as parameter
func spamFilter(name string) string{
	if name == "Anjing"{
		return "..."
	}else{
		return name
	}
}

//! type declaration biar gak kepanjangan nanti di paramnya
type Is_even func(int) bool

func genap(nilai int, is_even Is_even){
	answer := is_even(nilai)
	fmt.Printf("%t %d \n",answer,nilai)
}

func is_even(nilai int) bool{
	if nilai % 2 == 0{
		return true
	}else{
		return false
	}
}

//! anonymous function
func registerUser(name string, blacklist func(string) bool){
	if blacklist(name) {
		fmt.Println("youre blocked", name)
	}else{
		fmt.Println("welcome ",name)
	}
}



func main(){

	// memanggil fungsi sayhello
	for i := 0; i < 10; i++ {
		sayHello("rizky","darmadi",i)
	}

	//  memanggil fungsi get hello
	fmt.Println(getHello("Rayi"))

	// memanggil fungsi getFullName
	fmt.Println(getFullName("raisa","andriana"))

	// deklarasi multiple return function
	firstname, lastname := getFullName("ryan","dahl")
	fmt.Printf("%s %s \n",firstname,lastname)

	// ignore multiple return function
	firstname_, _:= getFullName("park","saeroyi")
	fmt.Println(firstname_)

	// memanggil function named return values
	fmt.Println(getCompleteName())

	//todo variadic function
	sumAll(1,2,5,3,4)

	//! variadic function dengan slice
	slice := []int{1,2,3,4,5,6,7,8}
	sumAll(slice...)

	//! memanggil function value
	goodbye := getGoodBye
	fmt.Println((goodbye("ariel")))

	//! call function as parameter
	//! function spamfillter di call barengan
	sayHelloWithFilter("Rizky",spamFilter)
	sayHelloWithFilter("Anjing",spamFilter)

	//! fungsi is even di call
	genap(4,is_even)

	//! memanggil anonymous function
	blacklist := func (name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("rizky",blacklist)
}