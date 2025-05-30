// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse
{
    public static class GetEventHubDataConnection
    {
        /// <summary>
        /// Returns a data connection.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Task<GetEventHubDataConnectionResult> InvokeAsync(GetEventHubDataConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEventHubDataConnectionResult>("azure-native:synapse:getEventHubDataConnection", args ?? new GetEventHubDataConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a data connection.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Output<GetEventHubDataConnectionResult> Invoke(GetEventHubDataConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventHubDataConnectionResult>("azure-native:synapse:getEventHubDataConnection", args ?? new GetEventHubDataConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a data connection.
        /// 
        /// Uses Azure REST API version 2021-06-01-preview.
        /// </summary>
        public static Output<GetEventHubDataConnectionResult> Invoke(GetEventHubDataConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventHubDataConnectionResult>("azure-native:synapse:getEventHubDataConnection", args ?? new GetEventHubDataConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventHubDataConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the data connection.
        /// </summary>
        [Input("dataConnectionName", required: true)]
        public string DataConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the database in the Kusto pool.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName", required: true)]
        public string KustoPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetEventHubDataConnectionArgs()
        {
        }
        public static new GetEventHubDataConnectionArgs Empty => new GetEventHubDataConnectionArgs();
    }

    public sealed class GetEventHubDataConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the data connection.
        /// </summary>
        [Input("dataConnectionName", required: true)]
        public Input<string> DataConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the database in the Kusto pool.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName", required: true)]
        public Input<string> KustoPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetEventHubDataConnectionInvokeArgs()
        {
        }
        public static new GetEventHubDataConnectionInvokeArgs Empty => new GetEventHubDataConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetEventHubDataConnectionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The event hub messages compression type
        /// </summary>
        public readonly string? Compression;
        /// <summary>
        /// The event hub consumer group.
        /// </summary>
        public readonly string ConsumerGroup;
        /// <summary>
        /// The data format of the message. Optionally the data format can be added to each message.
        /// </summary>
        public readonly string? DataFormat;
        /// <summary>
        /// The resource ID of the event hub to be used to create a data connection.
        /// </summary>
        public readonly string EventHubResourceId;
        /// <summary>
        /// System properties of the event hub
        /// </summary>
        public readonly ImmutableArray<string> EventSystemProperties;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of the endpoint for the data connection
        /// Expected value is 'EventHub'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The resource ID of a managed identity (system or user assigned) to be used to authenticate with event hub.
        /// </summary>
        public readonly string? ManagedIdentityResourceId;
        /// <summary>
        /// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
        /// </summary>
        public readonly string? MappingRuleName;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The table where the data should be ingested. Optionally the table information can be added to each message.
        /// </summary>
        public readonly string? TableName;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetEventHubDataConnectionResult(
            string azureApiVersion,

            string? compression,

            string consumerGroup,

            string? dataFormat,

            string eventHubResourceId,

            ImmutableArray<string> eventSystemProperties,

            string id,

            string kind,

            string? location,

            string? managedIdentityResourceId,

            string? mappingRuleName,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string? tableName,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Compression = compression;
            ConsumerGroup = consumerGroup;
            DataFormat = dataFormat;
            EventHubResourceId = eventHubResourceId;
            EventSystemProperties = eventSystemProperties;
            Id = id;
            Kind = kind;
            Location = location;
            ManagedIdentityResourceId = managedIdentityResourceId;
            MappingRuleName = mappingRuleName;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            TableName = tableName;
            Type = type;
        }
    }
}
