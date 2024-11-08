//This is a comment
package main

import ("fmt")



func numbers(){
    var a int
    var b int
    var c float32
    var d string
    var e bool
    var f , g ,h,i  = 2 ,4,"Afua",6
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
    fmt.Println(g)
    fmt.Println(h)
    fmt.Println(i)

}
const (
    A int = 1
    B = 3.14
    C = "Hi!"
  )

func main(){
    fmt.Println("Hello World")
    var name string = "Antwiwaa"
    names := "Afua Antwiwaa"

    fmt.Println(name)
    fmt.Println(names)
    
    fmt.Println(A)
    fmt.Println(B)
    fmt.Println(C)
    numbers()
}





//Variable Types in go
//1.int
//2.float32
//Sting - They are strictly in double quotes
//bool - Values that are either true or false

//var is used to declare a variable 
//syntax - var variablename type = value  or variablename := value


// You can also declare a variable without assigning a value to it 

//:= must be directly assigned a value, var can be used to declare a value and use later 



