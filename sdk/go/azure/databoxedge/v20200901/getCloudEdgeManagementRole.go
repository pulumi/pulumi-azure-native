


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudEdgeManagementRole(ctx *pulumi.Context, args *LookupCloudEdgeManagementRoleArgs, opts ...pulumi.InvokeOption) (*LookupCloudEdgeManagementRoleResult, error) {
	var rv LookupCloudEdgeManagementRoleResult
	err := ctx.Invoke("azure-native:databoxedge/v20200901:getCloudEdgeManagementRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudEdgeManagementRoleArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCloudEdgeManagementRoleResult struct {
	EdgeProfile           EdgeProfileResponse `pulumi:"edgeProfile"`
	Id                    string              `pulumi:"id"`
	Kind                  string              `pulumi:"kind"`
	LocalManagementStatus string              `pulumi:"localManagementStatus"`
	Name                  string              `pulumi:"name"`
	RoleStatus            string              `pulumi:"roleStatus"`
	SystemData            SystemDataResponse  `pulumi:"systemData"`
	Type                  string              `pulumi:"type"`
}
