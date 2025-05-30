// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.PowerPlatform
{
    public static class GetAccount
    {
        /// <summary>
        /// Get information about an account.
        /// 
        /// Uses Azure REST API version 2020-10-30-preview.
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azure-native:powerplatform:getAccount", args ?? new GetAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about an account.
        /// 
        /// Uses Azure REST API version 2020-10-30-preview.
        /// </summary>
        public static Output<GetAccountResult> Invoke(GetAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("azure-native:powerplatform:getAccount", args ?? new GetAccountInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about an account.
        /// 
        /// Uses Azure REST API version 2020-10-30-preview.
        /// </summary>
        public static Output<GetAccountResult> Invoke(GetAccountInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("azure-native:powerplatform:getAccount", args ?? new GetAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountArgs()
        {
        }
        public static new GetAccountArgs Empty => new GetAccountArgs();
    }

    public sealed class GetAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAccountInvokeArgs()
        {
        }
        public static new GetAccountInvokeArgs Empty => new GetAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The description of the account.
        /// </summary>
        public readonly string? Description;
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
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The internally assigned unique identifier of the resource.
        /// </summary>
        public readonly string SystemId;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAccountResult(
            string azureApiVersion,

            string? description,

            string id,

            string location,

            string name,

            Outputs.SystemDataResponse systemData,

            string systemId,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Description = description;
            Id = id;
            Location = location;
            Name = name;
            SystemData = systemData;
            SystemId = systemId;
            Tags = tags;
            Type = type;
        }
    }
}
