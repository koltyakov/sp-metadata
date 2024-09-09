# Entity Type: RecycleBinItem

> Namespace: SP

### Properties

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------:
AuthorEmail (Edm.String) | ✅ | ✅ | ✅ | ❌
AuthorName (Edm.String) | ✅ | ✅ | ✅ | ❌
DeletedByEmail (Edm.String) | ✅ | ✅ | ✅ | ❌
DeletedByName (Edm.String) | ✅ | ✅ | ✅ | ❌
DeletedDate (Edm.DateTime) | ✅ | ✅ | ✅ | ✅
DeletedDateLocalFormatted (Edm.String) | ✅ | ✅ | ✅ | ❌
DirName (Edm.String) | ✅ | ✅ | ✅ | ✅
DirNamePath (SP.ResourcePath) | ✅ | ✅ | ❌ | ❌
Id (Edm.Guid) | ✅ | ✅ | ✅ | ✅
ItemState (Edm.Int32) | ✅ | ✅ | ✅ | ✅
ItemType (Edm.Int32) | ✅ | ✅ | ✅ | ✅
LeafName (Edm.String) | ✅ | ✅ | ✅ | ✅
LeafNamePath (SP.ResourcePath) | ✅ | ✅ | ❌ | ❌
Size (Edm.Int64) | ✅ | ✅ | ✅ | ✅
Title (Edm.String) | ✅ | ✅ | ✅ | ✅
UniqueId (Edm.Guid) | ✅ | ❌ | ❌ | ❌

### Navigation Properties

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------:
Author | ✅ | ✅ | ✅ | ✅
DeletedBy | ✅ | ✅ | ✅ | ✅