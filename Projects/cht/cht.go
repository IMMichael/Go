package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "jaytaylor.com/html2text"
  "flag"
  "os"
)

const url = "https://cht.sh"
func main() {

  flag.Parse()

  lang := flag.Arg(0)
  topic := flag.Arg(1)

  if lang == "?" || lang == "help" {
    fmt.Println("Syntax is:")
    fmt.Println("cht arg1 arg2")
    fmt.Println("arg1 = Language or tool")
    fmt.Println("arg2 = Language only - Language topic")
    os.Exit(0)
  }

  var new_url string

//  fmt.Println(lang, topic)

  if lang := flag.Arg(0);  topic != "" {
    new_url = url + "/" + lang + "/" + topic
//    fmt.Println(new_url)
  } else if topic == "" {
    new_url = url + "/" + lang
//    fmt.Println(new_url)
  }

  resp, err := http.Get(new_url)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Response type: %T\n", resp) // The type will be a pointer named *http.Response
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes) //Converts the bytes received to a string using the string function
  text, err := html2text.FromString(content, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
//  fmt.Printf(content)
}
