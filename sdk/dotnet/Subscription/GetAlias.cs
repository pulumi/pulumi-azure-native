// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Subscription
{
    public static class GetAlias
    {
        /// <summary>
        /// Get Alias Subscription.
        /// 
        /// Uses Azure REST API version 2024-08-01-preview.
        /// 
        /// Other available API versions: 2021-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native subscription [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAliasResult> InvokeAsync(GetAliasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAliasResult>("azure-native:subscription:getAlias", args ?? new GetAliasArgs(), options.WithDefaults());

        /// <summary>
        /// Get Alias Subscription.
        /// 
        /// Uses Azure REST API version 2024-08-01-preview.
        /// 
        /// Other available API versions: 2021-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native subscription [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAliasResult> Invoke(GetAliasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAliasResult>("azure-native:subscription:getAlias", args ?? new GetAliasInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get Alias Subscription.
        /// 
        /// Uses Azure REST API version 2024-08-01-preview.
        /// 
        /// Other available API versions: 2021-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native subscription [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAliasResult> Invoke(GetAliasInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAliasResult>("azure-native:subscription:getAlias", args ?? new GetAliasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAliasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// AliasName is the name for the subscription creation request. Note that this is not the same as subscription name and this doesn’t have any other lifecycle need beyond the request for subscription creation.
        /// </summary>
        [Input("aliasName", required: true)]
        public string AliasName { get; set; } = null!;

        public GetAliasArgs()
        {
        }
        public static new GetAliasArgs Empty => new GetAliasArgs();
    }

    public sealed class GetAliasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// AliasName is the name for the subscription creation request. Note that this is not the same as subscription name and this doesn’t have any other lifecycle need beyond the request for subscription creation.
        /// </summary>
        [Input("aliasName", required: true)]
        public Input<string> AliasName { get; set; } = null!;

        public GetAliasInvokeArgs()
        {
        }
        public static new GetAliasInvokeArgs Empty => new GetAliasInvokeArgs();
    }


    [OutputType]
    public sealed class GetAliasResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified ID for the alias resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Alias ID.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Subscription Alias response properties.
        /// </summary>
        public readonly Outputs.SubscriptionAliasResponsePropertiesResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource type, Microsoft.Subscription/aliases.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAliasResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.SubscriptionAliasResponsePropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
