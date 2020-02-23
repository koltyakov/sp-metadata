# Namespace: SP
## Entity Type: SharingResult

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
ErrorMessage (Edm.String) | ✔ | ✔ | ✔ | ✖
Name (Edm.String) | ✔ | ✔ | ✔ | ✖
StatusCode (Edm.Int32) | ✔ | ✔ | ✔ | ✖
UsersAddedToGroup (Collection(SP.Sharing.UserSharingResult)) | ✔ | ✔ | ✔ | ✖
IconUrl (Edm.String) | ✔ | ✔ | ✔ | ✖
InvitedUsers (Collection(SP.SPInvitationCreationResult)) | ✔ | ✔ | ✔ | ✖
PermissionsPageRelativeUrl (Edm.String) | ✔ | ✔ | ✔ | ✖
UniquelyPermissionedUsers (Collection(SP.Sharing.UserSharingResult)) | ✔ | ✔ | ✔ | ✖
Url (Edm.String) | ✔ | ✔ | ✔ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
GroupsSharedWith | ✔ | ✔ | ✔ | ✖
GroupUsersAddedTo | ✔ | ✔ | ✔ | ✖
UsersWithAccessRequests | ✔ | ✔ | ✔ | ✖
