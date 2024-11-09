//The print(), prints each function with its argument

package main
import ("fmt")

func main() {
	var i , j = "Hello","World"

	//Prints on oneline
	fmt.Print(i)
	fmt.Print(j, "\n") 
	//Prints on 2 lines 
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
    //Prints on separate line but in one code
	fmt.Print(i, "\n", j)

	//We can aso also use " " to give spaces
	fmt.Print(i," ",j)

	//If neither are strings, the Print automatically inserts a space between them
	var e , a = 10 , 20 
	
	fmt.Print("\n",e,a)

	//Println addd white spaces
	fmt.Println("\n",e,a)

	//Printf is used to format argument based on the given formatting verb and prints them
	//%v used to print the value of the argument
	//%T used to print the type of the arguments 
	var g = 10
	var f = "Welcome to My World"
	
	
	fmt.Printf("g has a vlaue of: %v and a type of: %T \n",g,g)
	fmt.Printf("f  has a value of : %v and type of: %T \n",f,f)

}