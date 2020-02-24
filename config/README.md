# Config

The folder contains environment configuration and connection files:

File | Description
-----|------------
private.spo.json | SharePoint Online (Standard Release) environment connection file
private.spo.target.json | SharePoint Online (Target Release) environment connection file
private.2019.json | SharePoint 2019 (On-Premise) environment connection file
private.2016.json | SharePoint 2016 (On-Premise) environment connection file
private.2013.json | SharePoint 2013 (On-Premise) environment connection file

`private.json` files are JSON auth configs representation of [gosip](https://github.com/koltyakov/gosip) SharePoint client. The only exception is that `"strategy"` property is required. Strategy is the explicit name of the auth strategy to use with the environment.

`private.json` files are excluded from the source control.

## Samples

SharePoint Online:

```json
{
  "strategy": "saml",
  "siteUrl": "https://contoso.sharepoint.com/sites/meta",
  "username": "john.doe@contoso.onmicrosoft.com",
  "password": "__secret__"
}
```

SharePoint On-Premise:

```json
{
  "strategy": "ntlm",
  "siteUrl": "https://sp.contoso.com/sites/meta",
  "domain": "domain",
  "username": "john.doe",
  "password": "__secret__"
}
```

[Other](https://go.spflow.com/auth/overview) strategies can be used as well.

If a local configuration doesn't include some of the connection files, when running metadata get command these are skipped and not updated.
