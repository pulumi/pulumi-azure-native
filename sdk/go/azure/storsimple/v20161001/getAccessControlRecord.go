


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessControlRecord(ctx *pulumi.Context, args *LookupAccessControlRecordArgs, opts ...pulumi.InvokeOption) (*LookupAccessControlRecordResult, error) {
	var rv LookupAccessControlRecordResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getAccessControlRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessControlRecordArgs struct {
	AccessControlRecordName string `pulumi:"accessControlRecordName"`
	ManagerName             string `pulumi:"managerName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupAccessControlRecordResult struct {
	Id            string `pulumi:"id"`
	InitiatorName string `pulumi:"initiatorName"`
	Name          string `pulumi:"name"`
	Type          string `pulumi:"type"`
}
