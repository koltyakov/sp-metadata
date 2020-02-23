# Namespace: Microsoft.SharePoint.Comments
## Entity Type: comment

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
isReply (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
listId (Edm.Guid) | ✔ | ✔ | ✖ | ✖
mentions (Collection(Microsoft.SharePoint.Comments.Client.Identity)) | ✔ | ✔ | ✖ | ✖
parentId (Edm.String) | ✔ | ✔ | ✖ | ✖
replyCount (Edm.Int32) | ✔ | ✔ | ✖ | ✖
id (Edm.String) | ✔ | ✔ | ✖ | ✖
createdDate (Edm.DateTime) | ✔ | ✔ | ✖ | ✖
isLikedByUser (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
itemId (Edm.Int32) | ✔ | ✔ | ✖ | ✖
likeCount (Edm.Int32) | ✔ | ✔ | ✖ | ✖
relativeCreatedDate (Edm.String) | ✔ | ✖ | ✖ | ✖
text (Edm.String) | ✔ | ✔ | ✖ | ✖
author (SP.Sharing.Principal) | ✔ | ✔ | ✖ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
likedBy | ✔ | ✔ | ✖ | ✖
replies | ✔ | ✔ | ✖ | ✖
