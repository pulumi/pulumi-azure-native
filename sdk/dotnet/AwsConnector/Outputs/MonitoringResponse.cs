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
    /// Definition of Monitoring
    /// </summary>
    [OutputType]
    public sealed class MonitoringResponse
    {
        /// <summary>
        /// &lt;p&gt;Indicates whether detailed monitoring is enabled. Otherwise, basic monitoring is enabled.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.MonitoringStateEnumValueResponse? State;

        [OutputConstructor]
        private MonitoringResponse(Outputs.MonitoringStateEnumValueResponse? state)
        {
            State = state;
        }
    }
}
