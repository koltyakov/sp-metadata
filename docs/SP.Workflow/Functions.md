# Namespace: SP.Workflow

## Functions Imports

**Availability matrix**

Functions Imports | SPO | SP 2019 | SP 2016 | SP 2013
----------|:---:|:-------:|:-------:|:-------
Add (Collection(SP.Workflow.WorkflowAssociation)) | ✅ | ✅ | ✅ | ✅
BreakRoleInheritance (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
CancelWorkflow (SP.WorkflowServices.InteropService) | ✅ | ✅ | ✅ | ✅
CancelWorkflow (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
CountInstances (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
CountInstancesWithStatus (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
DeleteCollateral (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
DeleteDefinition (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
DeleteObject (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
DeleteObject (SP.Workflow.WorkflowAssociation) | ✅ | ✅ | ✅ | ✅
DeleteSubscription (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ✅
DeleteWithParameters (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
DeprecateDefinition (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
DisableEvents (SP.WorkflowServices.InteropService) | ✅ | ✅ | ✅ | ✅
EnableEvents (SP.WorkflowServices.InteropService) | ✅ | ✅ | ✅ | ✅
Enumerate (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
EnumerateDefinitions (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
EnumerateInstancesForListItem (SP.WorkflowServices.WorkflowInstanceService) | ✅ | ✅ | ✅ | ✅
EnumerateInstancesForListItemWithOffset (SP.WorkflowServices.WorkflowInstanceService) | ✅ | ✅ | ✅ | ✅
EnumerateInstancesForSite (SP.WorkflowServices.WorkflowInstanceService) | ✅ | ✅ | ✅ | ✅
EnumerateInstancesForSiteWithOffset (SP.WorkflowServices.WorkflowInstanceService) | ✅ | ✅ | ✅ | ✅
EnumerateIntegratedApps (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ❌
EnumerateSubscriptions (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ✅
EnumerateSubscriptionsByDefinition (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ✅
EnumerateSubscriptionsByEventSource (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ✅
EnumerateSubscriptionsByList (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ✅
EnumerateSubscriptionsByListAndParentContentType (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ❌
EnumerateSubscriptionsByListWithContentType (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ❌
EnumerateWithOffset (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
GetActivitySignatures (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
GetById (Collection(SP.Workflow.WorkflowAssociation)) | ✅ | ✅ | ✅ | ✅
GetById (Collection(SP.Workflow.WorkflowTemplate)) | ✅ | ✅ | ✅ | ✅
GetByName (Collection(SP.Workflow.WorkflowAssociation)) | ✅ | ✅ | ✅ | ✅
GetByName (Collection(SP.Workflow.WorkflowTemplate)) | ✅ | ✅ | ✅ | ✅
GetChanges (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
GetCollateralUri (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
GetDebugInfo (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
GetDefinition (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
GetDesignerActions (SP.WorkflowServices.WorkflowDeploymentService) | ❌ | ❌ | ❌ | ✅
GetExternalVariable (SP.WorkflowServices.WorkflowSubscription) | ✅ | ✅ | ✅ | ✅
GetHashtags (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
GetInstance (SP.WorkflowServices.WorkflowInstanceService) | ✅ | ✅ | ✅ | ✅
GetSubscription (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ✅
GetUserEffectivePermissions (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
GetWOPIFrameUrl (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
GetWorkflowDeploymentService (SP.WorkflowServices.WorkflowServicesManager) | ✅ | ✅ | ✅ | ✅
GetWorkflowInstanceService (SP.WorkflowServices.WorkflowServicesManager) | ✅ | ✅ | ✅ | ✅
GetWorkflowInteropService (SP.WorkflowServices.WorkflowServicesManager) | ✅ | ✅ | ✅ | ✅
GetWorkflowSubscriptionService (SP.WorkflowServices.WorkflowServicesManager) | ✅ | ✅ | ✅ | ✅
IsIntegratedApp (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ❌
IsIntegratedApp (SP.WorkflowServices.WorkflowServicesManager) | ✅ | ✅ | ✅ | ❌
MediaServiceUpdate (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
MediaServiceUpdateV2 (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
OverridePolicyTip (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
PackageDefinition (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
ParseAndSetFieldValue (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
PublishCustomEvent (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
PublishDefinition (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
PublishEvent (SP.WorkflowServices.WorkflowMessagingService) | ✅ | ❌ | ❌ | ❌
PublishSubscription (SP.WorkflowServices.WorkflowSubscriptionService) | ❌ | ❌ | ❌ | ✅
PublishSubscriptionForList (SP.WorkflowServices.WorkflowSubscriptionService) | ❌ | ❌ | ❌ | ✅
Recycle (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
RecycleWithParameters (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
RegisterInterestInHostWebList (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ❌
RegisterInterestInList (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ✅
ResetRoleInheritance (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
ResumeWorkflow (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
SP_WorkflowServices_InteropService_Current | ✅ | ✅ | ✅ | ✅
SP_WorkflowServices_WorkflowDefinition | ✅ | ✅ | ✅ | ✅
<span title="SP_WorkflowServices_WorkflowInstanceService_Current">SP_WorkflowServices_WorkflowInstanceService_Curren...</span> (SP WorkflowServices WorkflowInstanceService Current) | ✅ | ✅ | ✅ | ✅
<span title="SP_WorkflowServices_WorkflowServicesManager_Current">SP_WorkflowServices_WorkflowServicesManager_Curren...</span> (SP WorkflowServices WorkflowServicesManager Current) | ✅ | ✅ | ✅ | ✅
SP_WorkflowServices_WorkflowSubscription | ✅ | ✅ | ✅ | ✅
<span title="SP_WorkflowServices_WorkflowSubscriptionService_Current">SP_WorkflowServices_WorkflowSubscriptionService_Cu...</span> (SP WorkflowServices WorkflowSubscriptionService Current) | ✅ | ✅ | ✅ | ✅
SaveCollateral (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
SaveDefinition (SP.WorkflowServices.WorkflowDeploymentService) | ❌ | ❌ | ❌ | ✅
SetCommentsDisabled (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
SetComplianceTag (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
SetComplianceTagWithExplicitMetasUpdate (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
SetComplianceTagWithHold (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
SetComplianceTagWithMetaInfo (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
SetComplianceTagWithNoHold (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
SetComplianceTagWithRecord (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
SetExternalVariable (SP.WorkflowServices.WorkflowSubscription) | ✅ | ✅ | ✅ | ✅
SetProperty (SP.WorkflowServices.WorkflowDefinition) | ✅ | ✅ | ✅ | ✅
SetProperty (SP.WorkflowServices.WorkflowSubscription) | ✅ | ✅ | ✅ | ✅
Sort (Collection(SP.WorkflowServices.WorkflowDefinition)) | ✅ | ✅ | ✅ | ❌
Sort (Collection(SP.WorkflowServices.WorkflowSubscription)) | ✅ | ✅ | ✅ | ❌
StartWorkflow (SP.WorkflowServices.InteropService) | ✅ | ✅ | ✅ | ✅
StartWorkflow (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
StartWorkflowOnListItemBySubscriptionId (SP.WorkflowServices.WorkflowInstanceService) | ✅ | ✅ | ✅ | ✅
SuspendWorkflow (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
SystemUpdate (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
TerminateWorkflow (SP.WorkflowServices.WorkflowInstanceService) | ❌ | ❌ | ❌ | ✅
UnregisterInterestInHostWebList (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ❌
UnregisterInterestInList (SP.WorkflowServices.WorkflowSubscriptionService) | ✅ | ✅ | ✅ | ✅
Update (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
Update (SP.Workflow.WorkflowAssociation) | ✅ | ✅ | ✅ | ✅
UpdateEx (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
UpdateHashtags (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
UpdateOverwriteVersion (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
ValidateActivity (SP.WorkflowServices.WorkflowDeploymentService) | ✅ | ✅ | ✅ | ✅
ValidateUpdateListItem (SP.Workflow.SPWorkflowTask) | ✅ | ❌ | ❌ | ❌
