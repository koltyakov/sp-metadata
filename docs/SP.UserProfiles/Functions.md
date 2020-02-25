# Namespace: SP.UserProfiles

## Functions Imports

**Availability matrix**

Functions Imports | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------
AmIFollowedBy (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
AmIFollowing (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
CreatePersonalSite (SP.UserProfiles.UserProfile) | ✅ | ✅ | ✅ | ✅
CreatePersonalSiteEnque (SP.UserProfiles.UserProfile) | ✅ | ✅ | ✅ | ✅
CreatePersonalSiteEnqueueBulk (SP.UserProfiles.ProfileLoader) | ✅ | ✅ | ✅ | ❌
CreatePersonalSiteFromWorkItem (SP.UserProfiles.UserProfile) | ✅ | ✅ | ✅ | ❌
DeleteCacheItemsAsync (SP.UserProfiles.PersonalCache) | ✅ | ✅ | ❌ | ❌
DeleteCacheItemsAsync2 (SP.UserProfiles.PersonalCache) | ✅ | ❌ | ❌ | ❌
FindAndUpdateFollowedGroup (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ❌
FindAndUpdateFollowedItem (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ✅
Follow (SP.UserProfiles.FollowedContent) | ❌ | ❌ | ❌ | ✅
Follow (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
FollowItem (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ✅
FollowTag (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetDefaultDocumentLibrary (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ❌
GetFollowedStatus (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ✅
GetFollowedTags (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetFollowersFor (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetGroups (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ❌
GetItem (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ✅
GetItems (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ✅
GetMyFollowers (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetMyProperties (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetMySuggestions (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetPeopleFollowedBy (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetPeopleFollowedByMe (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetPropertiesFor (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
GetPropertyNames (SP.UserProfiles.UserProfilePropertiesForUser) | ✅ | ✅ | ✅ | ✅
GetSPUserInformation (SP.UserProfiles.PeopleManager) | ✅ | ❌ | ❌ | ❌
GetUserOneDriveQuotaMax (SP.UserProfiles.PeopleManager) | ✅ (❌) | ❌ | ❌ | ❌
GetUserProfile (SP.UserProfiles.ProfileLoader) | ✅ | ✅ | ✅ | ✅
GetUserProfileProperties (SP.UserProfiles.PeopleManager) | ✅ | ❌ | ❌ | ❌
GetUserProfilePropertiesFor (SP.UserProfiles.PeopleManager) | ❌ | ❌ | ❌ | ✅
GetUserProfilePropertyFor (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
HardDeleteUserProfile (SP.UserProfiles.PeopleManager) | ✅ | ❌ | ❌ | ❌
HasGroupMembershipChangedAndSyncChanges (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ❌
HideSuggestion (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
IsFollowed (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ✅
LoadUserProfile (SP.UserProfiles.PersonalCache) | ✅ | ✅ | ❌ | ❌
ReadCache (SP.UserProfiles.PersonalCache) | ✅ | ✅ | ❌ | ❌
ReadCache2 (SP.UserProfiles.PersonalCache) | ✅ | ❌ | ❌ | ❌
ReadCacheOrCreate (SP.UserProfiles.PersonalCache) | ✅ | ✅ | ❌ | ❌
ReadCacheOrCreate2 (SP.UserProfiles.PersonalCache) | ✅ | ❌ | ❌ | ❌
ReadCacheOrCreateOrderById (SP.UserProfiles.PersonalCache) | ✅ | ✅ | ❌ | ❌
ReadCacheOrCreateOrderById2 (SP.UserProfiles.PersonalCache) | ✅ | ❌ | ❌ | ❌
RefreshFollowedItem (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ✅
RemoveSPUserInformation (SP.UserProfiles.PeopleManager) | ✅ | ❌ | ❌ | ❌
ResetUserOneDriveQuotaToDefault (SP.UserProfiles.PeopleManager) | ✅ (❌) | ❌ | ❌ | ❌
[SP_UserProfiles_CrossGeoSync_ReadFullChangesBatch](./Functions/SP_UserProfiles_CrossGeoSync_ReadFullChangesBatch.md) | ✅ | ❌ | ❌ | ❌
[<span title="SP_UserProfiles_CrossGeoSync_ReadIncrementalChangesBatch">SP_UserProfiles_CrossGeoSync_ReadIncrementalChange...</span> (SP UserProfiles CrossGeoSync ReadIncrementalChangesBatch)](./Functions/SP_UserProfiles_CrossGeoSync_ReadIncrementalChangesBatch.md) | ✅ | ❌ | ❌ | ❌
SP_UserProfiles_FollowedContent | ✅ | ✅ | ✅ | ✅
[SP_UserProfiles_FollowedItemData](./Functions/SP_UserProfiles_FollowedItemData.md) | ✅ | ✅ | ✅ | ✅
SP_UserProfiles_PeopleManager | ✅ | ✅ | ✅ | ✅
SP_UserProfiles_PeopleManager_GetTrendingTags | ✅ | ✅ | ✅ | ✅
[SP_UserProfiles_PeopleManager_IsFollowing](./Functions/SP_UserProfiles_PeopleManager_IsFollowing.md) | ✅ | ✅ | ✅ | ✅
SP_UserProfiles_PersonalCache | ✅ | ✅ | ❌ | ❌
SP_UserProfiles_ProfileImageStore | ✅ | ✅ | ✅ | ✅
SP_UserProfiles_ProfileLoader_GetOwnerUserProfile | ✅ | ✅ | ✅ | ❌
SP_UserProfiles_ProfileLoader_GetProfileLoader | ✅ | ✅ | ✅ | ✅
[SP_UserProfiles_UserProfilePropertiesForUser](./Functions/SP_UserProfiles_UserProfilePropertiesForUser.md) | ✅ | ✅ | ✅ | ✅
[<span title="SP_UserProfiles_UserProfile_CreatePersonalSiteSyncFromWorkItem">SP_UserProfiles_UserProfile_CreatePersonalSiteSync...</span> (SP UserProfiles UserProfile CreatePersonalSiteSyncFromWorkItem)](./Functions/SP_UserProfiles_UserProfile_CreatePersonalSiteSyncFromWorkItem.md) | ✅ | ✅ | ✅ | ❌
SaveUploadedFile (SP.UserProfiles.ProfileImageStore) | ✅ | ✅ | ✅ | ✅
SetItemPinState (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ❌
SetMultiValuedProfileProperty (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ❌
SetMyProfilePicture (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
SetMySiteFirstRunExperience (SP.UserProfiles.UserProfile) | ✅ | ✅ | ✅ | ❌
SetSingleValueProfileProperty (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ❌
SetUserOneDriveQuota (SP.UserProfiles.PeopleManager) | ✅ (❌) | ❌ | ❌ | ❌
ShareAllSocialData (SP.UserProfiles.UserProfile) | ✅ | ✅ | ✅ | ✅
StopFollowing (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ✅
StopFollowing (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
StopFollowingTag (SP.UserProfiles.PeopleManager) | ✅ | ✅ | ✅ | ✅
UpdateData (SP.UserProfiles.FollowedContent) | ❌ | ❌ | ❌ | ✅
UpdateFollowedGroupForUser (SP.UserProfiles.FollowedContent) | ✅ | ✅ | ✅ | ❌
WriteCache (SP.UserProfiles.PersonalCache) | ✅ | ✅ | ❌ | ❌
WriteCache2 (SP.UserProfiles.PersonalCache) | ✅ | ❌ | ❌ | ❌