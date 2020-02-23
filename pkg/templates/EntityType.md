# Namespace: {{.Namespace}}
## Entity Type: {{.EntityType}}

{{ if (ne .PropsTable "") }}### Properties

**Availability matrix**

{{.PropsTable}}{{ end }}
{{ if (ne .NavPropsTable "") }}### Navigation Properties

**Availability matrix**

{{.NavPropsTable}}{{ end }}