// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData
{
    public static class GetSqlServerInstance
    {
        /// <summary>
        /// Retrieves a SQL Server Instance resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSqlServerInstanceResult> InvokeAsync(GetSqlServerInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSqlServerInstanceResult>("azure-native:azurearcdata:getSqlServerInstance", args ?? new GetSqlServerInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a SQL Server Instance resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSqlServerInstanceResult> Invoke(GetSqlServerInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSqlServerInstanceResult>("azure-native:azurearcdata:getSqlServerInstance", args ?? new GetSqlServerInstanceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a SQL Server Instance resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSqlServerInstanceResult> Invoke(GetSqlServerInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSqlServerInstanceResult>("azure-native:azurearcdata:getSqlServerInstance", args ?? new GetSqlServerInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSqlServerInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of SQL Server Instance
        /// </summary>
        [Input("sqlServerInstanceName", required: true)]
        public string SqlServerInstanceName { get; set; } = null!;

        public GetSqlServerInstanceArgs()
        {
        }
        public static new GetSqlServerInstanceArgs Empty => new GetSqlServerInstanceArgs();
    }

    public sealed class GetSqlServerInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of SQL Server Instance
        /// </summary>
        [Input("sqlServerInstanceName", required: true)]
        public Input<string> SqlServerInstanceName { get; set; } = null!;

        public GetSqlServerInstanceInvokeArgs()
        {
        }
        public static new GetSqlServerInstanceInvokeArgs Empty => new GetSqlServerInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSqlServerInstanceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// null
        /// </summary>
        public readonly Outputs.SqlServerInstancePropertiesResponse Properties;
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
        private GetSqlServerInstanceResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.SqlServerInstancePropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
