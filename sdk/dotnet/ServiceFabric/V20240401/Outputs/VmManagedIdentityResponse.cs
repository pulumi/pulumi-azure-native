// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.V20240401.Outputs
{

    /// <summary>
    /// Identities for the virtual machine scale set under the node type.
    /// </summary>
    [OutputType]
    public sealed class VmManagedIdentityResponse
    {
        /// <summary>
        /// The list of user identities associated with the virtual machine scale set under the node type. Each entry will be an ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public readonly ImmutableArray<string> UserAssignedIdentities;

        [OutputConstructor]
        private VmManagedIdentityResponse(ImmutableArray<string> userAssignedIdentities)
        {
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
