// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector
{
    public static class GetMacieAllowList
    {
        /// <summary>
        /// Get a MacieAllowList
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Task<GetMacieAllowListResult> InvokeAsync(GetMacieAllowListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMacieAllowListResult>("azure-native:awsconnector:getMacieAllowList", args ?? new GetMacieAllowListArgs(), options.WithDefaults());

        /// <summary>
        /// Get a MacieAllowList
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Output<GetMacieAllowListResult> Invoke(GetMacieAllowListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMacieAllowListResult>("azure-native:awsconnector:getMacieAllowList", args ?? new GetMacieAllowListInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a MacieAllowList
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Output<GetMacieAllowListResult> Invoke(GetMacieAllowListInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMacieAllowListResult>("azure-native:awsconnector:getMacieAllowList", args ?? new GetMacieAllowListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMacieAllowListArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of MacieAllowList
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMacieAllowListArgs()
        {
        }
        public static new GetMacieAllowListArgs Empty => new GetMacieAllowListArgs();
    }

    public sealed class GetMacieAllowListInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of MacieAllowList
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMacieAllowListInvokeArgs()
        {
        }
        public static new GetMacieAllowListInvokeArgs Empty => new GetMacieAllowListInvokeArgs();
    }


    [OutputType]
    public sealed class GetMacieAllowListResult
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
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        public readonly Outputs.MacieAllowListPropertiesResponse Properties;
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
        private GetMacieAllowListResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.MacieAllowListPropertiesResponse properties,

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
