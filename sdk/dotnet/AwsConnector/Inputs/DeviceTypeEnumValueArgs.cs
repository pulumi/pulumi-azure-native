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
    /// Definition of DeviceTypeEnumValue
    /// </summary>
    public sealed class DeviceTypeEnumValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property value
        /// </summary>
        [Input("value")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.DeviceType>? Value { get; set; }

        public DeviceTypeEnumValueArgs()
        {
        }
        public static new DeviceTypeEnumValueArgs Empty => new DeviceTypeEnumValueArgs();
    }
}
