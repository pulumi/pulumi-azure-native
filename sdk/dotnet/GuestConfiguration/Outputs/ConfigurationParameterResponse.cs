// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.GuestConfiguration.Outputs
{

    /// <summary>
    /// Represents a configuration parameter.
    /// </summary>
    [OutputType]
    public sealed class ConfigurationParameterResponse
    {
        /// <summary>
        /// Name of the configuration parameter.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Value of the configuration parameter.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ConfigurationParameterResponse(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
