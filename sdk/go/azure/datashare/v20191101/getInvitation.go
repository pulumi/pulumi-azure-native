


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInvitation(ctx *pulumi.Context, args *LookupInvitationArgs, opts ...pulumi.InvokeOption) (*LookupInvitationResult, error) {
	var rv LookupInvitationResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getInvitation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInvitationArgs struct {
	AccountName       string `pulumi:"accountName"`
	InvitationName    string `pulumi:"invitationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupInvitationResult struct {
	Id                      string  `pulumi:"id"`
	InvitationId            string  `pulumi:"invitationId"`
	InvitationStatus        string  `pulumi:"invitationStatus"`
	Name                    string  `pulumi:"name"`
	RespondedAt             string  `pulumi:"respondedAt"`
	SentAt                  string  `pulumi:"sentAt"`
	TargetActiveDirectoryId *string `pulumi:"targetActiveDirectoryId"`
	TargetEmail             *string `pulumi:"targetEmail"`
	TargetObjectId          *string `pulumi:"targetObjectId"`
	Type                    string  `pulumi:"type"`
	UserEmail               string  `pulumi:"userEmail"`
	UserName                string  `pulumi:"userName"`
}
