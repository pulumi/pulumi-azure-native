// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of ConnectionTrackingSpecificationResponse
    /// </summary>
    [OutputType]
    public sealed class ConnectionTrackingSpecificationResponseResponse
    {
        /// <summary>
        /// &lt;p&gt;Timeout (in seconds) for idle TCP connections in an established state. Min: 60 seconds. Max: 432000 seconds (5 days). Default: 432000 seconds. Recommended: Less than 432000 seconds.&lt;/p&gt;
        /// </summary>
        public readonly int? TcpEstablishedTimeout;
        /// <summary>
        /// &lt;p&gt;Timeout (in seconds) for idle UDP flows classified as streams which have seen more than one request-response transaction. Min: 60 seconds. Max: 180 seconds (3 minutes). Default: 180 seconds.&lt;/p&gt;
        /// </summary>
        public readonly int? UdpStreamTimeout;
        /// <summary>
        /// &lt;p&gt;Timeout (in seconds) for idle UDP flows that have seen traffic only in a single direction or a single request-response transaction. Min: 30 seconds. Max: 60 seconds. Default: 30 seconds.&lt;/p&gt;
        /// </summary>
        public readonly int? UdpTimeout;

        [OutputConstructor]
        private ConnectionTrackingSpecificationResponseResponse(
            int? tcpEstablishedTimeout,

            int? udpStreamTimeout,

            int? udpTimeout)
        {
            TcpEstablishedTimeout = tcpEstablishedTimeout;
            UdpStreamTimeout = udpStreamTimeout;
            UdpTimeout = udpTimeout;
        }
    }
}
