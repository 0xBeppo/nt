package cmd

type MyNote struct {
	Title    string
	Date     string
	Tags     []string
	OldTasks []string
}

const NoteTemplate = `
---
title: {{ .Title }}
date: {{ .Date }}
tags:
{{- range .Tags }}
  - {{ . }}
{{- end }}
---

`

const TodoTemplate = `---
title: {{ .Title }} TODOs
date: {{ .Date }}
tags:
{{- range .Tags }}
  - {{ . }}
{{- end }}
---

## TODOs

{{ range .OldTasks -}}
{{ . }}
{{- end }}
- [ ] 

## In Progress

## Completed ✓

`
