// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetNspAssociation
    {
        /// <summary>
        /// Gets the specified NSP association by name.
        /// 
        /// Uses Azure REST API version 2023-08-01-preview.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2023-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetNspAssociationResult> InvokeAsync(GetNspAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNspAssociationResult>("azure-native:network:getNspAssociation", args ?? new GetNspAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified NSP association by name.
        /// 
        /// Uses Azure REST API version 2023-08-01-preview.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2023-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNspAssociationResult> Invoke(GetNspAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNspAssociationResult>("azure-native:network:getNspAssociation", args ?? new GetNspAssociationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified NSP association by name.
        /// 
        /// Uses Azure REST API version 2023-08-01-preview.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2023-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNspAssociationResult> Invoke(GetNspAssociationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNspAssociationResult>("azure-native:network:getNspAssociation", args ?? new GetNspAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNspAssociationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NSP association.
        /// </summary>
        [Input("associationName", required: true)]
        public string AssociationName { get; set; } = null!;

        /// <summary>
        /// The name of the network security perimeter.
        /// </summary>
        [Input("networkSecurityPerimeterName", required: true)]
        public string NetworkSecurityPerimeterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNspAssociationArgs()
        {
        }
        public static new GetNspAssociationArgs Empty => new GetNspAssociationArgs();
    }

    public sealed class GetNspAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NSP association.
        /// </summary>
        [Input("associationName", required: true)]
        public Input<string> AssociationName { get; set; } = null!;

        /// <summary>
        /// The name of the network security perimeter.
        /// </summary>
        [Input("networkSecurityPerimeterName", required: true)]
        public Input<string> NetworkSecurityPerimeterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNspAssociationInvokeArgs()
        {
        }
        public static new GetNspAssociationInvokeArgs Empty => new GetNspAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetNspAssociationResult
    {
        /// <summary>
        /// Access mode on the association.
        /// </summary>
        public readonly string? AccessMode;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Specifies if there are provisioning issues
        /// </summary>
        public readonly string HasProvisioningIssues;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The PaaS resource to be associated.
        /// </summary>
        public readonly Outputs.SubResourceResponse? PrivateLinkResource;
        /// <summary>
        /// Profile id to which the PaaS resource is associated.
        /// </summary>
        public readonly Outputs.SubResourceResponse? Profile;
        /// <summary>
        /// The provisioning state of the resource  association resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNspAssociationResult(
            string? accessMode,

            string azureApiVersion,

            string hasProvisioningIssues,

            string id,

            string? location,

            string name,

            Outputs.SubResourceResponse? privateLinkResource,

            Outputs.SubResourceResponse? profile,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AccessMode = accessMode;
            AzureApiVersion = azureApiVersion;
            HasProvisioningIssues = hasProvisioningIssues;
            Id = id;
            Location = location;
            Name = name;
            PrivateLinkResource = privateLinkResource;
            Profile = profile;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
        }
    }
}
