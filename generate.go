package triematcher

import (
	"os"
	"text/template"
)

var templateFile = template.Must(template.New("file").Parse(`
// DO NOT EDIT!
// Code generated by go-triematcher <https://github.com/Maki-Daisuke/go-triematcher>
// DO NOT EDIT!

package {{ .PackageName }}

import (
  "strings"
)

func {{ .FuncName }}(str string) bool {
  defer func(){
    recover() // Must be "index out of range" error for string.
              // Ignore and return false.
  }()

  goto STATE_{{ with index .States 0 }}{{ .Id }}{{ end }}
{{ range .States }}
  STATE_{{ .Id }}:
  {{ if .IsGoal }}
      return len(str) == 0
  {{ else }}
    switch{
    {{ range .OutBounds }}
      case strings.HasPrefix(str, {{ printf "%q" .Key }}):
        str = str[{{ .Key | len }}:]
        goto STATE_{{ .State.Id }}
    {{ end }}
      default:
        return false
    }
  {{ end }}
{{ end }}
}
`))

func generate(pkg_name, func_name string, st *state) {
	err := templateFile.Execute(os.Stdout, map[string]interface{}{
		"PackageName": pkg_name,
		"FuncName":    func_name,
		"States":      allStates(st),
	})
	if err != nil {
		panic(err)
	}
}
