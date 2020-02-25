# Entity Type: {{.EntityType}}

> Namespace: {{.Namespace}}

{{ if (ne .PropsTable "") }}### Properties

{{.PropsTable}}{{ end }}
{{ if (ne .NavPropsTable "") }}### Navigation Properties

{{.NavPropsTable}}{{ end }}