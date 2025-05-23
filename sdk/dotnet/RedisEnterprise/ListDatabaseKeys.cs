// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedisEnterprise
{
    public static class ListDatabaseKeys
    {
        /// <summary>
        /// Retrieves the access keys for the RedisEnterprise database.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2020-10-01-preview, 2021-02-01-preview, 2021-03-01, 2021-08-01, 2022-01-01, 2022-11-01-preview, 2023-03-01-preview, 2023-07-01, 2023-08-01-preview, 2023-10-01-preview, 2023-11-01, 2024-02-01, 2024-06-01-preview, 2024-09-01-preview, 2024-10-01, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redisenterprise [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListDatabaseKeysResult> InvokeAsync(ListDatabaseKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListDatabaseKeysResult>("azure-native:redisenterprise:listDatabaseKeys", args ?? new ListDatabaseKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the access keys for the RedisEnterprise database.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2020-10-01-preview, 2021-02-01-preview, 2021-03-01, 2021-08-01, 2022-01-01, 2022-11-01-preview, 2023-03-01-preview, 2023-07-01, 2023-08-01-preview, 2023-10-01-preview, 2023-11-01, 2024-02-01, 2024-06-01-preview, 2024-09-01-preview, 2024-10-01, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redisenterprise [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListDatabaseKeysResult> Invoke(ListDatabaseKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListDatabaseKeysResult>("azure-native:redisenterprise:listDatabaseKeys", args ?? new ListDatabaseKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the access keys for the RedisEnterprise database.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2020-10-01-preview, 2021-02-01-preview, 2021-03-01, 2021-08-01, 2022-01-01, 2022-11-01-preview, 2023-03-01-preview, 2023-07-01, 2023-08-01-preview, 2023-10-01-preview, 2023-11-01, 2024-02-01, 2024-06-01-preview, 2024-09-01-preview, 2024-10-01, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redisenterprise [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListDatabaseKeysResult> Invoke(ListDatabaseKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListDatabaseKeysResult>("azure-native:redisenterprise:listDatabaseKeys", args ?? new ListDatabaseKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListDatabaseKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis Enterprise cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the Redis Enterprise database.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListDatabaseKeysArgs()
        {
        }
        public static new ListDatabaseKeysArgs Empty => new ListDatabaseKeysArgs();
    }

    public sealed class ListDatabaseKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis Enterprise cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the Redis Enterprise database.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListDatabaseKeysInvokeArgs()
        {
        }
        public static new ListDatabaseKeysInvokeArgs Empty => new ListDatabaseKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListDatabaseKeysResult
    {
        /// <summary>
        /// The current primary key that clients can use to authenticate
        /// </summary>
        public readonly string PrimaryKey;
        /// <summary>
        /// The current secondary key that clients can use to authenticate
        /// </summary>
        public readonly string SecondaryKey;

        [OutputConstructor]
        private ListDatabaseKeysResult(
            string primaryKey,

            string secondaryKey)
        {
            PrimaryKey = primaryKey;
            SecondaryKey = secondaryKey;
        }
    }
}
