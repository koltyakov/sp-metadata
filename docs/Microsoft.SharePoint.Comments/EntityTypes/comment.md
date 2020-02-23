# Namespace: Microsoft.SharePoint.Comments
## Entity Type: comment

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
isReply (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
itemId (Edm.Int32) | ✔ | ✔ | ✖ | ✖
likeCount (Edm.Int32) | ✔ | ✔ | ✖ | ✖
text (Edm.String) | ✔ | ✔ | ✖ | ✖
mentions (Collection(Microsoft.SharePoint.Comments.Client.Identity)) | ✔ | ✔ | ✖ | ✖
parentId (Edm.String) | ✔ | ✔ | ✖ | ✖
relativeCreatedDate (Edm.String) | ✔ | ✖ | ✖ | ✖
author (SP.Sharing.Principal) | ✔ | ✔ | ✖ | ✖
createdDate (Edm.DateTime) | ✔ | ✔ | ✖ | ✖
id (Edm.String) | ✔ | ✔ | ✖ | ✖
isLikedByUser (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
listId (Edm.Guid) | ✔ | ✔ | ✖ | ✖
replyCount (Edm.Int32) | ✔ | ✔ | ✖ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
likedBy | ✔ | ✔ | ✖ | ✖
replies | ✔ | ✔ | ✖ | ✖
