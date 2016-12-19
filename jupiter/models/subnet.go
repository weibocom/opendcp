/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */



package models

// Contains the output of DescribeSubnets.
type SubnetsResp struct {

	// Information about one or more subnets.
	Subnets []Subnet `locationName:"subnetSet" locationNameList:"item" type:"list"`
}

// Describes a subnet.
type Subnet struct {

	// The Availability Zone of the subnet.
	AvailabilityZone string `locationName:"availabilityZone" type:"string"`

	// The number of unused IP addresses in the subnet. Note that the IP addresses
	// for any stopped instances are considered unavailable.
	AvailableIpAddressCount int64 `locationName:"availableIpAddressCount" type:"integer"`

	// The CIDR block assigned to the subnet.
	CidrBlock string `locationName:"cidrBlock" type:"string"`

	// Indicates whether this is the default subnet for the Availability Zone.
	DefaultForAz bool `locationName:"defaultForAz" type:"boolean"`

	// Indicates whether instances launched in this subnet receive a public IP address.
	MapPublicIpOnLaunch bool `locationName:"mapPublicIpOnLaunch" type:"boolean"`

	// The current state of the subnet.
	State string `locationName:"state" type:"string" enum:"SubnetState"`

	// The ID of the subnet.
	SubnetId string `locationName:"subnetId" type:"string"`

	// Any tags assigned to the subnet.
	Tags []Tag `locationName:"tagSet" locationNameList:"item" type:"list"`

	// The ID of the VPC the subnet is in.
	VpcId string `locationName:"vpcId" type:"string"`
}

// Contains the parameters for CreateSubnet.
type CreateSubnetParam struct {

	// The name of provider
	Provider string `locationName:"provider" type:"string" required:"true"`

	// The Availability Zone for the subnet.
	//
	// Default: AWS selects one for you. If you create more than one subnet in
	// your VPC, we may not necessarily select a different zone for each subnet.
	AvailabilityZone string `type:"string"`

	// The network range for the subnet, in CIDR notation. For example, 10.0.0.0/24.
	CidrBlock string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	VpcId string `type:"string" required:"true"`
}

// Contains the output of CreateSubnet.
type SubnetResp struct {

	// Information about the subnet.
	Subnet Subnet `locationName:"subnet" type:"structure"`
}
