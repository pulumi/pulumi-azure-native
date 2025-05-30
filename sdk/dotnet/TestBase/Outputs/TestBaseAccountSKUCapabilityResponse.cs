// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase.Outputs
{

    /// <summary>
    /// Properties of the Test Base Account SKU Capability.
    /// </summary>
    [OutputType]
    public sealed class TestBaseAccountSKUCapabilityResponse
    {
        /// <summary>
        /// An invariant to describe the feature, such as 'SLA'.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An invariant if the feature is measured by quantity, such as 99.9%.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private TestBaseAccountSKUCapabilityResponse(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
