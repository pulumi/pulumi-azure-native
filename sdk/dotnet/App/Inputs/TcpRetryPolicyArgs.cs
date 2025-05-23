// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Inputs
{

    /// <summary>
    /// Policy that defines tcp request retry conditions
    /// </summary>
    public sealed class TcpRetryPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of attempts to connect to the tcp service
        /// </summary>
        [Input("maxConnectAttempts")]
        public Input<int>? MaxConnectAttempts { get; set; }

        public TcpRetryPolicyArgs()
        {
        }
        public static new TcpRetryPolicyArgs Empty => new TcpRetryPolicyArgs();
    }
}
