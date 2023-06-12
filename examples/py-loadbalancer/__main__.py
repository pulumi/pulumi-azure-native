# Copyright 2023, Pulumi Corporation.  All rights reserved.

from pulumi_azure_native import (resources, network)

# Create a resource group.
resource_group = resources.ResourceGroup("rg")

# Create a virtual network.
vnet = network.VirtualNetwork(
    "network",
    resource_group_name=resource_group.name,
    address_space=network.AddressSpaceArgs(
        address_prefixes=[
            "10.0.0.0/16",
        ],
    ),
    subnets=[
        network.SubnetArgs(
            name="default",
            address_prefix="10.0.1.0/24",
        ),
    ],
)

# Create a public IP address for the VM.
lb_public_ip = network.PublicIPAddress(
    "lb-public-ip",
    resource_group_name=resource_group.name,
    public_ip_allocation_method=network.IpAllocationMethod.STATIC,
    public_ip_address_version="IPV4",
    # required to be paired with below LB type
    sku=network.PublicIPAddressSkuArgs(
        name="Standard"
    )
)

# we are required to create the IDs for these resources unfortunately due to how the Azure API structures LoadBalancers as god objects.
# however, they are relative IDs - the actual physical IDs will be resolved by the provider.
lb_backend_name = "lb-backend"
lb_backend_id = "$self/backendAddressPools/" + lb_backend_name

lb_frontend_name = "lb-frontend"
lb_frontend_id = "$self/frontendIPConfigurations/" + lb_frontend_name

lb_probe_name = "lb-probe"
lb_probe_id = "$self/probes/" + lb_probe_name

lb = network.LoadBalancer(
    "lb",
    backend_address_pools=[
        network.BackendAddressPoolArgs(
            name=lb_backend_name,
        ),
    ],
    frontend_ip_configurations=[network.FrontendIPConfigurationArgs(
        name=lb_frontend_name,
        public_ip_address=network.PublicIPAddressArgs(
            id=lb_public_ip.id,
        ),
    )],
    load_balancing_rules=[network.LoadBalancingRuleArgs(
        backend_address_pool=network.SubResourceArgs(
            id=lb_backend_id,
        ),
        backend_port=80,
        disable_outbound_snat=True,
        enable_floating_ip=False,
        enable_tcp_reset=True,
        frontend_ip_configuration=network.SubResourceArgs(
            id=lb_frontend_id,
        ),
        frontend_port=80,
        idle_timeout_in_minutes=4,
        load_distribution="Default",
        name="lb-rule1",
        probe=network.SubResourceArgs(
            id=lb_probe_id,
        ),
        protocol="TCP",
    )],
    probes=[network.ProbeArgs(
        interval_in_seconds=5,
        name=lb_probe_name,
        number_of_probes=1,
        port=80,
        protocol="HTTP",
        request_path="/",
    )],
    resource_group_name=resource_group.name,
    sku=network.LoadBalancerSkuArgs(
        name="Standard",
    ),
)
