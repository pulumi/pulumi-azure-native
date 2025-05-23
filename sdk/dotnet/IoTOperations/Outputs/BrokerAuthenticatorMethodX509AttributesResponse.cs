// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// BrokerAuthenticatorMethodX509Attributes properties.
    /// </summary>
    [OutputType]
    public sealed class BrokerAuthenticatorMethodX509AttributesResponse
    {
        /// <summary>
        /// Attributes object.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Attributes;
        /// <summary>
        /// Subject of the X509 attribute.
        /// </summary>
        public readonly string Subject;

        [OutputConstructor]
        private BrokerAuthenticatorMethodX509AttributesResponse(
            ImmutableDictionary<string, string> attributes,

            string subject)
        {
            Attributes = attributes;
            Subject = subject;
        }
    }
}
