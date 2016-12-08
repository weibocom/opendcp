// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
