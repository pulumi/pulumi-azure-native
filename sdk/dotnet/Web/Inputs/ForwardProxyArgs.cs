// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// The configuration settings of a forward proxy used to make the requests.
    /// </summary>
    public sealed class ForwardProxyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The convention used to determine the url of the request made.
        /// </summary>
        [Input("convention")]
        public Input<Pulumi.AzureNative.Web.ForwardProxyConvention>? Convention { get; set; }

        /// <summary>
        /// The name of the header containing the host of the request.
        /// </summary>
        [Input("customHostHeaderName")]
        public Input<string>? CustomHostHeaderName { get; set; }

        /// <summary>
        /// The name of the header containing the scheme of the request.
        /// </summary>
        [Input("customProtoHeaderName")]
        public Input<string>? CustomProtoHeaderName { get; set; }

        public ForwardProxyArgs()
        {
        }
        public static new ForwardProxyArgs Empty => new ForwardProxyArgs();
    }
}
