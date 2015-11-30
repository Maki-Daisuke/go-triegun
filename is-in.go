package triegun

import (
	"io"
	"text/template"
)

func (p *Plant) genIsIn(w io.Writer, st *state) error {
	return templIsIn.Execute(w, map[string]interface{}{
		"TagName": p.TagName,
		"States":  allStates(st),
	})
}

var templIsIn = template.Must(template.New("isin").Parse(`
func IsIn{{.TagName}}String(str string) bool {
  return IsIn{{.TagName}}(*(*[]byte)(unsafe.Pointer(&str)))
}

func IsIn{{.TagName}}(bytes []byte) bool {
  defer func(){
    recover() // Must be "index out of range" error for string.
              // Ignore and return false.
  }()

  i := 0
  goto STATE_{{ with index .States 0 }}{{ print .Id }}{{ end }}

{{ range .States }}
  STATE_{{ .Id }}:
  {{ if .IsGoal }}
    if i == len(bytes) {
      return true
    }
  {{ end }}
    switch bytes[i] {
    {{ range $key, $next := .Nexts }}
    case {{ printf "%q" $key }}:
      i++
      goto STATE_{{ $next.Id }}
    {{ end }}
    default:
      return false
    }
{{ end }}
}
`))
