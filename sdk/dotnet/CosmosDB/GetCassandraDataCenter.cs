// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    public static class GetCassandraDataCenter
    {
        /// <summary>
        /// Get the properties of a managed Cassandra data center.
        /// 
        /// Uses Azure REST API version 2024-11-15.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2021-04-01-preview, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetCassandraDataCenterResult> InvokeAsync(GetCassandraDataCenterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCassandraDataCenterResult>("azure-native:cosmosdb:getCassandraDataCenter", args ?? new GetCassandraDataCenterArgs(), options.WithDefaults());

        /// <summary>
        /// Get the properties of a managed Cassandra data center.
        /// 
        /// Uses Azure REST API version 2024-11-15.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2021-04-01-preview, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetCassandraDataCenterResult> Invoke(GetCassandraDataCenterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCassandraDataCenterResult>("azure-native:cosmosdb:getCassandraDataCenter", args ?? new GetCassandraDataCenterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the properties of a managed Cassandra data center.
        /// 
        /// Uses Azure REST API version 2024-11-15.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2021-04-01-preview, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetCassandraDataCenterResult> Invoke(GetCassandraDataCenterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCassandraDataCenterResult>("azure-native:cosmosdb:getCassandraDataCenter", args ?? new GetCassandraDataCenterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCassandraDataCenterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Managed Cassandra cluster name.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// Data center name in a managed Cassandra cluster.
        /// </summary>
        [Input("dataCenterName", required: true)]
        public string DataCenterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCassandraDataCenterArgs()
        {
        }
        public static new GetCassandraDataCenterArgs Empty => new GetCassandraDataCenterArgs();
    }

    public sealed class GetCassandraDataCenterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Managed Cassandra cluster name.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Data center name in a managed Cassandra cluster.
        /// </summary>
        [Input("dataCenterName", required: true)]
        public Input<string> DataCenterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetCassandraDataCenterInvokeArgs()
        {
        }
        public static new GetCassandraDataCenterInvokeArgs Empty => new GetCassandraDataCenterInvokeArgs();
    }


    [OutputType]
    public sealed class GetCassandraDataCenterResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The unique resource identifier of the database account.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the database account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of a managed Cassandra data center.
        /// </summary>
        public readonly Outputs.DataCenterResourceResponseProperties Properties;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCassandraDataCenterResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.DataCenterResourceResponseProperties properties,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
