


package batch

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:batch:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPoolArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPoolResult struct {
	AllocationState                 string                                `pulumi:"allocationState"`
	AllocationStateTransitionTime   string                                `pulumi:"allocationStateTransitionTime"`
	ApplicationLicenses             []string                              `pulumi:"applicationLicenses"`
	ApplicationPackages             []ApplicationPackageReferenceResponse `pulumi:"applicationPackages"`
	AutoScaleRun                    AutoScaleRunResponse                  `pulumi:"autoScaleRun"`
	Certificates                    []CertificateReferenceResponse        `pulumi:"certificates"`
	CreationTime                    string                                `pulumi:"creationTime"`
	CurrentDedicatedNodes           int                                   `pulumi:"currentDedicatedNodes"`
	CurrentLowPriorityNodes         int                                   `pulumi:"currentLowPriorityNodes"`
	DeploymentConfiguration         *DeploymentConfigurationResponse      `pulumi:"deploymentConfiguration"`
	DisplayName                     *string                               `pulumi:"displayName"`
	Etag                            string                                `pulumi:"etag"`
	Id                              string                                `pulumi:"id"`
	Identity                        *BatchPoolIdentityResponse            `pulumi:"identity"`
	InterNodeCommunication          *string                               `pulumi:"interNodeCommunication"`
	LastModified                    string                                `pulumi:"lastModified"`
	Metadata                        []MetadataItemResponse                `pulumi:"metadata"`
	MountConfiguration              []MountConfigurationResponse          `pulumi:"mountConfiguration"`
	Name                            string                                `pulumi:"name"`
	NetworkConfiguration            *NetworkConfigurationResponse         `pulumi:"networkConfiguration"`
	ProvisioningState               string                                `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime string                                `pulumi:"provisioningStateTransitionTime"`
	ResizeOperationStatus           ResizeOperationStatusResponse         `pulumi:"resizeOperationStatus"`
	ScaleSettings                   *ScaleSettingsResponse                `pulumi:"scaleSettings"`
	StartTask                       *StartTaskResponse                    `pulumi:"startTask"`
	TaskSchedulingPolicy            *TaskSchedulingPolicyResponse         `pulumi:"taskSchedulingPolicy"`
	TaskSlotsPerNode                *int                                  `pulumi:"taskSlotsPerNode"`
	Type                            string                                `pulumi:"type"`
	UserAccounts                    []UserAccountResponse                 `pulumi:"userAccounts"`
	VmSize                          *string                               `pulumi:"vmSize"`
}
