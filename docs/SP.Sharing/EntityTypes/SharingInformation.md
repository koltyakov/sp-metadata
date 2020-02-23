# Namespace: SP.Sharing
## Entity Type: SharingInformation

### Properties

**Availability matrix**

Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
showExternalSharingWarning (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
canUseSimplifiedRoles (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
sharedObjectType (Edm.Int32) | ✔ | ✔ | ✖ | ✖
shareUiUrl (Edm.String) | ✔ | ✔ | ✖ | ✖
accessRequestSettings (SP.Sharing.AccessRequestSettings) | ✔ | ✔ | ✖ | ✖
defaultShareLinkPermission (Edm.Int32) | ✔ | ✔ | ✖ | ✖
hasUniquePermissions (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
permissionsInformation (SP.Sharing.PermissionCollection) | ✔ | ✔ | ✖ | ✖
canRequestAccessForGrantAccess (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
defaultLinkKind (Edm.Int32) | ✔ | ✔ | ✖ | ✖
effectiveLimitedAccessFileType (Edm.Int32) | ✔ | ✖ | ✖ | ✖
fileExtension (Edm.String) | ✔ | ✖ | ✖ | ✖
sharingLinkTemplates (SP.Sharing.SharingLinkDefaultTemplatesCollection) | ✔ | ✖ | ✖ | ✖
addressBarLinkSettings (SP.Sharing.AddressBarLinkSettings) | ✔ | ✖ | ✖ | ✖
canAddExternalPrincipal (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
currentRole (Edm.Int32) | ✔ | ✔ | ✖ | ✖
defaultShareLinkScope (Edm.Int32) | ✔ | ✖ | ✖ | ✖
defaultShareLinkToExistingAccess (Edm.Boolean) | ✔ | ✖ | ✖ | ✖
doesUserHaveIBSegment (Edm.Boolean) | ✔ | ✖ | ✖ | ✖
microserviceShareUiUrl (Edm.String) | ✔ | ✔ | ✖ | ✖
sharingAbilities (SP.Sharing.SharingAbilities) | ✔ | ✔ | ✖ | ✖
anonymousLinkExpirationRestrictionDays (Edm.Int32) | ✔ | ✔ | ✖ | ✖
canSendEmail (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
customizedExternalSharingServiceUrl (Edm.String) | ✔ | ✖ | ✖ | ✖
webUrl (Edm.String) | ✔ | ✔ | ✖ | ✖
blockPeoplePickerAndSharing (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
directUrl (Edm.String) | ✔ | ✔ | ✖ | ✖
webTemplateId (Edm.Int32) | ✔ | ✖ | ✖ | ✖
sharingStatus (Edm.Int32) | ✔ | ✔ | ✖ | ✖
siteIBSegmentIDs (Collection(Edm.String)) | ✔ | ✖ | ✖ | ✖
canAddInternalPrincipal (Edm.Boolean) | ✔ | ✔ | ✖ | ✖
domainRestrictionSettings (SP.Sharing.DomainRestrictionSettings) | ✔ | ✔ | ✖ | ✖
itemUniqueId (Edm.Guid) | ✔ | ✖ | ✖ | ✖

### Navigation Properties

**Availability matrix**

Navigation Property | SPO | SP 2019 | SP 2016 | SP 2013
----------|-----|---------|---------|--------
pickerSettings | ✔ | ✔ | ✖ | ✖
