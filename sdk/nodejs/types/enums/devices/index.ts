// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20200901preview from "./v20200901preview";
import * as v20220430preview from "./v20220430preview";
import * as v20221115preview from "./v20221115preview";
import * as v20221212 from "./v20221212";
import * as v20230301preview from "./v20230301preview";
import * as v20230630 from "./v20230630";
import * as v20230630preview from "./v20230630preview";
import * as v20250201preview from "./v20250201preview";

export {
    v20200901preview,
    v20220430preview,
    v20221115preview,
    v20221212,
    v20230301preview,
    v20230630,
    v20230630preview,
    v20250201preview,
};

export const AccessRights = {
    RegistryRead: "RegistryRead",
    RegistryWrite: "RegistryWrite",
    ServiceConnect: "ServiceConnect",
    DeviceConnect: "DeviceConnect",
    RegistryRead_RegistryWrite: "RegistryRead, RegistryWrite",
    RegistryRead_ServiceConnect: "RegistryRead, ServiceConnect",
    RegistryRead_DeviceConnect: "RegistryRead, DeviceConnect",
    RegistryWrite_ServiceConnect: "RegistryWrite, ServiceConnect",
    RegistryWrite_DeviceConnect: "RegistryWrite, DeviceConnect",
    ServiceConnect_DeviceConnect: "ServiceConnect, DeviceConnect",
    RegistryRead_RegistryWrite_ServiceConnect: "RegistryRead, RegistryWrite, ServiceConnect",
    RegistryRead_RegistryWrite_DeviceConnect: "RegistryRead, RegistryWrite, DeviceConnect",
    RegistryRead_ServiceConnect_DeviceConnect: "RegistryRead, ServiceConnect, DeviceConnect",
    RegistryWrite_ServiceConnect_DeviceConnect: "RegistryWrite, ServiceConnect, DeviceConnect",
    RegistryRead_RegistryWrite_ServiceConnect_DeviceConnect: "RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect",
} as const;

/**
 * The permissions assigned to the shared access policy.
 */
export type AccessRights = (typeof AccessRights)[keyof typeof AccessRights];

export const AccessRightsDescription = {
    ServiceConfig: "ServiceConfig",
    EnrollmentRead: "EnrollmentRead",
    EnrollmentWrite: "EnrollmentWrite",
    DeviceConnect: "DeviceConnect",
    RegistrationStatusRead: "RegistrationStatusRead",
    RegistrationStatusWrite: "RegistrationStatusWrite",
} as const;

/**
 * Rights that this key has.
 */
export type AccessRightsDescription = (typeof AccessRightsDescription)[keyof typeof AccessRightsDescription];

export const AllocationPolicy = {
    Hashed: "Hashed",
    GeoLatency: "GeoLatency",
    Static: "Static",
} as const;

/**
 * Allocation policy to be used by this provisioning service.
 */
export type AllocationPolicy = (typeof AllocationPolicy)[keyof typeof AllocationPolicy];

export const AuthenticationType = {
    KeyBased: "keyBased",
    IdentityBased: "identityBased",
} as const;

/**
 * Specifies authentication type being used for connecting to the storage account.
 */
export type AuthenticationType = (typeof AuthenticationType)[keyof typeof AuthenticationType];

export const Capabilities = {
    None: "None",
    DeviceManagement: "DeviceManagement",
} as const;

/**
 * The capabilities and features enabled for the IoT hub.
 */
export type Capabilities = (typeof Capabilities)[keyof typeof Capabilities];

export const DefaultAction = {
    Deny: "Deny",
    Allow: "Allow",
} as const;

/**
 * Default Action for Network Rule Set
 */
export type DefaultAction = (typeof DefaultAction)[keyof typeof DefaultAction];

export const IotDpsSku = {
    S1: "S1",
} as const;

/**
 * Sku name.
 */
export type IotDpsSku = (typeof IotDpsSku)[keyof typeof IotDpsSku];

export const IotHubSku = {
    F1: "F1",
    S1: "S1",
    S2: "S2",
    S3: "S3",
    B1: "B1",
    B2: "B2",
    B3: "B3",
} as const;

/**
 * The name of the SKU.
 */
export type IotHubSku = (typeof IotHubSku)[keyof typeof IotHubSku];

export const IpFilterActionType = {
    Accept: "Accept",
    Reject: "Reject",
} as const;

/**
 * The desired action for requests captured by this rule.
 */
export type IpFilterActionType = (typeof IpFilterActionType)[keyof typeof IpFilterActionType];

export const IpFilterTargetType = {
    All: "all",
    ServiceApi: "serviceApi",
    DeviceApi: "deviceApi",
} as const;

/**
 * Target for requests captured by this rule.
 */
export type IpFilterTargetType = (typeof IpFilterTargetType)[keyof typeof IpFilterTargetType];

export const NetworkRuleIPAction = {
    Allow: "Allow",
} as const;

/**
 * IP Filter Action
 */
export type NetworkRuleIPAction = (typeof NetworkRuleIPAction)[keyof typeof NetworkRuleIPAction];

export const PrivateLinkServiceConnectionStatus = {
    Pending: "Pending",
    Approved: "Approved",
    Rejected: "Rejected",
    Disconnected: "Disconnected",
} as const;

/**
 * The status of a private endpoint connection
 */
export type PrivateLinkServiceConnectionStatus = (typeof PrivateLinkServiceConnectionStatus)[keyof typeof PrivateLinkServiceConnectionStatus];

export const PublicNetworkAccess = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Whether requests from Public Network are allowed
 */
export type PublicNetworkAccess = (typeof PublicNetworkAccess)[keyof typeof PublicNetworkAccess];

export const ResourceIdentityType = {
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned, UserAssigned",
    None: "None",
} as const;

/**
 * The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
 */
export type ResourceIdentityType = (typeof ResourceIdentityType)[keyof typeof ResourceIdentityType];

export const RoutingSource = {
    Invalid: "Invalid",
    DeviceMessages: "DeviceMessages",
    TwinChangeEvents: "TwinChangeEvents",
    DeviceLifecycleEvents: "DeviceLifecycleEvents",
    DeviceJobLifecycleEvents: "DeviceJobLifecycleEvents",
    DigitalTwinChangeEvents: "DigitalTwinChangeEvents",
    DeviceConnectionStateEvents: "DeviceConnectionStateEvents",
    MqttBrokerMessages: "MqttBrokerMessages",
} as const;

/**
 * The source that the routing rule is to be applied to, such as DeviceMessages.
 */
export type RoutingSource = (typeof RoutingSource)[keyof typeof RoutingSource];

export const State = {
    Activating: "Activating",
    Active: "Active",
    Deleting: "Deleting",
    Deleted: "Deleted",
    ActivationFailed: "ActivationFailed",
    DeletionFailed: "DeletionFailed",
    Transitioning: "Transitioning",
    Suspending: "Suspending",
    Suspended: "Suspended",
    Resuming: "Resuming",
    FailingOver: "FailingOver",
    FailoverFailed: "FailoverFailed",
} as const;

/**
 * Current state of the provisioning service.
 */
export type State = (typeof State)[keyof typeof State];
