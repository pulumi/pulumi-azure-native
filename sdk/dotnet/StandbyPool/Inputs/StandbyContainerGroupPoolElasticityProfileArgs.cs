// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StandbyPool.Inputs
{

    /// <summary>
    /// Specifies the elasticity profile of the standby container group pools.
    /// </summary>
    public sealed class StandbyContainerGroupPoolElasticityProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies maximum number of standby container groups in the standby pool.
        /// </summary>
        [Input("maxReadyCapacity", required: true)]
        public Input<double> MaxReadyCapacity { get; set; } = null!;

        /// <summary>
        /// Specifies refill policy of the pool.
        /// </summary>
        [Input("refillPolicy")]
        public InputUnion<string, Pulumi.AzureNative.StandbyPool.RefillPolicy>? RefillPolicy { get; set; }

        public StandbyContainerGroupPoolElasticityProfileArgs()
        {
        }
        public static new StandbyContainerGroupPoolElasticityProfileArgs Empty => new StandbyContainerGroupPoolElasticityProfileArgs();
    }
}
