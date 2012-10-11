package buttgenerator

import (
	"html/template"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

var html = template.Must(template.New("html").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<title>Butt Generator</title>
		<style>
			html {
				background: url({{.Image}}) no-repeat center center fixed;
				background-size: cover;
			}
		</style>
	</head>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	image := "data:image/svg+xml;utf8," + butt_inline
	if (r.URL.Path == "/original") {
		image = "images/butt.jpg"
	}

	err := html.Execute(w, struct {
		Image template.URL
	}{
		Image: template.URL(image),
	})

	if err != nil {
		log.Println("error executing template: ", err)
	}
}
