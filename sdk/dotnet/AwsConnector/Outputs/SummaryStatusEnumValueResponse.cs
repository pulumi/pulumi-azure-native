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
    /// Definition of SummaryStatusEnumValue
    /// </summary>
    [OutputType]
    public sealed class SummaryStatusEnumValueResponse
    {
        /// <summary>
        /// Property value
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private SummaryStatusEnumValueResponse(string? value)
        {
            Value = value;
        }
    }
}
