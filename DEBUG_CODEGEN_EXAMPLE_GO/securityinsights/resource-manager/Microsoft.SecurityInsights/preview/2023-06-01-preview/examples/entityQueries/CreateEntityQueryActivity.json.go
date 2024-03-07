package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewActivityCustomEntityQuery(ctx, "activityCustomEntityQuery", &securityinsights.ActivityCustomEntityQueryArgs{
			Content:     pulumi.String("On '{{Computer}}' the account '{{TargetAccount}}' was deleted by '{{AddedBy}}'"),
			Description: pulumi.String("Account deleted on host"),
			Enabled:     pulumi.Bool(true),
			EntitiesFilter: pulumi.StringArrayMap{
				"Host_OsFamily": pulumi.StringArray{
					pulumi.String("Windows"),
				},
			},
			EntityQueryId:   pulumi.String("07da3cc8-c8ad-4710-a44e-334cdcb7882b"),
			InputEntityType: pulumi.String("Host"),
			Kind:            pulumi.String("Activity"),
			QueryDefinitions: &securityinsights.ActivityEntityQueriesPropertiesQueryDefinitionsArgs{
				Query: pulumi.String(`let GetAccountActions = (v_Host_Name:string, v_Host_NTDomain:string, v_Host_DnsDomain:string, v_Host_AzureID:string, v_Host_OMSAgentID:string){
SecurityEvent
| where EventID in (4725, 4726, 4767, 4720, 4722, 4723, 4724)
// parsing for Host to handle variety of conventions coming from data
| extend Host_HostName = case(
Computer has '@', tostring(split(Computer, '@')[0]),
Computer has '\\', tostring(split(Computer, '\\')[1]),
Computer has '.', tostring(split(Computer, '.')[0]),
Computer
)
| extend Host_NTDomain = case(
Computer has '\\', tostring(split(Computer, '\\')[0]), 
Computer has '.', tostring(split(Computer, '.')[-2]), 
Computer
)
| extend Host_DnsDomain = case(
Computer has '\\', tostring(split(Computer, '\\')[0]), 
Computer has '.', strcat_array(array_slice(split(Computer,'.'),-2,-1),'.'), 
Computer
)
| where (Host_HostName =~ v_Host_Name and Host_NTDomain =~ v_Host_NTDomain) 
or (Host_HostName =~ v_Host_Name and Host_DnsDomain =~ v_Host_DnsDomain) 
or v_Host_AzureID =~ _ResourceId 
or v_Host_OMSAgentID == SourceComputerId
| project TimeGenerated, EventID, Activity, Computer, TargetAccount, TargetUserName, TargetDomainName, TargetSid, SubjectUserName, SubjectUserSid, _ResourceId, SourceComputerId
| extend AddedBy = SubjectUserName
// Future support for Activities
| extend timestamp = TimeGenerated, HostCustomEntity = Computer, AccountCustomEntity = TargetAccount
};
GetAccountActions('{{Host_HostName}}', '{{Host_NTDomain}}', '{{Host_DnsDomain}}', '{{Host_AzureID}}', '{{Host_OMSAgentID}}')
 
| where EventID == 4726 `),
			},
			RequiredInputFieldsSets: pulumi.StringArrayArray{
				pulumi.StringArray{
					pulumi.String("Host_HostName"),
					pulumi.String("Host_NTDomain"),
				},
				pulumi.StringArray{
					pulumi.String("Host_HostName"),
					pulumi.String("Host_DnsDomain"),
				},
				pulumi.StringArray{
					pulumi.String("Host_AzureID"),
				},
				pulumi.StringArray{
					pulumi.String("Host_OMSAgentID"),
				},
			},
			ResourceGroupName: pulumi.String("myRg"),
			Title:             pulumi.String("An account was deleted on this host"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
