# Namespace: SP
## Entity Type: ObjectSharingInformation

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
AnonymousEditLink (Edm.String) | ✔ | ✔ | ✔ | ✔
CanBeShared (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
AnonymousViewLink (Edm.String) | ✔ | ✔ | ✔ | ✔
CanBeUnshared (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
IsFolder (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
HasPendingAccessRequests (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsSharedWithGuest (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsSharedWithMany (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
CanManagePermissions (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
HasPermissionLevels (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsSharedWithCurrentUser (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsSharedWithSecurityGroup (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
PendingAccessRequestsLink (Edm.String) | ✔ | ✔ | ✔ | ✔
SharingLinks (Collection(SP.SharingLinkInfo)) | ✔ | ✔ | ✔ | ✖
TotalFileCount (Edm.Int64) | ✔ | ✔ | ✔ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
SharedWithUsersCollection | ✔ | ✔ | ✔ | ✖
