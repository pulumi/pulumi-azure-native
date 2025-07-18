// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw
{
    public static class ListGlobalRulestackAppIds
    {
        /// <summary>
        /// List of AppIds for GlobalRulestack ApiVersion
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListGlobalRulestackAppIdsResult> InvokeAsync(ListGlobalRulestackAppIdsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListGlobalRulestackAppIdsResult>("azure-native:cloudngfw:listGlobalRulestackAppIds", args ?? new ListGlobalRulestackAppIdsArgs(), options.WithDefaults());

        /// <summary>
        /// List of AppIds for GlobalRulestack ApiVersion
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListGlobalRulestackAppIdsResult> Invoke(ListGlobalRulestackAppIdsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListGlobalRulestackAppIdsResult>("azure-native:cloudngfw:listGlobalRulestackAppIds", args ?? new ListGlobalRulestackAppIdsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List of AppIds for GlobalRulestack ApiVersion
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListGlobalRulestackAppIdsResult> Invoke(ListGlobalRulestackAppIdsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListGlobalRulestackAppIdsResult>("azure-native:cloudngfw:listGlobalRulestackAppIds", args ?? new ListGlobalRulestackAppIdsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListGlobalRulestackAppIdsArgs : global::Pulumi.InvokeArgs
    {
        [Input("appIdVersion")]
        public string? AppIdVersion { get; set; }

        [Input("appPrefix")]
        public string? AppPrefix { get; set; }

        /// <summary>
        /// GlobalRulestack resource name
        /// </summary>
        [Input("globalRulestackName", required: true)]
        public string GlobalRulestackName { get; set; } = null!;

        [Input("skip")]
        public string? Skip { get; set; }

        [Input("top")]
        public int? Top { get; set; }

        public ListGlobalRulestackAppIdsArgs()
        {
        }
        public static new ListGlobalRulestackAppIdsArgs Empty => new ListGlobalRulestackAppIdsArgs();
    }

    public sealed class ListGlobalRulestackAppIdsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("appIdVersion")]
        public Input<string>? AppIdVersion { get; set; }

        [Input("appPrefix")]
        public Input<string>? AppPrefix { get; set; }

        /// <summary>
        /// GlobalRulestack resource name
        /// </summary>
        [Input("globalRulestackName", required: true)]
        public Input<string> GlobalRulestackName { get; set; } = null!;

        [Input("skip")]
        public Input<string>? Skip { get; set; }

        [Input("top")]
        public Input<int>? Top { get; set; }

        public ListGlobalRulestackAppIdsInvokeArgs()
        {
        }
        public static new ListGlobalRulestackAppIdsInvokeArgs Empty => new ListGlobalRulestackAppIdsInvokeArgs();
    }


    [OutputType]
    public sealed class ListGlobalRulestackAppIdsResult
    {
        /// <summary>
        /// next Link
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// List of AppIds
        /// </summary>
        public readonly ImmutableArray<string> Value;

        [OutputConstructor]
        private ListGlobalRulestackAppIdsResult(
            string? nextLink,

            ImmutableArray<string> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
