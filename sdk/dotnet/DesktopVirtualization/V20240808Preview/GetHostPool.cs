// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization.V20240808Preview
{
    public static class GetHostPool
    {
        /// <summary>
        /// Get a host pool.
        /// </summary>
        public static Task<GetHostPoolResult> InvokeAsync(GetHostPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHostPoolResult>("azure-native:desktopvirtualization/v20240808preview:getHostPool", args ?? new GetHostPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Get a host pool.
        /// </summary>
        public static Output<GetHostPoolResult> Invoke(GetHostPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostPoolResult>("azure-native:desktopvirtualization/v20240808preview:getHostPool", args ?? new GetHostPoolInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a host pool.
        /// </summary>
        public static Output<GetHostPoolResult> Invoke(GetHostPoolInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostPoolResult>("azure-native:desktopvirtualization/v20240808preview:getHostPool", args ?? new GetHostPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHostPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the host pool within the specified resource group
        /// </summary>
        [Input("hostPoolName", required: true)]
        public string HostPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetHostPoolArgs()
        {
        }
        public static new GetHostPoolArgs Empty => new GetHostPoolArgs();
    }

    public sealed class GetHostPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the host pool within the specified resource group
        /// </summary>
        [Input("hostPoolName", required: true)]
        public Input<string> HostPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetHostPoolInvokeArgs()
        {
        }
        public static new GetHostPoolInvokeArgs Empty => new GetHostPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetHostPoolResult
    {
        /// <summary>
        /// The session host configuration for updating agent, monitoring agent, and stack component.
        /// </summary>
        public readonly Outputs.AgentUpdatePropertiesResponse? AgentUpdate;
        /// <summary>
        /// List of App Attach Package links.
        /// </summary>
        public readonly ImmutableArray<string> AppAttachPackageReferences;
        /// <summary>
        /// List of applicationGroup links.
        /// </summary>
        public readonly ImmutableArray<string> ApplicationGroupReferences;
        /// <summary>
        /// Is cloud pc resource.
        /// </summary>
        public readonly bool CloudPcResource;
        /// <summary>
        /// Custom rdp property of HostPool.
        /// </summary>
        public readonly string? CustomRdpProperty;
        /// <summary>
        /// Description of HostPool.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Default: AVD-wide settings are used to determine connection availability, Enabled: UDP will attempt this connection type when making connections. This means that this connection is possible, but is not guaranteed, as there are other factors that may prevent this connection type, Disabled: UDP will not attempt this connection type when making connections
        /// </summary>
        public readonly string? DirectUDP;
        /// <summary>
        /// If etag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Friendly name of HostPool.
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// HostPool type for desktop.
        /// </summary>
        public readonly string HostPoolType;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed service identities assigned to this resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The type of the load balancer.
        /// </summary>
        public readonly string LoadBalancerType;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
        /// </summary>
        public readonly string? ManagedBy;
        /// <summary>
        /// Default: AVD-wide settings are used to determine connection availability, Enabled: UDP will attempt this connection type when making connections. This means that this connection is possible, but is not guaranteed, as there are other factors that may prevent this connection type, Disabled: UDP will not attempt this connection type when making connections
        /// </summary>
        public readonly string? ManagedPrivateUDP;
        /// <summary>
        /// The type of management for this hostpool, Automated or Standard. The default value is Automated.
        /// </summary>
        public readonly string? ManagementType;
        /// <summary>
        /// The max session limit of HostPool.
        /// </summary>
        public readonly int? MaxSessionLimit;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ObjectId of HostPool. (internal use)
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// PersonalDesktopAssignment type for HostPool.
        /// </summary>
        public readonly string? PersonalDesktopAssignmentType;
        /// <summary>
        /// Details of the resource plan.
        /// </summary>
        public readonly Outputs.PlanResponse? Plan;
        /// <summary>
        /// The type of preferred application group type, default to Desktop Application Group
        /// </summary>
        public readonly string PreferredAppGroupType;
        /// <summary>
        /// List of private endpoint connection associated with the specified resource
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// Default: AVD-wide settings are used to determine connection availability, Enabled: UDP will attempt this connection type when making connections. This means that this connection is possible, but is not guaranteed, as there are other factors that may prevent this connection type, Disabled: UDP will not attempt this connection type when making connections
        /// </summary>
        public readonly string? PublicUDP;
        /// <summary>
        /// The registration info of HostPool.
        /// </summary>
        public readonly Outputs.RegistrationInfoResponse? RegistrationInfo;
        /// <summary>
        /// Default: AVD-wide settings are used to determine connection availability, Enabled: UDP will attempt this connection type when making connections. This means that this connection is possible, but is not guaranteed, as there are other factors that may prevent this connection type, Disabled: UDP will not attempt this connection type when making connections
        /// </summary>
        public readonly string? RelayUDP;
        /// <summary>
        /// The ring number of HostPool.
        /// </summary>
        public readonly int? Ring;
        /// <summary>
        /// The SKU (Stock Keeping Unit) assigned to this resource.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// ClientId for the registered Relying Party used to issue WVD SSO certificates.
        /// </summary>
        public readonly string? SsoClientId;
        /// <summary>
        /// Path to Azure KeyVault storing the secret used for communication to ADFS.
        /// </summary>
        public readonly string? SsoClientSecretKeyVaultPath;
        /// <summary>
        /// The type of single sign on Secret Type.
        /// </summary>
        public readonly string? SsoSecretType;
        /// <summary>
        /// URL to customer ADFS server for signing WVD SSO certificates.
        /// </summary>
        public readonly string? SsoadfsAuthority;
        /// <summary>
        /// The flag to turn on/off StartVMOnConnect feature.
        /// </summary>
        public readonly bool? StartVMOnConnect;
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
        /// Is validation environment.
        /// </summary>
        public readonly bool? ValidationEnvironment;
        /// <summary>
        /// VM template for sessionhosts configuration within hostpool.
        /// </summary>
        public readonly string? VmTemplate;

        [OutputConstructor]
        private GetHostPoolResult(
            Outputs.AgentUpdatePropertiesResponse? agentUpdate,

            ImmutableArray<string> appAttachPackageReferences,

            ImmutableArray<string> applicationGroupReferences,

            bool cloudPcResource,

            string? customRdpProperty,

            string? description,

            string? directUDP,

            string etag,

            string? friendlyName,

            string hostPoolType,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string? kind,

            string loadBalancerType,

            string location,

            string? managedBy,

            string? managedPrivateUDP,

            string? managementType,

            int? maxSessionLimit,

            string name,

            string objectId,

            string? personalDesktopAssignmentType,

            Outputs.PlanResponse? plan,

            string preferredAppGroupType,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string? publicNetworkAccess,

            string? publicUDP,

            Outputs.RegistrationInfoResponse? registrationInfo,

            string? relayUDP,

            int? ring,

            Outputs.SkuResponse? sku,

            string? ssoClientId,

            string? ssoClientSecretKeyVaultPath,

            string? ssoSecretType,

            string? ssoadfsAuthority,

            bool? startVMOnConnect,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            bool? validationEnvironment,

            string? vmTemplate)
        {
            AgentUpdate = agentUpdate;
            AppAttachPackageReferences = appAttachPackageReferences;
            ApplicationGroupReferences = applicationGroupReferences;
            CloudPcResource = cloudPcResource;
            CustomRdpProperty = customRdpProperty;
            Description = description;
            DirectUDP = directUDP;
            Etag = etag;
            FriendlyName = friendlyName;
            HostPoolType = hostPoolType;
            Id = id;
            Identity = identity;
            Kind = kind;
            LoadBalancerType = loadBalancerType;
            Location = location;
            ManagedBy = managedBy;
            ManagedPrivateUDP = managedPrivateUDP;
            ManagementType = managementType;
            MaxSessionLimit = maxSessionLimit;
            Name = name;
            ObjectId = objectId;
            PersonalDesktopAssignmentType = personalDesktopAssignmentType;
            Plan = plan;
            PreferredAppGroupType = preferredAppGroupType;
            PrivateEndpointConnections = privateEndpointConnections;
            PublicNetworkAccess = publicNetworkAccess;
            PublicUDP = publicUDP;
            RegistrationInfo = registrationInfo;
            RelayUDP = relayUDP;
            Ring = ring;
            Sku = sku;
            SsoClientId = ssoClientId;
            SsoClientSecretKeyVaultPath = ssoClientSecretKeyVaultPath;
            SsoSecretType = ssoSecretType;
            SsoadfsAuthority = ssoadfsAuthority;
            StartVMOnConnect = startVMOnConnect;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            ValidationEnvironment = validationEnvironment;
            VmTemplate = vmTemplate;
        }
    }
}
