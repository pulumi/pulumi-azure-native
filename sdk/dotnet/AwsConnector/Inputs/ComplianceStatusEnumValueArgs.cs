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
    /// Definition of ComplianceStatusEnumValue
    /// </summary>
    public sealed class ComplianceStatusEnumValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property value
        /// </summary>
        [Input("value")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.ComplianceStatus>? Value { get; set; }

        public ComplianceStatusEnumValueArgs()
        {
        }
        public static new ComplianceStatusEnumValueArgs Empty => new ComplianceStatusEnumValueArgs();
    }
}
