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

// Contains the output of DescribeVpcs.
type VpcsResp struct {

	// Information about one or more VPCs.
	Vpcs []Vpc `locationName:"vpcSet" locationNameList:"item" type:"list"`
}

// Contains the output of CreateVpc
type VpcResp struct {
	Vpc Vpc `locationName:"vpc" type:"structure"`
}

// Describes a VPC.
type Vpc struct {

	// The CIDR block for the VPC.
	CidrBlock string `locationName:"cidrBlock" type:"string"`

	// The ID of the set of DHCP options you've associated with the VPC (or default
	// if the default options are associated with the VPC).
	DhcpOptionsId string `locationName:"dhcpOptionsId" type:"string"`

	// The allowed tenancy of instances launched into the VPC.
	InstanceTenancy string `locationName:"instanceTenancy" type:"string" enum:"Tenancy"`

	// Indicates whether the VPC is the default VPC.
	IsDefault bool `locationName:"isDefault" type:"boolean"`

	// The current state of the VPC.
	State string `locationName:"state" type:"string" enum:"VpcState"`

	// Any tags assigned to the VPC.
	Tags []Tag `locationName:"tagSet" locationNameList:"item" type:"list"`

	// The ID of the VPC.
	VpcId string `locationName:"vpcId" type:"string"`
}

// Contains the parameters for CreateVpc.
type CreateVpcParam struct {

	// The name of provider
	Provider string `locationName:"provider" type:"string" required:"true"`

	// The network range for the VPC, in CIDR notation. For example, 10.0.0.0/16.
	CidrBlock string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun bool `locationName:"dryRun" type:"boolean"`

	// The tenancy options for instances launched into the VPC. For default, instances
	// are launched with shared tenancy by default. You can launch instances with
	// any tenancy into a shared tenancy VPC. For dedicated, instances are launched
	// as dedicated tenancy instances by default. You can only launch instances
	// with a tenancy of dedicated or host into a dedicated tenancy VPC.
	//
	//  Important: The host value cannot be used with this parameter. Use the default
	// or dedicated values only.
	//
	// Default: default
	InstanceTenancy string `locationName:"instanceTenancy" type:"string" enum:"Tenancy"`
}
