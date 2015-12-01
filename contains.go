package triegun

import (
	"io"
	"text/template"
)

func (p *Plant) genContains(w io.Writer, st *state) error {
	allowSubmatch(st)
	states := allStatesUpToGoal(st)
	return tmplContains.Execute(w, map[string]interface{}{
		"TagName": p.TagName,
		"Start":   states[0],
		"States":  states[1:],
	})
}

var tmplContains = template.Must(template.New("contains").Parse(`
func Contains{{ .TagName }}String(str string) bool {
  return Contains{{ .TagName}}(*(*[]byte)(unsafe.Pointer(&str)))
}

func Contains{{ .TagName }}(bytes []byte) bool {
  defer func(){
    recover() // Must be "index out of range" error for string.
              // Ignore and return false.
  }()

  i := 0

{{ $start := .Start }}
  STATE_{{ $start.Id }}:
{{ if .IsGoal }}
    return true
{{ else }}
    switch bytes[i] {
  {{ range $key, $next := $start.Nexts }}
    case {{ printf "%q" $key }}:
      i++
      goto STATE_{{ $next.Id }}
  {{ end }}
    default:
      i++
      goto STATE_{{ $start.Id }}
    }
{{ end }}

{{ range .States }}
  STATE_{{ .Id }}:
  {{ if .IsGoal }}
      return true
  {{ else }}
    switch bytes[i] {
    {{ range $key, $next := .Nexts }}
    case {{ printf "%q" $key }}:
      i++
      goto STATE_{{ $next.Id }}
    {{ end }}
    default:
      goto STATE_{{ $start.Id }}
    }
  {{ end }}
{{ end }}
}
`))
