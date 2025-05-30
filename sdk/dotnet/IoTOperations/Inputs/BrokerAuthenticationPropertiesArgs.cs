// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Inputs
{

    /// <summary>
    /// BrokerAuthentication Resource properties
    /// </summary>
    public sealed class BrokerAuthenticationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticationMethods", required: true)]
        private InputList<Inputs.BrokerAuthenticatorMethodsArgs>? _authenticationMethods;

        /// <summary>
        /// Defines a set of Broker authentication methods to be used on `BrokerListeners`. For each array element one authenticator type supported.
        /// </summary>
        public InputList<Inputs.BrokerAuthenticatorMethodsArgs> AuthenticationMethods
        {
            get => _authenticationMethods ?? (_authenticationMethods = new InputList<Inputs.BrokerAuthenticatorMethodsArgs>());
            set => _authenticationMethods = value;
        }

        public BrokerAuthenticationPropertiesArgs()
        {
        }
        public static new BrokerAuthenticationPropertiesArgs Empty => new BrokerAuthenticationPropertiesArgs();
    }
}
