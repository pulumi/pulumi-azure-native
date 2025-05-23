// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Marketplace
{
    public static class ListPrivateStoreSubscriptionsContext
    {
        /// <summary>
        /// List all the subscriptions in the private store context
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// 
        /// Other available API versions: 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native marketplace [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListPrivateStoreSubscriptionsContextResult> InvokeAsync(ListPrivateStoreSubscriptionsContextArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListPrivateStoreSubscriptionsContextResult>("azure-native:marketplace:listPrivateStoreSubscriptionsContext", args ?? new ListPrivateStoreSubscriptionsContextArgs(), options.WithDefaults());

        /// <summary>
        /// List all the subscriptions in the private store context
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// 
        /// Other available API versions: 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native marketplace [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListPrivateStoreSubscriptionsContextResult> Invoke(ListPrivateStoreSubscriptionsContextInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListPrivateStoreSubscriptionsContextResult>("azure-native:marketplace:listPrivateStoreSubscriptionsContext", args ?? new ListPrivateStoreSubscriptionsContextInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List all the subscriptions in the private store context
        /// 
        /// Uses Azure REST API version 2023-01-01.
        /// 
        /// Other available API versions: 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native marketplace [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListPrivateStoreSubscriptionsContextResult> Invoke(ListPrivateStoreSubscriptionsContextInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListPrivateStoreSubscriptionsContextResult>("azure-native:marketplace:listPrivateStoreSubscriptionsContext", args ?? new ListPrivateStoreSubscriptionsContextInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListPrivateStoreSubscriptionsContextArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The store ID - must use the tenant ID
        /// </summary>
        [Input("privateStoreId", required: true)]
        public string PrivateStoreId { get; set; } = null!;

        public ListPrivateStoreSubscriptionsContextArgs()
        {
        }
        public static new ListPrivateStoreSubscriptionsContextArgs Empty => new ListPrivateStoreSubscriptionsContextArgs();
    }

    public sealed class ListPrivateStoreSubscriptionsContextInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The store ID - must use the tenant ID
        /// </summary>
        [Input("privateStoreId", required: true)]
        public Input<string> PrivateStoreId { get; set; } = null!;

        public ListPrivateStoreSubscriptionsContextInvokeArgs()
        {
        }
        public static new ListPrivateStoreSubscriptionsContextInvokeArgs Empty => new ListPrivateStoreSubscriptionsContextInvokeArgs();
    }


    [OutputType]
    public sealed class ListPrivateStoreSubscriptionsContextResult
    {
        public readonly ImmutableArray<string> SubscriptionsIds;

        [OutputConstructor]
        private ListPrivateStoreSubscriptionsContextResult(ImmutableArray<string> subscriptionsIds)
        {
            SubscriptionsIds = subscriptionsIds;
        }
    }
}
