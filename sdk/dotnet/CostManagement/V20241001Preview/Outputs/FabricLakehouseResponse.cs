// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.V20241001Preview.Outputs
{

    /// <summary>
    /// This represents the fabric lakehouse identifier.
    /// </summary>
    [OutputType]
    public sealed class FabricLakehouseResponse
    {
        /// <summary>
        /// The display name of the fabric lakehouse.  This is only returned in Get operations.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The identifier of the fabric lakehouse.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private FabricLakehouseResponse(
            string? displayName,

            string? id)
        {
            DisplayName = displayName;
            Id = id;
        }
    }
}
