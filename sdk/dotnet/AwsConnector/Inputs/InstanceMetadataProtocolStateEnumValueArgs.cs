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
    /// Definition of InstanceMetadataProtocolStateEnumValue
    /// </summary>
    public sealed class InstanceMetadataProtocolStateEnumValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property value
        /// </summary>
        [Input("value")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.InstanceMetadataProtocolState>? Value { get; set; }

        public InstanceMetadataProtocolStateEnumValueArgs()
        {
            Value = "disabled";
        }
        public static new InstanceMetadataProtocolStateEnumValueArgs Empty => new InstanceMetadataProtocolStateEnumValueArgs();
    }
}
