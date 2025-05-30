// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase
{
    /// <summary>
    /// The Test Base Package resource.
    /// 
    /// Uses Azure REST API version 2023-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-04-01-preview.
    /// 
    /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:testbase:Package")]
    public partial class Package : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application name
        /// </summary>
        [Output("applicationName")]
        public Output<string> ApplicationName { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The file path of the package.
        /// </summary>
        [Output("blobPath")]
        public Output<string?> BlobPath { get; private set; } = null!;

        /// <summary>
        /// The id of draft package. Used to create or update this package from a draft package.
        /// </summary>
        [Output("draftPackageId")]
        public Output<string?> DraftPackageId { get; private set; } = null!;

        /// <summary>
        /// The list of first party applications to test along with user application.
        /// </summary>
        [Output("firstPartyApps")]
        public Output<ImmutableArray<Outputs.FirstPartyAppDefinitionResponse>> FirstPartyApps { get; private set; } = null!;

        /// <summary>
        /// The flighting ring for feature update.
        /// </summary>
        [Output("flightingRing")]
        public Output<string?> FlightingRing { get; private set; } = null!;

        /// <summary>
        /// The list of gallery apps to test along with user application.
        /// </summary>
        [Output("galleryApps")]
        public Output<ImmutableArray<Outputs.GalleryAppDefinitionResponse>> GalleryApps { get; private set; } = null!;

        /// <summary>
        /// Specifies the baseline os and target os for inplace upgrade.
        /// </summary>
        [Output("inplaceUpgradeOSPair")]
        public Output<Outputs.InplaceUpgradeOSInfoResponse?> InplaceUpgradeOSPair { get; private set; } = null!;

        /// <summary>
        /// The metadata of Intune enrollment.
        /// </summary>
        [Output("intuneEnrollmentMetadata")]
        public Output<Outputs.IntuneEnrollmentMetadataResponse?> IntuneEnrollmentMetadata { get; private set; } = null!;

        /// <summary>
        /// Flag showing that whether the package is enabled. It doesn't schedule test for package which is not enabled.
        /// </summary>
        [Output("isEnabled")]
        public Output<bool> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// The UTC timestamp when the package was last modified.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The status of the package.
        /// </summary>
        [Output("packageStatus")]
        public Output<string> PackageStatus { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the target OSs of specific OS Update types.
        /// </summary>
        [Output("targetOSList")]
        public Output<ImmutableArray<Outputs.TargetOSInfoResponse>> TargetOSList { get; private set; } = null!;

        /// <summary>
        /// OOB, functional or flow driven. Mapped to the data in 'tests' property.
        /// </summary>
        [Output("testTypes")]
        public Output<ImmutableArray<string>> TestTypes { get; private set; } = null!;

        /// <summary>
        /// The detailed test information.
        /// </summary>
        [Output("tests")]
        public Output<ImmutableArray<Outputs.TestResponse>> Tests { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The validation results. There's validation on package when it's created or updated.
        /// </summary>
        [Output("validationResults")]
        public Output<ImmutableArray<Outputs.PackageValidationResultResponse>> ValidationResults { get; private set; } = null!;

        /// <summary>
        /// Application version
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Package resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Package(string name, PackageArgs args, CustomResourceOptions? options = null)
            : base("azure-native:testbase:Package", name, args ?? new PackageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Package(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:testbase:Package", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:testbase/v20201216preview:Package" },
                    new global::Pulumi.Alias { Type = "azure-native:testbase/v20220401preview:Package" },
                    new global::Pulumi.Alias { Type = "azure-native:testbase/v20231101preview:Package" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Package resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Package Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Package(name, id, options);
        }
    }

    public sealed class PackageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application name
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// The file path of the package.
        /// </summary>
        [Input("blobPath")]
        public Input<string>? BlobPath { get; set; }

        /// <summary>
        /// The id of draft package. Used to create or update this package from a draft package.
        /// </summary>
        [Input("draftPackageId")]
        public Input<string>? DraftPackageId { get; set; }

        [Input("firstPartyApps")]
        private InputList<Inputs.FirstPartyAppDefinitionArgs>? _firstPartyApps;

        /// <summary>
        /// The list of first party applications to test along with user application.
        /// </summary>
        public InputList<Inputs.FirstPartyAppDefinitionArgs> FirstPartyApps
        {
            get => _firstPartyApps ?? (_firstPartyApps = new InputList<Inputs.FirstPartyAppDefinitionArgs>());
            set => _firstPartyApps = value;
        }

        /// <summary>
        /// The flighting ring for feature update.
        /// </summary>
        [Input("flightingRing")]
        public Input<string>? FlightingRing { get; set; }

        /// <summary>
        /// Specifies the baseline os and target os for inplace upgrade.
        /// </summary>
        [Input("inplaceUpgradeOSPair")]
        public Input<Inputs.InplaceUpgradeOSInfoArgs>? InplaceUpgradeOSPair { get; set; }

        /// <summary>
        /// The metadata of Intune enrollment.
        /// </summary>
        [Input("intuneEnrollmentMetadata")]
        public Input<Inputs.IntuneEnrollmentMetadataArgs>? IntuneEnrollmentMetadata { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the Test Base Package.
        /// </summary>
        [Input("packageName")]
        public Input<string>? PackageName { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targetOSList")]
        private InputList<Inputs.TargetOSInfoArgs>? _targetOSList;

        /// <summary>
        /// Specifies the target OSs of specific OS Update types.
        /// </summary>
        public InputList<Inputs.TargetOSInfoArgs> TargetOSList
        {
            get => _targetOSList ?? (_targetOSList = new InputList<Inputs.TargetOSInfoArgs>());
            set => _targetOSList = value;
        }

        /// <summary>
        /// The resource name of the Test Base Account.
        /// </summary>
        [Input("testBaseAccountName", required: true)]
        public Input<string> TestBaseAccountName { get; set; } = null!;

        [Input("tests")]
        private InputList<Inputs.TestArgs>? _tests;

        /// <summary>
        /// The detailed test information.
        /// </summary>
        public InputList<Inputs.TestArgs> Tests
        {
            get => _tests ?? (_tests = new InputList<Inputs.TestArgs>());
            set => _tests = value;
        }

        /// <summary>
        /// Application version
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public PackageArgs()
        {
        }
        public static new PackageArgs Empty => new PackageArgs();
    }
}
