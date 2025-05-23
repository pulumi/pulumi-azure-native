// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Inputs
{

    /// <summary>
    /// Environment variables of Spring Cloud Gateway
    /// </summary>
    public sealed class GatewayPropertiesEnvironmentVariablesArgs : global::Pulumi.ResourceArgs
    {
        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// Non-sensitive properties
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("secrets")]
        private InputMap<string>? _secrets;

        /// <summary>
        /// Sensitive properties
        /// </summary>
        public InputMap<string> Secrets
        {
            get => _secrets ?? (_secrets = new InputMap<string>());
            set => _secrets = value;
        }

        public GatewayPropertiesEnvironmentVariablesArgs()
        {
        }
        public static new GatewayPropertiesEnvironmentVariablesArgs Empty => new GatewayPropertiesEnvironmentVariablesArgs();
    }
}
