// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Inputs
{

    /// <summary>
    /// A user that can be assigned to a compute instance.
    /// </summary>
    public sealed class AssignedUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User’s AAD Object Id.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        /// <summary>
        /// User’s AAD Tenant Id.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public AssignedUserArgs()
        {
        }
        public static new AssignedUserArgs Empty => new AssignedUserArgs();
    }
}
