// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.Inputs
{

    /// <summary>
    /// Authentication mechanism for IoT devices.
    /// </summary>
    public sealed class AuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Symmetric key for authentication.
        /// </summary>
        [Input("symmetricKey")]
        public Input<Inputs.SymmetricKeyArgs>? SymmetricKey { get; set; }

        public AuthenticationArgs()
        {
        }
        public static new AuthenticationArgs Empty => new AuthenticationArgs();
    }
}
