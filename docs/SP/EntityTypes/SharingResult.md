# Namespace: SP
## Entity Type: SharingResult

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
Name (Edm.String) | ✔ | ✔ | ✔ | ✖
IconUrl (Edm.String) | ✔ | ✔ | ✔ | ✖
InvitedUsers (Collection(SP.SPInvitationCreationResult)) | ✔ | ✔ | ✔ | ✖
PermissionsPageRelativeUrl (Edm.String) | ✔ | ✔ | ✔ | ✖
StatusCode (Edm.Int32) | ✔ | ✔ | ✔ | ✖
UniquelyPermissionedUsers (Collection(SP.Sharing.UserSharingResult)) | ✔ | ✔ | ✔ | ✖
Url (Edm.String) | ✔ | ✔ | ✔ | ✖
UsersAddedToGroup (Collection(SP.Sharing.UserSharingResult)) | ✔ | ✔ | ✔ | ✖
ErrorMessage (Edm.String) | ✔ | ✔ | ✔ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
UsersWithAccessRequests | ✔ | ✔ | ✔ | ✖
GroupsSharedWith | ✔ | ✔ | ✔ | ✖
GroupUsersAddedTo | ✔ | ✔ | ✔ | ✖
