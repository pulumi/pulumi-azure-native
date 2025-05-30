// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerInstance.Outputs
{

    /// <summary>
    /// The container group profile reference.
    /// </summary>
    [OutputType]
    public sealed class ContainerGroupProfileReferenceDefinitionResponse
    {
        /// <summary>
        /// The container group profile reference id.This will be an ARM resource id in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroupProfiles/{containerGroupProfileName}'.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The container group profile reference revision.
        /// </summary>
        public readonly int? Revision;

        [OutputConstructor]
        private ContainerGroupProfileReferenceDefinitionResponse(
            string? id,

            int? revision)
        {
            Id = id;
            Revision = revision;
        }
    }
}
