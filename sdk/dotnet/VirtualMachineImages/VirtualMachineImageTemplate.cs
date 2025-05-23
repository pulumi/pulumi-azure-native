// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VirtualMachineImages
{
    /// <summary>
    /// Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
    /// 
    /// Uses Azure REST API version 2024-02-01. In version 2.x of the Azure Native provider, it used API version 2022-07-01.
    /// 
    /// Other available API versions: 2022-07-01, 2023-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native virtualmachineimages [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:virtualmachineimages:VirtualMachineImageTemplate")]
    public partial class VirtualMachineImageTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether or not to automatically run the image template build on template creation or update.
        /// </summary>
        [Output("autoRun")]
        public Output<Outputs.ImageTemplateAutoRunResponse?> AutoRun { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Maximum duration to wait while building the image template (includes all customizations, optimization, validations, and distributions). Omit or specify 0 to use the default (4 hours).
        /// </summary>
        [Output("buildTimeoutInMinutes")]
        public Output<int?> BuildTimeoutInMinutes { get; private set; } = null!;

        /// <summary>
        /// Specifies the properties used to describe the customization steps of the image, like Image source etc
        /// </summary>
        [Output("customize")]
        public Output<ImmutableArray<object>> Customize { get; private set; } = null!;

        /// <summary>
        /// The distribution targets where the image output needs to go to.
        /// </summary>
        [Output("distribute")]
        public Output<ImmutableArray<object>> Distribute { get; private set; } = null!;

        /// <summary>
        /// Error handling options upon a build failure
        /// </summary>
        [Output("errorHandling")]
        public Output<Outputs.ImageTemplatePropertiesResponseErrorHandling?> ErrorHandling { get; private set; } = null!;

        /// <summary>
        /// The staging resource group id in the same subscription as the image template that will be used to build the image. This read-only field differs from 'stagingResourceGroup' only if the value specified in the 'stagingResourceGroup' field is empty.
        /// </summary>
        [Output("exactStagingResourceGroup")]
        public Output<string> ExactStagingResourceGroup { get; private set; } = null!;

        /// <summary>
        /// The identity of the image template, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ImageTemplateIdentityResponse> Identity { get; private set; } = null!;

        /// <summary>
        /// State of 'run' that is currently executing or was last executed.
        /// </summary>
        [Output("lastRunStatus")]
        public Output<Outputs.ImageTemplateLastRunStatusResponse> LastRunStatus { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Tags that will be applied to the resource group and/or resources created by the service.
        /// </summary>
        [Output("managedResourceTags")]
        public Output<ImmutableDictionary<string, string>?> ManagedResourceTags { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies optimization to be performed on image.
        /// </summary>
        [Output("optimize")]
        public Output<Outputs.ImageTemplatePropertiesResponseOptimize?> Optimize { get; private set; } = null!;

        /// <summary>
        /// Provisioning error, if any
        /// </summary>
        [Output("provisioningError")]
        public Output<Outputs.ProvisioningErrorResponse> ProvisioningError { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Specifies the properties used to describe the source image.
        /// </summary>
        [Output("source")]
        public Output<object> Source { get; private set; } = null!;

        /// <summary>
        /// The staging resource group id in the same subscription as the image template that will be used to build the image. If this field is empty, a resource group with a random name will be created. If the resource group specified in this field doesn't exist, it will be created with the same name. If the resource group specified exists, it must be empty and in the same region as the image template. The resource group created will be deleted during template deletion if this field is empty or the resource group specified doesn't exist, but if the resource group specified exists the resources created in the resource group will be deleted during template deletion and the resource group itself will remain.
        /// </summary>
        [Output("stagingResourceGroup")]
        public Output<string?> StagingResourceGroup { get; private set; } = null!;

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
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Configuration options and list of validations to be performed on the resulting image.
        /// </summary>
        [Output("validate")]
        public Output<Outputs.ImageTemplatePropertiesResponseValidate?> Validate { get; private set; } = null!;

        /// <summary>
        /// Describes how virtual machine is set up to build images
        /// </summary>
        [Output("vmProfile")]
        public Output<Outputs.ImageTemplateVmProfileResponse?> VmProfile { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachineImageTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachineImageTemplate(string name, VirtualMachineImageTemplateArgs args, CustomResourceOptions? options = null)
            : base("azure-native:virtualmachineimages:VirtualMachineImageTemplate", name, args ?? new VirtualMachineImageTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachineImageTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:virtualmachineimages:VirtualMachineImageTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20180201preview:VirtualMachineImageTemplate" },
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20190201preview:VirtualMachineImageTemplate" },
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20190501preview:VirtualMachineImageTemplate" },
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20200214:VirtualMachineImageTemplate" },
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20211001:VirtualMachineImageTemplate" },
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20220214:VirtualMachineImageTemplate" },
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20220701:VirtualMachineImageTemplate" },
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20230701:VirtualMachineImageTemplate" },
                    new global::Pulumi.Alias { Type = "azure-native:virtualmachineimages/v20240201:VirtualMachineImageTemplate" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualMachineImageTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachineImageTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualMachineImageTemplate(name, id, options);
        }
    }

    public sealed class VirtualMachineImageTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether or not to automatically run the image template build on template creation or update.
        /// </summary>
        [Input("autoRun")]
        public Input<Inputs.ImageTemplateAutoRunArgs>? AutoRun { get; set; }

        /// <summary>
        /// Maximum duration to wait while building the image template (includes all customizations, optimization, validations, and distributions). Omit or specify 0 to use the default (4 hours).
        /// </summary>
        [Input("buildTimeoutInMinutes")]
        public Input<int>? BuildTimeoutInMinutes { get; set; }

        [Input("customize")]
        private InputList<object>? _customize;

        /// <summary>
        /// Specifies the properties used to describe the customization steps of the image, like Image source etc
        /// </summary>
        public InputList<object> Customize
        {
            get => _customize ?? (_customize = new InputList<object>());
            set => _customize = value;
        }

        [Input("distribute", required: true)]
        private InputList<object>? _distribute;

        /// <summary>
        /// The distribution targets where the image output needs to go to.
        /// </summary>
        public InputList<object> Distribute
        {
            get => _distribute ?? (_distribute = new InputList<object>());
            set => _distribute = value;
        }

        /// <summary>
        /// Error handling options upon a build failure
        /// </summary>
        [Input("errorHandling")]
        public Input<Inputs.ImageTemplatePropertiesErrorHandlingArgs>? ErrorHandling { get; set; }

        /// <summary>
        /// The identity of the image template, if configured.
        /// </summary>
        [Input("identity", required: true)]
        public Input<Inputs.ImageTemplateIdentityArgs> Identity { get; set; } = null!;

        /// <summary>
        /// The name of the image Template
        /// </summary>
        [Input("imageTemplateName")]
        public Input<string>? ImageTemplateName { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("managedResourceTags")]
        private InputMap<string>? _managedResourceTags;

        /// <summary>
        /// Tags that will be applied to the resource group and/or resources created by the service.
        /// </summary>
        public InputMap<string> ManagedResourceTags
        {
            get => _managedResourceTags ?? (_managedResourceTags = new InputMap<string>());
            set => _managedResourceTags = value;
        }

        /// <summary>
        /// Specifies optimization to be performed on image.
        /// </summary>
        [Input("optimize")]
        public Input<Inputs.ImageTemplatePropertiesOptimizeArgs>? Optimize { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the properties used to describe the source image.
        /// </summary>
        [Input("source", required: true)]
        public object Source { get; set; } = null!;

        /// <summary>
        /// The staging resource group id in the same subscription as the image template that will be used to build the image. If this field is empty, a resource group with a random name will be created. If the resource group specified in this field doesn't exist, it will be created with the same name. If the resource group specified exists, it must be empty and in the same region as the image template. The resource group created will be deleted during template deletion if this field is empty or the resource group specified doesn't exist, but if the resource group specified exists the resources created in the resource group will be deleted during template deletion and the resource group itself will remain.
        /// </summary>
        [Input("stagingResourceGroup")]
        public Input<string>? StagingResourceGroup { get; set; }

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

        /// <summary>
        /// Configuration options and list of validations to be performed on the resulting image.
        /// </summary>
        [Input("validate")]
        public Input<Inputs.ImageTemplatePropertiesValidateArgs>? Validate { get; set; }

        /// <summary>
        /// Describes how virtual machine is set up to build images
        /// </summary>
        [Input("vmProfile")]
        public Input<Inputs.ImageTemplateVmProfileArgs>? VmProfile { get; set; }

        public VirtualMachineImageTemplateArgs()
        {
            BuildTimeoutInMinutes = 0;
        }
        public static new VirtualMachineImageTemplateArgs Empty => new VirtualMachineImageTemplateArgs();
    }
}
