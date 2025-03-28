// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Purview.V20210701
{
    public static class ListAccountKeys
    {
        /// <summary>
        /// List the authorization keys associated with this account.
        /// </summary>
        public static Task<ListAccountKeysResult> InvokeAsync(ListAccountKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListAccountKeysResult>("azure-native:purview/v20210701:listAccountKeys", args ?? new ListAccountKeysArgs(), options.WithDefaults());

        /// <summary>
        /// List the authorization keys associated with this account.
        /// </summary>
        public static Output<ListAccountKeysResult> Invoke(ListAccountKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListAccountKeysResult>("azure-native:purview/v20210701:listAccountKeys", args ?? new ListAccountKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List the authorization keys associated with this account.
        /// </summary>
        public static Output<ListAccountKeysResult> Invoke(ListAccountKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListAccountKeysResult>("azure-native:purview/v20210701:listAccountKeys", args ?? new ListAccountKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListAccountKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListAccountKeysArgs()
        {
        }
        public static new ListAccountKeysArgs Empty => new ListAccountKeysArgs();
    }

    public sealed class ListAccountKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListAccountKeysInvokeArgs()
        {
        }
        public static new ListAccountKeysInvokeArgs Empty => new ListAccountKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListAccountKeysResult
    {
        /// <summary>
        /// Gets or sets the primary connection string.
        /// </summary>
        public readonly string? AtlasKafkaPrimaryEndpoint;
        /// <summary>
        /// Gets or sets the secondary connection string.
        /// </summary>
        public readonly string? AtlasKafkaSecondaryEndpoint;

        [OutputConstructor]
        private ListAccountKeysResult(
            string? atlasKafkaPrimaryEndpoint,

            string? atlasKafkaSecondaryEndpoint)
        {
            AtlasKafkaPrimaryEndpoint = atlasKafkaPrimaryEndpoint;
            AtlasKafkaSecondaryEndpoint = atlasKafkaSecondaryEndpoint;
        }
    }
}
