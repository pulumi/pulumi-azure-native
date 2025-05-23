// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetNspProfile
    {
        /// <summary>
        /// Gets the specified NSP profile.
        /// 
        /// Uses Azure REST API version 2023-08-01-preview.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2023-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetNspProfileResult> InvokeAsync(GetNspProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNspProfileResult>("azure-native:network:getNspProfile", args ?? new GetNspProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified NSP profile.
        /// 
        /// Uses Azure REST API version 2023-08-01-preview.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2023-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNspProfileResult> Invoke(GetNspProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNspProfileResult>("azure-native:network:getNspProfile", args ?? new GetNspProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified NSP profile.
        /// 
        /// Uses Azure REST API version 2023-08-01-preview.
        /// 
        /// Other available API versions: 2021-02-01-preview, 2023-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetNspProfileResult> Invoke(GetNspProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNspProfileResult>("azure-native:network:getNspProfile", args ?? new GetNspProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNspProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network security perimeter.
        /// </summary>
        [Input("networkSecurityPerimeterName", required: true)]
        public string NetworkSecurityPerimeterName { get; set; } = null!;

        /// <summary>
        /// The name of the NSP profile.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNspProfileArgs()
        {
        }
        public static new GetNspProfileArgs Empty => new GetNspProfileArgs();
    }

    public sealed class GetNspProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network security perimeter.
        /// </summary>
        [Input("networkSecurityPerimeterName", required: true)]
        public Input<string> NetworkSecurityPerimeterName { get; set; } = null!;

        /// <summary>
        /// The name of the NSP profile.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNspProfileInvokeArgs()
        {
        }
        public static new GetNspProfileInvokeArgs Empty => new GetNspProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetNspProfileResult
    {
        /// <summary>
        /// Version number that increases with every update to access rules within the profile.
        /// </summary>
        public readonly string AccessRulesVersion;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Version number that increases with every update to diagnostic settings within the profile.
        /// </summary>
        public readonly string DiagnosticSettingsVersion;
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
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNspProfileResult(
            string accessRulesVersion,

            string azureApiVersion,

            string diagnosticSettingsVersion,

            string id,

            string? location,

            string name,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AccessRulesVersion = accessRulesVersion;
            AzureApiVersion = azureApiVersion;
            DiagnosticSettingsVersion = diagnosticSettingsVersion;
            Id = id;
            Location = location;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
}
