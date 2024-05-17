# Entity Type: comment

> Namespace: Microsoft.SharePoint.Comments

### Properties

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------:
author (SP.Sharing.Principal) | ✅ | ✅ | ❌ | ❌
contentAnchor (Microsoft.SharePoint.Comments.ContentAnchor) | ✅ | ❌ | ❌ | ❌
createdDate (Edm.DateTime) | ✅ | ✅ | ❌ | ❌
id (Edm.String) | ✅ | ✅ | ❌ | ❌
isLikedByUser (Edm.Boolean) | ✅ | ✅ | ❌ | ❌
isReply (Edm.Boolean) | ✅ | ✅ | ❌ | ❌
itemId (Edm.Int32) | ✅ | ✅ | ❌ | ❌
likeCount (Edm.Int32) | ✅ | ✅ | ❌ | ❌
listId (Edm.Guid) | ✅ | ✅ | ❌ | ❌
mentions (Collection(Microsoft.SharePoint.Comments.Client.Identity)) | ✅ | ✅ | ❌ | ❌
modifiedDate (Edm.DateTime) | ✅ | ❌ | ❌ | ❌
parentId (Edm.String) | ✅ | ✅ | ❌ | ❌
relativeCreatedDate (Edm.String) | ✅ | ❌ | ❌ | ❌
replyCount (Edm.Int32) | ✅ | ✅ | ❌ | ❌
text (Edm.String) | ✅ | ✅ | ❌ | ❌

### Navigation Properties

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------:
likedBy | ✅ | ✅ | ❌ | ❌
replies | ✅ | ✅ | ❌ | ❌