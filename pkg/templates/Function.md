# Function Import: {{.Function}}

> Namespace: {{.Namespace}}

{{ if (ne .EntityType "") }}- Entity type: {{.EntityType}}
{{ end }}{{ if (ne .ReturnType "") }}- Return type: {{.ReturnType}}
{{ end }}{{ if (ne .EntitySet "") }}- Entity set: {{.EntitySet}}
{{ end }}- Is composable: {{.IsComposable}}
- Is bindable: {{.IsBindable}}

{{ if (ne .ParamsTable "") }}### Parameters

{{.ParamsTable}}{{ end }}