// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cache.V20230401
{
    public static class ListRedisKeys
    {
        /// <summary>
        /// Retrieve a Redis cache's access keys. This operation requires write permission to the cache resource.
        /// </summary>
        public static Task<ListRedisKeysResult> InvokeAsync(ListRedisKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListRedisKeysResult>("azure-native:cache/v20230401:listRedisKeys", args ?? new ListRedisKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve a Redis cache's access keys. This operation requires write permission to the cache resource.
        /// </summary>
        public static Output<ListRedisKeysResult> Invoke(ListRedisKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListRedisKeysResult>("azure-native:cache/v20230401:listRedisKeys", args ?? new ListRedisKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve a Redis cache's access keys. This operation requires write permission to the cache resource.
        /// </summary>
        public static Output<ListRedisKeysResult> Invoke(ListRedisKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListRedisKeysResult>("azure-native:cache/v20230401:listRedisKeys", args ?? new ListRedisKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListRedisKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListRedisKeysArgs()
        {
        }
        public static new ListRedisKeysArgs Empty => new ListRedisKeysArgs();
    }

    public sealed class ListRedisKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListRedisKeysInvokeArgs()
        {
        }
        public static new ListRedisKeysInvokeArgs Empty => new ListRedisKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListRedisKeysResult
    {
        /// <summary>
        /// The current primary key that clients can use to authenticate with Redis cache.
        /// </summary>
        public readonly string PrimaryKey;
        /// <summary>
        /// The current secondary key that clients can use to authenticate with Redis cache.
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private ListRedisKeysResult(
            string primaryKey,

            string secondaryKey)
        {
            PrimaryKey = primaryKey;
            SecondaryKey = secondaryKey;
        }
    }
}
