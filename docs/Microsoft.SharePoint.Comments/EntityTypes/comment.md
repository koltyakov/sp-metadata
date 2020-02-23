# Namespace: Microsoft.SharePoint.Comments
## Entity Type: comment

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
id (Edm.String) | ✔ | ✔ | ✖ | ✖
isLikedByUser (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
mentions (Collection(Microsoft.SharePoint.Comments.Client.Identity)) | ✔ | ✔ | ✖ | ✖
parentId (Edm.String) | ✔ | ✔ | ✖ | ✖
relativeCreatedDate (Edm.String) | ✔ | ✖ | ✖ | ✖
text (Edm.String) | ✔ | ✔ | ✖ | ✖
author (SP.Sharing.Principal) | ✔ | ✔ | ✖ | ✖
createdDate (Edm.DateTime) | ✔ | ✔ | ✖ | ✖
isReply (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
itemId (Edm.Int32) | ✔ | ✔ | ✖ | ✖
likeCount (Edm.Int32) | ✔ | ✔ | ✖ | ✖
listId (Edm.Guid) | ✔ | ✔ | ✖ | ✖
replyCount (Edm.Int32) | ✔ | ✔ | ✖ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
likedBy | ✔ | ✔ | ✖ | ✖
replies | ✔ | ✔ | ✖ | ✖
