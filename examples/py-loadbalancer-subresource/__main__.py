from pulumi_azure_native import resources, network

# Create a new resource group
resource_group = resources.ResourceGroup('resource_group')

# Create public IP
public_ip = network.PublicIPAddress('public_ip',
                                    resource_group_name=resource_group.name,
                                    location=resource_group.location,
                                    sku=network.PublicIPAddressSkuArgs(
                                        name="Standard"),
                                    public_ip_allocation_method="Static"
                                    )

# Create a new LoadBalancer
load_balancer = network.LoadBalancer('load_balancer',
                                     location=resource_group.location,
                                     sku=network.LoadBalancerSkuArgs(
                                         name="Standard"),
                                     resource_group_name=resource_group.name,
                                     frontend_ip_configurations=[{
                                         "name": "PublicIPAddressFrontendIP",
                                         "public_ip_address": {
                                             "id": public_ip.id,
                                         },
                                     }],
                                     )

# Create the BackendAddressPool as a separate resource rather than a property of the load balancer.
backend_address_pool1 = network.LoadBalancerBackendAddressPool(
    'backend_pool1',
    resource_group_name=resource_group.name,
    load_balancer_name=load_balancer.name
)
