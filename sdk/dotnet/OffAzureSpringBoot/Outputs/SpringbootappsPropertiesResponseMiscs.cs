// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OffAzureSpringBoot.Outputs
{

    [OutputType]
    public sealed class SpringbootappsPropertiesResponseMiscs
    {
        /// <summary>
        /// The miscs. key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The miscs. value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private SpringbootappsPropertiesResponseMiscs(
            string key,

            string? value)
        {
            Key = key;
            Value = value;
        }
    }
}
