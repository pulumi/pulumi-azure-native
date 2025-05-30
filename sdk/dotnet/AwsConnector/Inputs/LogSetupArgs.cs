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
    /// Definition of LogSetup
    /// </summary>
    public sealed class LogSetupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;If a log type is enabled, that log type exports its control plane logs to CloudWatch Logs. If a log type isn't enabled, that log type doesn't export its control plane logs. Each individual log type can be enabled or disabled independently.&lt;/p&gt;
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("types")]
        private InputList<Union<string, Pulumi.AzureNative.AwsConnector.LogType>>? _types;

        /// <summary>
        /// &lt;p&gt;The available cluster control plane log types.&lt;/p&gt;
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.AwsConnector.LogType>> Types
        {
            get => _types ?? (_types = new InputList<Union<string, Pulumi.AzureNative.AwsConnector.LogType>>());
            set => _types = value;
        }

        public LogSetupArgs()
        {
        }
        public static new LogSetupArgs Empty => new LogSetupArgs();
    }
}
