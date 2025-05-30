// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// The configuration settings of a forward proxy used to make the requests.
    /// </summary>
    [OutputType]
    public sealed class ForwardProxyResponse
    {
        /// <summary>
        /// The convention used to determine the url of the request made.
        /// </summary>
        public readonly string? Convention;
        /// <summary>
        /// The name of the header containing the host of the request.
        /// </summary>
        public readonly string? CustomHostHeaderName;
        /// <summary>
        /// The name of the header containing the scheme of the request.
        /// </summary>
        public readonly string? CustomProtoHeaderName;

        [OutputConstructor]
        private ForwardProxyResponse(
            string? convention,

            string? customHostHeaderName,

            string? customProtoHeaderName)
        {
            Convention = convention;
            CustomHostHeaderName = customHostHeaderName;
            CustomProtoHeaderName = customProtoHeaderName;
        }
    }
}
