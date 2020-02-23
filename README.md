# SharePoint Metadata Tracker

The Metadata Tracker collects SharePoint API entity definitions (`/_api/$metadata`).

Entity definitions (`edmx` models) are stored in `./meta` folder.

Schema comparison allows tracking "what's new in SharePoint Online" in comparison to previous releases and On-Premise versions.

SharePoint Online is tracked for Standard and Target Releases which allows detecting what potential changes API-terms would be applied shortly.

## APIs drill-down entry point

Check out [REST APIs Namespaces](./docs/Namespaces.md) availability then drill-down to a specific API.

## Similar projects

In a contract to [SharePoint REST API Metadata Explorer](https://github.com/s-KaiNet/sp-rest-explorer) by [Sergei Sergeev](https://github.com/s-KaiNet), SharePoint Metadata Tracker aims to compare APIs availability with previous versions.