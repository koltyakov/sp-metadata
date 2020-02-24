# Metadata

The folder contains EDMX models exported from the corresponding environments of different SharePoint versions:

File | Description
-----|------------
[spo.xml](./spo.xml) | SharePoint Online (Standard Release)
[spo.target.xml](./spo.target.xml) | SharePoint Online (Target Release)
[2019.xml](./2019.xml) | SharePoint 2019 (On-Premise)
[2016.xml](./2016.xml) | SharePoint 2016 (On-Premise)
[2013.xml](./2013.xml) | SharePoint 2013 (On-Premise)

SPO schemas are tracked on schedule basis and updated automatically in the repository.

On-Premise schemas, especially for the old versions, are usually not change at all.

To manually update a scheme (in a fork or clone), e.g. to verify if on a moment of time an environment(s) you use is identical or not API model-wise, create one or many connection files in `./config` folder ([see more](../config/README.md)) then run `go run ./cmd/get`.