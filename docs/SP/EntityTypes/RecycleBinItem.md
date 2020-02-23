# Namespace: SP
## Entity Type: RecycleBinItem

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
DeletedDate (Edm.DateTime) | ✔ | ✔ | ✔ | ✔
LeafName (Edm.String) | ✔ | ✔ | ✔ | ✔
Title (Edm.String) | ✔ | ✔ | ✔ | ✔
AuthorName (Edm.String) | ✔ | ✔ | ✔ | ✖
DeletedByEmail (Edm.String) | ✔ | ✔ | ✔ | ✖
DeletedDateLocalFormatted (Edm.String) | ✔ | ✔ | ✔ | ✖
ItemType (Edm.Int32) | ✔ | ✔ | ✔ | ✔
Size (Edm.Int64) | ✔ | ✔ | ✔ | ✔
DirName (Edm.String) | ✔ | ✔ | ✔ | ✔
DirNamePath (SP.ResourcePath) | ✔ | ✔ | ✖ | ✖
AuthorEmail (Edm.String) | ✔ | ✔ | ✔ | ✖
DeletedByName (Edm.String) | ✔ | ✔ | ✔ | ✖
Id (Edm.Guid) | ✔ | ✔ | ✔ | ✔
ItemState (Edm.Int32) | ✔ | ✔ | ✔ | ✔
LeafNamePath (SP.ResourcePath) | ✔ | ✔ | ✖ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
Author | ✔ | ✔ | ✔ | ✔
DeletedBy | ✔ | ✔ | ✔ | ✔
