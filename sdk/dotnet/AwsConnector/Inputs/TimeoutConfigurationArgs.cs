// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of TimeoutConfiguration
    /// </summary>
    public sealed class TimeoutConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of time in seconds a connection will stay active while idle. A value of ``0`` can be set to disable ``idleTimeout``. The ``idleTimeout`` default for ``HTTP``/``HTTP2``/``GRPC`` is 5 minutes. The ``idleTimeout`` default for ``TCP`` is 1 hour.
        /// </summary>
        [Input("idleTimeoutSeconds")]
        public Input<int>? IdleTimeoutSeconds { get; set; }

        /// <summary>
        /// The amount of time waiting for the upstream to respond with a complete response per request. A value of ``0`` can be set to disable ``perRequestTimeout``. ``perRequestTimeout`` can only be set if Service Connect ``appProtocol`` isn't ``TCP``. Only ``idleTimeout`` is allowed for ``TCP`` ``appProtocol``.
        /// </summary>
        [Input("perRequestTimeoutSeconds")]
        public Input<int>? PerRequestTimeoutSeconds { get; set; }

        public TimeoutConfigurationArgs()
        {
        }
        public static new TimeoutConfigurationArgs Empty => new TimeoutConfigurationArgs();
    }
}
