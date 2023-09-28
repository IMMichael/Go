package main

import (
  "fmt"
  "net/http"
  "os"
  "io/ioutil"
  "jaytaylor.com/html2text"
  "flag"
)

const url = "https://cht.sh"
func main() {

  lang := os.Args[1]
//  topic := os.Args[2]
  var topic string

  flag.StringVar(&topic, "topic", "", "Topic")
  flag.Parse()


  if lang == "?" || lang == "help" {
    fmt.Println("Syntax is:")
    fmt.Println("cht arg1 arg2")
    fmt.Println("arg1 = Language or tool")
    fmt.Println("arg2 = Language only - Language topic")
    os.Exit(0)
  }

  var new_url string
  fmt.Println(lang, topic)
  if len(topic) > 0 {
    new_url = url + "/" + lang + "/" + topic
  } else if len(topic) == 0 {
    new_url = url + "/" + lang
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
 // strippedHtml := re.ReplaceAllString(content, "")
  text, err := html2text.FromString(content, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
//  fmt.Printf(content)
}
