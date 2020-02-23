# Namespace: SP
## Entity Type: ObjectSharingInformation

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
PendingAccessRequestsLink (Edm.String) | ✔ | ✔ | ✔ | ✔
CanManagePermissions (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
HasPendingAccessRequests (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsFolder (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
IsSharedWithGuest (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsSharedWithMany (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
AnonymousEditLink (Edm.String) | ✔ | ✔ | ✔ | ✔
HasPermissionLevels (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsSharedWithCurrentUser (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
SharingLinks (Collection(SP.SharingLinkInfo)) | ✔ | ✔ | ✔ | ✖
AnonymousViewLink (Edm.String) | ✔ | ✔ | ✔ | ✔
CanBeShared (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
CanBeUnshared (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
IsSharedWithSecurityGroup (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
TotalFileCount (Edm.Int64) | ✔ | ✔ | ✔ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
SharedWithUsersCollection | ✔ | ✔ | ✔ | ✖
