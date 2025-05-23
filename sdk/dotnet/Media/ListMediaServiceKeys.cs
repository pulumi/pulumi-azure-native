// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media
{
    public static class ListMediaServiceKeys
    {
        /// <summary>
        /// Lists the keys for a Media Service.
        /// 
        /// Uses Azure REST API version 2015-10-01.
        /// </summary>
        public static Task<ListMediaServiceKeysResult> InvokeAsync(ListMediaServiceKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListMediaServiceKeysResult>("azure-native:media:listMediaServiceKeys", args ?? new ListMediaServiceKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the keys for a Media Service.
        /// 
        /// Uses Azure REST API version 2015-10-01.
        /// </summary>
        public static Output<ListMediaServiceKeysResult> Invoke(ListMediaServiceKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListMediaServiceKeysResult>("azure-native:media:listMediaServiceKeys", args ?? new ListMediaServiceKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the keys for a Media Service.
        /// 
        /// Uses Azure REST API version 2015-10-01.
        /// </summary>
        public static Output<ListMediaServiceKeysResult> Invoke(ListMediaServiceKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListMediaServiceKeysResult>("azure-native:media:listMediaServiceKeys", args ?? new ListMediaServiceKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListMediaServiceKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Media Service.
        /// </summary>
        [Input("mediaServiceName", required: true)]
        public string MediaServiceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListMediaServiceKeysArgs()
        {
        }
        public static new ListMediaServiceKeysArgs Empty => new ListMediaServiceKeysArgs();
    }

    public sealed class ListMediaServiceKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Media Service.
        /// </summary>
        [Input("mediaServiceName", required: true)]
        public Input<string> MediaServiceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListMediaServiceKeysInvokeArgs()
        {
        }
        public static new ListMediaServiceKeysInvokeArgs Empty => new ListMediaServiceKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListMediaServiceKeysResult
    {
        /// <summary>
        /// The primary authorization endpoint.
        /// </summary>
        public readonly string? PrimaryAuthEndpoint;
        /// <summary>
        /// The primary key for the Media Service resource.
        /// </summary>
        public readonly string? PrimaryKey;
        /// <summary>
        /// The authorization scope.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// The secondary authorization endpoint.
        /// </summary>
        public readonly string? SecondaryAuthEndpoint;
        /// <summary>
        /// The secondary key for the Media Service resource.
        /// </summary>
        public readonly string? SecondaryKey;

        [OutputConstructor]
        private ListMediaServiceKeysResult(
            string? primaryAuthEndpoint,

            string? primaryKey,

            string? scope,

            string? secondaryAuthEndpoint,

            string? secondaryKey)
        {
            PrimaryAuthEndpoint = primaryAuthEndpoint;
            PrimaryKey = primaryKey;
            Scope = scope;
            SecondaryAuthEndpoint = secondaryAuthEndpoint;
            SecondaryKey = secondaryKey;
        }
    }
}
