# Namespace: SP
## Entity Type: RecycleBinItem

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
DeletedDate (Edm.DateTime) | ✔ | ✔ | ✔ | ✔
DirName (Edm.String) | ✔ | ✔ | ✔ | ✔
ItemState (Edm.Int32) | ✔ | ✔ | ✔ | ✔
AuthorEmail (Edm.String) | ✔ | ✔ | ✔ | ✖
DeletedByEmail (Edm.String) | ✔ | ✔ | ✔ | ✖
DeletedByName (Edm.String) | ✔ | ✔ | ✔ | ✖
LeafName (Edm.String) | ✔ | ✔ | ✔ | ✔
LeafNamePath (SP.ResourcePath) | ✔ | ✔ | ✖ | ✖
Size (Edm.Int64) | ✔ | ✔ | ✔ | ✔
Title (Edm.String) | ✔ | ✔ | ✔ | ✔
ItemType (Edm.Int32) | ✔ | ✔ | ✔ | ✔
AuthorName (Edm.String) | ✔ | ✔ | ✔ | ✖
DeletedDateLocalFormatted (Edm.String) | ✔ | ✔ | ✔ | ✖
DirNamePath (SP.ResourcePath) | ✔ | ✔ | ✖ | ✖
Id (Edm.Guid) | ✔ | ✔ | ✔ | ✔

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
Author | ✔ | ✔ | ✔ | ✔
DeletedBy | ✔ | ✔ | ✔ | ✔
