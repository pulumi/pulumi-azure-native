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
    /// Definition of Address
    /// </summary>
    public sealed class AddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property addressDefinition
        /// </summary>
        [Input("addressDefinition")]
        public Input<string>? AddressDefinition { get; set; }

        public AddressArgs()
        {
        }
        public static new AddressArgs Empty => new AddressArgs();
    }
}
