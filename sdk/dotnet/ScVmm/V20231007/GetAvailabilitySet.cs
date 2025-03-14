// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ScVmm.V20231007
{
    public static class GetAvailabilitySet
    {
        /// <summary>
        /// Implements AvailabilitySet GET method.
        /// </summary>
        public static Task<GetAvailabilitySetResult> InvokeAsync(GetAvailabilitySetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilitySetResult>("azure-native:scvmm/v20231007:getAvailabilitySet", args ?? new GetAvailabilitySetArgs(), options.WithDefaults());

        /// <summary>
        /// Implements AvailabilitySet GET method.
        /// </summary>
        public static Output<GetAvailabilitySetResult> Invoke(GetAvailabilitySetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilitySetResult>("azure-native:scvmm/v20231007:getAvailabilitySet", args ?? new GetAvailabilitySetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Implements AvailabilitySet GET method.
        /// </summary>
        public static Output<GetAvailabilitySetResult> Invoke(GetAvailabilitySetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailabilitySetResult>("azure-native:scvmm/v20231007:getAvailabilitySet", args ?? new GetAvailabilitySetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAvailabilitySetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the AvailabilitySet.
        /// </summary>
        [Input("availabilitySetResourceName", required: true)]
        public string AvailabilitySetResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAvailabilitySetArgs()
        {
        }
        public static new GetAvailabilitySetArgs Empty => new GetAvailabilitySetArgs();
    }

    public sealed class GetAvailabilitySetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the AvailabilitySet.
        /// </summary>
        [Input("availabilitySetResourceName", required: true)]
        public Input<string> AvailabilitySetResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAvailabilitySetInvokeArgs()
        {
        }
        public static new GetAvailabilitySetInvokeArgs Empty => new GetAvailabilitySetInvokeArgs();
    }


    [OutputType]
    public sealed class GetAvailabilitySetResult
    {
        /// <summary>
        /// Name of the availability set.
        /// </summary>
        public readonly string? AvailabilitySetName;
        /// <summary>
        /// The extended location.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse ExtendedLocation;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// ARM Id of the vmmServer resource in which this resource resides.
        /// </summary>
        public readonly string? VmmServerId;

        [OutputConstructor]
        private GetAvailabilitySetResult(
            string? availabilitySetName,

            Outputs.ExtendedLocationResponse extendedLocation,

            string id,

            string location,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? vmmServerId)
        {
            AvailabilitySetName = availabilitySetName;
            ExtendedLocation = extendedLocation;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            VmmServerId = vmmServerId;
        }
    }
}
