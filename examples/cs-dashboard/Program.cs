// Copyright 2024, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Text.Json.Nodes;

using Pulumi;
using Pulumi.AzureNative.Resources;
using Pulumi.AzureNative.Portal;
using Pulumi.AzureNative.Portal.Inputs;

return await Pulumi.Deployment.RunAsync(() =>
{
    var resourceGroup = new ResourceGroup("dashrg");

    var dashboard = new Dashboard("dash", new()
    {
        ResourceGroupName = resourceGroup.Name,
        Lenses = new[]
        {
            new DashboardLensArgs
            {
                Order = 0,
                Parts = new[]
                {
                    new DashboardPartsArgs
                    {
                        Metadata = new DashboardPartMetadataArgs
                        {
                            Inputs = new InputList<object>(),
                            Settings = {
                                ["content"] = new Dictionary<string, object> {
                                    ["settings"] = new Dictionary<string, object> {
                                        ["content"] =  @"<div style='line-height:50px;'>
<span style='font-size:16px;font-weight:bold'>AZURE RESOURCE INVENTORY - </span>
<span>This section gives you an overview of all your Azure resources across all subscriptions that you can access.</span>
</div>",
                                        ["subtitle"] = "",
                                        ["title"] = "",
                                    }
                                }
                            },
                            Type = "Extension/HubsExtension/PartType/MarkdownPart",
                        },
                        Position = new DashboardPartsPositionArgs
                        {
                            ColSpan = 18,
                            RowSpan = 1,
                            X = 0,
                            Y = 0,
                        },
                    },
                    new DashboardPartsArgs
                    {
                        Metadata = new DashboardPartMetadataArgs
                        {
                            Inputs = new[]
                            {
                                new Dictionary<string, object>
                                {
                                    { "name", "partTitle" },
                                    { "value", "Count of all my Azure resources" },
                                },
                                new Dictionary<string, object>                                
                                {
                                    { "name", "query" },
                                    { "value", "summarize Resources=count()" },
                                },
                                new Dictionary<string, object>                                
                                {
                                    { "isOptional", true },
                                    { "name", "chartType" },
                                },
                                new Dictionary<string, object>                                
                                {
                                    { "isOptional", true },
                                    { "name", "isShared" },
                                },
                                new Dictionary<string, object>                                
                                {
                                    { "isOptional", true },
                                    { "name", "queryId" },
                                },
                                new Dictionary<string, object>                                
                                {
                                    { "isOptional", true },
                                    { "name", "formatResults" },
                                },
                                new Dictionary<string, object>                                
                                {
                                    { "isOptional", true },
                                    { "name", "queryScope" },
                                },
                            },
                            Settings = new Dictionary<string, object>(),
                            Type = "Extension/HubsExtension/PartType/ArgQuerySingleValueTile",
                        },
                        Position = new DashboardPartsPositionArgs
                        {
                            ColSpan = 3,
                            RowSpan = 4,
                            X = 0,
                            Y = 1,
                        },
                    },
                },
            },
        },
        Metadata = new Dictionary<string, object> {
            ["model"] = new Dictionary<string, object> {
                ["timeRange"] = new Dictionary<string, object> {
                    ["type"] = "MsPortalFx.Composition.Configuration.ValueTypes.TimeRange",
                    ["value"] = new Dictionary<string, object> {
                        ["relative"] = new Dictionary<string, object> {
                            ["duration"] = 24,
                            ["timeUnit"] = 1
                        }
                    }
                }
            },
        },
        Tags = 
        {
            ["hidden-title"] = "Test Dashboard",
        },
    });

    return new Dictionary<string, object?>
    {
    };
});
