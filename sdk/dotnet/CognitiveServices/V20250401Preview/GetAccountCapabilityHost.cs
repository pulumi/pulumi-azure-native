// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices.V20250401Preview
{
    public static class GetAccountCapabilityHost
    {
        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// </summary>
        public static Task<GetAccountCapabilityHostResult> InvokeAsync(GetAccountCapabilityHostArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountCapabilityHostResult>("azure-native:cognitiveservices/v20250401preview:getAccountCapabilityHost", args ?? new GetAccountCapabilityHostArgs(), options.WithDefaults());

        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// </summary>
        public static Output<GetAccountCapabilityHostResult> Invoke(GetAccountCapabilityHostInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountCapabilityHostResult>("azure-native:cognitiveservices/v20250401preview:getAccountCapabilityHost", args ?? new GetAccountCapabilityHostInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Azure Resource Manager resource envelope.
        /// </summary>
        public static Output<GetAccountCapabilityHostResult> Invoke(GetAccountCapabilityHostInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountCapabilityHostResult>("azure-native:cognitiveservices/v20250401preview:getAccountCapabilityHost", args ?? new GetAccountCapabilityHostInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountCapabilityHostArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of Cognitive Services account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the capability host associated with the Cognitive Services Resource
        /// </summary>
        [Input("capabilityHostName", required: true)]
        public string CapabilityHostName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountCapabilityHostArgs()
        {
        }
        public static new GetAccountCapabilityHostArgs Empty => new GetAccountCapabilityHostArgs();
    }

    public sealed class GetAccountCapabilityHostInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of Cognitive Services account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the capability host associated with the Cognitive Services Resource
        /// </summary>
        [Input("capabilityHostName", required: true)]
        public Input<string> CapabilityHostName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAccountCapabilityHostInvokeArgs()
        {
        }
        public static new GetAccountCapabilityHostInvokeArgs Empty => new GetAccountCapabilityHostInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountCapabilityHostResult
    {
        /// <summary>
        /// [Required] Additional attributes of the entity.
        /// </summary>
        public readonly Outputs.CapabilityHostResponse CapabilityHostProperties;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAccountCapabilityHostResult(
            Outputs.CapabilityHostResponse capabilityHostProperties,

            string id,

            string name,

            string type)
        {
            CapabilityHostProperties = capabilityHostProperties;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
