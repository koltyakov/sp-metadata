# Entity Type: SharingResult

> Namespace: SP

### Properties

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------:
ErrorMessage (Edm.String) | ✅ | ✅ | ✅ | ❌
IconUrl (Edm.String) | ✅ | ✅ | ✅ | ❌
InvitedUsers (Collection(SP.SPInvitationCreationResult)) | ✅ | ✅ | ✅ | ❌
Name (Edm.String) | ✅ | ✅ | ✅ | ❌
PermissionsPageRelativeUrl (Edm.String) | ✅ | ✅ | ✅ | ❌
StatusCode (Edm.Int32) | ✅ | ✅ | ✅ | ❌
UniquelyPermissionedUsers (Collection(SP.Sharing.UserSharingResult)) | ✅ | ✅ | ✅ | ❌
Url (Edm.String) | ✅ | ✅ | ✅ | ❌
UsersAddedToGroup (Collection(SP.Sharing.UserSharingResult)) | ✅ | ✅ | ✅ | ❌

### Navigation Properties

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------:
GroupUsersAddedTo | ✅ | ✅ | ✅ | ❌
GroupsSharedWith | ✅ | ✅ | ✅ | ❌
UsersWithAccessRequests | ✅ | ✅ | ✅ | ❌