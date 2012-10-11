package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

const html = `
<!DOCTYPE html>
<html>

	<head>
		<title>Butt Generator</title>

		<style>
		html {
			background: url(images/butt.jpg) no-repeat center center fixed;
			-webkit-background-size: cover;
			-moz-background-size: cover;
			-o-background-size: cover;
			background-size: cover;
		}
		</style>
	</head>

	<body> </body>

</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html)
}
