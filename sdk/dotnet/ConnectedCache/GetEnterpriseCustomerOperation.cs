// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedCache
{
    public static class GetEnterpriseCustomerOperation
    {
        /// <summary>
        /// Retrieves the properties of a Enterprise customer
        /// 
        /// Uses Azure REST API version 2023-05-01-preview.
        /// </summary>
        public static Task<GetEnterpriseCustomerOperationResult> InvokeAsync(GetEnterpriseCustomerOperationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnterpriseCustomerOperationResult>("azure-native:connectedcache:getEnterpriseCustomerOperation", args ?? new GetEnterpriseCustomerOperationArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of a Enterprise customer
        /// 
        /// Uses Azure REST API version 2023-05-01-preview.
        /// </summary>
        public static Output<GetEnterpriseCustomerOperationResult> Invoke(GetEnterpriseCustomerOperationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnterpriseCustomerOperationResult>("azure-native:connectedcache:getEnterpriseCustomerOperation", args ?? new GetEnterpriseCustomerOperationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of a Enterprise customer
        /// 
        /// Uses Azure REST API version 2023-05-01-preview.
        /// </summary>
        public static Output<GetEnterpriseCustomerOperationResult> Invoke(GetEnterpriseCustomerOperationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnterpriseCustomerOperationResult>("azure-native:connectedcache:getEnterpriseCustomerOperation", args ?? new GetEnterpriseCustomerOperationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnterpriseCustomerOperationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Customer resource
        /// </summary>
        [Input("customerResourceName", required: true)]
        public string CustomerResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetEnterpriseCustomerOperationArgs()
        {
        }
        public static new GetEnterpriseCustomerOperationArgs Empty => new GetEnterpriseCustomerOperationArgs();
    }

    public sealed class GetEnterpriseCustomerOperationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Customer resource
        /// </summary>
        [Input("customerResourceName", required: true)]
        public Input<string> CustomerResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetEnterpriseCustomerOperationInvokeArgs()
        {
        }
        public static new GetEnterpriseCustomerOperationInvokeArgs Empty => new GetEnterpriseCustomerOperationInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnterpriseCustomerOperationResult
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
        public readonly Outputs.CacheNodeOldResponseResponse Properties;
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
        private GetEnterpriseCustomerOperationResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.CacheNodeOldResponseResponse properties,

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
