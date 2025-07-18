// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SqlVirtualMachine
{
    public static class GetSqlVirtualMachineGroup
    {
        /// <summary>
        /// Gets a SQL virtual machine group.
        /// 
        /// Uses Azure REST API version 2023-10-01.
        /// 
        /// Other available API versions: 2022-02-01, 2022-07-01-preview, 2022-08-01-preview, 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sqlvirtualmachine [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSqlVirtualMachineGroupResult> InvokeAsync(GetSqlVirtualMachineGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSqlVirtualMachineGroupResult>("azure-native:sqlvirtualmachine:getSqlVirtualMachineGroup", args ?? new GetSqlVirtualMachineGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a SQL virtual machine group.
        /// 
        /// Uses Azure REST API version 2023-10-01.
        /// 
        /// Other available API versions: 2022-02-01, 2022-07-01-preview, 2022-08-01-preview, 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sqlvirtualmachine [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSqlVirtualMachineGroupResult> Invoke(GetSqlVirtualMachineGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSqlVirtualMachineGroupResult>("azure-native:sqlvirtualmachine:getSqlVirtualMachineGroup", args ?? new GetSqlVirtualMachineGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a SQL virtual machine group.
        /// 
        /// Uses Azure REST API version 2023-10-01.
        /// 
        /// Other available API versions: 2022-02-01, 2022-07-01-preview, 2022-08-01-preview, 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sqlvirtualmachine [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSqlVirtualMachineGroupResult> Invoke(GetSqlVirtualMachineGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSqlVirtualMachineGroupResult>("azure-native:sqlvirtualmachine:getSqlVirtualMachineGroup", args ?? new GetSqlVirtualMachineGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSqlVirtualMachineGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the SQL virtual machine group.
        /// </summary>
        [Input("sqlVirtualMachineGroupName", required: true)]
        public string SqlVirtualMachineGroupName { get; set; } = null!;

        public GetSqlVirtualMachineGroupArgs()
        {
        }
        public static new GetSqlVirtualMachineGroupArgs Empty => new GetSqlVirtualMachineGroupArgs();
    }

    public sealed class GetSqlVirtualMachineGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the SQL virtual machine group.
        /// </summary>
        [Input("sqlVirtualMachineGroupName", required: true)]
        public Input<string> SqlVirtualMachineGroupName { get; set; } = null!;

        public GetSqlVirtualMachineGroupInvokeArgs()
        {
        }
        public static new GetSqlVirtualMachineGroupInvokeArgs Empty => new GetSqlVirtualMachineGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetSqlVirtualMachineGroupResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Cluster type.
        /// </summary>
        public readonly string ClusterConfiguration;
        /// <summary>
        /// Type of cluster manager: Windows Server Failover Cluster (WSFC), implied by the scale type of the group and the OS type.
        /// </summary>
        public readonly string ClusterManagerType;
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
        /// Provisioning state to track the async operation status.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Scale type.
        /// </summary>
        public readonly string ScaleType;
        /// <summary>
        /// SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
        /// </summary>
        public readonly string? SqlImageOffer;
        /// <summary>
        /// SQL image sku.
        /// </summary>
        public readonly string? SqlImageSku;
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
        /// <summary>
        /// Cluster Active Directory domain profile.
        /// </summary>
        public readonly Outputs.WsfcDomainProfileResponse? WsfcDomainProfile;

        [OutputConstructor]
        private GetSqlVirtualMachineGroupResult(
            string azureApiVersion,

            string clusterConfiguration,

            string clusterManagerType,

            string id,

            string location,

            string name,

            string provisioningState,

            string scaleType,

            string? sqlImageOffer,

            string? sqlImageSku,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.WsfcDomainProfileResponse? wsfcDomainProfile)
        {
            AzureApiVersion = azureApiVersion;
            ClusterConfiguration = clusterConfiguration;
            ClusterManagerType = clusterManagerType;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ScaleType = scaleType;
            SqlImageOffer = sqlImageOffer;
            SqlImageSku = sqlImageSku;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            WsfcDomainProfile = wsfcDomainProfile;
        }
    }
}
