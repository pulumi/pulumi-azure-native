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
    /// Definition of DomainTypeEnumValue
    /// </summary>
    public sealed class DomainTypeEnumValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property value
        /// </summary>
        [Input("value")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.DomainType>? Value { get; set; }

        public DomainTypeEnumValueArgs()
        {
        }
        public static new DomainTypeEnumValueArgs Empty => new DomainTypeEnumValueArgs();
    }
}
