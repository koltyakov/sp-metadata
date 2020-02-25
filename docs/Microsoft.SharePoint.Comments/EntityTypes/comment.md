# Namespace: Microsoft.SharePoint.Comments

## Entity Type: comment

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------:
author (SP.Sharing.Principal) | ✅ | ✅ | ❌ | ❌
createdDate (Edm.DateTime) | ✅ | ✅ | ❌ | ❌
id (Edm.String) | ✅ | ✅ | ❌ | ❌
isLikedByUser (Edm.Boolean) | ✅ | ✅ | ❌ | ❌
isReply (Edm.Boolean) | ✅ | ✅ | ❌ | ❌
itemId (Edm.Int32) | ✅ | ✅ | ❌ | ❌
likeCount (Edm.Int32) | ✅ | ✅ | ❌ | ❌
listId (Edm.Guid) | ✅ | ✅ | ❌ | ❌
mentions (Collection(Microsoft.SharePoint.Comments.Client.Identity)) | ✅ | ✅ | ❌ | ❌
parentId (Edm.String) | ✅ | ✅ | ❌ | ❌
relativeCreatedDate (Edm.String) | ✅ | ❌ | ❌ | ❌
replyCount (Edm.Int32) | ✅ | ✅ | ❌ | ❌
text (Edm.String) | ✅ | ✅ | ❌ | ❌

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------:
likedBy | ✅ | ✅ | ❌ | ❌
replies | ✅ | ✅ | ❌ | ❌