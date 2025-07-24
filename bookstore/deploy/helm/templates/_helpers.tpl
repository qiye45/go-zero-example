{{/*
通用label模板，可用于后续扩展
*/}}
{{- define "bookstore.labels" -}}
app.kubernetes.io/name: {{ include "bookstore.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}} 