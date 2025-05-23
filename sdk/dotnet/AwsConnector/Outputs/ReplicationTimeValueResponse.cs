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
    /// Definition of ReplicationTimeValue
    /// </summary>
    [OutputType]
    public sealed class ReplicationTimeValueResponse
    {
        /// <summary>
        /// Contains an integer specifying time in minutes.   Valid value: 15
        /// </summary>
        public readonly int? Minutes;

        [OutputConstructor]
        private ReplicationTimeValueResponse(int? minutes)
        {
            Minutes = minutes;
        }
    }
}
