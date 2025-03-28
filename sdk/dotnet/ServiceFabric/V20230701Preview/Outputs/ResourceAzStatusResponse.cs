// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.V20230701Preview.Outputs
{

    /// <summary>
    /// Describes Az Resiliency status of Base resources
    /// </summary>
    [OutputType]
    public sealed class ResourceAzStatusResponse
    {
        /// <summary>
        /// VM Size name.
        /// </summary>
        public readonly bool IsZoneResilient;
        /// <summary>
        /// VM Size properties.
        /// </summary>
        public readonly string ResourceName;
        /// <summary>
        /// VM Size id.
        /// </summary>
        public readonly string ResourceType;

        [OutputConstructor]
        private ResourceAzStatusResponse(
            bool isZoneResilient,

            string resourceName,

            string resourceType)
        {
            IsZoneResilient = isZoneResilient;
            ResourceName = resourceName;
            ResourceType = resourceType;
        }
    }
}
