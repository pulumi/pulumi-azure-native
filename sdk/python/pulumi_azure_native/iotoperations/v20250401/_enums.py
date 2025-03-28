# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'BrokerAuthenticationMethod',
    'BrokerMemoryProfile',
    'BrokerProtocolType',
    'BrokerResourceDefinitionMethods',
    'CertManagerIssuerKind',
    'CloudEventAttributeType',
    'DataLakeStorageAuthMethod',
    'DataflowEndpointAuthenticationSaslType',
    'DataflowEndpointFabricPathType',
    'DataflowEndpointKafkaAcks',
    'DataflowEndpointKafkaCompression',
    'DataflowEndpointKafkaPartitionStrategy',
    'DataflowMappingType',
    'EndpointType',
    'ExtendedLocationType',
    'FilterType',
    'InstanceFeatureMode',
    'KafkaAuthMethod',
    'ManagedServiceIdentityType',
    'MqttAuthMethod',
    'MqttRetainType',
    'OperationType',
    'OperationalMode',
    'OperatorValues',
    'PrivateKeyAlgorithm',
    'PrivateKeyRotationPolicy',
    'ServiceType',
    'SourceSerializationFormat',
    'StateStoreResourceDefinitionMethods',
    'StateStoreResourceKeyTypes',
    'SubscriberMessageDropStrategy',
    'TlsCertMethodMode',
    'TransformationSerializationFormat',
]


class BrokerAuthenticationMethod(str, Enum):
    """
    Custom authentication configuration.
    """
    CUSTOM = "Custom"
    """
    Custom authentication configuration.
    """
    SERVICE_ACCOUNT_TOKEN = "ServiceAccountToken"
    """
    ServiceAccountToken authentication configuration.
    """
    X509 = "X509"
    """
    X.509 authentication configuration.
    """


class BrokerMemoryProfile(str, Enum):
    """
    Memory profile of Broker.
    """
    TINY = "Tiny"
    """
    Tiny memory profile.
    """
    LOW = "Low"
    """
    Low memory profile.
    """
    MEDIUM = "Medium"
    """
    Medium memory profile.
    """
    HIGH = "High"
    """
    High memory profile.
    """


class BrokerProtocolType(str, Enum):
    """
    Enable or disable websockets.
    """
    MQTT = "Mqtt"
    """
    protocol broker
    """
    WEB_SOCKETS = "WebSockets"
    """
    protocol websocket
    """


class BrokerResourceDefinitionMethods(str, Enum):
    """
    Give access for a Broker method (i.e., Connect, Subscribe, or Publish).
    """
    CONNECT = "Connect"
    """
    Allowed Connecting to Broker
    """
    PUBLISH = "Publish"
    """
    Allowed Publishing to Broker
    """
    SUBSCRIBE = "Subscribe"
    """
    Allowed Subscribing to Broker
    """


class CertManagerIssuerKind(str, Enum):
    """
    kind of issuer (Issuer or ClusterIssuer).
    """
    ISSUER = "Issuer"
    """
    Issuer kind.
    """
    CLUSTER_ISSUER = "ClusterIssuer"
    """
    ClusterIssuer kind.
    """


class CloudEventAttributeType(str, Enum):
    """
    Cloud event mapping config.
    """
    PROPAGATE = "Propagate"
    """
    Propagate type
    """
    CREATE_OR_REMAP = "CreateOrRemap"
    """
    CreateOrRemap type
    """


class DataLakeStorageAuthMethod(str, Enum):
    """
    Mode of Authentication.
    """
    SYSTEM_ASSIGNED_MANAGED_IDENTITY = "SystemAssignedManagedIdentity"
    """
    SystemAssignedManagedIdentity type
    """
    USER_ASSIGNED_MANAGED_IDENTITY = "UserAssignedManagedIdentity"
    """
    UserAssignedManagedIdentity type
    """
    ACCESS_TOKEN = "AccessToken"
    """
    AccessToken Option
    """


class DataflowEndpointAuthenticationSaslType(str, Enum):
    """
    Type of SASL authentication. Can be PLAIN, SCRAM-SHA-256, or SCRAM-SHA-512.
    """
    PLAIN = "Plain"
    """
    PLAIN Type
    """
    SCRAM_SHA256 = "ScramSha256"
    """
    SCRAM_SHA_256 Type
    """
    SCRAM_SHA512 = "ScramSha512"
    """
    SCRAM_SHA_512 Type
    """


class DataflowEndpointFabricPathType(str, Enum):
    """
    Type of location of the data in the workspace. Can be either tables or files.
    """
    FILES = "Files"
    """
    FILES Type
    """
    TABLES = "Tables"
    """
    TABLES Type
    """


class DataflowEndpointKafkaAcks(str, Enum):
    """
    Kafka acks. Can be all, one, or zero. No effect if the endpoint is used as a source.
    """
    ZERO = "Zero"
    """
    ZERO Option
    """
    ONE = "One"
    """
    ONE Option
    """
    ALL = "All"
    """
    ALL Option
    """


class DataflowEndpointKafkaCompression(str, Enum):
    """
    Compression. Can be none, gzip, lz4, or snappy. No effect if the endpoint is used as a source.
    """
    NONE = "None"
    """
    NONE Option
    """
    GZIP = "Gzip"
    """
    Gzip Option
    """
    SNAPPY = "Snappy"
    """
    SNAPPY Option
    """
    LZ4 = "Lz4"
    """
    LZ4 Option
    """


class DataflowEndpointKafkaPartitionStrategy(str, Enum):
    """
    Partition handling strategy. Can be default or static. No effect if the endpoint is used as a source.
    """
    DEFAULT = "Default"
    """
    Default: Assigns messages to random partitions, using a round-robin algorithm.
    """
    STATIC = "Static"
    """
    Static: Assigns messages to a fixed partition number that's derived from the instance ID of the dataflow.
    """
    TOPIC = "Topic"
    """
    TOPIC Option
    """
    PROPERTY = "Property"
    """
    PROPERTY Option
    """


class DataflowMappingType(str, Enum):
    """
    Type of transformation.
    """
    NEW_PROPERTIES = "NewProperties"
    """
    New Properties type
    """
    RENAME = "Rename"
    """
    Rename type
    """
    COMPUTE = "Compute"
    """
    Compute type
    """
    PASS_THROUGH = "PassThrough"
    """
    Pass-through type
    """
    BUILT_IN_FUNCTION = "BuiltInFunction"
    """
    Built in function type
    """


class EndpointType(str, Enum):
    """
    Endpoint Type.
    """
    DATA_EXPLORER = "DataExplorer"
    """
    Azure Data Explorer Type
    """
    DATA_LAKE_STORAGE = "DataLakeStorage"
    """
    Azure Data Lake Type
    """
    FABRIC_ONE_LAKE = "FabricOneLake"
    """
    Microsoft Fabric Type
    """
    KAFKA = "Kafka"
    """
    Kafka Type
    """
    LOCAL_STORAGE = "LocalStorage"
    """
    Local Storage Type
    """
    MQTT = "Mqtt"
    """
    Broker Type
    """


class ExtendedLocationType(str, Enum):
    """
    Type of ExtendedLocation.
    """
    CUSTOM_LOCATION = "CustomLocation"
    """
    CustomLocation type
    """


class FilterType(str, Enum):
    """
    The type of dataflow operation.
    """
    FILTER = "Filter"
    """
    Filter type
    """


class InstanceFeatureMode(str, Enum):
    """
    The state of the feature.
    """
    STABLE = "Stable"
    """
    Opt in to enable a stable feature
    """
    PREVIEW = "Preview"
    """
    Opt in to enable a preview feature
    """
    DISABLED = "Disabled"
    """
    Opt out of a feature
    """


class KafkaAuthMethod(str, Enum):
    """
    Mode of Authentication.
    """
    SYSTEM_ASSIGNED_MANAGED_IDENTITY = "SystemAssignedManagedIdentity"
    """
    SystemAssignedManagedIdentity type
    """
    USER_ASSIGNED_MANAGED_IDENTITY = "UserAssignedManagedIdentity"
    """
    UserAssignedManagedIdentity type
    """
    SASL = "Sasl"
    """
    Sasl Option
    """
    X509_CERTIFICATE = "X509Certificate"
    """
    x509Certificate Option
    """
    ANONYMOUS = "Anonymous"
    """
    Anonymous Option
    """


class ManagedServiceIdentityType(str, Enum):
    """
    Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned,UserAssigned"


class MqttAuthMethod(str, Enum):
    """
    Mode of Authentication.
    """
    SYSTEM_ASSIGNED_MANAGED_IDENTITY = "SystemAssignedManagedIdentity"
    """
    SystemAssignedManagedIdentity type
    """
    USER_ASSIGNED_MANAGED_IDENTITY = "UserAssignedManagedIdentity"
    """
    UserAssignedManagedIdentity type
    """
    SERVICE_ACCOUNT_TOKEN = "ServiceAccountToken"
    """
    ServiceAccountToken Option
    """
    X509_CERTIFICATE = "X509Certificate"
    """
    x509Certificate Option
    """
    ANONYMOUS = "Anonymous"
    """
    Anonymous Option
    """


class MqttRetainType(str, Enum):
    """
    Whether or not to keep the retain setting.
    """
    KEEP = "Keep"
    """
    Retain the messages.
    """
    NEVER = "Never"
    """
    Never retain messages.
    """


class OperationType(str, Enum):
    """
    Type of operation.
    """
    SOURCE = "Source"
    """
    Dataflow Source Operation
    """
    DESTINATION = "Destination"
    """
    Dataflow Destination Operation
    """
    BUILT_IN_TRANSFORMATION = "BuiltInTransformation"
    """
    Dataflow BuiltIn Transformation Operation
    """


class OperationalMode(str, Enum):
    """
    Mode properties
    """
    ENABLED = "Enabled"
    """
    Enabled is equivalent to True
    """
    DISABLED = "Disabled"
    """
    Disabled is equivalent to False.
    """


class OperatorValues(str, Enum):
    """
    operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
    """
    IN_ = "In"
    """
    In operator.
    """
    NOT_IN = "NotIn"
    """
    NotIn operator.
    """
    EXISTS = "Exists"
    """
    Exists operator.
    """
    DOES_NOT_EXIST = "DoesNotExist"
    """
    DoesNotExist operator.
    """


class PrivateKeyAlgorithm(str, Enum):
    """
    algorithm for private key.
    """
    EC256 = "Ec256"
    """
    Algorithm - ec256.
    """
    EC384 = "Ec384"
    """
    Algorithm - ec384.
    """
    EC521 = "Ec521"
    """
    Algorithm - ec521.
    """
    ED25519 = "Ed25519"
    """
    Algorithm - ed25519.
    """
    RSA2048 = "Rsa2048"
    """
    Algorithm - rsa2048.
    """
    RSA4096 = "Rsa4096"
    """
    Algorithm - rsa4096.
    """
    RSA8192 = "Rsa8192"
    """
    Algorithm - rsa8192.
    """


class PrivateKeyRotationPolicy(str, Enum):
    """
    cert-manager private key rotationPolicy.
    """
    ALWAYS = "Always"
    """
    Rotation Policy - Always.
    """
    NEVER = "Never"
    """
    Rotation Policy - Never.
    """


class ServiceType(str, Enum):
    """
    Kubernetes Service type of this listener.
    """
    CLUSTER_IP = "ClusterIp"
    """
    Cluster IP Service.
    """
    LOAD_BALANCER = "LoadBalancer"
    """
    Load Balancer Service.
    """
    NODE_PORT = "NodePort"
    """
    Node Port Service.
    """


class SourceSerializationFormat(str, Enum):
    """
    Content is a JSON Schema. Allowed: JSON Schema/draft-7.
    """
    JSON = "Json"
    """
    JSON Format
    """


class StateStoreResourceDefinitionMethods(str, Enum):
    """
    Give access for `Read`, `Write` and `ReadWrite` access level.
    """
    READ = "Read"
    """
    Get/KeyNotify from Store
    """
    WRITE = "Write"
    """
    Set/Delete in Store
    """
    READ_WRITE = "ReadWrite"
    """
    Allowed all operations on Store - Get/KeyNotify/Set/Delete
    """


class StateStoreResourceKeyTypes(str, Enum):
    """
    Allowed keyTypes pattern, string, binary. The key type used for matching, for example pattern tries to match the key to a glob-style pattern and string checks key is equal to value provided in keys.
    """
    PATTERN = "Pattern"
    """
    Key type - pattern
    """
    STRING = "String"
    """
    Key type - string
    """
    BINARY = "Binary"
    """
    Key type - binary
    """


class SubscriberMessageDropStrategy(str, Enum):
    """
    The strategy to use for dropping messages from the queue.
    """
    NONE = "None"
    """
    Messages are never dropped.
    """
    DROP_OLDEST = "DropOldest"
    """
    The oldest message is dropped.
    """


class TlsCertMethodMode(str, Enum):
    """
    Mode of TLS server certificate management.
    """
    AUTOMATIC = "Automatic"
    """
    Automatic TLS server certificate configuration.
    """
    MANUAL = "Manual"
    """
    Manual TLS server certificate configuration.
    """


class TransformationSerializationFormat(str, Enum):
    """
    Serialization format. Optional; defaults to JSON. Allowed value JSON Schema/draft-7, Parquet. Default: Json
    """
    DELTA = "Delta"
    """
    Delta Format
    """
    JSON = "Json"
    """
    JSON Format
    """
    PARQUET = "Parquet"
    """
    Parquet Format
    """
