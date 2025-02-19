// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights.V20230311.Outputs
{

    /// <summary>
    /// A setting used to control an agent behavior on a host machine
    /// </summary>
    [OutputType]
    public sealed class AgentSettingResponse
    {
        /// <summary>
        /// The name of the setting. 
        /// Must be part of the list of supported settings
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The value of the setting
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private AgentSettingResponse(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
