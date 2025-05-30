// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Purview
{
    public static class ListFeatureSubscription
    {
        /// <summary>
        /// Gets details from a list of feature names.
        /// 
        /// Uses Azure REST API version 2024-04-01-preview.
        /// 
        /// Other available API versions: 2021-12-01, 2023-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native purview [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListFeatureSubscriptionResult> InvokeAsync(ListFeatureSubscriptionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListFeatureSubscriptionResult>("azure-native:purview:listFeatureSubscription", args ?? new ListFeatureSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details from a list of feature names.
        /// 
        /// Uses Azure REST API version 2024-04-01-preview.
        /// 
        /// Other available API versions: 2021-12-01, 2023-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native purview [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListFeatureSubscriptionResult> Invoke(ListFeatureSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListFeatureSubscriptionResult>("azure-native:purview:listFeatureSubscription", args ?? new ListFeatureSubscriptionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details from a list of feature names.
        /// 
        /// Uses Azure REST API version 2024-04-01-preview.
        /// 
        /// Other available API versions: 2021-12-01, 2023-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native purview [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListFeatureSubscriptionResult> Invoke(ListFeatureSubscriptionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListFeatureSubscriptionResult>("azure-native:purview:listFeatureSubscription", args ?? new ListFeatureSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListFeatureSubscriptionArgs : global::Pulumi.InvokeArgs
    {
        [Input("features")]
        private List<string>? _features;

        /// <summary>
        /// Set of features
        /// </summary>
        public List<string> Features
        {
            get => _features ?? (_features = new List<string>());
            set => _features = value;
        }

        /// <summary>
        /// Location of feature.
        /// </summary>
        [Input("locations", required: true)]
        public string Locations { get; set; } = null!;

        public ListFeatureSubscriptionArgs()
        {
        }
        public static new ListFeatureSubscriptionArgs Empty => new ListFeatureSubscriptionArgs();
    }

    public sealed class ListFeatureSubscriptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("features")]
        private InputList<string>? _features;

        /// <summary>
        /// Set of features
        /// </summary>
        public InputList<string> Features
        {
            get => _features ?? (_features = new InputList<string>());
            set => _features = value;
        }

        /// <summary>
        /// Location of feature.
        /// </summary>
        [Input("locations", required: true)]
        public Input<string> Locations { get; set; } = null!;

        public ListFeatureSubscriptionInvokeArgs()
        {
        }
        public static new ListFeatureSubscriptionInvokeArgs Empty => new ListFeatureSubscriptionInvokeArgs();
    }


    [OutputType]
    public sealed class ListFeatureSubscriptionResult
    {
        /// <summary>
        /// Features with enabled status
        /// </summary>
        public readonly ImmutableDictionary<string, bool> Features;

        [OutputConstructor]
        private ListFeatureSubscriptionResult(ImmutableDictionary<string, bool> features)
        {
            Features = features;
        }
    }
}
