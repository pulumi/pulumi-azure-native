// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApplicationGroupType = {
    RemoteApp: "RemoteApp",
    Desktop: "Desktop",
} as const;

/**
 * Resource Type of ApplicationGroup.
 */
export type ApplicationGroupType = (typeof ApplicationGroupType)[keyof typeof ApplicationGroupType];

export const CommandLineSetting = {
    DoNotAllow: "DoNotAllow",
    Allow: "Allow",
    Require: "Require",
} as const;

/**
 * Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
 */
export type CommandLineSetting = (typeof CommandLineSetting)[keyof typeof CommandLineSetting];

export const DayOfWeek = {
    Sunday: "Sunday",
    Monday: "Monday",
    Tuesday: "Tuesday",
    Wednesday: "Wednesday",
    Thursday: "Thursday",
    Friday: "Friday",
    Saturday: "Saturday",
} as const;

/**
 * Day of the week.
 */
export type DayOfWeek = (typeof DayOfWeek)[keyof typeof DayOfWeek];

export const DirectUDP = {
    Default: "Default",
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Default: AVD-wide settings are used to determine connection availability, Enabled: UDP will attempt this connection type when making connections. This means that this connection is possible, but is not guaranteed, as there are other factors that may prevent this connection type, Disabled: UDP will not attempt this connection type when making connections
 */
export type DirectUDP = (typeof DirectUDP)[keyof typeof DirectUDP];

export const FailHealthCheckOnStagingFailure = {
    Unhealthy: "Unhealthy",
    NeedsAssistance: "NeedsAssistance",
    DoNotFail: "DoNotFail",
} as const;

/**
 * Parameter indicating how the health check should behave if this package fails staging
 */
export type FailHealthCheckOnStagingFailure = (typeof FailHealthCheckOnStagingFailure)[keyof typeof FailHealthCheckOnStagingFailure];

export const HostPoolType = {
    /**
     * Users will be assigned a SessionHost either by administrators (PersonalDesktopAssignmentType = Direct) or upon connecting to the pool (PersonalDesktopAssignmentType = Automatic). They will always be redirected to their assigned SessionHost.
     */
    Personal: "Personal",
    /**
     * Users get a new (random) SessionHost every time it connects to the HostPool.
     */
    Pooled: "Pooled",
    /**
     * Users assign their own machines, load balancing logic remains the same as Personal. PersonalDesktopAssignmentType must be Direct.
     */
    BYODesktop: "BYODesktop",
} as const;

/**
 * HostPool type for desktop.
 */
export type HostPoolType = (typeof HostPoolType)[keyof typeof HostPoolType];

export const HostpoolPublicNetworkAccess = {
    Enabled: "Enabled",
    Disabled: "Disabled",
    EnabledForSessionHostsOnly: "EnabledForSessionHostsOnly",
    EnabledForClientsOnly: "EnabledForClientsOnly",
} as const;

/**
 * Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
 */
export type HostpoolPublicNetworkAccess = (typeof HostpoolPublicNetworkAccess)[keyof typeof HostpoolPublicNetworkAccess];

export const LoadBalancerType = {
    BreadthFirst: "BreadthFirst",
    DepthFirst: "DepthFirst",
    Persistent: "Persistent",
    MultiplePersistent: "MultiplePersistent",
} as const;

/**
 * The type of the load balancer.
 */
export type LoadBalancerType = (typeof LoadBalancerType)[keyof typeof LoadBalancerType];

export const ManagedPrivateUDP = {
    Default: "Default",
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Default: AVD-wide settings are used to determine connection availability, Enabled: UDP will attempt this connection type when making connections. This means that this connection is possible, but is not guaranteed, as there are other factors that may prevent this connection type, Disabled: UDP will not attempt this connection type when making connections
 */
export type ManagedPrivateUDP = (typeof ManagedPrivateUDP)[keyof typeof ManagedPrivateUDP];

export const ManagedServiceIdentityType = {
    None: "None",
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned,UserAssigned",
} as const;

/**
 * Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
 */
export type ManagedServiceIdentityType = (typeof ManagedServiceIdentityType)[keyof typeof ManagedServiceIdentityType];

export const ManagementType = {
    Automated: "Automated",
    Standard: "Standard",
} as const;

/**
 * The type of management for this hostpool, Automated or Standard. The default value is Automated.
 */
export type ManagementType = (typeof ManagementType)[keyof typeof ManagementType];

export const PackageTimestamped = {
    Timestamped: "Timestamped",
    NotTimestamped: "NotTimestamped",
} as const;

/**
 * Is package timestamped so it can ignore the certificate expiry date
 */
export type PackageTimestamped = (typeof PackageTimestamped)[keyof typeof PackageTimestamped];

export const PersonalDesktopAssignmentType = {
    Automatic: "Automatic",
    Direct: "Direct",
} as const;

/**
 * PersonalDesktopAssignment type for HostPool.
 */
export type PersonalDesktopAssignmentType = (typeof PersonalDesktopAssignmentType)[keyof typeof PersonalDesktopAssignmentType];

export const PreferredAppGroupType = {
    /**
     * This value is read only, it is not accepted on input.
     */
    None: "None",
    /**
     * Users access the full Windows desktop from a session host. Available with pooled or personal host pools.
     */
    Desktop: "Desktop",
    /**
     * Users access individual applications you select and publish to the application group. Available with pooled host pools only.
     */
    RailApplications: "RailApplications",
} as const;

/**
 * The type of preferred application group type, default to Desktop Application Group
 */
export type PreferredAppGroupType = (typeof PreferredAppGroupType)[keyof typeof PreferredAppGroupType];

export const PrivateEndpointServiceConnectionStatus = {
    Pending: "Pending",
    Approved: "Approved",
    Rejected: "Rejected",
} as const;

/**
 * Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
 */
export type PrivateEndpointServiceConnectionStatus = (typeof PrivateEndpointServiceConnectionStatus)[keyof typeof PrivateEndpointServiceConnectionStatus];

export const PublicNetworkAccess = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
 */
export type PublicNetworkAccess = (typeof PublicNetworkAccess)[keyof typeof PublicNetworkAccess];

export const PublicUDP = {
    Default: "Default",
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Default: AVD-wide settings are used to determine connection availability, Enabled: UDP will attempt this connection type when making connections. This means that this connection is possible, but is not guaranteed, as there are other factors that may prevent this connection type, Disabled: UDP will not attempt this connection type when making connections
 */
export type PublicUDP = (typeof PublicUDP)[keyof typeof PublicUDP];

export const RegistrationTokenOperation = {
    Delete: "Delete",
    None: "None",
    Update: "Update",
} as const;

/**
 * The type of resetting the token.
 */
export type RegistrationTokenOperation = (typeof RegistrationTokenOperation)[keyof typeof RegistrationTokenOperation];

export const RelayUDP = {
    Default: "Default",
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Default: AVD-wide settings are used to determine connection availability, Enabled: UDP will attempt this connection type when making connections. This means that this connection is possible, but is not guaranteed, as there are other factors that may prevent this connection type, Disabled: UDP will not attempt this connection type when making connections
 */
export type RelayUDP = (typeof RelayUDP)[keyof typeof RelayUDP];

export const RemoteApplicationType = {
    InBuilt: "InBuilt",
    MsixApplication: "MsixApplication",
} as const;

/**
 * Resource Type of Application.
 */
export type RemoteApplicationType = (typeof RemoteApplicationType)[keyof typeof RemoteApplicationType];

export const SSOSecretType = {
    SharedKey: "SharedKey",
    Certificate: "Certificate",
    SharedKeyInKeyVault: "SharedKeyInKeyVault",
    CertificateInKeyVault: "CertificateInKeyVault",
} as const;

/**
 * The type of single sign on Secret Type.
 */
export type SSOSecretType = (typeof SSOSecretType)[keyof typeof SSOSecretType];

export const ScalingHostPoolType = {
    /**
     * Users get a new (random) SessionHost every time it connects to the HostPool.
     */
    Pooled: "Pooled",
    /**
     * Users will be assigned a SessionHost either by administrators (PersonalDesktopAssignmentType = Direct) or upon connecting to the pool (PersonalDesktopAssignmentType = Automatic). They will always be redirected to their assigned SessionHost.
     */
    Personal: "Personal",
} as const;

/**
 * HostPool type for desktop.
 */
export type ScalingHostPoolType = (typeof ScalingHostPoolType)[keyof typeof ScalingHostPoolType];

export const ScalingMethod = {
    /**
     * Scaling will manage hosts in the host pool by power managing the hosts, but will not change the host pool size.
     */
    PowerManage: "PowerManage",
    /**
     * Scaling will manage the hosts in the host pool by power managing the hosts, as well as creating and deleting hosts to modify the host pool size. This requires the create delete object to be set, and the assigned hostpool to have a session host config property.
     */
    CreateDeletePowerManage: "CreateDeletePowerManage",
} as const;

/**
 * The desired scaling method to be used to scale the hosts in the assigned host pool.
 */
export type ScalingMethod = (typeof ScalingMethod)[keyof typeof ScalingMethod];

export const SessionHandlingOperation = {
    None: "None",
    Deallocate: "Deallocate",
    Hibernate: "Hibernate",
} as const;

/**
 * Action to be taken after a logoff during the ramp up period.
 */
export type SessionHandlingOperation = (typeof SessionHandlingOperation)[keyof typeof SessionHandlingOperation];

export const SessionHostComponentUpdateType = {
    /**
     * Agent and other agent side components are delivery schedule is controlled by WVD Infra.
     */
    Default: "Default",
    /**
     * TenantAdmin have opted in for Scheduled Component Update feature.
     */
    Scheduled: "Scheduled",
} as const;

/**
 * The type of maintenance for session host components.
 */
export type SessionHostComponentUpdateType = (typeof SessionHostComponentUpdateType)[keyof typeof SessionHostComponentUpdateType];

export const SessionHostLoadBalancingAlgorithm = {
    BreadthFirst: "BreadthFirst",
    DepthFirst: "DepthFirst",
} as const;

/**
 * Load balancing algorithm for ramp up period.
 */
export type SessionHostLoadBalancingAlgorithm = (typeof SessionHostLoadBalancingAlgorithm)[keyof typeof SessionHostLoadBalancingAlgorithm];

export const SetStartVMOnConnect = {
    Enable: "Enable",
    Disable: "Disable",
} as const;

/**
 * The desired configuration of Start VM On Connect for the hostpool during the ramp up phase. If this is disabled, session hosts must be turned on using rampUpAutoStartHosts or by turning them on manually.
 */
export type SetStartVMOnConnect = (typeof SetStartVMOnConnect)[keyof typeof SetStartVMOnConnect];

export const SkuTier = {
    Free: "Free",
    Basic: "Basic",
    Standard: "Standard",
    Premium: "Premium",
} as const;

/**
 * This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
 */
export type SkuTier = (typeof SkuTier)[keyof typeof SkuTier];

export const StartupBehavior = {
    /**
     * Session hosts will not be started by the service. This setting depends on Start VM on Connect to be enabled to start the session hosts.
     */
    None: "None",
    /**
     * Session hosts with an assigned user will be started during Ramp Up
     */
    WithAssignedUser: "WithAssignedUser",
    /**
     * All personal session hosts in the hostpool will be started during ramp up.
     */
    All: "All",
} as const;

/**
 * The desired startup behavior during the ramp up period for personal vms in the hostpool.
 */
export type StartupBehavior = (typeof StartupBehavior)[keyof typeof StartupBehavior];

export const StopHostsWhen = {
    ZeroSessions: "ZeroSessions",
    ZeroActiveSessions: "ZeroActiveSessions",
} as const;

/**
 * Specifies when to stop hosts during ramp down period.
 */
export type StopHostsWhen = (typeof StopHostsWhen)[keyof typeof StopHostsWhen];
