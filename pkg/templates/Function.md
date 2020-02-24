# Namespace: {{.Namespace}}

## Function Import: {{.Function}}

{{ if (ne .ReturnType "") }}- ReturnType: {{.ReturnType}}{{ end }}
- IsComposable: {{.IsComposable}}
- IsBindable: {{.IsBindable}}
{{ if (ne .EntitySet "") }}- EntitySet: {{.EntitySet}}{{ end }}

{{ if (ne .ParamsTable "") }}### Parameters

**Availability matrix**

{{.ParamsTable}}{{ end }}