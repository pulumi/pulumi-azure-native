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
    /// Definition of InstanceBootModeValuesEnumValue
    /// </summary>
    [OutputType]
    public sealed class InstanceBootModeValuesEnumValueResponse
    {
        /// <summary>
        /// Property value
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private InstanceBootModeValuesEnumValueResponse(string? value)
        {
            Value = value;
        }
    }
}
