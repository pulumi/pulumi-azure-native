// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of OptionSetting
    /// </summary>
    [OutputType]
    public sealed class OptionSettingResponse
    {
        /// <summary>
        /// A unique namespace that identifies the option's associated AWS resource.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// The name of the configuration option.
        /// </summary>
        public readonly string? OptionName;
        /// <summary>
        /// A unique resource name for the option setting. Use it for a time–based scaling configuration option.
        /// </summary>
        public readonly string? ResourceName;
        /// <summary>
        /// The current value for the configuration option.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private OptionSettingResponse(
            string? @namespace,

            string? optionName,

            string? resourceName,

            string? value)
        {
            Namespace = @namespace;
            OptionName = optionName;
            ResourceName = resourceName;
            Value = value;
        }
    }
}
