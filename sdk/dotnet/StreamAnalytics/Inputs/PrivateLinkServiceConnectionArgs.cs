// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StreamAnalytics.Inputs
{

    /// <summary>
    /// A grouping of information about the connection to the remote resource.
    /// </summary>
    public sealed class PrivateLinkServiceConnectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// The ID(s) of the group(s) obtained from the remote resource that this private endpoint should connect to. Required on PUT (CreateOrUpdate) requests.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The resource id of the private link service. Required on PUT (CreateOrUpdate) requests.
        /// </summary>
        [Input("privateLinkServiceId")]
        public Input<string>? PrivateLinkServiceId { get; set; }

        public PrivateLinkServiceConnectionArgs()
        {
        }
        public static new PrivateLinkServiceConnectionArgs Empty => new PrivateLinkServiceConnectionArgs();
    }
}
