# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AccessKeyPermissions',
    'ActionType',
    'AllowType',
    'AuthType',
    'AzureResourceType',
    'ClientType',
    'DeleteOrUpdateBehavior',
    'DryrunActionName',
    'SecretType',
    'TargetServiceType',
    'VNetSolutionType',
]


class AccessKeyPermissions(str, Enum):
    READ = "Read"
    WRITE = "Write"
    LISTEN = "Listen"
    SEND = "Send"
    MANAGE = "Manage"


class ActionType(str, Enum):
    """
    Optional. Indicates public network solution. If enable, enable public network access of target service with best try. Default is enable. If optOut, opt out public network access configuration.
    """
    ENABLE = "enable"
    OPT_OUT = "optOut"


class AllowType(str, Enum):
    """
    Allow caller client IP to access the target service if true. the property is used when connecting local application to target service.
    """
    TRUE = "true"
    FALSE = "false"


class AuthType(str, Enum):
    """
    The authentication type.
    """
    SYSTEM_ASSIGNED_IDENTITY = "systemAssignedIdentity"
    USER_ASSIGNED_IDENTITY = "userAssignedIdentity"
    SERVICE_PRINCIPAL_SECRET = "servicePrincipalSecret"
    SERVICE_PRINCIPAL_CERTIFICATE = "servicePrincipalCertificate"
    SECRET = "secret"
    ACCESS_KEY = "accessKey"
    USER_ACCOUNT = "userAccount"


class AzureResourceType(str, Enum):
    """
    The azure resource type.
    """
    KEY_VAULT = "KeyVault"


class ClientType(str, Enum):
    """
    The application client type
    """
    NONE = "none"
    DOTNET = "dotnet"
    JAVA = "java"
    PYTHON = "python"
    GO = "go"
    PHP = "php"
    RUBY = "ruby"
    DJANGO = "django"
    NODEJS = "nodejs"
    SPRING_BOOT = "springBoot"
    KAFKA_SPRING_BOOT = "kafka-springBoot"
    DAPR = "dapr"


class DeleteOrUpdateBehavior(str, Enum):
    """
    Indicates whether to clean up previous operation when Linker is updating or deleting
    """
    DEFAULT = "Default"
    FORCED_CLEANUP = "ForcedCleanup"


class DryrunActionName(str, Enum):
    """
    The name of action for you dryrun job.
    """
    CREATE_OR_UPDATE = "createOrUpdate"


class SecretType(str, Enum):
    """
    The secret type.
    """
    RAW_VALUE = "rawValue"
    KEY_VAULT_SECRET_URI = "keyVaultSecretUri"
    KEY_VAULT_SECRET_REFERENCE = "keyVaultSecretReference"


class TargetServiceType(str, Enum):
    """
    The target service type.
    """
    AZURE_RESOURCE = "AzureResource"
    CONFLUENT_BOOTSTRAP_SERVER = "ConfluentBootstrapServer"
    CONFLUENT_SCHEMA_REGISTRY = "ConfluentSchemaRegistry"
    SELF_HOSTED_SERVER = "SelfHostedServer"


class VNetSolutionType(str, Enum):
    """
    Type of VNet solution.
    """
    SERVICE_ENDPOINT = "serviceEndpoint"
    PRIVATE_LINK = "privateLink"
