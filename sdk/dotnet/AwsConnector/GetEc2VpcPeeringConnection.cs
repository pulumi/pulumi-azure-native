// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector
{
    public static class GetEc2VpcPeeringConnection
    {
        /// <summary>
        /// Get a Ec2VPCPeeringConnection
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Task<GetEc2VpcPeeringConnectionResult> InvokeAsync(GetEc2VpcPeeringConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEc2VpcPeeringConnectionResult>("azure-native:awsconnector:getEc2VpcPeeringConnection", args ?? new GetEc2VpcPeeringConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Ec2VPCPeeringConnection
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Output<GetEc2VpcPeeringConnectionResult> Invoke(GetEc2VpcPeeringConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEc2VpcPeeringConnectionResult>("azure-native:awsconnector:getEc2VpcPeeringConnection", args ?? new GetEc2VpcPeeringConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Ec2VPCPeeringConnection
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Output<GetEc2VpcPeeringConnectionResult> Invoke(GetEc2VpcPeeringConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEc2VpcPeeringConnectionResult>("azure-native:awsconnector:getEc2VpcPeeringConnection", args ?? new GetEc2VpcPeeringConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEc2VpcPeeringConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Ec2VPCPeeringConnection
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetEc2VpcPeeringConnectionArgs()
        {
        }
        public static new GetEc2VpcPeeringConnectionArgs Empty => new GetEc2VpcPeeringConnectionArgs();
    }

    public sealed class GetEc2VpcPeeringConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Ec2VPCPeeringConnection
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetEc2VpcPeeringConnectionInvokeArgs()
        {
        }
        public static new GetEc2VpcPeeringConnectionInvokeArgs Empty => new GetEc2VpcPeeringConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetEc2VpcPeeringConnectionResult
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
        public readonly Outputs.Ec2VPCPeeringConnectionPropertiesResponse Properties;
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
        private GetEc2VpcPeeringConnectionResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.Ec2VPCPeeringConnectionPropertiesResponse properties,

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
