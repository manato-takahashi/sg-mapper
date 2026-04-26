package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := ec2.NewFromConfig(cfg)

	output, err := client.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, sg := range output.SecurityGroups {
		fmt.Println("===", *sg.GroupId, *sg.GroupName, "===")
		fmt.Println("ingress:")
		for _, ingress := range sg.IpPermissions {
			for _, ipRange := range ingress.IpRanges {
				if ingress.FromPort != nil {
					fmt.Println(*ingress.IpProtocol, *ingress.FromPort, *ingress.ToPort, *ipRange.CidrIp)
				} else {
					fmt.Println("* * * ", *ipRange.CidrIp)
				}
			}
		}

		fmt.Println("egress:")
		for _, egress := range sg.IpPermissionsEgress {
			for _, ipRange := range egress.IpRanges {
				if egress.FromPort != nil {
					fmt.Println(*egress.IpProtocol, *egress.FromPort, *egress.ToPort, *ipRange.CidrIp)
				} else {
					fmt.Println("* * * ", *ipRange.CidrIp)
				}
			}
		}
	}
}
