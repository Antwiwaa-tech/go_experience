//%v is used for printing the vlaue of a variable
//

package main

import ("fmt")

func main(){
    var i = 15.5
    var txt = "Hello World!"

    fmt.Printf("%v\n", i)
    fmt.Printf("%#v\n", i)
    fmt.Printf("%v%%\n", i)
    fmt.Printf("%T\n", i)

    fmt.Printf("%v\n", txt)
    fmt.Printf("%#v\n", txt)
    fmt.Printf("%T\n", txt)

    fmt.Printf("\n%b\n", i)
  fmt.Printf("%d\n", i)
  fmt.Printf("%+d\n", i)
  fmt.Printf("%o\n", i)
  fmt.Printf("%O\n", i)
  fmt.Printf("%x\n", i)
  fmt.Printf("%X\n", i)
  fmt.Printf("%#x\n", i)
  fmt.Printf("%4d\n", i)
  fmt.Printf("%-4d\n", i)
  fmt.Printf("%04d\n", i)

  fmt.Printf("\n%s\n", txt)
  fmt.Printf("%q\n", txt)
  fmt.Printf("%8s\n", txt)
  fmt.Printf("%-8s\n", txt)
  fmt.Printf("%x\n", txt)
  fmt.Printf("% x\n", txt)

  var a = true
  var b = false

  fmt.Printf("\n%t\n", a)
  fmt.Printf("%t\n", b)
  var c = 3.141

  fmt.Printf("\n%e\n", c)
  fmt.Printf("%f\n", c)
  fmt.Printf("%.2f\n", c)
  fmt.Printf("%6.2f\n",c)
  fmt.Printf("%g\n", c)

  var d bool = true
  var e int = 5
  var f float32 = 45.88
  var g string = "Antwiwaa"

  fmt.Printf("\nBoolean:", d)
  fmt.Printf("\nInteger:", e)
  fmt.Printf("\nFloat:", f)
  fmt.Printf("\nString:", g)

  fmt.Printf("\nBoolean: %v", d)  // Use %v for boolean
  fmt.Printf("\nInteger: %d", e)   // Use %d for integer
  fmt.Printf("\nFloat: %f", f)     // Use %f for float
  fmt.Printf("\nString: %s\n", g)    // Use %s for string

  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}