# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AgentPoolMode',
    'AgentPoolType',
    'Code',
    'GPUInstanceProfile',
    'KubeletDiskType',
    'OSDiskType',
    'OSSKU',
    'OSType',
    'ScaleDownMode',
    'ScaleSetEvictionPolicy',
    'ScaleSetPriority',
    'WorkloadRuntime',
]


class AgentPoolMode(str, Enum):
    """
    A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
    """
    SYSTEM = "System"
    """
    System agent pools are primarily for hosting critical system pods such as CoreDNS and metrics-server. System agent pools osType must be Linux. System agent pools VM SKU must have at least 2vCPUs and 4GB of memory.
    """
    USER = "User"
    """
    User agent pools are primarily for hosting your application pods.
    """


class AgentPoolType(str, Enum):
    """
    The type of Agent Pool.
    """
    VIRTUAL_MACHINE_SCALE_SETS = "VirtualMachineScaleSets"
    """
    Create an Agent Pool backed by a Virtual Machine Scale Set.
    """
    AVAILABILITY_SET = "AvailabilitySet"
    """
    Use of this is strongly discouraged.
    """


class Code(str, Enum):
    """
    Tells whether the cluster is Running or Stopped
    """
    RUNNING = "Running"
    """
    The cluster is running.
    """
    STOPPED = "Stopped"
    """
    The cluster is stopped.
    """


class GPUInstanceProfile(str, Enum):
    """
    GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
    """
    MIG1G = "MIG1g"
    MIG2G = "MIG2g"
    MIG3G = "MIG3g"
    MIG4G = "MIG4g"
    MIG7G = "MIG7g"


class KubeletDiskType(str, Enum):
    """
    Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
    """
    OS = "OS"
    """
    Kubelet will use the OS disk for its data.
    """
    TEMPORARY = "Temporary"
    """
    Kubelet will use the temporary disk for its data.
    """


class OSDiskType(str, Enum):
    """
    The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
    """
    MANAGED = "Managed"
    """
    Azure replicates the operating system disk for a virtual machine to Azure storage to avoid data loss should the VM need to be relocated to another host. Since containers aren't designed to have local state persisted, this behavior offers limited value while providing some drawbacks, including slower node provisioning and higher read/write latency.
    """
    EPHEMERAL = "Ephemeral"
    """
    Ephemeral OS disks are stored only on the host machine, just like a temporary disk. This provides lower read/write latency, along with faster node scaling and cluster upgrades.
    """


class OSSKU(str, Enum):
    """
    Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.
    """
    UBUNTU = "Ubuntu"
    CBL_MARINER = "CBLMariner"
    WINDOWS2019 = "Windows2019"
    WINDOWS2022 = "Windows2022"


class OSType(str, Enum):
    """
    The operating system type. The default is Linux.
    """
    LINUX = "Linux"
    """
    Use Linux.
    """
    WINDOWS = "Windows"
    """
    Use Windows.
    """


class ScaleDownMode(str, Enum):
    """
    This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
    """
    DELETE = "Delete"
    """
    Create new instances during scale up and remove instances during scale down.
    """
    DEALLOCATE = "Deallocate"
    """
    Attempt to start deallocated instances (if they exist) during scale up and deallocate instances during scale down.
    """


class ScaleSetEvictionPolicy(str, Enum):
    """
    This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.
    """
    DELETE = "Delete"
    """
    Nodes in the underlying Scale Set of the node pool are deleted when they're evicted.
    """
    DEALLOCATE = "Deallocate"
    """
    Nodes in the underlying Scale Set of the node pool are set to the stopped-deallocated state upon eviction. Nodes in the stopped-deallocated state count against your compute quota and can cause issues with cluster scaling or upgrading.
    """


class ScaleSetPriority(str, Enum):
    """
    The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
    """
    SPOT = "Spot"
    """
    Spot priority VMs will be used. There is no SLA for spot nodes. See [spot on AKS](https://docs.microsoft.com/azure/aks/spot-node-pool) for more information.
    """
    REGULAR = "Regular"
    """
    Regular VMs will be used.
    """


class WorkloadRuntime(str, Enum):
    """
    Determines the type of workload a node can run.
    """
    OCI_CONTAINER = "OCIContainer"
    """
    Nodes will use Kubelet to run standard OCI container workloads.
    """
    WASM_WASI = "WasmWasi"
    """
    Nodes will use Krustlet to run WASM workloads using the WASI provider (Preview).
    """
