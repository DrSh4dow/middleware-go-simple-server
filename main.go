package main

import (
	"html/template"
	"os"
)

var someHtml = `
<html>
  <body>
    <h1>Hello Black Gophers!</h1>
    {{.}}
  </body>
</html>
`

func main() {
	t, err := template.New("hello").Parse(someHtml)
	if err != nil {
		panic(err)
	}

	_ = t.Execute(os.Stdout, "<script>alert('molly')</script>")

}
