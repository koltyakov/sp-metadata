# Namespace: SP
## Entity Type: ObjectSharingInformation

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
HasPendingAccessRequests (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
HasPermissionLevels (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsSharedWithCurrentUser (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsSharedWithMany (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
TotalFileCount (Edm.Int64) | ✔ | ✔ | ✔ | ✖
AnonymousEditLink (Edm.String) | ✔ | ✔ | ✔ | ✔
CanBeShared (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
CanBeUnshared (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
IsSharedWithSecurityGroup (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
AnonymousViewLink (Edm.String) | ✔ | ✔ | ✔ | ✔
CanManagePermissions (Edm.Boolean) | ✔ | ✔ | ✔ | ✔
IsFolder (Edm.Boolean) | ✔ | ✔ | ✔ | ✖
PendingAccessRequestsLink (Edm.String) | ✔ | ✔ | ✔ | ✔
SharingLinks (Collection(SP.SharingLinkInfo)) | ✔ | ✔ | ✔ | ✖
IsSharedWithGuest (Edm.Boolean) | ✔ | ✔ | ✔ | ✔

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
SharedWithUsersCollection | ✔ | ✔ | ✔ | ✖
