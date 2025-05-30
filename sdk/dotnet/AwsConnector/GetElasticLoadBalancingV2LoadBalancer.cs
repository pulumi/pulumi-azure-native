// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector
{
    public static class GetElasticLoadBalancingV2LoadBalancer
    {
        /// <summary>
        /// Get a ElasticLoadBalancingV2LoadBalancer
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Task<GetElasticLoadBalancingV2LoadBalancerResult> InvokeAsync(GetElasticLoadBalancingV2LoadBalancerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetElasticLoadBalancingV2LoadBalancerResult>("azure-native:awsconnector:getElasticLoadBalancingV2LoadBalancer", args ?? new GetElasticLoadBalancingV2LoadBalancerArgs(), options.WithDefaults());

        /// <summary>
        /// Get a ElasticLoadBalancingV2LoadBalancer
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Output<GetElasticLoadBalancingV2LoadBalancerResult> Invoke(GetElasticLoadBalancingV2LoadBalancerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetElasticLoadBalancingV2LoadBalancerResult>("azure-native:awsconnector:getElasticLoadBalancingV2LoadBalancer", args ?? new GetElasticLoadBalancingV2LoadBalancerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a ElasticLoadBalancingV2LoadBalancer
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// </summary>
        public static Output<GetElasticLoadBalancingV2LoadBalancerResult> Invoke(GetElasticLoadBalancingV2LoadBalancerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetElasticLoadBalancingV2LoadBalancerResult>("azure-native:awsconnector:getElasticLoadBalancingV2LoadBalancer", args ?? new GetElasticLoadBalancingV2LoadBalancerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetElasticLoadBalancingV2LoadBalancerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of ElasticLoadBalancingV2LoadBalancer
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetElasticLoadBalancingV2LoadBalancerArgs()
        {
        }
        public static new GetElasticLoadBalancingV2LoadBalancerArgs Empty => new GetElasticLoadBalancingV2LoadBalancerArgs();
    }

    public sealed class GetElasticLoadBalancingV2LoadBalancerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of ElasticLoadBalancingV2LoadBalancer
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetElasticLoadBalancingV2LoadBalancerInvokeArgs()
        {
        }
        public static new GetElasticLoadBalancingV2LoadBalancerInvokeArgs Empty => new GetElasticLoadBalancingV2LoadBalancerInvokeArgs();
    }


    [OutputType]
    public sealed class GetElasticLoadBalancingV2LoadBalancerResult
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
        public readonly Outputs.ElasticLoadBalancingV2LoadBalancerPropertiesResponse Properties;
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
        private GetElasticLoadBalancingV2LoadBalancerResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.ElasticLoadBalancingV2LoadBalancerPropertiesResponse properties,

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
