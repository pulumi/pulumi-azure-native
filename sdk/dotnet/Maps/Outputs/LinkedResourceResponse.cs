// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maps.Outputs
{

    /// <summary>
    /// Linked resource is reference to a resource deployed in an Azure subscription, add the linked resource `uniqueName` value as an optional parameter for operations on Azure Maps Geospatial REST APIs.
    /// </summary>
    [OutputType]
    public sealed class LinkedResourceResponse
    {
        /// <summary>
        /// ARM resource id in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/accounts/{storageName}'.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A provided name which uniquely identifies the linked resource.
        /// </summary>
        public readonly string UniqueName;

        [OutputConstructor]
        private LinkedResourceResponse(
            string id,

            string uniqueName)
        {
            Id = id;
            UniqueName = uniqueName;
        }
    }
}
