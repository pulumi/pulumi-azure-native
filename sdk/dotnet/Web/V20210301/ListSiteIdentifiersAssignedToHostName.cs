// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20210301
{
    public static class ListSiteIdentifiersAssignedToHostName
    {
        /// <summary>
        /// List all apps that are assigned to a hostname.
        /// </summary>
        public static Task<ListSiteIdentifiersAssignedToHostNameResult> InvokeAsync(ListSiteIdentifiersAssignedToHostNameArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSiteIdentifiersAssignedToHostNameResult>("azure-native:web/v20210301:listSiteIdentifiersAssignedToHostName", args ?? new ListSiteIdentifiersAssignedToHostNameArgs(), options.WithDefaults());

        /// <summary>
        /// List all apps that are assigned to a hostname.
        /// </summary>
        public static Output<ListSiteIdentifiersAssignedToHostNameResult> Invoke(ListSiteIdentifiersAssignedToHostNameInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSiteIdentifiersAssignedToHostNameResult>("azure-native:web/v20210301:listSiteIdentifiersAssignedToHostName", args ?? new ListSiteIdentifiersAssignedToHostNameInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List all apps that are assigned to a hostname.
        /// </summary>
        public static Output<ListSiteIdentifiersAssignedToHostNameResult> Invoke(ListSiteIdentifiersAssignedToHostNameInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListSiteIdentifiersAssignedToHostNameResult>("azure-native:web/v20210301:listSiteIdentifiersAssignedToHostName", args ?? new ListSiteIdentifiersAssignedToHostNameInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListSiteIdentifiersAssignedToHostNameArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the object.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public ListSiteIdentifiersAssignedToHostNameArgs()
        {
        }
        public static new ListSiteIdentifiersAssignedToHostNameArgs Empty => new ListSiteIdentifiersAssignedToHostNameArgs();
    }

    public sealed class ListSiteIdentifiersAssignedToHostNameInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the object.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ListSiteIdentifiersAssignedToHostNameInvokeArgs()
        {
        }
        public static new ListSiteIdentifiersAssignedToHostNameInvokeArgs Empty => new ListSiteIdentifiersAssignedToHostNameInvokeArgs();
    }


    [OutputType]
    public sealed class ListSiteIdentifiersAssignedToHostNameResult
    {
        /// <summary>
        /// Link to next page of resources.
        /// </summary>
        public readonly string NextLink;
        /// <summary>
        /// Collection of resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.IdentifierResponse> Value;

        [OutputConstructor]
        private ListSiteIdentifiersAssignedToHostNameResult(
            string nextLink,

            ImmutableArray<Outputs.IdentifierResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
