# Namespace: SP

## Entity Type: ObjectSharingInformation

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
AnonymousEditLink (Edm.String) | ✅ | ✅ | ✅ | ✅
AnonymousViewLink (Edm.String) | ✅ | ✅ | ✅ | ✅
CanBeShared (Edm.Boolean) | ✅ | ✅ | ✅ | ❌
CanBeUnshared (Edm.Boolean) | ✅ | ✅ | ✅ | ❌
CanManagePermissions (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
HasPendingAccessRequests (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
HasPermissionLevels (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
IsFolder (Edm.Boolean) | ✅ | ✅ | ✅ | ❌
IsSharedWithCurrentUser (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
IsSharedWithGuest (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
IsSharedWithMany (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
IsSharedWithSecurityGroup (Edm.Boolean) | ✅ | ✅ | ✅ | ✅
PendingAccessRequestsLink (Edm.String) | ✅ | ✅ | ✅ | ✅
SharingLinks (Collection(SP.SharingLinkInfo)) | ✅ | ✅ | ✅ | ❌
TotalFileCount (Edm.Int64) | ✅ | ✅ | ✅ | ❌

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
SharedWithUsersCollection | ✅ | ✅ | ✅ | ❌
