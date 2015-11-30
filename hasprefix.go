package triegun

import (
	"io"
	"text/template"
)

func (p *Plant) genHasPrefix(w io.Writer, st *state) error {
	return templPrefix.Execute(w, map[string]interface{}{
		"TagName": p.TagName,
		"States":  allStatesUpToGoal(st),
	})
}

var templPrefix = template.Must(template.New("prefix").Parse(`
func HasPrefix{{.TagName}}String(str string) bool {
  return HasPrefix{{.TagName}}(*(*[]byte)(unsafe.Pointer(&str)))
}

func HasPrefix{{.TagName}}(bytes []byte) bool {
  defer func(){
    recover() // Must be "index out of range" error for string.
              // Ignore and return false.
  }()

  i := 0
  goto STATE_{{ with index .States 0 }}{{ print .Id }}{{ end }}

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
      return false
    }
  {{ end }}
{{ end }}
}
`))
