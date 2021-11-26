


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureBackupServerContainer struct {
	BackupManagementType *string                   `pulumi:"backupManagementType"`
	CanReRegister        *bool                     `pulumi:"canReRegister"`
	ContainerId          *string                   `pulumi:"containerId"`
	ContainerType        string                    `pulumi:"containerType"`
	DpmAgentVersion      *string                   `pulumi:"dpmAgentVersion"`
	DpmServers           []string                  `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                   `pulumi:"friendlyName"`
	HealthStatus         *string                   `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                  `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                   `pulumi:"protectionStatus"`
	RegistrationStatus   *string                   `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                     `pulumi:"upgradeAvailable"`
}





type AzureBackupServerContainerInput interface {
	pulumi.Input

	ToAzureBackupServerContainerOutput() AzureBackupServerContainerOutput
	ToAzureBackupServerContainerOutputWithContext(context.Context) AzureBackupServerContainerOutput
}

type AzureBackupServerContainerArgs struct {
	BackupManagementType pulumi.StringPtrInput            `pulumi:"backupManagementType"`
	CanReRegister        pulumi.BoolPtrInput              `pulumi:"canReRegister"`
	ContainerId          pulumi.StringPtrInput            `pulumi:"containerId"`
	ContainerType        pulumi.StringInput               `pulumi:"containerType"`
	DpmAgentVersion      pulumi.StringPtrInput            `pulumi:"dpmAgentVersion"`
	DpmServers           pulumi.StringArrayInput          `pulumi:"dpmServers"`
	ExtendedInfo         DPMContainerExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput            `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput            `pulumi:"healthStatus"`
	ProtectedItemCount   pulumi.Float64PtrInput           `pulumi:"protectedItemCount"`
	ProtectionStatus     pulumi.StringPtrInput            `pulumi:"protectionStatus"`
	RegistrationStatus   pulumi.StringPtrInput            `pulumi:"registrationStatus"`
	UpgradeAvailable     pulumi.BoolPtrInput              `pulumi:"upgradeAvailable"`
}

func (AzureBackupServerContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBackupServerContainer)(nil)).Elem()
}

func (i AzureBackupServerContainerArgs) ToAzureBackupServerContainerOutput() AzureBackupServerContainerOutput {
	return i.ToAzureBackupServerContainerOutputWithContext(context.Background())
}

func (i AzureBackupServerContainerArgs) ToAzureBackupServerContainerOutputWithContext(ctx context.Context) AzureBackupServerContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBackupServerContainerOutput)
}

type AzureBackupServerContainerOutput struct{ *pulumi.OutputState }

func (AzureBackupServerContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBackupServerContainer)(nil)).Elem()
}

func (o AzureBackupServerContainerOutput) ToAzureBackupServerContainerOutput() AzureBackupServerContainerOutput {
	return o
}

func (o AzureBackupServerContainerOutput) ToAzureBackupServerContainerOutputWithContext(ctx context.Context) AzureBackupServerContainerOutput {
	return o
}

func (o AzureBackupServerContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerOutput) CanReRegister() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *bool { return v.CanReRegister }).(pulumi.BoolPtrOutput)
}

func (o AzureBackupServerContainerOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *string { return v.ContainerId }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureBackupServerContainerOutput) DpmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *string { return v.DpmAgentVersion }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerOutput) DpmServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) []string { return v.DpmServers }).(pulumi.StringArrayOutput)
}

func (o AzureBackupServerContainerOutput) ExtendedInfo() DPMContainerExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *DPMContainerExtendedInfo { return v.ExtendedInfo }).(DPMContainerExtendedInfoPtrOutput)
}

func (o AzureBackupServerContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerOutput) ProtectedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *float64 { return v.ProtectedItemCount }).(pulumi.Float64PtrOutput)
}

func (o AzureBackupServerContainerOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerOutput) UpgradeAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainer) *bool { return v.UpgradeAvailable }).(pulumi.BoolPtrOutput)
}

type AzureBackupServerContainerResponse struct {
	BackupManagementType *string                           `pulumi:"backupManagementType"`
	CanReRegister        *bool                             `pulumi:"canReRegister"`
	ContainerId          *string                           `pulumi:"containerId"`
	ContainerType        string                            `pulumi:"containerType"`
	DpmAgentVersion      *string                           `pulumi:"dpmAgentVersion"`
	DpmServers           []string                          `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                           `pulumi:"friendlyName"`
	HealthStatus         *string                           `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                          `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                           `pulumi:"protectionStatus"`
	RegistrationStatus   *string                           `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                             `pulumi:"upgradeAvailable"`
}





type AzureBackupServerContainerResponseInput interface {
	pulumi.Input

	ToAzureBackupServerContainerResponseOutput() AzureBackupServerContainerResponseOutput
	ToAzureBackupServerContainerResponseOutputWithContext(context.Context) AzureBackupServerContainerResponseOutput
}

type AzureBackupServerContainerResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput                    `pulumi:"backupManagementType"`
	CanReRegister        pulumi.BoolPtrInput                      `pulumi:"canReRegister"`
	ContainerId          pulumi.StringPtrInput                    `pulumi:"containerId"`
	ContainerType        pulumi.StringInput                       `pulumi:"containerType"`
	DpmAgentVersion      pulumi.StringPtrInput                    `pulumi:"dpmAgentVersion"`
	DpmServers           pulumi.StringArrayInput                  `pulumi:"dpmServers"`
	ExtendedInfo         DPMContainerExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput                    `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                    `pulumi:"healthStatus"`
	ProtectedItemCount   pulumi.Float64PtrInput                   `pulumi:"protectedItemCount"`
	ProtectionStatus     pulumi.StringPtrInput                    `pulumi:"protectionStatus"`
	RegistrationStatus   pulumi.StringPtrInput                    `pulumi:"registrationStatus"`
	UpgradeAvailable     pulumi.BoolPtrInput                      `pulumi:"upgradeAvailable"`
}

func (AzureBackupServerContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBackupServerContainerResponse)(nil)).Elem()
}

func (i AzureBackupServerContainerResponseArgs) ToAzureBackupServerContainerResponseOutput() AzureBackupServerContainerResponseOutput {
	return i.ToAzureBackupServerContainerResponseOutputWithContext(context.Background())
}

func (i AzureBackupServerContainerResponseArgs) ToAzureBackupServerContainerResponseOutputWithContext(ctx context.Context) AzureBackupServerContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBackupServerContainerResponseOutput)
}

type AzureBackupServerContainerResponseOutput struct{ *pulumi.OutputState }

func (AzureBackupServerContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBackupServerContainerResponse)(nil)).Elem()
}

func (o AzureBackupServerContainerResponseOutput) ToAzureBackupServerContainerResponseOutput() AzureBackupServerContainerResponseOutput {
	return o
}

func (o AzureBackupServerContainerResponseOutput) ToAzureBackupServerContainerResponseOutputWithContext(ctx context.Context) AzureBackupServerContainerResponseOutput {
	return o
}

func (o AzureBackupServerContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) CanReRegister() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *bool { return v.CanReRegister }).(pulumi.BoolPtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *string { return v.ContainerId }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureBackupServerContainerResponseOutput) DpmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *string { return v.DpmAgentVersion }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) DpmServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) []string { return v.DpmServers }).(pulumi.StringArrayOutput)
}

func (o AzureBackupServerContainerResponseOutput) ExtendedInfo() DPMContainerExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *DPMContainerExtendedInfoResponse { return v.ExtendedInfo }).(DPMContainerExtendedInfoResponsePtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) ProtectedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *float64 { return v.ProtectedItemCount }).(pulumi.Float64PtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureBackupServerContainerResponseOutput) UpgradeAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBackupServerContainerResponse) *bool { return v.UpgradeAvailable }).(pulumi.BoolPtrOutput)
}

type AzureFileShareProtectionPolicy struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{} `pulumi:"schedulePolicy"`
	TimeZone                       *string     `pulumi:"timeZone"`
	WorkLoadType                   *string     `pulumi:"workLoadType"`
}





type AzureFileShareProtectionPolicyInput interface {
	pulumi.Input

	ToAzureFileShareProtectionPolicyOutput() AzureFileShareProtectionPolicyOutput
	ToAzureFileShareProtectionPolicyOutputWithContext(context.Context) AzureFileShareProtectionPolicyOutput
}

type AzureFileShareProtectionPolicyArgs struct {
	BackupManagementType           pulumi.StringInput      `pulumi:"backupManagementType"`
	ProtectedItemsCount            pulumi.IntPtrInput      `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                pulumi.Input            `pulumi:"retentionPolicy"`
	SchedulePolicy                 pulumi.Input            `pulumi:"schedulePolicy"`
	TimeZone                       pulumi.StringPtrInput   `pulumi:"timeZone"`
	WorkLoadType                   pulumi.StringPtrInput   `pulumi:"workLoadType"`
}

func (AzureFileShareProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileShareProtectionPolicy)(nil)).Elem()
}

func (i AzureFileShareProtectionPolicyArgs) ToAzureFileShareProtectionPolicyOutput() AzureFileShareProtectionPolicyOutput {
	return i.ToAzureFileShareProtectionPolicyOutputWithContext(context.Background())
}

func (i AzureFileShareProtectionPolicyArgs) ToAzureFileShareProtectionPolicyOutputWithContext(ctx context.Context) AzureFileShareProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileShareProtectionPolicyOutput)
}

type AzureFileShareProtectionPolicyOutput struct{ *pulumi.OutputState }

func (AzureFileShareProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileShareProtectionPolicy)(nil)).Elem()
}

func (o AzureFileShareProtectionPolicyOutput) ToAzureFileShareProtectionPolicyOutput() AzureFileShareProtectionPolicyOutput {
	return o
}

func (o AzureFileShareProtectionPolicyOutput) ToAzureFileShareProtectionPolicyOutputWithContext(ctx context.Context) AzureFileShareProtectionPolicyOutput {
	return o
}

func (o AzureFileShareProtectionPolicyOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicy) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o AzureFileShareProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureFileShareProtectionPolicyOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicy) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureFileShareProtectionPolicyOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicy) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o AzureFileShareProtectionPolicyOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicy) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

func (o AzureFileShareProtectionPolicyOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicy) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o AzureFileShareProtectionPolicyOutput) WorkLoadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicy) *string { return v.WorkLoadType }).(pulumi.StringPtrOutput)
}

type AzureFileShareProtectionPolicyResponse struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{} `pulumi:"schedulePolicy"`
	TimeZone                       *string     `pulumi:"timeZone"`
	WorkLoadType                   *string     `pulumi:"workLoadType"`
}





type AzureFileShareProtectionPolicyResponseInput interface {
	pulumi.Input

	ToAzureFileShareProtectionPolicyResponseOutput() AzureFileShareProtectionPolicyResponseOutput
	ToAzureFileShareProtectionPolicyResponseOutputWithContext(context.Context) AzureFileShareProtectionPolicyResponseOutput
}

type AzureFileShareProtectionPolicyResponseArgs struct {
	BackupManagementType           pulumi.StringInput      `pulumi:"backupManagementType"`
	ProtectedItemsCount            pulumi.IntPtrInput      `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                pulumi.Input            `pulumi:"retentionPolicy"`
	SchedulePolicy                 pulumi.Input            `pulumi:"schedulePolicy"`
	TimeZone                       pulumi.StringPtrInput   `pulumi:"timeZone"`
	WorkLoadType                   pulumi.StringPtrInput   `pulumi:"workLoadType"`
}

func (AzureFileShareProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileShareProtectionPolicyResponse)(nil)).Elem()
}

func (i AzureFileShareProtectionPolicyResponseArgs) ToAzureFileShareProtectionPolicyResponseOutput() AzureFileShareProtectionPolicyResponseOutput {
	return i.ToAzureFileShareProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i AzureFileShareProtectionPolicyResponseArgs) ToAzureFileShareProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureFileShareProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileShareProtectionPolicyResponseOutput)
}

type AzureFileShareProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (AzureFileShareProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileShareProtectionPolicyResponse)(nil)).Elem()
}

func (o AzureFileShareProtectionPolicyResponseOutput) ToAzureFileShareProtectionPolicyResponseOutput() AzureFileShareProtectionPolicyResponseOutput {
	return o
}

func (o AzureFileShareProtectionPolicyResponseOutput) ToAzureFileShareProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureFileShareProtectionPolicyResponseOutput {
	return o
}

func (o AzureFileShareProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicyResponse) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o AzureFileShareProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureFileShareProtectionPolicyResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicyResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureFileShareProtectionPolicyResponseOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicyResponse) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o AzureFileShareProtectionPolicyResponseOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicyResponse) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

func (o AzureFileShareProtectionPolicyResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicyResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o AzureFileShareProtectionPolicyResponseOutput) WorkLoadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileShareProtectionPolicyResponse) *string { return v.WorkLoadType }).(pulumi.StringPtrOutput)
}

type AzureFileshareProtectedItem struct {
	BackupManagementType             *string                                  `pulumi:"backupManagementType"`
	BackupSetName                    *string                                  `pulumi:"backupSetName"`
	ContainerName                    *string                                  `pulumi:"containerName"`
	CreateMode                       *string                                  `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                  `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                  `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureFileshareProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                  `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                    `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                    `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                    `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails      `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                  `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                  `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                  `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                  `pulumi:"policyId"`
	ProtectedItemType                string                                   `pulumi:"protectedItemType"`
	ProtectionState                  *string                                  `pulumi:"protectionState"`
	ProtectionStatus                 *string                                  `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                 `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                                  `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                  `pulumi:"workloadType"`
}





type AzureFileshareProtectedItemInput interface {
	pulumi.Input

	ToAzureFileshareProtectedItemOutput() AzureFileshareProtectedItemOutput
	ToAzureFileshareProtectedItemOutputWithContext(context.Context) AzureFileshareProtectedItemOutput
}

type AzureFileshareProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                           `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                           `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                           `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureFileshareProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsMapInput                `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                           `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                           `pulumi:"policyId"`
	ProtectedItemType                pulumi.StringInput                              `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                           `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                           `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                         `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                           `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                           `pulumi:"workloadType"`
}

func (AzureFileshareProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileshareProtectedItem)(nil)).Elem()
}

func (i AzureFileshareProtectedItemArgs) ToAzureFileshareProtectedItemOutput() AzureFileshareProtectedItemOutput {
	return i.ToAzureFileshareProtectedItemOutputWithContext(context.Background())
}

func (i AzureFileshareProtectedItemArgs) ToAzureFileshareProtectedItemOutputWithContext(ctx context.Context) AzureFileshareProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileshareProtectedItemOutput)
}

type AzureFileshareProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureFileshareProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileshareProtectedItem)(nil)).Elem()
}

func (o AzureFileshareProtectedItemOutput) ToAzureFileshareProtectedItemOutput() AzureFileshareProtectedItemOutput {
	return o
}

func (o AzureFileshareProtectedItemOutput) ToAzureFileshareProtectedItemOutputWithContext(ctx context.Context) AzureFileshareProtectedItemOutput {
	return o
}

func (o AzureFileshareProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) ExtendedInfo() AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *AzureFileshareProtectedItemExtendedInfo { return v.ExtendedInfo }).(AzureFileshareProtectedItemExtendedInfoPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) KpisHealths() KPIResourceHealthDetailsMapOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) map[string]KPIResourceHealthDetails { return v.KpisHealths }).(KPIResourceHealthDetailsMapOutput)
}

func (o AzureFileshareProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureFileshareProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureFileshareProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureFileshareProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type AzureFileshareProtectedItemExtendedInfoInput interface {
	pulumi.Input

	ToAzureFileshareProtectedItemExtendedInfoOutput() AzureFileshareProtectedItemExtendedInfoOutput
	ToAzureFileshareProtectedItemExtendedInfoOutputWithContext(context.Context) AzureFileshareProtectedItemExtendedInfoOutput
}

type AzureFileshareProtectedItemExtendedInfoArgs struct {
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	PolicyState         pulumi.StringPtrInput `pulumi:"policyState"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (AzureFileshareProtectedItemExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileshareProtectedItemExtendedInfo)(nil)).Elem()
}

func (i AzureFileshareProtectedItemExtendedInfoArgs) ToAzureFileshareProtectedItemExtendedInfoOutput() AzureFileshareProtectedItemExtendedInfoOutput {
	return i.ToAzureFileshareProtectedItemExtendedInfoOutputWithContext(context.Background())
}

func (i AzureFileshareProtectedItemExtendedInfoArgs) ToAzureFileshareProtectedItemExtendedInfoOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileshareProtectedItemExtendedInfoOutput)
}

func (i AzureFileshareProtectedItemExtendedInfoArgs) ToAzureFileshareProtectedItemExtendedInfoPtrOutput() AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return i.ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i AzureFileshareProtectedItemExtendedInfoArgs) ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileshareProtectedItemExtendedInfoOutput).ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(ctx)
}









type AzureFileshareProtectedItemExtendedInfoPtrInput interface {
	pulumi.Input

	ToAzureFileshareProtectedItemExtendedInfoPtrOutput() AzureFileshareProtectedItemExtendedInfoPtrOutput
	ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(context.Context) AzureFileshareProtectedItemExtendedInfoPtrOutput
}

type azureFileshareProtectedItemExtendedInfoPtrType AzureFileshareProtectedItemExtendedInfoArgs

func AzureFileshareProtectedItemExtendedInfoPtr(v *AzureFileshareProtectedItemExtendedInfoArgs) AzureFileshareProtectedItemExtendedInfoPtrInput {
	return (*azureFileshareProtectedItemExtendedInfoPtrType)(v)
}

func (*azureFileshareProtectedItemExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileshareProtectedItemExtendedInfo)(nil)).Elem()
}

func (i *azureFileshareProtectedItemExtendedInfoPtrType) ToAzureFileshareProtectedItemExtendedInfoPtrOutput() AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return i.ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *azureFileshareProtectedItemExtendedInfoPtrType) ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileshareProtectedItemExtendedInfoPtrOutput)
}

type AzureFileshareProtectedItemExtendedInfoOutput struct{ *pulumi.OutputState }

func (AzureFileshareProtectedItemExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileshareProtectedItemExtendedInfo)(nil)).Elem()
}

func (o AzureFileshareProtectedItemExtendedInfoOutput) ToAzureFileshareProtectedItemExtendedInfoOutput() AzureFileshareProtectedItemExtendedInfoOutput {
	return o
}

func (o AzureFileshareProtectedItemExtendedInfoOutput) ToAzureFileshareProtectedItemExtendedInfoOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoOutput {
	return o
}

func (o AzureFileshareProtectedItemExtendedInfoOutput) ToAzureFileshareProtectedItemExtendedInfoPtrOutput() AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return o.ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (o AzureFileshareProtectedItemExtendedInfoOutput) ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFileshareProtectedItemExtendedInfo) *AzureFileshareProtectedItemExtendedInfo {
		return &v
	}).(AzureFileshareProtectedItemExtendedInfoPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemExtendedInfo) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemExtendedInfo) *string { return v.PolicyState }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemExtendedInfo) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type AzureFileshareProtectedItemExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (AzureFileshareProtectedItemExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileshareProtectedItemExtendedInfo)(nil)).Elem()
}

func (o AzureFileshareProtectedItemExtendedInfoPtrOutput) ToAzureFileshareProtectedItemExtendedInfoPtrOutput() AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o AzureFileshareProtectedItemExtendedInfoPtrOutput) ToAzureFileshareProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o AzureFileshareProtectedItemExtendedInfoPtrOutput) Elem() AzureFileshareProtectedItemExtendedInfoOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfo) AzureFileshareProtectedItemExtendedInfo {
		if v != nil {
			return *v
		}
		var ret AzureFileshareProtectedItemExtendedInfo
		return ret
	}).(AzureFileshareProtectedItemExtendedInfoOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoPtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoPtrOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.PolicyState
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoPtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfo) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type AzureFileshareProtectedItemExtendedInfoResponse struct {
	OldestRecoveryPoint   *string `pulumi:"oldestRecoveryPoint"`
	PolicyState           *string `pulumi:"policyState"`
	RecoveryPointCount    *int    `pulumi:"recoveryPointCount"`
	ResourceState         string  `pulumi:"resourceState"`
	ResourceStateSyncTime string  `pulumi:"resourceStateSyncTime"`
}





type AzureFileshareProtectedItemExtendedInfoResponseInput interface {
	pulumi.Input

	ToAzureFileshareProtectedItemExtendedInfoResponseOutput() AzureFileshareProtectedItemExtendedInfoResponseOutput
	ToAzureFileshareProtectedItemExtendedInfoResponseOutputWithContext(context.Context) AzureFileshareProtectedItemExtendedInfoResponseOutput
}

type AzureFileshareProtectedItemExtendedInfoResponseArgs struct {
	OldestRecoveryPoint   pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	PolicyState           pulumi.StringPtrInput `pulumi:"policyState"`
	RecoveryPointCount    pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
	ResourceState         pulumi.StringInput    `pulumi:"resourceState"`
	ResourceStateSyncTime pulumi.StringInput    `pulumi:"resourceStateSyncTime"`
}

func (AzureFileshareProtectedItemExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileshareProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i AzureFileshareProtectedItemExtendedInfoResponseArgs) ToAzureFileshareProtectedItemExtendedInfoResponseOutput() AzureFileshareProtectedItemExtendedInfoResponseOutput {
	return i.ToAzureFileshareProtectedItemExtendedInfoResponseOutputWithContext(context.Background())
}

func (i AzureFileshareProtectedItemExtendedInfoResponseArgs) ToAzureFileshareProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileshareProtectedItemExtendedInfoResponseOutput)
}

func (i AzureFileshareProtectedItemExtendedInfoResponseArgs) ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutput() AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i AzureFileshareProtectedItemExtendedInfoResponseArgs) ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileshareProtectedItemExtendedInfoResponseOutput).ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx)
}









type AzureFileshareProtectedItemExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutput() AzureFileshareProtectedItemExtendedInfoResponsePtrOutput
	ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Context) AzureFileshareProtectedItemExtendedInfoResponsePtrOutput
}

type azureFileshareProtectedItemExtendedInfoResponsePtrType AzureFileshareProtectedItemExtendedInfoResponseArgs

func AzureFileshareProtectedItemExtendedInfoResponsePtr(v *AzureFileshareProtectedItemExtendedInfoResponseArgs) AzureFileshareProtectedItemExtendedInfoResponsePtrInput {
	return (*azureFileshareProtectedItemExtendedInfoResponsePtrType)(v)
}

func (*azureFileshareProtectedItemExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileshareProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i *azureFileshareProtectedItemExtendedInfoResponsePtrType) ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutput() AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *azureFileshareProtectedItemExtendedInfoResponsePtrType) ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileshareProtectedItemExtendedInfoResponsePtrOutput)
}

type AzureFileshareProtectedItemExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureFileshareProtectedItemExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileshareProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) ToAzureFileshareProtectedItemExtendedInfoResponseOutput() AzureFileshareProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) ToAzureFileshareProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutput() AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return o.ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFileshareProtectedItemExtendedInfoResponse) *AzureFileshareProtectedItemExtendedInfoResponse {
		return &v
	}).(AzureFileshareProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemExtendedInfoResponse) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemExtendedInfoResponse) *string { return v.PolicyState }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemExtendedInfoResponse) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemExtendedInfoResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponseOutput) ResourceStateSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemExtendedInfoResponse) string { return v.ResourceStateSyncTime }).(pulumi.StringOutput)
}

type AzureFileshareProtectedItemExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileshareProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutput() AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) ToAzureFileshareProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) Elem() AzureFileshareProtectedItemExtendedInfoResponseOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfoResponse) AzureFileshareProtectedItemExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureFileshareProtectedItemExtendedInfoResponse
		return ret
	}).(AzureFileshareProtectedItemExtendedInfoResponseOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyState
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) ResourceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceState
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemExtendedInfoResponsePtrOutput) ResourceStateSyncTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileshareProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceStateSyncTime
	}).(pulumi.StringPtrOutput)
}

type AzureFileshareProtectedItemResponse struct {
	BackupManagementType             *string                                          `pulumi:"backupManagementType"`
	BackupSetName                    *string                                          `pulumi:"backupSetName"`
	ContainerName                    *string                                          `pulumi:"containerName"`
	CreateMode                       *string                                          `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                          `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                          `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureFileshareProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                          `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                            `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                            `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                            `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse      `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                          `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                          `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                          `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                          `pulumi:"policyId"`
	ProtectedItemType                string                                           `pulumi:"protectedItemType"`
	ProtectionState                  *string                                          `pulumi:"protectionState"`
	ProtectionStatus                 *string                                          `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                         `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                                          `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                          `pulumi:"workloadType"`
}





type AzureFileshareProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureFileshareProtectedItemResponseOutput() AzureFileshareProtectedItemResponseOutput
	ToAzureFileshareProtectedItemResponseOutputWithContext(context.Context) AzureFileshareProtectedItemResponseOutput
}

type AzureFileshareProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                   `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                   `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                                   `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureFileshareProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsResponseMapInput                `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                   `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                                   `pulumi:"policyId"`
	ProtectedItemType                pulumi.StringInput                                      `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                   `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                                   `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                                 `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                                   `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                                   `pulumi:"workloadType"`
}

func (AzureFileshareProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileshareProtectedItemResponse)(nil)).Elem()
}

func (i AzureFileshareProtectedItemResponseArgs) ToAzureFileshareProtectedItemResponseOutput() AzureFileshareProtectedItemResponseOutput {
	return i.ToAzureFileshareProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureFileshareProtectedItemResponseArgs) ToAzureFileshareProtectedItemResponseOutputWithContext(ctx context.Context) AzureFileshareProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileshareProtectedItemResponseOutput)
}

type AzureFileshareProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureFileshareProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileshareProtectedItemResponse)(nil)).Elem()
}

func (o AzureFileshareProtectedItemResponseOutput) ToAzureFileshareProtectedItemResponseOutput() AzureFileshareProtectedItemResponseOutput {
	return o
}

func (o AzureFileshareProtectedItemResponseOutput) ToAzureFileshareProtectedItemResponseOutputWithContext(ctx context.Context) AzureFileshareProtectedItemResponseOutput {
	return o
}

func (o AzureFileshareProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) ExtendedInfo() AzureFileshareProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *AzureFileshareProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureFileshareProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) KpisHealths() KPIResourceHealthDetailsResponseMapOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) map[string]KPIResourceHealthDetailsResponse {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsResponseMapOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureFileshareProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileshareProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureIaaSClassicComputeVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}





type AzureIaaSClassicComputeVMContainerInput interface {
	pulumi.Input

	ToAzureIaaSClassicComputeVMContainerOutput() AzureIaaSClassicComputeVMContainerOutput
	ToAzureIaaSClassicComputeVMContainerOutputWithContext(context.Context) AzureIaaSClassicComputeVMContainerOutput
}

type AzureIaaSClassicComputeVMContainerArgs struct {
	BackupManagementType  pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ContainerType         pulumi.StringInput    `pulumi:"containerType"`
	FriendlyName          pulumi.StringPtrInput `pulumi:"friendlyName"`
	HealthStatus          pulumi.StringPtrInput `pulumi:"healthStatus"`
	RegistrationStatus    pulumi.StringPtrInput `pulumi:"registrationStatus"`
	ResourceGroup         pulumi.StringPtrInput `pulumi:"resourceGroup"`
	VirtualMachineId      pulumi.StringPtrInput `pulumi:"virtualMachineId"`
	VirtualMachineVersion pulumi.StringPtrInput `pulumi:"virtualMachineVersion"`
}

func (AzureIaaSClassicComputeVMContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSClassicComputeVMContainer)(nil)).Elem()
}

func (i AzureIaaSClassicComputeVMContainerArgs) ToAzureIaaSClassicComputeVMContainerOutput() AzureIaaSClassicComputeVMContainerOutput {
	return i.ToAzureIaaSClassicComputeVMContainerOutputWithContext(context.Background())
}

func (i AzureIaaSClassicComputeVMContainerArgs) ToAzureIaaSClassicComputeVMContainerOutputWithContext(ctx context.Context) AzureIaaSClassicComputeVMContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSClassicComputeVMContainerOutput)
}

type AzureIaaSClassicComputeVMContainerOutput struct{ *pulumi.OutputState }

func (AzureIaaSClassicComputeVMContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSClassicComputeVMContainer)(nil)).Elem()
}

func (o AzureIaaSClassicComputeVMContainerOutput) ToAzureIaaSClassicComputeVMContainerOutput() AzureIaaSClassicComputeVMContainerOutput {
	return o
}

func (o AzureIaaSClassicComputeVMContainerOutput) ToAzureIaaSClassicComputeVMContainerOutputWithContext(ctx context.Context) AzureIaaSClassicComputeVMContainerOutput {
	return o
}

func (o AzureIaaSClassicComputeVMContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureIaaSClassicComputeVMContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainer) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainer) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerOutput) VirtualMachineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainer) *string { return v.VirtualMachineVersion }).(pulumi.StringPtrOutput)
}

type AzureIaaSClassicComputeVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}





type AzureIaaSClassicComputeVMContainerResponseInput interface {
	pulumi.Input

	ToAzureIaaSClassicComputeVMContainerResponseOutput() AzureIaaSClassicComputeVMContainerResponseOutput
	ToAzureIaaSClassicComputeVMContainerResponseOutputWithContext(context.Context) AzureIaaSClassicComputeVMContainerResponseOutput
}

type AzureIaaSClassicComputeVMContainerResponseArgs struct {
	BackupManagementType  pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ContainerType         pulumi.StringInput    `pulumi:"containerType"`
	FriendlyName          pulumi.StringPtrInput `pulumi:"friendlyName"`
	HealthStatus          pulumi.StringPtrInput `pulumi:"healthStatus"`
	RegistrationStatus    pulumi.StringPtrInput `pulumi:"registrationStatus"`
	ResourceGroup         pulumi.StringPtrInput `pulumi:"resourceGroup"`
	VirtualMachineId      pulumi.StringPtrInput `pulumi:"virtualMachineId"`
	VirtualMachineVersion pulumi.StringPtrInput `pulumi:"virtualMachineVersion"`
}

func (AzureIaaSClassicComputeVMContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSClassicComputeVMContainerResponse)(nil)).Elem()
}

func (i AzureIaaSClassicComputeVMContainerResponseArgs) ToAzureIaaSClassicComputeVMContainerResponseOutput() AzureIaaSClassicComputeVMContainerResponseOutput {
	return i.ToAzureIaaSClassicComputeVMContainerResponseOutputWithContext(context.Background())
}

func (i AzureIaaSClassicComputeVMContainerResponseArgs) ToAzureIaaSClassicComputeVMContainerResponseOutputWithContext(ctx context.Context) AzureIaaSClassicComputeVMContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSClassicComputeVMContainerResponseOutput)
}

type AzureIaaSClassicComputeVMContainerResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSClassicComputeVMContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSClassicComputeVMContainerResponse)(nil)).Elem()
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) ToAzureIaaSClassicComputeVMContainerResponseOutput() AzureIaaSClassicComputeVMContainerResponseOutput {
	return o
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) ToAzureIaaSClassicComputeVMContainerResponseOutputWithContext(ctx context.Context) AzureIaaSClassicComputeVMContainerResponseOutput {
	return o
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainerResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainerResponse) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMContainerResponseOutput) VirtualMachineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMContainerResponse) *string { return v.VirtualMachineVersion }).(pulumi.StringPtrOutput)
}

type AzureIaaSClassicComputeVMProtectedItem struct {
	BackupManagementType             *string                               `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	HealthStatus                     *string                               `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                               `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	ProtectedItemDataId              *string                               `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                              `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                               `pulumi:"virtualMachineId"`
	WorkloadType                     *string                               `pulumi:"workloadType"`
}





type AzureIaaSClassicComputeVMProtectedItemInput interface {
	pulumi.Input

	ToAzureIaaSClassicComputeVMProtectedItemOutput() AzureIaaSClassicComputeVMProtectedItemOutput
	ToAzureIaaSClassicComputeVMProtectedItemOutputWithContext(context.Context) AzureIaaSClassicComputeVMProtectedItemOutput
}

type AzureIaaSClassicComputeVMProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                        `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                        `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                        `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                        `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                        `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                        `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureIaaSVMProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	ExtendedProperties               ExtendedPropertiesPtrInput                   `pulumi:"extendedProperties"`
	FriendlyName                     pulumi.StringPtrInput                        `pulumi:"friendlyName"`
	HealthStatus                     pulumi.StringPtrInput                        `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                          `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                          `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                          `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsMapInput             `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                        `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                        `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                        `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                        `pulumi:"policyId"`
	ProtectedItemDataId              pulumi.StringPtrInput                        `pulumi:"protectedItemDataId"`
	ProtectedItemType                pulumi.StringInput                           `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                        `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                        `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                      `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                        `pulumi:"sourceResourceId"`
	VirtualMachineId                 pulumi.StringPtrInput                        `pulumi:"virtualMachineId"`
	WorkloadType                     pulumi.StringPtrInput                        `pulumi:"workloadType"`
}

func (AzureIaaSClassicComputeVMProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSClassicComputeVMProtectedItem)(nil)).Elem()
}

func (i AzureIaaSClassicComputeVMProtectedItemArgs) ToAzureIaaSClassicComputeVMProtectedItemOutput() AzureIaaSClassicComputeVMProtectedItemOutput {
	return i.ToAzureIaaSClassicComputeVMProtectedItemOutputWithContext(context.Background())
}

func (i AzureIaaSClassicComputeVMProtectedItemArgs) ToAzureIaaSClassicComputeVMProtectedItemOutputWithContext(ctx context.Context) AzureIaaSClassicComputeVMProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSClassicComputeVMProtectedItemOutput)
}

type AzureIaaSClassicComputeVMProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureIaaSClassicComputeVMProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSClassicComputeVMProtectedItem)(nil)).Elem()
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ToAzureIaaSClassicComputeVMProtectedItemOutput() AzureIaaSClassicComputeVMProtectedItemOutput {
	return o
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ToAzureIaaSClassicComputeVMProtectedItemOutputWithContext(ctx context.Context) AzureIaaSClassicComputeVMProtectedItemOutput {
	return o
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ExtendedInfo() AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *AzureIaaSVMProtectedItemExtendedInfo {
		return v.ExtendedInfo
	}).(AzureIaaSVMProtectedItemExtendedInfoPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ExtendedProperties() ExtendedPropertiesPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *ExtendedProperties { return v.ExtendedProperties }).(ExtendedPropertiesPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) KpisHealths() KPIResourceHealthDetailsMapOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) map[string]KPIResourceHealthDetails {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsMapOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ProtectedItemDataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.ProtectedItemDataId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureIaaSClassicComputeVMProtectedItemResponse struct {
	BackupManagementType             *string                                       `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                                       `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     *string                                       `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                       `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	ProtectedItemDataId              *string                                       `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                      `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                                       `pulumi:"virtualMachineId"`
	WorkloadType                     *string                                       `pulumi:"workloadType"`
}





type AzureIaaSClassicComputeVMProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureIaaSClassicComputeVMProtectedItemResponseOutput() AzureIaaSClassicComputeVMProtectedItemResponseOutput
	ToAzureIaaSClassicComputeVMProtectedItemResponseOutputWithContext(context.Context) AzureIaaSClassicComputeVMProtectedItemResponseOutput
}

type AzureIaaSClassicComputeVMProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                                `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureIaaSVMProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	ExtendedProperties               ExtendedPropertiesResponsePtrInput                   `pulumi:"extendedProperties"`
	FriendlyName                     pulumi.StringPtrInput                                `pulumi:"friendlyName"`
	HealthDetails                    AzureIaaSVMHealthDetailsResponseArrayInput           `pulumi:"healthDetails"`
	HealthStatus                     pulumi.StringPtrInput                                `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                  `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                  `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                  `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsResponseMapInput             `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                                `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                                `pulumi:"policyId"`
	ProtectedItemDataId              pulumi.StringPtrInput                                `pulumi:"protectedItemDataId"`
	ProtectedItemType                pulumi.StringInput                                   `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                                `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                              `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                                `pulumi:"sourceResourceId"`
	VirtualMachineId                 pulumi.StringPtrInput                                `pulumi:"virtualMachineId"`
	WorkloadType                     pulumi.StringPtrInput                                `pulumi:"workloadType"`
}

func (AzureIaaSClassicComputeVMProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSClassicComputeVMProtectedItemResponse)(nil)).Elem()
}

func (i AzureIaaSClassicComputeVMProtectedItemResponseArgs) ToAzureIaaSClassicComputeVMProtectedItemResponseOutput() AzureIaaSClassicComputeVMProtectedItemResponseOutput {
	return i.ToAzureIaaSClassicComputeVMProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureIaaSClassicComputeVMProtectedItemResponseArgs) ToAzureIaaSClassicComputeVMProtectedItemResponseOutputWithContext(ctx context.Context) AzureIaaSClassicComputeVMProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSClassicComputeVMProtectedItemResponseOutput)
}

type AzureIaaSClassicComputeVMProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSClassicComputeVMProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSClassicComputeVMProtectedItemResponse)(nil)).Elem()
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ToAzureIaaSClassicComputeVMProtectedItemResponseOutput() AzureIaaSClassicComputeVMProtectedItemResponseOutput {
	return o
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ToAzureIaaSClassicComputeVMProtectedItemResponseOutputWithContext(ctx context.Context) AzureIaaSClassicComputeVMProtectedItemResponseOutput {
	return o
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ExtendedInfo() AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *AzureIaaSVMProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ExtendedProperties() ExtendedPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *ExtendedPropertiesResponse {
		return v.ExtendedProperties
	}).(ExtendedPropertiesResponsePtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) HealthDetails() AzureIaaSVMHealthDetailsResponseArrayOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) []AzureIaaSVMHealthDetailsResponse {
		return v.HealthDetails
	}).(AzureIaaSVMHealthDetailsResponseArrayOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *bool {
		return v.IsDeferredDeleteScheduleUpcoming
	}).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) KpisHealths() KPIResourceHealthDetailsResponseMapOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) map[string]KPIResourceHealthDetailsResponse {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsResponseMapOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ProtectedItemDataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.ProtectedItemDataId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) []string {
		return v.ResourceGuardOperationRequests
	}).(pulumi.StringArrayOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSClassicComputeVMProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSClassicComputeVMProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureIaaSComputeVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}





type AzureIaaSComputeVMContainerInput interface {
	pulumi.Input

	ToAzureIaaSComputeVMContainerOutput() AzureIaaSComputeVMContainerOutput
	ToAzureIaaSComputeVMContainerOutputWithContext(context.Context) AzureIaaSComputeVMContainerOutput
}

type AzureIaaSComputeVMContainerArgs struct {
	BackupManagementType  pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ContainerType         pulumi.StringInput    `pulumi:"containerType"`
	FriendlyName          pulumi.StringPtrInput `pulumi:"friendlyName"`
	HealthStatus          pulumi.StringPtrInput `pulumi:"healthStatus"`
	RegistrationStatus    pulumi.StringPtrInput `pulumi:"registrationStatus"`
	ResourceGroup         pulumi.StringPtrInput `pulumi:"resourceGroup"`
	VirtualMachineId      pulumi.StringPtrInput `pulumi:"virtualMachineId"`
	VirtualMachineVersion pulumi.StringPtrInput `pulumi:"virtualMachineVersion"`
}

func (AzureIaaSComputeVMContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSComputeVMContainer)(nil)).Elem()
}

func (i AzureIaaSComputeVMContainerArgs) ToAzureIaaSComputeVMContainerOutput() AzureIaaSComputeVMContainerOutput {
	return i.ToAzureIaaSComputeVMContainerOutputWithContext(context.Background())
}

func (i AzureIaaSComputeVMContainerArgs) ToAzureIaaSComputeVMContainerOutputWithContext(ctx context.Context) AzureIaaSComputeVMContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSComputeVMContainerOutput)
}

type AzureIaaSComputeVMContainerOutput struct{ *pulumi.OutputState }

func (AzureIaaSComputeVMContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSComputeVMContainer)(nil)).Elem()
}

func (o AzureIaaSComputeVMContainerOutput) ToAzureIaaSComputeVMContainerOutput() AzureIaaSComputeVMContainerOutput {
	return o
}

func (o AzureIaaSComputeVMContainerOutput) ToAzureIaaSComputeVMContainerOutputWithContext(ctx context.Context) AzureIaaSComputeVMContainerOutput {
	return o
}

func (o AzureIaaSComputeVMContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureIaaSComputeVMContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainer) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainer) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerOutput) VirtualMachineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainer) *string { return v.VirtualMachineVersion }).(pulumi.StringPtrOutput)
}

type AzureIaaSComputeVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}





type AzureIaaSComputeVMContainerResponseInput interface {
	pulumi.Input

	ToAzureIaaSComputeVMContainerResponseOutput() AzureIaaSComputeVMContainerResponseOutput
	ToAzureIaaSComputeVMContainerResponseOutputWithContext(context.Context) AzureIaaSComputeVMContainerResponseOutput
}

type AzureIaaSComputeVMContainerResponseArgs struct {
	BackupManagementType  pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ContainerType         pulumi.StringInput    `pulumi:"containerType"`
	FriendlyName          pulumi.StringPtrInput `pulumi:"friendlyName"`
	HealthStatus          pulumi.StringPtrInput `pulumi:"healthStatus"`
	RegistrationStatus    pulumi.StringPtrInput `pulumi:"registrationStatus"`
	ResourceGroup         pulumi.StringPtrInput `pulumi:"resourceGroup"`
	VirtualMachineId      pulumi.StringPtrInput `pulumi:"virtualMachineId"`
	VirtualMachineVersion pulumi.StringPtrInput `pulumi:"virtualMachineVersion"`
}

func (AzureIaaSComputeVMContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSComputeVMContainerResponse)(nil)).Elem()
}

func (i AzureIaaSComputeVMContainerResponseArgs) ToAzureIaaSComputeVMContainerResponseOutput() AzureIaaSComputeVMContainerResponseOutput {
	return i.ToAzureIaaSComputeVMContainerResponseOutputWithContext(context.Background())
}

func (i AzureIaaSComputeVMContainerResponseArgs) ToAzureIaaSComputeVMContainerResponseOutputWithContext(ctx context.Context) AzureIaaSComputeVMContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSComputeVMContainerResponseOutput)
}

type AzureIaaSComputeVMContainerResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSComputeVMContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSComputeVMContainerResponse)(nil)).Elem()
}

func (o AzureIaaSComputeVMContainerResponseOutput) ToAzureIaaSComputeVMContainerResponseOutput() AzureIaaSComputeVMContainerResponseOutput {
	return o
}

func (o AzureIaaSComputeVMContainerResponseOutput) ToAzureIaaSComputeVMContainerResponseOutputWithContext(ctx context.Context) AzureIaaSComputeVMContainerResponseOutput {
	return o
}

func (o AzureIaaSComputeVMContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureIaaSComputeVMContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainerResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerResponseOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainerResponse) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMContainerResponseOutput) VirtualMachineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMContainerResponse) *string { return v.VirtualMachineVersion }).(pulumi.StringPtrOutput)
}

type AzureIaaSComputeVMProtectedItem struct {
	BackupManagementType             *string                               `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	HealthStatus                     *string                               `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                               `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	ProtectedItemDataId              *string                               `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                              `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                               `pulumi:"virtualMachineId"`
	WorkloadType                     *string                               `pulumi:"workloadType"`
}





type AzureIaaSComputeVMProtectedItemInput interface {
	pulumi.Input

	ToAzureIaaSComputeVMProtectedItemOutput() AzureIaaSComputeVMProtectedItemOutput
	ToAzureIaaSComputeVMProtectedItemOutputWithContext(context.Context) AzureIaaSComputeVMProtectedItemOutput
}

type AzureIaaSComputeVMProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                        `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                        `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                        `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                        `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                        `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                        `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureIaaSVMProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	ExtendedProperties               ExtendedPropertiesPtrInput                   `pulumi:"extendedProperties"`
	FriendlyName                     pulumi.StringPtrInput                        `pulumi:"friendlyName"`
	HealthStatus                     pulumi.StringPtrInput                        `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                          `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                          `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                          `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsMapInput             `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                        `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                        `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                        `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                        `pulumi:"policyId"`
	ProtectedItemDataId              pulumi.StringPtrInput                        `pulumi:"protectedItemDataId"`
	ProtectedItemType                pulumi.StringInput                           `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                        `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                        `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                      `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                        `pulumi:"sourceResourceId"`
	VirtualMachineId                 pulumi.StringPtrInput                        `pulumi:"virtualMachineId"`
	WorkloadType                     pulumi.StringPtrInput                        `pulumi:"workloadType"`
}

func (AzureIaaSComputeVMProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSComputeVMProtectedItem)(nil)).Elem()
}

func (i AzureIaaSComputeVMProtectedItemArgs) ToAzureIaaSComputeVMProtectedItemOutput() AzureIaaSComputeVMProtectedItemOutput {
	return i.ToAzureIaaSComputeVMProtectedItemOutputWithContext(context.Background())
}

func (i AzureIaaSComputeVMProtectedItemArgs) ToAzureIaaSComputeVMProtectedItemOutputWithContext(ctx context.Context) AzureIaaSComputeVMProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSComputeVMProtectedItemOutput)
}

type AzureIaaSComputeVMProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureIaaSComputeVMProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSComputeVMProtectedItem)(nil)).Elem()
}

func (o AzureIaaSComputeVMProtectedItemOutput) ToAzureIaaSComputeVMProtectedItemOutput() AzureIaaSComputeVMProtectedItemOutput {
	return o
}

func (o AzureIaaSComputeVMProtectedItemOutput) ToAzureIaaSComputeVMProtectedItemOutputWithContext(ctx context.Context) AzureIaaSComputeVMProtectedItemOutput {
	return o
}

func (o AzureIaaSComputeVMProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) ExtendedInfo() AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *AzureIaaSVMProtectedItemExtendedInfo { return v.ExtendedInfo }).(AzureIaaSVMProtectedItemExtendedInfoPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) ExtendedProperties() ExtendedPropertiesPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *ExtendedProperties { return v.ExtendedProperties }).(ExtendedPropertiesPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) KpisHealths() KPIResourceHealthDetailsMapOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) map[string]KPIResourceHealthDetails { return v.KpisHealths }).(KPIResourceHealthDetailsMapOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) ProtectedItemDataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.ProtectedItemDataId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureIaaSComputeVMProtectedItemResponse struct {
	BackupManagementType             *string                                       `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                                       `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     *string                                       `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                       `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	ProtectedItemDataId              *string                                       `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                      `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                                       `pulumi:"virtualMachineId"`
	WorkloadType                     *string                                       `pulumi:"workloadType"`
}





type AzureIaaSComputeVMProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureIaaSComputeVMProtectedItemResponseOutput() AzureIaaSComputeVMProtectedItemResponseOutput
	ToAzureIaaSComputeVMProtectedItemResponseOutputWithContext(context.Context) AzureIaaSComputeVMProtectedItemResponseOutput
}

type AzureIaaSComputeVMProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                                `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureIaaSVMProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	ExtendedProperties               ExtendedPropertiesResponsePtrInput                   `pulumi:"extendedProperties"`
	FriendlyName                     pulumi.StringPtrInput                                `pulumi:"friendlyName"`
	HealthDetails                    AzureIaaSVMHealthDetailsResponseArrayInput           `pulumi:"healthDetails"`
	HealthStatus                     pulumi.StringPtrInput                                `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                  `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                  `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                  `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsResponseMapInput             `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                                `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                                `pulumi:"policyId"`
	ProtectedItemDataId              pulumi.StringPtrInput                                `pulumi:"protectedItemDataId"`
	ProtectedItemType                pulumi.StringInput                                   `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                                `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                              `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                                `pulumi:"sourceResourceId"`
	VirtualMachineId                 pulumi.StringPtrInput                                `pulumi:"virtualMachineId"`
	WorkloadType                     pulumi.StringPtrInput                                `pulumi:"workloadType"`
}

func (AzureIaaSComputeVMProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSComputeVMProtectedItemResponse)(nil)).Elem()
}

func (i AzureIaaSComputeVMProtectedItemResponseArgs) ToAzureIaaSComputeVMProtectedItemResponseOutput() AzureIaaSComputeVMProtectedItemResponseOutput {
	return i.ToAzureIaaSComputeVMProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureIaaSComputeVMProtectedItemResponseArgs) ToAzureIaaSComputeVMProtectedItemResponseOutputWithContext(ctx context.Context) AzureIaaSComputeVMProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSComputeVMProtectedItemResponseOutput)
}

type AzureIaaSComputeVMProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSComputeVMProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSComputeVMProtectedItemResponse)(nil)).Elem()
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ToAzureIaaSComputeVMProtectedItemResponseOutput() AzureIaaSComputeVMProtectedItemResponseOutput {
	return o
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ToAzureIaaSComputeVMProtectedItemResponseOutputWithContext(ctx context.Context) AzureIaaSComputeVMProtectedItemResponseOutput {
	return o
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ExtendedInfo() AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *AzureIaaSVMProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ExtendedProperties() ExtendedPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *ExtendedPropertiesResponse {
		return v.ExtendedProperties
	}).(ExtendedPropertiesResponsePtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) HealthDetails() AzureIaaSVMHealthDetailsResponseArrayOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) []AzureIaaSVMHealthDetailsResponse {
		return v.HealthDetails
	}).(AzureIaaSVMHealthDetailsResponseArrayOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) KpisHealths() KPIResourceHealthDetailsResponseMapOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) map[string]KPIResourceHealthDetailsResponse {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsResponseMapOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ProtectedItemDataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.ProtectedItemDataId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSComputeVMProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSComputeVMProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureIaaSVMHealthDetailsResponse struct {
	Code            int      `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           string   `pulumi:"title"`
}





type AzureIaaSVMHealthDetailsResponseInput interface {
	pulumi.Input

	ToAzureIaaSVMHealthDetailsResponseOutput() AzureIaaSVMHealthDetailsResponseOutput
	ToAzureIaaSVMHealthDetailsResponseOutputWithContext(context.Context) AzureIaaSVMHealthDetailsResponseOutput
}

type AzureIaaSVMHealthDetailsResponseArgs struct {
	Code            pulumi.IntInput         `pulumi:"code"`
	Message         pulumi.StringInput      `pulumi:"message"`
	Recommendations pulumi.StringArrayInput `pulumi:"recommendations"`
	Title           pulumi.StringInput      `pulumi:"title"`
}

func (AzureIaaSVMHealthDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMHealthDetailsResponse)(nil)).Elem()
}

func (i AzureIaaSVMHealthDetailsResponseArgs) ToAzureIaaSVMHealthDetailsResponseOutput() AzureIaaSVMHealthDetailsResponseOutput {
	return i.ToAzureIaaSVMHealthDetailsResponseOutputWithContext(context.Background())
}

func (i AzureIaaSVMHealthDetailsResponseArgs) ToAzureIaaSVMHealthDetailsResponseOutputWithContext(ctx context.Context) AzureIaaSVMHealthDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMHealthDetailsResponseOutput)
}





type AzureIaaSVMHealthDetailsResponseArrayInput interface {
	pulumi.Input

	ToAzureIaaSVMHealthDetailsResponseArrayOutput() AzureIaaSVMHealthDetailsResponseArrayOutput
	ToAzureIaaSVMHealthDetailsResponseArrayOutputWithContext(context.Context) AzureIaaSVMHealthDetailsResponseArrayOutput
}

type AzureIaaSVMHealthDetailsResponseArray []AzureIaaSVMHealthDetailsResponseInput

func (AzureIaaSVMHealthDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureIaaSVMHealthDetailsResponse)(nil)).Elem()
}

func (i AzureIaaSVMHealthDetailsResponseArray) ToAzureIaaSVMHealthDetailsResponseArrayOutput() AzureIaaSVMHealthDetailsResponseArrayOutput {
	return i.ToAzureIaaSVMHealthDetailsResponseArrayOutputWithContext(context.Background())
}

func (i AzureIaaSVMHealthDetailsResponseArray) ToAzureIaaSVMHealthDetailsResponseArrayOutputWithContext(ctx context.Context) AzureIaaSVMHealthDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMHealthDetailsResponseArrayOutput)
}

type AzureIaaSVMHealthDetailsResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMHealthDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMHealthDetailsResponse)(nil)).Elem()
}

func (o AzureIaaSVMHealthDetailsResponseOutput) ToAzureIaaSVMHealthDetailsResponseOutput() AzureIaaSVMHealthDetailsResponseOutput {
	return o
}

func (o AzureIaaSVMHealthDetailsResponseOutput) ToAzureIaaSVMHealthDetailsResponseOutputWithContext(ctx context.Context) AzureIaaSVMHealthDetailsResponseOutput {
	return o
}

func (o AzureIaaSVMHealthDetailsResponseOutput) Code() pulumi.IntOutput {
	return o.ApplyT(func(v AzureIaaSVMHealthDetailsResponse) int { return v.Code }).(pulumi.IntOutput)
}

func (o AzureIaaSVMHealthDetailsResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSVMHealthDetailsResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o AzureIaaSVMHealthDetailsResponseOutput) Recommendations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSVMHealthDetailsResponse) []string { return v.Recommendations }).(pulumi.StringArrayOutput)
}

func (o AzureIaaSVMHealthDetailsResponseOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSVMHealthDetailsResponse) string { return v.Title }).(pulumi.StringOutput)
}

type AzureIaaSVMHealthDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMHealthDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureIaaSVMHealthDetailsResponse)(nil)).Elem()
}

func (o AzureIaaSVMHealthDetailsResponseArrayOutput) ToAzureIaaSVMHealthDetailsResponseArrayOutput() AzureIaaSVMHealthDetailsResponseArrayOutput {
	return o
}

func (o AzureIaaSVMHealthDetailsResponseArrayOutput) ToAzureIaaSVMHealthDetailsResponseArrayOutputWithContext(ctx context.Context) AzureIaaSVMHealthDetailsResponseArrayOutput {
	return o
}

func (o AzureIaaSVMHealthDetailsResponseArrayOutput) Index(i pulumi.IntInput) AzureIaaSVMHealthDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureIaaSVMHealthDetailsResponse {
		return vs[0].([]AzureIaaSVMHealthDetailsResponse)[vs[1].(int)]
	}).(AzureIaaSVMHealthDetailsResponseOutput)
}

type AzureIaaSVMProtectedItem struct {
	BackupManagementType             *string                               `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	HealthStatus                     *string                               `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                               `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	ProtectedItemDataId              *string                               `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                              `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                               `pulumi:"virtualMachineId"`
	WorkloadType                     *string                               `pulumi:"workloadType"`
}





type AzureIaaSVMProtectedItemInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectedItemOutput() AzureIaaSVMProtectedItemOutput
	ToAzureIaaSVMProtectedItemOutputWithContext(context.Context) AzureIaaSVMProtectedItemOutput
}

type AzureIaaSVMProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                        `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                        `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                        `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                        `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                        `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                        `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureIaaSVMProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	ExtendedProperties               ExtendedPropertiesPtrInput                   `pulumi:"extendedProperties"`
	FriendlyName                     pulumi.StringPtrInput                        `pulumi:"friendlyName"`
	HealthStatus                     pulumi.StringPtrInput                        `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                          `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                          `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                          `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsMapInput             `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                        `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                        `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                        `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                        `pulumi:"policyId"`
	ProtectedItemDataId              pulumi.StringPtrInput                        `pulumi:"protectedItemDataId"`
	ProtectedItemType                pulumi.StringInput                           `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                        `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                        `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                      `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                        `pulumi:"sourceResourceId"`
	VirtualMachineId                 pulumi.StringPtrInput                        `pulumi:"virtualMachineId"`
	WorkloadType                     pulumi.StringPtrInput                        `pulumi:"workloadType"`
}

func (AzureIaaSVMProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectedItem)(nil)).Elem()
}

func (i AzureIaaSVMProtectedItemArgs) ToAzureIaaSVMProtectedItemOutput() AzureIaaSVMProtectedItemOutput {
	return i.ToAzureIaaSVMProtectedItemOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectedItemArgs) ToAzureIaaSVMProtectedItemOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectedItemOutput)
}

type AzureIaaSVMProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectedItem)(nil)).Elem()
}

func (o AzureIaaSVMProtectedItemOutput) ToAzureIaaSVMProtectedItemOutput() AzureIaaSVMProtectedItemOutput {
	return o
}

func (o AzureIaaSVMProtectedItemOutput) ToAzureIaaSVMProtectedItemOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemOutput {
	return o
}

func (o AzureIaaSVMProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) ExtendedInfo() AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *AzureIaaSVMProtectedItemExtendedInfo { return v.ExtendedInfo }).(AzureIaaSVMProtectedItemExtendedInfoPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) ExtendedProperties() ExtendedPropertiesPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *ExtendedProperties { return v.ExtendedProperties }).(ExtendedPropertiesPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) KpisHealths() KPIResourceHealthDetailsMapOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) map[string]KPIResourceHealthDetails { return v.KpisHealths }).(KPIResourceHealthDetailsMapOutput)
}

func (o AzureIaaSVMProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) ProtectedItemDataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.ProtectedItemDataId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureIaaSVMProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureIaaSVMProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureIaaSVMProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyInconsistent  *bool   `pulumi:"policyInconsistent"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type AzureIaaSVMProtectedItemExtendedInfoInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectedItemExtendedInfoOutput() AzureIaaSVMProtectedItemExtendedInfoOutput
	ToAzureIaaSVMProtectedItemExtendedInfoOutputWithContext(context.Context) AzureIaaSVMProtectedItemExtendedInfoOutput
}

type AzureIaaSVMProtectedItemExtendedInfoArgs struct {
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	PolicyInconsistent  pulumi.BoolPtrInput   `pulumi:"policyInconsistent"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (AzureIaaSVMProtectedItemExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectedItemExtendedInfo)(nil)).Elem()
}

func (i AzureIaaSVMProtectedItemExtendedInfoArgs) ToAzureIaaSVMProtectedItemExtendedInfoOutput() AzureIaaSVMProtectedItemExtendedInfoOutput {
	return i.ToAzureIaaSVMProtectedItemExtendedInfoOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectedItemExtendedInfoArgs) ToAzureIaaSVMProtectedItemExtendedInfoOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectedItemExtendedInfoOutput)
}

func (i AzureIaaSVMProtectedItemExtendedInfoArgs) ToAzureIaaSVMProtectedItemExtendedInfoPtrOutput() AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return i.ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectedItemExtendedInfoArgs) ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectedItemExtendedInfoOutput).ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(ctx)
}









type AzureIaaSVMProtectedItemExtendedInfoPtrInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectedItemExtendedInfoPtrOutput() AzureIaaSVMProtectedItemExtendedInfoPtrOutput
	ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(context.Context) AzureIaaSVMProtectedItemExtendedInfoPtrOutput
}

type azureIaaSVMProtectedItemExtendedInfoPtrType AzureIaaSVMProtectedItemExtendedInfoArgs

func AzureIaaSVMProtectedItemExtendedInfoPtr(v *AzureIaaSVMProtectedItemExtendedInfoArgs) AzureIaaSVMProtectedItemExtendedInfoPtrInput {
	return (*azureIaaSVMProtectedItemExtendedInfoPtrType)(v)
}

func (*azureIaaSVMProtectedItemExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureIaaSVMProtectedItemExtendedInfo)(nil)).Elem()
}

func (i *azureIaaSVMProtectedItemExtendedInfoPtrType) ToAzureIaaSVMProtectedItemExtendedInfoPtrOutput() AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return i.ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *azureIaaSVMProtectedItemExtendedInfoPtrType) ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectedItemExtendedInfoPtrOutput)
}

type AzureIaaSVMProtectedItemExtendedInfoOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectedItemExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectedItemExtendedInfo)(nil)).Elem()
}

func (o AzureIaaSVMProtectedItemExtendedInfoOutput) ToAzureIaaSVMProtectedItemExtendedInfoOutput() AzureIaaSVMProtectedItemExtendedInfoOutput {
	return o
}

func (o AzureIaaSVMProtectedItemExtendedInfoOutput) ToAzureIaaSVMProtectedItemExtendedInfoOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoOutput {
	return o
}

func (o AzureIaaSVMProtectedItemExtendedInfoOutput) ToAzureIaaSVMProtectedItemExtendedInfoPtrOutput() AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return o.ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (o AzureIaaSVMProtectedItemExtendedInfoOutput) ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureIaaSVMProtectedItemExtendedInfo) *AzureIaaSVMProtectedItemExtendedInfo {
		return &v
	}).(AzureIaaSVMProtectedItemExtendedInfoPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemExtendedInfo) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoOutput) PolicyInconsistent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemExtendedInfo) *bool { return v.PolicyInconsistent }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemExtendedInfo) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type AzureIaaSVMProtectedItemExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectedItemExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureIaaSVMProtectedItemExtendedInfo)(nil)).Elem()
}

func (o AzureIaaSVMProtectedItemExtendedInfoPtrOutput) ToAzureIaaSVMProtectedItemExtendedInfoPtrOutput() AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o AzureIaaSVMProtectedItemExtendedInfoPtrOutput) ToAzureIaaSVMProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o AzureIaaSVMProtectedItemExtendedInfoPtrOutput) Elem() AzureIaaSVMProtectedItemExtendedInfoOutput {
	return o.ApplyT(func(v *AzureIaaSVMProtectedItemExtendedInfo) AzureIaaSVMProtectedItemExtendedInfo {
		if v != nil {
			return *v
		}
		var ret AzureIaaSVMProtectedItemExtendedInfo
		return ret
	}).(AzureIaaSVMProtectedItemExtendedInfoOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoPtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureIaaSVMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoPtrOutput) PolicyInconsistent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureIaaSVMProtectedItemExtendedInfo) *bool {
		if v == nil {
			return nil
		}
		return v.PolicyInconsistent
	}).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoPtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureIaaSVMProtectedItemExtendedInfo) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type AzureIaaSVMProtectedItemExtendedInfoResponse struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyInconsistent  *bool   `pulumi:"policyInconsistent"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type AzureIaaSVMProtectedItemExtendedInfoResponseInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectedItemExtendedInfoResponseOutput() AzureIaaSVMProtectedItemExtendedInfoResponseOutput
	ToAzureIaaSVMProtectedItemExtendedInfoResponseOutputWithContext(context.Context) AzureIaaSVMProtectedItemExtendedInfoResponseOutput
}

type AzureIaaSVMProtectedItemExtendedInfoResponseArgs struct {
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	PolicyInconsistent  pulumi.BoolPtrInput   `pulumi:"policyInconsistent"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (AzureIaaSVMProtectedItemExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i AzureIaaSVMProtectedItemExtendedInfoResponseArgs) ToAzureIaaSVMProtectedItemExtendedInfoResponseOutput() AzureIaaSVMProtectedItemExtendedInfoResponseOutput {
	return i.ToAzureIaaSVMProtectedItemExtendedInfoResponseOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectedItemExtendedInfoResponseArgs) ToAzureIaaSVMProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectedItemExtendedInfoResponseOutput)
}

func (i AzureIaaSVMProtectedItemExtendedInfoResponseArgs) ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput() AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectedItemExtendedInfoResponseArgs) ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectedItemExtendedInfoResponseOutput).ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx)
}









type AzureIaaSVMProtectedItemExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput() AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput
	ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Context) AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput
}

type azureIaaSVMProtectedItemExtendedInfoResponsePtrType AzureIaaSVMProtectedItemExtendedInfoResponseArgs

func AzureIaaSVMProtectedItemExtendedInfoResponsePtr(v *AzureIaaSVMProtectedItemExtendedInfoResponseArgs) AzureIaaSVMProtectedItemExtendedInfoResponsePtrInput {
	return (*azureIaaSVMProtectedItemExtendedInfoResponsePtrType)(v)
}

func (*azureIaaSVMProtectedItemExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureIaaSVMProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i *azureIaaSVMProtectedItemExtendedInfoResponsePtrType) ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput() AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *azureIaaSVMProtectedItemExtendedInfoResponsePtrType) ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput)
}

type AzureIaaSVMProtectedItemExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectedItemExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponseOutput) ToAzureIaaSVMProtectedItemExtendedInfoResponseOutput() AzureIaaSVMProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponseOutput) ToAzureIaaSVMProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponseOutput) ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput() AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return o.ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponseOutput) ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureIaaSVMProtectedItemExtendedInfoResponse) *AzureIaaSVMProtectedItemExtendedInfoResponse {
		return &v
	}).(AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponseOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemExtendedInfoResponse) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponseOutput) PolicyInconsistent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemExtendedInfoResponse) *bool { return v.PolicyInconsistent }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponseOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemExtendedInfoResponse) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureIaaSVMProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput) ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput() AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput) ToAzureIaaSVMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput) Elem() AzureIaaSVMProtectedItemExtendedInfoResponseOutput {
	return o.ApplyT(func(v *AzureIaaSVMProtectedItemExtendedInfoResponse) AzureIaaSVMProtectedItemExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureIaaSVMProtectedItemExtendedInfoResponse
		return ret
	}).(AzureIaaSVMProtectedItemExtendedInfoResponseOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureIaaSVMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput) PolicyInconsistent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureIaaSVMProtectedItemExtendedInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PolicyInconsistent
	}).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureIaaSVMProtectedItemExtendedInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type AzureIaaSVMProtectedItemResponse struct {
	BackupManagementType             *string                                       `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                                       `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     *string                                       `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                       `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	ProtectedItemDataId              *string                                       `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                      `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                                       `pulumi:"virtualMachineId"`
	WorkloadType                     *string                                       `pulumi:"workloadType"`
}





type AzureIaaSVMProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectedItemResponseOutput() AzureIaaSVMProtectedItemResponseOutput
	ToAzureIaaSVMProtectedItemResponseOutputWithContext(context.Context) AzureIaaSVMProtectedItemResponseOutput
}

type AzureIaaSVMProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                                `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureIaaSVMProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	ExtendedProperties               ExtendedPropertiesResponsePtrInput                   `pulumi:"extendedProperties"`
	FriendlyName                     pulumi.StringPtrInput                                `pulumi:"friendlyName"`
	HealthDetails                    AzureIaaSVMHealthDetailsResponseArrayInput           `pulumi:"healthDetails"`
	HealthStatus                     pulumi.StringPtrInput                                `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                  `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                  `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                  `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsResponseMapInput             `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                                `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                                `pulumi:"policyId"`
	ProtectedItemDataId              pulumi.StringPtrInput                                `pulumi:"protectedItemDataId"`
	ProtectedItemType                pulumi.StringInput                                   `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                                `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                              `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                                `pulumi:"sourceResourceId"`
	VirtualMachineId                 pulumi.StringPtrInput                                `pulumi:"virtualMachineId"`
	WorkloadType                     pulumi.StringPtrInput                                `pulumi:"workloadType"`
}

func (AzureIaaSVMProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectedItemResponse)(nil)).Elem()
}

func (i AzureIaaSVMProtectedItemResponseArgs) ToAzureIaaSVMProtectedItemResponseOutput() AzureIaaSVMProtectedItemResponseOutput {
	return i.ToAzureIaaSVMProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectedItemResponseArgs) ToAzureIaaSVMProtectedItemResponseOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectedItemResponseOutput)
}

type AzureIaaSVMProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectedItemResponse)(nil)).Elem()
}

func (o AzureIaaSVMProtectedItemResponseOutput) ToAzureIaaSVMProtectedItemResponseOutput() AzureIaaSVMProtectedItemResponseOutput {
	return o
}

func (o AzureIaaSVMProtectedItemResponseOutput) ToAzureIaaSVMProtectedItemResponseOutputWithContext(ctx context.Context) AzureIaaSVMProtectedItemResponseOutput {
	return o
}

func (o AzureIaaSVMProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) ExtendedInfo() AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *AzureIaaSVMProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) ExtendedProperties() ExtendedPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *ExtendedPropertiesResponse { return v.ExtendedProperties }).(ExtendedPropertiesResponsePtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) HealthDetails() AzureIaaSVMHealthDetailsResponseArrayOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) []AzureIaaSVMHealthDetailsResponse { return v.HealthDetails }).(AzureIaaSVMHealthDetailsResponseArrayOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) KpisHealths() KPIResourceHealthDetailsResponseMapOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) map[string]KPIResourceHealthDetailsResponse {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsResponseMapOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) ProtectedItemDataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.ProtectedItemDataId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureIaaSVMProtectionPolicy struct {
	BackupManagementType           string                      `pulumi:"backupManagementType"`
	InstantRPDetails               *InstantRPAdditionalDetails `pulumi:"instantRPDetails"`
	InstantRpRetentionRangeInDays  *int                        `pulumi:"instantRpRetentionRangeInDays"`
	ProtectedItemsCount            *int                        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string                    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{}                 `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{}                 `pulumi:"schedulePolicy"`
	TimeZone                       *string                     `pulumi:"timeZone"`
}





type AzureIaaSVMProtectionPolicyInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectionPolicyOutput() AzureIaaSVMProtectionPolicyOutput
	ToAzureIaaSVMProtectionPolicyOutputWithContext(context.Context) AzureIaaSVMProtectionPolicyOutput
}

type AzureIaaSVMProtectionPolicyArgs struct {
	BackupManagementType           pulumi.StringInput                 `pulumi:"backupManagementType"`
	InstantRPDetails               InstantRPAdditionalDetailsPtrInput `pulumi:"instantRPDetails"`
	InstantRpRetentionRangeInDays  pulumi.IntPtrInput                 `pulumi:"instantRpRetentionRangeInDays"`
	ProtectedItemsCount            pulumi.IntPtrInput                 `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput            `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                pulumi.Input                       `pulumi:"retentionPolicy"`
	SchedulePolicy                 pulumi.Input                       `pulumi:"schedulePolicy"`
	TimeZone                       pulumi.StringPtrInput              `pulumi:"timeZone"`
}

func (AzureIaaSVMProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectionPolicy)(nil)).Elem()
}

func (i AzureIaaSVMProtectionPolicyArgs) ToAzureIaaSVMProtectionPolicyOutput() AzureIaaSVMProtectionPolicyOutput {
	return i.ToAzureIaaSVMProtectionPolicyOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectionPolicyArgs) ToAzureIaaSVMProtectionPolicyOutputWithContext(ctx context.Context) AzureIaaSVMProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectionPolicyOutput)
}

type AzureIaaSVMProtectionPolicyOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectionPolicy)(nil)).Elem()
}

func (o AzureIaaSVMProtectionPolicyOutput) ToAzureIaaSVMProtectionPolicyOutput() AzureIaaSVMProtectionPolicyOutput {
	return o
}

func (o AzureIaaSVMProtectionPolicyOutput) ToAzureIaaSVMProtectionPolicyOutputWithContext(ctx context.Context) AzureIaaSVMProtectionPolicyOutput {
	return o
}

func (o AzureIaaSVMProtectionPolicyOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) InstantRPDetails() InstantRPAdditionalDetailsPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) *InstantRPAdditionalDetails { return v.InstantRPDetails }).(InstantRPAdditionalDetailsPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) InstantRpRetentionRangeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) *int { return v.InstantRpRetentionRangeInDays }).(pulumi.IntPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type AzureIaaSVMProtectionPolicyResponse struct {
	BackupManagementType           string                              `pulumi:"backupManagementType"`
	InstantRPDetails               *InstantRPAdditionalDetailsResponse `pulumi:"instantRPDetails"`
	InstantRpRetentionRangeInDays  *int                                `pulumi:"instantRpRetentionRangeInDays"`
	ProtectedItemsCount            *int                                `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string                            `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{}                         `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{}                         `pulumi:"schedulePolicy"`
	TimeZone                       *string                             `pulumi:"timeZone"`
}





type AzureIaaSVMProtectionPolicyResponseInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectionPolicyResponseOutput() AzureIaaSVMProtectionPolicyResponseOutput
	ToAzureIaaSVMProtectionPolicyResponseOutputWithContext(context.Context) AzureIaaSVMProtectionPolicyResponseOutput
}

type AzureIaaSVMProtectionPolicyResponseArgs struct {
	BackupManagementType           pulumi.StringInput                         `pulumi:"backupManagementType"`
	InstantRPDetails               InstantRPAdditionalDetailsResponsePtrInput `pulumi:"instantRPDetails"`
	InstantRpRetentionRangeInDays  pulumi.IntPtrInput                         `pulumi:"instantRpRetentionRangeInDays"`
	ProtectedItemsCount            pulumi.IntPtrInput                         `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput                    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                pulumi.Input                               `pulumi:"retentionPolicy"`
	SchedulePolicy                 pulumi.Input                               `pulumi:"schedulePolicy"`
	TimeZone                       pulumi.StringPtrInput                      `pulumi:"timeZone"`
}

func (AzureIaaSVMProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectionPolicyResponse)(nil)).Elem()
}

func (i AzureIaaSVMProtectionPolicyResponseArgs) ToAzureIaaSVMProtectionPolicyResponseOutput() AzureIaaSVMProtectionPolicyResponseOutput {
	return i.ToAzureIaaSVMProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectionPolicyResponseArgs) ToAzureIaaSVMProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureIaaSVMProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectionPolicyResponseOutput)
}

type AzureIaaSVMProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectionPolicyResponse)(nil)).Elem()
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) ToAzureIaaSVMProtectionPolicyResponseOutput() AzureIaaSVMProtectionPolicyResponseOutput {
	return o
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) ToAzureIaaSVMProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureIaaSVMProtectionPolicyResponseOutput {
	return o
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) InstantRPDetails() InstantRPAdditionalDetailsResponsePtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) *InstantRPAdditionalDetailsResponse {
		return v.InstantRPDetails
	}).(InstantRPAdditionalDetailsResponsePtrOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) InstantRpRetentionRangeInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) *int { return v.InstantRpRetentionRangeInDays }).(pulumi.IntPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type AzureRecoveryServiceVaultProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureRecoveryServiceVaultProtectionIntentInput interface {
	pulumi.Input

	ToAzureRecoveryServiceVaultProtectionIntentOutput() AzureRecoveryServiceVaultProtectionIntentOutput
	ToAzureRecoveryServiceVaultProtectionIntentOutputWithContext(context.Context) AzureRecoveryServiceVaultProtectionIntentOutput
}

type AzureRecoveryServiceVaultProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureRecoveryServiceVaultProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntent)(nil)).Elem()
}

func (i AzureRecoveryServiceVaultProtectionIntentArgs) ToAzureRecoveryServiceVaultProtectionIntentOutput() AzureRecoveryServiceVaultProtectionIntentOutput {
	return i.ToAzureRecoveryServiceVaultProtectionIntentOutputWithContext(context.Background())
}

func (i AzureRecoveryServiceVaultProtectionIntentArgs) ToAzureRecoveryServiceVaultProtectionIntentOutputWithContext(ctx context.Context) AzureRecoveryServiceVaultProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRecoveryServiceVaultProtectionIntentOutput)
}

type AzureRecoveryServiceVaultProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureRecoveryServiceVaultProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntent)(nil)).Elem()
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ToAzureRecoveryServiceVaultProtectionIntentOutput() AzureRecoveryServiceVaultProtectionIntentOutput {
	return o
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ToAzureRecoveryServiceVaultProtectionIntentOutputWithContext(ctx context.Context) AzureRecoveryServiceVaultProtectionIntentOutput {
	return o
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureRecoveryServiceVaultProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureRecoveryServiceVaultProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureRecoveryServiceVaultProtectionIntentResponseOutput() AzureRecoveryServiceVaultProtectionIntentResponseOutput
	ToAzureRecoveryServiceVaultProtectionIntentResponseOutputWithContext(context.Context) AzureRecoveryServiceVaultProtectionIntentResponseOutput
}

type AzureRecoveryServiceVaultProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureRecoveryServiceVaultProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntentResponse)(nil)).Elem()
}

func (i AzureRecoveryServiceVaultProtectionIntentResponseArgs) ToAzureRecoveryServiceVaultProtectionIntentResponseOutput() AzureRecoveryServiceVaultProtectionIntentResponseOutput {
	return i.ToAzureRecoveryServiceVaultProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureRecoveryServiceVaultProtectionIntentResponseArgs) ToAzureRecoveryServiceVaultProtectionIntentResponseOutputWithContext(ctx context.Context) AzureRecoveryServiceVaultProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRecoveryServiceVaultProtectionIntentResponseOutput)
}

type AzureRecoveryServiceVaultProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureRecoveryServiceVaultProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntentResponse)(nil)).Elem()
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ToAzureRecoveryServiceVaultProtectionIntentResponseOutput() AzureRecoveryServiceVaultProtectionIntentResponseOutput {
	return o
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ToAzureRecoveryServiceVaultProtectionIntentResponseOutputWithContext(ctx context.Context) AzureRecoveryServiceVaultProtectionIntentResponseOutput {
	return o
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureResourceProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	FriendlyName             *string `pulumi:"friendlyName"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureResourceProtectionIntentInput interface {
	pulumi.Input

	ToAzureResourceProtectionIntentOutput() AzureResourceProtectionIntentOutput
	ToAzureResourceProtectionIntentOutputWithContext(context.Context) AzureResourceProtectionIntentOutput
}

type AzureResourceProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	FriendlyName             pulumi.StringPtrInput `pulumi:"friendlyName"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureResourceProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceProtectionIntent)(nil)).Elem()
}

func (i AzureResourceProtectionIntentArgs) ToAzureResourceProtectionIntentOutput() AzureResourceProtectionIntentOutput {
	return i.ToAzureResourceProtectionIntentOutputWithContext(context.Background())
}

func (i AzureResourceProtectionIntentArgs) ToAzureResourceProtectionIntentOutputWithContext(ctx context.Context) AzureResourceProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureResourceProtectionIntentOutput)
}

type AzureResourceProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureResourceProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceProtectionIntent)(nil)).Elem()
}

func (o AzureResourceProtectionIntentOutput) ToAzureResourceProtectionIntentOutput() AzureResourceProtectionIntentOutput {
	return o
}

func (o AzureResourceProtectionIntentOutput) ToAzureResourceProtectionIntentOutputWithContext(ctx context.Context) AzureResourceProtectionIntentOutput {
	return o
}

func (o AzureResourceProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureResourceProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureResourceProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	FriendlyName             *string `pulumi:"friendlyName"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureResourceProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureResourceProtectionIntentResponseOutput() AzureResourceProtectionIntentResponseOutput
	ToAzureResourceProtectionIntentResponseOutputWithContext(context.Context) AzureResourceProtectionIntentResponseOutput
}

type AzureResourceProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	FriendlyName             pulumi.StringPtrInput `pulumi:"friendlyName"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureResourceProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceProtectionIntentResponse)(nil)).Elem()
}

func (i AzureResourceProtectionIntentResponseArgs) ToAzureResourceProtectionIntentResponseOutput() AzureResourceProtectionIntentResponseOutput {
	return i.ToAzureResourceProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureResourceProtectionIntentResponseArgs) ToAzureResourceProtectionIntentResponseOutputWithContext(ctx context.Context) AzureResourceProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureResourceProtectionIntentResponseOutput)
}

type AzureResourceProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceProtectionIntentResponse)(nil)).Elem()
}

func (o AzureResourceProtectionIntentResponseOutput) ToAzureResourceProtectionIntentResponseOutput() AzureResourceProtectionIntentResponseOutput {
	return o
}

func (o AzureResourceProtectionIntentResponseOutput) ToAzureResourceProtectionIntentResponseOutputWithContext(ctx context.Context) AzureResourceProtectionIntentResponseOutput {
	return o
}

func (o AzureResourceProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureSQLAGWorkloadContainerProtectionContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        string                              `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}





type AzureSQLAGWorkloadContainerProtectionContainerInput interface {
	pulumi.Input

	ToAzureSQLAGWorkloadContainerProtectionContainerOutput() AzureSQLAGWorkloadContainerProtectionContainerOutput
	ToAzureSQLAGWorkloadContainerProtectionContainerOutputWithContext(context.Context) AzureSQLAGWorkloadContainerProtectionContainerOutput
}

type AzureSQLAGWorkloadContainerProtectionContainerArgs struct {
	BackupManagementType pulumi.StringPtrInput                      `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput                         `pulumi:"containerType"`
	ExtendedInfo         AzureWorkloadContainerExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput                      `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                      `pulumi:"healthStatus"`
	LastUpdatedTime      pulumi.StringPtrInput                      `pulumi:"lastUpdatedTime"`
	OperationType        pulumi.StringPtrInput                      `pulumi:"operationType"`
	RegistrationStatus   pulumi.StringPtrInput                      `pulumi:"registrationStatus"`
	SourceResourceId     pulumi.StringPtrInput                      `pulumi:"sourceResourceId"`
	WorkloadType         pulumi.StringPtrInput                      `pulumi:"workloadType"`
}

func (AzureSQLAGWorkloadContainerProtectionContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSQLAGWorkloadContainerProtectionContainer)(nil)).Elem()
}

func (i AzureSQLAGWorkloadContainerProtectionContainerArgs) ToAzureSQLAGWorkloadContainerProtectionContainerOutput() AzureSQLAGWorkloadContainerProtectionContainerOutput {
	return i.ToAzureSQLAGWorkloadContainerProtectionContainerOutputWithContext(context.Background())
}

func (i AzureSQLAGWorkloadContainerProtectionContainerArgs) ToAzureSQLAGWorkloadContainerProtectionContainerOutputWithContext(ctx context.Context) AzureSQLAGWorkloadContainerProtectionContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSQLAGWorkloadContainerProtectionContainerOutput)
}

type AzureSQLAGWorkloadContainerProtectionContainerOutput struct{ *pulumi.OutputState }

func (AzureSQLAGWorkloadContainerProtectionContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSQLAGWorkloadContainerProtectionContainer)(nil)).Elem()
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) ToAzureSQLAGWorkloadContainerProtectionContainerOutput() AzureSQLAGWorkloadContainerProtectionContainerOutput {
	return o
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) ToAzureSQLAGWorkloadContainerProtectionContainerOutputWithContext(ctx context.Context) AzureSQLAGWorkloadContainerProtectionContainerOutput {
	return o
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) ExtendedInfo() AzureWorkloadContainerExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *AzureWorkloadContainerExtendedInfo {
		return v.ExtendedInfo
	}).(AzureWorkloadContainerExtendedInfoPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *string { return v.OperationType }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainer) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureSQLAGWorkloadContainerProtectionContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        string                                      `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}





type AzureSQLAGWorkloadContainerProtectionContainerResponseInput interface {
	pulumi.Input

	ToAzureSQLAGWorkloadContainerProtectionContainerResponseOutput() AzureSQLAGWorkloadContainerProtectionContainerResponseOutput
	ToAzureSQLAGWorkloadContainerProtectionContainerResponseOutputWithContext(context.Context) AzureSQLAGWorkloadContainerProtectionContainerResponseOutput
}

type AzureSQLAGWorkloadContainerProtectionContainerResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput                              `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput                                 `pulumi:"containerType"`
	ExtendedInfo         AzureWorkloadContainerExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput                              `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                              `pulumi:"healthStatus"`
	LastUpdatedTime      pulumi.StringPtrInput                              `pulumi:"lastUpdatedTime"`
	OperationType        pulumi.StringPtrInput                              `pulumi:"operationType"`
	RegistrationStatus   pulumi.StringPtrInput                              `pulumi:"registrationStatus"`
	SourceResourceId     pulumi.StringPtrInput                              `pulumi:"sourceResourceId"`
	WorkloadType         pulumi.StringPtrInput                              `pulumi:"workloadType"`
}

func (AzureSQLAGWorkloadContainerProtectionContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSQLAGWorkloadContainerProtectionContainerResponse)(nil)).Elem()
}

func (i AzureSQLAGWorkloadContainerProtectionContainerResponseArgs) ToAzureSQLAGWorkloadContainerProtectionContainerResponseOutput() AzureSQLAGWorkloadContainerProtectionContainerResponseOutput {
	return i.ToAzureSQLAGWorkloadContainerProtectionContainerResponseOutputWithContext(context.Background())
}

func (i AzureSQLAGWorkloadContainerProtectionContainerResponseArgs) ToAzureSQLAGWorkloadContainerProtectionContainerResponseOutputWithContext(ctx context.Context) AzureSQLAGWorkloadContainerProtectionContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSQLAGWorkloadContainerProtectionContainerResponseOutput)
}

type AzureSQLAGWorkloadContainerProtectionContainerResponseOutput struct{ *pulumi.OutputState }

func (AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSQLAGWorkloadContainerProtectionContainerResponse)(nil)).Elem()
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) ToAzureSQLAGWorkloadContainerProtectionContainerResponseOutput() AzureSQLAGWorkloadContainerProtectionContainerResponseOutput {
	return o
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) ToAzureSQLAGWorkloadContainerProtectionContainerResponseOutputWithContext(ctx context.Context) AzureSQLAGWorkloadContainerProtectionContainerResponseOutput {
	return o
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) ExtendedInfo() AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *AzureWorkloadContainerExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureWorkloadContainerExtendedInfoResponsePtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *string { return v.OperationType }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureSQLAGWorkloadContainerProtectionContainerResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSQLAGWorkloadContainerProtectionContainerResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureSqlContainer struct {
	BackupManagementType *string `pulumi:"backupManagementType"`
	ContainerType        string  `pulumi:"containerType"`
	FriendlyName         *string `pulumi:"friendlyName"`
	HealthStatus         *string `pulumi:"healthStatus"`
	RegistrationStatus   *string `pulumi:"registrationStatus"`
}





type AzureSqlContainerInput interface {
	pulumi.Input

	ToAzureSqlContainerOutput() AzureSqlContainerOutput
	ToAzureSqlContainerOutputWithContext(context.Context) AzureSqlContainerOutput
}

type AzureSqlContainerArgs struct {
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput    `pulumi:"containerType"`
	FriendlyName         pulumi.StringPtrInput `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput `pulumi:"healthStatus"`
	RegistrationStatus   pulumi.StringPtrInput `pulumi:"registrationStatus"`
}

func (AzureSqlContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlContainer)(nil)).Elem()
}

func (i AzureSqlContainerArgs) ToAzureSqlContainerOutput() AzureSqlContainerOutput {
	return i.ToAzureSqlContainerOutputWithContext(context.Background())
}

func (i AzureSqlContainerArgs) ToAzureSqlContainerOutputWithContext(ctx context.Context) AzureSqlContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlContainerOutput)
}

type AzureSqlContainerOutput struct{ *pulumi.OutputState }

func (AzureSqlContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlContainer)(nil)).Elem()
}

func (o AzureSqlContainerOutput) ToAzureSqlContainerOutput() AzureSqlContainerOutput {
	return o
}

func (o AzureSqlContainerOutput) ToAzureSqlContainerOutputWithContext(ctx context.Context) AzureSqlContainerOutput {
	return o
}

func (o AzureSqlContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureSqlContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureSqlContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureSqlContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

type AzureSqlContainerResponse struct {
	BackupManagementType *string `pulumi:"backupManagementType"`
	ContainerType        string  `pulumi:"containerType"`
	FriendlyName         *string `pulumi:"friendlyName"`
	HealthStatus         *string `pulumi:"healthStatus"`
	RegistrationStatus   *string `pulumi:"registrationStatus"`
}





type AzureSqlContainerResponseInput interface {
	pulumi.Input

	ToAzureSqlContainerResponseOutput() AzureSqlContainerResponseOutput
	ToAzureSqlContainerResponseOutputWithContext(context.Context) AzureSqlContainerResponseOutput
}

type AzureSqlContainerResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput    `pulumi:"containerType"`
	FriendlyName         pulumi.StringPtrInput `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput `pulumi:"healthStatus"`
	RegistrationStatus   pulumi.StringPtrInput `pulumi:"registrationStatus"`
}

func (AzureSqlContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlContainerResponse)(nil)).Elem()
}

func (i AzureSqlContainerResponseArgs) ToAzureSqlContainerResponseOutput() AzureSqlContainerResponseOutput {
	return i.ToAzureSqlContainerResponseOutputWithContext(context.Background())
}

func (i AzureSqlContainerResponseArgs) ToAzureSqlContainerResponseOutputWithContext(ctx context.Context) AzureSqlContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlContainerResponseOutput)
}

type AzureSqlContainerResponseOutput struct{ *pulumi.OutputState }

func (AzureSqlContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlContainerResponse)(nil)).Elem()
}

func (o AzureSqlContainerResponseOutput) ToAzureSqlContainerResponseOutput() AzureSqlContainerResponseOutput {
	return o
}

func (o AzureSqlContainerResponseOutput) ToAzureSqlContainerResponseOutputWithContext(ctx context.Context) AzureSqlContainerResponseOutput {
	return o
}

func (o AzureSqlContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureSqlContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureSqlContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureSqlContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

type AzureSqlProtectedItem struct {
	BackupManagementType             *string                            `pulumi:"backupManagementType"`
	BackupSetName                    *string                            `pulumi:"backupSetName"`
	ContainerName                    *string                            `pulumi:"containerName"`
	CreateMode                       *string                            `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                            `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                            `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureSqlProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	IsDeferredDeleteScheduleUpcoming *bool                              `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                              `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                              `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                            `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                            `pulumi:"policyId"`
	ProtectedItemDataId              *string                            `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                             `pulumi:"protectedItemType"`
	ProtectionState                  *string                            `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                           `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                            `pulumi:"sourceResourceId"`
	WorkloadType                     *string                            `pulumi:"workloadType"`
}





type AzureSqlProtectedItemInput interface {
	pulumi.Input

	ToAzureSqlProtectedItemOutput() AzureSqlProtectedItemOutput
	ToAzureSqlProtectedItemOutputWithContext(context.Context) AzureSqlProtectedItemOutput
}

type AzureSqlProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                     `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                     `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                     `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                     `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                     `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                     `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureSqlProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                       `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                       `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                       `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                pulumi.StringPtrInput                     `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                     `pulumi:"policyId"`
	ProtectedItemDataId              pulumi.StringPtrInput                     `pulumi:"protectedItemDataId"`
	ProtectedItemType                pulumi.StringInput                        `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                     `pulumi:"protectionState"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                   `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                     `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                     `pulumi:"workloadType"`
}

func (AzureSqlProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectedItem)(nil)).Elem()
}

func (i AzureSqlProtectedItemArgs) ToAzureSqlProtectedItemOutput() AzureSqlProtectedItemOutput {
	return i.ToAzureSqlProtectedItemOutputWithContext(context.Background())
}

func (i AzureSqlProtectedItemArgs) ToAzureSqlProtectedItemOutputWithContext(ctx context.Context) AzureSqlProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectedItemOutput)
}

type AzureSqlProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectedItem)(nil)).Elem()
}

func (o AzureSqlProtectedItemOutput) ToAzureSqlProtectedItemOutput() AzureSqlProtectedItemOutput {
	return o
}

func (o AzureSqlProtectedItemOutput) ToAzureSqlProtectedItemOutputWithContext(ctx context.Context) AzureSqlProtectedItemOutput {
	return o
}

func (o AzureSqlProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) ExtendedInfo() AzureSqlProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *AzureSqlProtectedItemExtendedInfo { return v.ExtendedInfo }).(AzureSqlProtectedItemExtendedInfoPtrOutput)
}

func (o AzureSqlProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureSqlProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureSqlProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureSqlProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) ProtectedItemDataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.ProtectedItemDataId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureSqlProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureSqlProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureSqlProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type AzureSqlProtectedItemExtendedInfoInput interface {
	pulumi.Input

	ToAzureSqlProtectedItemExtendedInfoOutput() AzureSqlProtectedItemExtendedInfoOutput
	ToAzureSqlProtectedItemExtendedInfoOutputWithContext(context.Context) AzureSqlProtectedItemExtendedInfoOutput
}

type AzureSqlProtectedItemExtendedInfoArgs struct {
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	PolicyState         pulumi.StringPtrInput `pulumi:"policyState"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (AzureSqlProtectedItemExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectedItemExtendedInfo)(nil)).Elem()
}

func (i AzureSqlProtectedItemExtendedInfoArgs) ToAzureSqlProtectedItemExtendedInfoOutput() AzureSqlProtectedItemExtendedInfoOutput {
	return i.ToAzureSqlProtectedItemExtendedInfoOutputWithContext(context.Background())
}

func (i AzureSqlProtectedItemExtendedInfoArgs) ToAzureSqlProtectedItemExtendedInfoOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectedItemExtendedInfoOutput)
}

func (i AzureSqlProtectedItemExtendedInfoArgs) ToAzureSqlProtectedItemExtendedInfoPtrOutput() AzureSqlProtectedItemExtendedInfoPtrOutput {
	return i.ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i AzureSqlProtectedItemExtendedInfoArgs) ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectedItemExtendedInfoOutput).ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(ctx)
}









type AzureSqlProtectedItemExtendedInfoPtrInput interface {
	pulumi.Input

	ToAzureSqlProtectedItemExtendedInfoPtrOutput() AzureSqlProtectedItemExtendedInfoPtrOutput
	ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(context.Context) AzureSqlProtectedItemExtendedInfoPtrOutput
}

type azureSqlProtectedItemExtendedInfoPtrType AzureSqlProtectedItemExtendedInfoArgs

func AzureSqlProtectedItemExtendedInfoPtr(v *AzureSqlProtectedItemExtendedInfoArgs) AzureSqlProtectedItemExtendedInfoPtrInput {
	return (*azureSqlProtectedItemExtendedInfoPtrType)(v)
}

func (*azureSqlProtectedItemExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSqlProtectedItemExtendedInfo)(nil)).Elem()
}

func (i *azureSqlProtectedItemExtendedInfoPtrType) ToAzureSqlProtectedItemExtendedInfoPtrOutput() AzureSqlProtectedItemExtendedInfoPtrOutput {
	return i.ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *azureSqlProtectedItemExtendedInfoPtrType) ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectedItemExtendedInfoPtrOutput)
}

type AzureSqlProtectedItemExtendedInfoOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectedItemExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectedItemExtendedInfo)(nil)).Elem()
}

func (o AzureSqlProtectedItemExtendedInfoOutput) ToAzureSqlProtectedItemExtendedInfoOutput() AzureSqlProtectedItemExtendedInfoOutput {
	return o
}

func (o AzureSqlProtectedItemExtendedInfoOutput) ToAzureSqlProtectedItemExtendedInfoOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoOutput {
	return o
}

func (o AzureSqlProtectedItemExtendedInfoOutput) ToAzureSqlProtectedItemExtendedInfoPtrOutput() AzureSqlProtectedItemExtendedInfoPtrOutput {
	return o.ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (o AzureSqlProtectedItemExtendedInfoOutput) ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSqlProtectedItemExtendedInfo) *AzureSqlProtectedItemExtendedInfo {
		return &v
	}).(AzureSqlProtectedItemExtendedInfoPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemExtendedInfo) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemExtendedInfo) *string { return v.PolicyState }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemExtendedInfo) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type AzureSqlProtectedItemExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectedItemExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSqlProtectedItemExtendedInfo)(nil)).Elem()
}

func (o AzureSqlProtectedItemExtendedInfoPtrOutput) ToAzureSqlProtectedItemExtendedInfoPtrOutput() AzureSqlProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o AzureSqlProtectedItemExtendedInfoPtrOutput) ToAzureSqlProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o AzureSqlProtectedItemExtendedInfoPtrOutput) Elem() AzureSqlProtectedItemExtendedInfoOutput {
	return o.ApplyT(func(v *AzureSqlProtectedItemExtendedInfo) AzureSqlProtectedItemExtendedInfo {
		if v != nil {
			return *v
		}
		var ret AzureSqlProtectedItemExtendedInfo
		return ret
	}).(AzureSqlProtectedItemExtendedInfoOutput)
}

func (o AzureSqlProtectedItemExtendedInfoPtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoPtrOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.PolicyState
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoPtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureSqlProtectedItemExtendedInfo) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type AzureSqlProtectedItemExtendedInfoResponse struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type AzureSqlProtectedItemExtendedInfoResponseInput interface {
	pulumi.Input

	ToAzureSqlProtectedItemExtendedInfoResponseOutput() AzureSqlProtectedItemExtendedInfoResponseOutput
	ToAzureSqlProtectedItemExtendedInfoResponseOutputWithContext(context.Context) AzureSqlProtectedItemExtendedInfoResponseOutput
}

type AzureSqlProtectedItemExtendedInfoResponseArgs struct {
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	PolicyState         pulumi.StringPtrInput `pulumi:"policyState"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (AzureSqlProtectedItemExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i AzureSqlProtectedItemExtendedInfoResponseArgs) ToAzureSqlProtectedItemExtendedInfoResponseOutput() AzureSqlProtectedItemExtendedInfoResponseOutput {
	return i.ToAzureSqlProtectedItemExtendedInfoResponseOutputWithContext(context.Background())
}

func (i AzureSqlProtectedItemExtendedInfoResponseArgs) ToAzureSqlProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectedItemExtendedInfoResponseOutput)
}

func (i AzureSqlProtectedItemExtendedInfoResponseArgs) ToAzureSqlProtectedItemExtendedInfoResponsePtrOutput() AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i AzureSqlProtectedItemExtendedInfoResponseArgs) ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectedItemExtendedInfoResponseOutput).ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx)
}









type AzureSqlProtectedItemExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToAzureSqlProtectedItemExtendedInfoResponsePtrOutput() AzureSqlProtectedItemExtendedInfoResponsePtrOutput
	ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Context) AzureSqlProtectedItemExtendedInfoResponsePtrOutput
}

type azureSqlProtectedItemExtendedInfoResponsePtrType AzureSqlProtectedItemExtendedInfoResponseArgs

func AzureSqlProtectedItemExtendedInfoResponsePtr(v *AzureSqlProtectedItemExtendedInfoResponseArgs) AzureSqlProtectedItemExtendedInfoResponsePtrInput {
	return (*azureSqlProtectedItemExtendedInfoResponsePtrType)(v)
}

func (*azureSqlProtectedItemExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSqlProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i *azureSqlProtectedItemExtendedInfoResponsePtrType) ToAzureSqlProtectedItemExtendedInfoResponsePtrOutput() AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *azureSqlProtectedItemExtendedInfoResponsePtrType) ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectedItemExtendedInfoResponsePtrOutput)
}

type AzureSqlProtectedItemExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectedItemExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o AzureSqlProtectedItemExtendedInfoResponseOutput) ToAzureSqlProtectedItemExtendedInfoResponseOutput() AzureSqlProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o AzureSqlProtectedItemExtendedInfoResponseOutput) ToAzureSqlProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o AzureSqlProtectedItemExtendedInfoResponseOutput) ToAzureSqlProtectedItemExtendedInfoResponsePtrOutput() AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return o.ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o AzureSqlProtectedItemExtendedInfoResponseOutput) ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSqlProtectedItemExtendedInfoResponse) *AzureSqlProtectedItemExtendedInfoResponse {
		return &v
	}).(AzureSqlProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoResponseOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemExtendedInfoResponse) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoResponseOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemExtendedInfoResponse) *string { return v.PolicyState }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoResponseOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemExtendedInfoResponse) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type AzureSqlProtectedItemExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectedItemExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSqlProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o AzureSqlProtectedItemExtendedInfoResponsePtrOutput) ToAzureSqlProtectedItemExtendedInfoResponsePtrOutput() AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureSqlProtectedItemExtendedInfoResponsePtrOutput) ToAzureSqlProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureSqlProtectedItemExtendedInfoResponsePtrOutput) Elem() AzureSqlProtectedItemExtendedInfoResponseOutput {
	return o.ApplyT(func(v *AzureSqlProtectedItemExtendedInfoResponse) AzureSqlProtectedItemExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureSqlProtectedItemExtendedInfoResponse
		return ret
	}).(AzureSqlProtectedItemExtendedInfoResponseOutput)
}

func (o AzureSqlProtectedItemExtendedInfoResponsePtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoResponsePtrOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyState
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemExtendedInfoResponsePtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureSqlProtectedItemExtendedInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type AzureSqlProtectedItemResponse struct {
	BackupManagementType             *string                                    `pulumi:"backupManagementType"`
	BackupSetName                    *string                                    `pulumi:"backupSetName"`
	ContainerName                    *string                                    `pulumi:"containerName"`
	CreateMode                       *string                                    `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                    `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                    `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureSqlProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	IsDeferredDeleteScheduleUpcoming *bool                                      `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                      `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                      `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                                    `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                    `pulumi:"policyId"`
	ProtectedItemDataId              *string                                    `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                     `pulumi:"protectedItemType"`
	ProtectionState                  *string                                    `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                   `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                                    `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                    `pulumi:"workloadType"`
}





type AzureSqlProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureSqlProtectedItemResponseOutput() AzureSqlProtectedItemResponseOutput
	ToAzureSqlProtectedItemResponseOutputWithContext(context.Context) AzureSqlProtectedItemResponseOutput
}

type AzureSqlProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                             `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                             `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                             `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                             `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                             `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                             `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureSqlProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                               `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                               `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                               `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                pulumi.StringPtrInput                             `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                             `pulumi:"policyId"`
	ProtectedItemDataId              pulumi.StringPtrInput                             `pulumi:"protectedItemDataId"`
	ProtectedItemType                pulumi.StringInput                                `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                             `pulumi:"protectionState"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                           `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                             `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                             `pulumi:"workloadType"`
}

func (AzureSqlProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectedItemResponse)(nil)).Elem()
}

func (i AzureSqlProtectedItemResponseArgs) ToAzureSqlProtectedItemResponseOutput() AzureSqlProtectedItemResponseOutput {
	return i.ToAzureSqlProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureSqlProtectedItemResponseArgs) ToAzureSqlProtectedItemResponseOutputWithContext(ctx context.Context) AzureSqlProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectedItemResponseOutput)
}

type AzureSqlProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectedItemResponse)(nil)).Elem()
}

func (o AzureSqlProtectedItemResponseOutput) ToAzureSqlProtectedItemResponseOutput() AzureSqlProtectedItemResponseOutput {
	return o
}

func (o AzureSqlProtectedItemResponseOutput) ToAzureSqlProtectedItemResponseOutputWithContext(ctx context.Context) AzureSqlProtectedItemResponseOutput {
	return o
}

func (o AzureSqlProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) ExtendedInfo() AzureSqlProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *AzureSqlProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureSqlProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) ProtectedItemDataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.ProtectedItemDataId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureSqlProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureSqlProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureSqlProtectionPolicy struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
}





type AzureSqlProtectionPolicyInput interface {
	pulumi.Input

	ToAzureSqlProtectionPolicyOutput() AzureSqlProtectionPolicyOutput
	ToAzureSqlProtectionPolicyOutputWithContext(context.Context) AzureSqlProtectionPolicyOutput
}

type AzureSqlProtectionPolicyArgs struct {
	BackupManagementType           pulumi.StringInput      `pulumi:"backupManagementType"`
	ProtectedItemsCount            pulumi.IntPtrInput      `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                pulumi.Input            `pulumi:"retentionPolicy"`
}

func (AzureSqlProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectionPolicy)(nil)).Elem()
}

func (i AzureSqlProtectionPolicyArgs) ToAzureSqlProtectionPolicyOutput() AzureSqlProtectionPolicyOutput {
	return i.ToAzureSqlProtectionPolicyOutputWithContext(context.Background())
}

func (i AzureSqlProtectionPolicyArgs) ToAzureSqlProtectionPolicyOutputWithContext(ctx context.Context) AzureSqlProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectionPolicyOutput)
}

type AzureSqlProtectionPolicyOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectionPolicy)(nil)).Elem()
}

func (o AzureSqlProtectionPolicyOutput) ToAzureSqlProtectionPolicyOutput() AzureSqlProtectionPolicyOutput {
	return o
}

func (o AzureSqlProtectionPolicyOutput) ToAzureSqlProtectionPolicyOutputWithContext(ctx context.Context) AzureSqlProtectionPolicyOutput {
	return o
}

func (o AzureSqlProtectionPolicyOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicy) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o AzureSqlProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureSqlProtectionPolicyOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicy) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureSqlProtectionPolicyOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicy) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

type AzureSqlProtectionPolicyResponse struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
}





type AzureSqlProtectionPolicyResponseInput interface {
	pulumi.Input

	ToAzureSqlProtectionPolicyResponseOutput() AzureSqlProtectionPolicyResponseOutput
	ToAzureSqlProtectionPolicyResponseOutputWithContext(context.Context) AzureSqlProtectionPolicyResponseOutput
}

type AzureSqlProtectionPolicyResponseArgs struct {
	BackupManagementType           pulumi.StringInput      `pulumi:"backupManagementType"`
	ProtectedItemsCount            pulumi.IntPtrInput      `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                pulumi.Input            `pulumi:"retentionPolicy"`
}

func (AzureSqlProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectionPolicyResponse)(nil)).Elem()
}

func (i AzureSqlProtectionPolicyResponseArgs) ToAzureSqlProtectionPolicyResponseOutput() AzureSqlProtectionPolicyResponseOutput {
	return i.ToAzureSqlProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i AzureSqlProtectionPolicyResponseArgs) ToAzureSqlProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureSqlProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectionPolicyResponseOutput)
}

type AzureSqlProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectionPolicyResponse)(nil)).Elem()
}

func (o AzureSqlProtectionPolicyResponseOutput) ToAzureSqlProtectionPolicyResponseOutput() AzureSqlProtectionPolicyResponseOutput {
	return o
}

func (o AzureSqlProtectionPolicyResponseOutput) ToAzureSqlProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureSqlProtectionPolicyResponseOutput {
	return o
}

func (o AzureSqlProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicyResponse) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o AzureSqlProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureSqlProtectionPolicyResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicyResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureSqlProtectionPolicyResponseOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicyResponse) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

type AzureStorageContainer struct {
	AcquireStorageAccountLock *string  `pulumi:"acquireStorageAccountLock"`
	BackupManagementType      *string  `pulumi:"backupManagementType"`
	ContainerType             string   `pulumi:"containerType"`
	FriendlyName              *string  `pulumi:"friendlyName"`
	HealthStatus              *string  `pulumi:"healthStatus"`
	ProtectedItemCount        *float64 `pulumi:"protectedItemCount"`
	RegistrationStatus        *string  `pulumi:"registrationStatus"`
	ResourceGroup             *string  `pulumi:"resourceGroup"`
	SourceResourceId          *string  `pulumi:"sourceResourceId"`
	StorageAccountVersion     *string  `pulumi:"storageAccountVersion"`
}





type AzureStorageContainerInput interface {
	pulumi.Input

	ToAzureStorageContainerOutput() AzureStorageContainerOutput
	ToAzureStorageContainerOutputWithContext(context.Context) AzureStorageContainerOutput
}

type AzureStorageContainerArgs struct {
	AcquireStorageAccountLock pulumi.StringPtrInput  `pulumi:"acquireStorageAccountLock"`
	BackupManagementType      pulumi.StringPtrInput  `pulumi:"backupManagementType"`
	ContainerType             pulumi.StringInput     `pulumi:"containerType"`
	FriendlyName              pulumi.StringPtrInput  `pulumi:"friendlyName"`
	HealthStatus              pulumi.StringPtrInput  `pulumi:"healthStatus"`
	ProtectedItemCount        pulumi.Float64PtrInput `pulumi:"protectedItemCount"`
	RegistrationStatus        pulumi.StringPtrInput  `pulumi:"registrationStatus"`
	ResourceGroup             pulumi.StringPtrInput  `pulumi:"resourceGroup"`
	SourceResourceId          pulumi.StringPtrInput  `pulumi:"sourceResourceId"`
	StorageAccountVersion     pulumi.StringPtrInput  `pulumi:"storageAccountVersion"`
}

func (AzureStorageContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageContainer)(nil)).Elem()
}

func (i AzureStorageContainerArgs) ToAzureStorageContainerOutput() AzureStorageContainerOutput {
	return i.ToAzureStorageContainerOutputWithContext(context.Background())
}

func (i AzureStorageContainerArgs) ToAzureStorageContainerOutputWithContext(ctx context.Context) AzureStorageContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStorageContainerOutput)
}

type AzureStorageContainerOutput struct{ *pulumi.OutputState }

func (AzureStorageContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageContainer)(nil)).Elem()
}

func (o AzureStorageContainerOutput) ToAzureStorageContainerOutput() AzureStorageContainerOutput {
	return o
}

func (o AzureStorageContainerOutput) ToAzureStorageContainerOutputWithContext(ctx context.Context) AzureStorageContainerOutput {
	return o
}

func (o AzureStorageContainerOutput) AcquireStorageAccountLock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *string { return v.AcquireStorageAccountLock }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureStorageContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureStorageContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerOutput) ProtectedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *float64 { return v.ProtectedItemCount }).(pulumi.Float64PtrOutput)
}

func (o AzureStorageContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerOutput) StorageAccountVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainer) *string { return v.StorageAccountVersion }).(pulumi.StringPtrOutput)
}

type AzureStorageContainerResponse struct {
	AcquireStorageAccountLock *string  `pulumi:"acquireStorageAccountLock"`
	BackupManagementType      *string  `pulumi:"backupManagementType"`
	ContainerType             string   `pulumi:"containerType"`
	FriendlyName              *string  `pulumi:"friendlyName"`
	HealthStatus              *string  `pulumi:"healthStatus"`
	ProtectedItemCount        *float64 `pulumi:"protectedItemCount"`
	RegistrationStatus        *string  `pulumi:"registrationStatus"`
	ResourceGroup             *string  `pulumi:"resourceGroup"`
	SourceResourceId          *string  `pulumi:"sourceResourceId"`
	StorageAccountVersion     *string  `pulumi:"storageAccountVersion"`
}





type AzureStorageContainerResponseInput interface {
	pulumi.Input

	ToAzureStorageContainerResponseOutput() AzureStorageContainerResponseOutput
	ToAzureStorageContainerResponseOutputWithContext(context.Context) AzureStorageContainerResponseOutput
}

type AzureStorageContainerResponseArgs struct {
	AcquireStorageAccountLock pulumi.StringPtrInput  `pulumi:"acquireStorageAccountLock"`
	BackupManagementType      pulumi.StringPtrInput  `pulumi:"backupManagementType"`
	ContainerType             pulumi.StringInput     `pulumi:"containerType"`
	FriendlyName              pulumi.StringPtrInput  `pulumi:"friendlyName"`
	HealthStatus              pulumi.StringPtrInput  `pulumi:"healthStatus"`
	ProtectedItemCount        pulumi.Float64PtrInput `pulumi:"protectedItemCount"`
	RegistrationStatus        pulumi.StringPtrInput  `pulumi:"registrationStatus"`
	ResourceGroup             pulumi.StringPtrInput  `pulumi:"resourceGroup"`
	SourceResourceId          pulumi.StringPtrInput  `pulumi:"sourceResourceId"`
	StorageAccountVersion     pulumi.StringPtrInput  `pulumi:"storageAccountVersion"`
}

func (AzureStorageContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageContainerResponse)(nil)).Elem()
}

func (i AzureStorageContainerResponseArgs) ToAzureStorageContainerResponseOutput() AzureStorageContainerResponseOutput {
	return i.ToAzureStorageContainerResponseOutputWithContext(context.Background())
}

func (i AzureStorageContainerResponseArgs) ToAzureStorageContainerResponseOutputWithContext(ctx context.Context) AzureStorageContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStorageContainerResponseOutput)
}

type AzureStorageContainerResponseOutput struct{ *pulumi.OutputState }

func (AzureStorageContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageContainerResponse)(nil)).Elem()
}

func (o AzureStorageContainerResponseOutput) ToAzureStorageContainerResponseOutput() AzureStorageContainerResponseOutput {
	return o
}

func (o AzureStorageContainerResponseOutput) ToAzureStorageContainerResponseOutputWithContext(ctx context.Context) AzureStorageContainerResponseOutput {
	return o
}

func (o AzureStorageContainerResponseOutput) AcquireStorageAccountLock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *string { return v.AcquireStorageAccountLock }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureStorageContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerResponseOutput) ProtectedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *float64 { return v.ProtectedItemCount }).(pulumi.Float64PtrOutput)
}

func (o AzureStorageContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureStorageContainerResponseOutput) StorageAccountVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageContainerResponse) *string { return v.StorageAccountVersion }).(pulumi.StringPtrOutput)
}

type AzureVMAppContainerProtectionContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        string                              `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}





type AzureVMAppContainerProtectionContainerInput interface {
	pulumi.Input

	ToAzureVMAppContainerProtectionContainerOutput() AzureVMAppContainerProtectionContainerOutput
	ToAzureVMAppContainerProtectionContainerOutputWithContext(context.Context) AzureVMAppContainerProtectionContainerOutput
}

type AzureVMAppContainerProtectionContainerArgs struct {
	BackupManagementType pulumi.StringPtrInput                      `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput                         `pulumi:"containerType"`
	ExtendedInfo         AzureWorkloadContainerExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput                      `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                      `pulumi:"healthStatus"`
	LastUpdatedTime      pulumi.StringPtrInput                      `pulumi:"lastUpdatedTime"`
	OperationType        pulumi.StringPtrInput                      `pulumi:"operationType"`
	RegistrationStatus   pulumi.StringPtrInput                      `pulumi:"registrationStatus"`
	SourceResourceId     pulumi.StringPtrInput                      `pulumi:"sourceResourceId"`
	WorkloadType         pulumi.StringPtrInput                      `pulumi:"workloadType"`
}

func (AzureVMAppContainerProtectionContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVMAppContainerProtectionContainer)(nil)).Elem()
}

func (i AzureVMAppContainerProtectionContainerArgs) ToAzureVMAppContainerProtectionContainerOutput() AzureVMAppContainerProtectionContainerOutput {
	return i.ToAzureVMAppContainerProtectionContainerOutputWithContext(context.Background())
}

func (i AzureVMAppContainerProtectionContainerArgs) ToAzureVMAppContainerProtectionContainerOutputWithContext(ctx context.Context) AzureVMAppContainerProtectionContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVMAppContainerProtectionContainerOutput)
}

type AzureVMAppContainerProtectionContainerOutput struct{ *pulumi.OutputState }

func (AzureVMAppContainerProtectionContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVMAppContainerProtectionContainer)(nil)).Elem()
}

func (o AzureVMAppContainerProtectionContainerOutput) ToAzureVMAppContainerProtectionContainerOutput() AzureVMAppContainerProtectionContainerOutput {
	return o
}

func (o AzureVMAppContainerProtectionContainerOutput) ToAzureVMAppContainerProtectionContainerOutputWithContext(ctx context.Context) AzureVMAppContainerProtectionContainerOutput {
	return o
}

func (o AzureVMAppContainerProtectionContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) ExtendedInfo() AzureWorkloadContainerExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *AzureWorkloadContainerExtendedInfo {
		return v.ExtendedInfo
	}).(AzureWorkloadContainerExtendedInfoPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *string { return v.OperationType }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainer) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVMAppContainerProtectionContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        string                                      `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}





type AzureVMAppContainerProtectionContainerResponseInput interface {
	pulumi.Input

	ToAzureVMAppContainerProtectionContainerResponseOutput() AzureVMAppContainerProtectionContainerResponseOutput
	ToAzureVMAppContainerProtectionContainerResponseOutputWithContext(context.Context) AzureVMAppContainerProtectionContainerResponseOutput
}

type AzureVMAppContainerProtectionContainerResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput                              `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput                                 `pulumi:"containerType"`
	ExtendedInfo         AzureWorkloadContainerExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput                              `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                              `pulumi:"healthStatus"`
	LastUpdatedTime      pulumi.StringPtrInput                              `pulumi:"lastUpdatedTime"`
	OperationType        pulumi.StringPtrInput                              `pulumi:"operationType"`
	RegistrationStatus   pulumi.StringPtrInput                              `pulumi:"registrationStatus"`
	SourceResourceId     pulumi.StringPtrInput                              `pulumi:"sourceResourceId"`
	WorkloadType         pulumi.StringPtrInput                              `pulumi:"workloadType"`
}

func (AzureVMAppContainerProtectionContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVMAppContainerProtectionContainerResponse)(nil)).Elem()
}

func (i AzureVMAppContainerProtectionContainerResponseArgs) ToAzureVMAppContainerProtectionContainerResponseOutput() AzureVMAppContainerProtectionContainerResponseOutput {
	return i.ToAzureVMAppContainerProtectionContainerResponseOutputWithContext(context.Background())
}

func (i AzureVMAppContainerProtectionContainerResponseArgs) ToAzureVMAppContainerProtectionContainerResponseOutputWithContext(ctx context.Context) AzureVMAppContainerProtectionContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVMAppContainerProtectionContainerResponseOutput)
}

type AzureVMAppContainerProtectionContainerResponseOutput struct{ *pulumi.OutputState }

func (AzureVMAppContainerProtectionContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVMAppContainerProtectionContainerResponse)(nil)).Elem()
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) ToAzureVMAppContainerProtectionContainerResponseOutput() AzureVMAppContainerProtectionContainerResponseOutput {
	return o
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) ToAzureVMAppContainerProtectionContainerResponseOutputWithContext(ctx context.Context) AzureVMAppContainerProtectionContainerResponseOutput {
	return o
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) ExtendedInfo() AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *AzureWorkloadContainerExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureWorkloadContainerExtendedInfoResponsePtrOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *string { return v.OperationType }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVMAppContainerProtectionContainerResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVMAppContainerProtectionContainerResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadProtectedItem struct {
	BackupManagementType             *string                                   `pulumi:"backupManagementType"`
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails       `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                   `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                   `pulumi:"parentName"`
	ParentType                       *string                                   `pulumi:"parentType"`
	PolicyId                         *string                                   `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ProtectionStatus                 *string                                   `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                   `pulumi:"workloadType"`
}





type AzureVmWorkloadProtectedItemInput interface {
	pulumi.Input

	ToAzureVmWorkloadProtectedItemOutput() AzureVmWorkloadProtectedItemOutput
	ToAzureVmWorkloadProtectedItemOutputWithContext(context.Context) AzureVmWorkloadProtectedItemOutput
}

type AzureVmWorkloadProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                            `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                            `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                            `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                            `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                            `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                            `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureVmWorkloadProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                            `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                              `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                              `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                              `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsMapInput                 `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                            `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                            `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                            `pulumi:"lastRecoveryPoint"`
	ParentName                       pulumi.StringPtrInput                            `pulumi:"parentName"`
	ParentType                       pulumi.StringPtrInput                            `pulumi:"parentType"`
	PolicyId                         pulumi.StringPtrInput                            `pulumi:"policyId"`
	ProtectedItemDataSourceId        pulumi.StringPtrInput                            `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        pulumi.StringPtrInput                            `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                pulumi.StringInput                               `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                            `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       pulumi.StringPtrInput                            `pulumi:"serverName"`
	SourceResourceId                 pulumi.StringPtrInput                            `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                            `pulumi:"workloadType"`
}

func (AzureVmWorkloadProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectedItem)(nil)).Elem()
}

func (i AzureVmWorkloadProtectedItemArgs) ToAzureVmWorkloadProtectedItemOutput() AzureVmWorkloadProtectedItemOutput {
	return i.ToAzureVmWorkloadProtectedItemOutputWithContext(context.Background())
}

func (i AzureVmWorkloadProtectedItemArgs) ToAzureVmWorkloadProtectedItemOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectedItemOutput)
}

type AzureVmWorkloadProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectedItem)(nil)).Elem()
}

func (o AzureVmWorkloadProtectedItemOutput) ToAzureVmWorkloadProtectedItemOutput() AzureVmWorkloadProtectedItemOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemOutput) ToAzureVmWorkloadProtectedItemOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ExtendedInfo() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *AzureVmWorkloadProtectedItemExtendedInfo { return v.ExtendedInfo }).(AzureVmWorkloadProtectedItemExtendedInfoPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) KpisHealths() KPIResourceHealthDetailsMapOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) map[string]KPIResourceHealthDetails { return v.KpisHealths }).(KPIResourceHealthDetailsMapOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ParentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.ParentType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ProtectedItemDataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.ProtectedItemDataSourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ProtectedItemHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.ProtectedItemHealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type AzureVmWorkloadProtectedItemExtendedInfoInput interface {
	pulumi.Input

	ToAzureVmWorkloadProtectedItemExtendedInfoOutput() AzureVmWorkloadProtectedItemExtendedInfoOutput
	ToAzureVmWorkloadProtectedItemExtendedInfoOutputWithContext(context.Context) AzureVmWorkloadProtectedItemExtendedInfoOutput
}

type AzureVmWorkloadProtectedItemExtendedInfoArgs struct {
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	PolicyState         pulumi.StringPtrInput `pulumi:"policyState"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (AzureVmWorkloadProtectedItemExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectedItemExtendedInfo)(nil)).Elem()
}

func (i AzureVmWorkloadProtectedItemExtendedInfoArgs) ToAzureVmWorkloadProtectedItemExtendedInfoOutput() AzureVmWorkloadProtectedItemExtendedInfoOutput {
	return i.ToAzureVmWorkloadProtectedItemExtendedInfoOutputWithContext(context.Background())
}

func (i AzureVmWorkloadProtectedItemExtendedInfoArgs) ToAzureVmWorkloadProtectedItemExtendedInfoOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectedItemExtendedInfoOutput)
}

func (i AzureVmWorkloadProtectedItemExtendedInfoArgs) ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutput() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return i.ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i AzureVmWorkloadProtectedItemExtendedInfoArgs) ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectedItemExtendedInfoOutput).ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(ctx)
}









type AzureVmWorkloadProtectedItemExtendedInfoPtrInput interface {
	pulumi.Input

	ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutput() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput
	ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(context.Context) AzureVmWorkloadProtectedItemExtendedInfoPtrOutput
}

type azureVmWorkloadProtectedItemExtendedInfoPtrType AzureVmWorkloadProtectedItemExtendedInfoArgs

func AzureVmWorkloadProtectedItemExtendedInfoPtr(v *AzureVmWorkloadProtectedItemExtendedInfoArgs) AzureVmWorkloadProtectedItemExtendedInfoPtrInput {
	return (*azureVmWorkloadProtectedItemExtendedInfoPtrType)(v)
}

func (*azureVmWorkloadProtectedItemExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureVmWorkloadProtectedItemExtendedInfo)(nil)).Elem()
}

func (i *azureVmWorkloadProtectedItemExtendedInfoPtrType) ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutput() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return i.ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *azureVmWorkloadProtectedItemExtendedInfoPtrType) ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectedItemExtendedInfoPtrOutput)
}

type AzureVmWorkloadProtectedItemExtendedInfoOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadProtectedItemExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectedItemExtendedInfo)(nil)).Elem()
}

func (o AzureVmWorkloadProtectedItemExtendedInfoOutput) ToAzureVmWorkloadProtectedItemExtendedInfoOutput() AzureVmWorkloadProtectedItemExtendedInfoOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemExtendedInfoOutput) ToAzureVmWorkloadProtectedItemExtendedInfoOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemExtendedInfoOutput) ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutput() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return o.ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (o AzureVmWorkloadProtectedItemExtendedInfoOutput) ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureVmWorkloadProtectedItemExtendedInfo) *AzureVmWorkloadProtectedItemExtendedInfo {
		return &v
	}).(AzureVmWorkloadProtectedItemExtendedInfoPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemExtendedInfo) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemExtendedInfo) *string { return v.PolicyState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemExtendedInfo) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type AzureVmWorkloadProtectedItemExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadProtectedItemExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureVmWorkloadProtectedItemExtendedInfo)(nil)).Elem()
}

func (o AzureVmWorkloadProtectedItemExtendedInfoPtrOutput) ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutput() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemExtendedInfoPtrOutput) ToAzureVmWorkloadProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemExtendedInfoPtrOutput) Elem() AzureVmWorkloadProtectedItemExtendedInfoOutput {
	return o.ApplyT(func(v *AzureVmWorkloadProtectedItemExtendedInfo) AzureVmWorkloadProtectedItemExtendedInfo {
		if v != nil {
			return *v
		}
		var ret AzureVmWorkloadProtectedItemExtendedInfo
		return ret
	}).(AzureVmWorkloadProtectedItemExtendedInfoOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoPtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureVmWorkloadProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoPtrOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureVmWorkloadProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.PolicyState
	}).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoPtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureVmWorkloadProtectedItemExtendedInfo) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type AzureVmWorkloadProtectedItemExtendedInfoResponse struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type AzureVmWorkloadProtectedItemExtendedInfoResponseInput interface {
	pulumi.Input

	ToAzureVmWorkloadProtectedItemExtendedInfoResponseOutput() AzureVmWorkloadProtectedItemExtendedInfoResponseOutput
	ToAzureVmWorkloadProtectedItemExtendedInfoResponseOutputWithContext(context.Context) AzureVmWorkloadProtectedItemExtendedInfoResponseOutput
}

type AzureVmWorkloadProtectedItemExtendedInfoResponseArgs struct {
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	PolicyState         pulumi.StringPtrInput `pulumi:"policyState"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (AzureVmWorkloadProtectedItemExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i AzureVmWorkloadProtectedItemExtendedInfoResponseArgs) ToAzureVmWorkloadProtectedItemExtendedInfoResponseOutput() AzureVmWorkloadProtectedItemExtendedInfoResponseOutput {
	return i.ToAzureVmWorkloadProtectedItemExtendedInfoResponseOutputWithContext(context.Background())
}

func (i AzureVmWorkloadProtectedItemExtendedInfoResponseArgs) ToAzureVmWorkloadProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectedItemExtendedInfoResponseOutput)
}

func (i AzureVmWorkloadProtectedItemExtendedInfoResponseArgs) ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i AzureVmWorkloadProtectedItemExtendedInfoResponseArgs) ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectedItemExtendedInfoResponseOutput).ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx)
}









type AzureVmWorkloadProtectedItemExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput
	ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Context) AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput
}

type azureVmWorkloadProtectedItemExtendedInfoResponsePtrType AzureVmWorkloadProtectedItemExtendedInfoResponseArgs

func AzureVmWorkloadProtectedItemExtendedInfoResponsePtr(v *AzureVmWorkloadProtectedItemExtendedInfoResponseArgs) AzureVmWorkloadProtectedItemExtendedInfoResponsePtrInput {
	return (*azureVmWorkloadProtectedItemExtendedInfoResponsePtrType)(v)
}

func (*azureVmWorkloadProtectedItemExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureVmWorkloadProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i *azureVmWorkloadProtectedItemExtendedInfoResponsePtrType) ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *azureVmWorkloadProtectedItemExtendedInfoResponsePtrType) ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput)
}

type AzureVmWorkloadProtectedItemExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadProtectedItemExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponseOutput) ToAzureVmWorkloadProtectedItemExtendedInfoResponseOutput() AzureVmWorkloadProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponseOutput) ToAzureVmWorkloadProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponseOutput) ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return o.ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponseOutput) ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureVmWorkloadProtectedItemExtendedInfoResponse) *AzureVmWorkloadProtectedItemExtendedInfoResponse {
		return &v
	}).(AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponseOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemExtendedInfoResponse) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponseOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemExtendedInfoResponse) *string { return v.PolicyState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponseOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemExtendedInfoResponse) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureVmWorkloadProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput) ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput) ToAzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput) Elem() AzureVmWorkloadProtectedItemExtendedInfoResponseOutput {
	return o.ApplyT(func(v *AzureVmWorkloadProtectedItemExtendedInfoResponse) AzureVmWorkloadProtectedItemExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureVmWorkloadProtectedItemExtendedInfoResponse
		return ret
	}).(AzureVmWorkloadProtectedItemExtendedInfoResponseOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureVmWorkloadProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureVmWorkloadProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyState
	}).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureVmWorkloadProtectedItemExtendedInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type AzureVmWorkloadProtectedItemResponse struct {
	BackupManagementType             *string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse       `pulumi:"kpisHealths"`
	LastBackupErrorDetail            *ErrorDetailResponse                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 *string                                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                           `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                           `pulumi:"parentName"`
	ParentType                       *string                                           `pulumi:"parentType"`
	PolicyId                         *string                                           `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 *string                                           `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                           `pulumi:"workloadType"`
}





type AzureVmWorkloadProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureVmWorkloadProtectedItemResponseOutput() AzureVmWorkloadProtectedItemResponseOutput
	ToAzureVmWorkloadProtectedItemResponseOutputWithContext(context.Context) AzureVmWorkloadProtectedItemResponseOutput
}

type AzureVmWorkloadProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                    `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                    `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                                    `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                    `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                    `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                    `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureVmWorkloadProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                                    `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                      `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                      `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                      `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsResponseMapInput                 `pulumi:"kpisHealths"`
	LastBackupErrorDetail            ErrorDetailResponsePtrInput                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 pulumi.StringPtrInput                                    `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                    `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                    `pulumi:"lastRecoveryPoint"`
	ParentName                       pulumi.StringPtrInput                                    `pulumi:"parentName"`
	ParentType                       pulumi.StringPtrInput                                    `pulumi:"parentType"`
	PolicyId                         pulumi.StringPtrInput                                    `pulumi:"policyId"`
	ProtectedItemDataSourceId        pulumi.StringPtrInput                                    `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        pulumi.StringPtrInput                                    `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                pulumi.StringInput                                       `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                    `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                                    `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       pulumi.StringPtrInput                                    `pulumi:"serverName"`
	SourceResourceId                 pulumi.StringPtrInput                                    `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                                    `pulumi:"workloadType"`
}

func (AzureVmWorkloadProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectedItemResponse)(nil)).Elem()
}

func (i AzureVmWorkloadProtectedItemResponseArgs) ToAzureVmWorkloadProtectedItemResponseOutput() AzureVmWorkloadProtectedItemResponseOutput {
	return i.ToAzureVmWorkloadProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureVmWorkloadProtectedItemResponseArgs) ToAzureVmWorkloadProtectedItemResponseOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectedItemResponseOutput)
}

type AzureVmWorkloadProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectedItemResponse)(nil)).Elem()
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ToAzureVmWorkloadProtectedItemResponseOutput() AzureVmWorkloadProtectedItemResponseOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ToAzureVmWorkloadProtectedItemResponseOutputWithContext(ctx context.Context) AzureVmWorkloadProtectedItemResponseOutput {
	return o
}

func (o AzureVmWorkloadProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ExtendedInfo() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *AzureVmWorkloadProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) KpisHealths() KPIResourceHealthDetailsResponseMapOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) map[string]KPIResourceHealthDetailsResponse {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsResponseMapOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) LastBackupErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *ErrorDetailResponse { return v.LastBackupErrorDetail }).(ErrorDetailResponsePtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ParentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.ParentType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ProtectedItemDataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.ProtectedItemDataSourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ProtectedItemHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.ProtectedItemHealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadProtectionPolicy struct {
	BackupManagementType           string                `pulumi:"backupManagementType"`
	MakePolicyConsistent           *bool                 `pulumi:"makePolicyConsistent"`
	ProtectedItemsCount            *int                  `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string              `pulumi:"resourceGuardOperationRequests"`
	Settings                       *Settings             `pulumi:"settings"`
	SubProtectionPolicy            []SubProtectionPolicy `pulumi:"subProtectionPolicy"`
	WorkLoadType                   *string               `pulumi:"workLoadType"`
}





type AzureVmWorkloadProtectionPolicyInput interface {
	pulumi.Input

	ToAzureVmWorkloadProtectionPolicyOutput() AzureVmWorkloadProtectionPolicyOutput
	ToAzureVmWorkloadProtectionPolicyOutputWithContext(context.Context) AzureVmWorkloadProtectionPolicyOutput
}

type AzureVmWorkloadProtectionPolicyArgs struct {
	BackupManagementType           pulumi.StringInput            `pulumi:"backupManagementType"`
	MakePolicyConsistent           pulumi.BoolPtrInput           `pulumi:"makePolicyConsistent"`
	ProtectedItemsCount            pulumi.IntPtrInput            `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput       `pulumi:"resourceGuardOperationRequests"`
	Settings                       SettingsPtrInput              `pulumi:"settings"`
	SubProtectionPolicy            SubProtectionPolicyArrayInput `pulumi:"subProtectionPolicy"`
	WorkLoadType                   pulumi.StringPtrInput         `pulumi:"workLoadType"`
}

func (AzureVmWorkloadProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectionPolicy)(nil)).Elem()
}

func (i AzureVmWorkloadProtectionPolicyArgs) ToAzureVmWorkloadProtectionPolicyOutput() AzureVmWorkloadProtectionPolicyOutput {
	return i.ToAzureVmWorkloadProtectionPolicyOutputWithContext(context.Background())
}

func (i AzureVmWorkloadProtectionPolicyArgs) ToAzureVmWorkloadProtectionPolicyOutputWithContext(ctx context.Context) AzureVmWorkloadProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectionPolicyOutput)
}

type AzureVmWorkloadProtectionPolicyOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectionPolicy)(nil)).Elem()
}

func (o AzureVmWorkloadProtectionPolicyOutput) ToAzureVmWorkloadProtectionPolicyOutput() AzureVmWorkloadProtectionPolicyOutput {
	return o
}

func (o AzureVmWorkloadProtectionPolicyOutput) ToAzureVmWorkloadProtectionPolicyOutputWithContext(ctx context.Context) AzureVmWorkloadProtectionPolicyOutput {
	return o
}

func (o AzureVmWorkloadProtectionPolicyOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicy) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadProtectionPolicyOutput) MakePolicyConsistent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicy) *bool { return v.MakePolicyConsistent }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureVmWorkloadProtectionPolicyOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicy) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadProtectionPolicyOutput) Settings() SettingsPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicy) *Settings { return v.Settings }).(SettingsPtrOutput)
}

func (o AzureVmWorkloadProtectionPolicyOutput) SubProtectionPolicy() SubProtectionPolicyArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicy) []SubProtectionPolicy { return v.SubProtectionPolicy }).(SubProtectionPolicyArrayOutput)
}

func (o AzureVmWorkloadProtectionPolicyOutput) WorkLoadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicy) *string { return v.WorkLoadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadProtectionPolicyResponse struct {
	BackupManagementType           string                        `pulumi:"backupManagementType"`
	MakePolicyConsistent           *bool                         `pulumi:"makePolicyConsistent"`
	ProtectedItemsCount            *int                          `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string                      `pulumi:"resourceGuardOperationRequests"`
	Settings                       *SettingsResponse             `pulumi:"settings"`
	SubProtectionPolicy            []SubProtectionPolicyResponse `pulumi:"subProtectionPolicy"`
	WorkLoadType                   *string                       `pulumi:"workLoadType"`
}





type AzureVmWorkloadProtectionPolicyResponseInput interface {
	pulumi.Input

	ToAzureVmWorkloadProtectionPolicyResponseOutput() AzureVmWorkloadProtectionPolicyResponseOutput
	ToAzureVmWorkloadProtectionPolicyResponseOutputWithContext(context.Context) AzureVmWorkloadProtectionPolicyResponseOutput
}

type AzureVmWorkloadProtectionPolicyResponseArgs struct {
	BackupManagementType           pulumi.StringInput                    `pulumi:"backupManagementType"`
	MakePolicyConsistent           pulumi.BoolPtrInput                   `pulumi:"makePolicyConsistent"`
	ProtectedItemsCount            pulumi.IntPtrInput                    `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput               `pulumi:"resourceGuardOperationRequests"`
	Settings                       SettingsResponsePtrInput              `pulumi:"settings"`
	SubProtectionPolicy            SubProtectionPolicyResponseArrayInput `pulumi:"subProtectionPolicy"`
	WorkLoadType                   pulumi.StringPtrInput                 `pulumi:"workLoadType"`
}

func (AzureVmWorkloadProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectionPolicyResponse)(nil)).Elem()
}

func (i AzureVmWorkloadProtectionPolicyResponseArgs) ToAzureVmWorkloadProtectionPolicyResponseOutput() AzureVmWorkloadProtectionPolicyResponseOutput {
	return i.ToAzureVmWorkloadProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i AzureVmWorkloadProtectionPolicyResponseArgs) ToAzureVmWorkloadProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureVmWorkloadProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadProtectionPolicyResponseOutput)
}

type AzureVmWorkloadProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadProtectionPolicyResponse)(nil)).Elem()
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) ToAzureVmWorkloadProtectionPolicyResponseOutput() AzureVmWorkloadProtectionPolicyResponseOutput {
	return o
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) ToAzureVmWorkloadProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureVmWorkloadProtectionPolicyResponseOutput {
	return o
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicyResponse) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) MakePolicyConsistent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicyResponse) *bool { return v.MakePolicyConsistent }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicyResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) Settings() SettingsResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicyResponse) *SettingsResponse { return v.Settings }).(SettingsResponsePtrOutput)
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) SubProtectionPolicy() SubProtectionPolicyResponseArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicyResponse) []SubProtectionPolicyResponse {
		return v.SubProtectionPolicy
	}).(SubProtectionPolicyResponseArrayOutput)
}

func (o AzureVmWorkloadProtectionPolicyResponseOutput) WorkLoadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadProtectionPolicyResponse) *string { return v.WorkLoadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadSAPAseDatabaseProtectedItem struct {
	BackupManagementType             *string                                   `pulumi:"backupManagementType"`
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails       `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                   `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                   `pulumi:"parentName"`
	ParentType                       *string                                   `pulumi:"parentType"`
	PolicyId                         *string                                   `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ProtectionStatus                 *string                                   `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                   `pulumi:"workloadType"`
}





type AzureVmWorkloadSAPAseDatabaseProtectedItemInput interface {
	pulumi.Input

	ToAzureVmWorkloadSAPAseDatabaseProtectedItemOutput() AzureVmWorkloadSAPAseDatabaseProtectedItemOutput
	ToAzureVmWorkloadSAPAseDatabaseProtectedItemOutputWithContext(context.Context) AzureVmWorkloadSAPAseDatabaseProtectedItemOutput
}

type AzureVmWorkloadSAPAseDatabaseProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                            `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                            `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                            `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                            `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                            `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                            `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureVmWorkloadProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                            `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                              `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                              `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                              `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsMapInput                 `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                            `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                            `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                            `pulumi:"lastRecoveryPoint"`
	ParentName                       pulumi.StringPtrInput                            `pulumi:"parentName"`
	ParentType                       pulumi.StringPtrInput                            `pulumi:"parentType"`
	PolicyId                         pulumi.StringPtrInput                            `pulumi:"policyId"`
	ProtectedItemDataSourceId        pulumi.StringPtrInput                            `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        pulumi.StringPtrInput                            `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                pulumi.StringInput                               `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                            `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       pulumi.StringPtrInput                            `pulumi:"serverName"`
	SourceResourceId                 pulumi.StringPtrInput                            `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                            `pulumi:"workloadType"`
}

func (AzureVmWorkloadSAPAseDatabaseProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSAPAseDatabaseProtectedItem)(nil)).Elem()
}

func (i AzureVmWorkloadSAPAseDatabaseProtectedItemArgs) ToAzureVmWorkloadSAPAseDatabaseProtectedItemOutput() AzureVmWorkloadSAPAseDatabaseProtectedItemOutput {
	return i.ToAzureVmWorkloadSAPAseDatabaseProtectedItemOutputWithContext(context.Background())
}

func (i AzureVmWorkloadSAPAseDatabaseProtectedItemArgs) ToAzureVmWorkloadSAPAseDatabaseProtectedItemOutputWithContext(ctx context.Context) AzureVmWorkloadSAPAseDatabaseProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadSAPAseDatabaseProtectedItemOutput)
}

type AzureVmWorkloadSAPAseDatabaseProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSAPAseDatabaseProtectedItem)(nil)).Elem()
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ToAzureVmWorkloadSAPAseDatabaseProtectedItemOutput() AzureVmWorkloadSAPAseDatabaseProtectedItemOutput {
	return o
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ToAzureVmWorkloadSAPAseDatabaseProtectedItemOutputWithContext(ctx context.Context) AzureVmWorkloadSAPAseDatabaseProtectedItemOutput {
	return o
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ExtendedInfo() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *AzureVmWorkloadProtectedItemExtendedInfo {
		return v.ExtendedInfo
	}).(AzureVmWorkloadProtectedItemExtendedInfoPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) KpisHealths() KPIResourceHealthDetailsMapOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) map[string]KPIResourceHealthDetails {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsMapOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ParentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.ParentType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ProtectedItemDataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.ProtectedItemDataSourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ProtectedItemHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.ProtectedItemHealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadSAPAseDatabaseProtectedItemResponse struct {
	BackupManagementType             *string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse       `pulumi:"kpisHealths"`
	LastBackupErrorDetail            *ErrorDetailResponse                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 *string                                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                           `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                           `pulumi:"parentName"`
	ParentType                       *string                                           `pulumi:"parentType"`
	PolicyId                         *string                                           `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 *string                                           `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                           `pulumi:"workloadType"`
}





type AzureVmWorkloadSAPAseDatabaseProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput() AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput
	ToAzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutputWithContext(context.Context) AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput
}

type AzureVmWorkloadSAPAseDatabaseProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                    `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                    `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                                    `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                    `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                    `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                    `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureVmWorkloadProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                                    `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                      `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                      `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                      `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsResponseMapInput                 `pulumi:"kpisHealths"`
	LastBackupErrorDetail            ErrorDetailResponsePtrInput                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 pulumi.StringPtrInput                                    `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                    `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                    `pulumi:"lastRecoveryPoint"`
	ParentName                       pulumi.StringPtrInput                                    `pulumi:"parentName"`
	ParentType                       pulumi.StringPtrInput                                    `pulumi:"parentType"`
	PolicyId                         pulumi.StringPtrInput                                    `pulumi:"policyId"`
	ProtectedItemDataSourceId        pulumi.StringPtrInput                                    `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        pulumi.StringPtrInput                                    `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                pulumi.StringInput                                       `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                    `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                                    `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       pulumi.StringPtrInput                                    `pulumi:"serverName"`
	SourceResourceId                 pulumi.StringPtrInput                                    `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                                    `pulumi:"workloadType"`
}

func (AzureVmWorkloadSAPAseDatabaseProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSAPAseDatabaseProtectedItemResponse)(nil)).Elem()
}

func (i AzureVmWorkloadSAPAseDatabaseProtectedItemResponseArgs) ToAzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput() AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput {
	return i.ToAzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureVmWorkloadSAPAseDatabaseProtectedItemResponseArgs) ToAzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutputWithContext(ctx context.Context) AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput)
}

type AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSAPAseDatabaseProtectedItemResponse)(nil)).Elem()
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ToAzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput() AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput {
	return o
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ToAzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutputWithContext(ctx context.Context) AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput {
	return o
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string {
		return v.DeferredDeleteTimeRemaining
	}).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ExtendedInfo() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *AzureVmWorkloadProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *bool {
		return v.IsDeferredDeleteScheduleUpcoming
	}).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *bool {
		return v.IsScheduledForDeferredDelete
	}).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) KpisHealths() KPIResourceHealthDetailsResponseMapOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) map[string]KPIResourceHealthDetailsResponse {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsResponseMapOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) LastBackupErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *ErrorDetailResponse {
		return v.LastBackupErrorDetail
	}).(ErrorDetailResponsePtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ParentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.ParentType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ProtectedItemDataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.ProtectedItemDataSourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ProtectedItemHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.ProtectedItemHealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) []string {
		return v.ResourceGuardOperationRequests
	}).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPAseDatabaseProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItem struct {
	BackupManagementType             *string                                   `pulumi:"backupManagementType"`
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails       `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                   `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                   `pulumi:"parentName"`
	ParentType                       *string                                   `pulumi:"parentType"`
	PolicyId                         *string                                   `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ProtectionStatus                 *string                                   `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                   `pulumi:"workloadType"`
}





type AzureVmWorkloadSAPHanaDatabaseProtectedItemInput interface {
	pulumi.Input

	ToAzureVmWorkloadSAPHanaDatabaseProtectedItemOutput() AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput
	ToAzureVmWorkloadSAPHanaDatabaseProtectedItemOutputWithContext(context.Context) AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                            `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                            `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                            `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                            `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                            `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                            `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureVmWorkloadProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                            `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                              `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                              `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                              `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsMapInput                 `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                            `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                            `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                            `pulumi:"lastRecoveryPoint"`
	ParentName                       pulumi.StringPtrInput                            `pulumi:"parentName"`
	ParentType                       pulumi.StringPtrInput                            `pulumi:"parentType"`
	PolicyId                         pulumi.StringPtrInput                            `pulumi:"policyId"`
	ProtectedItemDataSourceId        pulumi.StringPtrInput                            `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        pulumi.StringPtrInput                            `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                pulumi.StringInput                               `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                            `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       pulumi.StringPtrInput                            `pulumi:"serverName"`
	SourceResourceId                 pulumi.StringPtrInput                            `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                            `pulumi:"workloadType"`
}

func (AzureVmWorkloadSAPHanaDatabaseProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSAPHanaDatabaseProtectedItem)(nil)).Elem()
}

func (i AzureVmWorkloadSAPHanaDatabaseProtectedItemArgs) ToAzureVmWorkloadSAPHanaDatabaseProtectedItemOutput() AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput {
	return i.ToAzureVmWorkloadSAPHanaDatabaseProtectedItemOutputWithContext(context.Background())
}

func (i AzureVmWorkloadSAPHanaDatabaseProtectedItemArgs) ToAzureVmWorkloadSAPHanaDatabaseProtectedItemOutputWithContext(ctx context.Context) AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput)
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSAPHanaDatabaseProtectedItem)(nil)).Elem()
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ToAzureVmWorkloadSAPHanaDatabaseProtectedItemOutput() AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput {
	return o
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ToAzureVmWorkloadSAPHanaDatabaseProtectedItemOutputWithContext(ctx context.Context) AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput {
	return o
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ExtendedInfo() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *AzureVmWorkloadProtectedItemExtendedInfo {
		return v.ExtendedInfo
	}).(AzureVmWorkloadProtectedItemExtendedInfoPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) KpisHealths() KPIResourceHealthDetailsMapOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) map[string]KPIResourceHealthDetails {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsMapOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ParentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.ParentType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ProtectedItemDataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.ProtectedItemDataSourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ProtectedItemHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.ProtectedItemHealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse struct {
	BackupManagementType             *string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse       `pulumi:"kpisHealths"`
	LastBackupErrorDetail            *ErrorDetailResponse                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 *string                                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                           `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                           `pulumi:"parentName"`
	ParentType                       *string                                           `pulumi:"parentType"`
	PolicyId                         *string                                           `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 *string                                           `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                           `pulumi:"workloadType"`
}





type AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput() AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput
	ToAzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutputWithContext(context.Context) AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                    `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                    `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                                    `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                    `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                    `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                    `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureVmWorkloadProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                                    `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                      `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                      `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                      `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsResponseMapInput                 `pulumi:"kpisHealths"`
	LastBackupErrorDetail            ErrorDetailResponsePtrInput                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 pulumi.StringPtrInput                                    `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                    `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                    `pulumi:"lastRecoveryPoint"`
	ParentName                       pulumi.StringPtrInput                                    `pulumi:"parentName"`
	ParentType                       pulumi.StringPtrInput                                    `pulumi:"parentType"`
	PolicyId                         pulumi.StringPtrInput                                    `pulumi:"policyId"`
	ProtectedItemDataSourceId        pulumi.StringPtrInput                                    `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        pulumi.StringPtrInput                                    `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                pulumi.StringInput                                       `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                    `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                                    `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       pulumi.StringPtrInput                                    `pulumi:"serverName"`
	SourceResourceId                 pulumi.StringPtrInput                                    `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                                    `pulumi:"workloadType"`
}

func (AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse)(nil)).Elem()
}

func (i AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseArgs) ToAzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput() AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput {
	return i.ToAzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseArgs) ToAzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutputWithContext(ctx context.Context) AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput)
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse)(nil)).Elem()
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ToAzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput() AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput {
	return o
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ToAzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutputWithContext(ctx context.Context) AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput {
	return o
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string {
		return v.DeferredDeleteTimeRemaining
	}).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ExtendedInfo() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *AzureVmWorkloadProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *bool {
		return v.IsDeferredDeleteScheduleUpcoming
	}).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *bool {
		return v.IsScheduledForDeferredDelete
	}).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) KpisHealths() KPIResourceHealthDetailsResponseMapOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) map[string]KPIResourceHealthDetailsResponse {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsResponseMapOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) LastBackupErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *ErrorDetailResponse {
		return v.LastBackupErrorDetail
	}).(ErrorDetailResponsePtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ParentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.ParentType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ProtectedItemDataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string {
		return v.ProtectedItemDataSourceId
	}).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ProtectedItemHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string {
		return v.ProtectedItemHealthStatus
	}).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) []string {
		return v.ResourceGuardOperationRequests
	}).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadSQLDatabaseProtectedItem struct {
	BackupManagementType             *string                                   `pulumi:"backupManagementType"`
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails       `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                   `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                   `pulumi:"parentName"`
	ParentType                       *string                                   `pulumi:"parentType"`
	PolicyId                         *string                                   `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ProtectionStatus                 *string                                   `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                   `pulumi:"workloadType"`
}





type AzureVmWorkloadSQLDatabaseProtectedItemInput interface {
	pulumi.Input

	ToAzureVmWorkloadSQLDatabaseProtectedItemOutput() AzureVmWorkloadSQLDatabaseProtectedItemOutput
	ToAzureVmWorkloadSQLDatabaseProtectedItemOutputWithContext(context.Context) AzureVmWorkloadSQLDatabaseProtectedItemOutput
}

type AzureVmWorkloadSQLDatabaseProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                            `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                            `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                            `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                            `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                            `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                            `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureVmWorkloadProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                            `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                              `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                              `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                              `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsMapInput                 `pulumi:"kpisHealths"`
	LastBackupStatus                 pulumi.StringPtrInput                            `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                            `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                            `pulumi:"lastRecoveryPoint"`
	ParentName                       pulumi.StringPtrInput                            `pulumi:"parentName"`
	ParentType                       pulumi.StringPtrInput                            `pulumi:"parentType"`
	PolicyId                         pulumi.StringPtrInput                            `pulumi:"policyId"`
	ProtectedItemDataSourceId        pulumi.StringPtrInput                            `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        pulumi.StringPtrInput                            `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                pulumi.StringInput                               `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                            `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       pulumi.StringPtrInput                            `pulumi:"serverName"`
	SourceResourceId                 pulumi.StringPtrInput                            `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                            `pulumi:"workloadType"`
}

func (AzureVmWorkloadSQLDatabaseProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSQLDatabaseProtectedItem)(nil)).Elem()
}

func (i AzureVmWorkloadSQLDatabaseProtectedItemArgs) ToAzureVmWorkloadSQLDatabaseProtectedItemOutput() AzureVmWorkloadSQLDatabaseProtectedItemOutput {
	return i.ToAzureVmWorkloadSQLDatabaseProtectedItemOutputWithContext(context.Background())
}

func (i AzureVmWorkloadSQLDatabaseProtectedItemArgs) ToAzureVmWorkloadSQLDatabaseProtectedItemOutputWithContext(ctx context.Context) AzureVmWorkloadSQLDatabaseProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadSQLDatabaseProtectedItemOutput)
}

type AzureVmWorkloadSQLDatabaseProtectedItemOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadSQLDatabaseProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSQLDatabaseProtectedItem)(nil)).Elem()
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ToAzureVmWorkloadSQLDatabaseProtectedItemOutput() AzureVmWorkloadSQLDatabaseProtectedItemOutput {
	return o
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ToAzureVmWorkloadSQLDatabaseProtectedItemOutputWithContext(ctx context.Context) AzureVmWorkloadSQLDatabaseProtectedItemOutput {
	return o
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ExtendedInfo() AzureVmWorkloadProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *AzureVmWorkloadProtectedItemExtendedInfo {
		return v.ExtendedInfo
	}).(AzureVmWorkloadProtectedItemExtendedInfoPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) KpisHealths() KPIResourceHealthDetailsMapOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) map[string]KPIResourceHealthDetails {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsMapOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ParentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.ParentType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ProtectedItemDataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.ProtectedItemDataSourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ProtectedItemHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.ProtectedItemHealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureVmWorkloadSQLDatabaseProtectedItemResponse struct {
	BackupManagementType             *string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse       `pulumi:"kpisHealths"`
	LastBackupErrorDetail            *ErrorDetailResponse                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 *string                                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                           `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                           `pulumi:"parentName"`
	ParentType                       *string                                           `pulumi:"parentType"`
	PolicyId                         *string                                           `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 *string                                           `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                           `pulumi:"workloadType"`
}





type AzureVmWorkloadSQLDatabaseProtectedItemResponseInput interface {
	pulumi.Input

	ToAzureVmWorkloadSQLDatabaseProtectedItemResponseOutput() AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput
	ToAzureVmWorkloadSQLDatabaseProtectedItemResponseOutputWithContext(context.Context) AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput
}

type AzureVmWorkloadSQLDatabaseProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                    `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                    `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                                    `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                    `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                    `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                    `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     AzureVmWorkloadProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                                    `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                      `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                      `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                      `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      KPIResourceHealthDetailsResponseMapInput                 `pulumi:"kpisHealths"`
	LastBackupErrorDetail            ErrorDetailResponsePtrInput                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 pulumi.StringPtrInput                                    `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                    `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                    `pulumi:"lastRecoveryPoint"`
	ParentName                       pulumi.StringPtrInput                                    `pulumi:"parentName"`
	ParentType                       pulumi.StringPtrInput                                    `pulumi:"parentType"`
	PolicyId                         pulumi.StringPtrInput                                    `pulumi:"policyId"`
	ProtectedItemDataSourceId        pulumi.StringPtrInput                                    `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        pulumi.StringPtrInput                                    `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                pulumi.StringInput                                       `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                    `pulumi:"protectionState"`
	ProtectionStatus                 pulumi.StringPtrInput                                    `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       pulumi.StringPtrInput                                    `pulumi:"serverName"`
	SourceResourceId                 pulumi.StringPtrInput                                    `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                                    `pulumi:"workloadType"`
}

func (AzureVmWorkloadSQLDatabaseProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSQLDatabaseProtectedItemResponse)(nil)).Elem()
}

func (i AzureVmWorkloadSQLDatabaseProtectedItemResponseArgs) ToAzureVmWorkloadSQLDatabaseProtectedItemResponseOutput() AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput {
	return i.ToAzureVmWorkloadSQLDatabaseProtectedItemResponseOutputWithContext(context.Background())
}

func (i AzureVmWorkloadSQLDatabaseProtectedItemResponseArgs) ToAzureVmWorkloadSQLDatabaseProtectedItemResponseOutputWithContext(ctx context.Context) AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput)
}

type AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmWorkloadSQLDatabaseProtectedItemResponse)(nil)).Elem()
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ToAzureVmWorkloadSQLDatabaseProtectedItemResponseOutput() AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput {
	return o
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ToAzureVmWorkloadSQLDatabaseProtectedItemResponseOutputWithContext(ctx context.Context) AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput {
	return o
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ExtendedInfo() AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *AzureVmWorkloadProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *bool {
		return v.IsDeferredDeleteScheduleUpcoming
	}).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) KpisHealths() KPIResourceHealthDetailsResponseMapOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) map[string]KPIResourceHealthDetailsResponse {
		return v.KpisHealths
	}).(KPIResourceHealthDetailsResponseMapOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) LastBackupErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *ErrorDetailResponse {
		return v.LastBackupErrorDetail
	}).(ErrorDetailResponsePtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ParentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.ParentType }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ProtectedItemDataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.ProtectedItemDataSourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ProtectedItemHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.ProtectedItemHealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) []string {
		return v.ResourceGuardOperationRequests
	}).(pulumi.StringArrayOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmWorkloadSQLDatabaseProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureWorkloadAutoProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureWorkloadAutoProtectionIntentInput interface {
	pulumi.Input

	ToAzureWorkloadAutoProtectionIntentOutput() AzureWorkloadAutoProtectionIntentOutput
	ToAzureWorkloadAutoProtectionIntentOutputWithContext(context.Context) AzureWorkloadAutoProtectionIntentOutput
}

type AzureWorkloadAutoProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureWorkloadAutoProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadAutoProtectionIntent)(nil)).Elem()
}

func (i AzureWorkloadAutoProtectionIntentArgs) ToAzureWorkloadAutoProtectionIntentOutput() AzureWorkloadAutoProtectionIntentOutput {
	return i.ToAzureWorkloadAutoProtectionIntentOutputWithContext(context.Background())
}

func (i AzureWorkloadAutoProtectionIntentArgs) ToAzureWorkloadAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadAutoProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadAutoProtectionIntentOutput)
}

type AzureWorkloadAutoProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureWorkloadAutoProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadAutoProtectionIntent)(nil)).Elem()
}

func (o AzureWorkloadAutoProtectionIntentOutput) ToAzureWorkloadAutoProtectionIntentOutput() AzureWorkloadAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadAutoProtectionIntentOutput) ToAzureWorkloadAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadAutoProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureWorkloadAutoProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureWorkloadAutoProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureWorkloadAutoProtectionIntentResponseOutput() AzureWorkloadAutoProtectionIntentResponseOutput
	ToAzureWorkloadAutoProtectionIntentResponseOutputWithContext(context.Context) AzureWorkloadAutoProtectionIntentResponseOutput
}

type AzureWorkloadAutoProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureWorkloadAutoProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadAutoProtectionIntentResponse)(nil)).Elem()
}

func (i AzureWorkloadAutoProtectionIntentResponseArgs) ToAzureWorkloadAutoProtectionIntentResponseOutput() AzureWorkloadAutoProtectionIntentResponseOutput {
	return i.ToAzureWorkloadAutoProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureWorkloadAutoProtectionIntentResponseArgs) ToAzureWorkloadAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadAutoProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadAutoProtectionIntentResponseOutput)
}

type AzureWorkloadAutoProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureWorkloadAutoProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadAutoProtectionIntentResponse)(nil)).Elem()
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ToAzureWorkloadAutoProtectionIntentResponseOutput() AzureWorkloadAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ToAzureWorkloadAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureWorkloadContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        string                              `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}





type AzureWorkloadContainerInput interface {
	pulumi.Input

	ToAzureWorkloadContainerOutput() AzureWorkloadContainerOutput
	ToAzureWorkloadContainerOutputWithContext(context.Context) AzureWorkloadContainerOutput
}

type AzureWorkloadContainerArgs struct {
	BackupManagementType pulumi.StringPtrInput                      `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput                         `pulumi:"containerType"`
	ExtendedInfo         AzureWorkloadContainerExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput                      `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                      `pulumi:"healthStatus"`
	LastUpdatedTime      pulumi.StringPtrInput                      `pulumi:"lastUpdatedTime"`
	OperationType        pulumi.StringPtrInput                      `pulumi:"operationType"`
	RegistrationStatus   pulumi.StringPtrInput                      `pulumi:"registrationStatus"`
	SourceResourceId     pulumi.StringPtrInput                      `pulumi:"sourceResourceId"`
	WorkloadType         pulumi.StringPtrInput                      `pulumi:"workloadType"`
}

func (AzureWorkloadContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainer)(nil)).Elem()
}

func (i AzureWorkloadContainerArgs) ToAzureWorkloadContainerOutput() AzureWorkloadContainerOutput {
	return i.ToAzureWorkloadContainerOutputWithContext(context.Background())
}

func (i AzureWorkloadContainerArgs) ToAzureWorkloadContainerOutputWithContext(ctx context.Context) AzureWorkloadContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerOutput)
}

type AzureWorkloadContainerOutput struct{ *pulumi.OutputState }

func (AzureWorkloadContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainer)(nil)).Elem()
}

func (o AzureWorkloadContainerOutput) ToAzureWorkloadContainerOutput() AzureWorkloadContainerOutput {
	return o
}

func (o AzureWorkloadContainerOutput) ToAzureWorkloadContainerOutputWithContext(ctx context.Context) AzureWorkloadContainerOutput {
	return o
}

func (o AzureWorkloadContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureWorkloadContainerOutput) ExtendedInfo() AzureWorkloadContainerExtendedInfoPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *AzureWorkloadContainerExtendedInfo { return v.ExtendedInfo }).(AzureWorkloadContainerExtendedInfoPtrOutput)
}

func (o AzureWorkloadContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *string { return v.OperationType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainer) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureWorkloadContainerAutoProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureWorkloadContainerAutoProtectionIntentInput interface {
	pulumi.Input

	ToAzureWorkloadContainerAutoProtectionIntentOutput() AzureWorkloadContainerAutoProtectionIntentOutput
	ToAzureWorkloadContainerAutoProtectionIntentOutputWithContext(context.Context) AzureWorkloadContainerAutoProtectionIntentOutput
}

type AzureWorkloadContainerAutoProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureWorkloadContainerAutoProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerAutoProtectionIntent)(nil)).Elem()
}

func (i AzureWorkloadContainerAutoProtectionIntentArgs) ToAzureWorkloadContainerAutoProtectionIntentOutput() AzureWorkloadContainerAutoProtectionIntentOutput {
	return i.ToAzureWorkloadContainerAutoProtectionIntentOutputWithContext(context.Background())
}

func (i AzureWorkloadContainerAutoProtectionIntentArgs) ToAzureWorkloadContainerAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadContainerAutoProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerAutoProtectionIntentOutput)
}

type AzureWorkloadContainerAutoProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureWorkloadContainerAutoProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerAutoProtectionIntent)(nil)).Elem()
}

func (o AzureWorkloadContainerAutoProtectionIntentOutput) ToAzureWorkloadContainerAutoProtectionIntentOutput() AzureWorkloadContainerAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadContainerAutoProtectionIntentOutput) ToAzureWorkloadContainerAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadContainerAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadContainerAutoProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntent) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureWorkloadContainerAutoProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureWorkloadContainerAutoProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureWorkloadContainerAutoProtectionIntentResponseOutput() AzureWorkloadContainerAutoProtectionIntentResponseOutput
	ToAzureWorkloadContainerAutoProtectionIntentResponseOutputWithContext(context.Context) AzureWorkloadContainerAutoProtectionIntentResponseOutput
}

type AzureWorkloadContainerAutoProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureWorkloadContainerAutoProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerAutoProtectionIntentResponse)(nil)).Elem()
}

func (i AzureWorkloadContainerAutoProtectionIntentResponseArgs) ToAzureWorkloadContainerAutoProtectionIntentResponseOutput() AzureWorkloadContainerAutoProtectionIntentResponseOutput {
	return i.ToAzureWorkloadContainerAutoProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureWorkloadContainerAutoProtectionIntentResponseArgs) ToAzureWorkloadContainerAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadContainerAutoProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerAutoProtectionIntentResponseOutput)
}

type AzureWorkloadContainerAutoProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureWorkloadContainerAutoProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerAutoProtectionIntentResponse)(nil)).Elem()
}

func (o AzureWorkloadContainerAutoProtectionIntentResponseOutput) ToAzureWorkloadContainerAutoProtectionIntentResponseOutput() AzureWorkloadContainerAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadContainerAutoProtectionIntentResponseOutput) ToAzureWorkloadContainerAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadContainerAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadContainerAutoProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntentResponse) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerAutoProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerAutoProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureWorkloadContainerExtendedInfo struct {
	HostServerName *string                `pulumi:"hostServerName"`
	InquiryInfo    *InquiryInfo           `pulumi:"inquiryInfo"`
	NodesList      []DistributedNodesInfo `pulumi:"nodesList"`
}





type AzureWorkloadContainerExtendedInfoInput interface {
	pulumi.Input

	ToAzureWorkloadContainerExtendedInfoOutput() AzureWorkloadContainerExtendedInfoOutput
	ToAzureWorkloadContainerExtendedInfoOutputWithContext(context.Context) AzureWorkloadContainerExtendedInfoOutput
}

type AzureWorkloadContainerExtendedInfoArgs struct {
	HostServerName pulumi.StringPtrInput          `pulumi:"hostServerName"`
	InquiryInfo    InquiryInfoPtrInput            `pulumi:"inquiryInfo"`
	NodesList      DistributedNodesInfoArrayInput `pulumi:"nodesList"`
}

func (AzureWorkloadContainerExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerExtendedInfo)(nil)).Elem()
}

func (i AzureWorkloadContainerExtendedInfoArgs) ToAzureWorkloadContainerExtendedInfoOutput() AzureWorkloadContainerExtendedInfoOutput {
	return i.ToAzureWorkloadContainerExtendedInfoOutputWithContext(context.Background())
}

func (i AzureWorkloadContainerExtendedInfoArgs) ToAzureWorkloadContainerExtendedInfoOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerExtendedInfoOutput)
}

func (i AzureWorkloadContainerExtendedInfoArgs) ToAzureWorkloadContainerExtendedInfoPtrOutput() AzureWorkloadContainerExtendedInfoPtrOutput {
	return i.ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (i AzureWorkloadContainerExtendedInfoArgs) ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerExtendedInfoOutput).ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(ctx)
}









type AzureWorkloadContainerExtendedInfoPtrInput interface {
	pulumi.Input

	ToAzureWorkloadContainerExtendedInfoPtrOutput() AzureWorkloadContainerExtendedInfoPtrOutput
	ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(context.Context) AzureWorkloadContainerExtendedInfoPtrOutput
}

type azureWorkloadContainerExtendedInfoPtrType AzureWorkloadContainerExtendedInfoArgs

func AzureWorkloadContainerExtendedInfoPtr(v *AzureWorkloadContainerExtendedInfoArgs) AzureWorkloadContainerExtendedInfoPtrInput {
	return (*azureWorkloadContainerExtendedInfoPtrType)(v)
}

func (*azureWorkloadContainerExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureWorkloadContainerExtendedInfo)(nil)).Elem()
}

func (i *azureWorkloadContainerExtendedInfoPtrType) ToAzureWorkloadContainerExtendedInfoPtrOutput() AzureWorkloadContainerExtendedInfoPtrOutput {
	return i.ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *azureWorkloadContainerExtendedInfoPtrType) ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerExtendedInfoPtrOutput)
}

type AzureWorkloadContainerExtendedInfoOutput struct{ *pulumi.OutputState }

func (AzureWorkloadContainerExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerExtendedInfo)(nil)).Elem()
}

func (o AzureWorkloadContainerExtendedInfoOutput) ToAzureWorkloadContainerExtendedInfoOutput() AzureWorkloadContainerExtendedInfoOutput {
	return o
}

func (o AzureWorkloadContainerExtendedInfoOutput) ToAzureWorkloadContainerExtendedInfoOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoOutput {
	return o
}

func (o AzureWorkloadContainerExtendedInfoOutput) ToAzureWorkloadContainerExtendedInfoPtrOutput() AzureWorkloadContainerExtendedInfoPtrOutput {
	return o.ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (o AzureWorkloadContainerExtendedInfoOutput) ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureWorkloadContainerExtendedInfo) *AzureWorkloadContainerExtendedInfo {
		return &v
	}).(AzureWorkloadContainerExtendedInfoPtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoOutput) HostServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerExtendedInfo) *string { return v.HostServerName }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoOutput) InquiryInfo() InquiryInfoPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerExtendedInfo) *InquiryInfo { return v.InquiryInfo }).(InquiryInfoPtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoOutput) NodesList() DistributedNodesInfoArrayOutput {
	return o.ApplyT(func(v AzureWorkloadContainerExtendedInfo) []DistributedNodesInfo { return v.NodesList }).(DistributedNodesInfoArrayOutput)
}

type AzureWorkloadContainerExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (AzureWorkloadContainerExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureWorkloadContainerExtendedInfo)(nil)).Elem()
}

func (o AzureWorkloadContainerExtendedInfoPtrOutput) ToAzureWorkloadContainerExtendedInfoPtrOutput() AzureWorkloadContainerExtendedInfoPtrOutput {
	return o
}

func (o AzureWorkloadContainerExtendedInfoPtrOutput) ToAzureWorkloadContainerExtendedInfoPtrOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoPtrOutput {
	return o
}

func (o AzureWorkloadContainerExtendedInfoPtrOutput) Elem() AzureWorkloadContainerExtendedInfoOutput {
	return o.ApplyT(func(v *AzureWorkloadContainerExtendedInfo) AzureWorkloadContainerExtendedInfo {
		if v != nil {
			return *v
		}
		var ret AzureWorkloadContainerExtendedInfo
		return ret
	}).(AzureWorkloadContainerExtendedInfoOutput)
}

func (o AzureWorkloadContainerExtendedInfoPtrOutput) HostServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureWorkloadContainerExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.HostServerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoPtrOutput) InquiryInfo() InquiryInfoPtrOutput {
	return o.ApplyT(func(v *AzureWorkloadContainerExtendedInfo) *InquiryInfo {
		if v == nil {
			return nil
		}
		return v.InquiryInfo
	}).(InquiryInfoPtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoPtrOutput) NodesList() DistributedNodesInfoArrayOutput {
	return o.ApplyT(func(v *AzureWorkloadContainerExtendedInfo) []DistributedNodesInfo {
		if v == nil {
			return nil
		}
		return v.NodesList
	}).(DistributedNodesInfoArrayOutput)
}

type AzureWorkloadContainerExtendedInfoResponse struct {
	HostServerName *string                        `pulumi:"hostServerName"`
	InquiryInfo    *InquiryInfoResponse           `pulumi:"inquiryInfo"`
	NodesList      []DistributedNodesInfoResponse `pulumi:"nodesList"`
}





type AzureWorkloadContainerExtendedInfoResponseInput interface {
	pulumi.Input

	ToAzureWorkloadContainerExtendedInfoResponseOutput() AzureWorkloadContainerExtendedInfoResponseOutput
	ToAzureWorkloadContainerExtendedInfoResponseOutputWithContext(context.Context) AzureWorkloadContainerExtendedInfoResponseOutput
}

type AzureWorkloadContainerExtendedInfoResponseArgs struct {
	HostServerName pulumi.StringPtrInput                  `pulumi:"hostServerName"`
	InquiryInfo    InquiryInfoResponsePtrInput            `pulumi:"inquiryInfo"`
	NodesList      DistributedNodesInfoResponseArrayInput `pulumi:"nodesList"`
}

func (AzureWorkloadContainerExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerExtendedInfoResponse)(nil)).Elem()
}

func (i AzureWorkloadContainerExtendedInfoResponseArgs) ToAzureWorkloadContainerExtendedInfoResponseOutput() AzureWorkloadContainerExtendedInfoResponseOutput {
	return i.ToAzureWorkloadContainerExtendedInfoResponseOutputWithContext(context.Background())
}

func (i AzureWorkloadContainerExtendedInfoResponseArgs) ToAzureWorkloadContainerExtendedInfoResponseOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerExtendedInfoResponseOutput)
}

func (i AzureWorkloadContainerExtendedInfoResponseArgs) ToAzureWorkloadContainerExtendedInfoResponsePtrOutput() AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return i.ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i AzureWorkloadContainerExtendedInfoResponseArgs) ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerExtendedInfoResponseOutput).ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(ctx)
}









type AzureWorkloadContainerExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToAzureWorkloadContainerExtendedInfoResponsePtrOutput() AzureWorkloadContainerExtendedInfoResponsePtrOutput
	ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(context.Context) AzureWorkloadContainerExtendedInfoResponsePtrOutput
}

type azureWorkloadContainerExtendedInfoResponsePtrType AzureWorkloadContainerExtendedInfoResponseArgs

func AzureWorkloadContainerExtendedInfoResponsePtr(v *AzureWorkloadContainerExtendedInfoResponseArgs) AzureWorkloadContainerExtendedInfoResponsePtrInput {
	return (*azureWorkloadContainerExtendedInfoResponsePtrType)(v)
}

func (*azureWorkloadContainerExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureWorkloadContainerExtendedInfoResponse)(nil)).Elem()
}

func (i *azureWorkloadContainerExtendedInfoResponsePtrType) ToAzureWorkloadContainerExtendedInfoResponsePtrOutput() AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return i.ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *azureWorkloadContainerExtendedInfoResponsePtrType) ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerExtendedInfoResponsePtrOutput)
}

type AzureWorkloadContainerExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureWorkloadContainerExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerExtendedInfoResponse)(nil)).Elem()
}

func (o AzureWorkloadContainerExtendedInfoResponseOutput) ToAzureWorkloadContainerExtendedInfoResponseOutput() AzureWorkloadContainerExtendedInfoResponseOutput {
	return o
}

func (o AzureWorkloadContainerExtendedInfoResponseOutput) ToAzureWorkloadContainerExtendedInfoResponseOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoResponseOutput {
	return o
}

func (o AzureWorkloadContainerExtendedInfoResponseOutput) ToAzureWorkloadContainerExtendedInfoResponsePtrOutput() AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return o.ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o AzureWorkloadContainerExtendedInfoResponseOutput) ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureWorkloadContainerExtendedInfoResponse) *AzureWorkloadContainerExtendedInfoResponse {
		return &v
	}).(AzureWorkloadContainerExtendedInfoResponsePtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoResponseOutput) HostServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerExtendedInfoResponse) *string { return v.HostServerName }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoResponseOutput) InquiryInfo() InquiryInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerExtendedInfoResponse) *InquiryInfoResponse { return v.InquiryInfo }).(InquiryInfoResponsePtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoResponseOutput) NodesList() DistributedNodesInfoResponseArrayOutput {
	return o.ApplyT(func(v AzureWorkloadContainerExtendedInfoResponse) []DistributedNodesInfoResponse { return v.NodesList }).(DistributedNodesInfoResponseArrayOutput)
}

type AzureWorkloadContainerExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureWorkloadContainerExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureWorkloadContainerExtendedInfoResponse)(nil)).Elem()
}

func (o AzureWorkloadContainerExtendedInfoResponsePtrOutput) ToAzureWorkloadContainerExtendedInfoResponsePtrOutput() AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureWorkloadContainerExtendedInfoResponsePtrOutput) ToAzureWorkloadContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return o
}

func (o AzureWorkloadContainerExtendedInfoResponsePtrOutput) Elem() AzureWorkloadContainerExtendedInfoResponseOutput {
	return o.ApplyT(func(v *AzureWorkloadContainerExtendedInfoResponse) AzureWorkloadContainerExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureWorkloadContainerExtendedInfoResponse
		return ret
	}).(AzureWorkloadContainerExtendedInfoResponseOutput)
}

func (o AzureWorkloadContainerExtendedInfoResponsePtrOutput) HostServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureWorkloadContainerExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostServerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoResponsePtrOutput) InquiryInfo() InquiryInfoResponsePtrOutput {
	return o.ApplyT(func(v *AzureWorkloadContainerExtendedInfoResponse) *InquiryInfoResponse {
		if v == nil {
			return nil
		}
		return v.InquiryInfo
	}).(InquiryInfoResponsePtrOutput)
}

func (o AzureWorkloadContainerExtendedInfoResponsePtrOutput) NodesList() DistributedNodesInfoResponseArrayOutput {
	return o.ApplyT(func(v *AzureWorkloadContainerExtendedInfoResponse) []DistributedNodesInfoResponse {
		if v == nil {
			return nil
		}
		return v.NodesList
	}).(DistributedNodesInfoResponseArrayOutput)
}

type AzureWorkloadContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        string                                      `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}





type AzureWorkloadContainerResponseInput interface {
	pulumi.Input

	ToAzureWorkloadContainerResponseOutput() AzureWorkloadContainerResponseOutput
	ToAzureWorkloadContainerResponseOutputWithContext(context.Context) AzureWorkloadContainerResponseOutput
}

type AzureWorkloadContainerResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput                              `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput                                 `pulumi:"containerType"`
	ExtendedInfo         AzureWorkloadContainerExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput                              `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                              `pulumi:"healthStatus"`
	LastUpdatedTime      pulumi.StringPtrInput                              `pulumi:"lastUpdatedTime"`
	OperationType        pulumi.StringPtrInput                              `pulumi:"operationType"`
	RegistrationStatus   pulumi.StringPtrInput                              `pulumi:"registrationStatus"`
	SourceResourceId     pulumi.StringPtrInput                              `pulumi:"sourceResourceId"`
	WorkloadType         pulumi.StringPtrInput                              `pulumi:"workloadType"`
}

func (AzureWorkloadContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerResponse)(nil)).Elem()
}

func (i AzureWorkloadContainerResponseArgs) ToAzureWorkloadContainerResponseOutput() AzureWorkloadContainerResponseOutput {
	return i.ToAzureWorkloadContainerResponseOutputWithContext(context.Background())
}

func (i AzureWorkloadContainerResponseArgs) ToAzureWorkloadContainerResponseOutputWithContext(ctx context.Context) AzureWorkloadContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadContainerResponseOutput)
}

type AzureWorkloadContainerResponseOutput struct{ *pulumi.OutputState }

func (AzureWorkloadContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadContainerResponse)(nil)).Elem()
}

func (o AzureWorkloadContainerResponseOutput) ToAzureWorkloadContainerResponseOutput() AzureWorkloadContainerResponseOutput {
	return o
}

func (o AzureWorkloadContainerResponseOutput) ToAzureWorkloadContainerResponseOutputWithContext(ctx context.Context) AzureWorkloadContainerResponseOutput {
	return o
}

func (o AzureWorkloadContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o AzureWorkloadContainerResponseOutput) ExtendedInfo() AzureWorkloadContainerExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *AzureWorkloadContainerExtendedInfoResponse {
		return v.ExtendedInfo
	}).(AzureWorkloadContainerExtendedInfoResponsePtrOutput)
}

func (o AzureWorkloadContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerResponseOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerResponseOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *string { return v.OperationType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadContainerResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadContainerResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type AzureWorkloadSQLAutoProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
	WorkloadItemType         *string `pulumi:"workloadItemType"`
}





type AzureWorkloadSQLAutoProtectionIntentInput interface {
	pulumi.Input

	ToAzureWorkloadSQLAutoProtectionIntentOutput() AzureWorkloadSQLAutoProtectionIntentOutput
	ToAzureWorkloadSQLAutoProtectionIntentOutputWithContext(context.Context) AzureWorkloadSQLAutoProtectionIntentOutput
}

type AzureWorkloadSQLAutoProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
	WorkloadItemType         pulumi.StringPtrInput `pulumi:"workloadItemType"`
}

func (AzureWorkloadSQLAutoProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntent)(nil)).Elem()
}

func (i AzureWorkloadSQLAutoProtectionIntentArgs) ToAzureWorkloadSQLAutoProtectionIntentOutput() AzureWorkloadSQLAutoProtectionIntentOutput {
	return i.ToAzureWorkloadSQLAutoProtectionIntentOutputWithContext(context.Background())
}

func (i AzureWorkloadSQLAutoProtectionIntentArgs) ToAzureWorkloadSQLAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadSQLAutoProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadSQLAutoProtectionIntentOutput)
}

type AzureWorkloadSQLAutoProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureWorkloadSQLAutoProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntent)(nil)).Elem()
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ToAzureWorkloadSQLAutoProtectionIntentOutput() AzureWorkloadSQLAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ToAzureWorkloadSQLAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadSQLAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) WorkloadItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.WorkloadItemType }).(pulumi.StringPtrOutput)
}

type AzureWorkloadSQLAutoProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
	WorkloadItemType         *string `pulumi:"workloadItemType"`
}





type AzureWorkloadSQLAutoProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureWorkloadSQLAutoProtectionIntentResponseOutput() AzureWorkloadSQLAutoProtectionIntentResponseOutput
	ToAzureWorkloadSQLAutoProtectionIntentResponseOutputWithContext(context.Context) AzureWorkloadSQLAutoProtectionIntentResponseOutput
}

type AzureWorkloadSQLAutoProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringInput    `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
	WorkloadItemType         pulumi.StringPtrInput `pulumi:"workloadItemType"`
}

func (AzureWorkloadSQLAutoProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntentResponse)(nil)).Elem()
}

func (i AzureWorkloadSQLAutoProtectionIntentResponseArgs) ToAzureWorkloadSQLAutoProtectionIntentResponseOutput() AzureWorkloadSQLAutoProtectionIntentResponseOutput {
	return i.ToAzureWorkloadSQLAutoProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureWorkloadSQLAutoProtectionIntentResponseArgs) ToAzureWorkloadSQLAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadSQLAutoProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadSQLAutoProtectionIntentResponseOutput)
}

type AzureWorkloadSQLAutoProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureWorkloadSQLAutoProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntentResponse)(nil)).Elem()
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ToAzureWorkloadSQLAutoProtectionIntentResponseOutput() AzureWorkloadSQLAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ToAzureWorkloadSQLAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadSQLAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) string { return v.ProtectionIntentItemType }).(pulumi.StringOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) WorkloadItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.WorkloadItemType }).(pulumi.StringPtrOutput)
}

type ContainerIdentityInfo struct {
	AadTenantId              *string `pulumi:"aadTenantId"`
	Audience                 *string `pulumi:"audience"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	UniqueName               *string `pulumi:"uniqueName"`
}





type ContainerIdentityInfoInput interface {
	pulumi.Input

	ToContainerIdentityInfoOutput() ContainerIdentityInfoOutput
	ToContainerIdentityInfoOutputWithContext(context.Context) ContainerIdentityInfoOutput
}

type ContainerIdentityInfoArgs struct {
	AadTenantId              pulumi.StringPtrInput `pulumi:"aadTenantId"`
	Audience                 pulumi.StringPtrInput `pulumi:"audience"`
	ServicePrincipalClientId pulumi.StringPtrInput `pulumi:"servicePrincipalClientId"`
	UniqueName               pulumi.StringPtrInput `pulumi:"uniqueName"`
}

func (ContainerIdentityInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerIdentityInfo)(nil)).Elem()
}

func (i ContainerIdentityInfoArgs) ToContainerIdentityInfoOutput() ContainerIdentityInfoOutput {
	return i.ToContainerIdentityInfoOutputWithContext(context.Background())
}

func (i ContainerIdentityInfoArgs) ToContainerIdentityInfoOutputWithContext(ctx context.Context) ContainerIdentityInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerIdentityInfoOutput)
}

func (i ContainerIdentityInfoArgs) ToContainerIdentityInfoPtrOutput() ContainerIdentityInfoPtrOutput {
	return i.ToContainerIdentityInfoPtrOutputWithContext(context.Background())
}

func (i ContainerIdentityInfoArgs) ToContainerIdentityInfoPtrOutputWithContext(ctx context.Context) ContainerIdentityInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerIdentityInfoOutput).ToContainerIdentityInfoPtrOutputWithContext(ctx)
}









type ContainerIdentityInfoPtrInput interface {
	pulumi.Input

	ToContainerIdentityInfoPtrOutput() ContainerIdentityInfoPtrOutput
	ToContainerIdentityInfoPtrOutputWithContext(context.Context) ContainerIdentityInfoPtrOutput
}

type containerIdentityInfoPtrType ContainerIdentityInfoArgs

func ContainerIdentityInfoPtr(v *ContainerIdentityInfoArgs) ContainerIdentityInfoPtrInput {
	return (*containerIdentityInfoPtrType)(v)
}

func (*containerIdentityInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerIdentityInfo)(nil)).Elem()
}

func (i *containerIdentityInfoPtrType) ToContainerIdentityInfoPtrOutput() ContainerIdentityInfoPtrOutput {
	return i.ToContainerIdentityInfoPtrOutputWithContext(context.Background())
}

func (i *containerIdentityInfoPtrType) ToContainerIdentityInfoPtrOutputWithContext(ctx context.Context) ContainerIdentityInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerIdentityInfoPtrOutput)
}

type ContainerIdentityInfoOutput struct{ *pulumi.OutputState }

func (ContainerIdentityInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerIdentityInfo)(nil)).Elem()
}

func (o ContainerIdentityInfoOutput) ToContainerIdentityInfoOutput() ContainerIdentityInfoOutput {
	return o
}

func (o ContainerIdentityInfoOutput) ToContainerIdentityInfoOutputWithContext(ctx context.Context) ContainerIdentityInfoOutput {
	return o
}

func (o ContainerIdentityInfoOutput) ToContainerIdentityInfoPtrOutput() ContainerIdentityInfoPtrOutput {
	return o.ToContainerIdentityInfoPtrOutputWithContext(context.Background())
}

func (o ContainerIdentityInfoOutput) ToContainerIdentityInfoPtrOutputWithContext(ctx context.Context) ContainerIdentityInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerIdentityInfo) *ContainerIdentityInfo {
		return &v
	}).(ContainerIdentityInfoPtrOutput)
}

func (o ContainerIdentityInfoOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerIdentityInfo) *string { return v.AadTenantId }).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerIdentityInfo) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoOutput) ServicePrincipalClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerIdentityInfo) *string { return v.ServicePrincipalClientId }).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoOutput) UniqueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerIdentityInfo) *string { return v.UniqueName }).(pulumi.StringPtrOutput)
}

type ContainerIdentityInfoPtrOutput struct{ *pulumi.OutputState }

func (ContainerIdentityInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerIdentityInfo)(nil)).Elem()
}

func (o ContainerIdentityInfoPtrOutput) ToContainerIdentityInfoPtrOutput() ContainerIdentityInfoPtrOutput {
	return o
}

func (o ContainerIdentityInfoPtrOutput) ToContainerIdentityInfoPtrOutputWithContext(ctx context.Context) ContainerIdentityInfoPtrOutput {
	return o
}

func (o ContainerIdentityInfoPtrOutput) Elem() ContainerIdentityInfoOutput {
	return o.ApplyT(func(v *ContainerIdentityInfo) ContainerIdentityInfo {
		if v != nil {
			return *v
		}
		var ret ContainerIdentityInfo
		return ret
	}).(ContainerIdentityInfoOutput)
}

func (o ContainerIdentityInfoPtrOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerIdentityInfo) *string {
		if v == nil {
			return nil
		}
		return v.AadTenantId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerIdentityInfo) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoPtrOutput) ServicePrincipalClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerIdentityInfo) *string {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalClientId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoPtrOutput) UniqueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerIdentityInfo) *string {
		if v == nil {
			return nil
		}
		return v.UniqueName
	}).(pulumi.StringPtrOutput)
}

type ContainerIdentityInfoResponse struct {
	AadTenantId              *string `pulumi:"aadTenantId"`
	Audience                 *string `pulumi:"audience"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	UniqueName               *string `pulumi:"uniqueName"`
}





type ContainerIdentityInfoResponseInput interface {
	pulumi.Input

	ToContainerIdentityInfoResponseOutput() ContainerIdentityInfoResponseOutput
	ToContainerIdentityInfoResponseOutputWithContext(context.Context) ContainerIdentityInfoResponseOutput
}

type ContainerIdentityInfoResponseArgs struct {
	AadTenantId              pulumi.StringPtrInput `pulumi:"aadTenantId"`
	Audience                 pulumi.StringPtrInput `pulumi:"audience"`
	ServicePrincipalClientId pulumi.StringPtrInput `pulumi:"servicePrincipalClientId"`
	UniqueName               pulumi.StringPtrInput `pulumi:"uniqueName"`
}

func (ContainerIdentityInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerIdentityInfoResponse)(nil)).Elem()
}

func (i ContainerIdentityInfoResponseArgs) ToContainerIdentityInfoResponseOutput() ContainerIdentityInfoResponseOutput {
	return i.ToContainerIdentityInfoResponseOutputWithContext(context.Background())
}

func (i ContainerIdentityInfoResponseArgs) ToContainerIdentityInfoResponseOutputWithContext(ctx context.Context) ContainerIdentityInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerIdentityInfoResponseOutput)
}

func (i ContainerIdentityInfoResponseArgs) ToContainerIdentityInfoResponsePtrOutput() ContainerIdentityInfoResponsePtrOutput {
	return i.ToContainerIdentityInfoResponsePtrOutputWithContext(context.Background())
}

func (i ContainerIdentityInfoResponseArgs) ToContainerIdentityInfoResponsePtrOutputWithContext(ctx context.Context) ContainerIdentityInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerIdentityInfoResponseOutput).ToContainerIdentityInfoResponsePtrOutputWithContext(ctx)
}









type ContainerIdentityInfoResponsePtrInput interface {
	pulumi.Input

	ToContainerIdentityInfoResponsePtrOutput() ContainerIdentityInfoResponsePtrOutput
	ToContainerIdentityInfoResponsePtrOutputWithContext(context.Context) ContainerIdentityInfoResponsePtrOutput
}

type containerIdentityInfoResponsePtrType ContainerIdentityInfoResponseArgs

func ContainerIdentityInfoResponsePtr(v *ContainerIdentityInfoResponseArgs) ContainerIdentityInfoResponsePtrInput {
	return (*containerIdentityInfoResponsePtrType)(v)
}

func (*containerIdentityInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerIdentityInfoResponse)(nil)).Elem()
}

func (i *containerIdentityInfoResponsePtrType) ToContainerIdentityInfoResponsePtrOutput() ContainerIdentityInfoResponsePtrOutput {
	return i.ToContainerIdentityInfoResponsePtrOutputWithContext(context.Background())
}

func (i *containerIdentityInfoResponsePtrType) ToContainerIdentityInfoResponsePtrOutputWithContext(ctx context.Context) ContainerIdentityInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerIdentityInfoResponsePtrOutput)
}

type ContainerIdentityInfoResponseOutput struct{ *pulumi.OutputState }

func (ContainerIdentityInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerIdentityInfoResponse)(nil)).Elem()
}

func (o ContainerIdentityInfoResponseOutput) ToContainerIdentityInfoResponseOutput() ContainerIdentityInfoResponseOutput {
	return o
}

func (o ContainerIdentityInfoResponseOutput) ToContainerIdentityInfoResponseOutputWithContext(ctx context.Context) ContainerIdentityInfoResponseOutput {
	return o
}

func (o ContainerIdentityInfoResponseOutput) ToContainerIdentityInfoResponsePtrOutput() ContainerIdentityInfoResponsePtrOutput {
	return o.ToContainerIdentityInfoResponsePtrOutputWithContext(context.Background())
}

func (o ContainerIdentityInfoResponseOutput) ToContainerIdentityInfoResponsePtrOutputWithContext(ctx context.Context) ContainerIdentityInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerIdentityInfoResponse) *ContainerIdentityInfoResponse {
		return &v
	}).(ContainerIdentityInfoResponsePtrOutput)
}

func (o ContainerIdentityInfoResponseOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerIdentityInfoResponse) *string { return v.AadTenantId }).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoResponseOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerIdentityInfoResponse) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoResponseOutput) ServicePrincipalClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerIdentityInfoResponse) *string { return v.ServicePrincipalClientId }).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoResponseOutput) UniqueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerIdentityInfoResponse) *string { return v.UniqueName }).(pulumi.StringPtrOutput)
}

type ContainerIdentityInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerIdentityInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerIdentityInfoResponse)(nil)).Elem()
}

func (o ContainerIdentityInfoResponsePtrOutput) ToContainerIdentityInfoResponsePtrOutput() ContainerIdentityInfoResponsePtrOutput {
	return o
}

func (o ContainerIdentityInfoResponsePtrOutput) ToContainerIdentityInfoResponsePtrOutputWithContext(ctx context.Context) ContainerIdentityInfoResponsePtrOutput {
	return o
}

func (o ContainerIdentityInfoResponsePtrOutput) Elem() ContainerIdentityInfoResponseOutput {
	return o.ApplyT(func(v *ContainerIdentityInfoResponse) ContainerIdentityInfoResponse {
		if v != nil {
			return *v
		}
		var ret ContainerIdentityInfoResponse
		return ret
	}).(ContainerIdentityInfoResponseOutput)
}

func (o ContainerIdentityInfoResponsePtrOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerIdentityInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadTenantId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoResponsePtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerIdentityInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoResponsePtrOutput) ServicePrincipalClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerIdentityInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalClientId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerIdentityInfoResponsePtrOutput) UniqueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerIdentityInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UniqueName
	}).(pulumi.StringPtrOutput)
}

type DPMContainerExtendedInfo struct {
	LastRefreshedAt *string `pulumi:"lastRefreshedAt"`
}





type DPMContainerExtendedInfoInput interface {
	pulumi.Input

	ToDPMContainerExtendedInfoOutput() DPMContainerExtendedInfoOutput
	ToDPMContainerExtendedInfoOutputWithContext(context.Context) DPMContainerExtendedInfoOutput
}

type DPMContainerExtendedInfoArgs struct {
	LastRefreshedAt pulumi.StringPtrInput `pulumi:"lastRefreshedAt"`
}

func (DPMContainerExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMContainerExtendedInfo)(nil)).Elem()
}

func (i DPMContainerExtendedInfoArgs) ToDPMContainerExtendedInfoOutput() DPMContainerExtendedInfoOutput {
	return i.ToDPMContainerExtendedInfoOutputWithContext(context.Background())
}

func (i DPMContainerExtendedInfoArgs) ToDPMContainerExtendedInfoOutputWithContext(ctx context.Context) DPMContainerExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMContainerExtendedInfoOutput)
}

func (i DPMContainerExtendedInfoArgs) ToDPMContainerExtendedInfoPtrOutput() DPMContainerExtendedInfoPtrOutput {
	return i.ToDPMContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (i DPMContainerExtendedInfoArgs) ToDPMContainerExtendedInfoPtrOutputWithContext(ctx context.Context) DPMContainerExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMContainerExtendedInfoOutput).ToDPMContainerExtendedInfoPtrOutputWithContext(ctx)
}









type DPMContainerExtendedInfoPtrInput interface {
	pulumi.Input

	ToDPMContainerExtendedInfoPtrOutput() DPMContainerExtendedInfoPtrOutput
	ToDPMContainerExtendedInfoPtrOutputWithContext(context.Context) DPMContainerExtendedInfoPtrOutput
}

type dpmcontainerExtendedInfoPtrType DPMContainerExtendedInfoArgs

func DPMContainerExtendedInfoPtr(v *DPMContainerExtendedInfoArgs) DPMContainerExtendedInfoPtrInput {
	return (*dpmcontainerExtendedInfoPtrType)(v)
}

func (*dpmcontainerExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DPMContainerExtendedInfo)(nil)).Elem()
}

func (i *dpmcontainerExtendedInfoPtrType) ToDPMContainerExtendedInfoPtrOutput() DPMContainerExtendedInfoPtrOutput {
	return i.ToDPMContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *dpmcontainerExtendedInfoPtrType) ToDPMContainerExtendedInfoPtrOutputWithContext(ctx context.Context) DPMContainerExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMContainerExtendedInfoPtrOutput)
}

type DPMContainerExtendedInfoOutput struct{ *pulumi.OutputState }

func (DPMContainerExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMContainerExtendedInfo)(nil)).Elem()
}

func (o DPMContainerExtendedInfoOutput) ToDPMContainerExtendedInfoOutput() DPMContainerExtendedInfoOutput {
	return o
}

func (o DPMContainerExtendedInfoOutput) ToDPMContainerExtendedInfoOutputWithContext(ctx context.Context) DPMContainerExtendedInfoOutput {
	return o
}

func (o DPMContainerExtendedInfoOutput) ToDPMContainerExtendedInfoPtrOutput() DPMContainerExtendedInfoPtrOutput {
	return o.ToDPMContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (o DPMContainerExtendedInfoOutput) ToDPMContainerExtendedInfoPtrOutputWithContext(ctx context.Context) DPMContainerExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DPMContainerExtendedInfo) *DPMContainerExtendedInfo {
		return &v
	}).(DPMContainerExtendedInfoPtrOutput)
}

func (o DPMContainerExtendedInfoOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMContainerExtendedInfo) *string { return v.LastRefreshedAt }).(pulumi.StringPtrOutput)
}

type DPMContainerExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (DPMContainerExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DPMContainerExtendedInfo)(nil)).Elem()
}

func (o DPMContainerExtendedInfoPtrOutput) ToDPMContainerExtendedInfoPtrOutput() DPMContainerExtendedInfoPtrOutput {
	return o
}

func (o DPMContainerExtendedInfoPtrOutput) ToDPMContainerExtendedInfoPtrOutputWithContext(ctx context.Context) DPMContainerExtendedInfoPtrOutput {
	return o
}

func (o DPMContainerExtendedInfoPtrOutput) Elem() DPMContainerExtendedInfoOutput {
	return o.ApplyT(func(v *DPMContainerExtendedInfo) DPMContainerExtendedInfo {
		if v != nil {
			return *v
		}
		var ret DPMContainerExtendedInfo
		return ret
	}).(DPMContainerExtendedInfoOutput)
}

func (o DPMContainerExtendedInfoPtrOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMContainerExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.LastRefreshedAt
	}).(pulumi.StringPtrOutput)
}

type DPMContainerExtendedInfoResponse struct {
	LastRefreshedAt *string `pulumi:"lastRefreshedAt"`
}





type DPMContainerExtendedInfoResponseInput interface {
	pulumi.Input

	ToDPMContainerExtendedInfoResponseOutput() DPMContainerExtendedInfoResponseOutput
	ToDPMContainerExtendedInfoResponseOutputWithContext(context.Context) DPMContainerExtendedInfoResponseOutput
}

type DPMContainerExtendedInfoResponseArgs struct {
	LastRefreshedAt pulumi.StringPtrInput `pulumi:"lastRefreshedAt"`
}

func (DPMContainerExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMContainerExtendedInfoResponse)(nil)).Elem()
}

func (i DPMContainerExtendedInfoResponseArgs) ToDPMContainerExtendedInfoResponseOutput() DPMContainerExtendedInfoResponseOutput {
	return i.ToDPMContainerExtendedInfoResponseOutputWithContext(context.Background())
}

func (i DPMContainerExtendedInfoResponseArgs) ToDPMContainerExtendedInfoResponseOutputWithContext(ctx context.Context) DPMContainerExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMContainerExtendedInfoResponseOutput)
}

func (i DPMContainerExtendedInfoResponseArgs) ToDPMContainerExtendedInfoResponsePtrOutput() DPMContainerExtendedInfoResponsePtrOutput {
	return i.ToDPMContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i DPMContainerExtendedInfoResponseArgs) ToDPMContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) DPMContainerExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMContainerExtendedInfoResponseOutput).ToDPMContainerExtendedInfoResponsePtrOutputWithContext(ctx)
}









type DPMContainerExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToDPMContainerExtendedInfoResponsePtrOutput() DPMContainerExtendedInfoResponsePtrOutput
	ToDPMContainerExtendedInfoResponsePtrOutputWithContext(context.Context) DPMContainerExtendedInfoResponsePtrOutput
}

type dpmcontainerExtendedInfoResponsePtrType DPMContainerExtendedInfoResponseArgs

func DPMContainerExtendedInfoResponsePtr(v *DPMContainerExtendedInfoResponseArgs) DPMContainerExtendedInfoResponsePtrInput {
	return (*dpmcontainerExtendedInfoResponsePtrType)(v)
}

func (*dpmcontainerExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DPMContainerExtendedInfoResponse)(nil)).Elem()
}

func (i *dpmcontainerExtendedInfoResponsePtrType) ToDPMContainerExtendedInfoResponsePtrOutput() DPMContainerExtendedInfoResponsePtrOutput {
	return i.ToDPMContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *dpmcontainerExtendedInfoResponsePtrType) ToDPMContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) DPMContainerExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMContainerExtendedInfoResponsePtrOutput)
}

type DPMContainerExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (DPMContainerExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMContainerExtendedInfoResponse)(nil)).Elem()
}

func (o DPMContainerExtendedInfoResponseOutput) ToDPMContainerExtendedInfoResponseOutput() DPMContainerExtendedInfoResponseOutput {
	return o
}

func (o DPMContainerExtendedInfoResponseOutput) ToDPMContainerExtendedInfoResponseOutputWithContext(ctx context.Context) DPMContainerExtendedInfoResponseOutput {
	return o
}

func (o DPMContainerExtendedInfoResponseOutput) ToDPMContainerExtendedInfoResponsePtrOutput() DPMContainerExtendedInfoResponsePtrOutput {
	return o.ToDPMContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o DPMContainerExtendedInfoResponseOutput) ToDPMContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) DPMContainerExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DPMContainerExtendedInfoResponse) *DPMContainerExtendedInfoResponse {
		return &v
	}).(DPMContainerExtendedInfoResponsePtrOutput)
}

func (o DPMContainerExtendedInfoResponseOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMContainerExtendedInfoResponse) *string { return v.LastRefreshedAt }).(pulumi.StringPtrOutput)
}

type DPMContainerExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (DPMContainerExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DPMContainerExtendedInfoResponse)(nil)).Elem()
}

func (o DPMContainerExtendedInfoResponsePtrOutput) ToDPMContainerExtendedInfoResponsePtrOutput() DPMContainerExtendedInfoResponsePtrOutput {
	return o
}

func (o DPMContainerExtendedInfoResponsePtrOutput) ToDPMContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) DPMContainerExtendedInfoResponsePtrOutput {
	return o
}

func (o DPMContainerExtendedInfoResponsePtrOutput) Elem() DPMContainerExtendedInfoResponseOutput {
	return o.ApplyT(func(v *DPMContainerExtendedInfoResponse) DPMContainerExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret DPMContainerExtendedInfoResponse
		return ret
	}).(DPMContainerExtendedInfoResponseOutput)
}

func (o DPMContainerExtendedInfoResponsePtrOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMContainerExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastRefreshedAt
	}).(pulumi.StringPtrOutput)
}

type DPMProtectedItem struct {
	BackupEngineName                 *string                       `pulumi:"backupEngineName"`
	BackupManagementType             *string                       `pulumi:"backupManagementType"`
	BackupSetName                    *string                       `pulumi:"backupSetName"`
	ContainerName                    *string                       `pulumi:"containerName"`
	CreateMode                       *string                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *DPMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                       `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                         `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                       `pulumi:"policyId"`
	ProtectedItemType                string                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                       `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                      `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                       `pulumi:"sourceResourceId"`
	WorkloadType                     *string                       `pulumi:"workloadType"`
}





type DPMProtectedItemInput interface {
	pulumi.Input

	ToDPMProtectedItemOutput() DPMProtectedItemOutput
	ToDPMProtectedItemOutputWithContext(context.Context) DPMProtectedItemOutput
}

type DPMProtectedItemArgs struct {
	BackupEngineName                 pulumi.StringPtrInput                `pulumi:"backupEngineName"`
	BackupManagementType             pulumi.StringPtrInput                `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     DPMProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                  `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                  `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                  `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                pulumi.StringPtrInput                `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                `pulumi:"policyId"`
	ProtectedItemType                pulumi.StringInput                   `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                `pulumi:"protectionState"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput              `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                `pulumi:"workloadType"`
}

func (DPMProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMProtectedItem)(nil)).Elem()
}

func (i DPMProtectedItemArgs) ToDPMProtectedItemOutput() DPMProtectedItemOutput {
	return i.ToDPMProtectedItemOutputWithContext(context.Background())
}

func (i DPMProtectedItemArgs) ToDPMProtectedItemOutputWithContext(ctx context.Context) DPMProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMProtectedItemOutput)
}

type DPMProtectedItemOutput struct{ *pulumi.OutputState }

func (DPMProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMProtectedItem)(nil)).Elem()
}

func (o DPMProtectedItemOutput) ToDPMProtectedItemOutput() DPMProtectedItemOutput {
	return o
}

func (o DPMProtectedItemOutput) ToDPMProtectedItemOutputWithContext(ctx context.Context) DPMProtectedItemOutput {
	return o
}

func (o DPMProtectedItemOutput) BackupEngineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.BackupEngineName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) ExtendedInfo() DPMProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *DPMProtectedItemExtendedInfo { return v.ExtendedInfo }).(DPMProtectedItemExtendedInfoPtrOutput)
}

func (o DPMProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v DPMProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o DPMProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DPMProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o DPMProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type DPMProtectedItemExtendedInfo struct {
	DiskStorageUsedInBytes       *string           `pulumi:"diskStorageUsedInBytes"`
	IsCollocated                 *bool             `pulumi:"isCollocated"`
	IsPresentOnCloud             *bool             `pulumi:"isPresentOnCloud"`
	LastBackupStatus             *string           `pulumi:"lastBackupStatus"`
	LastRefreshedAt              *string           `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint          *string           `pulumi:"oldestRecoveryPoint"`
	OnPremiseLatestRecoveryPoint *string           `pulumi:"onPremiseLatestRecoveryPoint"`
	OnPremiseOldestRecoveryPoint *string           `pulumi:"onPremiseOldestRecoveryPoint"`
	OnPremiseRecoveryPointCount  *int              `pulumi:"onPremiseRecoveryPointCount"`
	ProtectableObjectLoadPath    map[string]string `pulumi:"protectableObjectLoadPath"`
	Protected                    *bool             `pulumi:"protected"`
	ProtectionGroupName          *string           `pulumi:"protectionGroupName"`
	RecoveryPointCount           *int              `pulumi:"recoveryPointCount"`
	TotalDiskStorageSizeInBytes  *string           `pulumi:"totalDiskStorageSizeInBytes"`
}





type DPMProtectedItemExtendedInfoInput interface {
	pulumi.Input

	ToDPMProtectedItemExtendedInfoOutput() DPMProtectedItemExtendedInfoOutput
	ToDPMProtectedItemExtendedInfoOutputWithContext(context.Context) DPMProtectedItemExtendedInfoOutput
}

type DPMProtectedItemExtendedInfoArgs struct {
	DiskStorageUsedInBytes       pulumi.StringPtrInput `pulumi:"diskStorageUsedInBytes"`
	IsCollocated                 pulumi.BoolPtrInput   `pulumi:"isCollocated"`
	IsPresentOnCloud             pulumi.BoolPtrInput   `pulumi:"isPresentOnCloud"`
	LastBackupStatus             pulumi.StringPtrInput `pulumi:"lastBackupStatus"`
	LastRefreshedAt              pulumi.StringPtrInput `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint          pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	OnPremiseLatestRecoveryPoint pulumi.StringPtrInput `pulumi:"onPremiseLatestRecoveryPoint"`
	OnPremiseOldestRecoveryPoint pulumi.StringPtrInput `pulumi:"onPremiseOldestRecoveryPoint"`
	OnPremiseRecoveryPointCount  pulumi.IntPtrInput    `pulumi:"onPremiseRecoveryPointCount"`
	ProtectableObjectLoadPath    pulumi.StringMapInput `pulumi:"protectableObjectLoadPath"`
	Protected                    pulumi.BoolPtrInput   `pulumi:"protected"`
	ProtectionGroupName          pulumi.StringPtrInput `pulumi:"protectionGroupName"`
	RecoveryPointCount           pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
	TotalDiskStorageSizeInBytes  pulumi.StringPtrInput `pulumi:"totalDiskStorageSizeInBytes"`
}

func (DPMProtectedItemExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMProtectedItemExtendedInfo)(nil)).Elem()
}

func (i DPMProtectedItemExtendedInfoArgs) ToDPMProtectedItemExtendedInfoOutput() DPMProtectedItemExtendedInfoOutput {
	return i.ToDPMProtectedItemExtendedInfoOutputWithContext(context.Background())
}

func (i DPMProtectedItemExtendedInfoArgs) ToDPMProtectedItemExtendedInfoOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMProtectedItemExtendedInfoOutput)
}

func (i DPMProtectedItemExtendedInfoArgs) ToDPMProtectedItemExtendedInfoPtrOutput() DPMProtectedItemExtendedInfoPtrOutput {
	return i.ToDPMProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i DPMProtectedItemExtendedInfoArgs) ToDPMProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMProtectedItemExtendedInfoOutput).ToDPMProtectedItemExtendedInfoPtrOutputWithContext(ctx)
}









type DPMProtectedItemExtendedInfoPtrInput interface {
	pulumi.Input

	ToDPMProtectedItemExtendedInfoPtrOutput() DPMProtectedItemExtendedInfoPtrOutput
	ToDPMProtectedItemExtendedInfoPtrOutputWithContext(context.Context) DPMProtectedItemExtendedInfoPtrOutput
}

type dpmprotectedItemExtendedInfoPtrType DPMProtectedItemExtendedInfoArgs

func DPMProtectedItemExtendedInfoPtr(v *DPMProtectedItemExtendedInfoArgs) DPMProtectedItemExtendedInfoPtrInput {
	return (*dpmprotectedItemExtendedInfoPtrType)(v)
}

func (*dpmprotectedItemExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DPMProtectedItemExtendedInfo)(nil)).Elem()
}

func (i *dpmprotectedItemExtendedInfoPtrType) ToDPMProtectedItemExtendedInfoPtrOutput() DPMProtectedItemExtendedInfoPtrOutput {
	return i.ToDPMProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *dpmprotectedItemExtendedInfoPtrType) ToDPMProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMProtectedItemExtendedInfoPtrOutput)
}

type DPMProtectedItemExtendedInfoOutput struct{ *pulumi.OutputState }

func (DPMProtectedItemExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMProtectedItemExtendedInfo)(nil)).Elem()
}

func (o DPMProtectedItemExtendedInfoOutput) ToDPMProtectedItemExtendedInfoOutput() DPMProtectedItemExtendedInfoOutput {
	return o
}

func (o DPMProtectedItemExtendedInfoOutput) ToDPMProtectedItemExtendedInfoOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoOutput {
	return o
}

func (o DPMProtectedItemExtendedInfoOutput) ToDPMProtectedItemExtendedInfoPtrOutput() DPMProtectedItemExtendedInfoPtrOutput {
	return o.ToDPMProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (o DPMProtectedItemExtendedInfoOutput) ToDPMProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DPMProtectedItemExtendedInfo) *DPMProtectedItemExtendedInfo {
		return &v
	}).(DPMProtectedItemExtendedInfoPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) DiskStorageUsedInBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *string { return v.DiskStorageUsedInBytes }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) IsCollocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *bool { return v.IsCollocated }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) IsPresentOnCloud() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *bool { return v.IsPresentOnCloud }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *string { return v.LastRefreshedAt }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) OnPremiseLatestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *string { return v.OnPremiseLatestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) OnPremiseOldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *string { return v.OnPremiseOldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) OnPremiseRecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *int { return v.OnPremiseRecoveryPointCount }).(pulumi.IntPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) ProtectableObjectLoadPath() pulumi.StringMapOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) map[string]string { return v.ProtectableObjectLoadPath }).(pulumi.StringMapOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) Protected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *bool { return v.Protected }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) ProtectionGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *string { return v.ProtectionGroupName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

func (o DPMProtectedItemExtendedInfoOutput) TotalDiskStorageSizeInBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfo) *string { return v.TotalDiskStorageSizeInBytes }).(pulumi.StringPtrOutput)
}

type DPMProtectedItemExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (DPMProtectedItemExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DPMProtectedItemExtendedInfo)(nil)).Elem()
}

func (o DPMProtectedItemExtendedInfoPtrOutput) ToDPMProtectedItemExtendedInfoPtrOutput() DPMProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o DPMProtectedItemExtendedInfoPtrOutput) ToDPMProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o DPMProtectedItemExtendedInfoPtrOutput) Elem() DPMProtectedItemExtendedInfoOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) DPMProtectedItemExtendedInfo {
		if v != nil {
			return *v
		}
		var ret DPMProtectedItemExtendedInfo
		return ret
	}).(DPMProtectedItemExtendedInfoOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) DiskStorageUsedInBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.DiskStorageUsedInBytes
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) IsCollocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *bool {
		if v == nil {
			return nil
		}
		return v.IsCollocated
	}).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) IsPresentOnCloud() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *bool {
		if v == nil {
			return nil
		}
		return v.IsPresentOnCloud
	}).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.LastBackupStatus
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.LastRefreshedAt
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) OnPremiseLatestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.OnPremiseLatestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) OnPremiseOldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.OnPremiseOldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) OnPremiseRecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *int {
		if v == nil {
			return nil
		}
		return v.OnPremiseRecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) ProtectableObjectLoadPath() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) map[string]string {
		if v == nil {
			return nil
		}
		return v.ProtectableObjectLoadPath
	}).(pulumi.StringMapOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) Protected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *bool {
		if v == nil {
			return nil
		}
		return v.Protected
	}).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) ProtectionGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.ProtectionGroupName
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

func (o DPMProtectedItemExtendedInfoPtrOutput) TotalDiskStorageSizeInBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.TotalDiskStorageSizeInBytes
	}).(pulumi.StringPtrOutput)
}

type DPMProtectedItemExtendedInfoResponse struct {
	DiskStorageUsedInBytes       *string           `pulumi:"diskStorageUsedInBytes"`
	IsCollocated                 *bool             `pulumi:"isCollocated"`
	IsPresentOnCloud             *bool             `pulumi:"isPresentOnCloud"`
	LastBackupStatus             *string           `pulumi:"lastBackupStatus"`
	LastRefreshedAt              *string           `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint          *string           `pulumi:"oldestRecoveryPoint"`
	OnPremiseLatestRecoveryPoint *string           `pulumi:"onPremiseLatestRecoveryPoint"`
	OnPremiseOldestRecoveryPoint *string           `pulumi:"onPremiseOldestRecoveryPoint"`
	OnPremiseRecoveryPointCount  *int              `pulumi:"onPremiseRecoveryPointCount"`
	ProtectableObjectLoadPath    map[string]string `pulumi:"protectableObjectLoadPath"`
	Protected                    *bool             `pulumi:"protected"`
	ProtectionGroupName          *string           `pulumi:"protectionGroupName"`
	RecoveryPointCount           *int              `pulumi:"recoveryPointCount"`
	TotalDiskStorageSizeInBytes  *string           `pulumi:"totalDiskStorageSizeInBytes"`
}





type DPMProtectedItemExtendedInfoResponseInput interface {
	pulumi.Input

	ToDPMProtectedItemExtendedInfoResponseOutput() DPMProtectedItemExtendedInfoResponseOutput
	ToDPMProtectedItemExtendedInfoResponseOutputWithContext(context.Context) DPMProtectedItemExtendedInfoResponseOutput
}

type DPMProtectedItemExtendedInfoResponseArgs struct {
	DiskStorageUsedInBytes       pulumi.StringPtrInput `pulumi:"diskStorageUsedInBytes"`
	IsCollocated                 pulumi.BoolPtrInput   `pulumi:"isCollocated"`
	IsPresentOnCloud             pulumi.BoolPtrInput   `pulumi:"isPresentOnCloud"`
	LastBackupStatus             pulumi.StringPtrInput `pulumi:"lastBackupStatus"`
	LastRefreshedAt              pulumi.StringPtrInput `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint          pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	OnPremiseLatestRecoveryPoint pulumi.StringPtrInput `pulumi:"onPremiseLatestRecoveryPoint"`
	OnPremiseOldestRecoveryPoint pulumi.StringPtrInput `pulumi:"onPremiseOldestRecoveryPoint"`
	OnPremiseRecoveryPointCount  pulumi.IntPtrInput    `pulumi:"onPremiseRecoveryPointCount"`
	ProtectableObjectLoadPath    pulumi.StringMapInput `pulumi:"protectableObjectLoadPath"`
	Protected                    pulumi.BoolPtrInput   `pulumi:"protected"`
	ProtectionGroupName          pulumi.StringPtrInput `pulumi:"protectionGroupName"`
	RecoveryPointCount           pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
	TotalDiskStorageSizeInBytes  pulumi.StringPtrInput `pulumi:"totalDiskStorageSizeInBytes"`
}

func (DPMProtectedItemExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i DPMProtectedItemExtendedInfoResponseArgs) ToDPMProtectedItemExtendedInfoResponseOutput() DPMProtectedItemExtendedInfoResponseOutput {
	return i.ToDPMProtectedItemExtendedInfoResponseOutputWithContext(context.Background())
}

func (i DPMProtectedItemExtendedInfoResponseArgs) ToDPMProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMProtectedItemExtendedInfoResponseOutput)
}

func (i DPMProtectedItemExtendedInfoResponseArgs) ToDPMProtectedItemExtendedInfoResponsePtrOutput() DPMProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i DPMProtectedItemExtendedInfoResponseArgs) ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMProtectedItemExtendedInfoResponseOutput).ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx)
}









type DPMProtectedItemExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToDPMProtectedItemExtendedInfoResponsePtrOutput() DPMProtectedItemExtendedInfoResponsePtrOutput
	ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Context) DPMProtectedItemExtendedInfoResponsePtrOutput
}

type dpmprotectedItemExtendedInfoResponsePtrType DPMProtectedItemExtendedInfoResponseArgs

func DPMProtectedItemExtendedInfoResponsePtr(v *DPMProtectedItemExtendedInfoResponseArgs) DPMProtectedItemExtendedInfoResponsePtrInput {
	return (*dpmprotectedItemExtendedInfoResponsePtrType)(v)
}

func (*dpmprotectedItemExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DPMProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i *dpmprotectedItemExtendedInfoResponsePtrType) ToDPMProtectedItemExtendedInfoResponsePtrOutput() DPMProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *dpmprotectedItemExtendedInfoResponsePtrType) ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMProtectedItemExtendedInfoResponsePtrOutput)
}

type DPMProtectedItemExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (DPMProtectedItemExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o DPMProtectedItemExtendedInfoResponseOutput) ToDPMProtectedItemExtendedInfoResponseOutput() DPMProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o DPMProtectedItemExtendedInfoResponseOutput) ToDPMProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o DPMProtectedItemExtendedInfoResponseOutput) ToDPMProtectedItemExtendedInfoResponsePtrOutput() DPMProtectedItemExtendedInfoResponsePtrOutput {
	return o.ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o DPMProtectedItemExtendedInfoResponseOutput) ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DPMProtectedItemExtendedInfoResponse) *DPMProtectedItemExtendedInfoResponse {
		return &v
	}).(DPMProtectedItemExtendedInfoResponsePtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) DiskStorageUsedInBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *string { return v.DiskStorageUsedInBytes }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) IsCollocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *bool { return v.IsCollocated }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) IsPresentOnCloud() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *bool { return v.IsPresentOnCloud }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *string { return v.LastRefreshedAt }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) OnPremiseLatestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *string { return v.OnPremiseLatestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) OnPremiseOldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *string { return v.OnPremiseOldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) OnPremiseRecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *int { return v.OnPremiseRecoveryPointCount }).(pulumi.IntPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) ProtectableObjectLoadPath() pulumi.StringMapOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) map[string]string { return v.ProtectableObjectLoadPath }).(pulumi.StringMapOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) Protected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *bool { return v.Protected }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) ProtectionGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *string { return v.ProtectionGroupName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponseOutput) TotalDiskStorageSizeInBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemExtendedInfoResponse) *string { return v.TotalDiskStorageSizeInBytes }).(pulumi.StringPtrOutput)
}

type DPMProtectedItemExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (DPMProtectedItemExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DPMProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) ToDPMProtectedItemExtendedInfoResponsePtrOutput() DPMProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) ToDPMProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) DPMProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) Elem() DPMProtectedItemExtendedInfoResponseOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) DPMProtectedItemExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret DPMProtectedItemExtendedInfoResponse
		return ret
	}).(DPMProtectedItemExtendedInfoResponseOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) DiskStorageUsedInBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiskStorageUsedInBytes
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) IsCollocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCollocated
	}).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) IsPresentOnCloud() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsPresentOnCloud
	}).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastBackupStatus
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastRefreshedAt
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) OnPremiseLatestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OnPremiseLatestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) OnPremiseOldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OnPremiseOldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) OnPremiseRecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.OnPremiseRecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) ProtectableObjectLoadPath() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.ProtectableObjectLoadPath
	}).(pulumi.StringMapOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) Protected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Protected
	}).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) ProtectionGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProtectionGroupName
	}).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

func (o DPMProtectedItemExtendedInfoResponsePtrOutput) TotalDiskStorageSizeInBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DPMProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.TotalDiskStorageSizeInBytes
	}).(pulumi.StringPtrOutput)
}

type DPMProtectedItemResponse struct {
	BackupEngineName                 *string                               `pulumi:"backupEngineName"`
	BackupManagementType             *string                               `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *DPMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                              `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	WorkloadType                     *string                               `pulumi:"workloadType"`
}





type DPMProtectedItemResponseInput interface {
	pulumi.Input

	ToDPMProtectedItemResponseOutput() DPMProtectedItemResponseOutput
	ToDPMProtectedItemResponseOutputWithContext(context.Context) DPMProtectedItemResponseOutput
}

type DPMProtectedItemResponseArgs struct {
	BackupEngineName                 pulumi.StringPtrInput                        `pulumi:"backupEngineName"`
	BackupManagementType             pulumi.StringPtrInput                        `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                        `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput                        `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                        `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                        `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                        `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     DPMProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                        `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                          `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                          `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                          `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                pulumi.StringPtrInput                        `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                        `pulumi:"policyId"`
	ProtectedItemType                pulumi.StringInput                           `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                        `pulumi:"protectionState"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                      `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                        `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                        `pulumi:"workloadType"`
}

func (DPMProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMProtectedItemResponse)(nil)).Elem()
}

func (i DPMProtectedItemResponseArgs) ToDPMProtectedItemResponseOutput() DPMProtectedItemResponseOutput {
	return i.ToDPMProtectedItemResponseOutputWithContext(context.Background())
}

func (i DPMProtectedItemResponseArgs) ToDPMProtectedItemResponseOutputWithContext(ctx context.Context) DPMProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DPMProtectedItemResponseOutput)
}

type DPMProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (DPMProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DPMProtectedItemResponse)(nil)).Elem()
}

func (o DPMProtectedItemResponseOutput) ToDPMProtectedItemResponseOutput() DPMProtectedItemResponseOutput {
	return o
}

func (o DPMProtectedItemResponseOutput) ToDPMProtectedItemResponseOutputWithContext(ctx context.Context) DPMProtectedItemResponseOutput {
	return o
}

func (o DPMProtectedItemResponseOutput) BackupEngineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.BackupEngineName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) ExtendedInfo() DPMProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *DPMProtectedItemExtendedInfoResponse { return v.ExtendedInfo }).(DPMProtectedItemExtendedInfoResponsePtrOutput)
}

func (o DPMProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o DPMProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o DPMProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o DPMProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o DPMProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DPMProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type DailyRetentionFormat struct {
	DaysOfTheMonth []Day `pulumi:"daysOfTheMonth"`
}





type DailyRetentionFormatInput interface {
	pulumi.Input

	ToDailyRetentionFormatOutput() DailyRetentionFormatOutput
	ToDailyRetentionFormatOutputWithContext(context.Context) DailyRetentionFormatOutput
}

type DailyRetentionFormatArgs struct {
	DaysOfTheMonth DayArrayInput `pulumi:"daysOfTheMonth"`
}

func (DailyRetentionFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionFormat)(nil)).Elem()
}

func (i DailyRetentionFormatArgs) ToDailyRetentionFormatOutput() DailyRetentionFormatOutput {
	return i.ToDailyRetentionFormatOutputWithContext(context.Background())
}

func (i DailyRetentionFormatArgs) ToDailyRetentionFormatOutputWithContext(ctx context.Context) DailyRetentionFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatOutput)
}

func (i DailyRetentionFormatArgs) ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput {
	return i.ToDailyRetentionFormatPtrOutputWithContext(context.Background())
}

func (i DailyRetentionFormatArgs) ToDailyRetentionFormatPtrOutputWithContext(ctx context.Context) DailyRetentionFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatOutput).ToDailyRetentionFormatPtrOutputWithContext(ctx)
}









type DailyRetentionFormatPtrInput interface {
	pulumi.Input

	ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput
	ToDailyRetentionFormatPtrOutputWithContext(context.Context) DailyRetentionFormatPtrOutput
}

type dailyRetentionFormatPtrType DailyRetentionFormatArgs

func DailyRetentionFormatPtr(v *DailyRetentionFormatArgs) DailyRetentionFormatPtrInput {
	return (*dailyRetentionFormatPtrType)(v)
}

func (*dailyRetentionFormatPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionFormat)(nil)).Elem()
}

func (i *dailyRetentionFormatPtrType) ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput {
	return i.ToDailyRetentionFormatPtrOutputWithContext(context.Background())
}

func (i *dailyRetentionFormatPtrType) ToDailyRetentionFormatPtrOutputWithContext(ctx context.Context) DailyRetentionFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatPtrOutput)
}

type DailyRetentionFormatOutput struct{ *pulumi.OutputState }

func (DailyRetentionFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionFormat)(nil)).Elem()
}

func (o DailyRetentionFormatOutput) ToDailyRetentionFormatOutput() DailyRetentionFormatOutput {
	return o
}

func (o DailyRetentionFormatOutput) ToDailyRetentionFormatOutputWithContext(ctx context.Context) DailyRetentionFormatOutput {
	return o
}

func (o DailyRetentionFormatOutput) ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput {
	return o.ToDailyRetentionFormatPtrOutputWithContext(context.Background())
}

func (o DailyRetentionFormatOutput) ToDailyRetentionFormatPtrOutputWithContext(ctx context.Context) DailyRetentionFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyRetentionFormat) *DailyRetentionFormat {
		return &v
	}).(DailyRetentionFormatPtrOutput)
}

func (o DailyRetentionFormatOutput) DaysOfTheMonth() DayArrayOutput {
	return o.ApplyT(func(v DailyRetentionFormat) []Day { return v.DaysOfTheMonth }).(DayArrayOutput)
}

type DailyRetentionFormatPtrOutput struct{ *pulumi.OutputState }

func (DailyRetentionFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionFormat)(nil)).Elem()
}

func (o DailyRetentionFormatPtrOutput) ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput {
	return o
}

func (o DailyRetentionFormatPtrOutput) ToDailyRetentionFormatPtrOutputWithContext(ctx context.Context) DailyRetentionFormatPtrOutput {
	return o
}

func (o DailyRetentionFormatPtrOutput) Elem() DailyRetentionFormatOutput {
	return o.ApplyT(func(v *DailyRetentionFormat) DailyRetentionFormat {
		if v != nil {
			return *v
		}
		var ret DailyRetentionFormat
		return ret
	}).(DailyRetentionFormatOutput)
}

func (o DailyRetentionFormatPtrOutput) DaysOfTheMonth() DayArrayOutput {
	return o.ApplyT(func(v *DailyRetentionFormat) []Day {
		if v == nil {
			return nil
		}
		return v.DaysOfTheMonth
	}).(DayArrayOutput)
}

type DailyRetentionFormatResponse struct {
	DaysOfTheMonth []DayResponse `pulumi:"daysOfTheMonth"`
}





type DailyRetentionFormatResponseInput interface {
	pulumi.Input

	ToDailyRetentionFormatResponseOutput() DailyRetentionFormatResponseOutput
	ToDailyRetentionFormatResponseOutputWithContext(context.Context) DailyRetentionFormatResponseOutput
}

type DailyRetentionFormatResponseArgs struct {
	DaysOfTheMonth DayResponseArrayInput `pulumi:"daysOfTheMonth"`
}

func (DailyRetentionFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionFormatResponse)(nil)).Elem()
}

func (i DailyRetentionFormatResponseArgs) ToDailyRetentionFormatResponseOutput() DailyRetentionFormatResponseOutput {
	return i.ToDailyRetentionFormatResponseOutputWithContext(context.Background())
}

func (i DailyRetentionFormatResponseArgs) ToDailyRetentionFormatResponseOutputWithContext(ctx context.Context) DailyRetentionFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatResponseOutput)
}

func (i DailyRetentionFormatResponseArgs) ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput {
	return i.ToDailyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (i DailyRetentionFormatResponseArgs) ToDailyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) DailyRetentionFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatResponseOutput).ToDailyRetentionFormatResponsePtrOutputWithContext(ctx)
}









type DailyRetentionFormatResponsePtrInput interface {
	pulumi.Input

	ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput
	ToDailyRetentionFormatResponsePtrOutputWithContext(context.Context) DailyRetentionFormatResponsePtrOutput
}

type dailyRetentionFormatResponsePtrType DailyRetentionFormatResponseArgs

func DailyRetentionFormatResponsePtr(v *DailyRetentionFormatResponseArgs) DailyRetentionFormatResponsePtrInput {
	return (*dailyRetentionFormatResponsePtrType)(v)
}

func (*dailyRetentionFormatResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionFormatResponse)(nil)).Elem()
}

func (i *dailyRetentionFormatResponsePtrType) ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput {
	return i.ToDailyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (i *dailyRetentionFormatResponsePtrType) ToDailyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) DailyRetentionFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatResponsePtrOutput)
}

type DailyRetentionFormatResponseOutput struct{ *pulumi.OutputState }

func (DailyRetentionFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionFormatResponse)(nil)).Elem()
}

func (o DailyRetentionFormatResponseOutput) ToDailyRetentionFormatResponseOutput() DailyRetentionFormatResponseOutput {
	return o
}

func (o DailyRetentionFormatResponseOutput) ToDailyRetentionFormatResponseOutputWithContext(ctx context.Context) DailyRetentionFormatResponseOutput {
	return o
}

func (o DailyRetentionFormatResponseOutput) ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput {
	return o.ToDailyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (o DailyRetentionFormatResponseOutput) ToDailyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) DailyRetentionFormatResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyRetentionFormatResponse) *DailyRetentionFormatResponse {
		return &v
	}).(DailyRetentionFormatResponsePtrOutput)
}

func (o DailyRetentionFormatResponseOutput) DaysOfTheMonth() DayResponseArrayOutput {
	return o.ApplyT(func(v DailyRetentionFormatResponse) []DayResponse { return v.DaysOfTheMonth }).(DayResponseArrayOutput)
}

type DailyRetentionFormatResponsePtrOutput struct{ *pulumi.OutputState }

func (DailyRetentionFormatResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionFormatResponse)(nil)).Elem()
}

func (o DailyRetentionFormatResponsePtrOutput) ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput {
	return o
}

func (o DailyRetentionFormatResponsePtrOutput) ToDailyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) DailyRetentionFormatResponsePtrOutput {
	return o
}

func (o DailyRetentionFormatResponsePtrOutput) Elem() DailyRetentionFormatResponseOutput {
	return o.ApplyT(func(v *DailyRetentionFormatResponse) DailyRetentionFormatResponse {
		if v != nil {
			return *v
		}
		var ret DailyRetentionFormatResponse
		return ret
	}).(DailyRetentionFormatResponseOutput)
}

func (o DailyRetentionFormatResponsePtrOutput) DaysOfTheMonth() DayResponseArrayOutput {
	return o.ApplyT(func(v *DailyRetentionFormatResponse) []DayResponse {
		if v == nil {
			return nil
		}
		return v.DaysOfTheMonth
	}).(DayResponseArrayOutput)
}

type DailyRetentionSchedule struct {
	RetentionDuration *RetentionDuration `pulumi:"retentionDuration"`
	RetentionTimes    []string           `pulumi:"retentionTimes"`
}





type DailyRetentionScheduleInput interface {
	pulumi.Input

	ToDailyRetentionScheduleOutput() DailyRetentionScheduleOutput
	ToDailyRetentionScheduleOutputWithContext(context.Context) DailyRetentionScheduleOutput
}

type DailyRetentionScheduleArgs struct {
	RetentionDuration RetentionDurationPtrInput `pulumi:"retentionDuration"`
	RetentionTimes    pulumi.StringArrayInput   `pulumi:"retentionTimes"`
}

func (DailyRetentionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionSchedule)(nil)).Elem()
}

func (i DailyRetentionScheduleArgs) ToDailyRetentionScheduleOutput() DailyRetentionScheduleOutput {
	return i.ToDailyRetentionScheduleOutputWithContext(context.Background())
}

func (i DailyRetentionScheduleArgs) ToDailyRetentionScheduleOutputWithContext(ctx context.Context) DailyRetentionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleOutput)
}

func (i DailyRetentionScheduleArgs) ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput {
	return i.ToDailyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i DailyRetentionScheduleArgs) ToDailyRetentionSchedulePtrOutputWithContext(ctx context.Context) DailyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleOutput).ToDailyRetentionSchedulePtrOutputWithContext(ctx)
}









type DailyRetentionSchedulePtrInput interface {
	pulumi.Input

	ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput
	ToDailyRetentionSchedulePtrOutputWithContext(context.Context) DailyRetentionSchedulePtrOutput
}

type dailyRetentionSchedulePtrType DailyRetentionScheduleArgs

func DailyRetentionSchedulePtr(v *DailyRetentionScheduleArgs) DailyRetentionSchedulePtrInput {
	return (*dailyRetentionSchedulePtrType)(v)
}

func (*dailyRetentionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionSchedule)(nil)).Elem()
}

func (i *dailyRetentionSchedulePtrType) ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput {
	return i.ToDailyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i *dailyRetentionSchedulePtrType) ToDailyRetentionSchedulePtrOutputWithContext(ctx context.Context) DailyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionSchedulePtrOutput)
}

type DailyRetentionScheduleOutput struct{ *pulumi.OutputState }

func (DailyRetentionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionSchedule)(nil)).Elem()
}

func (o DailyRetentionScheduleOutput) ToDailyRetentionScheduleOutput() DailyRetentionScheduleOutput {
	return o
}

func (o DailyRetentionScheduleOutput) ToDailyRetentionScheduleOutputWithContext(ctx context.Context) DailyRetentionScheduleOutput {
	return o
}

func (o DailyRetentionScheduleOutput) ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput {
	return o.ToDailyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (o DailyRetentionScheduleOutput) ToDailyRetentionSchedulePtrOutputWithContext(ctx context.Context) DailyRetentionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyRetentionSchedule) *DailyRetentionSchedule {
		return &v
	}).(DailyRetentionSchedulePtrOutput)
}

func (o DailyRetentionScheduleOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v DailyRetentionSchedule) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o DailyRetentionScheduleOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DailyRetentionSchedule) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type DailyRetentionSchedulePtrOutput struct{ *pulumi.OutputState }

func (DailyRetentionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionSchedule)(nil)).Elem()
}

func (o DailyRetentionSchedulePtrOutput) ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput {
	return o
}

func (o DailyRetentionSchedulePtrOutput) ToDailyRetentionSchedulePtrOutputWithContext(ctx context.Context) DailyRetentionSchedulePtrOutput {
	return o
}

func (o DailyRetentionSchedulePtrOutput) Elem() DailyRetentionScheduleOutput {
	return o.ApplyT(func(v *DailyRetentionSchedule) DailyRetentionSchedule {
		if v != nil {
			return *v
		}
		var ret DailyRetentionSchedule
		return ret
	}).(DailyRetentionScheduleOutput)
}

func (o DailyRetentionSchedulePtrOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v *DailyRetentionSchedule) *RetentionDuration {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationPtrOutput)
}

func (o DailyRetentionSchedulePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DailyRetentionSchedule) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type DailyRetentionScheduleResponse struct {
	RetentionDuration *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionTimes    []string                   `pulumi:"retentionTimes"`
}





type DailyRetentionScheduleResponseInput interface {
	pulumi.Input

	ToDailyRetentionScheduleResponseOutput() DailyRetentionScheduleResponseOutput
	ToDailyRetentionScheduleResponseOutputWithContext(context.Context) DailyRetentionScheduleResponseOutput
}

type DailyRetentionScheduleResponseArgs struct {
	RetentionDuration RetentionDurationResponsePtrInput `pulumi:"retentionDuration"`
	RetentionTimes    pulumi.StringArrayInput           `pulumi:"retentionTimes"`
}

func (DailyRetentionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionScheduleResponse)(nil)).Elem()
}

func (i DailyRetentionScheduleResponseArgs) ToDailyRetentionScheduleResponseOutput() DailyRetentionScheduleResponseOutput {
	return i.ToDailyRetentionScheduleResponseOutputWithContext(context.Background())
}

func (i DailyRetentionScheduleResponseArgs) ToDailyRetentionScheduleResponseOutputWithContext(ctx context.Context) DailyRetentionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleResponseOutput)
}

func (i DailyRetentionScheduleResponseArgs) ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput {
	return i.ToDailyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i DailyRetentionScheduleResponseArgs) ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) DailyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleResponseOutput).ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx)
}









type DailyRetentionScheduleResponsePtrInput interface {
	pulumi.Input

	ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput
	ToDailyRetentionScheduleResponsePtrOutputWithContext(context.Context) DailyRetentionScheduleResponsePtrOutput
}

type dailyRetentionScheduleResponsePtrType DailyRetentionScheduleResponseArgs

func DailyRetentionScheduleResponsePtr(v *DailyRetentionScheduleResponseArgs) DailyRetentionScheduleResponsePtrInput {
	return (*dailyRetentionScheduleResponsePtrType)(v)
}

func (*dailyRetentionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionScheduleResponse)(nil)).Elem()
}

func (i *dailyRetentionScheduleResponsePtrType) ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput {
	return i.ToDailyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *dailyRetentionScheduleResponsePtrType) ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) DailyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleResponsePtrOutput)
}

type DailyRetentionScheduleResponseOutput struct{ *pulumi.OutputState }

func (DailyRetentionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionScheduleResponse)(nil)).Elem()
}

func (o DailyRetentionScheduleResponseOutput) ToDailyRetentionScheduleResponseOutput() DailyRetentionScheduleResponseOutput {
	return o
}

func (o DailyRetentionScheduleResponseOutput) ToDailyRetentionScheduleResponseOutputWithContext(ctx context.Context) DailyRetentionScheduleResponseOutput {
	return o
}

func (o DailyRetentionScheduleResponseOutput) ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput {
	return o.ToDailyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o DailyRetentionScheduleResponseOutput) ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) DailyRetentionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyRetentionScheduleResponse) *DailyRetentionScheduleResponse {
		return &v
	}).(DailyRetentionScheduleResponsePtrOutput)
}

func (o DailyRetentionScheduleResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v DailyRetentionScheduleResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o DailyRetentionScheduleResponseOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DailyRetentionScheduleResponse) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type DailyRetentionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (DailyRetentionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionScheduleResponse)(nil)).Elem()
}

func (o DailyRetentionScheduleResponsePtrOutput) ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput {
	return o
}

func (o DailyRetentionScheduleResponsePtrOutput) ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) DailyRetentionScheduleResponsePtrOutput {
	return o
}

func (o DailyRetentionScheduleResponsePtrOutput) Elem() DailyRetentionScheduleResponseOutput {
	return o.ApplyT(func(v *DailyRetentionScheduleResponse) DailyRetentionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret DailyRetentionScheduleResponse
		return ret
	}).(DailyRetentionScheduleResponseOutput)
}

func (o DailyRetentionScheduleResponsePtrOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v *DailyRetentionScheduleResponse) *RetentionDurationResponse {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationResponsePtrOutput)
}

func (o DailyRetentionScheduleResponsePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DailyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type Day struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}





type DayInput interface {
	pulumi.Input

	ToDayOutput() DayOutput
	ToDayOutputWithContext(context.Context) DayOutput
}

type DayArgs struct {
	Date   pulumi.IntPtrInput  `pulumi:"date"`
	IsLast pulumi.BoolPtrInput `pulumi:"isLast"`
}

func (DayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Day)(nil)).Elem()
}

func (i DayArgs) ToDayOutput() DayOutput {
	return i.ToDayOutputWithContext(context.Background())
}

func (i DayArgs) ToDayOutputWithContext(ctx context.Context) DayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayOutput)
}





type DayArrayInput interface {
	pulumi.Input

	ToDayArrayOutput() DayArrayOutput
	ToDayArrayOutputWithContext(context.Context) DayArrayOutput
}

type DayArray []DayInput

func (DayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Day)(nil)).Elem()
}

func (i DayArray) ToDayArrayOutput() DayArrayOutput {
	return i.ToDayArrayOutputWithContext(context.Background())
}

func (i DayArray) ToDayArrayOutputWithContext(ctx context.Context) DayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayArrayOutput)
}

type DayOutput struct{ *pulumi.OutputState }

func (DayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Day)(nil)).Elem()
}

func (o DayOutput) ToDayOutput() DayOutput {
	return o
}

func (o DayOutput) ToDayOutputWithContext(ctx context.Context) DayOutput {
	return o
}

func (o DayOutput) Date() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Day) *int { return v.Date }).(pulumi.IntPtrOutput)
}

func (o DayOutput) IsLast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Day) *bool { return v.IsLast }).(pulumi.BoolPtrOutput)
}

type DayArrayOutput struct{ *pulumi.OutputState }

func (DayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Day)(nil)).Elem()
}

func (o DayArrayOutput) ToDayArrayOutput() DayArrayOutput {
	return o
}

func (o DayArrayOutput) ToDayArrayOutputWithContext(ctx context.Context) DayArrayOutput {
	return o
}

func (o DayArrayOutput) Index(i pulumi.IntInput) DayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Day {
		return vs[0].([]Day)[vs[1].(int)]
	}).(DayOutput)
}

type DayResponse struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}





type DayResponseInput interface {
	pulumi.Input

	ToDayResponseOutput() DayResponseOutput
	ToDayResponseOutputWithContext(context.Context) DayResponseOutput
}

type DayResponseArgs struct {
	Date   pulumi.IntPtrInput  `pulumi:"date"`
	IsLast pulumi.BoolPtrInput `pulumi:"isLast"`
}

func (DayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DayResponse)(nil)).Elem()
}

func (i DayResponseArgs) ToDayResponseOutput() DayResponseOutput {
	return i.ToDayResponseOutputWithContext(context.Background())
}

func (i DayResponseArgs) ToDayResponseOutputWithContext(ctx context.Context) DayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayResponseOutput)
}





type DayResponseArrayInput interface {
	pulumi.Input

	ToDayResponseArrayOutput() DayResponseArrayOutput
	ToDayResponseArrayOutputWithContext(context.Context) DayResponseArrayOutput
}

type DayResponseArray []DayResponseInput

func (DayResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayResponse)(nil)).Elem()
}

func (i DayResponseArray) ToDayResponseArrayOutput() DayResponseArrayOutput {
	return i.ToDayResponseArrayOutputWithContext(context.Background())
}

func (i DayResponseArray) ToDayResponseArrayOutputWithContext(ctx context.Context) DayResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayResponseArrayOutput)
}

type DayResponseOutput struct{ *pulumi.OutputState }

func (DayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayResponse)(nil)).Elem()
}

func (o DayResponseOutput) ToDayResponseOutput() DayResponseOutput {
	return o
}

func (o DayResponseOutput) ToDayResponseOutputWithContext(ctx context.Context) DayResponseOutput {
	return o
}

func (o DayResponseOutput) Date() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DayResponse) *int { return v.Date }).(pulumi.IntPtrOutput)
}

func (o DayResponseOutput) IsLast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DayResponse) *bool { return v.IsLast }).(pulumi.BoolPtrOutput)
}

type DayResponseArrayOutput struct{ *pulumi.OutputState }

func (DayResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayResponse)(nil)).Elem()
}

func (o DayResponseArrayOutput) ToDayResponseArrayOutput() DayResponseArrayOutput {
	return o
}

func (o DayResponseArrayOutput) ToDayResponseArrayOutputWithContext(ctx context.Context) DayResponseArrayOutput {
	return o
}

func (o DayResponseArrayOutput) Index(i pulumi.IntInput) DayResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DayResponse {
		return vs[0].([]DayResponse)[vs[1].(int)]
	}).(DayResponseOutput)
}

type DiskExclusionProperties struct {
	DiskLunList     []int `pulumi:"diskLunList"`
	IsInclusionList *bool `pulumi:"isInclusionList"`
}





type DiskExclusionPropertiesInput interface {
	pulumi.Input

	ToDiskExclusionPropertiesOutput() DiskExclusionPropertiesOutput
	ToDiskExclusionPropertiesOutputWithContext(context.Context) DiskExclusionPropertiesOutput
}

type DiskExclusionPropertiesArgs struct {
	DiskLunList     pulumi.IntArrayInput `pulumi:"diskLunList"`
	IsInclusionList pulumi.BoolPtrInput  `pulumi:"isInclusionList"`
}

func (DiskExclusionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskExclusionProperties)(nil)).Elem()
}

func (i DiskExclusionPropertiesArgs) ToDiskExclusionPropertiesOutput() DiskExclusionPropertiesOutput {
	return i.ToDiskExclusionPropertiesOutputWithContext(context.Background())
}

func (i DiskExclusionPropertiesArgs) ToDiskExclusionPropertiesOutputWithContext(ctx context.Context) DiskExclusionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskExclusionPropertiesOutput)
}

func (i DiskExclusionPropertiesArgs) ToDiskExclusionPropertiesPtrOutput() DiskExclusionPropertiesPtrOutput {
	return i.ToDiskExclusionPropertiesPtrOutputWithContext(context.Background())
}

func (i DiskExclusionPropertiesArgs) ToDiskExclusionPropertiesPtrOutputWithContext(ctx context.Context) DiskExclusionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskExclusionPropertiesOutput).ToDiskExclusionPropertiesPtrOutputWithContext(ctx)
}









type DiskExclusionPropertiesPtrInput interface {
	pulumi.Input

	ToDiskExclusionPropertiesPtrOutput() DiskExclusionPropertiesPtrOutput
	ToDiskExclusionPropertiesPtrOutputWithContext(context.Context) DiskExclusionPropertiesPtrOutput
}

type diskExclusionPropertiesPtrType DiskExclusionPropertiesArgs

func DiskExclusionPropertiesPtr(v *DiskExclusionPropertiesArgs) DiskExclusionPropertiesPtrInput {
	return (*diskExclusionPropertiesPtrType)(v)
}

func (*diskExclusionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskExclusionProperties)(nil)).Elem()
}

func (i *diskExclusionPropertiesPtrType) ToDiskExclusionPropertiesPtrOutput() DiskExclusionPropertiesPtrOutput {
	return i.ToDiskExclusionPropertiesPtrOutputWithContext(context.Background())
}

func (i *diskExclusionPropertiesPtrType) ToDiskExclusionPropertiesPtrOutputWithContext(ctx context.Context) DiskExclusionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskExclusionPropertiesPtrOutput)
}

type DiskExclusionPropertiesOutput struct{ *pulumi.OutputState }

func (DiskExclusionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskExclusionProperties)(nil)).Elem()
}

func (o DiskExclusionPropertiesOutput) ToDiskExclusionPropertiesOutput() DiskExclusionPropertiesOutput {
	return o
}

func (o DiskExclusionPropertiesOutput) ToDiskExclusionPropertiesOutputWithContext(ctx context.Context) DiskExclusionPropertiesOutput {
	return o
}

func (o DiskExclusionPropertiesOutput) ToDiskExclusionPropertiesPtrOutput() DiskExclusionPropertiesPtrOutput {
	return o.ToDiskExclusionPropertiesPtrOutputWithContext(context.Background())
}

func (o DiskExclusionPropertiesOutput) ToDiskExclusionPropertiesPtrOutputWithContext(ctx context.Context) DiskExclusionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskExclusionProperties) *DiskExclusionProperties {
		return &v
	}).(DiskExclusionPropertiesPtrOutput)
}

func (o DiskExclusionPropertiesOutput) DiskLunList() pulumi.IntArrayOutput {
	return o.ApplyT(func(v DiskExclusionProperties) []int { return v.DiskLunList }).(pulumi.IntArrayOutput)
}

func (o DiskExclusionPropertiesOutput) IsInclusionList() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskExclusionProperties) *bool { return v.IsInclusionList }).(pulumi.BoolPtrOutput)
}

type DiskExclusionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DiskExclusionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskExclusionProperties)(nil)).Elem()
}

func (o DiskExclusionPropertiesPtrOutput) ToDiskExclusionPropertiesPtrOutput() DiskExclusionPropertiesPtrOutput {
	return o
}

func (o DiskExclusionPropertiesPtrOutput) ToDiskExclusionPropertiesPtrOutputWithContext(ctx context.Context) DiskExclusionPropertiesPtrOutput {
	return o
}

func (o DiskExclusionPropertiesPtrOutput) Elem() DiskExclusionPropertiesOutput {
	return o.ApplyT(func(v *DiskExclusionProperties) DiskExclusionProperties {
		if v != nil {
			return *v
		}
		var ret DiskExclusionProperties
		return ret
	}).(DiskExclusionPropertiesOutput)
}

func (o DiskExclusionPropertiesPtrOutput) DiskLunList() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *DiskExclusionProperties) []int {
		if v == nil {
			return nil
		}
		return v.DiskLunList
	}).(pulumi.IntArrayOutput)
}

func (o DiskExclusionPropertiesPtrOutput) IsInclusionList() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskExclusionProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsInclusionList
	}).(pulumi.BoolPtrOutput)
}

type DiskExclusionPropertiesResponse struct {
	DiskLunList     []int `pulumi:"diskLunList"`
	IsInclusionList *bool `pulumi:"isInclusionList"`
}





type DiskExclusionPropertiesResponseInput interface {
	pulumi.Input

	ToDiskExclusionPropertiesResponseOutput() DiskExclusionPropertiesResponseOutput
	ToDiskExclusionPropertiesResponseOutputWithContext(context.Context) DiskExclusionPropertiesResponseOutput
}

type DiskExclusionPropertiesResponseArgs struct {
	DiskLunList     pulumi.IntArrayInput `pulumi:"diskLunList"`
	IsInclusionList pulumi.BoolPtrInput  `pulumi:"isInclusionList"`
}

func (DiskExclusionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskExclusionPropertiesResponse)(nil)).Elem()
}

func (i DiskExclusionPropertiesResponseArgs) ToDiskExclusionPropertiesResponseOutput() DiskExclusionPropertiesResponseOutput {
	return i.ToDiskExclusionPropertiesResponseOutputWithContext(context.Background())
}

func (i DiskExclusionPropertiesResponseArgs) ToDiskExclusionPropertiesResponseOutputWithContext(ctx context.Context) DiskExclusionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskExclusionPropertiesResponseOutput)
}

func (i DiskExclusionPropertiesResponseArgs) ToDiskExclusionPropertiesResponsePtrOutput() DiskExclusionPropertiesResponsePtrOutput {
	return i.ToDiskExclusionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DiskExclusionPropertiesResponseArgs) ToDiskExclusionPropertiesResponsePtrOutputWithContext(ctx context.Context) DiskExclusionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskExclusionPropertiesResponseOutput).ToDiskExclusionPropertiesResponsePtrOutputWithContext(ctx)
}









type DiskExclusionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToDiskExclusionPropertiesResponsePtrOutput() DiskExclusionPropertiesResponsePtrOutput
	ToDiskExclusionPropertiesResponsePtrOutputWithContext(context.Context) DiskExclusionPropertiesResponsePtrOutput
}

type diskExclusionPropertiesResponsePtrType DiskExclusionPropertiesResponseArgs

func DiskExclusionPropertiesResponsePtr(v *DiskExclusionPropertiesResponseArgs) DiskExclusionPropertiesResponsePtrInput {
	return (*diskExclusionPropertiesResponsePtrType)(v)
}

func (*diskExclusionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskExclusionPropertiesResponse)(nil)).Elem()
}

func (i *diskExclusionPropertiesResponsePtrType) ToDiskExclusionPropertiesResponsePtrOutput() DiskExclusionPropertiesResponsePtrOutput {
	return i.ToDiskExclusionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *diskExclusionPropertiesResponsePtrType) ToDiskExclusionPropertiesResponsePtrOutputWithContext(ctx context.Context) DiskExclusionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskExclusionPropertiesResponsePtrOutput)
}

type DiskExclusionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DiskExclusionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskExclusionPropertiesResponse)(nil)).Elem()
}

func (o DiskExclusionPropertiesResponseOutput) ToDiskExclusionPropertiesResponseOutput() DiskExclusionPropertiesResponseOutput {
	return o
}

func (o DiskExclusionPropertiesResponseOutput) ToDiskExclusionPropertiesResponseOutputWithContext(ctx context.Context) DiskExclusionPropertiesResponseOutput {
	return o
}

func (o DiskExclusionPropertiesResponseOutput) ToDiskExclusionPropertiesResponsePtrOutput() DiskExclusionPropertiesResponsePtrOutput {
	return o.ToDiskExclusionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DiskExclusionPropertiesResponseOutput) ToDiskExclusionPropertiesResponsePtrOutputWithContext(ctx context.Context) DiskExclusionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskExclusionPropertiesResponse) *DiskExclusionPropertiesResponse {
		return &v
	}).(DiskExclusionPropertiesResponsePtrOutput)
}

func (o DiskExclusionPropertiesResponseOutput) DiskLunList() pulumi.IntArrayOutput {
	return o.ApplyT(func(v DiskExclusionPropertiesResponse) []int { return v.DiskLunList }).(pulumi.IntArrayOutput)
}

func (o DiskExclusionPropertiesResponseOutput) IsInclusionList() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskExclusionPropertiesResponse) *bool { return v.IsInclusionList }).(pulumi.BoolPtrOutput)
}

type DiskExclusionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskExclusionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskExclusionPropertiesResponse)(nil)).Elem()
}

func (o DiskExclusionPropertiesResponsePtrOutput) ToDiskExclusionPropertiesResponsePtrOutput() DiskExclusionPropertiesResponsePtrOutput {
	return o
}

func (o DiskExclusionPropertiesResponsePtrOutput) ToDiskExclusionPropertiesResponsePtrOutputWithContext(ctx context.Context) DiskExclusionPropertiesResponsePtrOutput {
	return o
}

func (o DiskExclusionPropertiesResponsePtrOutput) Elem() DiskExclusionPropertiesResponseOutput {
	return o.ApplyT(func(v *DiskExclusionPropertiesResponse) DiskExclusionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DiskExclusionPropertiesResponse
		return ret
	}).(DiskExclusionPropertiesResponseOutput)
}

func (o DiskExclusionPropertiesResponsePtrOutput) DiskLunList() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *DiskExclusionPropertiesResponse) []int {
		if v == nil {
			return nil
		}
		return v.DiskLunList
	}).(pulumi.IntArrayOutput)
}

func (o DiskExclusionPropertiesResponsePtrOutput) IsInclusionList() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskExclusionPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsInclusionList
	}).(pulumi.BoolPtrOutput)
}

type DistributedNodesInfo struct {
	NodeName *string `pulumi:"nodeName"`
	Status   *string `pulumi:"status"`
}





type DistributedNodesInfoInput interface {
	pulumi.Input

	ToDistributedNodesInfoOutput() DistributedNodesInfoOutput
	ToDistributedNodesInfoOutputWithContext(context.Context) DistributedNodesInfoOutput
}

type DistributedNodesInfoArgs struct {
	NodeName pulumi.StringPtrInput `pulumi:"nodeName"`
	Status   pulumi.StringPtrInput `pulumi:"status"`
}

func (DistributedNodesInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributedNodesInfo)(nil)).Elem()
}

func (i DistributedNodesInfoArgs) ToDistributedNodesInfoOutput() DistributedNodesInfoOutput {
	return i.ToDistributedNodesInfoOutputWithContext(context.Background())
}

func (i DistributedNodesInfoArgs) ToDistributedNodesInfoOutputWithContext(ctx context.Context) DistributedNodesInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributedNodesInfoOutput)
}





type DistributedNodesInfoArrayInput interface {
	pulumi.Input

	ToDistributedNodesInfoArrayOutput() DistributedNodesInfoArrayOutput
	ToDistributedNodesInfoArrayOutputWithContext(context.Context) DistributedNodesInfoArrayOutput
}

type DistributedNodesInfoArray []DistributedNodesInfoInput

func (DistributedNodesInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DistributedNodesInfo)(nil)).Elem()
}

func (i DistributedNodesInfoArray) ToDistributedNodesInfoArrayOutput() DistributedNodesInfoArrayOutput {
	return i.ToDistributedNodesInfoArrayOutputWithContext(context.Background())
}

func (i DistributedNodesInfoArray) ToDistributedNodesInfoArrayOutputWithContext(ctx context.Context) DistributedNodesInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributedNodesInfoArrayOutput)
}

type DistributedNodesInfoOutput struct{ *pulumi.OutputState }

func (DistributedNodesInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributedNodesInfo)(nil)).Elem()
}

func (o DistributedNodesInfoOutput) ToDistributedNodesInfoOutput() DistributedNodesInfoOutput {
	return o
}

func (o DistributedNodesInfoOutput) ToDistributedNodesInfoOutputWithContext(ctx context.Context) DistributedNodesInfoOutput {
	return o
}

func (o DistributedNodesInfoOutput) NodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DistributedNodesInfo) *string { return v.NodeName }).(pulumi.StringPtrOutput)
}

func (o DistributedNodesInfoOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DistributedNodesInfo) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type DistributedNodesInfoArrayOutput struct{ *pulumi.OutputState }

func (DistributedNodesInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DistributedNodesInfo)(nil)).Elem()
}

func (o DistributedNodesInfoArrayOutput) ToDistributedNodesInfoArrayOutput() DistributedNodesInfoArrayOutput {
	return o
}

func (o DistributedNodesInfoArrayOutput) ToDistributedNodesInfoArrayOutputWithContext(ctx context.Context) DistributedNodesInfoArrayOutput {
	return o
}

func (o DistributedNodesInfoArrayOutput) Index(i pulumi.IntInput) DistributedNodesInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DistributedNodesInfo {
		return vs[0].([]DistributedNodesInfo)[vs[1].(int)]
	}).(DistributedNodesInfoOutput)
}

type DistributedNodesInfoResponse struct {
	ErrorDetail *ErrorDetailResponse `pulumi:"errorDetail"`
	NodeName    *string              `pulumi:"nodeName"`
	Status      *string              `pulumi:"status"`
}





type DistributedNodesInfoResponseInput interface {
	pulumi.Input

	ToDistributedNodesInfoResponseOutput() DistributedNodesInfoResponseOutput
	ToDistributedNodesInfoResponseOutputWithContext(context.Context) DistributedNodesInfoResponseOutput
}

type DistributedNodesInfoResponseArgs struct {
	ErrorDetail ErrorDetailResponsePtrInput `pulumi:"errorDetail"`
	NodeName    pulumi.StringPtrInput       `pulumi:"nodeName"`
	Status      pulumi.StringPtrInput       `pulumi:"status"`
}

func (DistributedNodesInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributedNodesInfoResponse)(nil)).Elem()
}

func (i DistributedNodesInfoResponseArgs) ToDistributedNodesInfoResponseOutput() DistributedNodesInfoResponseOutput {
	return i.ToDistributedNodesInfoResponseOutputWithContext(context.Background())
}

func (i DistributedNodesInfoResponseArgs) ToDistributedNodesInfoResponseOutputWithContext(ctx context.Context) DistributedNodesInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributedNodesInfoResponseOutput)
}





type DistributedNodesInfoResponseArrayInput interface {
	pulumi.Input

	ToDistributedNodesInfoResponseArrayOutput() DistributedNodesInfoResponseArrayOutput
	ToDistributedNodesInfoResponseArrayOutputWithContext(context.Context) DistributedNodesInfoResponseArrayOutput
}

type DistributedNodesInfoResponseArray []DistributedNodesInfoResponseInput

func (DistributedNodesInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DistributedNodesInfoResponse)(nil)).Elem()
}

func (i DistributedNodesInfoResponseArray) ToDistributedNodesInfoResponseArrayOutput() DistributedNodesInfoResponseArrayOutput {
	return i.ToDistributedNodesInfoResponseArrayOutputWithContext(context.Background())
}

func (i DistributedNodesInfoResponseArray) ToDistributedNodesInfoResponseArrayOutputWithContext(ctx context.Context) DistributedNodesInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributedNodesInfoResponseArrayOutput)
}

type DistributedNodesInfoResponseOutput struct{ *pulumi.OutputState }

func (DistributedNodesInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributedNodesInfoResponse)(nil)).Elem()
}

func (o DistributedNodesInfoResponseOutput) ToDistributedNodesInfoResponseOutput() DistributedNodesInfoResponseOutput {
	return o
}

func (o DistributedNodesInfoResponseOutput) ToDistributedNodesInfoResponseOutputWithContext(ctx context.Context) DistributedNodesInfoResponseOutput {
	return o
}

func (o DistributedNodesInfoResponseOutput) ErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v DistributedNodesInfoResponse) *ErrorDetailResponse { return v.ErrorDetail }).(ErrorDetailResponsePtrOutput)
}

func (o DistributedNodesInfoResponseOutput) NodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DistributedNodesInfoResponse) *string { return v.NodeName }).(pulumi.StringPtrOutput)
}

func (o DistributedNodesInfoResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DistributedNodesInfoResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type DistributedNodesInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (DistributedNodesInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DistributedNodesInfoResponse)(nil)).Elem()
}

func (o DistributedNodesInfoResponseArrayOutput) ToDistributedNodesInfoResponseArrayOutput() DistributedNodesInfoResponseArrayOutput {
	return o
}

func (o DistributedNodesInfoResponseArrayOutput) ToDistributedNodesInfoResponseArrayOutputWithContext(ctx context.Context) DistributedNodesInfoResponseArrayOutput {
	return o
}

func (o DistributedNodesInfoResponseArrayOutput) Index(i pulumi.IntInput) DistributedNodesInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DistributedNodesInfoResponse {
		return vs[0].([]DistributedNodesInfoResponse)[vs[1].(int)]
	}).(DistributedNodesInfoResponseOutput)
}

type DpmContainer struct {
	BackupManagementType *string                   `pulumi:"backupManagementType"`
	CanReRegister        *bool                     `pulumi:"canReRegister"`
	ContainerId          *string                   `pulumi:"containerId"`
	ContainerType        string                    `pulumi:"containerType"`
	DpmAgentVersion      *string                   `pulumi:"dpmAgentVersion"`
	DpmServers           []string                  `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                   `pulumi:"friendlyName"`
	HealthStatus         *string                   `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                  `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                   `pulumi:"protectionStatus"`
	RegistrationStatus   *string                   `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                     `pulumi:"upgradeAvailable"`
}





type DpmContainerInput interface {
	pulumi.Input

	ToDpmContainerOutput() DpmContainerOutput
	ToDpmContainerOutputWithContext(context.Context) DpmContainerOutput
}

type DpmContainerArgs struct {
	BackupManagementType pulumi.StringPtrInput            `pulumi:"backupManagementType"`
	CanReRegister        pulumi.BoolPtrInput              `pulumi:"canReRegister"`
	ContainerId          pulumi.StringPtrInput            `pulumi:"containerId"`
	ContainerType        pulumi.StringInput               `pulumi:"containerType"`
	DpmAgentVersion      pulumi.StringPtrInput            `pulumi:"dpmAgentVersion"`
	DpmServers           pulumi.StringArrayInput          `pulumi:"dpmServers"`
	ExtendedInfo         DPMContainerExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput            `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput            `pulumi:"healthStatus"`
	ProtectedItemCount   pulumi.Float64PtrInput           `pulumi:"protectedItemCount"`
	ProtectionStatus     pulumi.StringPtrInput            `pulumi:"protectionStatus"`
	RegistrationStatus   pulumi.StringPtrInput            `pulumi:"registrationStatus"`
	UpgradeAvailable     pulumi.BoolPtrInput              `pulumi:"upgradeAvailable"`
}

func (DpmContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DpmContainer)(nil)).Elem()
}

func (i DpmContainerArgs) ToDpmContainerOutput() DpmContainerOutput {
	return i.ToDpmContainerOutputWithContext(context.Background())
}

func (i DpmContainerArgs) ToDpmContainerOutputWithContext(ctx context.Context) DpmContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpmContainerOutput)
}

type DpmContainerOutput struct{ *pulumi.OutputState }

func (DpmContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DpmContainer)(nil)).Elem()
}

func (o DpmContainerOutput) ToDpmContainerOutput() DpmContainerOutput {
	return o
}

func (o DpmContainerOutput) ToDpmContainerOutputWithContext(ctx context.Context) DpmContainerOutput {
	return o
}

func (o DpmContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o DpmContainerOutput) CanReRegister() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DpmContainer) *bool { return v.CanReRegister }).(pulumi.BoolPtrOutput)
}

func (o DpmContainerOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainer) *string { return v.ContainerId }).(pulumi.StringPtrOutput)
}

func (o DpmContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v DpmContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o DpmContainerOutput) DpmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainer) *string { return v.DpmAgentVersion }).(pulumi.StringPtrOutput)
}

func (o DpmContainerOutput) DpmServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DpmContainer) []string { return v.DpmServers }).(pulumi.StringArrayOutput)
}

func (o DpmContainerOutput) ExtendedInfo() DPMContainerExtendedInfoPtrOutput {
	return o.ApplyT(func(v DpmContainer) *DPMContainerExtendedInfo { return v.ExtendedInfo }).(DPMContainerExtendedInfoPtrOutput)
}

func (o DpmContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o DpmContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o DpmContainerOutput) ProtectedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DpmContainer) *float64 { return v.ProtectedItemCount }).(pulumi.Float64PtrOutput)
}

func (o DpmContainerOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainer) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o DpmContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o DpmContainerOutput) UpgradeAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DpmContainer) *bool { return v.UpgradeAvailable }).(pulumi.BoolPtrOutput)
}

type DpmContainerResponse struct {
	BackupManagementType *string                           `pulumi:"backupManagementType"`
	CanReRegister        *bool                             `pulumi:"canReRegister"`
	ContainerId          *string                           `pulumi:"containerId"`
	ContainerType        string                            `pulumi:"containerType"`
	DpmAgentVersion      *string                           `pulumi:"dpmAgentVersion"`
	DpmServers           []string                          `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                           `pulumi:"friendlyName"`
	HealthStatus         *string                           `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                          `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                           `pulumi:"protectionStatus"`
	RegistrationStatus   *string                           `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                             `pulumi:"upgradeAvailable"`
}





type DpmContainerResponseInput interface {
	pulumi.Input

	ToDpmContainerResponseOutput() DpmContainerResponseOutput
	ToDpmContainerResponseOutputWithContext(context.Context) DpmContainerResponseOutput
}

type DpmContainerResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput                    `pulumi:"backupManagementType"`
	CanReRegister        pulumi.BoolPtrInput                      `pulumi:"canReRegister"`
	ContainerId          pulumi.StringPtrInput                    `pulumi:"containerId"`
	ContainerType        pulumi.StringInput                       `pulumi:"containerType"`
	DpmAgentVersion      pulumi.StringPtrInput                    `pulumi:"dpmAgentVersion"`
	DpmServers           pulumi.StringArrayInput                  `pulumi:"dpmServers"`
	ExtendedInfo         DPMContainerExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName         pulumi.StringPtrInput                    `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                    `pulumi:"healthStatus"`
	ProtectedItemCount   pulumi.Float64PtrInput                   `pulumi:"protectedItemCount"`
	ProtectionStatus     pulumi.StringPtrInput                    `pulumi:"protectionStatus"`
	RegistrationStatus   pulumi.StringPtrInput                    `pulumi:"registrationStatus"`
	UpgradeAvailable     pulumi.BoolPtrInput                      `pulumi:"upgradeAvailable"`
}

func (DpmContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DpmContainerResponse)(nil)).Elem()
}

func (i DpmContainerResponseArgs) ToDpmContainerResponseOutput() DpmContainerResponseOutput {
	return i.ToDpmContainerResponseOutputWithContext(context.Background())
}

func (i DpmContainerResponseArgs) ToDpmContainerResponseOutputWithContext(ctx context.Context) DpmContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpmContainerResponseOutput)
}

type DpmContainerResponseOutput struct{ *pulumi.OutputState }

func (DpmContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DpmContainerResponse)(nil)).Elem()
}

func (o DpmContainerResponseOutput) ToDpmContainerResponseOutput() DpmContainerResponseOutput {
	return o
}

func (o DpmContainerResponseOutput) ToDpmContainerResponseOutputWithContext(ctx context.Context) DpmContainerResponseOutput {
	return o
}

func (o DpmContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o DpmContainerResponseOutput) CanReRegister() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *bool { return v.CanReRegister }).(pulumi.BoolPtrOutput)
}

func (o DpmContainerResponseOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *string { return v.ContainerId }).(pulumi.StringPtrOutput)
}

func (o DpmContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v DpmContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o DpmContainerResponseOutput) DpmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *string { return v.DpmAgentVersion }).(pulumi.StringPtrOutput)
}

func (o DpmContainerResponseOutput) DpmServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DpmContainerResponse) []string { return v.DpmServers }).(pulumi.StringArrayOutput)
}

func (o DpmContainerResponseOutput) ExtendedInfo() DPMContainerExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *DPMContainerExtendedInfoResponse { return v.ExtendedInfo }).(DPMContainerExtendedInfoResponsePtrOutput)
}

func (o DpmContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o DpmContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o DpmContainerResponseOutput) ProtectedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *float64 { return v.ProtectedItemCount }).(pulumi.Float64PtrOutput)
}

func (o DpmContainerResponseOutput) ProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *string { return v.ProtectionStatus }).(pulumi.StringPtrOutput)
}

func (o DpmContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o DpmContainerResponseOutput) UpgradeAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DpmContainerResponse) *bool { return v.UpgradeAvailable }).(pulumi.BoolPtrOutput)
}

type ErrorDetailResponse struct {
	Code            string   `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
}





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	Code            pulumi.StringInput      `pulumi:"code"`
	Message         pulumi.StringInput      `pulumi:"message"`
	Recommendations pulumi.StringArrayInput `pulumi:"recommendations"`
}

func (ErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return i.ToErrorDetailResponseOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput)
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return i.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput).ToErrorDetailResponsePtrOutputWithContext(ctx)
}









type ErrorDetailResponsePtrInput interface {
	pulumi.Input

	ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput
	ToErrorDetailResponsePtrOutputWithContext(context.Context) ErrorDetailResponsePtrOutput
}

type errorDetailResponsePtrType ErrorDetailResponseArgs

func ErrorDetailResponsePtr(v *ErrorDetailResponseArgs) ErrorDetailResponsePtrInput {
	return (*errorDetailResponsePtrType)(v)
}

func (*errorDetailResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (i *errorDetailResponsePtrType) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return i.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (i *errorDetailResponsePtrType) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponsePtrOutput)
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorDetailResponse) *ErrorDetailResponse {
		return &v
	}).(ErrorDetailResponsePtrOutput)
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Recommendations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []string { return v.Recommendations }).(pulumi.StringArrayOutput)
}

type ErrorDetailResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) Elem() ErrorDetailResponseOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) ErrorDetailResponse {
		if v != nil {
			return *v
		}
		var ret ErrorDetailResponse
		return ret
	}).(ErrorDetailResponseOutput)
}

func (o ErrorDetailResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Recommendations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []string {
		if v == nil {
			return nil
		}
		return v.Recommendations
	}).(pulumi.StringArrayOutput)
}

type ExtendedProperties struct {
	DiskExclusionProperties *DiskExclusionProperties `pulumi:"diskExclusionProperties"`
}





type ExtendedPropertiesInput interface {
	pulumi.Input

	ToExtendedPropertiesOutput() ExtendedPropertiesOutput
	ToExtendedPropertiesOutputWithContext(context.Context) ExtendedPropertiesOutput
}

type ExtendedPropertiesArgs struct {
	DiskExclusionProperties DiskExclusionPropertiesPtrInput `pulumi:"diskExclusionProperties"`
}

func (ExtendedPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedProperties)(nil)).Elem()
}

func (i ExtendedPropertiesArgs) ToExtendedPropertiesOutput() ExtendedPropertiesOutput {
	return i.ToExtendedPropertiesOutputWithContext(context.Background())
}

func (i ExtendedPropertiesArgs) ToExtendedPropertiesOutputWithContext(ctx context.Context) ExtendedPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedPropertiesOutput)
}

func (i ExtendedPropertiesArgs) ToExtendedPropertiesPtrOutput() ExtendedPropertiesPtrOutput {
	return i.ToExtendedPropertiesPtrOutputWithContext(context.Background())
}

func (i ExtendedPropertiesArgs) ToExtendedPropertiesPtrOutputWithContext(ctx context.Context) ExtendedPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedPropertiesOutput).ToExtendedPropertiesPtrOutputWithContext(ctx)
}









type ExtendedPropertiesPtrInput interface {
	pulumi.Input

	ToExtendedPropertiesPtrOutput() ExtendedPropertiesPtrOutput
	ToExtendedPropertiesPtrOutputWithContext(context.Context) ExtendedPropertiesPtrOutput
}

type extendedPropertiesPtrType ExtendedPropertiesArgs

func ExtendedPropertiesPtr(v *ExtendedPropertiesArgs) ExtendedPropertiesPtrInput {
	return (*extendedPropertiesPtrType)(v)
}

func (*extendedPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedProperties)(nil)).Elem()
}

func (i *extendedPropertiesPtrType) ToExtendedPropertiesPtrOutput() ExtendedPropertiesPtrOutput {
	return i.ToExtendedPropertiesPtrOutputWithContext(context.Background())
}

func (i *extendedPropertiesPtrType) ToExtendedPropertiesPtrOutputWithContext(ctx context.Context) ExtendedPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedPropertiesPtrOutput)
}

type ExtendedPropertiesOutput struct{ *pulumi.OutputState }

func (ExtendedPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedProperties)(nil)).Elem()
}

func (o ExtendedPropertiesOutput) ToExtendedPropertiesOutput() ExtendedPropertiesOutput {
	return o
}

func (o ExtendedPropertiesOutput) ToExtendedPropertiesOutputWithContext(ctx context.Context) ExtendedPropertiesOutput {
	return o
}

func (o ExtendedPropertiesOutput) ToExtendedPropertiesPtrOutput() ExtendedPropertiesPtrOutput {
	return o.ToExtendedPropertiesPtrOutputWithContext(context.Background())
}

func (o ExtendedPropertiesOutput) ToExtendedPropertiesPtrOutputWithContext(ctx context.Context) ExtendedPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedProperties) *ExtendedProperties {
		return &v
	}).(ExtendedPropertiesPtrOutput)
}

func (o ExtendedPropertiesOutput) DiskExclusionProperties() DiskExclusionPropertiesPtrOutput {
	return o.ApplyT(func(v ExtendedProperties) *DiskExclusionProperties { return v.DiskExclusionProperties }).(DiskExclusionPropertiesPtrOutput)
}

type ExtendedPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ExtendedPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedProperties)(nil)).Elem()
}

func (o ExtendedPropertiesPtrOutput) ToExtendedPropertiesPtrOutput() ExtendedPropertiesPtrOutput {
	return o
}

func (o ExtendedPropertiesPtrOutput) ToExtendedPropertiesPtrOutputWithContext(ctx context.Context) ExtendedPropertiesPtrOutput {
	return o
}

func (o ExtendedPropertiesPtrOutput) Elem() ExtendedPropertiesOutput {
	return o.ApplyT(func(v *ExtendedProperties) ExtendedProperties {
		if v != nil {
			return *v
		}
		var ret ExtendedProperties
		return ret
	}).(ExtendedPropertiesOutput)
}

func (o ExtendedPropertiesPtrOutput) DiskExclusionProperties() DiskExclusionPropertiesPtrOutput {
	return o.ApplyT(func(v *ExtendedProperties) *DiskExclusionProperties {
		if v == nil {
			return nil
		}
		return v.DiskExclusionProperties
	}).(DiskExclusionPropertiesPtrOutput)
}

type ExtendedPropertiesResponse struct {
	DiskExclusionProperties *DiskExclusionPropertiesResponse `pulumi:"diskExclusionProperties"`
}





type ExtendedPropertiesResponseInput interface {
	pulumi.Input

	ToExtendedPropertiesResponseOutput() ExtendedPropertiesResponseOutput
	ToExtendedPropertiesResponseOutputWithContext(context.Context) ExtendedPropertiesResponseOutput
}

type ExtendedPropertiesResponseArgs struct {
	DiskExclusionProperties DiskExclusionPropertiesResponsePtrInput `pulumi:"diskExclusionProperties"`
}

func (ExtendedPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedPropertiesResponse)(nil)).Elem()
}

func (i ExtendedPropertiesResponseArgs) ToExtendedPropertiesResponseOutput() ExtendedPropertiesResponseOutput {
	return i.ToExtendedPropertiesResponseOutputWithContext(context.Background())
}

func (i ExtendedPropertiesResponseArgs) ToExtendedPropertiesResponseOutputWithContext(ctx context.Context) ExtendedPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedPropertiesResponseOutput)
}

func (i ExtendedPropertiesResponseArgs) ToExtendedPropertiesResponsePtrOutput() ExtendedPropertiesResponsePtrOutput {
	return i.ToExtendedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ExtendedPropertiesResponseArgs) ToExtendedPropertiesResponsePtrOutputWithContext(ctx context.Context) ExtendedPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedPropertiesResponseOutput).ToExtendedPropertiesResponsePtrOutputWithContext(ctx)
}









type ExtendedPropertiesResponsePtrInput interface {
	pulumi.Input

	ToExtendedPropertiesResponsePtrOutput() ExtendedPropertiesResponsePtrOutput
	ToExtendedPropertiesResponsePtrOutputWithContext(context.Context) ExtendedPropertiesResponsePtrOutput
}

type extendedPropertiesResponsePtrType ExtendedPropertiesResponseArgs

func ExtendedPropertiesResponsePtr(v *ExtendedPropertiesResponseArgs) ExtendedPropertiesResponsePtrInput {
	return (*extendedPropertiesResponsePtrType)(v)
}

func (*extendedPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedPropertiesResponse)(nil)).Elem()
}

func (i *extendedPropertiesResponsePtrType) ToExtendedPropertiesResponsePtrOutput() ExtendedPropertiesResponsePtrOutput {
	return i.ToExtendedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *extendedPropertiesResponsePtrType) ToExtendedPropertiesResponsePtrOutputWithContext(ctx context.Context) ExtendedPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedPropertiesResponsePtrOutput)
}

type ExtendedPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ExtendedPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedPropertiesResponse)(nil)).Elem()
}

func (o ExtendedPropertiesResponseOutput) ToExtendedPropertiesResponseOutput() ExtendedPropertiesResponseOutput {
	return o
}

func (o ExtendedPropertiesResponseOutput) ToExtendedPropertiesResponseOutputWithContext(ctx context.Context) ExtendedPropertiesResponseOutput {
	return o
}

func (o ExtendedPropertiesResponseOutput) ToExtendedPropertiesResponsePtrOutput() ExtendedPropertiesResponsePtrOutput {
	return o.ToExtendedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ExtendedPropertiesResponseOutput) ToExtendedPropertiesResponsePtrOutputWithContext(ctx context.Context) ExtendedPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedPropertiesResponse) *ExtendedPropertiesResponse {
		return &v
	}).(ExtendedPropertiesResponsePtrOutput)
}

func (o ExtendedPropertiesResponseOutput) DiskExclusionProperties() DiskExclusionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ExtendedPropertiesResponse) *DiskExclusionPropertiesResponse { return v.DiskExclusionProperties }).(DiskExclusionPropertiesResponsePtrOutput)
}

type ExtendedPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedPropertiesResponse)(nil)).Elem()
}

func (o ExtendedPropertiesResponsePtrOutput) ToExtendedPropertiesResponsePtrOutput() ExtendedPropertiesResponsePtrOutput {
	return o
}

func (o ExtendedPropertiesResponsePtrOutput) ToExtendedPropertiesResponsePtrOutputWithContext(ctx context.Context) ExtendedPropertiesResponsePtrOutput {
	return o
}

func (o ExtendedPropertiesResponsePtrOutput) Elem() ExtendedPropertiesResponseOutput {
	return o.ApplyT(func(v *ExtendedPropertiesResponse) ExtendedPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedPropertiesResponse
		return ret
	}).(ExtendedPropertiesResponseOutput)
}

func (o ExtendedPropertiesResponsePtrOutput) DiskExclusionProperties() DiskExclusionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ExtendedPropertiesResponse) *DiskExclusionPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.DiskExclusionProperties
	}).(DiskExclusionPropertiesResponsePtrOutput)
}

type GenericContainer struct {
	BackupManagementType *string                       `pulumi:"backupManagementType"`
	ContainerType        string                        `pulumi:"containerType"`
	ExtendedInformation  *GenericContainerExtendedInfo `pulumi:"extendedInformation"`
	FabricName           *string                       `pulumi:"fabricName"`
	FriendlyName         *string                       `pulumi:"friendlyName"`
	HealthStatus         *string                       `pulumi:"healthStatus"`
	RegistrationStatus   *string                       `pulumi:"registrationStatus"`
}





type GenericContainerInput interface {
	pulumi.Input

	ToGenericContainerOutput() GenericContainerOutput
	ToGenericContainerOutputWithContext(context.Context) GenericContainerOutput
}

type GenericContainerArgs struct {
	BackupManagementType pulumi.StringPtrInput                `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput                   `pulumi:"containerType"`
	ExtendedInformation  GenericContainerExtendedInfoPtrInput `pulumi:"extendedInformation"`
	FabricName           pulumi.StringPtrInput                `pulumi:"fabricName"`
	FriendlyName         pulumi.StringPtrInput                `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                `pulumi:"healthStatus"`
	RegistrationStatus   pulumi.StringPtrInput                `pulumi:"registrationStatus"`
}

func (GenericContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericContainer)(nil)).Elem()
}

func (i GenericContainerArgs) ToGenericContainerOutput() GenericContainerOutput {
	return i.ToGenericContainerOutputWithContext(context.Background())
}

func (i GenericContainerArgs) ToGenericContainerOutputWithContext(ctx context.Context) GenericContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericContainerOutput)
}

type GenericContainerOutput struct{ *pulumi.OutputState }

func (GenericContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericContainer)(nil)).Elem()
}

func (o GenericContainerOutput) ToGenericContainerOutput() GenericContainerOutput {
	return o
}

func (o GenericContainerOutput) ToGenericContainerOutputWithContext(ctx context.Context) GenericContainerOutput {
	return o
}

func (o GenericContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o GenericContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v GenericContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o GenericContainerOutput) ExtendedInformation() GenericContainerExtendedInfoPtrOutput {
	return o.ApplyT(func(v GenericContainer) *GenericContainerExtendedInfo { return v.ExtendedInformation }).(GenericContainerExtendedInfoPtrOutput)
}

func (o GenericContainerOutput) FabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainer) *string { return v.FabricName }).(pulumi.StringPtrOutput)
}

func (o GenericContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o GenericContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o GenericContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

type GenericContainerExtendedInfo struct {
	ContainerIdentityInfo *ContainerIdentityInfo `pulumi:"containerIdentityInfo"`
	RawCertData           *string                `pulumi:"rawCertData"`
	ServiceEndpoints      map[string]string      `pulumi:"serviceEndpoints"`
}





type GenericContainerExtendedInfoInput interface {
	pulumi.Input

	ToGenericContainerExtendedInfoOutput() GenericContainerExtendedInfoOutput
	ToGenericContainerExtendedInfoOutputWithContext(context.Context) GenericContainerExtendedInfoOutput
}

type GenericContainerExtendedInfoArgs struct {
	ContainerIdentityInfo ContainerIdentityInfoPtrInput `pulumi:"containerIdentityInfo"`
	RawCertData           pulumi.StringPtrInput         `pulumi:"rawCertData"`
	ServiceEndpoints      pulumi.StringMapInput         `pulumi:"serviceEndpoints"`
}

func (GenericContainerExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericContainerExtendedInfo)(nil)).Elem()
}

func (i GenericContainerExtendedInfoArgs) ToGenericContainerExtendedInfoOutput() GenericContainerExtendedInfoOutput {
	return i.ToGenericContainerExtendedInfoOutputWithContext(context.Background())
}

func (i GenericContainerExtendedInfoArgs) ToGenericContainerExtendedInfoOutputWithContext(ctx context.Context) GenericContainerExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericContainerExtendedInfoOutput)
}

func (i GenericContainerExtendedInfoArgs) ToGenericContainerExtendedInfoPtrOutput() GenericContainerExtendedInfoPtrOutput {
	return i.ToGenericContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (i GenericContainerExtendedInfoArgs) ToGenericContainerExtendedInfoPtrOutputWithContext(ctx context.Context) GenericContainerExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericContainerExtendedInfoOutput).ToGenericContainerExtendedInfoPtrOutputWithContext(ctx)
}









type GenericContainerExtendedInfoPtrInput interface {
	pulumi.Input

	ToGenericContainerExtendedInfoPtrOutput() GenericContainerExtendedInfoPtrOutput
	ToGenericContainerExtendedInfoPtrOutputWithContext(context.Context) GenericContainerExtendedInfoPtrOutput
}

type genericContainerExtendedInfoPtrType GenericContainerExtendedInfoArgs

func GenericContainerExtendedInfoPtr(v *GenericContainerExtendedInfoArgs) GenericContainerExtendedInfoPtrInput {
	return (*genericContainerExtendedInfoPtrType)(v)
}

func (*genericContainerExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericContainerExtendedInfo)(nil)).Elem()
}

func (i *genericContainerExtendedInfoPtrType) ToGenericContainerExtendedInfoPtrOutput() GenericContainerExtendedInfoPtrOutput {
	return i.ToGenericContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *genericContainerExtendedInfoPtrType) ToGenericContainerExtendedInfoPtrOutputWithContext(ctx context.Context) GenericContainerExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericContainerExtendedInfoPtrOutput)
}

type GenericContainerExtendedInfoOutput struct{ *pulumi.OutputState }

func (GenericContainerExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericContainerExtendedInfo)(nil)).Elem()
}

func (o GenericContainerExtendedInfoOutput) ToGenericContainerExtendedInfoOutput() GenericContainerExtendedInfoOutput {
	return o
}

func (o GenericContainerExtendedInfoOutput) ToGenericContainerExtendedInfoOutputWithContext(ctx context.Context) GenericContainerExtendedInfoOutput {
	return o
}

func (o GenericContainerExtendedInfoOutput) ToGenericContainerExtendedInfoPtrOutput() GenericContainerExtendedInfoPtrOutput {
	return o.ToGenericContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (o GenericContainerExtendedInfoOutput) ToGenericContainerExtendedInfoPtrOutputWithContext(ctx context.Context) GenericContainerExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GenericContainerExtendedInfo) *GenericContainerExtendedInfo {
		return &v
	}).(GenericContainerExtendedInfoPtrOutput)
}

func (o GenericContainerExtendedInfoOutput) ContainerIdentityInfo() ContainerIdentityInfoPtrOutput {
	return o.ApplyT(func(v GenericContainerExtendedInfo) *ContainerIdentityInfo { return v.ContainerIdentityInfo }).(ContainerIdentityInfoPtrOutput)
}

func (o GenericContainerExtendedInfoOutput) RawCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainerExtendedInfo) *string { return v.RawCertData }).(pulumi.StringPtrOutput)
}

func (o GenericContainerExtendedInfoOutput) ServiceEndpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v GenericContainerExtendedInfo) map[string]string { return v.ServiceEndpoints }).(pulumi.StringMapOutput)
}

type GenericContainerExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (GenericContainerExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericContainerExtendedInfo)(nil)).Elem()
}

func (o GenericContainerExtendedInfoPtrOutput) ToGenericContainerExtendedInfoPtrOutput() GenericContainerExtendedInfoPtrOutput {
	return o
}

func (o GenericContainerExtendedInfoPtrOutput) ToGenericContainerExtendedInfoPtrOutputWithContext(ctx context.Context) GenericContainerExtendedInfoPtrOutput {
	return o
}

func (o GenericContainerExtendedInfoPtrOutput) Elem() GenericContainerExtendedInfoOutput {
	return o.ApplyT(func(v *GenericContainerExtendedInfo) GenericContainerExtendedInfo {
		if v != nil {
			return *v
		}
		var ret GenericContainerExtendedInfo
		return ret
	}).(GenericContainerExtendedInfoOutput)
}

func (o GenericContainerExtendedInfoPtrOutput) ContainerIdentityInfo() ContainerIdentityInfoPtrOutput {
	return o.ApplyT(func(v *GenericContainerExtendedInfo) *ContainerIdentityInfo {
		if v == nil {
			return nil
		}
		return v.ContainerIdentityInfo
	}).(ContainerIdentityInfoPtrOutput)
}

func (o GenericContainerExtendedInfoPtrOutput) RawCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GenericContainerExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.RawCertData
	}).(pulumi.StringPtrOutput)
}

func (o GenericContainerExtendedInfoPtrOutput) ServiceEndpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GenericContainerExtendedInfo) map[string]string {
		if v == nil {
			return nil
		}
		return v.ServiceEndpoints
	}).(pulumi.StringMapOutput)
}

type GenericContainerExtendedInfoResponse struct {
	ContainerIdentityInfo *ContainerIdentityInfoResponse `pulumi:"containerIdentityInfo"`
	RawCertData           *string                        `pulumi:"rawCertData"`
	ServiceEndpoints      map[string]string              `pulumi:"serviceEndpoints"`
}





type GenericContainerExtendedInfoResponseInput interface {
	pulumi.Input

	ToGenericContainerExtendedInfoResponseOutput() GenericContainerExtendedInfoResponseOutput
	ToGenericContainerExtendedInfoResponseOutputWithContext(context.Context) GenericContainerExtendedInfoResponseOutput
}

type GenericContainerExtendedInfoResponseArgs struct {
	ContainerIdentityInfo ContainerIdentityInfoResponsePtrInput `pulumi:"containerIdentityInfo"`
	RawCertData           pulumi.StringPtrInput                 `pulumi:"rawCertData"`
	ServiceEndpoints      pulumi.StringMapInput                 `pulumi:"serviceEndpoints"`
}

func (GenericContainerExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericContainerExtendedInfoResponse)(nil)).Elem()
}

func (i GenericContainerExtendedInfoResponseArgs) ToGenericContainerExtendedInfoResponseOutput() GenericContainerExtendedInfoResponseOutput {
	return i.ToGenericContainerExtendedInfoResponseOutputWithContext(context.Background())
}

func (i GenericContainerExtendedInfoResponseArgs) ToGenericContainerExtendedInfoResponseOutputWithContext(ctx context.Context) GenericContainerExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericContainerExtendedInfoResponseOutput)
}

func (i GenericContainerExtendedInfoResponseArgs) ToGenericContainerExtendedInfoResponsePtrOutput() GenericContainerExtendedInfoResponsePtrOutput {
	return i.ToGenericContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i GenericContainerExtendedInfoResponseArgs) ToGenericContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) GenericContainerExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericContainerExtendedInfoResponseOutput).ToGenericContainerExtendedInfoResponsePtrOutputWithContext(ctx)
}









type GenericContainerExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToGenericContainerExtendedInfoResponsePtrOutput() GenericContainerExtendedInfoResponsePtrOutput
	ToGenericContainerExtendedInfoResponsePtrOutputWithContext(context.Context) GenericContainerExtendedInfoResponsePtrOutput
}

type genericContainerExtendedInfoResponsePtrType GenericContainerExtendedInfoResponseArgs

func GenericContainerExtendedInfoResponsePtr(v *GenericContainerExtendedInfoResponseArgs) GenericContainerExtendedInfoResponsePtrInput {
	return (*genericContainerExtendedInfoResponsePtrType)(v)
}

func (*genericContainerExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericContainerExtendedInfoResponse)(nil)).Elem()
}

func (i *genericContainerExtendedInfoResponsePtrType) ToGenericContainerExtendedInfoResponsePtrOutput() GenericContainerExtendedInfoResponsePtrOutput {
	return i.ToGenericContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *genericContainerExtendedInfoResponsePtrType) ToGenericContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) GenericContainerExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericContainerExtendedInfoResponsePtrOutput)
}

type GenericContainerExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (GenericContainerExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericContainerExtendedInfoResponse)(nil)).Elem()
}

func (o GenericContainerExtendedInfoResponseOutput) ToGenericContainerExtendedInfoResponseOutput() GenericContainerExtendedInfoResponseOutput {
	return o
}

func (o GenericContainerExtendedInfoResponseOutput) ToGenericContainerExtendedInfoResponseOutputWithContext(ctx context.Context) GenericContainerExtendedInfoResponseOutput {
	return o
}

func (o GenericContainerExtendedInfoResponseOutput) ToGenericContainerExtendedInfoResponsePtrOutput() GenericContainerExtendedInfoResponsePtrOutput {
	return o.ToGenericContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o GenericContainerExtendedInfoResponseOutput) ToGenericContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) GenericContainerExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GenericContainerExtendedInfoResponse) *GenericContainerExtendedInfoResponse {
		return &v
	}).(GenericContainerExtendedInfoResponsePtrOutput)
}

func (o GenericContainerExtendedInfoResponseOutput) ContainerIdentityInfo() ContainerIdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v GenericContainerExtendedInfoResponse) *ContainerIdentityInfoResponse {
		return v.ContainerIdentityInfo
	}).(ContainerIdentityInfoResponsePtrOutput)
}

func (o GenericContainerExtendedInfoResponseOutput) RawCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainerExtendedInfoResponse) *string { return v.RawCertData }).(pulumi.StringPtrOutput)
}

func (o GenericContainerExtendedInfoResponseOutput) ServiceEndpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v GenericContainerExtendedInfoResponse) map[string]string { return v.ServiceEndpoints }).(pulumi.StringMapOutput)
}

type GenericContainerExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (GenericContainerExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericContainerExtendedInfoResponse)(nil)).Elem()
}

func (o GenericContainerExtendedInfoResponsePtrOutput) ToGenericContainerExtendedInfoResponsePtrOutput() GenericContainerExtendedInfoResponsePtrOutput {
	return o
}

func (o GenericContainerExtendedInfoResponsePtrOutput) ToGenericContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) GenericContainerExtendedInfoResponsePtrOutput {
	return o
}

func (o GenericContainerExtendedInfoResponsePtrOutput) Elem() GenericContainerExtendedInfoResponseOutput {
	return o.ApplyT(func(v *GenericContainerExtendedInfoResponse) GenericContainerExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret GenericContainerExtendedInfoResponse
		return ret
	}).(GenericContainerExtendedInfoResponseOutput)
}

func (o GenericContainerExtendedInfoResponsePtrOutput) ContainerIdentityInfo() ContainerIdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v *GenericContainerExtendedInfoResponse) *ContainerIdentityInfoResponse {
		if v == nil {
			return nil
		}
		return v.ContainerIdentityInfo
	}).(ContainerIdentityInfoResponsePtrOutput)
}

func (o GenericContainerExtendedInfoResponsePtrOutput) RawCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GenericContainerExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RawCertData
	}).(pulumi.StringPtrOutput)
}

func (o GenericContainerExtendedInfoResponsePtrOutput) ServiceEndpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GenericContainerExtendedInfoResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.ServiceEndpoints
	}).(pulumi.StringMapOutput)
}

type GenericContainerResponse struct {
	BackupManagementType *string                               `pulumi:"backupManagementType"`
	ContainerType        string                                `pulumi:"containerType"`
	ExtendedInformation  *GenericContainerExtendedInfoResponse `pulumi:"extendedInformation"`
	FabricName           *string                               `pulumi:"fabricName"`
	FriendlyName         *string                               `pulumi:"friendlyName"`
	HealthStatus         *string                               `pulumi:"healthStatus"`
	RegistrationStatus   *string                               `pulumi:"registrationStatus"`
}





type GenericContainerResponseInput interface {
	pulumi.Input

	ToGenericContainerResponseOutput() GenericContainerResponseOutput
	ToGenericContainerResponseOutputWithContext(context.Context) GenericContainerResponseOutput
}

type GenericContainerResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput                        `pulumi:"backupManagementType"`
	ContainerType        pulumi.StringInput                           `pulumi:"containerType"`
	ExtendedInformation  GenericContainerExtendedInfoResponsePtrInput `pulumi:"extendedInformation"`
	FabricName           pulumi.StringPtrInput                        `pulumi:"fabricName"`
	FriendlyName         pulumi.StringPtrInput                        `pulumi:"friendlyName"`
	HealthStatus         pulumi.StringPtrInput                        `pulumi:"healthStatus"`
	RegistrationStatus   pulumi.StringPtrInput                        `pulumi:"registrationStatus"`
}

func (GenericContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericContainerResponse)(nil)).Elem()
}

func (i GenericContainerResponseArgs) ToGenericContainerResponseOutput() GenericContainerResponseOutput {
	return i.ToGenericContainerResponseOutputWithContext(context.Background())
}

func (i GenericContainerResponseArgs) ToGenericContainerResponseOutputWithContext(ctx context.Context) GenericContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericContainerResponseOutput)
}

type GenericContainerResponseOutput struct{ *pulumi.OutputState }

func (GenericContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericContainerResponse)(nil)).Elem()
}

func (o GenericContainerResponseOutput) ToGenericContainerResponseOutput() GenericContainerResponseOutput {
	return o
}

func (o GenericContainerResponseOutput) ToGenericContainerResponseOutputWithContext(ctx context.Context) GenericContainerResponseOutput {
	return o
}

func (o GenericContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o GenericContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v GenericContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o GenericContainerResponseOutput) ExtendedInformation() GenericContainerExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v GenericContainerResponse) *GenericContainerExtendedInfoResponse { return v.ExtendedInformation }).(GenericContainerExtendedInfoResponsePtrOutput)
}

func (o GenericContainerResponseOutput) FabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainerResponse) *string { return v.FabricName }).(pulumi.StringPtrOutput)
}

func (o GenericContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o GenericContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o GenericContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

type GenericProtectedItem struct {
	BackupManagementType             *string           `pulumi:"backupManagementType"`
	BackupSetName                    *string           `pulumi:"backupSetName"`
	ContainerName                    *string           `pulumi:"containerName"`
	CreateMode                       *string           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string           `pulumi:"deferredDeleteTimeRemaining"`
	FabricName                       *string           `pulumi:"fabricName"`
	FriendlyName                     *string           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool             `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string           `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string           `pulumi:"policyId"`
	PolicyState                      *string           `pulumi:"policyState"`
	ProtectedItemId                  *float64          `pulumi:"protectedItemId"`
	ProtectedItemType                string            `pulumi:"protectedItemType"`
	ProtectionState                  *string           `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string          `pulumi:"resourceGuardOperationRequests"`
	SourceAssociations               map[string]string `pulumi:"sourceAssociations"`
	SourceResourceId                 *string           `pulumi:"sourceResourceId"`
	WorkloadType                     *string           `pulumi:"workloadType"`
}





type GenericProtectedItemInput interface {
	pulumi.Input

	ToGenericProtectedItemOutput() GenericProtectedItemOutput
	ToGenericProtectedItemOutputWithContext(context.Context) GenericProtectedItemOutput
}

type GenericProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput   `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput   `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput   `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput   `pulumi:"deferredDeleteTimeRemaining"`
	FabricName                       pulumi.StringPtrInput   `pulumi:"fabricName"`
	FriendlyName                     pulumi.StringPtrInput   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput     `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                pulumi.StringPtrInput   `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput   `pulumi:"policyId"`
	PolicyState                      pulumi.StringPtrInput   `pulumi:"policyState"`
	ProtectedItemId                  pulumi.Float64PtrInput  `pulumi:"protectedItemId"`
	ProtectedItemType                pulumi.StringInput      `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput   `pulumi:"protectionState"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput `pulumi:"resourceGuardOperationRequests"`
	SourceAssociations               pulumi.StringMapInput   `pulumi:"sourceAssociations"`
	SourceResourceId                 pulumi.StringPtrInput   `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput   `pulumi:"workloadType"`
}

func (GenericProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericProtectedItem)(nil)).Elem()
}

func (i GenericProtectedItemArgs) ToGenericProtectedItemOutput() GenericProtectedItemOutput {
	return i.ToGenericProtectedItemOutputWithContext(context.Background())
}

func (i GenericProtectedItemArgs) ToGenericProtectedItemOutputWithContext(ctx context.Context) GenericProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericProtectedItemOutput)
}

type GenericProtectedItemOutput struct{ *pulumi.OutputState }

func (GenericProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericProtectedItem)(nil)).Elem()
}

func (o GenericProtectedItemOutput) ToGenericProtectedItemOutput() GenericProtectedItemOutput {
	return o
}

func (o GenericProtectedItemOutput) ToGenericProtectedItemOutputWithContext(ctx context.Context) GenericProtectedItemOutput {
	return o
}

func (o GenericProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) FabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.FabricName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o GenericProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o GenericProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o GenericProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.PolicyState }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) ProtectedItemId() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *float64 { return v.ProtectedItemId }).(pulumi.Float64PtrOutput)
}

func (o GenericProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v GenericProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o GenericProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GenericProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o GenericProtectedItemOutput) SourceAssociations() pulumi.StringMapOutput {
	return o.ApplyT(func(v GenericProtectedItem) map[string]string { return v.SourceAssociations }).(pulumi.StringMapOutput)
}

func (o GenericProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type GenericProtectedItemResponse struct {
	BackupManagementType             *string           `pulumi:"backupManagementType"`
	BackupSetName                    *string           `pulumi:"backupSetName"`
	ContainerName                    *string           `pulumi:"containerName"`
	CreateMode                       *string           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string           `pulumi:"deferredDeleteTimeRemaining"`
	FabricName                       *string           `pulumi:"fabricName"`
	FriendlyName                     *string           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool             `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string           `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string           `pulumi:"policyId"`
	PolicyState                      *string           `pulumi:"policyState"`
	ProtectedItemId                  *float64          `pulumi:"protectedItemId"`
	ProtectedItemType                string            `pulumi:"protectedItemType"`
	ProtectionState                  *string           `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string          `pulumi:"resourceGuardOperationRequests"`
	SourceAssociations               map[string]string `pulumi:"sourceAssociations"`
	SourceResourceId                 *string           `pulumi:"sourceResourceId"`
	WorkloadType                     *string           `pulumi:"workloadType"`
}





type GenericProtectedItemResponseInput interface {
	pulumi.Input

	ToGenericProtectedItemResponseOutput() GenericProtectedItemResponseOutput
	ToGenericProtectedItemResponseOutputWithContext(context.Context) GenericProtectedItemResponseOutput
}

type GenericProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput   `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput   `pulumi:"backupSetName"`
	ContainerName                    pulumi.StringPtrInput   `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput   `pulumi:"deferredDeleteTimeRemaining"`
	FabricName                       pulumi.StringPtrInput   `pulumi:"fabricName"`
	FriendlyName                     pulumi.StringPtrInput   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput     `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                pulumi.StringPtrInput   `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput   `pulumi:"policyId"`
	PolicyState                      pulumi.StringPtrInput   `pulumi:"policyState"`
	ProtectedItemId                  pulumi.Float64PtrInput  `pulumi:"protectedItemId"`
	ProtectedItemType                pulumi.StringInput      `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput   `pulumi:"protectionState"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput `pulumi:"resourceGuardOperationRequests"`
	SourceAssociations               pulumi.StringMapInput   `pulumi:"sourceAssociations"`
	SourceResourceId                 pulumi.StringPtrInput   `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput   `pulumi:"workloadType"`
}

func (GenericProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericProtectedItemResponse)(nil)).Elem()
}

func (i GenericProtectedItemResponseArgs) ToGenericProtectedItemResponseOutput() GenericProtectedItemResponseOutput {
	return i.ToGenericProtectedItemResponseOutputWithContext(context.Background())
}

func (i GenericProtectedItemResponseArgs) ToGenericProtectedItemResponseOutputWithContext(ctx context.Context) GenericProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericProtectedItemResponseOutput)
}

type GenericProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (GenericProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericProtectedItemResponse)(nil)).Elem()
}

func (o GenericProtectedItemResponseOutput) ToGenericProtectedItemResponseOutput() GenericProtectedItemResponseOutput {
	return o
}

func (o GenericProtectedItemResponseOutput) ToGenericProtectedItemResponseOutputWithContext(ctx context.Context) GenericProtectedItemResponseOutput {
	return o
}

func (o GenericProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) FabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.FabricName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o GenericProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o GenericProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o GenericProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) PolicyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.PolicyState }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) ProtectedItemId() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *float64 { return v.ProtectedItemId }).(pulumi.Float64PtrOutput)
}

func (o GenericProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o GenericProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o GenericProtectedItemResponseOutput) SourceAssociations() pulumi.StringMapOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) map[string]string { return v.SourceAssociations }).(pulumi.StringMapOutput)
}

func (o GenericProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o GenericProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type GenericProtectionPolicy struct {
	BackupManagementType           string                `pulumi:"backupManagementType"`
	FabricName                     *string               `pulumi:"fabricName"`
	ProtectedItemsCount            *int                  `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string              `pulumi:"resourceGuardOperationRequests"`
	SubProtectionPolicy            []SubProtectionPolicy `pulumi:"subProtectionPolicy"`
	TimeZone                       *string               `pulumi:"timeZone"`
}





type GenericProtectionPolicyInput interface {
	pulumi.Input

	ToGenericProtectionPolicyOutput() GenericProtectionPolicyOutput
	ToGenericProtectionPolicyOutputWithContext(context.Context) GenericProtectionPolicyOutput
}

type GenericProtectionPolicyArgs struct {
	BackupManagementType           pulumi.StringInput            `pulumi:"backupManagementType"`
	FabricName                     pulumi.StringPtrInput         `pulumi:"fabricName"`
	ProtectedItemsCount            pulumi.IntPtrInput            `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput       `pulumi:"resourceGuardOperationRequests"`
	SubProtectionPolicy            SubProtectionPolicyArrayInput `pulumi:"subProtectionPolicy"`
	TimeZone                       pulumi.StringPtrInput         `pulumi:"timeZone"`
}

func (GenericProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericProtectionPolicy)(nil)).Elem()
}

func (i GenericProtectionPolicyArgs) ToGenericProtectionPolicyOutput() GenericProtectionPolicyOutput {
	return i.ToGenericProtectionPolicyOutputWithContext(context.Background())
}

func (i GenericProtectionPolicyArgs) ToGenericProtectionPolicyOutputWithContext(ctx context.Context) GenericProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericProtectionPolicyOutput)
}

type GenericProtectionPolicyOutput struct{ *pulumi.OutputState }

func (GenericProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericProtectionPolicy)(nil)).Elem()
}

func (o GenericProtectionPolicyOutput) ToGenericProtectionPolicyOutput() GenericProtectionPolicyOutput {
	return o
}

func (o GenericProtectionPolicyOutput) ToGenericProtectionPolicyOutputWithContext(ctx context.Context) GenericProtectionPolicyOutput {
	return o
}

func (o GenericProtectionPolicyOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v GenericProtectionPolicy) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o GenericProtectionPolicyOutput) FabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectionPolicy) *string { return v.FabricName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenericProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o GenericProtectionPolicyOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GenericProtectionPolicy) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o GenericProtectionPolicyOutput) SubProtectionPolicy() SubProtectionPolicyArrayOutput {
	return o.ApplyT(func(v GenericProtectionPolicy) []SubProtectionPolicy { return v.SubProtectionPolicy }).(SubProtectionPolicyArrayOutput)
}

func (o GenericProtectionPolicyOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectionPolicy) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type GenericProtectionPolicyResponse struct {
	BackupManagementType           string                        `pulumi:"backupManagementType"`
	FabricName                     *string                       `pulumi:"fabricName"`
	ProtectedItemsCount            *int                          `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string                      `pulumi:"resourceGuardOperationRequests"`
	SubProtectionPolicy            []SubProtectionPolicyResponse `pulumi:"subProtectionPolicy"`
	TimeZone                       *string                       `pulumi:"timeZone"`
}





type GenericProtectionPolicyResponseInput interface {
	pulumi.Input

	ToGenericProtectionPolicyResponseOutput() GenericProtectionPolicyResponseOutput
	ToGenericProtectionPolicyResponseOutputWithContext(context.Context) GenericProtectionPolicyResponseOutput
}

type GenericProtectionPolicyResponseArgs struct {
	BackupManagementType           pulumi.StringInput                    `pulumi:"backupManagementType"`
	FabricName                     pulumi.StringPtrInput                 `pulumi:"fabricName"`
	ProtectedItemsCount            pulumi.IntPtrInput                    `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput               `pulumi:"resourceGuardOperationRequests"`
	SubProtectionPolicy            SubProtectionPolicyResponseArrayInput `pulumi:"subProtectionPolicy"`
	TimeZone                       pulumi.StringPtrInput                 `pulumi:"timeZone"`
}

func (GenericProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericProtectionPolicyResponse)(nil)).Elem()
}

func (i GenericProtectionPolicyResponseArgs) ToGenericProtectionPolicyResponseOutput() GenericProtectionPolicyResponseOutput {
	return i.ToGenericProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i GenericProtectionPolicyResponseArgs) ToGenericProtectionPolicyResponseOutputWithContext(ctx context.Context) GenericProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericProtectionPolicyResponseOutput)
}

type GenericProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (GenericProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenericProtectionPolicyResponse)(nil)).Elem()
}

func (o GenericProtectionPolicyResponseOutput) ToGenericProtectionPolicyResponseOutput() GenericProtectionPolicyResponseOutput {
	return o
}

func (o GenericProtectionPolicyResponseOutput) ToGenericProtectionPolicyResponseOutputWithContext(ctx context.Context) GenericProtectionPolicyResponseOutput {
	return o
}

func (o GenericProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v GenericProtectionPolicyResponse) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o GenericProtectionPolicyResponseOutput) FabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectionPolicyResponse) *string { return v.FabricName }).(pulumi.StringPtrOutput)
}

func (o GenericProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenericProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o GenericProtectionPolicyResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GenericProtectionPolicyResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o GenericProtectionPolicyResponseOutput) SubProtectionPolicy() SubProtectionPolicyResponseArrayOutput {
	return o.ApplyT(func(v GenericProtectionPolicyResponse) []SubProtectionPolicyResponse { return v.SubProtectionPolicy }).(SubProtectionPolicyResponseArrayOutput)
}

func (o GenericProtectionPolicyResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenericProtectionPolicyResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type HourlySchedule struct {
	Interval                *int    `pulumi:"interval"`
	ScheduleWindowDuration  *int    `pulumi:"scheduleWindowDuration"`
	ScheduleWindowStartTime *string `pulumi:"scheduleWindowStartTime"`
}





type HourlyScheduleInput interface {
	pulumi.Input

	ToHourlyScheduleOutput() HourlyScheduleOutput
	ToHourlyScheduleOutputWithContext(context.Context) HourlyScheduleOutput
}

type HourlyScheduleArgs struct {
	Interval                pulumi.IntPtrInput    `pulumi:"interval"`
	ScheduleWindowDuration  pulumi.IntPtrInput    `pulumi:"scheduleWindowDuration"`
	ScheduleWindowStartTime pulumi.StringPtrInput `pulumi:"scheduleWindowStartTime"`
}

func (HourlyScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HourlySchedule)(nil)).Elem()
}

func (i HourlyScheduleArgs) ToHourlyScheduleOutput() HourlyScheduleOutput {
	return i.ToHourlyScheduleOutputWithContext(context.Background())
}

func (i HourlyScheduleArgs) ToHourlyScheduleOutputWithContext(ctx context.Context) HourlyScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleOutput)
}

func (i HourlyScheduleArgs) ToHourlySchedulePtrOutput() HourlySchedulePtrOutput {
	return i.ToHourlySchedulePtrOutputWithContext(context.Background())
}

func (i HourlyScheduleArgs) ToHourlySchedulePtrOutputWithContext(ctx context.Context) HourlySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleOutput).ToHourlySchedulePtrOutputWithContext(ctx)
}









type HourlySchedulePtrInput interface {
	pulumi.Input

	ToHourlySchedulePtrOutput() HourlySchedulePtrOutput
	ToHourlySchedulePtrOutputWithContext(context.Context) HourlySchedulePtrOutput
}

type hourlySchedulePtrType HourlyScheduleArgs

func HourlySchedulePtr(v *HourlyScheduleArgs) HourlySchedulePtrInput {
	return (*hourlySchedulePtrType)(v)
}

func (*hourlySchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HourlySchedule)(nil)).Elem()
}

func (i *hourlySchedulePtrType) ToHourlySchedulePtrOutput() HourlySchedulePtrOutput {
	return i.ToHourlySchedulePtrOutputWithContext(context.Background())
}

func (i *hourlySchedulePtrType) ToHourlySchedulePtrOutputWithContext(ctx context.Context) HourlySchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlySchedulePtrOutput)
}

type HourlyScheduleOutput struct{ *pulumi.OutputState }

func (HourlyScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HourlySchedule)(nil)).Elem()
}

func (o HourlyScheduleOutput) ToHourlyScheduleOutput() HourlyScheduleOutput {
	return o
}

func (o HourlyScheduleOutput) ToHourlyScheduleOutputWithContext(ctx context.Context) HourlyScheduleOutput {
	return o
}

func (o HourlyScheduleOutput) ToHourlySchedulePtrOutput() HourlySchedulePtrOutput {
	return o.ToHourlySchedulePtrOutputWithContext(context.Background())
}

func (o HourlyScheduleOutput) ToHourlySchedulePtrOutputWithContext(ctx context.Context) HourlySchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HourlySchedule) *HourlySchedule {
		return &v
	}).(HourlySchedulePtrOutput)
}

func (o HourlyScheduleOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourlySchedule) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleOutput) ScheduleWindowDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourlySchedule) *int { return v.ScheduleWindowDuration }).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleOutput) ScheduleWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HourlySchedule) *string { return v.ScheduleWindowStartTime }).(pulumi.StringPtrOutput)
}

type HourlySchedulePtrOutput struct{ *pulumi.OutputState }

func (HourlySchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourlySchedule)(nil)).Elem()
}

func (o HourlySchedulePtrOutput) ToHourlySchedulePtrOutput() HourlySchedulePtrOutput {
	return o
}

func (o HourlySchedulePtrOutput) ToHourlySchedulePtrOutputWithContext(ctx context.Context) HourlySchedulePtrOutput {
	return o
}

func (o HourlySchedulePtrOutput) Elem() HourlyScheduleOutput {
	return o.ApplyT(func(v *HourlySchedule) HourlySchedule {
		if v != nil {
			return *v
		}
		var ret HourlySchedule
		return ret
	}).(HourlyScheduleOutput)
}

func (o HourlySchedulePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourlySchedule) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o HourlySchedulePtrOutput) ScheduleWindowDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourlySchedule) *int {
		if v == nil {
			return nil
		}
		return v.ScheduleWindowDuration
	}).(pulumi.IntPtrOutput)
}

func (o HourlySchedulePtrOutput) ScheduleWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HourlySchedule) *string {
		if v == nil {
			return nil
		}
		return v.ScheduleWindowStartTime
	}).(pulumi.StringPtrOutput)
}

type HourlyScheduleResponse struct {
	Interval                *int    `pulumi:"interval"`
	ScheduleWindowDuration  *int    `pulumi:"scheduleWindowDuration"`
	ScheduleWindowStartTime *string `pulumi:"scheduleWindowStartTime"`
}





type HourlyScheduleResponseInput interface {
	pulumi.Input

	ToHourlyScheduleResponseOutput() HourlyScheduleResponseOutput
	ToHourlyScheduleResponseOutputWithContext(context.Context) HourlyScheduleResponseOutput
}

type HourlyScheduleResponseArgs struct {
	Interval                pulumi.IntPtrInput    `pulumi:"interval"`
	ScheduleWindowDuration  pulumi.IntPtrInput    `pulumi:"scheduleWindowDuration"`
	ScheduleWindowStartTime pulumi.StringPtrInput `pulumi:"scheduleWindowStartTime"`
}

func (HourlyScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HourlyScheduleResponse)(nil)).Elem()
}

func (i HourlyScheduleResponseArgs) ToHourlyScheduleResponseOutput() HourlyScheduleResponseOutput {
	return i.ToHourlyScheduleResponseOutputWithContext(context.Background())
}

func (i HourlyScheduleResponseArgs) ToHourlyScheduleResponseOutputWithContext(ctx context.Context) HourlyScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleResponseOutput)
}

func (i HourlyScheduleResponseArgs) ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput {
	return i.ToHourlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i HourlyScheduleResponseArgs) ToHourlyScheduleResponsePtrOutputWithContext(ctx context.Context) HourlyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleResponseOutput).ToHourlyScheduleResponsePtrOutputWithContext(ctx)
}









type HourlyScheduleResponsePtrInput interface {
	pulumi.Input

	ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput
	ToHourlyScheduleResponsePtrOutputWithContext(context.Context) HourlyScheduleResponsePtrOutput
}

type hourlyScheduleResponsePtrType HourlyScheduleResponseArgs

func HourlyScheduleResponsePtr(v *HourlyScheduleResponseArgs) HourlyScheduleResponsePtrInput {
	return (*hourlyScheduleResponsePtrType)(v)
}

func (*hourlyScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HourlyScheduleResponse)(nil)).Elem()
}

func (i *hourlyScheduleResponsePtrType) ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput {
	return i.ToHourlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *hourlyScheduleResponsePtrType) ToHourlyScheduleResponsePtrOutputWithContext(ctx context.Context) HourlyScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourlyScheduleResponsePtrOutput)
}

type HourlyScheduleResponseOutput struct{ *pulumi.OutputState }

func (HourlyScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HourlyScheduleResponse)(nil)).Elem()
}

func (o HourlyScheduleResponseOutput) ToHourlyScheduleResponseOutput() HourlyScheduleResponseOutput {
	return o
}

func (o HourlyScheduleResponseOutput) ToHourlyScheduleResponseOutputWithContext(ctx context.Context) HourlyScheduleResponseOutput {
	return o
}

func (o HourlyScheduleResponseOutput) ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput {
	return o.ToHourlyScheduleResponsePtrOutputWithContext(context.Background())
}

func (o HourlyScheduleResponseOutput) ToHourlyScheduleResponsePtrOutputWithContext(ctx context.Context) HourlyScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HourlyScheduleResponse) *HourlyScheduleResponse {
		return &v
	}).(HourlyScheduleResponsePtrOutput)
}

func (o HourlyScheduleResponseOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourlyScheduleResponse) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleResponseOutput) ScheduleWindowDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourlyScheduleResponse) *int { return v.ScheduleWindowDuration }).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleResponseOutput) ScheduleWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HourlyScheduleResponse) *string { return v.ScheduleWindowStartTime }).(pulumi.StringPtrOutput)
}

type HourlyScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (HourlyScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourlyScheduleResponse)(nil)).Elem()
}

func (o HourlyScheduleResponsePtrOutput) ToHourlyScheduleResponsePtrOutput() HourlyScheduleResponsePtrOutput {
	return o
}

func (o HourlyScheduleResponsePtrOutput) ToHourlyScheduleResponsePtrOutputWithContext(ctx context.Context) HourlyScheduleResponsePtrOutput {
	return o
}

func (o HourlyScheduleResponsePtrOutput) Elem() HourlyScheduleResponseOutput {
	return o.ApplyT(func(v *HourlyScheduleResponse) HourlyScheduleResponse {
		if v != nil {
			return *v
		}
		var ret HourlyScheduleResponse
		return ret
	}).(HourlyScheduleResponseOutput)
}

func (o HourlyScheduleResponsePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourlyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleResponsePtrOutput) ScheduleWindowDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourlyScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.ScheduleWindowDuration
	}).(pulumi.IntPtrOutput)
}

func (o HourlyScheduleResponsePtrOutput) ScheduleWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HourlyScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScheduleWindowStartTime
	}).(pulumi.StringPtrOutput)
}

type IaaSVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}





type IaaSVMContainerInput interface {
	pulumi.Input

	ToIaaSVMContainerOutput() IaaSVMContainerOutput
	ToIaaSVMContainerOutputWithContext(context.Context) IaaSVMContainerOutput
}

type IaaSVMContainerArgs struct {
	BackupManagementType  pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ContainerType         pulumi.StringInput    `pulumi:"containerType"`
	FriendlyName          pulumi.StringPtrInput `pulumi:"friendlyName"`
	HealthStatus          pulumi.StringPtrInput `pulumi:"healthStatus"`
	RegistrationStatus    pulumi.StringPtrInput `pulumi:"registrationStatus"`
	ResourceGroup         pulumi.StringPtrInput `pulumi:"resourceGroup"`
	VirtualMachineId      pulumi.StringPtrInput `pulumi:"virtualMachineId"`
	VirtualMachineVersion pulumi.StringPtrInput `pulumi:"virtualMachineVersion"`
}

func (IaaSVMContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IaaSVMContainer)(nil)).Elem()
}

func (i IaaSVMContainerArgs) ToIaaSVMContainerOutput() IaaSVMContainerOutput {
	return i.ToIaaSVMContainerOutputWithContext(context.Background())
}

func (i IaaSVMContainerArgs) ToIaaSVMContainerOutputWithContext(ctx context.Context) IaaSVMContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IaaSVMContainerOutput)
}

type IaaSVMContainerOutput struct{ *pulumi.OutputState }

func (IaaSVMContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IaaSVMContainer)(nil)).Elem()
}

func (o IaaSVMContainerOutput) ToIaaSVMContainerOutput() IaaSVMContainerOutput {
	return o
}

func (o IaaSVMContainerOutput) ToIaaSVMContainerOutputWithContext(ctx context.Context) IaaSVMContainerOutput {
	return o
}

func (o IaaSVMContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v IaaSVMContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o IaaSVMContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainer) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainer) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerOutput) VirtualMachineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainer) *string { return v.VirtualMachineVersion }).(pulumi.StringPtrOutput)
}

type IaaSVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}





type IaaSVMContainerResponseInput interface {
	pulumi.Input

	ToIaaSVMContainerResponseOutput() IaaSVMContainerResponseOutput
	ToIaaSVMContainerResponseOutputWithContext(context.Context) IaaSVMContainerResponseOutput
}

type IaaSVMContainerResponseArgs struct {
	BackupManagementType  pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ContainerType         pulumi.StringInput    `pulumi:"containerType"`
	FriendlyName          pulumi.StringPtrInput `pulumi:"friendlyName"`
	HealthStatus          pulumi.StringPtrInput `pulumi:"healthStatus"`
	RegistrationStatus    pulumi.StringPtrInput `pulumi:"registrationStatus"`
	ResourceGroup         pulumi.StringPtrInput `pulumi:"resourceGroup"`
	VirtualMachineId      pulumi.StringPtrInput `pulumi:"virtualMachineId"`
	VirtualMachineVersion pulumi.StringPtrInput `pulumi:"virtualMachineVersion"`
}

func (IaaSVMContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IaaSVMContainerResponse)(nil)).Elem()
}

func (i IaaSVMContainerResponseArgs) ToIaaSVMContainerResponseOutput() IaaSVMContainerResponseOutput {
	return i.ToIaaSVMContainerResponseOutputWithContext(context.Background())
}

func (i IaaSVMContainerResponseArgs) ToIaaSVMContainerResponseOutputWithContext(ctx context.Context) IaaSVMContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IaaSVMContainerResponseOutput)
}

type IaaSVMContainerResponseOutput struct{ *pulumi.OutputState }

func (IaaSVMContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IaaSVMContainerResponse)(nil)).Elem()
}

func (o IaaSVMContainerResponseOutput) ToIaaSVMContainerResponseOutput() IaaSVMContainerResponseOutput {
	return o
}

func (o IaaSVMContainerResponseOutput) ToIaaSVMContainerResponseOutputWithContext(ctx context.Context) IaaSVMContainerResponseOutput {
	return o
}

func (o IaaSVMContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v IaaSVMContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o IaaSVMContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainerResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerResponseOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainerResponse) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

func (o IaaSVMContainerResponseOutput) VirtualMachineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IaaSVMContainerResponse) *string { return v.VirtualMachineVersion }).(pulumi.StringPtrOutput)
}

type InquiryInfo struct {
	InquiryDetails []WorkloadInquiryDetails `pulumi:"inquiryDetails"`
	Status         *string                  `pulumi:"status"`
}





type InquiryInfoInput interface {
	pulumi.Input

	ToInquiryInfoOutput() InquiryInfoOutput
	ToInquiryInfoOutputWithContext(context.Context) InquiryInfoOutput
}

type InquiryInfoArgs struct {
	InquiryDetails WorkloadInquiryDetailsArrayInput `pulumi:"inquiryDetails"`
	Status         pulumi.StringPtrInput            `pulumi:"status"`
}

func (InquiryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InquiryInfo)(nil)).Elem()
}

func (i InquiryInfoArgs) ToInquiryInfoOutput() InquiryInfoOutput {
	return i.ToInquiryInfoOutputWithContext(context.Background())
}

func (i InquiryInfoArgs) ToInquiryInfoOutputWithContext(ctx context.Context) InquiryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryInfoOutput)
}

func (i InquiryInfoArgs) ToInquiryInfoPtrOutput() InquiryInfoPtrOutput {
	return i.ToInquiryInfoPtrOutputWithContext(context.Background())
}

func (i InquiryInfoArgs) ToInquiryInfoPtrOutputWithContext(ctx context.Context) InquiryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryInfoOutput).ToInquiryInfoPtrOutputWithContext(ctx)
}









type InquiryInfoPtrInput interface {
	pulumi.Input

	ToInquiryInfoPtrOutput() InquiryInfoPtrOutput
	ToInquiryInfoPtrOutputWithContext(context.Context) InquiryInfoPtrOutput
}

type inquiryInfoPtrType InquiryInfoArgs

func InquiryInfoPtr(v *InquiryInfoArgs) InquiryInfoPtrInput {
	return (*inquiryInfoPtrType)(v)
}

func (*inquiryInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InquiryInfo)(nil)).Elem()
}

func (i *inquiryInfoPtrType) ToInquiryInfoPtrOutput() InquiryInfoPtrOutput {
	return i.ToInquiryInfoPtrOutputWithContext(context.Background())
}

func (i *inquiryInfoPtrType) ToInquiryInfoPtrOutputWithContext(ctx context.Context) InquiryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryInfoPtrOutput)
}

type InquiryInfoOutput struct{ *pulumi.OutputState }

func (InquiryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InquiryInfo)(nil)).Elem()
}

func (o InquiryInfoOutput) ToInquiryInfoOutput() InquiryInfoOutput {
	return o
}

func (o InquiryInfoOutput) ToInquiryInfoOutputWithContext(ctx context.Context) InquiryInfoOutput {
	return o
}

func (o InquiryInfoOutput) ToInquiryInfoPtrOutput() InquiryInfoPtrOutput {
	return o.ToInquiryInfoPtrOutputWithContext(context.Background())
}

func (o InquiryInfoOutput) ToInquiryInfoPtrOutputWithContext(ctx context.Context) InquiryInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InquiryInfo) *InquiryInfo {
		return &v
	}).(InquiryInfoPtrOutput)
}

func (o InquiryInfoOutput) InquiryDetails() WorkloadInquiryDetailsArrayOutput {
	return o.ApplyT(func(v InquiryInfo) []WorkloadInquiryDetails { return v.InquiryDetails }).(WorkloadInquiryDetailsArrayOutput)
}

func (o InquiryInfoOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InquiryInfo) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type InquiryInfoPtrOutput struct{ *pulumi.OutputState }

func (InquiryInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InquiryInfo)(nil)).Elem()
}

func (o InquiryInfoPtrOutput) ToInquiryInfoPtrOutput() InquiryInfoPtrOutput {
	return o
}

func (o InquiryInfoPtrOutput) ToInquiryInfoPtrOutputWithContext(ctx context.Context) InquiryInfoPtrOutput {
	return o
}

func (o InquiryInfoPtrOutput) Elem() InquiryInfoOutput {
	return o.ApplyT(func(v *InquiryInfo) InquiryInfo {
		if v != nil {
			return *v
		}
		var ret InquiryInfo
		return ret
	}).(InquiryInfoOutput)
}

func (o InquiryInfoPtrOutput) InquiryDetails() WorkloadInquiryDetailsArrayOutput {
	return o.ApplyT(func(v *InquiryInfo) []WorkloadInquiryDetails {
		if v == nil {
			return nil
		}
		return v.InquiryDetails
	}).(WorkloadInquiryDetailsArrayOutput)
}

func (o InquiryInfoPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InquiryInfo) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type InquiryInfoResponse struct {
	ErrorDetail    *ErrorDetailResponse             `pulumi:"errorDetail"`
	InquiryDetails []WorkloadInquiryDetailsResponse `pulumi:"inquiryDetails"`
	Status         *string                          `pulumi:"status"`
}





type InquiryInfoResponseInput interface {
	pulumi.Input

	ToInquiryInfoResponseOutput() InquiryInfoResponseOutput
	ToInquiryInfoResponseOutputWithContext(context.Context) InquiryInfoResponseOutput
}

type InquiryInfoResponseArgs struct {
	ErrorDetail    ErrorDetailResponsePtrInput              `pulumi:"errorDetail"`
	InquiryDetails WorkloadInquiryDetailsResponseArrayInput `pulumi:"inquiryDetails"`
	Status         pulumi.StringPtrInput                    `pulumi:"status"`
}

func (InquiryInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InquiryInfoResponse)(nil)).Elem()
}

func (i InquiryInfoResponseArgs) ToInquiryInfoResponseOutput() InquiryInfoResponseOutput {
	return i.ToInquiryInfoResponseOutputWithContext(context.Background())
}

func (i InquiryInfoResponseArgs) ToInquiryInfoResponseOutputWithContext(ctx context.Context) InquiryInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryInfoResponseOutput)
}

func (i InquiryInfoResponseArgs) ToInquiryInfoResponsePtrOutput() InquiryInfoResponsePtrOutput {
	return i.ToInquiryInfoResponsePtrOutputWithContext(context.Background())
}

func (i InquiryInfoResponseArgs) ToInquiryInfoResponsePtrOutputWithContext(ctx context.Context) InquiryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryInfoResponseOutput).ToInquiryInfoResponsePtrOutputWithContext(ctx)
}









type InquiryInfoResponsePtrInput interface {
	pulumi.Input

	ToInquiryInfoResponsePtrOutput() InquiryInfoResponsePtrOutput
	ToInquiryInfoResponsePtrOutputWithContext(context.Context) InquiryInfoResponsePtrOutput
}

type inquiryInfoResponsePtrType InquiryInfoResponseArgs

func InquiryInfoResponsePtr(v *InquiryInfoResponseArgs) InquiryInfoResponsePtrInput {
	return (*inquiryInfoResponsePtrType)(v)
}

func (*inquiryInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InquiryInfoResponse)(nil)).Elem()
}

func (i *inquiryInfoResponsePtrType) ToInquiryInfoResponsePtrOutput() InquiryInfoResponsePtrOutput {
	return i.ToInquiryInfoResponsePtrOutputWithContext(context.Background())
}

func (i *inquiryInfoResponsePtrType) ToInquiryInfoResponsePtrOutputWithContext(ctx context.Context) InquiryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryInfoResponsePtrOutput)
}

type InquiryInfoResponseOutput struct{ *pulumi.OutputState }

func (InquiryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InquiryInfoResponse)(nil)).Elem()
}

func (o InquiryInfoResponseOutput) ToInquiryInfoResponseOutput() InquiryInfoResponseOutput {
	return o
}

func (o InquiryInfoResponseOutput) ToInquiryInfoResponseOutputWithContext(ctx context.Context) InquiryInfoResponseOutput {
	return o
}

func (o InquiryInfoResponseOutput) ToInquiryInfoResponsePtrOutput() InquiryInfoResponsePtrOutput {
	return o.ToInquiryInfoResponsePtrOutputWithContext(context.Background())
}

func (o InquiryInfoResponseOutput) ToInquiryInfoResponsePtrOutputWithContext(ctx context.Context) InquiryInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InquiryInfoResponse) *InquiryInfoResponse {
		return &v
	}).(InquiryInfoResponsePtrOutput)
}

func (o InquiryInfoResponseOutput) ErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v InquiryInfoResponse) *ErrorDetailResponse { return v.ErrorDetail }).(ErrorDetailResponsePtrOutput)
}

func (o InquiryInfoResponseOutput) InquiryDetails() WorkloadInquiryDetailsResponseArrayOutput {
	return o.ApplyT(func(v InquiryInfoResponse) []WorkloadInquiryDetailsResponse { return v.InquiryDetails }).(WorkloadInquiryDetailsResponseArrayOutput)
}

func (o InquiryInfoResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InquiryInfoResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type InquiryInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (InquiryInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InquiryInfoResponse)(nil)).Elem()
}

func (o InquiryInfoResponsePtrOutput) ToInquiryInfoResponsePtrOutput() InquiryInfoResponsePtrOutput {
	return o
}

func (o InquiryInfoResponsePtrOutput) ToInquiryInfoResponsePtrOutputWithContext(ctx context.Context) InquiryInfoResponsePtrOutput {
	return o
}

func (o InquiryInfoResponsePtrOutput) Elem() InquiryInfoResponseOutput {
	return o.ApplyT(func(v *InquiryInfoResponse) InquiryInfoResponse {
		if v != nil {
			return *v
		}
		var ret InquiryInfoResponse
		return ret
	}).(InquiryInfoResponseOutput)
}

func (o InquiryInfoResponsePtrOutput) ErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v *InquiryInfoResponse) *ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.ErrorDetail
	}).(ErrorDetailResponsePtrOutput)
}

func (o InquiryInfoResponsePtrOutput) InquiryDetails() WorkloadInquiryDetailsResponseArrayOutput {
	return o.ApplyT(func(v *InquiryInfoResponse) []WorkloadInquiryDetailsResponse {
		if v == nil {
			return nil
		}
		return v.InquiryDetails
	}).(WorkloadInquiryDetailsResponseArrayOutput)
}

func (o InquiryInfoResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InquiryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type InquiryValidation struct {
	Status *string `pulumi:"status"`
}





type InquiryValidationInput interface {
	pulumi.Input

	ToInquiryValidationOutput() InquiryValidationOutput
	ToInquiryValidationOutputWithContext(context.Context) InquiryValidationOutput
}

type InquiryValidationArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (InquiryValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InquiryValidation)(nil)).Elem()
}

func (i InquiryValidationArgs) ToInquiryValidationOutput() InquiryValidationOutput {
	return i.ToInquiryValidationOutputWithContext(context.Background())
}

func (i InquiryValidationArgs) ToInquiryValidationOutputWithContext(ctx context.Context) InquiryValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryValidationOutput)
}

func (i InquiryValidationArgs) ToInquiryValidationPtrOutput() InquiryValidationPtrOutput {
	return i.ToInquiryValidationPtrOutputWithContext(context.Background())
}

func (i InquiryValidationArgs) ToInquiryValidationPtrOutputWithContext(ctx context.Context) InquiryValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryValidationOutput).ToInquiryValidationPtrOutputWithContext(ctx)
}









type InquiryValidationPtrInput interface {
	pulumi.Input

	ToInquiryValidationPtrOutput() InquiryValidationPtrOutput
	ToInquiryValidationPtrOutputWithContext(context.Context) InquiryValidationPtrOutput
}

type inquiryValidationPtrType InquiryValidationArgs

func InquiryValidationPtr(v *InquiryValidationArgs) InquiryValidationPtrInput {
	return (*inquiryValidationPtrType)(v)
}

func (*inquiryValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InquiryValidation)(nil)).Elem()
}

func (i *inquiryValidationPtrType) ToInquiryValidationPtrOutput() InquiryValidationPtrOutput {
	return i.ToInquiryValidationPtrOutputWithContext(context.Background())
}

func (i *inquiryValidationPtrType) ToInquiryValidationPtrOutputWithContext(ctx context.Context) InquiryValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryValidationPtrOutput)
}

type InquiryValidationOutput struct{ *pulumi.OutputState }

func (InquiryValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InquiryValidation)(nil)).Elem()
}

func (o InquiryValidationOutput) ToInquiryValidationOutput() InquiryValidationOutput {
	return o
}

func (o InquiryValidationOutput) ToInquiryValidationOutputWithContext(ctx context.Context) InquiryValidationOutput {
	return o
}

func (o InquiryValidationOutput) ToInquiryValidationPtrOutput() InquiryValidationPtrOutput {
	return o.ToInquiryValidationPtrOutputWithContext(context.Background())
}

func (o InquiryValidationOutput) ToInquiryValidationPtrOutputWithContext(ctx context.Context) InquiryValidationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InquiryValidation) *InquiryValidation {
		return &v
	}).(InquiryValidationPtrOutput)
}

func (o InquiryValidationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InquiryValidation) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type InquiryValidationPtrOutput struct{ *pulumi.OutputState }

func (InquiryValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InquiryValidation)(nil)).Elem()
}

func (o InquiryValidationPtrOutput) ToInquiryValidationPtrOutput() InquiryValidationPtrOutput {
	return o
}

func (o InquiryValidationPtrOutput) ToInquiryValidationPtrOutputWithContext(ctx context.Context) InquiryValidationPtrOutput {
	return o
}

func (o InquiryValidationPtrOutput) Elem() InquiryValidationOutput {
	return o.ApplyT(func(v *InquiryValidation) InquiryValidation {
		if v != nil {
			return *v
		}
		var ret InquiryValidation
		return ret
	}).(InquiryValidationOutput)
}

func (o InquiryValidationPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InquiryValidation) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type InquiryValidationResponse struct {
	AdditionalDetail string               `pulumi:"additionalDetail"`
	ErrorDetail      *ErrorDetailResponse `pulumi:"errorDetail"`
	Status           *string              `pulumi:"status"`
}





type InquiryValidationResponseInput interface {
	pulumi.Input

	ToInquiryValidationResponseOutput() InquiryValidationResponseOutput
	ToInquiryValidationResponseOutputWithContext(context.Context) InquiryValidationResponseOutput
}

type InquiryValidationResponseArgs struct {
	AdditionalDetail pulumi.StringInput          `pulumi:"additionalDetail"`
	ErrorDetail      ErrorDetailResponsePtrInput `pulumi:"errorDetail"`
	Status           pulumi.StringPtrInput       `pulumi:"status"`
}

func (InquiryValidationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InquiryValidationResponse)(nil)).Elem()
}

func (i InquiryValidationResponseArgs) ToInquiryValidationResponseOutput() InquiryValidationResponseOutput {
	return i.ToInquiryValidationResponseOutputWithContext(context.Background())
}

func (i InquiryValidationResponseArgs) ToInquiryValidationResponseOutputWithContext(ctx context.Context) InquiryValidationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryValidationResponseOutput)
}

func (i InquiryValidationResponseArgs) ToInquiryValidationResponsePtrOutput() InquiryValidationResponsePtrOutput {
	return i.ToInquiryValidationResponsePtrOutputWithContext(context.Background())
}

func (i InquiryValidationResponseArgs) ToInquiryValidationResponsePtrOutputWithContext(ctx context.Context) InquiryValidationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryValidationResponseOutput).ToInquiryValidationResponsePtrOutputWithContext(ctx)
}









type InquiryValidationResponsePtrInput interface {
	pulumi.Input

	ToInquiryValidationResponsePtrOutput() InquiryValidationResponsePtrOutput
	ToInquiryValidationResponsePtrOutputWithContext(context.Context) InquiryValidationResponsePtrOutput
}

type inquiryValidationResponsePtrType InquiryValidationResponseArgs

func InquiryValidationResponsePtr(v *InquiryValidationResponseArgs) InquiryValidationResponsePtrInput {
	return (*inquiryValidationResponsePtrType)(v)
}

func (*inquiryValidationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InquiryValidationResponse)(nil)).Elem()
}

func (i *inquiryValidationResponsePtrType) ToInquiryValidationResponsePtrOutput() InquiryValidationResponsePtrOutput {
	return i.ToInquiryValidationResponsePtrOutputWithContext(context.Background())
}

func (i *inquiryValidationResponsePtrType) ToInquiryValidationResponsePtrOutputWithContext(ctx context.Context) InquiryValidationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InquiryValidationResponsePtrOutput)
}

type InquiryValidationResponseOutput struct{ *pulumi.OutputState }

func (InquiryValidationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InquiryValidationResponse)(nil)).Elem()
}

func (o InquiryValidationResponseOutput) ToInquiryValidationResponseOutput() InquiryValidationResponseOutput {
	return o
}

func (o InquiryValidationResponseOutput) ToInquiryValidationResponseOutputWithContext(ctx context.Context) InquiryValidationResponseOutput {
	return o
}

func (o InquiryValidationResponseOutput) ToInquiryValidationResponsePtrOutput() InquiryValidationResponsePtrOutput {
	return o.ToInquiryValidationResponsePtrOutputWithContext(context.Background())
}

func (o InquiryValidationResponseOutput) ToInquiryValidationResponsePtrOutputWithContext(ctx context.Context) InquiryValidationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InquiryValidationResponse) *InquiryValidationResponse {
		return &v
	}).(InquiryValidationResponsePtrOutput)
}

func (o InquiryValidationResponseOutput) AdditionalDetail() pulumi.StringOutput {
	return o.ApplyT(func(v InquiryValidationResponse) string { return v.AdditionalDetail }).(pulumi.StringOutput)
}

func (o InquiryValidationResponseOutput) ErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v InquiryValidationResponse) *ErrorDetailResponse { return v.ErrorDetail }).(ErrorDetailResponsePtrOutput)
}

func (o InquiryValidationResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InquiryValidationResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type InquiryValidationResponsePtrOutput struct{ *pulumi.OutputState }

func (InquiryValidationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InquiryValidationResponse)(nil)).Elem()
}

func (o InquiryValidationResponsePtrOutput) ToInquiryValidationResponsePtrOutput() InquiryValidationResponsePtrOutput {
	return o
}

func (o InquiryValidationResponsePtrOutput) ToInquiryValidationResponsePtrOutputWithContext(ctx context.Context) InquiryValidationResponsePtrOutput {
	return o
}

func (o InquiryValidationResponsePtrOutput) Elem() InquiryValidationResponseOutput {
	return o.ApplyT(func(v *InquiryValidationResponse) InquiryValidationResponse {
		if v != nil {
			return *v
		}
		var ret InquiryValidationResponse
		return ret
	}).(InquiryValidationResponseOutput)
}

func (o InquiryValidationResponsePtrOutput) AdditionalDetail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InquiryValidationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdditionalDetail
	}).(pulumi.StringPtrOutput)
}

func (o InquiryValidationResponsePtrOutput) ErrorDetail() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v *InquiryValidationResponse) *ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.ErrorDetail
	}).(ErrorDetailResponsePtrOutput)
}

func (o InquiryValidationResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InquiryValidationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type InstantRPAdditionalDetails struct {
	AzureBackupRGNamePrefix *string `pulumi:"azureBackupRGNamePrefix"`
	AzureBackupRGNameSuffix *string `pulumi:"azureBackupRGNameSuffix"`
}





type InstantRPAdditionalDetailsInput interface {
	pulumi.Input

	ToInstantRPAdditionalDetailsOutput() InstantRPAdditionalDetailsOutput
	ToInstantRPAdditionalDetailsOutputWithContext(context.Context) InstantRPAdditionalDetailsOutput
}

type InstantRPAdditionalDetailsArgs struct {
	AzureBackupRGNamePrefix pulumi.StringPtrInput `pulumi:"azureBackupRGNamePrefix"`
	AzureBackupRGNameSuffix pulumi.StringPtrInput `pulumi:"azureBackupRGNameSuffix"`
}

func (InstantRPAdditionalDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstantRPAdditionalDetails)(nil)).Elem()
}

func (i InstantRPAdditionalDetailsArgs) ToInstantRPAdditionalDetailsOutput() InstantRPAdditionalDetailsOutput {
	return i.ToInstantRPAdditionalDetailsOutputWithContext(context.Background())
}

func (i InstantRPAdditionalDetailsArgs) ToInstantRPAdditionalDetailsOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantRPAdditionalDetailsOutput)
}

func (i InstantRPAdditionalDetailsArgs) ToInstantRPAdditionalDetailsPtrOutput() InstantRPAdditionalDetailsPtrOutput {
	return i.ToInstantRPAdditionalDetailsPtrOutputWithContext(context.Background())
}

func (i InstantRPAdditionalDetailsArgs) ToInstantRPAdditionalDetailsPtrOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantRPAdditionalDetailsOutput).ToInstantRPAdditionalDetailsPtrOutputWithContext(ctx)
}









type InstantRPAdditionalDetailsPtrInput interface {
	pulumi.Input

	ToInstantRPAdditionalDetailsPtrOutput() InstantRPAdditionalDetailsPtrOutput
	ToInstantRPAdditionalDetailsPtrOutputWithContext(context.Context) InstantRPAdditionalDetailsPtrOutput
}

type instantRPAdditionalDetailsPtrType InstantRPAdditionalDetailsArgs

func InstantRPAdditionalDetailsPtr(v *InstantRPAdditionalDetailsArgs) InstantRPAdditionalDetailsPtrInput {
	return (*instantRPAdditionalDetailsPtrType)(v)
}

func (*instantRPAdditionalDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstantRPAdditionalDetails)(nil)).Elem()
}

func (i *instantRPAdditionalDetailsPtrType) ToInstantRPAdditionalDetailsPtrOutput() InstantRPAdditionalDetailsPtrOutput {
	return i.ToInstantRPAdditionalDetailsPtrOutputWithContext(context.Background())
}

func (i *instantRPAdditionalDetailsPtrType) ToInstantRPAdditionalDetailsPtrOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantRPAdditionalDetailsPtrOutput)
}

type InstantRPAdditionalDetailsOutput struct{ *pulumi.OutputState }

func (InstantRPAdditionalDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstantRPAdditionalDetails)(nil)).Elem()
}

func (o InstantRPAdditionalDetailsOutput) ToInstantRPAdditionalDetailsOutput() InstantRPAdditionalDetailsOutput {
	return o
}

func (o InstantRPAdditionalDetailsOutput) ToInstantRPAdditionalDetailsOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsOutput {
	return o
}

func (o InstantRPAdditionalDetailsOutput) ToInstantRPAdditionalDetailsPtrOutput() InstantRPAdditionalDetailsPtrOutput {
	return o.ToInstantRPAdditionalDetailsPtrOutputWithContext(context.Background())
}

func (o InstantRPAdditionalDetailsOutput) ToInstantRPAdditionalDetailsPtrOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstantRPAdditionalDetails) *InstantRPAdditionalDetails {
		return &v
	}).(InstantRPAdditionalDetailsPtrOutput)
}

func (o InstantRPAdditionalDetailsOutput) AzureBackupRGNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstantRPAdditionalDetails) *string { return v.AzureBackupRGNamePrefix }).(pulumi.StringPtrOutput)
}

func (o InstantRPAdditionalDetailsOutput) AzureBackupRGNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstantRPAdditionalDetails) *string { return v.AzureBackupRGNameSuffix }).(pulumi.StringPtrOutput)
}

type InstantRPAdditionalDetailsPtrOutput struct{ *pulumi.OutputState }

func (InstantRPAdditionalDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstantRPAdditionalDetails)(nil)).Elem()
}

func (o InstantRPAdditionalDetailsPtrOutput) ToInstantRPAdditionalDetailsPtrOutput() InstantRPAdditionalDetailsPtrOutput {
	return o
}

func (o InstantRPAdditionalDetailsPtrOutput) ToInstantRPAdditionalDetailsPtrOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsPtrOutput {
	return o
}

func (o InstantRPAdditionalDetailsPtrOutput) Elem() InstantRPAdditionalDetailsOutput {
	return o.ApplyT(func(v *InstantRPAdditionalDetails) InstantRPAdditionalDetails {
		if v != nil {
			return *v
		}
		var ret InstantRPAdditionalDetails
		return ret
	}).(InstantRPAdditionalDetailsOutput)
}

func (o InstantRPAdditionalDetailsPtrOutput) AzureBackupRGNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstantRPAdditionalDetails) *string {
		if v == nil {
			return nil
		}
		return v.AzureBackupRGNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o InstantRPAdditionalDetailsPtrOutput) AzureBackupRGNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstantRPAdditionalDetails) *string {
		if v == nil {
			return nil
		}
		return v.AzureBackupRGNameSuffix
	}).(pulumi.StringPtrOutput)
}

type InstantRPAdditionalDetailsResponse struct {
	AzureBackupRGNamePrefix *string `pulumi:"azureBackupRGNamePrefix"`
	AzureBackupRGNameSuffix *string `pulumi:"azureBackupRGNameSuffix"`
}





type InstantRPAdditionalDetailsResponseInput interface {
	pulumi.Input

	ToInstantRPAdditionalDetailsResponseOutput() InstantRPAdditionalDetailsResponseOutput
	ToInstantRPAdditionalDetailsResponseOutputWithContext(context.Context) InstantRPAdditionalDetailsResponseOutput
}

type InstantRPAdditionalDetailsResponseArgs struct {
	AzureBackupRGNamePrefix pulumi.StringPtrInput `pulumi:"azureBackupRGNamePrefix"`
	AzureBackupRGNameSuffix pulumi.StringPtrInput `pulumi:"azureBackupRGNameSuffix"`
}

func (InstantRPAdditionalDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstantRPAdditionalDetailsResponse)(nil)).Elem()
}

func (i InstantRPAdditionalDetailsResponseArgs) ToInstantRPAdditionalDetailsResponseOutput() InstantRPAdditionalDetailsResponseOutput {
	return i.ToInstantRPAdditionalDetailsResponseOutputWithContext(context.Background())
}

func (i InstantRPAdditionalDetailsResponseArgs) ToInstantRPAdditionalDetailsResponseOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantRPAdditionalDetailsResponseOutput)
}

func (i InstantRPAdditionalDetailsResponseArgs) ToInstantRPAdditionalDetailsResponsePtrOutput() InstantRPAdditionalDetailsResponsePtrOutput {
	return i.ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(context.Background())
}

func (i InstantRPAdditionalDetailsResponseArgs) ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantRPAdditionalDetailsResponseOutput).ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(ctx)
}









type InstantRPAdditionalDetailsResponsePtrInput interface {
	pulumi.Input

	ToInstantRPAdditionalDetailsResponsePtrOutput() InstantRPAdditionalDetailsResponsePtrOutput
	ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(context.Context) InstantRPAdditionalDetailsResponsePtrOutput
}

type instantRPAdditionalDetailsResponsePtrType InstantRPAdditionalDetailsResponseArgs

func InstantRPAdditionalDetailsResponsePtr(v *InstantRPAdditionalDetailsResponseArgs) InstantRPAdditionalDetailsResponsePtrInput {
	return (*instantRPAdditionalDetailsResponsePtrType)(v)
}

func (*instantRPAdditionalDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstantRPAdditionalDetailsResponse)(nil)).Elem()
}

func (i *instantRPAdditionalDetailsResponsePtrType) ToInstantRPAdditionalDetailsResponsePtrOutput() InstantRPAdditionalDetailsResponsePtrOutput {
	return i.ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *instantRPAdditionalDetailsResponsePtrType) ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstantRPAdditionalDetailsResponsePtrOutput)
}

type InstantRPAdditionalDetailsResponseOutput struct{ *pulumi.OutputState }

func (InstantRPAdditionalDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstantRPAdditionalDetailsResponse)(nil)).Elem()
}

func (o InstantRPAdditionalDetailsResponseOutput) ToInstantRPAdditionalDetailsResponseOutput() InstantRPAdditionalDetailsResponseOutput {
	return o
}

func (o InstantRPAdditionalDetailsResponseOutput) ToInstantRPAdditionalDetailsResponseOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsResponseOutput {
	return o
}

func (o InstantRPAdditionalDetailsResponseOutput) ToInstantRPAdditionalDetailsResponsePtrOutput() InstantRPAdditionalDetailsResponsePtrOutput {
	return o.ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(context.Background())
}

func (o InstantRPAdditionalDetailsResponseOutput) ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstantRPAdditionalDetailsResponse) *InstantRPAdditionalDetailsResponse {
		return &v
	}).(InstantRPAdditionalDetailsResponsePtrOutput)
}

func (o InstantRPAdditionalDetailsResponseOutput) AzureBackupRGNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstantRPAdditionalDetailsResponse) *string { return v.AzureBackupRGNamePrefix }).(pulumi.StringPtrOutput)
}

func (o InstantRPAdditionalDetailsResponseOutput) AzureBackupRGNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstantRPAdditionalDetailsResponse) *string { return v.AzureBackupRGNameSuffix }).(pulumi.StringPtrOutput)
}

type InstantRPAdditionalDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (InstantRPAdditionalDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstantRPAdditionalDetailsResponse)(nil)).Elem()
}

func (o InstantRPAdditionalDetailsResponsePtrOutput) ToInstantRPAdditionalDetailsResponsePtrOutput() InstantRPAdditionalDetailsResponsePtrOutput {
	return o
}

func (o InstantRPAdditionalDetailsResponsePtrOutput) ToInstantRPAdditionalDetailsResponsePtrOutputWithContext(ctx context.Context) InstantRPAdditionalDetailsResponsePtrOutput {
	return o
}

func (o InstantRPAdditionalDetailsResponsePtrOutput) Elem() InstantRPAdditionalDetailsResponseOutput {
	return o.ApplyT(func(v *InstantRPAdditionalDetailsResponse) InstantRPAdditionalDetailsResponse {
		if v != nil {
			return *v
		}
		var ret InstantRPAdditionalDetailsResponse
		return ret
	}).(InstantRPAdditionalDetailsResponseOutput)
}

func (o InstantRPAdditionalDetailsResponsePtrOutput) AzureBackupRGNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstantRPAdditionalDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureBackupRGNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o InstantRPAdditionalDetailsResponsePtrOutput) AzureBackupRGNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstantRPAdditionalDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureBackupRGNameSuffix
	}).(pulumi.StringPtrOutput)
}

type KPIResourceHealthDetails struct {
	ResourceHealthStatus *string `pulumi:"resourceHealthStatus"`
}





type KPIResourceHealthDetailsInput interface {
	pulumi.Input

	ToKPIResourceHealthDetailsOutput() KPIResourceHealthDetailsOutput
	ToKPIResourceHealthDetailsOutputWithContext(context.Context) KPIResourceHealthDetailsOutput
}

type KPIResourceHealthDetailsArgs struct {
	ResourceHealthStatus pulumi.StringPtrInput `pulumi:"resourceHealthStatus"`
}

func (KPIResourceHealthDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KPIResourceHealthDetails)(nil)).Elem()
}

func (i KPIResourceHealthDetailsArgs) ToKPIResourceHealthDetailsOutput() KPIResourceHealthDetailsOutput {
	return i.ToKPIResourceHealthDetailsOutputWithContext(context.Background())
}

func (i KPIResourceHealthDetailsArgs) ToKPIResourceHealthDetailsOutputWithContext(ctx context.Context) KPIResourceHealthDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KPIResourceHealthDetailsOutput)
}





type KPIResourceHealthDetailsMapInput interface {
	pulumi.Input

	ToKPIResourceHealthDetailsMapOutput() KPIResourceHealthDetailsMapOutput
	ToKPIResourceHealthDetailsMapOutputWithContext(context.Context) KPIResourceHealthDetailsMapOutput
}

type KPIResourceHealthDetailsMap map[string]KPIResourceHealthDetailsInput

func (KPIResourceHealthDetailsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KPIResourceHealthDetails)(nil)).Elem()
}

func (i KPIResourceHealthDetailsMap) ToKPIResourceHealthDetailsMapOutput() KPIResourceHealthDetailsMapOutput {
	return i.ToKPIResourceHealthDetailsMapOutputWithContext(context.Background())
}

func (i KPIResourceHealthDetailsMap) ToKPIResourceHealthDetailsMapOutputWithContext(ctx context.Context) KPIResourceHealthDetailsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KPIResourceHealthDetailsMapOutput)
}

type KPIResourceHealthDetailsOutput struct{ *pulumi.OutputState }

func (KPIResourceHealthDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KPIResourceHealthDetails)(nil)).Elem()
}

func (o KPIResourceHealthDetailsOutput) ToKPIResourceHealthDetailsOutput() KPIResourceHealthDetailsOutput {
	return o
}

func (o KPIResourceHealthDetailsOutput) ToKPIResourceHealthDetailsOutputWithContext(ctx context.Context) KPIResourceHealthDetailsOutput {
	return o
}

func (o KPIResourceHealthDetailsOutput) ResourceHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KPIResourceHealthDetails) *string { return v.ResourceHealthStatus }).(pulumi.StringPtrOutput)
}

type KPIResourceHealthDetailsMapOutput struct{ *pulumi.OutputState }

func (KPIResourceHealthDetailsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KPIResourceHealthDetails)(nil)).Elem()
}

func (o KPIResourceHealthDetailsMapOutput) ToKPIResourceHealthDetailsMapOutput() KPIResourceHealthDetailsMapOutput {
	return o
}

func (o KPIResourceHealthDetailsMapOutput) ToKPIResourceHealthDetailsMapOutputWithContext(ctx context.Context) KPIResourceHealthDetailsMapOutput {
	return o
}

func (o KPIResourceHealthDetailsMapOutput) MapIndex(k pulumi.StringInput) KPIResourceHealthDetailsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KPIResourceHealthDetails {
		return vs[0].(map[string]KPIResourceHealthDetails)[vs[1].(string)]
	}).(KPIResourceHealthDetailsOutput)
}

type KPIResourceHealthDetailsResponse struct {
	ResourceHealthDetails []ResourceHealthDetailsResponse `pulumi:"resourceHealthDetails"`
	ResourceHealthStatus  *string                         `pulumi:"resourceHealthStatus"`
}





type KPIResourceHealthDetailsResponseInput interface {
	pulumi.Input

	ToKPIResourceHealthDetailsResponseOutput() KPIResourceHealthDetailsResponseOutput
	ToKPIResourceHealthDetailsResponseOutputWithContext(context.Context) KPIResourceHealthDetailsResponseOutput
}

type KPIResourceHealthDetailsResponseArgs struct {
	ResourceHealthDetails ResourceHealthDetailsResponseArrayInput `pulumi:"resourceHealthDetails"`
	ResourceHealthStatus  pulumi.StringPtrInput                   `pulumi:"resourceHealthStatus"`
}

func (KPIResourceHealthDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KPIResourceHealthDetailsResponse)(nil)).Elem()
}

func (i KPIResourceHealthDetailsResponseArgs) ToKPIResourceHealthDetailsResponseOutput() KPIResourceHealthDetailsResponseOutput {
	return i.ToKPIResourceHealthDetailsResponseOutputWithContext(context.Background())
}

func (i KPIResourceHealthDetailsResponseArgs) ToKPIResourceHealthDetailsResponseOutputWithContext(ctx context.Context) KPIResourceHealthDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KPIResourceHealthDetailsResponseOutput)
}





type KPIResourceHealthDetailsResponseMapInput interface {
	pulumi.Input

	ToKPIResourceHealthDetailsResponseMapOutput() KPIResourceHealthDetailsResponseMapOutput
	ToKPIResourceHealthDetailsResponseMapOutputWithContext(context.Context) KPIResourceHealthDetailsResponseMapOutput
}

type KPIResourceHealthDetailsResponseMap map[string]KPIResourceHealthDetailsResponseInput

func (KPIResourceHealthDetailsResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KPIResourceHealthDetailsResponse)(nil)).Elem()
}

func (i KPIResourceHealthDetailsResponseMap) ToKPIResourceHealthDetailsResponseMapOutput() KPIResourceHealthDetailsResponseMapOutput {
	return i.ToKPIResourceHealthDetailsResponseMapOutputWithContext(context.Background())
}

func (i KPIResourceHealthDetailsResponseMap) ToKPIResourceHealthDetailsResponseMapOutputWithContext(ctx context.Context) KPIResourceHealthDetailsResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KPIResourceHealthDetailsResponseMapOutput)
}

type KPIResourceHealthDetailsResponseOutput struct{ *pulumi.OutputState }

func (KPIResourceHealthDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KPIResourceHealthDetailsResponse)(nil)).Elem()
}

func (o KPIResourceHealthDetailsResponseOutput) ToKPIResourceHealthDetailsResponseOutput() KPIResourceHealthDetailsResponseOutput {
	return o
}

func (o KPIResourceHealthDetailsResponseOutput) ToKPIResourceHealthDetailsResponseOutputWithContext(ctx context.Context) KPIResourceHealthDetailsResponseOutput {
	return o
}

func (o KPIResourceHealthDetailsResponseOutput) ResourceHealthDetails() ResourceHealthDetailsResponseArrayOutput {
	return o.ApplyT(func(v KPIResourceHealthDetailsResponse) []ResourceHealthDetailsResponse {
		return v.ResourceHealthDetails
	}).(ResourceHealthDetailsResponseArrayOutput)
}

func (o KPIResourceHealthDetailsResponseOutput) ResourceHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KPIResourceHealthDetailsResponse) *string { return v.ResourceHealthStatus }).(pulumi.StringPtrOutput)
}

type KPIResourceHealthDetailsResponseMapOutput struct{ *pulumi.OutputState }

func (KPIResourceHealthDetailsResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KPIResourceHealthDetailsResponse)(nil)).Elem()
}

func (o KPIResourceHealthDetailsResponseMapOutput) ToKPIResourceHealthDetailsResponseMapOutput() KPIResourceHealthDetailsResponseMapOutput {
	return o
}

func (o KPIResourceHealthDetailsResponseMapOutput) ToKPIResourceHealthDetailsResponseMapOutputWithContext(ctx context.Context) KPIResourceHealthDetailsResponseMapOutput {
	return o
}

func (o KPIResourceHealthDetailsResponseMapOutput) MapIndex(k pulumi.StringInput) KPIResourceHealthDetailsResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KPIResourceHealthDetailsResponse {
		return vs[0].(map[string]KPIResourceHealthDetailsResponse)[vs[1].(string)]
	}).(KPIResourceHealthDetailsResponseOutput)
}

type LogSchedulePolicy struct {
	ScheduleFrequencyInMins *int   `pulumi:"scheduleFrequencyInMins"`
	SchedulePolicyType      string `pulumi:"schedulePolicyType"`
}





type LogSchedulePolicyInput interface {
	pulumi.Input

	ToLogSchedulePolicyOutput() LogSchedulePolicyOutput
	ToLogSchedulePolicyOutputWithContext(context.Context) LogSchedulePolicyOutput
}

type LogSchedulePolicyArgs struct {
	ScheduleFrequencyInMins pulumi.IntPtrInput `pulumi:"scheduleFrequencyInMins"`
	SchedulePolicyType      pulumi.StringInput `pulumi:"schedulePolicyType"`
}

func (LogSchedulePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSchedulePolicy)(nil)).Elem()
}

func (i LogSchedulePolicyArgs) ToLogSchedulePolicyOutput() LogSchedulePolicyOutput {
	return i.ToLogSchedulePolicyOutputWithContext(context.Background())
}

func (i LogSchedulePolicyArgs) ToLogSchedulePolicyOutputWithContext(ctx context.Context) LogSchedulePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSchedulePolicyOutput)
}

type LogSchedulePolicyOutput struct{ *pulumi.OutputState }

func (LogSchedulePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSchedulePolicy)(nil)).Elem()
}

func (o LogSchedulePolicyOutput) ToLogSchedulePolicyOutput() LogSchedulePolicyOutput {
	return o
}

func (o LogSchedulePolicyOutput) ToLogSchedulePolicyOutputWithContext(ctx context.Context) LogSchedulePolicyOutput {
	return o
}

func (o LogSchedulePolicyOutput) ScheduleFrequencyInMins() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LogSchedulePolicy) *int { return v.ScheduleFrequencyInMins }).(pulumi.IntPtrOutput)
}

func (o LogSchedulePolicyOutput) SchedulePolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v LogSchedulePolicy) string { return v.SchedulePolicyType }).(pulumi.StringOutput)
}

type LogSchedulePolicyResponse struct {
	ScheduleFrequencyInMins *int   `pulumi:"scheduleFrequencyInMins"`
	SchedulePolicyType      string `pulumi:"schedulePolicyType"`
}





type LogSchedulePolicyResponseInput interface {
	pulumi.Input

	ToLogSchedulePolicyResponseOutput() LogSchedulePolicyResponseOutput
	ToLogSchedulePolicyResponseOutputWithContext(context.Context) LogSchedulePolicyResponseOutput
}

type LogSchedulePolicyResponseArgs struct {
	ScheduleFrequencyInMins pulumi.IntPtrInput `pulumi:"scheduleFrequencyInMins"`
	SchedulePolicyType      pulumi.StringInput `pulumi:"schedulePolicyType"`
}

func (LogSchedulePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSchedulePolicyResponse)(nil)).Elem()
}

func (i LogSchedulePolicyResponseArgs) ToLogSchedulePolicyResponseOutput() LogSchedulePolicyResponseOutput {
	return i.ToLogSchedulePolicyResponseOutputWithContext(context.Background())
}

func (i LogSchedulePolicyResponseArgs) ToLogSchedulePolicyResponseOutputWithContext(ctx context.Context) LogSchedulePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSchedulePolicyResponseOutput)
}

type LogSchedulePolicyResponseOutput struct{ *pulumi.OutputState }

func (LogSchedulePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSchedulePolicyResponse)(nil)).Elem()
}

func (o LogSchedulePolicyResponseOutput) ToLogSchedulePolicyResponseOutput() LogSchedulePolicyResponseOutput {
	return o
}

func (o LogSchedulePolicyResponseOutput) ToLogSchedulePolicyResponseOutputWithContext(ctx context.Context) LogSchedulePolicyResponseOutput {
	return o
}

func (o LogSchedulePolicyResponseOutput) ScheduleFrequencyInMins() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LogSchedulePolicyResponse) *int { return v.ScheduleFrequencyInMins }).(pulumi.IntPtrOutput)
}

func (o LogSchedulePolicyResponseOutput) SchedulePolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v LogSchedulePolicyResponse) string { return v.SchedulePolicyType }).(pulumi.StringOutput)
}

type LongTermRetentionPolicy struct {
	DailySchedule       *DailyRetentionSchedule   `pulumi:"dailySchedule"`
	MonthlySchedule     *MonthlyRetentionSchedule `pulumi:"monthlySchedule"`
	RetentionPolicyType string                    `pulumi:"retentionPolicyType"`
	WeeklySchedule      *WeeklyRetentionSchedule  `pulumi:"weeklySchedule"`
	YearlySchedule      *YearlyRetentionSchedule  `pulumi:"yearlySchedule"`
}





type LongTermRetentionPolicyInput interface {
	pulumi.Input

	ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput
	ToLongTermRetentionPolicyOutputWithContext(context.Context) LongTermRetentionPolicyOutput
}

type LongTermRetentionPolicyArgs struct {
	DailySchedule       DailyRetentionSchedulePtrInput   `pulumi:"dailySchedule"`
	MonthlySchedule     MonthlyRetentionSchedulePtrInput `pulumi:"monthlySchedule"`
	RetentionPolicyType pulumi.StringInput               `pulumi:"retentionPolicyType"`
	WeeklySchedule      WeeklyRetentionSchedulePtrInput  `pulumi:"weeklySchedule"`
	YearlySchedule      YearlyRetentionSchedulePtrInput  `pulumi:"yearlySchedule"`
}

func (LongTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicy)(nil)).Elem()
}

func (i LongTermRetentionPolicyArgs) ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput {
	return i.ToLongTermRetentionPolicyOutputWithContext(context.Background())
}

func (i LongTermRetentionPolicyArgs) ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermRetentionPolicyOutput)
}

type LongTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (LongTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicy)(nil)).Elem()
}

func (o LongTermRetentionPolicyOutput) ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput {
	return o
}

func (o LongTermRetentionPolicyOutput) ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput {
	return o
}

func (o LongTermRetentionPolicyOutput) DailySchedule() DailyRetentionSchedulePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *DailyRetentionSchedule { return v.DailySchedule }).(DailyRetentionSchedulePtrOutput)
}

func (o LongTermRetentionPolicyOutput) MonthlySchedule() MonthlyRetentionSchedulePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *MonthlyRetentionSchedule { return v.MonthlySchedule }).(MonthlyRetentionSchedulePtrOutput)
}

func (o LongTermRetentionPolicyOutput) RetentionPolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) string { return v.RetentionPolicyType }).(pulumi.StringOutput)
}

func (o LongTermRetentionPolicyOutput) WeeklySchedule() WeeklyRetentionSchedulePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *WeeklyRetentionSchedule { return v.WeeklySchedule }).(WeeklyRetentionSchedulePtrOutput)
}

func (o LongTermRetentionPolicyOutput) YearlySchedule() YearlyRetentionSchedulePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *YearlyRetentionSchedule { return v.YearlySchedule }).(YearlyRetentionSchedulePtrOutput)
}

type LongTermRetentionPolicyResponse struct {
	DailySchedule       *DailyRetentionScheduleResponse   `pulumi:"dailySchedule"`
	MonthlySchedule     *MonthlyRetentionScheduleResponse `pulumi:"monthlySchedule"`
	RetentionPolicyType string                            `pulumi:"retentionPolicyType"`
	WeeklySchedule      *WeeklyRetentionScheduleResponse  `pulumi:"weeklySchedule"`
	YearlySchedule      *YearlyRetentionScheduleResponse  `pulumi:"yearlySchedule"`
}





type LongTermRetentionPolicyResponseInput interface {
	pulumi.Input

	ToLongTermRetentionPolicyResponseOutput() LongTermRetentionPolicyResponseOutput
	ToLongTermRetentionPolicyResponseOutputWithContext(context.Context) LongTermRetentionPolicyResponseOutput
}

type LongTermRetentionPolicyResponseArgs struct {
	DailySchedule       DailyRetentionScheduleResponsePtrInput   `pulumi:"dailySchedule"`
	MonthlySchedule     MonthlyRetentionScheduleResponsePtrInput `pulumi:"monthlySchedule"`
	RetentionPolicyType pulumi.StringInput                       `pulumi:"retentionPolicyType"`
	WeeklySchedule      WeeklyRetentionScheduleResponsePtrInput  `pulumi:"weeklySchedule"`
	YearlySchedule      YearlyRetentionScheduleResponsePtrInput  `pulumi:"yearlySchedule"`
}

func (LongTermRetentionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicyResponse)(nil)).Elem()
}

func (i LongTermRetentionPolicyResponseArgs) ToLongTermRetentionPolicyResponseOutput() LongTermRetentionPolicyResponseOutput {
	return i.ToLongTermRetentionPolicyResponseOutputWithContext(context.Background())
}

func (i LongTermRetentionPolicyResponseArgs) ToLongTermRetentionPolicyResponseOutputWithContext(ctx context.Context) LongTermRetentionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermRetentionPolicyResponseOutput)
}

type LongTermRetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (LongTermRetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicyResponse)(nil)).Elem()
}

func (o LongTermRetentionPolicyResponseOutput) ToLongTermRetentionPolicyResponseOutput() LongTermRetentionPolicyResponseOutput {
	return o
}

func (o LongTermRetentionPolicyResponseOutput) ToLongTermRetentionPolicyResponseOutputWithContext(ctx context.Context) LongTermRetentionPolicyResponseOutput {
	return o
}

func (o LongTermRetentionPolicyResponseOutput) DailySchedule() DailyRetentionScheduleResponsePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *DailyRetentionScheduleResponse { return v.DailySchedule }).(DailyRetentionScheduleResponsePtrOutput)
}

func (o LongTermRetentionPolicyResponseOutput) MonthlySchedule() MonthlyRetentionScheduleResponsePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *MonthlyRetentionScheduleResponse { return v.MonthlySchedule }).(MonthlyRetentionScheduleResponsePtrOutput)
}

func (o LongTermRetentionPolicyResponseOutput) RetentionPolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) string { return v.RetentionPolicyType }).(pulumi.StringOutput)
}

func (o LongTermRetentionPolicyResponseOutput) WeeklySchedule() WeeklyRetentionScheduleResponsePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *WeeklyRetentionScheduleResponse { return v.WeeklySchedule }).(WeeklyRetentionScheduleResponsePtrOutput)
}

func (o LongTermRetentionPolicyResponseOutput) YearlySchedule() YearlyRetentionScheduleResponsePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *YearlyRetentionScheduleResponse { return v.YearlySchedule }).(YearlyRetentionScheduleResponsePtrOutput)
}

type LongTermSchedulePolicy struct {
	SchedulePolicyType string `pulumi:"schedulePolicyType"`
}





type LongTermSchedulePolicyInput interface {
	pulumi.Input

	ToLongTermSchedulePolicyOutput() LongTermSchedulePolicyOutput
	ToLongTermSchedulePolicyOutputWithContext(context.Context) LongTermSchedulePolicyOutput
}

type LongTermSchedulePolicyArgs struct {
	SchedulePolicyType pulumi.StringInput `pulumi:"schedulePolicyType"`
}

func (LongTermSchedulePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermSchedulePolicy)(nil)).Elem()
}

func (i LongTermSchedulePolicyArgs) ToLongTermSchedulePolicyOutput() LongTermSchedulePolicyOutput {
	return i.ToLongTermSchedulePolicyOutputWithContext(context.Background())
}

func (i LongTermSchedulePolicyArgs) ToLongTermSchedulePolicyOutputWithContext(ctx context.Context) LongTermSchedulePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermSchedulePolicyOutput)
}

type LongTermSchedulePolicyOutput struct{ *pulumi.OutputState }

func (LongTermSchedulePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermSchedulePolicy)(nil)).Elem()
}

func (o LongTermSchedulePolicyOutput) ToLongTermSchedulePolicyOutput() LongTermSchedulePolicyOutput {
	return o
}

func (o LongTermSchedulePolicyOutput) ToLongTermSchedulePolicyOutputWithContext(ctx context.Context) LongTermSchedulePolicyOutput {
	return o
}

func (o LongTermSchedulePolicyOutput) SchedulePolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v LongTermSchedulePolicy) string { return v.SchedulePolicyType }).(pulumi.StringOutput)
}

type LongTermSchedulePolicyResponse struct {
	SchedulePolicyType string `pulumi:"schedulePolicyType"`
}





type LongTermSchedulePolicyResponseInput interface {
	pulumi.Input

	ToLongTermSchedulePolicyResponseOutput() LongTermSchedulePolicyResponseOutput
	ToLongTermSchedulePolicyResponseOutputWithContext(context.Context) LongTermSchedulePolicyResponseOutput
}

type LongTermSchedulePolicyResponseArgs struct {
	SchedulePolicyType pulumi.StringInput `pulumi:"schedulePolicyType"`
}

func (LongTermSchedulePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermSchedulePolicyResponse)(nil)).Elem()
}

func (i LongTermSchedulePolicyResponseArgs) ToLongTermSchedulePolicyResponseOutput() LongTermSchedulePolicyResponseOutput {
	return i.ToLongTermSchedulePolicyResponseOutputWithContext(context.Background())
}

func (i LongTermSchedulePolicyResponseArgs) ToLongTermSchedulePolicyResponseOutputWithContext(ctx context.Context) LongTermSchedulePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermSchedulePolicyResponseOutput)
}

type LongTermSchedulePolicyResponseOutput struct{ *pulumi.OutputState }

func (LongTermSchedulePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermSchedulePolicyResponse)(nil)).Elem()
}

func (o LongTermSchedulePolicyResponseOutput) ToLongTermSchedulePolicyResponseOutput() LongTermSchedulePolicyResponseOutput {
	return o
}

func (o LongTermSchedulePolicyResponseOutput) ToLongTermSchedulePolicyResponseOutputWithContext(ctx context.Context) LongTermSchedulePolicyResponseOutput {
	return o
}

func (o LongTermSchedulePolicyResponseOutput) SchedulePolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v LongTermSchedulePolicyResponse) string { return v.SchedulePolicyType }).(pulumi.StringOutput)
}

type MABContainerHealthDetails struct {
	Code            *int     `pulumi:"code"`
	Message         *string  `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           *string  `pulumi:"title"`
}





type MABContainerHealthDetailsInput interface {
	pulumi.Input

	ToMABContainerHealthDetailsOutput() MABContainerHealthDetailsOutput
	ToMABContainerHealthDetailsOutputWithContext(context.Context) MABContainerHealthDetailsOutput
}

type MABContainerHealthDetailsArgs struct {
	Code            pulumi.IntPtrInput      `pulumi:"code"`
	Message         pulumi.StringPtrInput   `pulumi:"message"`
	Recommendations pulumi.StringArrayInput `pulumi:"recommendations"`
	Title           pulumi.StringPtrInput   `pulumi:"title"`
}

func (MABContainerHealthDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MABContainerHealthDetails)(nil)).Elem()
}

func (i MABContainerHealthDetailsArgs) ToMABContainerHealthDetailsOutput() MABContainerHealthDetailsOutput {
	return i.ToMABContainerHealthDetailsOutputWithContext(context.Background())
}

func (i MABContainerHealthDetailsArgs) ToMABContainerHealthDetailsOutputWithContext(ctx context.Context) MABContainerHealthDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MABContainerHealthDetailsOutput)
}





type MABContainerHealthDetailsArrayInput interface {
	pulumi.Input

	ToMABContainerHealthDetailsArrayOutput() MABContainerHealthDetailsArrayOutput
	ToMABContainerHealthDetailsArrayOutputWithContext(context.Context) MABContainerHealthDetailsArrayOutput
}

type MABContainerHealthDetailsArray []MABContainerHealthDetailsInput

func (MABContainerHealthDetailsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MABContainerHealthDetails)(nil)).Elem()
}

func (i MABContainerHealthDetailsArray) ToMABContainerHealthDetailsArrayOutput() MABContainerHealthDetailsArrayOutput {
	return i.ToMABContainerHealthDetailsArrayOutputWithContext(context.Background())
}

func (i MABContainerHealthDetailsArray) ToMABContainerHealthDetailsArrayOutputWithContext(ctx context.Context) MABContainerHealthDetailsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MABContainerHealthDetailsArrayOutput)
}

type MABContainerHealthDetailsOutput struct{ *pulumi.OutputState }

func (MABContainerHealthDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MABContainerHealthDetails)(nil)).Elem()
}

func (o MABContainerHealthDetailsOutput) ToMABContainerHealthDetailsOutput() MABContainerHealthDetailsOutput {
	return o
}

func (o MABContainerHealthDetailsOutput) ToMABContainerHealthDetailsOutputWithContext(ctx context.Context) MABContainerHealthDetailsOutput {
	return o
}

func (o MABContainerHealthDetailsOutput) Code() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MABContainerHealthDetails) *int { return v.Code }).(pulumi.IntPtrOutput)
}

func (o MABContainerHealthDetailsOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MABContainerHealthDetails) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o MABContainerHealthDetailsOutput) Recommendations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MABContainerHealthDetails) []string { return v.Recommendations }).(pulumi.StringArrayOutput)
}

func (o MABContainerHealthDetailsOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MABContainerHealthDetails) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type MABContainerHealthDetailsArrayOutput struct{ *pulumi.OutputState }

func (MABContainerHealthDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MABContainerHealthDetails)(nil)).Elem()
}

func (o MABContainerHealthDetailsArrayOutput) ToMABContainerHealthDetailsArrayOutput() MABContainerHealthDetailsArrayOutput {
	return o
}

func (o MABContainerHealthDetailsArrayOutput) ToMABContainerHealthDetailsArrayOutputWithContext(ctx context.Context) MABContainerHealthDetailsArrayOutput {
	return o
}

func (o MABContainerHealthDetailsArrayOutput) Index(i pulumi.IntInput) MABContainerHealthDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MABContainerHealthDetails {
		return vs[0].([]MABContainerHealthDetails)[vs[1].(int)]
	}).(MABContainerHealthDetailsOutput)
}

type MABContainerHealthDetailsResponse struct {
	Code            *int     `pulumi:"code"`
	Message         *string  `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           *string  `pulumi:"title"`
}





type MABContainerHealthDetailsResponseInput interface {
	pulumi.Input

	ToMABContainerHealthDetailsResponseOutput() MABContainerHealthDetailsResponseOutput
	ToMABContainerHealthDetailsResponseOutputWithContext(context.Context) MABContainerHealthDetailsResponseOutput
}

type MABContainerHealthDetailsResponseArgs struct {
	Code            pulumi.IntPtrInput      `pulumi:"code"`
	Message         pulumi.StringPtrInput   `pulumi:"message"`
	Recommendations pulumi.StringArrayInput `pulumi:"recommendations"`
	Title           pulumi.StringPtrInput   `pulumi:"title"`
}

func (MABContainerHealthDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MABContainerHealthDetailsResponse)(nil)).Elem()
}

func (i MABContainerHealthDetailsResponseArgs) ToMABContainerHealthDetailsResponseOutput() MABContainerHealthDetailsResponseOutput {
	return i.ToMABContainerHealthDetailsResponseOutputWithContext(context.Background())
}

func (i MABContainerHealthDetailsResponseArgs) ToMABContainerHealthDetailsResponseOutputWithContext(ctx context.Context) MABContainerHealthDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MABContainerHealthDetailsResponseOutput)
}





type MABContainerHealthDetailsResponseArrayInput interface {
	pulumi.Input

	ToMABContainerHealthDetailsResponseArrayOutput() MABContainerHealthDetailsResponseArrayOutput
	ToMABContainerHealthDetailsResponseArrayOutputWithContext(context.Context) MABContainerHealthDetailsResponseArrayOutput
}

type MABContainerHealthDetailsResponseArray []MABContainerHealthDetailsResponseInput

func (MABContainerHealthDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MABContainerHealthDetailsResponse)(nil)).Elem()
}

func (i MABContainerHealthDetailsResponseArray) ToMABContainerHealthDetailsResponseArrayOutput() MABContainerHealthDetailsResponseArrayOutput {
	return i.ToMABContainerHealthDetailsResponseArrayOutputWithContext(context.Background())
}

func (i MABContainerHealthDetailsResponseArray) ToMABContainerHealthDetailsResponseArrayOutputWithContext(ctx context.Context) MABContainerHealthDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MABContainerHealthDetailsResponseArrayOutput)
}

type MABContainerHealthDetailsResponseOutput struct{ *pulumi.OutputState }

func (MABContainerHealthDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MABContainerHealthDetailsResponse)(nil)).Elem()
}

func (o MABContainerHealthDetailsResponseOutput) ToMABContainerHealthDetailsResponseOutput() MABContainerHealthDetailsResponseOutput {
	return o
}

func (o MABContainerHealthDetailsResponseOutput) ToMABContainerHealthDetailsResponseOutputWithContext(ctx context.Context) MABContainerHealthDetailsResponseOutput {
	return o
}

func (o MABContainerHealthDetailsResponseOutput) Code() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MABContainerHealthDetailsResponse) *int { return v.Code }).(pulumi.IntPtrOutput)
}

func (o MABContainerHealthDetailsResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MABContainerHealthDetailsResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o MABContainerHealthDetailsResponseOutput) Recommendations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MABContainerHealthDetailsResponse) []string { return v.Recommendations }).(pulumi.StringArrayOutput)
}

func (o MABContainerHealthDetailsResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MABContainerHealthDetailsResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type MABContainerHealthDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (MABContainerHealthDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MABContainerHealthDetailsResponse)(nil)).Elem()
}

func (o MABContainerHealthDetailsResponseArrayOutput) ToMABContainerHealthDetailsResponseArrayOutput() MABContainerHealthDetailsResponseArrayOutput {
	return o
}

func (o MABContainerHealthDetailsResponseArrayOutput) ToMABContainerHealthDetailsResponseArrayOutputWithContext(ctx context.Context) MABContainerHealthDetailsResponseArrayOutput {
	return o
}

func (o MABContainerHealthDetailsResponseArrayOutput) Index(i pulumi.IntInput) MABContainerHealthDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MABContainerHealthDetailsResponse {
		return vs[0].([]MABContainerHealthDetailsResponse)[vs[1].(int)]
	}).(MABContainerHealthDetailsResponseOutput)
}

type MabContainer struct {
	AgentVersion              *string                     `pulumi:"agentVersion"`
	BackupManagementType      *string                     `pulumi:"backupManagementType"`
	CanReRegister             *bool                       `pulumi:"canReRegister"`
	ContainerHealthState      *string                     `pulumi:"containerHealthState"`
	ContainerId               *float64                    `pulumi:"containerId"`
	ContainerType             string                      `pulumi:"containerType"`
	ExtendedInfo              *MabContainerExtendedInfo   `pulumi:"extendedInfo"`
	FriendlyName              *string                     `pulumi:"friendlyName"`
	HealthStatus              *string                     `pulumi:"healthStatus"`
	MabContainerHealthDetails []MABContainerHealthDetails `pulumi:"mabContainerHealthDetails"`
	ProtectedItemCount        *float64                    `pulumi:"protectedItemCount"`
	RegistrationStatus        *string                     `pulumi:"registrationStatus"`
}





type MabContainerInput interface {
	pulumi.Input

	ToMabContainerOutput() MabContainerOutput
	ToMabContainerOutputWithContext(context.Context) MabContainerOutput
}

type MabContainerArgs struct {
	AgentVersion              pulumi.StringPtrInput               `pulumi:"agentVersion"`
	BackupManagementType      pulumi.StringPtrInput               `pulumi:"backupManagementType"`
	CanReRegister             pulumi.BoolPtrInput                 `pulumi:"canReRegister"`
	ContainerHealthState      pulumi.StringPtrInput               `pulumi:"containerHealthState"`
	ContainerId               pulumi.Float64PtrInput              `pulumi:"containerId"`
	ContainerType             pulumi.StringInput                  `pulumi:"containerType"`
	ExtendedInfo              MabContainerExtendedInfoPtrInput    `pulumi:"extendedInfo"`
	FriendlyName              pulumi.StringPtrInput               `pulumi:"friendlyName"`
	HealthStatus              pulumi.StringPtrInput               `pulumi:"healthStatus"`
	MabContainerHealthDetails MABContainerHealthDetailsArrayInput `pulumi:"mabContainerHealthDetails"`
	ProtectedItemCount        pulumi.Float64PtrInput              `pulumi:"protectedItemCount"`
	RegistrationStatus        pulumi.StringPtrInput               `pulumi:"registrationStatus"`
}

func (MabContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabContainer)(nil)).Elem()
}

func (i MabContainerArgs) ToMabContainerOutput() MabContainerOutput {
	return i.ToMabContainerOutputWithContext(context.Background())
}

func (i MabContainerArgs) ToMabContainerOutputWithContext(ctx context.Context) MabContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabContainerOutput)
}

type MabContainerOutput struct{ *pulumi.OutputState }

func (MabContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabContainer)(nil)).Elem()
}

func (o MabContainerOutput) ToMabContainerOutput() MabContainerOutput {
	return o
}

func (o MabContainerOutput) ToMabContainerOutputWithContext(ctx context.Context) MabContainerOutput {
	return o
}

func (o MabContainerOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainer) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o MabContainerOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainer) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o MabContainerOutput) CanReRegister() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MabContainer) *bool { return v.CanReRegister }).(pulumi.BoolPtrOutput)
}

func (o MabContainerOutput) ContainerHealthState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainer) *string { return v.ContainerHealthState }).(pulumi.StringPtrOutput)
}

func (o MabContainerOutput) ContainerId() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MabContainer) *float64 { return v.ContainerId }).(pulumi.Float64PtrOutput)
}

func (o MabContainerOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v MabContainer) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o MabContainerOutput) ExtendedInfo() MabContainerExtendedInfoPtrOutput {
	return o.ApplyT(func(v MabContainer) *MabContainerExtendedInfo { return v.ExtendedInfo }).(MabContainerExtendedInfoPtrOutput)
}

func (o MabContainerOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainer) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MabContainerOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainer) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o MabContainerOutput) MabContainerHealthDetails() MABContainerHealthDetailsArrayOutput {
	return o.ApplyT(func(v MabContainer) []MABContainerHealthDetails { return v.MabContainerHealthDetails }).(MABContainerHealthDetailsArrayOutput)
}

func (o MabContainerOutput) ProtectedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MabContainer) *float64 { return v.ProtectedItemCount }).(pulumi.Float64PtrOutput)
}

func (o MabContainerOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainer) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

type MabContainerExtendedInfo struct {
	BackupItemType   *string  `pulumi:"backupItemType"`
	BackupItems      []string `pulumi:"backupItems"`
	LastBackupStatus *string  `pulumi:"lastBackupStatus"`
	LastRefreshedAt  *string  `pulumi:"lastRefreshedAt"`
	PolicyName       *string  `pulumi:"policyName"`
}





type MabContainerExtendedInfoInput interface {
	pulumi.Input

	ToMabContainerExtendedInfoOutput() MabContainerExtendedInfoOutput
	ToMabContainerExtendedInfoOutputWithContext(context.Context) MabContainerExtendedInfoOutput
}

type MabContainerExtendedInfoArgs struct {
	BackupItemType   pulumi.StringPtrInput   `pulumi:"backupItemType"`
	BackupItems      pulumi.StringArrayInput `pulumi:"backupItems"`
	LastBackupStatus pulumi.StringPtrInput   `pulumi:"lastBackupStatus"`
	LastRefreshedAt  pulumi.StringPtrInput   `pulumi:"lastRefreshedAt"`
	PolicyName       pulumi.StringPtrInput   `pulumi:"policyName"`
}

func (MabContainerExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabContainerExtendedInfo)(nil)).Elem()
}

func (i MabContainerExtendedInfoArgs) ToMabContainerExtendedInfoOutput() MabContainerExtendedInfoOutput {
	return i.ToMabContainerExtendedInfoOutputWithContext(context.Background())
}

func (i MabContainerExtendedInfoArgs) ToMabContainerExtendedInfoOutputWithContext(ctx context.Context) MabContainerExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabContainerExtendedInfoOutput)
}

func (i MabContainerExtendedInfoArgs) ToMabContainerExtendedInfoPtrOutput() MabContainerExtendedInfoPtrOutput {
	return i.ToMabContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (i MabContainerExtendedInfoArgs) ToMabContainerExtendedInfoPtrOutputWithContext(ctx context.Context) MabContainerExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabContainerExtendedInfoOutput).ToMabContainerExtendedInfoPtrOutputWithContext(ctx)
}









type MabContainerExtendedInfoPtrInput interface {
	pulumi.Input

	ToMabContainerExtendedInfoPtrOutput() MabContainerExtendedInfoPtrOutput
	ToMabContainerExtendedInfoPtrOutputWithContext(context.Context) MabContainerExtendedInfoPtrOutput
}

type mabContainerExtendedInfoPtrType MabContainerExtendedInfoArgs

func MabContainerExtendedInfoPtr(v *MabContainerExtendedInfoArgs) MabContainerExtendedInfoPtrInput {
	return (*mabContainerExtendedInfoPtrType)(v)
}

func (*mabContainerExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MabContainerExtendedInfo)(nil)).Elem()
}

func (i *mabContainerExtendedInfoPtrType) ToMabContainerExtendedInfoPtrOutput() MabContainerExtendedInfoPtrOutput {
	return i.ToMabContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *mabContainerExtendedInfoPtrType) ToMabContainerExtendedInfoPtrOutputWithContext(ctx context.Context) MabContainerExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabContainerExtendedInfoPtrOutput)
}

type MabContainerExtendedInfoOutput struct{ *pulumi.OutputState }

func (MabContainerExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabContainerExtendedInfo)(nil)).Elem()
}

func (o MabContainerExtendedInfoOutput) ToMabContainerExtendedInfoOutput() MabContainerExtendedInfoOutput {
	return o
}

func (o MabContainerExtendedInfoOutput) ToMabContainerExtendedInfoOutputWithContext(ctx context.Context) MabContainerExtendedInfoOutput {
	return o
}

func (o MabContainerExtendedInfoOutput) ToMabContainerExtendedInfoPtrOutput() MabContainerExtendedInfoPtrOutput {
	return o.ToMabContainerExtendedInfoPtrOutputWithContext(context.Background())
}

func (o MabContainerExtendedInfoOutput) ToMabContainerExtendedInfoPtrOutputWithContext(ctx context.Context) MabContainerExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MabContainerExtendedInfo) *MabContainerExtendedInfo {
		return &v
	}).(MabContainerExtendedInfoPtrOutput)
}

func (o MabContainerExtendedInfoOutput) BackupItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerExtendedInfo) *string { return v.BackupItemType }).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoOutput) BackupItems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MabContainerExtendedInfo) []string { return v.BackupItems }).(pulumi.StringArrayOutput)
}

func (o MabContainerExtendedInfoOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerExtendedInfo) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerExtendedInfo) *string { return v.LastRefreshedAt }).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerExtendedInfo) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

type MabContainerExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (MabContainerExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MabContainerExtendedInfo)(nil)).Elem()
}

func (o MabContainerExtendedInfoPtrOutput) ToMabContainerExtendedInfoPtrOutput() MabContainerExtendedInfoPtrOutput {
	return o
}

func (o MabContainerExtendedInfoPtrOutput) ToMabContainerExtendedInfoPtrOutputWithContext(ctx context.Context) MabContainerExtendedInfoPtrOutput {
	return o
}

func (o MabContainerExtendedInfoPtrOutput) Elem() MabContainerExtendedInfoOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfo) MabContainerExtendedInfo {
		if v != nil {
			return *v
		}
		var ret MabContainerExtendedInfo
		return ret
	}).(MabContainerExtendedInfoOutput)
}

func (o MabContainerExtendedInfoPtrOutput) BackupItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.BackupItemType
	}).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoPtrOutput) BackupItems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfo) []string {
		if v == nil {
			return nil
		}
		return v.BackupItems
	}).(pulumi.StringArrayOutput)
}

func (o MabContainerExtendedInfoPtrOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.LastBackupStatus
	}).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoPtrOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.LastRefreshedAt
	}).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoPtrOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.PolicyName
	}).(pulumi.StringPtrOutput)
}

type MabContainerExtendedInfoResponse struct {
	BackupItemType   *string  `pulumi:"backupItemType"`
	BackupItems      []string `pulumi:"backupItems"`
	LastBackupStatus *string  `pulumi:"lastBackupStatus"`
	LastRefreshedAt  *string  `pulumi:"lastRefreshedAt"`
	PolicyName       *string  `pulumi:"policyName"`
}





type MabContainerExtendedInfoResponseInput interface {
	pulumi.Input

	ToMabContainerExtendedInfoResponseOutput() MabContainerExtendedInfoResponseOutput
	ToMabContainerExtendedInfoResponseOutputWithContext(context.Context) MabContainerExtendedInfoResponseOutput
}

type MabContainerExtendedInfoResponseArgs struct {
	BackupItemType   pulumi.StringPtrInput   `pulumi:"backupItemType"`
	BackupItems      pulumi.StringArrayInput `pulumi:"backupItems"`
	LastBackupStatus pulumi.StringPtrInput   `pulumi:"lastBackupStatus"`
	LastRefreshedAt  pulumi.StringPtrInput   `pulumi:"lastRefreshedAt"`
	PolicyName       pulumi.StringPtrInput   `pulumi:"policyName"`
}

func (MabContainerExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabContainerExtendedInfoResponse)(nil)).Elem()
}

func (i MabContainerExtendedInfoResponseArgs) ToMabContainerExtendedInfoResponseOutput() MabContainerExtendedInfoResponseOutput {
	return i.ToMabContainerExtendedInfoResponseOutputWithContext(context.Background())
}

func (i MabContainerExtendedInfoResponseArgs) ToMabContainerExtendedInfoResponseOutputWithContext(ctx context.Context) MabContainerExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabContainerExtendedInfoResponseOutput)
}

func (i MabContainerExtendedInfoResponseArgs) ToMabContainerExtendedInfoResponsePtrOutput() MabContainerExtendedInfoResponsePtrOutput {
	return i.ToMabContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i MabContainerExtendedInfoResponseArgs) ToMabContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) MabContainerExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabContainerExtendedInfoResponseOutput).ToMabContainerExtendedInfoResponsePtrOutputWithContext(ctx)
}









type MabContainerExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToMabContainerExtendedInfoResponsePtrOutput() MabContainerExtendedInfoResponsePtrOutput
	ToMabContainerExtendedInfoResponsePtrOutputWithContext(context.Context) MabContainerExtendedInfoResponsePtrOutput
}

type mabContainerExtendedInfoResponsePtrType MabContainerExtendedInfoResponseArgs

func MabContainerExtendedInfoResponsePtr(v *MabContainerExtendedInfoResponseArgs) MabContainerExtendedInfoResponsePtrInput {
	return (*mabContainerExtendedInfoResponsePtrType)(v)
}

func (*mabContainerExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MabContainerExtendedInfoResponse)(nil)).Elem()
}

func (i *mabContainerExtendedInfoResponsePtrType) ToMabContainerExtendedInfoResponsePtrOutput() MabContainerExtendedInfoResponsePtrOutput {
	return i.ToMabContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *mabContainerExtendedInfoResponsePtrType) ToMabContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) MabContainerExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabContainerExtendedInfoResponsePtrOutput)
}

type MabContainerExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (MabContainerExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabContainerExtendedInfoResponse)(nil)).Elem()
}

func (o MabContainerExtendedInfoResponseOutput) ToMabContainerExtendedInfoResponseOutput() MabContainerExtendedInfoResponseOutput {
	return o
}

func (o MabContainerExtendedInfoResponseOutput) ToMabContainerExtendedInfoResponseOutputWithContext(ctx context.Context) MabContainerExtendedInfoResponseOutput {
	return o
}

func (o MabContainerExtendedInfoResponseOutput) ToMabContainerExtendedInfoResponsePtrOutput() MabContainerExtendedInfoResponsePtrOutput {
	return o.ToMabContainerExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o MabContainerExtendedInfoResponseOutput) ToMabContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) MabContainerExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MabContainerExtendedInfoResponse) *MabContainerExtendedInfoResponse {
		return &v
	}).(MabContainerExtendedInfoResponsePtrOutput)
}

func (o MabContainerExtendedInfoResponseOutput) BackupItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerExtendedInfoResponse) *string { return v.BackupItemType }).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoResponseOutput) BackupItems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MabContainerExtendedInfoResponse) []string { return v.BackupItems }).(pulumi.StringArrayOutput)
}

func (o MabContainerExtendedInfoResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerExtendedInfoResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoResponseOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerExtendedInfoResponse) *string { return v.LastRefreshedAt }).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoResponseOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerExtendedInfoResponse) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

type MabContainerExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (MabContainerExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MabContainerExtendedInfoResponse)(nil)).Elem()
}

func (o MabContainerExtendedInfoResponsePtrOutput) ToMabContainerExtendedInfoResponsePtrOutput() MabContainerExtendedInfoResponsePtrOutput {
	return o
}

func (o MabContainerExtendedInfoResponsePtrOutput) ToMabContainerExtendedInfoResponsePtrOutputWithContext(ctx context.Context) MabContainerExtendedInfoResponsePtrOutput {
	return o
}

func (o MabContainerExtendedInfoResponsePtrOutput) Elem() MabContainerExtendedInfoResponseOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfoResponse) MabContainerExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret MabContainerExtendedInfoResponse
		return ret
	}).(MabContainerExtendedInfoResponseOutput)
}

func (o MabContainerExtendedInfoResponsePtrOutput) BackupItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.BackupItemType
	}).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoResponsePtrOutput) BackupItems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.BackupItems
	}).(pulumi.StringArrayOutput)
}

func (o MabContainerExtendedInfoResponsePtrOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastBackupStatus
	}).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoResponsePtrOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastRefreshedAt
	}).(pulumi.StringPtrOutput)
}

func (o MabContainerExtendedInfoResponsePtrOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabContainerExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyName
	}).(pulumi.StringPtrOutput)
}

type MabContainerResponse struct {
	AgentVersion              *string                             `pulumi:"agentVersion"`
	BackupManagementType      *string                             `pulumi:"backupManagementType"`
	CanReRegister             *bool                               `pulumi:"canReRegister"`
	ContainerHealthState      *string                             `pulumi:"containerHealthState"`
	ContainerId               *float64                            `pulumi:"containerId"`
	ContainerType             string                              `pulumi:"containerType"`
	ExtendedInfo              *MabContainerExtendedInfoResponse   `pulumi:"extendedInfo"`
	FriendlyName              *string                             `pulumi:"friendlyName"`
	HealthStatus              *string                             `pulumi:"healthStatus"`
	MabContainerHealthDetails []MABContainerHealthDetailsResponse `pulumi:"mabContainerHealthDetails"`
	ProtectedItemCount        *float64                            `pulumi:"protectedItemCount"`
	RegistrationStatus        *string                             `pulumi:"registrationStatus"`
}





type MabContainerResponseInput interface {
	pulumi.Input

	ToMabContainerResponseOutput() MabContainerResponseOutput
	ToMabContainerResponseOutputWithContext(context.Context) MabContainerResponseOutput
}

type MabContainerResponseArgs struct {
	AgentVersion              pulumi.StringPtrInput                       `pulumi:"agentVersion"`
	BackupManagementType      pulumi.StringPtrInput                       `pulumi:"backupManagementType"`
	CanReRegister             pulumi.BoolPtrInput                         `pulumi:"canReRegister"`
	ContainerHealthState      pulumi.StringPtrInput                       `pulumi:"containerHealthState"`
	ContainerId               pulumi.Float64PtrInput                      `pulumi:"containerId"`
	ContainerType             pulumi.StringInput                          `pulumi:"containerType"`
	ExtendedInfo              MabContainerExtendedInfoResponsePtrInput    `pulumi:"extendedInfo"`
	FriendlyName              pulumi.StringPtrInput                       `pulumi:"friendlyName"`
	HealthStatus              pulumi.StringPtrInput                       `pulumi:"healthStatus"`
	MabContainerHealthDetails MABContainerHealthDetailsResponseArrayInput `pulumi:"mabContainerHealthDetails"`
	ProtectedItemCount        pulumi.Float64PtrInput                      `pulumi:"protectedItemCount"`
	RegistrationStatus        pulumi.StringPtrInput                       `pulumi:"registrationStatus"`
}

func (MabContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabContainerResponse)(nil)).Elem()
}

func (i MabContainerResponseArgs) ToMabContainerResponseOutput() MabContainerResponseOutput {
	return i.ToMabContainerResponseOutputWithContext(context.Background())
}

func (i MabContainerResponseArgs) ToMabContainerResponseOutputWithContext(ctx context.Context) MabContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabContainerResponseOutput)
}

type MabContainerResponseOutput struct{ *pulumi.OutputState }

func (MabContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabContainerResponse)(nil)).Elem()
}

func (o MabContainerResponseOutput) ToMabContainerResponseOutput() MabContainerResponseOutput {
	return o
}

func (o MabContainerResponseOutput) ToMabContainerResponseOutputWithContext(ctx context.Context) MabContainerResponseOutput {
	return o
}

func (o MabContainerResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o MabContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o MabContainerResponseOutput) CanReRegister() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *bool { return v.CanReRegister }).(pulumi.BoolPtrOutput)
}

func (o MabContainerResponseOutput) ContainerHealthState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *string { return v.ContainerHealthState }).(pulumi.StringPtrOutput)
}

func (o MabContainerResponseOutput) ContainerId() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *float64 { return v.ContainerId }).(pulumi.Float64PtrOutput)
}

func (o MabContainerResponseOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v MabContainerResponse) string { return v.ContainerType }).(pulumi.StringOutput)
}

func (o MabContainerResponseOutput) ExtendedInfo() MabContainerExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *MabContainerExtendedInfoResponse { return v.ExtendedInfo }).(MabContainerExtendedInfoResponsePtrOutput)
}

func (o MabContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MabContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

func (o MabContainerResponseOutput) MabContainerHealthDetails() MABContainerHealthDetailsResponseArrayOutput {
	return o.ApplyT(func(v MabContainerResponse) []MABContainerHealthDetailsResponse { return v.MabContainerHealthDetails }).(MABContainerHealthDetailsResponseArrayOutput)
}

func (o MabContainerResponseOutput) ProtectedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *float64 { return v.ProtectedItemCount }).(pulumi.Float64PtrOutput)
}

func (o MabContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

type MabFileFolderProtectedItem struct {
	BackupManagementType             *string                                 `pulumi:"backupManagementType"`
	BackupSetName                    *string                                 `pulumi:"backupSetName"`
	ComputerName                     *string                                 `pulumi:"computerName"`
	ContainerName                    *string                                 `pulumi:"containerName"`
	CreateMode                       *string                                 `pulumi:"createMode"`
	DeferredDeleteSyncTimeInUTC      *float64                                `pulumi:"deferredDeleteSyncTimeInUTC"`
	DeferredDeleteTimeInUTC          *string                                 `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                 `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *MabFileFolderProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                 `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                   `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                   `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                   `pulumi:"isScheduledForDeferredDelete"`
	LastBackupStatus                 *string                                 `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                 `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                 `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                 `pulumi:"policyId"`
	ProtectedItemType                string                                  `pulumi:"protectedItemType"`
	ProtectionState                  *string                                 `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                                 `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                 `pulumi:"workloadType"`
}





type MabFileFolderProtectedItemInput interface {
	pulumi.Input

	ToMabFileFolderProtectedItemOutput() MabFileFolderProtectedItemOutput
	ToMabFileFolderProtectedItemOutputWithContext(context.Context) MabFileFolderProtectedItemOutput
}

type MabFileFolderProtectedItemArgs struct {
	BackupManagementType             pulumi.StringPtrInput                          `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                          `pulumi:"backupSetName"`
	ComputerName                     pulumi.StringPtrInput                          `pulumi:"computerName"`
	ContainerName                    pulumi.StringPtrInput                          `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                          `pulumi:"createMode"`
	DeferredDeleteSyncTimeInUTC      pulumi.Float64PtrInput                         `pulumi:"deferredDeleteSyncTimeInUTC"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                          `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                          `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     MabFileFolderProtectedItemExtendedInfoPtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                          `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                            `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                            `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                            `pulumi:"isScheduledForDeferredDelete"`
	LastBackupStatus                 pulumi.StringPtrInput                          `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                          `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                          `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                          `pulumi:"policyId"`
	ProtectedItemType                pulumi.StringInput                             `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                          `pulumi:"protectionState"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                        `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                          `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                          `pulumi:"workloadType"`
}

func (MabFileFolderProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabFileFolderProtectedItem)(nil)).Elem()
}

func (i MabFileFolderProtectedItemArgs) ToMabFileFolderProtectedItemOutput() MabFileFolderProtectedItemOutput {
	return i.ToMabFileFolderProtectedItemOutputWithContext(context.Background())
}

func (i MabFileFolderProtectedItemArgs) ToMabFileFolderProtectedItemOutputWithContext(ctx context.Context) MabFileFolderProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabFileFolderProtectedItemOutput)
}

type MabFileFolderProtectedItemOutput struct{ *pulumi.OutputState }

func (MabFileFolderProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabFileFolderProtectedItem)(nil)).Elem()
}

func (o MabFileFolderProtectedItemOutput) ToMabFileFolderProtectedItemOutput() MabFileFolderProtectedItemOutput {
	return o
}

func (o MabFileFolderProtectedItemOutput) ToMabFileFolderProtectedItemOutputWithContext(ctx context.Context) MabFileFolderProtectedItemOutput {
	return o
}

func (o MabFileFolderProtectedItemOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) DeferredDeleteSyncTimeInUTC() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *float64 { return v.DeferredDeleteSyncTimeInUTC }).(pulumi.Float64PtrOutput)
}

func (o MabFileFolderProtectedItemOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) ExtendedInfo() MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *MabFileFolderProtectedItemExtendedInfo { return v.ExtendedInfo }).(MabFileFolderProtectedItemExtendedInfoPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o MabFileFolderProtectedItemOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o MabFileFolderProtectedItemOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItem) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type MabFileFolderProtectedItemExtendedInfo struct {
	LastRefreshedAt     *string `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type MabFileFolderProtectedItemExtendedInfoInput interface {
	pulumi.Input

	ToMabFileFolderProtectedItemExtendedInfoOutput() MabFileFolderProtectedItemExtendedInfoOutput
	ToMabFileFolderProtectedItemExtendedInfoOutputWithContext(context.Context) MabFileFolderProtectedItemExtendedInfoOutput
}

type MabFileFolderProtectedItemExtendedInfoArgs struct {
	LastRefreshedAt     pulumi.StringPtrInput `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (MabFileFolderProtectedItemExtendedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabFileFolderProtectedItemExtendedInfo)(nil)).Elem()
}

func (i MabFileFolderProtectedItemExtendedInfoArgs) ToMabFileFolderProtectedItemExtendedInfoOutput() MabFileFolderProtectedItemExtendedInfoOutput {
	return i.ToMabFileFolderProtectedItemExtendedInfoOutputWithContext(context.Background())
}

func (i MabFileFolderProtectedItemExtendedInfoArgs) ToMabFileFolderProtectedItemExtendedInfoOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabFileFolderProtectedItemExtendedInfoOutput)
}

func (i MabFileFolderProtectedItemExtendedInfoArgs) ToMabFileFolderProtectedItemExtendedInfoPtrOutput() MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return i.ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i MabFileFolderProtectedItemExtendedInfoArgs) ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabFileFolderProtectedItemExtendedInfoOutput).ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(ctx)
}









type MabFileFolderProtectedItemExtendedInfoPtrInput interface {
	pulumi.Input

	ToMabFileFolderProtectedItemExtendedInfoPtrOutput() MabFileFolderProtectedItemExtendedInfoPtrOutput
	ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(context.Context) MabFileFolderProtectedItemExtendedInfoPtrOutput
}

type mabFileFolderProtectedItemExtendedInfoPtrType MabFileFolderProtectedItemExtendedInfoArgs

func MabFileFolderProtectedItemExtendedInfoPtr(v *MabFileFolderProtectedItemExtendedInfoArgs) MabFileFolderProtectedItemExtendedInfoPtrInput {
	return (*mabFileFolderProtectedItemExtendedInfoPtrType)(v)
}

func (*mabFileFolderProtectedItemExtendedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MabFileFolderProtectedItemExtendedInfo)(nil)).Elem()
}

func (i *mabFileFolderProtectedItemExtendedInfoPtrType) ToMabFileFolderProtectedItemExtendedInfoPtrOutput() MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return i.ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (i *mabFileFolderProtectedItemExtendedInfoPtrType) ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabFileFolderProtectedItemExtendedInfoPtrOutput)
}

type MabFileFolderProtectedItemExtendedInfoOutput struct{ *pulumi.OutputState }

func (MabFileFolderProtectedItemExtendedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabFileFolderProtectedItemExtendedInfo)(nil)).Elem()
}

func (o MabFileFolderProtectedItemExtendedInfoOutput) ToMabFileFolderProtectedItemExtendedInfoOutput() MabFileFolderProtectedItemExtendedInfoOutput {
	return o
}

func (o MabFileFolderProtectedItemExtendedInfoOutput) ToMabFileFolderProtectedItemExtendedInfoOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoOutput {
	return o
}

func (o MabFileFolderProtectedItemExtendedInfoOutput) ToMabFileFolderProtectedItemExtendedInfoPtrOutput() MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return o.ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(context.Background())
}

func (o MabFileFolderProtectedItemExtendedInfoOutput) ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MabFileFolderProtectedItemExtendedInfo) *MabFileFolderProtectedItemExtendedInfo {
		return &v
	}).(MabFileFolderProtectedItemExtendedInfoPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemExtendedInfo) *string { return v.LastRefreshedAt }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemExtendedInfo) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemExtendedInfo) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type MabFileFolderProtectedItemExtendedInfoPtrOutput struct{ *pulumi.OutputState }

func (MabFileFolderProtectedItemExtendedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MabFileFolderProtectedItemExtendedInfo)(nil)).Elem()
}

func (o MabFileFolderProtectedItemExtendedInfoPtrOutput) ToMabFileFolderProtectedItemExtendedInfoPtrOutput() MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o MabFileFolderProtectedItemExtendedInfoPtrOutput) ToMabFileFolderProtectedItemExtendedInfoPtrOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoPtrOutput {
	return o
}

func (o MabFileFolderProtectedItemExtendedInfoPtrOutput) Elem() MabFileFolderProtectedItemExtendedInfoOutput {
	return o.ApplyT(func(v *MabFileFolderProtectedItemExtendedInfo) MabFileFolderProtectedItemExtendedInfo {
		if v != nil {
			return *v
		}
		var ret MabFileFolderProtectedItemExtendedInfo
		return ret
	}).(MabFileFolderProtectedItemExtendedInfoOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoPtrOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabFileFolderProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.LastRefreshedAt
	}).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoPtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabFileFolderProtectedItemExtendedInfo) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoPtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MabFileFolderProtectedItemExtendedInfo) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type MabFileFolderProtectedItemExtendedInfoResponse struct {
	LastRefreshedAt     *string `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}





type MabFileFolderProtectedItemExtendedInfoResponseInput interface {
	pulumi.Input

	ToMabFileFolderProtectedItemExtendedInfoResponseOutput() MabFileFolderProtectedItemExtendedInfoResponseOutput
	ToMabFileFolderProtectedItemExtendedInfoResponseOutputWithContext(context.Context) MabFileFolderProtectedItemExtendedInfoResponseOutput
}

type MabFileFolderProtectedItemExtendedInfoResponseArgs struct {
	LastRefreshedAt     pulumi.StringPtrInput `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint pulumi.StringPtrInput `pulumi:"oldestRecoveryPoint"`
	RecoveryPointCount  pulumi.IntPtrInput    `pulumi:"recoveryPointCount"`
}

func (MabFileFolderProtectedItemExtendedInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabFileFolderProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i MabFileFolderProtectedItemExtendedInfoResponseArgs) ToMabFileFolderProtectedItemExtendedInfoResponseOutput() MabFileFolderProtectedItemExtendedInfoResponseOutput {
	return i.ToMabFileFolderProtectedItemExtendedInfoResponseOutputWithContext(context.Background())
}

func (i MabFileFolderProtectedItemExtendedInfoResponseArgs) ToMabFileFolderProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabFileFolderProtectedItemExtendedInfoResponseOutput)
}

func (i MabFileFolderProtectedItemExtendedInfoResponseArgs) ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutput() MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i MabFileFolderProtectedItemExtendedInfoResponseArgs) ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabFileFolderProtectedItemExtendedInfoResponseOutput).ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx)
}









type MabFileFolderProtectedItemExtendedInfoResponsePtrInput interface {
	pulumi.Input

	ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutput() MabFileFolderProtectedItemExtendedInfoResponsePtrOutput
	ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Context) MabFileFolderProtectedItemExtendedInfoResponsePtrOutput
}

type mabFileFolderProtectedItemExtendedInfoResponsePtrType MabFileFolderProtectedItemExtendedInfoResponseArgs

func MabFileFolderProtectedItemExtendedInfoResponsePtr(v *MabFileFolderProtectedItemExtendedInfoResponseArgs) MabFileFolderProtectedItemExtendedInfoResponsePtrInput {
	return (*mabFileFolderProtectedItemExtendedInfoResponsePtrType)(v)
}

func (*mabFileFolderProtectedItemExtendedInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MabFileFolderProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (i *mabFileFolderProtectedItemExtendedInfoResponsePtrType) ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutput() MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return i.ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (i *mabFileFolderProtectedItemExtendedInfoResponsePtrType) ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabFileFolderProtectedItemExtendedInfoResponsePtrOutput)
}

type MabFileFolderProtectedItemExtendedInfoResponseOutput struct{ *pulumi.OutputState }

func (MabFileFolderProtectedItemExtendedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabFileFolderProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o MabFileFolderProtectedItemExtendedInfoResponseOutput) ToMabFileFolderProtectedItemExtendedInfoResponseOutput() MabFileFolderProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o MabFileFolderProtectedItemExtendedInfoResponseOutput) ToMabFileFolderProtectedItemExtendedInfoResponseOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoResponseOutput {
	return o
}

func (o MabFileFolderProtectedItemExtendedInfoResponseOutput) ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutput() MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return o.ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(context.Background())
}

func (o MabFileFolderProtectedItemExtendedInfoResponseOutput) ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MabFileFolderProtectedItemExtendedInfoResponse) *MabFileFolderProtectedItemExtendedInfoResponse {
		return &v
	}).(MabFileFolderProtectedItemExtendedInfoResponsePtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoResponseOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemExtendedInfoResponse) *string { return v.LastRefreshedAt }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoResponseOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemExtendedInfoResponse) *string { return v.OldestRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoResponseOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemExtendedInfoResponse) *int { return v.RecoveryPointCount }).(pulumi.IntPtrOutput)
}

type MabFileFolderProtectedItemExtendedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (MabFileFolderProtectedItemExtendedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MabFileFolderProtectedItemExtendedInfoResponse)(nil)).Elem()
}

func (o MabFileFolderProtectedItemExtendedInfoResponsePtrOutput) ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutput() MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o MabFileFolderProtectedItemExtendedInfoResponsePtrOutput) ToMabFileFolderProtectedItemExtendedInfoResponsePtrOutputWithContext(ctx context.Context) MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return o
}

func (o MabFileFolderProtectedItemExtendedInfoResponsePtrOutput) Elem() MabFileFolderProtectedItemExtendedInfoResponseOutput {
	return o.ApplyT(func(v *MabFileFolderProtectedItemExtendedInfoResponse) MabFileFolderProtectedItemExtendedInfoResponse {
		if v != nil {
			return *v
		}
		var ret MabFileFolderProtectedItemExtendedInfoResponse
		return ret
	}).(MabFileFolderProtectedItemExtendedInfoResponseOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoResponsePtrOutput) LastRefreshedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabFileFolderProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastRefreshedAt
	}).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoResponsePtrOutput) OldestRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MabFileFolderProtectedItemExtendedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OldestRecoveryPoint
	}).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemExtendedInfoResponsePtrOutput) RecoveryPointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MabFileFolderProtectedItemExtendedInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.RecoveryPointCount
	}).(pulumi.IntPtrOutput)
}

type MabFileFolderProtectedItemResponse struct {
	BackupManagementType             *string                                         `pulumi:"backupManagementType"`
	BackupSetName                    *string                                         `pulumi:"backupSetName"`
	ComputerName                     *string                                         `pulumi:"computerName"`
	ContainerName                    *string                                         `pulumi:"containerName"`
	CreateMode                       *string                                         `pulumi:"createMode"`
	DeferredDeleteSyncTimeInUTC      *float64                                        `pulumi:"deferredDeleteSyncTimeInUTC"`
	DeferredDeleteTimeInUTC          *string                                         `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                         `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *MabFileFolderProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                         `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                           `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                           `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                           `pulumi:"isScheduledForDeferredDelete"`
	LastBackupStatus                 *string                                         `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                         `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                         `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                         `pulumi:"policyId"`
	ProtectedItemType                string                                          `pulumi:"protectedItemType"`
	ProtectionState                  *string                                         `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                        `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 *string                                         `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                         `pulumi:"workloadType"`
}





type MabFileFolderProtectedItemResponseInput interface {
	pulumi.Input

	ToMabFileFolderProtectedItemResponseOutput() MabFileFolderProtectedItemResponseOutput
	ToMabFileFolderProtectedItemResponseOutputWithContext(context.Context) MabFileFolderProtectedItemResponseOutput
}

type MabFileFolderProtectedItemResponseArgs struct {
	BackupManagementType             pulumi.StringPtrInput                                  `pulumi:"backupManagementType"`
	BackupSetName                    pulumi.StringPtrInput                                  `pulumi:"backupSetName"`
	ComputerName                     pulumi.StringPtrInput                                  `pulumi:"computerName"`
	ContainerName                    pulumi.StringPtrInput                                  `pulumi:"containerName"`
	CreateMode                       pulumi.StringPtrInput                                  `pulumi:"createMode"`
	DeferredDeleteSyncTimeInUTC      pulumi.Float64PtrInput                                 `pulumi:"deferredDeleteSyncTimeInUTC"`
	DeferredDeleteTimeInUTC          pulumi.StringPtrInput                                  `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      pulumi.StringPtrInput                                  `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     MabFileFolderProtectedItemExtendedInfoResponsePtrInput `pulumi:"extendedInfo"`
	FriendlyName                     pulumi.StringPtrInput                                  `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming pulumi.BoolPtrInput                                    `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      pulumi.BoolPtrInput                                    `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     pulumi.BoolPtrInput                                    `pulumi:"isScheduledForDeferredDelete"`
	LastBackupStatus                 pulumi.StringPtrInput                                  `pulumi:"lastBackupStatus"`
	LastBackupTime                   pulumi.StringPtrInput                                  `pulumi:"lastBackupTime"`
	LastRecoveryPoint                pulumi.StringPtrInput                                  `pulumi:"lastRecoveryPoint"`
	PolicyId                         pulumi.StringPtrInput                                  `pulumi:"policyId"`
	ProtectedItemType                pulumi.StringInput                                     `pulumi:"protectedItemType"`
	ProtectionState                  pulumi.StringPtrInput                                  `pulumi:"protectionState"`
	ResourceGuardOperationRequests   pulumi.StringArrayInput                                `pulumi:"resourceGuardOperationRequests"`
	SourceResourceId                 pulumi.StringPtrInput                                  `pulumi:"sourceResourceId"`
	WorkloadType                     pulumi.StringPtrInput                                  `pulumi:"workloadType"`
}

func (MabFileFolderProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabFileFolderProtectedItemResponse)(nil)).Elem()
}

func (i MabFileFolderProtectedItemResponseArgs) ToMabFileFolderProtectedItemResponseOutput() MabFileFolderProtectedItemResponseOutput {
	return i.ToMabFileFolderProtectedItemResponseOutputWithContext(context.Background())
}

func (i MabFileFolderProtectedItemResponseArgs) ToMabFileFolderProtectedItemResponseOutputWithContext(ctx context.Context) MabFileFolderProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabFileFolderProtectedItemResponseOutput)
}

type MabFileFolderProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (MabFileFolderProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabFileFolderProtectedItemResponse)(nil)).Elem()
}

func (o MabFileFolderProtectedItemResponseOutput) ToMabFileFolderProtectedItemResponseOutput() MabFileFolderProtectedItemResponseOutput {
	return o
}

func (o MabFileFolderProtectedItemResponseOutput) ToMabFileFolderProtectedItemResponseOutputWithContext(ctx context.Context) MabFileFolderProtectedItemResponseOutput {
	return o
}

func (o MabFileFolderProtectedItemResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) BackupSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.BackupSetName }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) DeferredDeleteSyncTimeInUTC() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *float64 { return v.DeferredDeleteSyncTimeInUTC }).(pulumi.Float64PtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) DeferredDeleteTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.DeferredDeleteTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) DeferredDeleteTimeRemaining() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.DeferredDeleteTimeRemaining }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) ExtendedInfo() MabFileFolderProtectedItemExtendedInfoResponsePtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *MabFileFolderProtectedItemExtendedInfoResponse {
		return v.ExtendedInfo
	}).(MabFileFolderProtectedItemExtendedInfoResponsePtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) IsDeferredDeleteScheduleUpcoming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *bool { return v.IsDeferredDeleteScheduleUpcoming }).(pulumi.BoolPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) IsRehydrate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *bool { return v.IsRehydrate }).(pulumi.BoolPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) IsScheduledForDeferredDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *bool { return v.IsScheduledForDeferredDelete }).(pulumi.BoolPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) LastBackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.LastBackupStatus }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) LastBackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.LastBackupTime }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) LastRecoveryPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.LastRecoveryPoint }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) ProtectedItemType() pulumi.StringOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) string { return v.ProtectedItemType }).(pulumi.StringOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o MabFileFolderProtectedItemResponseOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabFileFolderProtectedItemResponse) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

type MabProtectionPolicy struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{} `pulumi:"schedulePolicy"`
}





type MabProtectionPolicyInput interface {
	pulumi.Input

	ToMabProtectionPolicyOutput() MabProtectionPolicyOutput
	ToMabProtectionPolicyOutputWithContext(context.Context) MabProtectionPolicyOutput
}

type MabProtectionPolicyArgs struct {
	BackupManagementType           pulumi.StringInput      `pulumi:"backupManagementType"`
	ProtectedItemsCount            pulumi.IntPtrInput      `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                pulumi.Input            `pulumi:"retentionPolicy"`
	SchedulePolicy                 pulumi.Input            `pulumi:"schedulePolicy"`
}

func (MabProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabProtectionPolicy)(nil)).Elem()
}

func (i MabProtectionPolicyArgs) ToMabProtectionPolicyOutput() MabProtectionPolicyOutput {
	return i.ToMabProtectionPolicyOutputWithContext(context.Background())
}

func (i MabProtectionPolicyArgs) ToMabProtectionPolicyOutputWithContext(ctx context.Context) MabProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabProtectionPolicyOutput)
}

type MabProtectionPolicyOutput struct{ *pulumi.OutputState }

func (MabProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabProtectionPolicy)(nil)).Elem()
}

func (o MabProtectionPolicyOutput) ToMabProtectionPolicyOutput() MabProtectionPolicyOutput {
	return o
}

func (o MabProtectionPolicyOutput) ToMabProtectionPolicyOutputWithContext(ctx context.Context) MabProtectionPolicyOutput {
	return o
}

func (o MabProtectionPolicyOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v MabProtectionPolicy) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o MabProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MabProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o MabProtectionPolicyOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MabProtectionPolicy) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o MabProtectionPolicyOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v MabProtectionPolicy) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o MabProtectionPolicyOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v MabProtectionPolicy) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

type MabProtectionPolicyResponse struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{} `pulumi:"schedulePolicy"`
}





type MabProtectionPolicyResponseInput interface {
	pulumi.Input

	ToMabProtectionPolicyResponseOutput() MabProtectionPolicyResponseOutput
	ToMabProtectionPolicyResponseOutputWithContext(context.Context) MabProtectionPolicyResponseOutput
}

type MabProtectionPolicyResponseArgs struct {
	BackupManagementType           pulumi.StringInput      `pulumi:"backupManagementType"`
	ProtectedItemsCount            pulumi.IntPtrInput      `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests pulumi.StringArrayInput `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                pulumi.Input            `pulumi:"retentionPolicy"`
	SchedulePolicy                 pulumi.Input            `pulumi:"schedulePolicy"`
}

func (MabProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabProtectionPolicyResponse)(nil)).Elem()
}

func (i MabProtectionPolicyResponseArgs) ToMabProtectionPolicyResponseOutput() MabProtectionPolicyResponseOutput {
	return i.ToMabProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i MabProtectionPolicyResponseArgs) ToMabProtectionPolicyResponseOutputWithContext(ctx context.Context) MabProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabProtectionPolicyResponseOutput)
}

type MabProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (MabProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabProtectionPolicyResponse)(nil)).Elem()
}

func (o MabProtectionPolicyResponseOutput) ToMabProtectionPolicyResponseOutput() MabProtectionPolicyResponseOutput {
	return o
}

func (o MabProtectionPolicyResponseOutput) ToMabProtectionPolicyResponseOutputWithContext(ctx context.Context) MabProtectionPolicyResponseOutput {
	return o
}

func (o MabProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) string { return v.BackupManagementType }).(pulumi.StringOutput)
}

func (o MabProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o MabProtectionPolicyResponseOutput) ResourceGuardOperationRequests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) []string { return v.ResourceGuardOperationRequests }).(pulumi.StringArrayOutput)
}

func (o MabProtectionPolicyResponseOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o MabProtectionPolicyResponseOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

type MonthlyRetentionSchedule struct {
	RetentionDuration           *RetentionDuration     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormat  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormat `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string               `pulumi:"retentionTimes"`
}





type MonthlyRetentionScheduleInput interface {
	pulumi.Input

	ToMonthlyRetentionScheduleOutput() MonthlyRetentionScheduleOutput
	ToMonthlyRetentionScheduleOutputWithContext(context.Context) MonthlyRetentionScheduleOutput
}

type MonthlyRetentionScheduleArgs struct {
	RetentionDuration           RetentionDurationPtrInput     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      DailyRetentionFormatPtrInput  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType pulumi.StringPtrInput         `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     WeeklyRetentionFormatPtrInput `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              pulumi.StringArrayInput       `pulumi:"retentionTimes"`
}

func (MonthlyRetentionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRetentionSchedule)(nil)).Elem()
}

func (i MonthlyRetentionScheduleArgs) ToMonthlyRetentionScheduleOutput() MonthlyRetentionScheduleOutput {
	return i.ToMonthlyRetentionScheduleOutputWithContext(context.Background())
}

func (i MonthlyRetentionScheduleArgs) ToMonthlyRetentionScheduleOutputWithContext(ctx context.Context) MonthlyRetentionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleOutput)
}

func (i MonthlyRetentionScheduleArgs) ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput {
	return i.ToMonthlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i MonthlyRetentionScheduleArgs) ToMonthlyRetentionSchedulePtrOutputWithContext(ctx context.Context) MonthlyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleOutput).ToMonthlyRetentionSchedulePtrOutputWithContext(ctx)
}









type MonthlyRetentionSchedulePtrInput interface {
	pulumi.Input

	ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput
	ToMonthlyRetentionSchedulePtrOutputWithContext(context.Context) MonthlyRetentionSchedulePtrOutput
}

type monthlyRetentionSchedulePtrType MonthlyRetentionScheduleArgs

func MonthlyRetentionSchedulePtr(v *MonthlyRetentionScheduleArgs) MonthlyRetentionSchedulePtrInput {
	return (*monthlyRetentionSchedulePtrType)(v)
}

func (*monthlyRetentionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyRetentionSchedule)(nil)).Elem()
}

func (i *monthlyRetentionSchedulePtrType) ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput {
	return i.ToMonthlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i *monthlyRetentionSchedulePtrType) ToMonthlyRetentionSchedulePtrOutputWithContext(ctx context.Context) MonthlyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionSchedulePtrOutput)
}

type MonthlyRetentionScheduleOutput struct{ *pulumi.OutputState }

func (MonthlyRetentionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRetentionSchedule)(nil)).Elem()
}

func (o MonthlyRetentionScheduleOutput) ToMonthlyRetentionScheduleOutput() MonthlyRetentionScheduleOutput {
	return o
}

func (o MonthlyRetentionScheduleOutput) ToMonthlyRetentionScheduleOutputWithContext(ctx context.Context) MonthlyRetentionScheduleOutput {
	return o
}

func (o MonthlyRetentionScheduleOutput) ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput {
	return o.ToMonthlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (o MonthlyRetentionScheduleOutput) ToMonthlyRetentionSchedulePtrOutputWithContext(ctx context.Context) MonthlyRetentionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonthlyRetentionSchedule) *MonthlyRetentionSchedule {
		return &v
	}).(MonthlyRetentionSchedulePtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionScheduleDaily() DailyRetentionFormatPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) *DailyRetentionFormat { return v.RetentionScheduleDaily }).(DailyRetentionFormatPtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) *string { return v.RetentionScheduleFormatType }).(pulumi.StringPtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionScheduleWeekly() WeeklyRetentionFormatPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) *WeeklyRetentionFormat { return v.RetentionScheduleWeekly }).(WeeklyRetentionFormatPtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type MonthlyRetentionSchedulePtrOutput struct{ *pulumi.OutputState }

func (MonthlyRetentionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyRetentionSchedule)(nil)).Elem()
}

func (o MonthlyRetentionSchedulePtrOutput) ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput {
	return o
}

func (o MonthlyRetentionSchedulePtrOutput) ToMonthlyRetentionSchedulePtrOutputWithContext(ctx context.Context) MonthlyRetentionSchedulePtrOutput {
	return o
}

func (o MonthlyRetentionSchedulePtrOutput) Elem() MonthlyRetentionScheduleOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) MonthlyRetentionSchedule {
		if v != nil {
			return *v
		}
		var ret MonthlyRetentionSchedule
		return ret
	}).(MonthlyRetentionScheduleOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) *RetentionDuration {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationPtrOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionScheduleDaily() DailyRetentionFormatPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) *DailyRetentionFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatPtrOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) *string {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleFormatType
	}).(pulumi.StringPtrOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionScheduleWeekly() WeeklyRetentionFormatPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) *WeeklyRetentionFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatPtrOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type MonthlyRetentionScheduleResponse struct {
	RetentionDuration           *RetentionDurationResponse     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormatResponse  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                        `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormatResponse `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string                       `pulumi:"retentionTimes"`
}





type MonthlyRetentionScheduleResponseInput interface {
	pulumi.Input

	ToMonthlyRetentionScheduleResponseOutput() MonthlyRetentionScheduleResponseOutput
	ToMonthlyRetentionScheduleResponseOutputWithContext(context.Context) MonthlyRetentionScheduleResponseOutput
}

type MonthlyRetentionScheduleResponseArgs struct {
	RetentionDuration           RetentionDurationResponsePtrInput     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      DailyRetentionFormatResponsePtrInput  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType pulumi.StringPtrInput                 `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     WeeklyRetentionFormatResponsePtrInput `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              pulumi.StringArrayInput               `pulumi:"retentionTimes"`
}

func (MonthlyRetentionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRetentionScheduleResponse)(nil)).Elem()
}

func (i MonthlyRetentionScheduleResponseArgs) ToMonthlyRetentionScheduleResponseOutput() MonthlyRetentionScheduleResponseOutput {
	return i.ToMonthlyRetentionScheduleResponseOutputWithContext(context.Background())
}

func (i MonthlyRetentionScheduleResponseArgs) ToMonthlyRetentionScheduleResponseOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleResponseOutput)
}

func (i MonthlyRetentionScheduleResponseArgs) ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput {
	return i.ToMonthlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i MonthlyRetentionScheduleResponseArgs) ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleResponseOutput).ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx)
}









type MonthlyRetentionScheduleResponsePtrInput interface {
	pulumi.Input

	ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput
	ToMonthlyRetentionScheduleResponsePtrOutputWithContext(context.Context) MonthlyRetentionScheduleResponsePtrOutput
}

type monthlyRetentionScheduleResponsePtrType MonthlyRetentionScheduleResponseArgs

func MonthlyRetentionScheduleResponsePtr(v *MonthlyRetentionScheduleResponseArgs) MonthlyRetentionScheduleResponsePtrInput {
	return (*monthlyRetentionScheduleResponsePtrType)(v)
}

func (*monthlyRetentionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyRetentionScheduleResponse)(nil)).Elem()
}

func (i *monthlyRetentionScheduleResponsePtrType) ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput {
	return i.ToMonthlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *monthlyRetentionScheduleResponsePtrType) ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleResponsePtrOutput)
}

type MonthlyRetentionScheduleResponseOutput struct{ *pulumi.OutputState }

func (MonthlyRetentionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRetentionScheduleResponse)(nil)).Elem()
}

func (o MonthlyRetentionScheduleResponseOutput) ToMonthlyRetentionScheduleResponseOutput() MonthlyRetentionScheduleResponseOutput {
	return o
}

func (o MonthlyRetentionScheduleResponseOutput) ToMonthlyRetentionScheduleResponseOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponseOutput {
	return o
}

func (o MonthlyRetentionScheduleResponseOutput) ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput {
	return o.ToMonthlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o MonthlyRetentionScheduleResponseOutput) ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonthlyRetentionScheduleResponse) *MonthlyRetentionScheduleResponse {
		return &v
	}).(MonthlyRetentionScheduleResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionScheduleDaily() DailyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) *DailyRetentionFormatResponse {
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) *string { return v.RetentionScheduleFormatType }).(pulumi.StringPtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionScheduleWeekly() WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) *WeeklyRetentionFormatResponse {
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type MonthlyRetentionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (MonthlyRetentionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyRetentionScheduleResponse)(nil)).Elem()
}

func (o MonthlyRetentionScheduleResponsePtrOutput) ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput {
	return o
}

func (o MonthlyRetentionScheduleResponsePtrOutput) ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponsePtrOutput {
	return o
}

func (o MonthlyRetentionScheduleResponsePtrOutput) Elem() MonthlyRetentionScheduleResponseOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) MonthlyRetentionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret MonthlyRetentionScheduleResponse
		return ret
	}).(MonthlyRetentionScheduleResponseOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) *RetentionDurationResponse {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionScheduleDaily() DailyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) *DailyRetentionFormatResponse {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleFormatType
	}).(pulumi.StringPtrOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionScheduleWeekly() WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) *WeeklyRetentionFormatResponse {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionType struct {
	PrivateEndpoint                   *PrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                            `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringPtrInput                     `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput {
	return i.ToPrivateEndpointConnectionTypePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput).ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionTypePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput
	ToPrivateEndpointConnectionTypePtrOutputWithContext(context.Context) PrivateEndpointConnectionTypePtrOutput
}

type privateEndpointConnectionTypePtrType PrivateEndpointConnectionTypeArgs

func PrivateEndpointConnectionTypePtr(v *PrivateEndpointConnectionTypeArgs) PrivateEndpointConnectionTypePtrInput {
	return (*privateEndpointConnectionTypePtrType)(v)
}

func (*privateEndpointConnectionTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionType)(nil)).Elem()
}

func (i *privateEndpointConnectionTypePtrType) ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput {
	return i.ToPrivateEndpointConnectionTypePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionTypePtrType) ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypePtrOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput {
	return o.ToPrivateEndpointConnectionTypePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionType) *PrivateEndpointConnectionType {
		return &v
	}).(PrivateEndpointConnectionTypePtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateEndpoint { return v.PrivateEndpoint }).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionTypePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypePtrOutput) ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput {
	return o
}

func (o PrivateEndpointConnectionTypePtrOutput) ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypePtrOutput {
	return o
}

func (o PrivateEndpointConnectionTypePtrOutput) Elem() PrivateEndpointConnectionTypeOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionType) PrivateEndpointConnectionType {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionType
		return ret
	}).(PrivateEndpointConnectionTypeOutput)
}

func (o PrivateEndpointConnectionTypePtrOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionType) *PrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionTypePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionType) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                                    `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringPtrInput                             `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponsePtrOutput() PrivateEndpointConnectionResponsePtrOutput {
	return i.ToPrivateEndpointConnectionResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput).ToPrivateEndpointConnectionResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponsePtrOutput() PrivateEndpointConnectionResponsePtrOutput
	ToPrivateEndpointConnectionResponsePtrOutputWithContext(context.Context) PrivateEndpointConnectionResponsePtrOutput
}

type privateEndpointConnectionResponsePtrType PrivateEndpointConnectionResponseArgs

func PrivateEndpointConnectionResponsePtr(v *PrivateEndpointConnectionResponseArgs) PrivateEndpointConnectionResponsePtrInput {
	return (*privateEndpointConnectionResponsePtrType)(v)
}

func (*privateEndpointConnectionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i *privateEndpointConnectionResponsePtrType) ToPrivateEndpointConnectionResponsePtrOutput() PrivateEndpointConnectionResponsePtrOutput {
	return i.ToPrivateEndpointConnectionResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionResponsePtrType) ToPrivateEndpointConnectionResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponsePtrOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponsePtrOutput() PrivateEndpointConnectionResponsePtrOutput {
	return o.ToPrivateEndpointConnectionResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionResponse) *PrivateEndpointConnectionResponse {
		return &v
	}).(PrivateEndpointConnectionResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponsePtrOutput) ToPrivateEndpointConnectionResponsePtrOutput() PrivateEndpointConnectionResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionResponsePtrOutput) ToPrivateEndpointConnectionResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionResponsePtrOutput) Elem() PrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponse) PrivateEndpointConnectionResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionResponse
		return ret
	}).(PrivateEndpointConnectionResponseOutput)
}

func (o PrivateEndpointConnectionResponsePtrOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponse) *PrivateEndpointResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStateResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ResourceGuardOperationDetailResponse struct {
	DefaultResourceRequest *string `pulumi:"defaultResourceRequest"`
	VaultCriticalOperation *string `pulumi:"vaultCriticalOperation"`
}





type ResourceGuardOperationDetailResponseInput interface {
	pulumi.Input

	ToResourceGuardOperationDetailResponseOutput() ResourceGuardOperationDetailResponseOutput
	ToResourceGuardOperationDetailResponseOutputWithContext(context.Context) ResourceGuardOperationDetailResponseOutput
}

type ResourceGuardOperationDetailResponseArgs struct {
	DefaultResourceRequest pulumi.StringPtrInput `pulumi:"defaultResourceRequest"`
	VaultCriticalOperation pulumi.StringPtrInput `pulumi:"vaultCriticalOperation"`
}

func (ResourceGuardOperationDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardOperationDetailResponse)(nil)).Elem()
}

func (i ResourceGuardOperationDetailResponseArgs) ToResourceGuardOperationDetailResponseOutput() ResourceGuardOperationDetailResponseOutput {
	return i.ToResourceGuardOperationDetailResponseOutputWithContext(context.Background())
}

func (i ResourceGuardOperationDetailResponseArgs) ToResourceGuardOperationDetailResponseOutputWithContext(ctx context.Context) ResourceGuardOperationDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardOperationDetailResponseOutput)
}





type ResourceGuardOperationDetailResponseArrayInput interface {
	pulumi.Input

	ToResourceGuardOperationDetailResponseArrayOutput() ResourceGuardOperationDetailResponseArrayOutput
	ToResourceGuardOperationDetailResponseArrayOutputWithContext(context.Context) ResourceGuardOperationDetailResponseArrayOutput
}

type ResourceGuardOperationDetailResponseArray []ResourceGuardOperationDetailResponseInput

func (ResourceGuardOperationDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceGuardOperationDetailResponse)(nil)).Elem()
}

func (i ResourceGuardOperationDetailResponseArray) ToResourceGuardOperationDetailResponseArrayOutput() ResourceGuardOperationDetailResponseArrayOutput {
	return i.ToResourceGuardOperationDetailResponseArrayOutputWithContext(context.Background())
}

func (i ResourceGuardOperationDetailResponseArray) ToResourceGuardOperationDetailResponseArrayOutputWithContext(ctx context.Context) ResourceGuardOperationDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardOperationDetailResponseArrayOutput)
}

type ResourceGuardOperationDetailResponseOutput struct{ *pulumi.OutputState }

func (ResourceGuardOperationDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardOperationDetailResponse)(nil)).Elem()
}

func (o ResourceGuardOperationDetailResponseOutput) ToResourceGuardOperationDetailResponseOutput() ResourceGuardOperationDetailResponseOutput {
	return o
}

func (o ResourceGuardOperationDetailResponseOutput) ToResourceGuardOperationDetailResponseOutputWithContext(ctx context.Context) ResourceGuardOperationDetailResponseOutput {
	return o
}

func (o ResourceGuardOperationDetailResponseOutput) DefaultResourceRequest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardOperationDetailResponse) *string { return v.DefaultResourceRequest }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardOperationDetailResponseOutput) VaultCriticalOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardOperationDetailResponse) *string { return v.VaultCriticalOperation }).(pulumi.StringPtrOutput)
}

type ResourceGuardOperationDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceGuardOperationDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceGuardOperationDetailResponse)(nil)).Elem()
}

func (o ResourceGuardOperationDetailResponseArrayOutput) ToResourceGuardOperationDetailResponseArrayOutput() ResourceGuardOperationDetailResponseArrayOutput {
	return o
}

func (o ResourceGuardOperationDetailResponseArrayOutput) ToResourceGuardOperationDetailResponseArrayOutputWithContext(ctx context.Context) ResourceGuardOperationDetailResponseArrayOutput {
	return o
}

func (o ResourceGuardOperationDetailResponseArrayOutput) Index(i pulumi.IntInput) ResourceGuardOperationDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceGuardOperationDetailResponse {
		return vs[0].([]ResourceGuardOperationDetailResponse)[vs[1].(int)]
	}).(ResourceGuardOperationDetailResponseOutput)
}

type ResourceGuardProxyBaseResponse struct {
	Description                   *string                                `pulumi:"description"`
	LastUpdatedTime               *string                                `pulumi:"lastUpdatedTime"`
	ResourceGuardOperationDetails []ResourceGuardOperationDetailResponse `pulumi:"resourceGuardOperationDetails"`
	ResourceGuardResourceId       *string                                `pulumi:"resourceGuardResourceId"`
}





type ResourceGuardProxyBaseResponseInput interface {
	pulumi.Input

	ToResourceGuardProxyBaseResponseOutput() ResourceGuardProxyBaseResponseOutput
	ToResourceGuardProxyBaseResponseOutputWithContext(context.Context) ResourceGuardProxyBaseResponseOutput
}

type ResourceGuardProxyBaseResponseArgs struct {
	Description                   pulumi.StringPtrInput                          `pulumi:"description"`
	LastUpdatedTime               pulumi.StringPtrInput                          `pulumi:"lastUpdatedTime"`
	ResourceGuardOperationDetails ResourceGuardOperationDetailResponseArrayInput `pulumi:"resourceGuardOperationDetails"`
	ResourceGuardResourceId       pulumi.StringPtrInput                          `pulumi:"resourceGuardResourceId"`
}

func (ResourceGuardProxyBaseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardProxyBaseResponse)(nil)).Elem()
}

func (i ResourceGuardProxyBaseResponseArgs) ToResourceGuardProxyBaseResponseOutput() ResourceGuardProxyBaseResponseOutput {
	return i.ToResourceGuardProxyBaseResponseOutputWithContext(context.Background())
}

func (i ResourceGuardProxyBaseResponseArgs) ToResourceGuardProxyBaseResponseOutputWithContext(ctx context.Context) ResourceGuardProxyBaseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardProxyBaseResponseOutput)
}

func (i ResourceGuardProxyBaseResponseArgs) ToResourceGuardProxyBaseResponsePtrOutput() ResourceGuardProxyBaseResponsePtrOutput {
	return i.ToResourceGuardProxyBaseResponsePtrOutputWithContext(context.Background())
}

func (i ResourceGuardProxyBaseResponseArgs) ToResourceGuardProxyBaseResponsePtrOutputWithContext(ctx context.Context) ResourceGuardProxyBaseResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardProxyBaseResponseOutput).ToResourceGuardProxyBaseResponsePtrOutputWithContext(ctx)
}









type ResourceGuardProxyBaseResponsePtrInput interface {
	pulumi.Input

	ToResourceGuardProxyBaseResponsePtrOutput() ResourceGuardProxyBaseResponsePtrOutput
	ToResourceGuardProxyBaseResponsePtrOutputWithContext(context.Context) ResourceGuardProxyBaseResponsePtrOutput
}

type resourceGuardProxyBaseResponsePtrType ResourceGuardProxyBaseResponseArgs

func ResourceGuardProxyBaseResponsePtr(v *ResourceGuardProxyBaseResponseArgs) ResourceGuardProxyBaseResponsePtrInput {
	return (*resourceGuardProxyBaseResponsePtrType)(v)
}

func (*resourceGuardProxyBaseResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuardProxyBaseResponse)(nil)).Elem()
}

func (i *resourceGuardProxyBaseResponsePtrType) ToResourceGuardProxyBaseResponsePtrOutput() ResourceGuardProxyBaseResponsePtrOutput {
	return i.ToResourceGuardProxyBaseResponsePtrOutputWithContext(context.Background())
}

func (i *resourceGuardProxyBaseResponsePtrType) ToResourceGuardProxyBaseResponsePtrOutputWithContext(ctx context.Context) ResourceGuardProxyBaseResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardProxyBaseResponsePtrOutput)
}

type ResourceGuardProxyBaseResponseOutput struct{ *pulumi.OutputState }

func (ResourceGuardProxyBaseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardProxyBaseResponse)(nil)).Elem()
}

func (o ResourceGuardProxyBaseResponseOutput) ToResourceGuardProxyBaseResponseOutput() ResourceGuardProxyBaseResponseOutput {
	return o
}

func (o ResourceGuardProxyBaseResponseOutput) ToResourceGuardProxyBaseResponseOutputWithContext(ctx context.Context) ResourceGuardProxyBaseResponseOutput {
	return o
}

func (o ResourceGuardProxyBaseResponseOutput) ToResourceGuardProxyBaseResponsePtrOutput() ResourceGuardProxyBaseResponsePtrOutput {
	return o.ToResourceGuardProxyBaseResponsePtrOutputWithContext(context.Background())
}

func (o ResourceGuardProxyBaseResponseOutput) ToResourceGuardProxyBaseResponsePtrOutputWithContext(ctx context.Context) ResourceGuardProxyBaseResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceGuardProxyBaseResponse) *ResourceGuardProxyBaseResponse {
		return &v
	}).(ResourceGuardProxyBaseResponsePtrOutput)
}

func (o ResourceGuardProxyBaseResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBaseResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBaseResponseOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBaseResponse) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBaseResponseOutput) ResourceGuardOperationDetails() ResourceGuardOperationDetailResponseArrayOutput {
	return o.ApplyT(func(v ResourceGuardProxyBaseResponse) []ResourceGuardOperationDetailResponse {
		return v.ResourceGuardOperationDetails
	}).(ResourceGuardOperationDetailResponseArrayOutput)
}

func (o ResourceGuardProxyBaseResponseOutput) ResourceGuardResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBaseResponse) *string { return v.ResourceGuardResourceId }).(pulumi.StringPtrOutput)
}

type ResourceGuardProxyBaseResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceGuardProxyBaseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuardProxyBaseResponse)(nil)).Elem()
}

func (o ResourceGuardProxyBaseResponsePtrOutput) ToResourceGuardProxyBaseResponsePtrOutput() ResourceGuardProxyBaseResponsePtrOutput {
	return o
}

func (o ResourceGuardProxyBaseResponsePtrOutput) ToResourceGuardProxyBaseResponsePtrOutputWithContext(ctx context.Context) ResourceGuardProxyBaseResponsePtrOutput {
	return o
}

func (o ResourceGuardProxyBaseResponsePtrOutput) Elem() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBaseResponse) ResourceGuardProxyBaseResponse {
		if v != nil {
			return *v
		}
		var ret ResourceGuardProxyBaseResponse
		return ret
	}).(ResourceGuardProxyBaseResponseOutput)
}

func (o ResourceGuardProxyBaseResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBaseResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBaseResponsePtrOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBaseResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUpdatedTime
	}).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBaseResponsePtrOutput) ResourceGuardOperationDetails() ResourceGuardOperationDetailResponseArrayOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBaseResponse) []ResourceGuardOperationDetailResponse {
		if v == nil {
			return nil
		}
		return v.ResourceGuardOperationDetails
	}).(ResourceGuardOperationDetailResponseArrayOutput)
}

func (o ResourceGuardProxyBaseResponsePtrOutput) ResourceGuardResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBaseResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuardResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceHealthDetailsResponse struct {
	Code            int      `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           string   `pulumi:"title"`
}





type ResourceHealthDetailsResponseInput interface {
	pulumi.Input

	ToResourceHealthDetailsResponseOutput() ResourceHealthDetailsResponseOutput
	ToResourceHealthDetailsResponseOutputWithContext(context.Context) ResourceHealthDetailsResponseOutput
}

type ResourceHealthDetailsResponseArgs struct {
	Code            pulumi.IntInput         `pulumi:"code"`
	Message         pulumi.StringInput      `pulumi:"message"`
	Recommendations pulumi.StringArrayInput `pulumi:"recommendations"`
	Title           pulumi.StringInput      `pulumi:"title"`
}

func (ResourceHealthDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceHealthDetailsResponse)(nil)).Elem()
}

func (i ResourceHealthDetailsResponseArgs) ToResourceHealthDetailsResponseOutput() ResourceHealthDetailsResponseOutput {
	return i.ToResourceHealthDetailsResponseOutputWithContext(context.Background())
}

func (i ResourceHealthDetailsResponseArgs) ToResourceHealthDetailsResponseOutputWithContext(ctx context.Context) ResourceHealthDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceHealthDetailsResponseOutput)
}





type ResourceHealthDetailsResponseArrayInput interface {
	pulumi.Input

	ToResourceHealthDetailsResponseArrayOutput() ResourceHealthDetailsResponseArrayOutput
	ToResourceHealthDetailsResponseArrayOutputWithContext(context.Context) ResourceHealthDetailsResponseArrayOutput
}

type ResourceHealthDetailsResponseArray []ResourceHealthDetailsResponseInput

func (ResourceHealthDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceHealthDetailsResponse)(nil)).Elem()
}

func (i ResourceHealthDetailsResponseArray) ToResourceHealthDetailsResponseArrayOutput() ResourceHealthDetailsResponseArrayOutput {
	return i.ToResourceHealthDetailsResponseArrayOutputWithContext(context.Background())
}

func (i ResourceHealthDetailsResponseArray) ToResourceHealthDetailsResponseArrayOutputWithContext(ctx context.Context) ResourceHealthDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceHealthDetailsResponseArrayOutput)
}

type ResourceHealthDetailsResponseOutput struct{ *pulumi.OutputState }

func (ResourceHealthDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceHealthDetailsResponse)(nil)).Elem()
}

func (o ResourceHealthDetailsResponseOutput) ToResourceHealthDetailsResponseOutput() ResourceHealthDetailsResponseOutput {
	return o
}

func (o ResourceHealthDetailsResponseOutput) ToResourceHealthDetailsResponseOutputWithContext(ctx context.Context) ResourceHealthDetailsResponseOutput {
	return o
}

func (o ResourceHealthDetailsResponseOutput) Code() pulumi.IntOutput {
	return o.ApplyT(func(v ResourceHealthDetailsResponse) int { return v.Code }).(pulumi.IntOutput)
}

func (o ResourceHealthDetailsResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceHealthDetailsResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ResourceHealthDetailsResponseOutput) Recommendations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceHealthDetailsResponse) []string { return v.Recommendations }).(pulumi.StringArrayOutput)
}

func (o ResourceHealthDetailsResponseOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceHealthDetailsResponse) string { return v.Title }).(pulumi.StringOutput)
}

type ResourceHealthDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceHealthDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceHealthDetailsResponse)(nil)).Elem()
}

func (o ResourceHealthDetailsResponseArrayOutput) ToResourceHealthDetailsResponseArrayOutput() ResourceHealthDetailsResponseArrayOutput {
	return o
}

func (o ResourceHealthDetailsResponseArrayOutput) ToResourceHealthDetailsResponseArrayOutputWithContext(ctx context.Context) ResourceHealthDetailsResponseArrayOutput {
	return o
}

func (o ResourceHealthDetailsResponseArrayOutput) Index(i pulumi.IntInput) ResourceHealthDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceHealthDetailsResponse {
		return vs[0].([]ResourceHealthDetailsResponse)[vs[1].(int)]
	}).(ResourceHealthDetailsResponseOutput)
}

type RetentionDuration struct {
	Count        *int    `pulumi:"count"`
	DurationType *string `pulumi:"durationType"`
}





type RetentionDurationInput interface {
	pulumi.Input

	ToRetentionDurationOutput() RetentionDurationOutput
	ToRetentionDurationOutputWithContext(context.Context) RetentionDurationOutput
}

type RetentionDurationArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	DurationType pulumi.StringPtrInput `pulumi:"durationType"`
}

func (RetentionDurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDuration)(nil)).Elem()
}

func (i RetentionDurationArgs) ToRetentionDurationOutput() RetentionDurationOutput {
	return i.ToRetentionDurationOutputWithContext(context.Background())
}

func (i RetentionDurationArgs) ToRetentionDurationOutputWithContext(ctx context.Context) RetentionDurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationOutput)
}

func (i RetentionDurationArgs) ToRetentionDurationPtrOutput() RetentionDurationPtrOutput {
	return i.ToRetentionDurationPtrOutputWithContext(context.Background())
}

func (i RetentionDurationArgs) ToRetentionDurationPtrOutputWithContext(ctx context.Context) RetentionDurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationOutput).ToRetentionDurationPtrOutputWithContext(ctx)
}









type RetentionDurationPtrInput interface {
	pulumi.Input

	ToRetentionDurationPtrOutput() RetentionDurationPtrOutput
	ToRetentionDurationPtrOutputWithContext(context.Context) RetentionDurationPtrOutput
}

type retentionDurationPtrType RetentionDurationArgs

func RetentionDurationPtr(v *RetentionDurationArgs) RetentionDurationPtrInput {
	return (*retentionDurationPtrType)(v)
}

func (*retentionDurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDuration)(nil)).Elem()
}

func (i *retentionDurationPtrType) ToRetentionDurationPtrOutput() RetentionDurationPtrOutput {
	return i.ToRetentionDurationPtrOutputWithContext(context.Background())
}

func (i *retentionDurationPtrType) ToRetentionDurationPtrOutputWithContext(ctx context.Context) RetentionDurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationPtrOutput)
}

type RetentionDurationOutput struct{ *pulumi.OutputState }

func (RetentionDurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDuration)(nil)).Elem()
}

func (o RetentionDurationOutput) ToRetentionDurationOutput() RetentionDurationOutput {
	return o
}

func (o RetentionDurationOutput) ToRetentionDurationOutputWithContext(ctx context.Context) RetentionDurationOutput {
	return o
}

func (o RetentionDurationOutput) ToRetentionDurationPtrOutput() RetentionDurationPtrOutput {
	return o.ToRetentionDurationPtrOutputWithContext(context.Background())
}

func (o RetentionDurationOutput) ToRetentionDurationPtrOutputWithContext(ctx context.Context) RetentionDurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionDuration) *RetentionDuration {
		return &v
	}).(RetentionDurationPtrOutput)
}

func (o RetentionDurationOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionDuration) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RetentionDurationOutput) DurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionDuration) *string { return v.DurationType }).(pulumi.StringPtrOutput)
}

type RetentionDurationPtrOutput struct{ *pulumi.OutputState }

func (RetentionDurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDuration)(nil)).Elem()
}

func (o RetentionDurationPtrOutput) ToRetentionDurationPtrOutput() RetentionDurationPtrOutput {
	return o
}

func (o RetentionDurationPtrOutput) ToRetentionDurationPtrOutputWithContext(ctx context.Context) RetentionDurationPtrOutput {
	return o
}

func (o RetentionDurationPtrOutput) Elem() RetentionDurationOutput {
	return o.ApplyT(func(v *RetentionDuration) RetentionDuration {
		if v != nil {
			return *v
		}
		var ret RetentionDuration
		return ret
	}).(RetentionDurationOutput)
}

func (o RetentionDurationPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionDuration) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RetentionDurationPtrOutput) DurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionDuration) *string {
		if v == nil {
			return nil
		}
		return v.DurationType
	}).(pulumi.StringPtrOutput)
}

type RetentionDurationResponse struct {
	Count        *int    `pulumi:"count"`
	DurationType *string `pulumi:"durationType"`
}





type RetentionDurationResponseInput interface {
	pulumi.Input

	ToRetentionDurationResponseOutput() RetentionDurationResponseOutput
	ToRetentionDurationResponseOutputWithContext(context.Context) RetentionDurationResponseOutput
}

type RetentionDurationResponseArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	DurationType pulumi.StringPtrInput `pulumi:"durationType"`
}

func (RetentionDurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDurationResponse)(nil)).Elem()
}

func (i RetentionDurationResponseArgs) ToRetentionDurationResponseOutput() RetentionDurationResponseOutput {
	return i.ToRetentionDurationResponseOutputWithContext(context.Background())
}

func (i RetentionDurationResponseArgs) ToRetentionDurationResponseOutputWithContext(ctx context.Context) RetentionDurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationResponseOutput)
}

func (i RetentionDurationResponseArgs) ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput {
	return i.ToRetentionDurationResponsePtrOutputWithContext(context.Background())
}

func (i RetentionDurationResponseArgs) ToRetentionDurationResponsePtrOutputWithContext(ctx context.Context) RetentionDurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationResponseOutput).ToRetentionDurationResponsePtrOutputWithContext(ctx)
}









type RetentionDurationResponsePtrInput interface {
	pulumi.Input

	ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput
	ToRetentionDurationResponsePtrOutputWithContext(context.Context) RetentionDurationResponsePtrOutput
}

type retentionDurationResponsePtrType RetentionDurationResponseArgs

func RetentionDurationResponsePtr(v *RetentionDurationResponseArgs) RetentionDurationResponsePtrInput {
	return (*retentionDurationResponsePtrType)(v)
}

func (*retentionDurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDurationResponse)(nil)).Elem()
}

func (i *retentionDurationResponsePtrType) ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput {
	return i.ToRetentionDurationResponsePtrOutputWithContext(context.Background())
}

func (i *retentionDurationResponsePtrType) ToRetentionDurationResponsePtrOutputWithContext(ctx context.Context) RetentionDurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationResponsePtrOutput)
}

type RetentionDurationResponseOutput struct{ *pulumi.OutputState }

func (RetentionDurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDurationResponse)(nil)).Elem()
}

func (o RetentionDurationResponseOutput) ToRetentionDurationResponseOutput() RetentionDurationResponseOutput {
	return o
}

func (o RetentionDurationResponseOutput) ToRetentionDurationResponseOutputWithContext(ctx context.Context) RetentionDurationResponseOutput {
	return o
}

func (o RetentionDurationResponseOutput) ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput {
	return o.ToRetentionDurationResponsePtrOutputWithContext(context.Background())
}

func (o RetentionDurationResponseOutput) ToRetentionDurationResponsePtrOutputWithContext(ctx context.Context) RetentionDurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionDurationResponse) *RetentionDurationResponse {
		return &v
	}).(RetentionDurationResponsePtrOutput)
}

func (o RetentionDurationResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionDurationResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RetentionDurationResponseOutput) DurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionDurationResponse) *string { return v.DurationType }).(pulumi.StringPtrOutput)
}

type RetentionDurationResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionDurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDurationResponse)(nil)).Elem()
}

func (o RetentionDurationResponsePtrOutput) ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput {
	return o
}

func (o RetentionDurationResponsePtrOutput) ToRetentionDurationResponsePtrOutputWithContext(ctx context.Context) RetentionDurationResponsePtrOutput {
	return o
}

func (o RetentionDurationResponsePtrOutput) Elem() RetentionDurationResponseOutput {
	return o.ApplyT(func(v *RetentionDurationResponse) RetentionDurationResponse {
		if v != nil {
			return *v
		}
		var ret RetentionDurationResponse
		return ret
	}).(RetentionDurationResponseOutput)
}

func (o RetentionDurationResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionDurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RetentionDurationResponsePtrOutput) DurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionDurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DurationType
	}).(pulumi.StringPtrOutput)
}

type Settings struct {
	IsCompression    *bool   `pulumi:"isCompression"`
	Issqlcompression *bool   `pulumi:"issqlcompression"`
	TimeZone         *string `pulumi:"timeZone"`
}





type SettingsInput interface {
	pulumi.Input

	ToSettingsOutput() SettingsOutput
	ToSettingsOutputWithContext(context.Context) SettingsOutput
}

type SettingsArgs struct {
	IsCompression    pulumi.BoolPtrInput   `pulumi:"isCompression"`
	Issqlcompression pulumi.BoolPtrInput   `pulumi:"issqlcompression"`
	TimeZone         pulumi.StringPtrInput `pulumi:"timeZone"`
}

func (SettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Settings)(nil)).Elem()
}

func (i SettingsArgs) ToSettingsOutput() SettingsOutput {
	return i.ToSettingsOutputWithContext(context.Background())
}

func (i SettingsArgs) ToSettingsOutputWithContext(ctx context.Context) SettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsOutput)
}

func (i SettingsArgs) ToSettingsPtrOutput() SettingsPtrOutput {
	return i.ToSettingsPtrOutputWithContext(context.Background())
}

func (i SettingsArgs) ToSettingsPtrOutputWithContext(ctx context.Context) SettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsOutput).ToSettingsPtrOutputWithContext(ctx)
}









type SettingsPtrInput interface {
	pulumi.Input

	ToSettingsPtrOutput() SettingsPtrOutput
	ToSettingsPtrOutputWithContext(context.Context) SettingsPtrOutput
}

type settingsPtrType SettingsArgs

func SettingsPtr(v *SettingsArgs) SettingsPtrInput {
	return (*settingsPtrType)(v)
}

func (*settingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Settings)(nil)).Elem()
}

func (i *settingsPtrType) ToSettingsPtrOutput() SettingsPtrOutput {
	return i.ToSettingsPtrOutputWithContext(context.Background())
}

func (i *settingsPtrType) ToSettingsPtrOutputWithContext(ctx context.Context) SettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsPtrOutput)
}

type SettingsOutput struct{ *pulumi.OutputState }

func (SettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Settings)(nil)).Elem()
}

func (o SettingsOutput) ToSettingsOutput() SettingsOutput {
	return o
}

func (o SettingsOutput) ToSettingsOutputWithContext(ctx context.Context) SettingsOutput {
	return o
}

func (o SettingsOutput) ToSettingsPtrOutput() SettingsPtrOutput {
	return o.ToSettingsPtrOutputWithContext(context.Background())
}

func (o SettingsOutput) ToSettingsPtrOutputWithContext(ctx context.Context) SettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Settings) *Settings {
		return &v
	}).(SettingsPtrOutput)
}

func (o SettingsOutput) IsCompression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Settings) *bool { return v.IsCompression }).(pulumi.BoolPtrOutput)
}

func (o SettingsOutput) Issqlcompression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Settings) *bool { return v.Issqlcompression }).(pulumi.BoolPtrOutput)
}

func (o SettingsOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Settings) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type SettingsPtrOutput struct{ *pulumi.OutputState }

func (SettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Settings)(nil)).Elem()
}

func (o SettingsPtrOutput) ToSettingsPtrOutput() SettingsPtrOutput {
	return o
}

func (o SettingsPtrOutput) ToSettingsPtrOutputWithContext(ctx context.Context) SettingsPtrOutput {
	return o
}

func (o SettingsPtrOutput) Elem() SettingsOutput {
	return o.ApplyT(func(v *Settings) Settings {
		if v != nil {
			return *v
		}
		var ret Settings
		return ret
	}).(SettingsOutput)
}

func (o SettingsPtrOutput) IsCompression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Settings) *bool {
		if v == nil {
			return nil
		}
		return v.IsCompression
	}).(pulumi.BoolPtrOutput)
}

func (o SettingsPtrOutput) Issqlcompression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Settings) *bool {
		if v == nil {
			return nil
		}
		return v.Issqlcompression
	}).(pulumi.BoolPtrOutput)
}

func (o SettingsPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Settings) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type SettingsResponse struct {
	IsCompression    *bool   `pulumi:"isCompression"`
	Issqlcompression *bool   `pulumi:"issqlcompression"`
	TimeZone         *string `pulumi:"timeZone"`
}





type SettingsResponseInput interface {
	pulumi.Input

	ToSettingsResponseOutput() SettingsResponseOutput
	ToSettingsResponseOutputWithContext(context.Context) SettingsResponseOutput
}

type SettingsResponseArgs struct {
	IsCompression    pulumi.BoolPtrInput   `pulumi:"isCompression"`
	Issqlcompression pulumi.BoolPtrInput   `pulumi:"issqlcompression"`
	TimeZone         pulumi.StringPtrInput `pulumi:"timeZone"`
}

func (SettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsResponse)(nil)).Elem()
}

func (i SettingsResponseArgs) ToSettingsResponseOutput() SettingsResponseOutput {
	return i.ToSettingsResponseOutputWithContext(context.Background())
}

func (i SettingsResponseArgs) ToSettingsResponseOutputWithContext(ctx context.Context) SettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsResponseOutput)
}

func (i SettingsResponseArgs) ToSettingsResponsePtrOutput() SettingsResponsePtrOutput {
	return i.ToSettingsResponsePtrOutputWithContext(context.Background())
}

func (i SettingsResponseArgs) ToSettingsResponsePtrOutputWithContext(ctx context.Context) SettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsResponseOutput).ToSettingsResponsePtrOutputWithContext(ctx)
}









type SettingsResponsePtrInput interface {
	pulumi.Input

	ToSettingsResponsePtrOutput() SettingsResponsePtrOutput
	ToSettingsResponsePtrOutputWithContext(context.Context) SettingsResponsePtrOutput
}

type settingsResponsePtrType SettingsResponseArgs

func SettingsResponsePtr(v *SettingsResponseArgs) SettingsResponsePtrInput {
	return (*settingsResponsePtrType)(v)
}

func (*settingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingsResponse)(nil)).Elem()
}

func (i *settingsResponsePtrType) ToSettingsResponsePtrOutput() SettingsResponsePtrOutput {
	return i.ToSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *settingsResponsePtrType) ToSettingsResponsePtrOutputWithContext(ctx context.Context) SettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsResponsePtrOutput)
}

type SettingsResponseOutput struct{ *pulumi.OutputState }

func (SettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsResponse)(nil)).Elem()
}

func (o SettingsResponseOutput) ToSettingsResponseOutput() SettingsResponseOutput {
	return o
}

func (o SettingsResponseOutput) ToSettingsResponseOutputWithContext(ctx context.Context) SettingsResponseOutput {
	return o
}

func (o SettingsResponseOutput) ToSettingsResponsePtrOutput() SettingsResponsePtrOutput {
	return o.ToSettingsResponsePtrOutputWithContext(context.Background())
}

func (o SettingsResponseOutput) ToSettingsResponsePtrOutputWithContext(ctx context.Context) SettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SettingsResponse) *SettingsResponse {
		return &v
	}).(SettingsResponsePtrOutput)
}

func (o SettingsResponseOutput) IsCompression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SettingsResponse) *bool { return v.IsCompression }).(pulumi.BoolPtrOutput)
}

func (o SettingsResponseOutput) Issqlcompression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SettingsResponse) *bool { return v.Issqlcompression }).(pulumi.BoolPtrOutput)
}

func (o SettingsResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingsResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type SettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingsResponse)(nil)).Elem()
}

func (o SettingsResponsePtrOutput) ToSettingsResponsePtrOutput() SettingsResponsePtrOutput {
	return o
}

func (o SettingsResponsePtrOutput) ToSettingsResponsePtrOutputWithContext(ctx context.Context) SettingsResponsePtrOutput {
	return o
}

func (o SettingsResponsePtrOutput) Elem() SettingsResponseOutput {
	return o.ApplyT(func(v *SettingsResponse) SettingsResponse {
		if v != nil {
			return *v
		}
		var ret SettingsResponse
		return ret
	}).(SettingsResponseOutput)
}

func (o SettingsResponsePtrOutput) IsCompression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCompression
	}).(pulumi.BoolPtrOutput)
}

func (o SettingsResponsePtrOutput) Issqlcompression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Issqlcompression
	}).(pulumi.BoolPtrOutput)
}

func (o SettingsResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type SimpleRetentionPolicy struct {
	RetentionDuration   *RetentionDuration `pulumi:"retentionDuration"`
	RetentionPolicyType string             `pulumi:"retentionPolicyType"`
}





type SimpleRetentionPolicyInput interface {
	pulumi.Input

	ToSimpleRetentionPolicyOutput() SimpleRetentionPolicyOutput
	ToSimpleRetentionPolicyOutputWithContext(context.Context) SimpleRetentionPolicyOutput
}

type SimpleRetentionPolicyArgs struct {
	RetentionDuration   RetentionDurationPtrInput `pulumi:"retentionDuration"`
	RetentionPolicyType pulumi.StringInput        `pulumi:"retentionPolicyType"`
}

func (SimpleRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleRetentionPolicy)(nil)).Elem()
}

func (i SimpleRetentionPolicyArgs) ToSimpleRetentionPolicyOutput() SimpleRetentionPolicyOutput {
	return i.ToSimpleRetentionPolicyOutputWithContext(context.Background())
}

func (i SimpleRetentionPolicyArgs) ToSimpleRetentionPolicyOutputWithContext(ctx context.Context) SimpleRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleRetentionPolicyOutput)
}

type SimpleRetentionPolicyOutput struct{ *pulumi.OutputState }

func (SimpleRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleRetentionPolicy)(nil)).Elem()
}

func (o SimpleRetentionPolicyOutput) ToSimpleRetentionPolicyOutput() SimpleRetentionPolicyOutput {
	return o
}

func (o SimpleRetentionPolicyOutput) ToSimpleRetentionPolicyOutputWithContext(ctx context.Context) SimpleRetentionPolicyOutput {
	return o
}

func (o SimpleRetentionPolicyOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v SimpleRetentionPolicy) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o SimpleRetentionPolicyOutput) RetentionPolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v SimpleRetentionPolicy) string { return v.RetentionPolicyType }).(pulumi.StringOutput)
}

type SimpleRetentionPolicyResponse struct {
	RetentionDuration   *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionPolicyType string                     `pulumi:"retentionPolicyType"`
}





type SimpleRetentionPolicyResponseInput interface {
	pulumi.Input

	ToSimpleRetentionPolicyResponseOutput() SimpleRetentionPolicyResponseOutput
	ToSimpleRetentionPolicyResponseOutputWithContext(context.Context) SimpleRetentionPolicyResponseOutput
}

type SimpleRetentionPolicyResponseArgs struct {
	RetentionDuration   RetentionDurationResponsePtrInput `pulumi:"retentionDuration"`
	RetentionPolicyType pulumi.StringInput                `pulumi:"retentionPolicyType"`
}

func (SimpleRetentionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleRetentionPolicyResponse)(nil)).Elem()
}

func (i SimpleRetentionPolicyResponseArgs) ToSimpleRetentionPolicyResponseOutput() SimpleRetentionPolicyResponseOutput {
	return i.ToSimpleRetentionPolicyResponseOutputWithContext(context.Background())
}

func (i SimpleRetentionPolicyResponseArgs) ToSimpleRetentionPolicyResponseOutputWithContext(ctx context.Context) SimpleRetentionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleRetentionPolicyResponseOutput)
}

type SimpleRetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (SimpleRetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleRetentionPolicyResponse)(nil)).Elem()
}

func (o SimpleRetentionPolicyResponseOutput) ToSimpleRetentionPolicyResponseOutput() SimpleRetentionPolicyResponseOutput {
	return o
}

func (o SimpleRetentionPolicyResponseOutput) ToSimpleRetentionPolicyResponseOutputWithContext(ctx context.Context) SimpleRetentionPolicyResponseOutput {
	return o
}

func (o SimpleRetentionPolicyResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v SimpleRetentionPolicyResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o SimpleRetentionPolicyResponseOutput) RetentionPolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v SimpleRetentionPolicyResponse) string { return v.RetentionPolicyType }).(pulumi.StringOutput)
}

type SimpleSchedulePolicy struct {
	HourlySchedule          *HourlySchedule `pulumi:"hourlySchedule"`
	SchedulePolicyType      string          `pulumi:"schedulePolicyType"`
	ScheduleRunDays         []DayOfWeek     `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    *string         `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        []string        `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency *int            `pulumi:"scheduleWeeklyFrequency"`
}





type SimpleSchedulePolicyInput interface {
	pulumi.Input

	ToSimpleSchedulePolicyOutput() SimpleSchedulePolicyOutput
	ToSimpleSchedulePolicyOutputWithContext(context.Context) SimpleSchedulePolicyOutput
}

type SimpleSchedulePolicyArgs struct {
	HourlySchedule          HourlySchedulePtrInput  `pulumi:"hourlySchedule"`
	SchedulePolicyType      pulumi.StringInput      `pulumi:"schedulePolicyType"`
	ScheduleRunDays         DayOfWeekArrayInput     `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    pulumi.StringPtrInput   `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        pulumi.StringArrayInput `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency pulumi.IntPtrInput      `pulumi:"scheduleWeeklyFrequency"`
}

func (SimpleSchedulePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleSchedulePolicy)(nil)).Elem()
}

func (i SimpleSchedulePolicyArgs) ToSimpleSchedulePolicyOutput() SimpleSchedulePolicyOutput {
	return i.ToSimpleSchedulePolicyOutputWithContext(context.Background())
}

func (i SimpleSchedulePolicyArgs) ToSimpleSchedulePolicyOutputWithContext(ctx context.Context) SimpleSchedulePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleSchedulePolicyOutput)
}

type SimpleSchedulePolicyOutput struct{ *pulumi.OutputState }

func (SimpleSchedulePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleSchedulePolicy)(nil)).Elem()
}

func (o SimpleSchedulePolicyOutput) ToSimpleSchedulePolicyOutput() SimpleSchedulePolicyOutput {
	return o
}

func (o SimpleSchedulePolicyOutput) ToSimpleSchedulePolicyOutputWithContext(ctx context.Context) SimpleSchedulePolicyOutput {
	return o
}

func (o SimpleSchedulePolicyOutput) HourlySchedule() HourlySchedulePtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) *HourlySchedule { return v.HourlySchedule }).(HourlySchedulePtrOutput)
}

func (o SimpleSchedulePolicyOutput) SchedulePolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) string { return v.SchedulePolicyType }).(pulumi.StringOutput)
}

func (o SimpleSchedulePolicyOutput) ScheduleRunDays() DayOfWeekArrayOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) []DayOfWeek { return v.ScheduleRunDays }).(DayOfWeekArrayOutput)
}

func (o SimpleSchedulePolicyOutput) ScheduleRunFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) *string { return v.ScheduleRunFrequency }).(pulumi.StringPtrOutput)
}

func (o SimpleSchedulePolicyOutput) ScheduleRunTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) []string { return v.ScheduleRunTimes }).(pulumi.StringArrayOutput)
}

func (o SimpleSchedulePolicyOutput) ScheduleWeeklyFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) *int { return v.ScheduleWeeklyFrequency }).(pulumi.IntPtrOutput)
}

type SimpleSchedulePolicyResponse struct {
	HourlySchedule          *HourlyScheduleResponse `pulumi:"hourlySchedule"`
	SchedulePolicyType      string                  `pulumi:"schedulePolicyType"`
	ScheduleRunDays         []string                `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    *string                 `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        []string                `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency *int                    `pulumi:"scheduleWeeklyFrequency"`
}





type SimpleSchedulePolicyResponseInput interface {
	pulumi.Input

	ToSimpleSchedulePolicyResponseOutput() SimpleSchedulePolicyResponseOutput
	ToSimpleSchedulePolicyResponseOutputWithContext(context.Context) SimpleSchedulePolicyResponseOutput
}

type SimpleSchedulePolicyResponseArgs struct {
	HourlySchedule          HourlyScheduleResponsePtrInput `pulumi:"hourlySchedule"`
	SchedulePolicyType      pulumi.StringInput             `pulumi:"schedulePolicyType"`
	ScheduleRunDays         pulumi.StringArrayInput        `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    pulumi.StringPtrInput          `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        pulumi.StringArrayInput        `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency pulumi.IntPtrInput             `pulumi:"scheduleWeeklyFrequency"`
}

func (SimpleSchedulePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleSchedulePolicyResponse)(nil)).Elem()
}

func (i SimpleSchedulePolicyResponseArgs) ToSimpleSchedulePolicyResponseOutput() SimpleSchedulePolicyResponseOutput {
	return i.ToSimpleSchedulePolicyResponseOutputWithContext(context.Background())
}

func (i SimpleSchedulePolicyResponseArgs) ToSimpleSchedulePolicyResponseOutputWithContext(ctx context.Context) SimpleSchedulePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleSchedulePolicyResponseOutput)
}

type SimpleSchedulePolicyResponseOutput struct{ *pulumi.OutputState }

func (SimpleSchedulePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleSchedulePolicyResponse)(nil)).Elem()
}

func (o SimpleSchedulePolicyResponseOutput) ToSimpleSchedulePolicyResponseOutput() SimpleSchedulePolicyResponseOutput {
	return o
}

func (o SimpleSchedulePolicyResponseOutput) ToSimpleSchedulePolicyResponseOutputWithContext(ctx context.Context) SimpleSchedulePolicyResponseOutput {
	return o
}

func (o SimpleSchedulePolicyResponseOutput) HourlySchedule() HourlyScheduleResponsePtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) *HourlyScheduleResponse { return v.HourlySchedule }).(HourlyScheduleResponsePtrOutput)
}

func (o SimpleSchedulePolicyResponseOutput) SchedulePolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) string { return v.SchedulePolicyType }).(pulumi.StringOutput)
}

func (o SimpleSchedulePolicyResponseOutput) ScheduleRunDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) []string { return v.ScheduleRunDays }).(pulumi.StringArrayOutput)
}

func (o SimpleSchedulePolicyResponseOutput) ScheduleRunFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) *string { return v.ScheduleRunFrequency }).(pulumi.StringPtrOutput)
}

func (o SimpleSchedulePolicyResponseOutput) ScheduleRunTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) []string { return v.ScheduleRunTimes }).(pulumi.StringArrayOutput)
}

func (o SimpleSchedulePolicyResponseOutput) ScheduleWeeklyFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) *int { return v.ScheduleWeeklyFrequency }).(pulumi.IntPtrOutput)
}

type SubProtectionPolicy struct {
	PolicyType      *string     `pulumi:"policyType"`
	RetentionPolicy interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy  interface{} `pulumi:"schedulePolicy"`
}





type SubProtectionPolicyInput interface {
	pulumi.Input

	ToSubProtectionPolicyOutput() SubProtectionPolicyOutput
	ToSubProtectionPolicyOutputWithContext(context.Context) SubProtectionPolicyOutput
}

type SubProtectionPolicyArgs struct {
	PolicyType      pulumi.StringPtrInput `pulumi:"policyType"`
	RetentionPolicy pulumi.Input          `pulumi:"retentionPolicy"`
	SchedulePolicy  pulumi.Input          `pulumi:"schedulePolicy"`
}

func (SubProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubProtectionPolicy)(nil)).Elem()
}

func (i SubProtectionPolicyArgs) ToSubProtectionPolicyOutput() SubProtectionPolicyOutput {
	return i.ToSubProtectionPolicyOutputWithContext(context.Background())
}

func (i SubProtectionPolicyArgs) ToSubProtectionPolicyOutputWithContext(ctx context.Context) SubProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubProtectionPolicyOutput)
}





type SubProtectionPolicyArrayInput interface {
	pulumi.Input

	ToSubProtectionPolicyArrayOutput() SubProtectionPolicyArrayOutput
	ToSubProtectionPolicyArrayOutputWithContext(context.Context) SubProtectionPolicyArrayOutput
}

type SubProtectionPolicyArray []SubProtectionPolicyInput

func (SubProtectionPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubProtectionPolicy)(nil)).Elem()
}

func (i SubProtectionPolicyArray) ToSubProtectionPolicyArrayOutput() SubProtectionPolicyArrayOutput {
	return i.ToSubProtectionPolicyArrayOutputWithContext(context.Background())
}

func (i SubProtectionPolicyArray) ToSubProtectionPolicyArrayOutputWithContext(ctx context.Context) SubProtectionPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubProtectionPolicyArrayOutput)
}

type SubProtectionPolicyOutput struct{ *pulumi.OutputState }

func (SubProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubProtectionPolicy)(nil)).Elem()
}

func (o SubProtectionPolicyOutput) ToSubProtectionPolicyOutput() SubProtectionPolicyOutput {
	return o
}

func (o SubProtectionPolicyOutput) ToSubProtectionPolicyOutputWithContext(ctx context.Context) SubProtectionPolicyOutput {
	return o
}

func (o SubProtectionPolicyOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubProtectionPolicy) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o SubProtectionPolicyOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v SubProtectionPolicy) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o SubProtectionPolicyOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v SubProtectionPolicy) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

type SubProtectionPolicyArrayOutput struct{ *pulumi.OutputState }

func (SubProtectionPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubProtectionPolicy)(nil)).Elem()
}

func (o SubProtectionPolicyArrayOutput) ToSubProtectionPolicyArrayOutput() SubProtectionPolicyArrayOutput {
	return o
}

func (o SubProtectionPolicyArrayOutput) ToSubProtectionPolicyArrayOutputWithContext(ctx context.Context) SubProtectionPolicyArrayOutput {
	return o
}

func (o SubProtectionPolicyArrayOutput) Index(i pulumi.IntInput) SubProtectionPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubProtectionPolicy {
		return vs[0].([]SubProtectionPolicy)[vs[1].(int)]
	}).(SubProtectionPolicyOutput)
}

type SubProtectionPolicyResponse struct {
	PolicyType      *string     `pulumi:"policyType"`
	RetentionPolicy interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy  interface{} `pulumi:"schedulePolicy"`
}





type SubProtectionPolicyResponseInput interface {
	pulumi.Input

	ToSubProtectionPolicyResponseOutput() SubProtectionPolicyResponseOutput
	ToSubProtectionPolicyResponseOutputWithContext(context.Context) SubProtectionPolicyResponseOutput
}

type SubProtectionPolicyResponseArgs struct {
	PolicyType      pulumi.StringPtrInput `pulumi:"policyType"`
	RetentionPolicy pulumi.Input          `pulumi:"retentionPolicy"`
	SchedulePolicy  pulumi.Input          `pulumi:"schedulePolicy"`
}

func (SubProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubProtectionPolicyResponse)(nil)).Elem()
}

func (i SubProtectionPolicyResponseArgs) ToSubProtectionPolicyResponseOutput() SubProtectionPolicyResponseOutput {
	return i.ToSubProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i SubProtectionPolicyResponseArgs) ToSubProtectionPolicyResponseOutputWithContext(ctx context.Context) SubProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubProtectionPolicyResponseOutput)
}





type SubProtectionPolicyResponseArrayInput interface {
	pulumi.Input

	ToSubProtectionPolicyResponseArrayOutput() SubProtectionPolicyResponseArrayOutput
	ToSubProtectionPolicyResponseArrayOutputWithContext(context.Context) SubProtectionPolicyResponseArrayOutput
}

type SubProtectionPolicyResponseArray []SubProtectionPolicyResponseInput

func (SubProtectionPolicyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubProtectionPolicyResponse)(nil)).Elem()
}

func (i SubProtectionPolicyResponseArray) ToSubProtectionPolicyResponseArrayOutput() SubProtectionPolicyResponseArrayOutput {
	return i.ToSubProtectionPolicyResponseArrayOutputWithContext(context.Background())
}

func (i SubProtectionPolicyResponseArray) ToSubProtectionPolicyResponseArrayOutputWithContext(ctx context.Context) SubProtectionPolicyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubProtectionPolicyResponseArrayOutput)
}

type SubProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (SubProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubProtectionPolicyResponse)(nil)).Elem()
}

func (o SubProtectionPolicyResponseOutput) ToSubProtectionPolicyResponseOutput() SubProtectionPolicyResponseOutput {
	return o
}

func (o SubProtectionPolicyResponseOutput) ToSubProtectionPolicyResponseOutputWithContext(ctx context.Context) SubProtectionPolicyResponseOutput {
	return o
}

func (o SubProtectionPolicyResponseOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubProtectionPolicyResponse) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o SubProtectionPolicyResponseOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v SubProtectionPolicyResponse) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o SubProtectionPolicyResponseOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v SubProtectionPolicyResponse) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

type SubProtectionPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (SubProtectionPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubProtectionPolicyResponse)(nil)).Elem()
}

func (o SubProtectionPolicyResponseArrayOutput) ToSubProtectionPolicyResponseArrayOutput() SubProtectionPolicyResponseArrayOutput {
	return o
}

func (o SubProtectionPolicyResponseArrayOutput) ToSubProtectionPolicyResponseArrayOutputWithContext(ctx context.Context) SubProtectionPolicyResponseArrayOutput {
	return o
}

func (o SubProtectionPolicyResponseArrayOutput) Index(i pulumi.IntInput) SubProtectionPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubProtectionPolicyResponse {
		return vs[0].([]SubProtectionPolicyResponse)[vs[1].(int)]
	}).(SubProtectionPolicyResponseOutput)
}

type WeeklyRetentionFormat struct {
	DaysOfTheWeek   []DayOfWeek   `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth []WeekOfMonth `pulumi:"weeksOfTheMonth"`
}





type WeeklyRetentionFormatInput interface {
	pulumi.Input

	ToWeeklyRetentionFormatOutput() WeeklyRetentionFormatOutput
	ToWeeklyRetentionFormatOutputWithContext(context.Context) WeeklyRetentionFormatOutput
}

type WeeklyRetentionFormatArgs struct {
	DaysOfTheWeek   DayOfWeekArrayInput   `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth WeekOfMonthArrayInput `pulumi:"weeksOfTheMonth"`
}

func (WeeklyRetentionFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionFormat)(nil)).Elem()
}

func (i WeeklyRetentionFormatArgs) ToWeeklyRetentionFormatOutput() WeeklyRetentionFormatOutput {
	return i.ToWeeklyRetentionFormatOutputWithContext(context.Background())
}

func (i WeeklyRetentionFormatArgs) ToWeeklyRetentionFormatOutputWithContext(ctx context.Context) WeeklyRetentionFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatOutput)
}

func (i WeeklyRetentionFormatArgs) ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput {
	return i.ToWeeklyRetentionFormatPtrOutputWithContext(context.Background())
}

func (i WeeklyRetentionFormatArgs) ToWeeklyRetentionFormatPtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatOutput).ToWeeklyRetentionFormatPtrOutputWithContext(ctx)
}









type WeeklyRetentionFormatPtrInput interface {
	pulumi.Input

	ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput
	ToWeeklyRetentionFormatPtrOutputWithContext(context.Context) WeeklyRetentionFormatPtrOutput
}

type weeklyRetentionFormatPtrType WeeklyRetentionFormatArgs

func WeeklyRetentionFormatPtr(v *WeeklyRetentionFormatArgs) WeeklyRetentionFormatPtrInput {
	return (*weeklyRetentionFormatPtrType)(v)
}

func (*weeklyRetentionFormatPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionFormat)(nil)).Elem()
}

func (i *weeklyRetentionFormatPtrType) ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput {
	return i.ToWeeklyRetentionFormatPtrOutputWithContext(context.Background())
}

func (i *weeklyRetentionFormatPtrType) ToWeeklyRetentionFormatPtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatPtrOutput)
}

type WeeklyRetentionFormatOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionFormat)(nil)).Elem()
}

func (o WeeklyRetentionFormatOutput) ToWeeklyRetentionFormatOutput() WeeklyRetentionFormatOutput {
	return o
}

func (o WeeklyRetentionFormatOutput) ToWeeklyRetentionFormatOutputWithContext(ctx context.Context) WeeklyRetentionFormatOutput {
	return o
}

func (o WeeklyRetentionFormatOutput) ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput {
	return o.ToWeeklyRetentionFormatPtrOutputWithContext(context.Background())
}

func (o WeeklyRetentionFormatOutput) ToWeeklyRetentionFormatPtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyRetentionFormat) *WeeklyRetentionFormat {
		return &v
	}).(WeeklyRetentionFormatPtrOutput)
}

func (o WeeklyRetentionFormatOutput) DaysOfTheWeek() DayOfWeekArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionFormat) []DayOfWeek { return v.DaysOfTheWeek }).(DayOfWeekArrayOutput)
}

func (o WeeklyRetentionFormatOutput) WeeksOfTheMonth() WeekOfMonthArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionFormat) []WeekOfMonth { return v.WeeksOfTheMonth }).(WeekOfMonthArrayOutput)
}

type WeeklyRetentionFormatPtrOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionFormat)(nil)).Elem()
}

func (o WeeklyRetentionFormatPtrOutput) ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput {
	return o
}

func (o WeeklyRetentionFormatPtrOutput) ToWeeklyRetentionFormatPtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatPtrOutput {
	return o
}

func (o WeeklyRetentionFormatPtrOutput) Elem() WeeklyRetentionFormatOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormat) WeeklyRetentionFormat {
		if v != nil {
			return *v
		}
		var ret WeeklyRetentionFormat
		return ret
	}).(WeeklyRetentionFormatOutput)
}

func (o WeeklyRetentionFormatPtrOutput) DaysOfTheWeek() DayOfWeekArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormat) []DayOfWeek {
		if v == nil {
			return nil
		}
		return v.DaysOfTheWeek
	}).(DayOfWeekArrayOutput)
}

func (o WeeklyRetentionFormatPtrOutput) WeeksOfTheMonth() WeekOfMonthArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormat) []WeekOfMonth {
		if v == nil {
			return nil
		}
		return v.WeeksOfTheMonth
	}).(WeekOfMonthArrayOutput)
}

type WeeklyRetentionFormatResponse struct {
	DaysOfTheWeek   []string `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth []string `pulumi:"weeksOfTheMonth"`
}





type WeeklyRetentionFormatResponseInput interface {
	pulumi.Input

	ToWeeklyRetentionFormatResponseOutput() WeeklyRetentionFormatResponseOutput
	ToWeeklyRetentionFormatResponseOutputWithContext(context.Context) WeeklyRetentionFormatResponseOutput
}

type WeeklyRetentionFormatResponseArgs struct {
	DaysOfTheWeek   pulumi.StringArrayInput `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth pulumi.StringArrayInput `pulumi:"weeksOfTheMonth"`
}

func (WeeklyRetentionFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionFormatResponse)(nil)).Elem()
}

func (i WeeklyRetentionFormatResponseArgs) ToWeeklyRetentionFormatResponseOutput() WeeklyRetentionFormatResponseOutput {
	return i.ToWeeklyRetentionFormatResponseOutputWithContext(context.Background())
}

func (i WeeklyRetentionFormatResponseArgs) ToWeeklyRetentionFormatResponseOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatResponseOutput)
}

func (i WeeklyRetentionFormatResponseArgs) ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput {
	return i.ToWeeklyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (i WeeklyRetentionFormatResponseArgs) ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatResponseOutput).ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx)
}









type WeeklyRetentionFormatResponsePtrInput interface {
	pulumi.Input

	ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput
	ToWeeklyRetentionFormatResponsePtrOutputWithContext(context.Context) WeeklyRetentionFormatResponsePtrOutput
}

type weeklyRetentionFormatResponsePtrType WeeklyRetentionFormatResponseArgs

func WeeklyRetentionFormatResponsePtr(v *WeeklyRetentionFormatResponseArgs) WeeklyRetentionFormatResponsePtrInput {
	return (*weeklyRetentionFormatResponsePtrType)(v)
}

func (*weeklyRetentionFormatResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionFormatResponse)(nil)).Elem()
}

func (i *weeklyRetentionFormatResponsePtrType) ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput {
	return i.ToWeeklyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (i *weeklyRetentionFormatResponsePtrType) ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatResponsePtrOutput)
}

type WeeklyRetentionFormatResponseOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionFormatResponse)(nil)).Elem()
}

func (o WeeklyRetentionFormatResponseOutput) ToWeeklyRetentionFormatResponseOutput() WeeklyRetentionFormatResponseOutput {
	return o
}

func (o WeeklyRetentionFormatResponseOutput) ToWeeklyRetentionFormatResponseOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponseOutput {
	return o
}

func (o WeeklyRetentionFormatResponseOutput) ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput {
	return o.ToWeeklyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (o WeeklyRetentionFormatResponseOutput) ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyRetentionFormatResponse) *WeeklyRetentionFormatResponse {
		return &v
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o WeeklyRetentionFormatResponseOutput) DaysOfTheWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionFormatResponse) []string { return v.DaysOfTheWeek }).(pulumi.StringArrayOutput)
}

func (o WeeklyRetentionFormatResponseOutput) WeeksOfTheMonth() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionFormatResponse) []string { return v.WeeksOfTheMonth }).(pulumi.StringArrayOutput)
}

type WeeklyRetentionFormatResponsePtrOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionFormatResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionFormatResponse)(nil)).Elem()
}

func (o WeeklyRetentionFormatResponsePtrOutput) ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput {
	return o
}

func (o WeeklyRetentionFormatResponsePtrOutput) ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponsePtrOutput {
	return o
}

func (o WeeklyRetentionFormatResponsePtrOutput) Elem() WeeklyRetentionFormatResponseOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormatResponse) WeeklyRetentionFormatResponse {
		if v != nil {
			return *v
		}
		var ret WeeklyRetentionFormatResponse
		return ret
	}).(WeeklyRetentionFormatResponseOutput)
}

func (o WeeklyRetentionFormatResponsePtrOutput) DaysOfTheWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormatResponse) []string {
		if v == nil {
			return nil
		}
		return v.DaysOfTheWeek
	}).(pulumi.StringArrayOutput)
}

func (o WeeklyRetentionFormatResponsePtrOutput) WeeksOfTheMonth() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormatResponse) []string {
		if v == nil {
			return nil
		}
		return v.WeeksOfTheMonth
	}).(pulumi.StringArrayOutput)
}

type WeeklyRetentionSchedule struct {
	DaysOfTheWeek     []DayOfWeek        `pulumi:"daysOfTheWeek"`
	RetentionDuration *RetentionDuration `pulumi:"retentionDuration"`
	RetentionTimes    []string           `pulumi:"retentionTimes"`
}





type WeeklyRetentionScheduleInput interface {
	pulumi.Input

	ToWeeklyRetentionScheduleOutput() WeeklyRetentionScheduleOutput
	ToWeeklyRetentionScheduleOutputWithContext(context.Context) WeeklyRetentionScheduleOutput
}

type WeeklyRetentionScheduleArgs struct {
	DaysOfTheWeek     DayOfWeekArrayInput       `pulumi:"daysOfTheWeek"`
	RetentionDuration RetentionDurationPtrInput `pulumi:"retentionDuration"`
	RetentionTimes    pulumi.StringArrayInput   `pulumi:"retentionTimes"`
}

func (WeeklyRetentionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionSchedule)(nil)).Elem()
}

func (i WeeklyRetentionScheduleArgs) ToWeeklyRetentionScheduleOutput() WeeklyRetentionScheduleOutput {
	return i.ToWeeklyRetentionScheduleOutputWithContext(context.Background())
}

func (i WeeklyRetentionScheduleArgs) ToWeeklyRetentionScheduleOutputWithContext(ctx context.Context) WeeklyRetentionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleOutput)
}

func (i WeeklyRetentionScheduleArgs) ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput {
	return i.ToWeeklyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i WeeklyRetentionScheduleArgs) ToWeeklyRetentionSchedulePtrOutputWithContext(ctx context.Context) WeeklyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleOutput).ToWeeklyRetentionSchedulePtrOutputWithContext(ctx)
}









type WeeklyRetentionSchedulePtrInput interface {
	pulumi.Input

	ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput
	ToWeeklyRetentionSchedulePtrOutputWithContext(context.Context) WeeklyRetentionSchedulePtrOutput
}

type weeklyRetentionSchedulePtrType WeeklyRetentionScheduleArgs

func WeeklyRetentionSchedulePtr(v *WeeklyRetentionScheduleArgs) WeeklyRetentionSchedulePtrInput {
	return (*weeklyRetentionSchedulePtrType)(v)
}

func (*weeklyRetentionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionSchedule)(nil)).Elem()
}

func (i *weeklyRetentionSchedulePtrType) ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput {
	return i.ToWeeklyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i *weeklyRetentionSchedulePtrType) ToWeeklyRetentionSchedulePtrOutputWithContext(ctx context.Context) WeeklyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionSchedulePtrOutput)
}

type WeeklyRetentionScheduleOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionSchedule)(nil)).Elem()
}

func (o WeeklyRetentionScheduleOutput) ToWeeklyRetentionScheduleOutput() WeeklyRetentionScheduleOutput {
	return o
}

func (o WeeklyRetentionScheduleOutput) ToWeeklyRetentionScheduleOutputWithContext(ctx context.Context) WeeklyRetentionScheduleOutput {
	return o
}

func (o WeeklyRetentionScheduleOutput) ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput {
	return o.ToWeeklyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (o WeeklyRetentionScheduleOutput) ToWeeklyRetentionSchedulePtrOutputWithContext(ctx context.Context) WeeklyRetentionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyRetentionSchedule) *WeeklyRetentionSchedule {
		return &v
	}).(WeeklyRetentionSchedulePtrOutput)
}

func (o WeeklyRetentionScheduleOutput) DaysOfTheWeek() DayOfWeekArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionSchedule) []DayOfWeek { return v.DaysOfTheWeek }).(DayOfWeekArrayOutput)
}

func (o WeeklyRetentionScheduleOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v WeeklyRetentionSchedule) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o WeeklyRetentionScheduleOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionSchedule) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type WeeklyRetentionSchedulePtrOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionSchedule)(nil)).Elem()
}

func (o WeeklyRetentionSchedulePtrOutput) ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput {
	return o
}

func (o WeeklyRetentionSchedulePtrOutput) ToWeeklyRetentionSchedulePtrOutputWithContext(ctx context.Context) WeeklyRetentionSchedulePtrOutput {
	return o
}

func (o WeeklyRetentionSchedulePtrOutput) Elem() WeeklyRetentionScheduleOutput {
	return o.ApplyT(func(v *WeeklyRetentionSchedule) WeeklyRetentionSchedule {
		if v != nil {
			return *v
		}
		var ret WeeklyRetentionSchedule
		return ret
	}).(WeeklyRetentionScheduleOutput)
}

func (o WeeklyRetentionSchedulePtrOutput) DaysOfTheWeek() DayOfWeekArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionSchedule) []DayOfWeek {
		if v == nil {
			return nil
		}
		return v.DaysOfTheWeek
	}).(DayOfWeekArrayOutput)
}

func (o WeeklyRetentionSchedulePtrOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v *WeeklyRetentionSchedule) *RetentionDuration {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationPtrOutput)
}

func (o WeeklyRetentionSchedulePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionSchedule) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type WeeklyRetentionScheduleResponse struct {
	DaysOfTheWeek     []string                   `pulumi:"daysOfTheWeek"`
	RetentionDuration *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionTimes    []string                   `pulumi:"retentionTimes"`
}





type WeeklyRetentionScheduleResponseInput interface {
	pulumi.Input

	ToWeeklyRetentionScheduleResponseOutput() WeeklyRetentionScheduleResponseOutput
	ToWeeklyRetentionScheduleResponseOutputWithContext(context.Context) WeeklyRetentionScheduleResponseOutput
}

type WeeklyRetentionScheduleResponseArgs struct {
	DaysOfTheWeek     pulumi.StringArrayInput           `pulumi:"daysOfTheWeek"`
	RetentionDuration RetentionDurationResponsePtrInput `pulumi:"retentionDuration"`
	RetentionTimes    pulumi.StringArrayInput           `pulumi:"retentionTimes"`
}

func (WeeklyRetentionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionScheduleResponse)(nil)).Elem()
}

func (i WeeklyRetentionScheduleResponseArgs) ToWeeklyRetentionScheduleResponseOutput() WeeklyRetentionScheduleResponseOutput {
	return i.ToWeeklyRetentionScheduleResponseOutputWithContext(context.Background())
}

func (i WeeklyRetentionScheduleResponseArgs) ToWeeklyRetentionScheduleResponseOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleResponseOutput)
}

func (i WeeklyRetentionScheduleResponseArgs) ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput {
	return i.ToWeeklyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i WeeklyRetentionScheduleResponseArgs) ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleResponseOutput).ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx)
}









type WeeklyRetentionScheduleResponsePtrInput interface {
	pulumi.Input

	ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput
	ToWeeklyRetentionScheduleResponsePtrOutputWithContext(context.Context) WeeklyRetentionScheduleResponsePtrOutput
}

type weeklyRetentionScheduleResponsePtrType WeeklyRetentionScheduleResponseArgs

func WeeklyRetentionScheduleResponsePtr(v *WeeklyRetentionScheduleResponseArgs) WeeklyRetentionScheduleResponsePtrInput {
	return (*weeklyRetentionScheduleResponsePtrType)(v)
}

func (*weeklyRetentionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionScheduleResponse)(nil)).Elem()
}

func (i *weeklyRetentionScheduleResponsePtrType) ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput {
	return i.ToWeeklyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *weeklyRetentionScheduleResponsePtrType) ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleResponsePtrOutput)
}

type WeeklyRetentionScheduleResponseOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionScheduleResponse)(nil)).Elem()
}

func (o WeeklyRetentionScheduleResponseOutput) ToWeeklyRetentionScheduleResponseOutput() WeeklyRetentionScheduleResponseOutput {
	return o
}

func (o WeeklyRetentionScheduleResponseOutput) ToWeeklyRetentionScheduleResponseOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponseOutput {
	return o
}

func (o WeeklyRetentionScheduleResponseOutput) ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput {
	return o.ToWeeklyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o WeeklyRetentionScheduleResponseOutput) ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyRetentionScheduleResponse) *WeeklyRetentionScheduleResponse {
		return &v
	}).(WeeklyRetentionScheduleResponsePtrOutput)
}

func (o WeeklyRetentionScheduleResponseOutput) DaysOfTheWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionScheduleResponse) []string { return v.DaysOfTheWeek }).(pulumi.StringArrayOutput)
}

func (o WeeklyRetentionScheduleResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v WeeklyRetentionScheduleResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o WeeklyRetentionScheduleResponseOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionScheduleResponse) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type WeeklyRetentionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionScheduleResponse)(nil)).Elem()
}

func (o WeeklyRetentionScheduleResponsePtrOutput) ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput {
	return o
}

func (o WeeklyRetentionScheduleResponsePtrOutput) ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponsePtrOutput {
	return o
}

func (o WeeklyRetentionScheduleResponsePtrOutput) Elem() WeeklyRetentionScheduleResponseOutput {
	return o.ApplyT(func(v *WeeklyRetentionScheduleResponse) WeeklyRetentionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret WeeklyRetentionScheduleResponse
		return ret
	}).(WeeklyRetentionScheduleResponseOutput)
}

func (o WeeklyRetentionScheduleResponsePtrOutput) DaysOfTheWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.DaysOfTheWeek
	}).(pulumi.StringArrayOutput)
}

func (o WeeklyRetentionScheduleResponsePtrOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v *WeeklyRetentionScheduleResponse) *RetentionDurationResponse {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationResponsePtrOutput)
}

func (o WeeklyRetentionScheduleResponsePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type WorkloadInquiryDetails struct {
	InquiryValidation *InquiryValidation `pulumi:"inquiryValidation"`
	ItemCount         *float64           `pulumi:"itemCount"`
	Type              *string            `pulumi:"type"`
}





type WorkloadInquiryDetailsInput interface {
	pulumi.Input

	ToWorkloadInquiryDetailsOutput() WorkloadInquiryDetailsOutput
	ToWorkloadInquiryDetailsOutputWithContext(context.Context) WorkloadInquiryDetailsOutput
}

type WorkloadInquiryDetailsArgs struct {
	InquiryValidation InquiryValidationPtrInput `pulumi:"inquiryValidation"`
	ItemCount         pulumi.Float64PtrInput    `pulumi:"itemCount"`
	Type              pulumi.StringPtrInput     `pulumi:"type"`
}

func (WorkloadInquiryDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadInquiryDetails)(nil)).Elem()
}

func (i WorkloadInquiryDetailsArgs) ToWorkloadInquiryDetailsOutput() WorkloadInquiryDetailsOutput {
	return i.ToWorkloadInquiryDetailsOutputWithContext(context.Background())
}

func (i WorkloadInquiryDetailsArgs) ToWorkloadInquiryDetailsOutputWithContext(ctx context.Context) WorkloadInquiryDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadInquiryDetailsOutput)
}





type WorkloadInquiryDetailsArrayInput interface {
	pulumi.Input

	ToWorkloadInquiryDetailsArrayOutput() WorkloadInquiryDetailsArrayOutput
	ToWorkloadInquiryDetailsArrayOutputWithContext(context.Context) WorkloadInquiryDetailsArrayOutput
}

type WorkloadInquiryDetailsArray []WorkloadInquiryDetailsInput

func (WorkloadInquiryDetailsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadInquiryDetails)(nil)).Elem()
}

func (i WorkloadInquiryDetailsArray) ToWorkloadInquiryDetailsArrayOutput() WorkloadInquiryDetailsArrayOutput {
	return i.ToWorkloadInquiryDetailsArrayOutputWithContext(context.Background())
}

func (i WorkloadInquiryDetailsArray) ToWorkloadInquiryDetailsArrayOutputWithContext(ctx context.Context) WorkloadInquiryDetailsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadInquiryDetailsArrayOutput)
}

type WorkloadInquiryDetailsOutput struct{ *pulumi.OutputState }

func (WorkloadInquiryDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadInquiryDetails)(nil)).Elem()
}

func (o WorkloadInquiryDetailsOutput) ToWorkloadInquiryDetailsOutput() WorkloadInquiryDetailsOutput {
	return o
}

func (o WorkloadInquiryDetailsOutput) ToWorkloadInquiryDetailsOutputWithContext(ctx context.Context) WorkloadInquiryDetailsOutput {
	return o
}

func (o WorkloadInquiryDetailsOutput) InquiryValidation() InquiryValidationPtrOutput {
	return o.ApplyT(func(v WorkloadInquiryDetails) *InquiryValidation { return v.InquiryValidation }).(InquiryValidationPtrOutput)
}

func (o WorkloadInquiryDetailsOutput) ItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadInquiryDetails) *float64 { return v.ItemCount }).(pulumi.Float64PtrOutput)
}

func (o WorkloadInquiryDetailsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadInquiryDetails) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkloadInquiryDetailsArrayOutput struct{ *pulumi.OutputState }

func (WorkloadInquiryDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadInquiryDetails)(nil)).Elem()
}

func (o WorkloadInquiryDetailsArrayOutput) ToWorkloadInquiryDetailsArrayOutput() WorkloadInquiryDetailsArrayOutput {
	return o
}

func (o WorkloadInquiryDetailsArrayOutput) ToWorkloadInquiryDetailsArrayOutputWithContext(ctx context.Context) WorkloadInquiryDetailsArrayOutput {
	return o
}

func (o WorkloadInquiryDetailsArrayOutput) Index(i pulumi.IntInput) WorkloadInquiryDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkloadInquiryDetails {
		return vs[0].([]WorkloadInquiryDetails)[vs[1].(int)]
	}).(WorkloadInquiryDetailsOutput)
}

type WorkloadInquiryDetailsResponse struct {
	InquiryValidation *InquiryValidationResponse `pulumi:"inquiryValidation"`
	ItemCount         *float64                   `pulumi:"itemCount"`
	Type              *string                    `pulumi:"type"`
}





type WorkloadInquiryDetailsResponseInput interface {
	pulumi.Input

	ToWorkloadInquiryDetailsResponseOutput() WorkloadInquiryDetailsResponseOutput
	ToWorkloadInquiryDetailsResponseOutputWithContext(context.Context) WorkloadInquiryDetailsResponseOutput
}

type WorkloadInquiryDetailsResponseArgs struct {
	InquiryValidation InquiryValidationResponsePtrInput `pulumi:"inquiryValidation"`
	ItemCount         pulumi.Float64PtrInput            `pulumi:"itemCount"`
	Type              pulumi.StringPtrInput             `pulumi:"type"`
}

func (WorkloadInquiryDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadInquiryDetailsResponse)(nil)).Elem()
}

func (i WorkloadInquiryDetailsResponseArgs) ToWorkloadInquiryDetailsResponseOutput() WorkloadInquiryDetailsResponseOutput {
	return i.ToWorkloadInquiryDetailsResponseOutputWithContext(context.Background())
}

func (i WorkloadInquiryDetailsResponseArgs) ToWorkloadInquiryDetailsResponseOutputWithContext(ctx context.Context) WorkloadInquiryDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadInquiryDetailsResponseOutput)
}





type WorkloadInquiryDetailsResponseArrayInput interface {
	pulumi.Input

	ToWorkloadInquiryDetailsResponseArrayOutput() WorkloadInquiryDetailsResponseArrayOutput
	ToWorkloadInquiryDetailsResponseArrayOutputWithContext(context.Context) WorkloadInquiryDetailsResponseArrayOutput
}

type WorkloadInquiryDetailsResponseArray []WorkloadInquiryDetailsResponseInput

func (WorkloadInquiryDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadInquiryDetailsResponse)(nil)).Elem()
}

func (i WorkloadInquiryDetailsResponseArray) ToWorkloadInquiryDetailsResponseArrayOutput() WorkloadInquiryDetailsResponseArrayOutput {
	return i.ToWorkloadInquiryDetailsResponseArrayOutputWithContext(context.Background())
}

func (i WorkloadInquiryDetailsResponseArray) ToWorkloadInquiryDetailsResponseArrayOutputWithContext(ctx context.Context) WorkloadInquiryDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadInquiryDetailsResponseArrayOutput)
}

type WorkloadInquiryDetailsResponseOutput struct{ *pulumi.OutputState }

func (WorkloadInquiryDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadInquiryDetailsResponse)(nil)).Elem()
}

func (o WorkloadInquiryDetailsResponseOutput) ToWorkloadInquiryDetailsResponseOutput() WorkloadInquiryDetailsResponseOutput {
	return o
}

func (o WorkloadInquiryDetailsResponseOutput) ToWorkloadInquiryDetailsResponseOutputWithContext(ctx context.Context) WorkloadInquiryDetailsResponseOutput {
	return o
}

func (o WorkloadInquiryDetailsResponseOutput) InquiryValidation() InquiryValidationResponsePtrOutput {
	return o.ApplyT(func(v WorkloadInquiryDetailsResponse) *InquiryValidationResponse { return v.InquiryValidation }).(InquiryValidationResponsePtrOutput)
}

func (o WorkloadInquiryDetailsResponseOutput) ItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadInquiryDetailsResponse) *float64 { return v.ItemCount }).(pulumi.Float64PtrOutput)
}

func (o WorkloadInquiryDetailsResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadInquiryDetailsResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkloadInquiryDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkloadInquiryDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadInquiryDetailsResponse)(nil)).Elem()
}

func (o WorkloadInquiryDetailsResponseArrayOutput) ToWorkloadInquiryDetailsResponseArrayOutput() WorkloadInquiryDetailsResponseArrayOutput {
	return o
}

func (o WorkloadInquiryDetailsResponseArrayOutput) ToWorkloadInquiryDetailsResponseArrayOutputWithContext(ctx context.Context) WorkloadInquiryDetailsResponseArrayOutput {
	return o
}

func (o WorkloadInquiryDetailsResponseArrayOutput) Index(i pulumi.IntInput) WorkloadInquiryDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkloadInquiryDetailsResponse {
		return vs[0].([]WorkloadInquiryDetailsResponse)[vs[1].(int)]
	}).(WorkloadInquiryDetailsResponseOutput)
}

type YearlyRetentionSchedule struct {
	MonthsOfYear                []MonthOfYear          `pulumi:"monthsOfYear"`
	RetentionDuration           *RetentionDuration     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormat  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormat `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string               `pulumi:"retentionTimes"`
}





type YearlyRetentionScheduleInput interface {
	pulumi.Input

	ToYearlyRetentionScheduleOutput() YearlyRetentionScheduleOutput
	ToYearlyRetentionScheduleOutputWithContext(context.Context) YearlyRetentionScheduleOutput
}

type YearlyRetentionScheduleArgs struct {
	MonthsOfYear                MonthOfYearArrayInput         `pulumi:"monthsOfYear"`
	RetentionDuration           RetentionDurationPtrInput     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      DailyRetentionFormatPtrInput  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType pulumi.StringPtrInput         `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     WeeklyRetentionFormatPtrInput `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              pulumi.StringArrayInput       `pulumi:"retentionTimes"`
}

func (YearlyRetentionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*YearlyRetentionSchedule)(nil)).Elem()
}

func (i YearlyRetentionScheduleArgs) ToYearlyRetentionScheduleOutput() YearlyRetentionScheduleOutput {
	return i.ToYearlyRetentionScheduleOutputWithContext(context.Background())
}

func (i YearlyRetentionScheduleArgs) ToYearlyRetentionScheduleOutputWithContext(ctx context.Context) YearlyRetentionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleOutput)
}

func (i YearlyRetentionScheduleArgs) ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput {
	return i.ToYearlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i YearlyRetentionScheduleArgs) ToYearlyRetentionSchedulePtrOutputWithContext(ctx context.Context) YearlyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleOutput).ToYearlyRetentionSchedulePtrOutputWithContext(ctx)
}









type YearlyRetentionSchedulePtrInput interface {
	pulumi.Input

	ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput
	ToYearlyRetentionSchedulePtrOutputWithContext(context.Context) YearlyRetentionSchedulePtrOutput
}

type yearlyRetentionSchedulePtrType YearlyRetentionScheduleArgs

func YearlyRetentionSchedulePtr(v *YearlyRetentionScheduleArgs) YearlyRetentionSchedulePtrInput {
	return (*yearlyRetentionSchedulePtrType)(v)
}

func (*yearlyRetentionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**YearlyRetentionSchedule)(nil)).Elem()
}

func (i *yearlyRetentionSchedulePtrType) ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput {
	return i.ToYearlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i *yearlyRetentionSchedulePtrType) ToYearlyRetentionSchedulePtrOutputWithContext(ctx context.Context) YearlyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionSchedulePtrOutput)
}

type YearlyRetentionScheduleOutput struct{ *pulumi.OutputState }

func (YearlyRetentionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*YearlyRetentionSchedule)(nil)).Elem()
}

func (o YearlyRetentionScheduleOutput) ToYearlyRetentionScheduleOutput() YearlyRetentionScheduleOutput {
	return o
}

func (o YearlyRetentionScheduleOutput) ToYearlyRetentionScheduleOutputWithContext(ctx context.Context) YearlyRetentionScheduleOutput {
	return o
}

func (o YearlyRetentionScheduleOutput) ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput {
	return o.ToYearlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (o YearlyRetentionScheduleOutput) ToYearlyRetentionSchedulePtrOutputWithContext(ctx context.Context) YearlyRetentionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v YearlyRetentionSchedule) *YearlyRetentionSchedule {
		return &v
	}).(YearlyRetentionSchedulePtrOutput)
}

func (o YearlyRetentionScheduleOutput) MonthsOfYear() MonthOfYearArrayOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) []MonthOfYear { return v.MonthsOfYear }).(MonthOfYearArrayOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionScheduleDaily() DailyRetentionFormatPtrOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) *DailyRetentionFormat { return v.RetentionScheduleDaily }).(DailyRetentionFormatPtrOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) *string { return v.RetentionScheduleFormatType }).(pulumi.StringPtrOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionScheduleWeekly() WeeklyRetentionFormatPtrOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) *WeeklyRetentionFormat { return v.RetentionScheduleWeekly }).(WeeklyRetentionFormatPtrOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type YearlyRetentionSchedulePtrOutput struct{ *pulumi.OutputState }

func (YearlyRetentionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**YearlyRetentionSchedule)(nil)).Elem()
}

func (o YearlyRetentionSchedulePtrOutput) ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput {
	return o
}

func (o YearlyRetentionSchedulePtrOutput) ToYearlyRetentionSchedulePtrOutputWithContext(ctx context.Context) YearlyRetentionSchedulePtrOutput {
	return o
}

func (o YearlyRetentionSchedulePtrOutput) Elem() YearlyRetentionScheduleOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) YearlyRetentionSchedule {
		if v != nil {
			return *v
		}
		var ret YearlyRetentionSchedule
		return ret
	}).(YearlyRetentionScheduleOutput)
}

func (o YearlyRetentionSchedulePtrOutput) MonthsOfYear() MonthOfYearArrayOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) []MonthOfYear {
		if v == nil {
			return nil
		}
		return v.MonthsOfYear
	}).(MonthOfYearArrayOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) *RetentionDuration {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationPtrOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionScheduleDaily() DailyRetentionFormatPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) *DailyRetentionFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatPtrOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) *string {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleFormatType
	}).(pulumi.StringPtrOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionScheduleWeekly() WeeklyRetentionFormatPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) *WeeklyRetentionFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatPtrOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type YearlyRetentionScheduleResponse struct {
	MonthsOfYear                []string                       `pulumi:"monthsOfYear"`
	RetentionDuration           *RetentionDurationResponse     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormatResponse  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                        `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormatResponse `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string                       `pulumi:"retentionTimes"`
}





type YearlyRetentionScheduleResponseInput interface {
	pulumi.Input

	ToYearlyRetentionScheduleResponseOutput() YearlyRetentionScheduleResponseOutput
	ToYearlyRetentionScheduleResponseOutputWithContext(context.Context) YearlyRetentionScheduleResponseOutput
}

type YearlyRetentionScheduleResponseArgs struct {
	MonthsOfYear                pulumi.StringArrayInput               `pulumi:"monthsOfYear"`
	RetentionDuration           RetentionDurationResponsePtrInput     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      DailyRetentionFormatResponsePtrInput  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType pulumi.StringPtrInput                 `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     WeeklyRetentionFormatResponsePtrInput `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              pulumi.StringArrayInput               `pulumi:"retentionTimes"`
}

func (YearlyRetentionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*YearlyRetentionScheduleResponse)(nil)).Elem()
}

func (i YearlyRetentionScheduleResponseArgs) ToYearlyRetentionScheduleResponseOutput() YearlyRetentionScheduleResponseOutput {
	return i.ToYearlyRetentionScheduleResponseOutputWithContext(context.Background())
}

func (i YearlyRetentionScheduleResponseArgs) ToYearlyRetentionScheduleResponseOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleResponseOutput)
}

func (i YearlyRetentionScheduleResponseArgs) ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput {
	return i.ToYearlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i YearlyRetentionScheduleResponseArgs) ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleResponseOutput).ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx)
}









type YearlyRetentionScheduleResponsePtrInput interface {
	pulumi.Input

	ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput
	ToYearlyRetentionScheduleResponsePtrOutputWithContext(context.Context) YearlyRetentionScheduleResponsePtrOutput
}

type yearlyRetentionScheduleResponsePtrType YearlyRetentionScheduleResponseArgs

func YearlyRetentionScheduleResponsePtr(v *YearlyRetentionScheduleResponseArgs) YearlyRetentionScheduleResponsePtrInput {
	return (*yearlyRetentionScheduleResponsePtrType)(v)
}

func (*yearlyRetentionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**YearlyRetentionScheduleResponse)(nil)).Elem()
}

func (i *yearlyRetentionScheduleResponsePtrType) ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput {
	return i.ToYearlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *yearlyRetentionScheduleResponsePtrType) ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleResponsePtrOutput)
}

type YearlyRetentionScheduleResponseOutput struct{ *pulumi.OutputState }

func (YearlyRetentionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*YearlyRetentionScheduleResponse)(nil)).Elem()
}

func (o YearlyRetentionScheduleResponseOutput) ToYearlyRetentionScheduleResponseOutput() YearlyRetentionScheduleResponseOutput {
	return o
}

func (o YearlyRetentionScheduleResponseOutput) ToYearlyRetentionScheduleResponseOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponseOutput {
	return o
}

func (o YearlyRetentionScheduleResponseOutput) ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput {
	return o.ToYearlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o YearlyRetentionScheduleResponseOutput) ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v YearlyRetentionScheduleResponse) *YearlyRetentionScheduleResponse {
		return &v
	}).(YearlyRetentionScheduleResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) MonthsOfYear() pulumi.StringArrayOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) []string { return v.MonthsOfYear }).(pulumi.StringArrayOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionScheduleDaily() DailyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) *DailyRetentionFormatResponse { return v.RetentionScheduleDaily }).(DailyRetentionFormatResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) *string { return v.RetentionScheduleFormatType }).(pulumi.StringPtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionScheduleWeekly() WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) *WeeklyRetentionFormatResponse {
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type YearlyRetentionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (YearlyRetentionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**YearlyRetentionScheduleResponse)(nil)).Elem()
}

func (o YearlyRetentionScheduleResponsePtrOutput) ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput {
	return o
}

func (o YearlyRetentionScheduleResponsePtrOutput) ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponsePtrOutput {
	return o
}

func (o YearlyRetentionScheduleResponsePtrOutput) Elem() YearlyRetentionScheduleResponseOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) YearlyRetentionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret YearlyRetentionScheduleResponse
		return ret
	}).(YearlyRetentionScheduleResponseOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) MonthsOfYear() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.MonthsOfYear
	}).(pulumi.StringArrayOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) *RetentionDurationResponse {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionScheduleDaily() DailyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) *DailyRetentionFormatResponse {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleFormatType
	}).(pulumi.StringPtrOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionScheduleWeekly() WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) *WeeklyRetentionFormatResponse {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureBackupServerContainerOutput{})
	pulumi.RegisterOutputType(AzureBackupServerContainerResponseOutput{})
	pulumi.RegisterOutputType(AzureFileShareProtectionPolicyOutput{})
	pulumi.RegisterOutputType(AzureFileShareProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(AzureFileshareProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureFileshareProtectedItemExtendedInfoOutput{})
	pulumi.RegisterOutputType(AzureFileshareProtectedItemExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(AzureFileshareProtectedItemExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureFileshareProtectedItemExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureFileshareProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureIaaSClassicComputeVMContainerOutput{})
	pulumi.RegisterOutputType(AzureIaaSClassicComputeVMContainerResponseOutput{})
	pulumi.RegisterOutputType(AzureIaaSClassicComputeVMProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureIaaSClassicComputeVMProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureIaaSComputeVMContainerOutput{})
	pulumi.RegisterOutputType(AzureIaaSComputeVMContainerResponseOutput{})
	pulumi.RegisterOutputType(AzureIaaSComputeVMProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureIaaSComputeVMProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMHealthDetailsResponseOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMHealthDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectedItemExtendedInfoOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectedItemExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectedItemExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectedItemExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectionPolicyOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(AzureRecoveryServiceVaultProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureRecoveryServiceVaultProtectionIntentResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureResourceProtectionIntentResponseOutput{})
	pulumi.RegisterOutputType(AzureSQLAGWorkloadContainerProtectionContainerOutput{})
	pulumi.RegisterOutputType(AzureSQLAGWorkloadContainerProtectionContainerResponseOutput{})
	pulumi.RegisterOutputType(AzureSqlContainerOutput{})
	pulumi.RegisterOutputType(AzureSqlContainerResponseOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectedItemExtendedInfoOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectedItemExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectedItemExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectedItemExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectionPolicyOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(AzureStorageContainerOutput{})
	pulumi.RegisterOutputType(AzureStorageContainerResponseOutput{})
	pulumi.RegisterOutputType(AzureVMAppContainerProtectionContainerOutput{})
	pulumi.RegisterOutputType(AzureVMAppContainerProtectionContainerResponseOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadProtectedItemExtendedInfoOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadProtectedItemExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadProtectedItemExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadProtectedItemExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadProtectionPolicyOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadSAPAseDatabaseProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadSAPAseDatabaseProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadSAPHanaDatabaseProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadSAPHanaDatabaseProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadSQLDatabaseProtectedItemOutput{})
	pulumi.RegisterOutputType(AzureVmWorkloadSQLDatabaseProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(AzureWorkloadAutoProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureWorkloadAutoProtectionIntentResponseOutput{})
	pulumi.RegisterOutputType(AzureWorkloadContainerOutput{})
	pulumi.RegisterOutputType(AzureWorkloadContainerAutoProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureWorkloadContainerAutoProtectionIntentResponseOutput{})
	pulumi.RegisterOutputType(AzureWorkloadContainerExtendedInfoOutput{})
	pulumi.RegisterOutputType(AzureWorkloadContainerExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(AzureWorkloadContainerExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureWorkloadContainerExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureWorkloadContainerResponseOutput{})
	pulumi.RegisterOutputType(AzureWorkloadSQLAutoProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureWorkloadSQLAutoProtectionIntentResponseOutput{})
	pulumi.RegisterOutputType(ContainerIdentityInfoOutput{})
	pulumi.RegisterOutputType(ContainerIdentityInfoPtrOutput{})
	pulumi.RegisterOutputType(ContainerIdentityInfoResponseOutput{})
	pulumi.RegisterOutputType(ContainerIdentityInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(DPMContainerExtendedInfoOutput{})
	pulumi.RegisterOutputType(DPMContainerExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(DPMContainerExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(DPMContainerExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(DPMProtectedItemOutput{})
	pulumi.RegisterOutputType(DPMProtectedItemExtendedInfoOutput{})
	pulumi.RegisterOutputType(DPMProtectedItemExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(DPMProtectedItemExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(DPMProtectedItemExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(DPMProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(DailyRetentionFormatOutput{})
	pulumi.RegisterOutputType(DailyRetentionFormatPtrOutput{})
	pulumi.RegisterOutputType(DailyRetentionFormatResponseOutput{})
	pulumi.RegisterOutputType(DailyRetentionFormatResponsePtrOutput{})
	pulumi.RegisterOutputType(DailyRetentionScheduleOutput{})
	pulumi.RegisterOutputType(DailyRetentionSchedulePtrOutput{})
	pulumi.RegisterOutputType(DailyRetentionScheduleResponseOutput{})
	pulumi.RegisterOutputType(DailyRetentionScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(DayOutput{})
	pulumi.RegisterOutputType(DayArrayOutput{})
	pulumi.RegisterOutputType(DayResponseOutput{})
	pulumi.RegisterOutputType(DayResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskExclusionPropertiesOutput{})
	pulumi.RegisterOutputType(DiskExclusionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DiskExclusionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DiskExclusionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DistributedNodesInfoOutput{})
	pulumi.RegisterOutputType(DistributedNodesInfoArrayOutput{})
	pulumi.RegisterOutputType(DistributedNodesInfoResponseOutput{})
	pulumi.RegisterOutputType(DistributedNodesInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(DpmContainerOutput{})
	pulumi.RegisterOutputType(DpmContainerResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponsePtrOutput{})
	pulumi.RegisterOutputType(ExtendedPropertiesOutput{})
	pulumi.RegisterOutputType(ExtendedPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ExtendedPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ExtendedPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GenericContainerOutput{})
	pulumi.RegisterOutputType(GenericContainerExtendedInfoOutput{})
	pulumi.RegisterOutputType(GenericContainerExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(GenericContainerExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(GenericContainerExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(GenericContainerResponseOutput{})
	pulumi.RegisterOutputType(GenericProtectedItemOutput{})
	pulumi.RegisterOutputType(GenericProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(GenericProtectionPolicyOutput{})
	pulumi.RegisterOutputType(GenericProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(HourlyScheduleOutput{})
	pulumi.RegisterOutputType(HourlySchedulePtrOutput{})
	pulumi.RegisterOutputType(HourlyScheduleResponseOutput{})
	pulumi.RegisterOutputType(HourlyScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(IaaSVMContainerOutput{})
	pulumi.RegisterOutputType(IaaSVMContainerResponseOutput{})
	pulumi.RegisterOutputType(InquiryInfoOutput{})
	pulumi.RegisterOutputType(InquiryInfoPtrOutput{})
	pulumi.RegisterOutputType(InquiryInfoResponseOutput{})
	pulumi.RegisterOutputType(InquiryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(InquiryValidationOutput{})
	pulumi.RegisterOutputType(InquiryValidationPtrOutput{})
	pulumi.RegisterOutputType(InquiryValidationResponseOutput{})
	pulumi.RegisterOutputType(InquiryValidationResponsePtrOutput{})
	pulumi.RegisterOutputType(InstantRPAdditionalDetailsOutput{})
	pulumi.RegisterOutputType(InstantRPAdditionalDetailsPtrOutput{})
	pulumi.RegisterOutputType(InstantRPAdditionalDetailsResponseOutput{})
	pulumi.RegisterOutputType(InstantRPAdditionalDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(KPIResourceHealthDetailsOutput{})
	pulumi.RegisterOutputType(KPIResourceHealthDetailsMapOutput{})
	pulumi.RegisterOutputType(KPIResourceHealthDetailsResponseOutput{})
	pulumi.RegisterOutputType(KPIResourceHealthDetailsResponseMapOutput{})
	pulumi.RegisterOutputType(LogSchedulePolicyOutput{})
	pulumi.RegisterOutputType(LogSchedulePolicyResponseOutput{})
	pulumi.RegisterOutputType(LongTermRetentionPolicyOutput{})
	pulumi.RegisterOutputType(LongTermRetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(LongTermSchedulePolicyOutput{})
	pulumi.RegisterOutputType(LongTermSchedulePolicyResponseOutput{})
	pulumi.RegisterOutputType(MABContainerHealthDetailsOutput{})
	pulumi.RegisterOutputType(MABContainerHealthDetailsArrayOutput{})
	pulumi.RegisterOutputType(MABContainerHealthDetailsResponseOutput{})
	pulumi.RegisterOutputType(MABContainerHealthDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(MabContainerOutput{})
	pulumi.RegisterOutputType(MabContainerExtendedInfoOutput{})
	pulumi.RegisterOutputType(MabContainerExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(MabContainerExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(MabContainerExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(MabContainerResponseOutput{})
	pulumi.RegisterOutputType(MabFileFolderProtectedItemOutput{})
	pulumi.RegisterOutputType(MabFileFolderProtectedItemExtendedInfoOutput{})
	pulumi.RegisterOutputType(MabFileFolderProtectedItemExtendedInfoPtrOutput{})
	pulumi.RegisterOutputType(MabFileFolderProtectedItemExtendedInfoResponseOutput{})
	pulumi.RegisterOutputType(MabFileFolderProtectedItemExtendedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(MabFileFolderProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(MabProtectionPolicyOutput{})
	pulumi.RegisterOutputType(MabProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(MonthlyRetentionScheduleOutput{})
	pulumi.RegisterOutputType(MonthlyRetentionSchedulePtrOutput{})
	pulumi.RegisterOutputType(MonthlyRetentionScheduleResponseOutput{})
	pulumi.RegisterOutputType(MonthlyRetentionScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailResponseOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceGuardProxyBaseResponseOutput{})
	pulumi.RegisterOutputType(ResourceGuardProxyBaseResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceHealthDetailsResponseOutput{})
	pulumi.RegisterOutputType(ResourceHealthDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(RetentionDurationOutput{})
	pulumi.RegisterOutputType(RetentionDurationPtrOutput{})
	pulumi.RegisterOutputType(RetentionDurationResponseOutput{})
	pulumi.RegisterOutputType(RetentionDurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SettingsOutput{})
	pulumi.RegisterOutputType(SettingsPtrOutput{})
	pulumi.RegisterOutputType(SettingsResponseOutput{})
	pulumi.RegisterOutputType(SettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SimpleRetentionPolicyOutput{})
	pulumi.RegisterOutputType(SimpleRetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(SimpleSchedulePolicyOutput{})
	pulumi.RegisterOutputType(SimpleSchedulePolicyResponseOutput{})
	pulumi.RegisterOutputType(SubProtectionPolicyOutput{})
	pulumi.RegisterOutputType(SubProtectionPolicyArrayOutput{})
	pulumi.RegisterOutputType(SubProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(SubProtectionPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionFormatOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionFormatPtrOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionFormatResponseOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionFormatResponsePtrOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionScheduleOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionSchedulePtrOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionScheduleResponseOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkloadInquiryDetailsOutput{})
	pulumi.RegisterOutputType(WorkloadInquiryDetailsArrayOutput{})
	pulumi.RegisterOutputType(WorkloadInquiryDetailsResponseOutput{})
	pulumi.RegisterOutputType(WorkloadInquiryDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(YearlyRetentionScheduleOutput{})
	pulumi.RegisterOutputType(YearlyRetentionSchedulePtrOutput{})
	pulumi.RegisterOutputType(YearlyRetentionScheduleResponseOutput{})
	pulumi.RegisterOutputType(YearlyRetentionScheduleResponsePtrOutput{})
}
