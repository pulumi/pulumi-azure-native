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
    /// Definition of PlatformValuesEnumValue
    /// </summary>
    public sealed class PlatformValuesEnumValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property value
        /// </summary>
        [Input("value")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.PlatformValues>? Value { get; set; }

        public PlatformValuesEnumValueArgs()
        {
        }
        public static new PlatformValuesEnumValueArgs Empty => new PlatformValuesEnumValueArgs();
    }
}
