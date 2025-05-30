// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Automation.Outputs
{

    /// <summary>
    /// User identity used for CMK.
    /// </summary>
    [OutputType]
    public sealed class EncryptionPropertiesResponseIdentity
    {
        /// <summary>
        /// The user identity used for CMK. It will be an ARM resource id in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public readonly object? UserAssignedIdentity;

        [OutputConstructor]
        private EncryptionPropertiesResponseIdentity(object? userAssignedIdentity)
        {
            UserAssignedIdentity = userAssignedIdentity;
        }
    }
}
