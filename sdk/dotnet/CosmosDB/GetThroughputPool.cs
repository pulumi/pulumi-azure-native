// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    public static class GetThroughputPool
    {
        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Throughput Pool
        /// 
        /// Uses Azure REST API version 2024-12-01-preview.
        /// 
        /// Other available API versions: 2023-11-15-preview, 2024-02-15-preview, 2024-05-15-preview, 2024-09-01-preview, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetThroughputPoolResult> InvokeAsync(GetThroughputPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetThroughputPoolResult>("azure-native:cosmosdb:getThroughputPool", args ?? new GetThroughputPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Throughput Pool
        /// 
        /// Uses Azure REST API version 2024-12-01-preview.
        /// 
        /// Other available API versions: 2023-11-15-preview, 2024-02-15-preview, 2024-05-15-preview, 2024-09-01-preview, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetThroughputPoolResult> Invoke(GetThroughputPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetThroughputPoolResult>("azure-native:cosmosdb:getThroughputPool", args ?? new GetThroughputPoolInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of an existing Azure Cosmos DB Throughput Pool
        /// 
        /// Uses Azure REST API version 2024-12-01-preview.
        /// 
        /// Other available API versions: 2023-11-15-preview, 2024-02-15-preview, 2024-05-15-preview, 2024-09-01-preview, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetThroughputPoolResult> Invoke(GetThroughputPoolInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetThroughputPoolResult>("azure-native:cosmosdb:getThroughputPool", args ?? new GetThroughputPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetThroughputPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB Throughput Pool name.
        /// </summary>
        [Input("throughputPoolName", required: true)]
        public string ThroughputPoolName { get; set; } = null!;

        public GetThroughputPoolArgs()
        {
        }
        public static new GetThroughputPoolArgs Empty => new GetThroughputPoolArgs();
    }

    public sealed class GetThroughputPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB Throughput Pool name.
        /// </summary>
        [Input("throughputPoolName", required: true)]
        public Input<string> ThroughputPoolName { get; set; } = null!;

        public GetThroughputPoolInvokeArgs()
        {
        }
        public static new GetThroughputPoolInvokeArgs Empty => new GetThroughputPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetThroughputPoolResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Value for throughput to be shared among CosmosDB resources in the pool.
        /// </summary>
        public readonly int? MaxThroughput;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A provisioning state of the ThroughputPool.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetThroughputPoolResult(
            string azureApiVersion,

            string id,

            string location,

            int? maxThroughput,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            MaxThroughput = maxThroughput;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
