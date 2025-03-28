// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTFirmwareDefense.V20250401Preview.Outputs
{

    /// <summary>
    /// Workspace properties.
    /// </summary>
    [OutputType]
    public sealed class WorkspacePropertiesResponse
    {
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private WorkspacePropertiesResponse(string provisioningState)
        {
            ProvisioningState = provisioningState;
        }
    }
}
