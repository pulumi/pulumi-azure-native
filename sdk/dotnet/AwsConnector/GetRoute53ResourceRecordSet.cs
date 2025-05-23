// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector
{
    public static class GetRoute53ResourceRecordSet
    {
        /// <summary>
        /// Get a Route53ResourceRecordSet
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Task<GetRoute53ResourceRecordSetResult> InvokeAsync(GetRoute53ResourceRecordSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRoute53ResourceRecordSetResult>("azure-native:awsconnector:getRoute53ResourceRecordSet", args ?? new GetRoute53ResourceRecordSetArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Route53ResourceRecordSet
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Output<GetRoute53ResourceRecordSetResult> Invoke(GetRoute53ResourceRecordSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoute53ResourceRecordSetResult>("azure-native:awsconnector:getRoute53ResourceRecordSet", args ?? new GetRoute53ResourceRecordSetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Route53ResourceRecordSet
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Output<GetRoute53ResourceRecordSetResult> Invoke(GetRoute53ResourceRecordSetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoute53ResourceRecordSetResult>("azure-native:awsconnector:getRoute53ResourceRecordSet", args ?? new GetRoute53ResourceRecordSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRoute53ResourceRecordSetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Route53ResourceRecordSet
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRoute53ResourceRecordSetArgs()
        {
        }
        public static new GetRoute53ResourceRecordSetArgs Empty => new GetRoute53ResourceRecordSetArgs();
    }

    public sealed class GetRoute53ResourceRecordSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Route53ResourceRecordSet
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetRoute53ResourceRecordSetInvokeArgs()
        {
        }
        public static new GetRoute53ResourceRecordSetInvokeArgs Empty => new GetRoute53ResourceRecordSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetRoute53ResourceRecordSetResult
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
        public readonly Outputs.Route53ResourceRecordSetPropertiesResponse Properties;
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
        private GetRoute53ResourceRecordSetResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.Route53ResourceRecordSetPropertiesResponse properties,

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
