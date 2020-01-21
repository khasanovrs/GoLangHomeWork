package main

import (
	"html/template"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	const tmpl = `<doctype html>
<html>
	<head>
    	<title>Hello World!</title>
		<meta charset="utf-8">
	</head>
	<body>
    	Hello World!
		<div>
			Hello {{.Name}}!
		</div>
	</body>
</html>`

	data := struct {
		Name string
	}{req.URL.Query().Get("name")}

	t := template.Must(template.New("").Parse(tmpl))

	err := t.Execute(res, data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8181", nil)
}
