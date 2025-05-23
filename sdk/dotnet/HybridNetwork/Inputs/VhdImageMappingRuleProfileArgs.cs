// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// Vhd mapping rule profile
    /// </summary>
    public sealed class VhdImageMappingRuleProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of values.
        /// </summary>
        [Input("userConfiguration")]
        public Input<string>? UserConfiguration { get; set; }

        public VhdImageMappingRuleProfileArgs()
        {
        }
        public static new VhdImageMappingRuleProfileArgs Empty => new VhdImageMappingRuleProfileArgs();
    }
}
