package html

import (
	"html/template"
	"os"
)

func GenerateHTML(images []string, filename string) error {
	tmpl := `<!doctype html>
<html lang="fr">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Tests</title>
		<link rel="stylesheet" href="https://atontour.eu/coleoptera/css/coleoptera-12.min.css" />
    </head>
	<body class="container container-xx">

	<section class="row">
{{range .}}
<figure class="col-3">
	<a href="{{.}}" target="_blank"><img src="{{.}}" alt="" /></a>
	<figcaption>{{.}}</figcaption>
</figure>
{{end}}
	</section>

</body>
</html>`

	t := template.Must(template.New("gallery").Parse(tmpl))
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, images)
}
