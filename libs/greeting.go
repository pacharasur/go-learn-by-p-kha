package libs

import "fmt"

func Greeting() {
    fmt.Println("hello developer")

}
func Login(username string, password string) bool {
    if username == "admin" && password == "1234" {
   	 return true
    }
    return false
}
