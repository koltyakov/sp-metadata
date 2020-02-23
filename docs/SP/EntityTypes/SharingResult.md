# Namespace: SP
## Entity Type: SharingResult

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
ErrorMessage (Edm.String) | ✔ | ✔ | ✔ | ✖
InvitedUsers (Collection(SP.SPInvitationCreationResult)) | ✔ | ✔ | ✔ | ✖
Name (Edm.String) | ✔ | ✔ | ✔ | ✖
UsersAddedToGroup (Collection(SP.Sharing.UserSharingResult)) | ✔ | ✔ | ✔ | ✖
IconUrl (Edm.String) | ✔ | ✔ | ✔ | ✖
PermissionsPageRelativeUrl (Edm.String) | ✔ | ✔ | ✔ | ✖
StatusCode (Edm.Int32) | ✔ | ✔ | ✔ | ✖
UniquelyPermissionedUsers (Collection(SP.Sharing.UserSharingResult)) | ✔ | ✔ | ✔ | ✖
Url (Edm.String) | ✔ | ✔ | ✔ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
GroupsSharedWith | ✔ | ✔ | ✔ | ✖
GroupUsersAddedTo | ✔ | ✔ | ✔ | ✖
UsersWithAccessRequests | ✔ | ✔ | ✔ | ✖
