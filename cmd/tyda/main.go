package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/alexflint/go-arg"

	"github.com/walle/tyda-api"
)

var args struct {
	Query     string   `arg:"positional,required"`
	Simple    bool     `arg:"-s,env,help:Simple output"`
	Languages []string `arg:"-l,env,help:Languages to translate to (en fr de es la nb sv)"`
}

const advanced = `Search term: {{ .SearchTerm }}
Language: {{ .Language }}
Word class: {{ .WordClass }}
Conjugations: {{ join .Conjugations ", " }}
Translations
{{ range .Translations }}{{ .Language }}{{ if .Description }} - {{ .Description }} {{ end }}{{ range .Words  }}
  {{ .Value }}{{ if .Context }} - {{ .Context }}{{ end }}{{ if .PronunciationURL }}
    Pronunciation: {{ .PronunciationURL }}{{ end }}{{ if .DictionaryURL }}
    Dictionary: {{ .DictionaryURL }}{{ end}}{{ end }}
{{ end }}{{ if .Synonyms }}Synonyms
{{ range .Synonyms }}  {{ .Value }} {{ if .Context }}- {{ .Context }} {{ end }}
{{ end }}{{ end }}`

const simple = `{{ range .Translations }}{{ range .Words }}{{ .Value }} {{ end}}{{ end }}
Synonyms: {{ range .Synonyms }}{{ .Value }} {{ end }}
`

var funcs = template.FuncMap{"join": strings.Join}

func main() {
	args.Languages = []string{"en"}
	arg.MustParse(&args)

	arg := []string{"-l"}
	arg = append(arg, args.Languages...)
	arg = append(arg, "--", args.Query)
	cmd := exec.Command("tyda-api", arg...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running tyda-api: %s : %s\n", err, stderr.String())
		os.Exit(1)
	}
	var resp tydaapi.Response
	err = json.Unmarshal(output, &resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing response: %s\n", err)
		os.Exit(1)
	}
	var t *template.Template
	if args.Simple {
		t = template.Must(template.New("simple").Funcs(funcs).Parse(simple))
	} else {
		t = template.Must(template.New("advanced").Funcs(funcs).Parse(advanced))
	}
	err = t.Execute(os.Stdout, resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error printing result: %s\n", err)
		os.Exit(1)
	}
}
