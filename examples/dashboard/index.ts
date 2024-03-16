// Copyright 2024, Pulumi Corporation.  All rights reserved.

import * as portal from "@pulumi/azure-native/portal";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("dashrg");

const dash = new portal.Dashboard("dash", {
    resourceGroupName: resourceGroup.name,
    tags: {
        "hidden-title": "Test Dashboard",
    },
    lenses: [{
        order: 0,
        parts: [
            {
                metadata: {
                    inputs: [],
                    settings: {
                        content: {
                            settings: {
                                content: `<div style='line-height:50px;'>
<span style='font-size:16px;font-weight:bold'>AZURE RESOURCE INVENTORY - </span>
<span>This section gives you an overview of all your Azure resources across all subscriptions that you can access.</span>
</div>`,
                                subtitle: "",
                                title: "",
                            },
                        },
                    },
                    type: "Extension/HubsExtension/PartType/MarkdownPart",
                },
                position: {
                    colSpan: 18,
                    rowSpan: 1,
                    x: 0,
                    y: 0,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of all my Azure resources",
                        },
                        {
                            name: "query",
                            value: "summarize Resources=count()",
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 3,
                    rowSpan: 4,
                    x: 0,
                    y: 1,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Top 10 resource counts by type",
                        },
                        {
                            name: "query",
                            value: `summarize ResourceCount=count() by type
| order by ResourceCount
| extend ['Resource count']=ResourceCount, ['Resource type']=type
| project ['Resource type'], ['Resource count']
| take 10`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 15,
                    rowSpan: 4,
                    x: 3,
                    y: 1,
                },
            },
            {
                metadata: {
                    inputs: [],
                    settings: {
                        content: {
                            settings: {
                                content: `<div style='line-height:50px'>
<span style='font-size:16px;font-weight:bold'>AZURE COMPUTE INVENTORY - </span>
<span>This section gives you an overview of your Azure Compute usage.</span>
</div>`,
                                subtitle: "",
                                title: "",
                            },
                        },
                    },
                    type: "Extension/HubsExtension/PartType/MarkdownPart",
                },
                position: {
                    colSpan: 18,
                    rowSpan: 1,
                    x: 0,
                    y: 5,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Virtual machines count (includes classic)",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.compute/virtualmachines" or type=="microsoft.classiccompute/virtualmachines"
| summarize VMCount=count()
| extend ['Count (Virtual Machines)']=VMCount
| project ['Count (Virtual Machines)']`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 4,
                    rowSpan: 3,
                    x: 0,
                    y: 6,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Virtual machines by operating system",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.compute/virtualmachines" or type == "microsoft.classiccompute/virtualmachines"
| extend OSType = iff(type == "microsoft.compute/virtualmachines", tostring(properties.storageProfile.osDisk.osType),tostring(properties.storageProfile.operatingSystemDisk.operatingSystem)) 
| summarize VMCount=count() by OSType
| order by VMCount desc
|extend ['Count (Virtual Machines)']=VMCount
| project OSType, ['Count (Virtual Machines)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 4,
                    rowSpan: 3,
                    x: 4,
                    y: 6,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Virtual machines by family",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.compute/virtualmachines" or type == "microsoft.classiccompute/virtualmachines"
| extend Size = iff(type == "microsoft.compute/virtualmachines", tostring(properties.hardwareProfile.vmSize),  tostring(properties.hardwareProfile.size))
| extend Family = extract('([^_]+)_', 1, Size, typeof(string))
| extend Family = iff(strlen(Family) == 0, Size, Family)
| summarize VMCount=count() by Family
| order by VMCount desc
|extend ['Count (Virtual Machines)']=VMCount
| project Family, ['Count (Virtual Machines)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 6,
                    rowSpan: 3,
                    x: 8,
                    y: 6,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of virtual machines by type",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.compute/virtualmachines" or type=="microsoft.classiccompute/virtualmachines"
| summarize VMCount=count() by type
| extend ['Count (Virtual Machines)']=VMCount
| extend Type = iff(type == "microsoft.compute/virtualmachines", "Resource Manager", "Classic")
| project Type,['Count (Virtual Machines)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 4,
                    rowSpan: 3,
                    x: 14,
                    y: 6,
                },
            },
            {
                metadata: {
                    inputs: [],
                    settings: {
                        content: {
                            settings: {
                                content: `<div style='line-height:50px'>
<span style='font-size:16px;font-weight:bold'>AZURE STORAGE INVENTORY - </span>
<span>This section gives you an overview of your Azure Storage usage.</span>
</div>`,
                                subtitle: "",
                                title: "",
                            },
                        },
                    },
                    type: "Extension/HubsExtension/PartType/MarkdownPart",
                },
                position: {
                    colSpan: 18,
                    rowSpan: 1,
                    x: 0,
                    y: 9,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of storage accounts by status",
                        },
                        {
                            name: "query",
                            value: ` where type == "microsoft.storage/storageaccounts"
 | summarize StorageCount=count() by PrimaryStatus=tostring(properties.statusOfPrimary), SecondaryStatus=tostring(properties.statusOfSecondary)
 | extend SecondaryStatus= iff(strlen(SecondaryStatus) == 0, "No secondary", SecondaryStatus)
 | extend ['Count (Storage accounts)']=StorageCount
 | project-away StorageCount`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryGridTile",
                },
                position: {
                    colSpan: 6,
                    rowSpan: 3,
                    x: 0,
                    y: 10,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Sum of all disk sizes (GB)",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.compute/disks"
| extend SizeGB = tolong(properties.diskSizeGB)
| summarize ['Total Disk Size (GB)']=sum(SizeGB)`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 3,
                    rowSpan: 3,
                    x: 6,
                    y: 10,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Disks (count) by disk state",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.compute/disks"
| summarize DiskCount=count() by State=tostring(properties.diskState)
| order by DiskCount desc
| extend ["Count (Disks)"]=DiskCount
| project State, ["Count (Disks)"]`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 5,
                    rowSpan: 3,
                    x: 9,
                    y: 10,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of storage accounts by type",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.storage/storageaccounts" or type=="microsoft.classicstorage/storageaccounts"
| summarize StorageCount=count() by type
| extend ['Count (Storage Accounts)']=StorageCount
| extend Type = iff(type == "microsoft.storage/storageaccounts", "Resource Manager", "Classic")
| project Type,['Count (Storage Accounts)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 4,
                    rowSpan: 3,
                    x: 14,
                    y: 10,
                },
            },
            {
                metadata: {
                    inputs: [],
                    settings: {
                        content: {
                            settings: {
                                content: `<div style='line-height:50px'>
<span style='font-size:16px;font-weight:bold'>AZURE NETWORKING INVENTORY - </span>
<span>This section gives you an overview of your Azure networking usage.</span>
</div>`,
                                subtitle: "",
                                title: "",
                            },
                        },
                    },
                    type: "Extension/HubsExtension/PartType/MarkdownPart",
                },
                position: {
                    colSpan: 18,
                    rowSpan: 1,
                    x: 0,
                    y: 13,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of virtual networks (includes classic)",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.network/virtualnetworks"or type == "microsoft.classicnetwork/virtualnetworks"
| summarize ['Virtual networks']=count()`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 4,
                    rowSpan: 3,
                    x: 0,
                    y: 14,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of network interfaces",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.network/networkinterfaces"
| summarize ['Network interfaces']=count()`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 3,
                    rowSpan: 3,
                    x: 4,
                    y: 14,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Total public IPs",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.network/publicipaddresses"
| summarize ['Number of public IP addresses']=count()`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 3,
                    rowSpan: 3,
                    x: 7,
                    y: 14,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Virtual networks by type",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.network/virtualnetworks" or type=="microsoft.classicnetwork/virtualnetworks"
| summarize VNetCount=count() by type
| extend ['Count (Virtual Networks)']=VNetCount
| extend Type = iff(type == "microsoft.network/virtualnetworks", "Resource Manager", "Classic")
| project Type,['Count (Virtual Networks)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 8,
                    rowSpan: 3,
                    x: 10,
                    y: 14,
                },
            },
            {
                metadata: {
                    inputs: [],
                    settings: {
                        content: {
                            settings: {
                                content: `<div style='line-height:50px'>
<span style='font-size:16px;font-weight:bold'>AZURE SQL DATABASES INVENTORY - </span>
<span>This section gives you an overview of your Azure SQL database usage.</span>
</div>`,
                                subtitle: "",
                                title: "",
                            },
                        },
                    },
                    type: "Extension/HubsExtension/PartType/MarkdownPart",
                },
                position: {
                    colSpan: 18,
                    rowSpan: 1,
                    x: 0,
                    y: 17,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of SQL databases",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.sql/servers/databases"
| summarize DBCount=count()`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 3,
                    rowSpan: 3,
                    x: 0,
                    y: 18,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "SQL databases (count) by tier",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.sql/servers/databases"
| summarize DBCount=count() by Tier=tostring(properties.currentSku.tier)
| order by DBCount desc
|extend ['Count (SQL Databases)']=DBCount
| project Tier, ['Count (SQL Databases)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 6,
                    rowSpan: 3,
                    x: 3,
                    y: 18,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "SQL databases (count) by max size (GB)",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.sql/servers/databases"
| extend GB = todouble(properties.maxSizeBytes) / (1024 * 1024 * 1024)
| summarize DBCount=count() by GB
| order by GB desc
|extend ['Count (SQL Databases)']=DBCount
| project GB=strcat(tostring(GB), " GB"), ['Count (SQL Databases)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 9,
                    rowSpan: 3,
                    x: 9,
                    y: 18,
                },
            },
            {
                metadata: {
                    inputs: [],
                    settings: {
                        content: {
                            settings: {
                                content: `<div style='line-height:50px'>
<span style='font-size:16px;font-weight:bold'>APP SERVICE INVENTORY - </span>
<span>This section gives you an overview of your Azure App Service usage.</span>
</div>`,
                                subtitle: "",
                                title: "",
                            },
                        },
                    },
                    type: "Extension/HubsExtension/PartType/MarkdownPart",
                },
                position: {
                    colSpan: 18,
                    rowSpan: 1,
                    x: 0,
                    y: 21,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of AppService apps",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.web/sites"
| summarize SiteCount=count() 
|extend ['Count (AppService Apps)']=SiteCount
| project ['Count (AppService Apps)']`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 3,
                    rowSpan: 3,
                    x: 0,
                    y: 22,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "Count of AppService plans",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.web/serverfarms"
| summarize serverFarmCount=count() 
|extend ['Count (AppService plans)']=serverFarmCount
| project ['Count (AppService plans)']`,
                        },
                        {
                            isOptional: true,
                            name: "chartType",
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                },
                position: {
                    colSpan: 3,
                    rowSpan: 3,
                    x: 3,
                    y: 22,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "AppService Apps by status",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.web/sites"
| summarize SiteCount=count() by Status=tostring(properties.state) 
| order by SiteCount desc
|extend ['Count (AppService Apps)']=SiteCount
| project Status, ['Count (AppService Apps)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 4,
                    rowSpan: 3,
                    x: 6,
                    y: 22,
                },
            },
            {
                metadata: {
                    inputs: [
                        {
                            name: "partTitle",
                            value: "AppService apps by kind",
                        },
                        {
                            name: "query",
                            value: `where type == "microsoft.web/sites"
| summarize SiteCount=count() by kind 
| order by SiteCount desc
|extend ['Count (AppService Apps)']=SiteCount
| project Kind=kind, ['Count (AppService Apps)']`,
                        },
                        {
                            name: "chartType",
                            value: 1,
                        },
                    ],
                    settings: {},
                    type: "Extension/HubsExtension/PartType/ArgQueryChartTile",
                },
                position: {
                    colSpan: 8,
                    rowSpan: 3,
                    x: 10,
                    y: 22,
                },
            },
        ],
    }],
    location: "westeurope",
    metadata: {
        model: {
            timeRange: {
                type: "MsPortalFx.Composition.Configuration.ValueTypes.TimeRange",
                value: {
                    relative: {
                        duration: 24,
                        timeUnit: 1,
                    },
                },
            },
        },
    },
});
