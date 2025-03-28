// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.V20220829
{
    public static class ListGlobalRulestackSecurityServices
    {
        /// <summary>
        /// List the security services for rulestack
        /// </summary>
        public static Task<ListGlobalRulestackSecurityServicesResult> InvokeAsync(ListGlobalRulestackSecurityServicesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListGlobalRulestackSecurityServicesResult>("azure-native:cloudngfw/v20220829:listGlobalRulestackSecurityServices", args ?? new ListGlobalRulestackSecurityServicesArgs(), options.WithDefaults());

        /// <summary>
        /// List the security services for rulestack
        /// </summary>
        public static Output<ListGlobalRulestackSecurityServicesResult> Invoke(ListGlobalRulestackSecurityServicesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListGlobalRulestackSecurityServicesResult>("azure-native:cloudngfw/v20220829:listGlobalRulestackSecurityServices", args ?? new ListGlobalRulestackSecurityServicesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List the security services for rulestack
        /// </summary>
        public static Output<ListGlobalRulestackSecurityServicesResult> Invoke(ListGlobalRulestackSecurityServicesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListGlobalRulestackSecurityServicesResult>("azure-native:cloudngfw/v20220829:listGlobalRulestackSecurityServices", args ?? new ListGlobalRulestackSecurityServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListGlobalRulestackSecurityServicesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// GlobalRulestack resource name
        /// </summary>
        [Input("globalRulestackName", required: true)]
        public string GlobalRulestackName { get; set; } = null!;

        [Input("skip")]
        public string? Skip { get; set; }

        [Input("top")]
        public int? Top { get; set; }

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public ListGlobalRulestackSecurityServicesArgs()
        {
        }
        public static new ListGlobalRulestackSecurityServicesArgs Empty => new ListGlobalRulestackSecurityServicesArgs();
    }

    public sealed class ListGlobalRulestackSecurityServicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// GlobalRulestack resource name
        /// </summary>
        [Input("globalRulestackName", required: true)]
        public Input<string> GlobalRulestackName { get; set; } = null!;

        [Input("skip")]
        public Input<string>? Skip { get; set; }

        [Input("top")]
        public Input<int>? Top { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ListGlobalRulestackSecurityServicesInvokeArgs()
        {
        }
        public static new ListGlobalRulestackSecurityServicesInvokeArgs Empty => new ListGlobalRulestackSecurityServicesInvokeArgs();
    }


    [OutputType]
    public sealed class ListGlobalRulestackSecurityServicesResult
    {
        /// <summary>
        /// next link
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// response value
        /// </summary>
        public readonly Outputs.SecurityServicesTypeListResponse Value;

        [OutputConstructor]
        private ListGlobalRulestackSecurityServicesResult(
            string? nextLink,

            Outputs.SecurityServicesTypeListResponse value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
