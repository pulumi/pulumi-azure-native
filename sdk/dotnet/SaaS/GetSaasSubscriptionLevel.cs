// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SaaS
{
    public static class GetSaasSubscriptionLevel
    {
        /// <summary>
        /// Gets information about the specified Subscription Level SaaS.
        /// 
        /// Uses Azure REST API version 2018-03-01-beta.
        /// </summary>
        public static Task<GetSaasSubscriptionLevelResult> InvokeAsync(GetSaasSubscriptionLevelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSaasSubscriptionLevelResult>("azure-native:saas:getSaasSubscriptionLevel", args ?? new GetSaasSubscriptionLevelArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified Subscription Level SaaS.
        /// 
        /// Uses Azure REST API version 2018-03-01-beta.
        /// </summary>
        public static Output<GetSaasSubscriptionLevelResult> Invoke(GetSaasSubscriptionLevelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSaasSubscriptionLevelResult>("azure-native:saas:getSaasSubscriptionLevel", args ?? new GetSaasSubscriptionLevelInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified Subscription Level SaaS.
        /// 
        /// Uses Azure REST API version 2018-03-01-beta.
        /// </summary>
        public static Output<GetSaasSubscriptionLevelResult> Invoke(GetSaasSubscriptionLevelInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSaasSubscriptionLevelResult>("azure-native:saas:getSaasSubscriptionLevel", args ?? new GetSaasSubscriptionLevelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSaasSubscriptionLevelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetSaasSubscriptionLevelArgs()
        {
        }
        public static new GetSaasSubscriptionLevelArgs Empty => new GetSaasSubscriptionLevelArgs();
    }

    public sealed class GetSaasSubscriptionLevelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetSaasSubscriptionLevelInvokeArgs()
        {
        }
        public static new GetSaasSubscriptionLevelInvokeArgs Empty => new GetSaasSubscriptionLevelInvokeArgs();
    }


    [OutputType]
    public sealed class GetSaasSubscriptionLevelResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The resource uri
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// saas properties
        /// </summary>
        public readonly Outputs.SaasResourceResponseProperties Properties;
        /// <summary>
        /// the resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSaasSubscriptionLevelResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.SaasResourceResponseProperties properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
