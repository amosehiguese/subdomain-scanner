from aws_cdk import Stack
from aws_cdk import aws_ec2 as ec2
from construct import Construct

class VpcStack(Stack):
    """
        This stack deploys a VPC with six subnets spread across two AZs
    """
    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        self.vpc = ec2.Vpc(
            self,
            id="Vpc",
            ip_addresses=ec2.IpAddresses.cidr("10.0.0.0/16"),
            max_azs=2,
        )
