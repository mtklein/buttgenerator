package buttgenerator

import (
	"html/template"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

// hi hi hi hi mike!
var html = template.Must(template.New("html").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<meta name="mobile-web-app-capable" content="yes">
		<link rel="icon" sizes="192x192" href="favicon.ico">

		<meta name="viewport" content="initial-scale=1, user-scalable=no">

		<title>Butt Generator</title>
		<script>
			{{.JS}}
		</script>
		<style>
			.butt {
				position: fixed;
				top: 0;
				left: 0;
				min-width: 100%;
				min-height: 100%;
				max-width: 100%;
				max-height: 100%;
			}
		</style>
	</head>

	<body>
		{{.Image}}
	</body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	image := Butt(r.URL.Path == "/spot")
	js := svg_js
	if (r.URL.Path == "/original") {
		image = `<img class="butt" src="images/butt.jpg" alt="A smiling butt.">`
		js = ""
	}

	err := html.Execute(w, struct {
		Image template.HTML
		JS    template.JS
	}{
		Image: template.HTML(image),
		JS: template.JS(js),
	})

	if err != nil {
		log.Println("error executing template: ", err)
	}
}
