# Namespace: SP
## Entity Type: RecycleBinItem

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
DeletedDateLocalFormatted (Edm.String) | ✔ | ✔ | ✔ | ✖
DirNamePath (SP.ResourcePath) | ✔ | ✔ | ✖ | ✖
Id (Edm.Guid) | ✔ | ✔ | ✔ | ✔
LeafNamePath (SP.ResourcePath) | ✔ | ✔ | ✖ | ✖
DeletedByName (Edm.String) | ✔ | ✔ | ✔ | ✖
DeletedByEmail (Edm.String) | ✔ | ✔ | ✔ | ✖
DirName (Edm.String) | ✔ | ✔ | ✔ | ✔
AuthorEmail (Edm.String) | ✔ | ✔ | ✔ | ✖
ItemState (Edm.Int32) | ✔ | ✔ | ✔ | ✔
ItemType (Edm.Int32) | ✔ | ✔ | ✔ | ✔
AuthorName (Edm.String) | ✔ | ✔ | ✔ | ✖
LeafName (Edm.String) | ✔ | ✔ | ✔ | ✔
Size (Edm.Int64) | ✔ | ✔ | ✔ | ✔
Title (Edm.String) | ✔ | ✔ | ✔ | ✔
DeletedDate (Edm.DateTime) | ✔ | ✔ | ✔ | ✔

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
Author | ✔ | ✔ | ✔ | ✔
DeletedBy | ✔ | ✔ | ✔ | ✔
