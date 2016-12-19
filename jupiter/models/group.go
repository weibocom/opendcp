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

// Describes a security group.
type GroupIdentifier struct {

	// The ID of the security group.
	GroupId string `locationName:"groupId" type:"string"`

	// The name of the security group.
	GroupName string `locationName:"groupName" type:"string"`
}

// Contains the parameters for CreateSecurityGroup.
type CreateSecurityGroupParam struct {

	// The name of provider
	Provider string `locationName:"provider" type:"string" required:"true"`

	// A description for the security group. This is informational only.
	//
	// Constraints: Up to 255 characters in length
	//
	// Constraints for Classic: ASCII characters
	//
	// Constraints for VPC: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	Description string `locationName:"GroupDescription" type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun bool `locationName:"dryRun" type:"boolean"`

	// The name of the security group.
	//
	// Constraints: Up to 255 characters in length
	//
	// Constraints for Classic: ASCII characters
	//
	// Constraints for VPC: a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=&;{}!$*
	GroupName string `type:"string" required:"true"`

	// [VPC] The ID of the VPC. Required for EC2-VPC.
	VpcId string `type:"string"`
}

// Contains the output of CreateSecurityGroups.
type CreateSecurityGroupsResp struct {

	// Information about one or more security groups.
	SecurityGroups []SecurityGroup `locationName:"securityGroupInfo" locationNameList:"item" type:"list"`
}

// Describes a security group
type SecurityGroup struct {

	// A description of the security group.
	Description string `locationName:"groupDescription" type:"string"`

	// The ID of the security group.
	GroupId string `locationName:"groupId" type:"string"`

	// The name of the security group.
	GroupName string `locationName:"groupName" type:"string"`

	// One or more inbound rules associated with the security group.
	IpPermissions []IpPermission `locationName:"ipPermissions" locationNameList:"item" type:"list"`

	// [EC2-VPC] One or more outbound rules associated with the security group.
	IpPermissionsEgress []IpPermission `locationName:"ipPermissionsEgress" locationNameList:"item" type:"list"`

	// The AWS account ID of the owner of the security group.
	OwnerId string `locationName:"ownerId" type:"string"`

	// Any tags assigned to the security group.
	Tags []Tag `locationName:"tagSet" locationNameList:"item" type:"list"`

	// [EC2-VPC] The ID of the VPC for the security group.
	VpcId string `locationName:"vpcId" type:"string"`
}

// Describes a security group and AWS account ID pair.
type UserIdGroupPair struct {

	// The ID of the security group.
	GroupId string `locationName:"groupId" type:"string"`

	// The name of the security group. In a request, use this parameter for a security
	// group in EC2-Classic or a default VPC only. For a security group in a nondefault
	// VPC, use the security group ID.
	GroupName string `locationName:"groupName" type:"string"`

	// The status of a VPC peering connection, if applicable.
	PeeringStatus string `locationName:"peeringStatus" type:"string"`

	// The ID of an AWS account. For a referenced security group in another VPC,
	// the account ID of the referenced security group is returned.
	//
	// [EC2-Classic] Required when adding or removing rules that reference a security
	// group in another AWS account.
	UserId string `locationName:"userId" type:"string"`

	// The ID of the VPC for the referenced security group, if applicable.
	VpcId string `locationName:"vpcId" type:"string"`

	// The ID of the VPC peering connection, if applicable.
	VpcPeeringConnectionId string `locationName:"vpcPeeringConnectionId" type:"string"`
}

// Contains the output of CreateSecurityGroup.
type SecurityGroupResp struct {

	// The ID of the security group.
	GroupId string `locationName:"groupId" type:"string"`
}

// Contains the parameters for AuthorizeSecurityGroupIngress.
type AuthorizeSecurityGroupIngress struct {

	// The name of provider
	Provider string `locationName:"provider" type:"string" required:"true"`

	// The CIDR IP address range. You can't specify this parameter when specifying
	// a source security group.
	CidrIp string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun bool `locationName:"dryRun" type:"boolean"`

	// The start of port range for the TCP and UDP protocols, or an ICMP type number.
	// For the ICMP type number, use -1 to specify all ICMP types.
	FromPort int64 `type:"integer"`

	// The ID of the security group. Required for a nondefault VPC.
	GroupId string `type:"string"`

	// [EC2-Classic, default VPC] The name of the security group.
	GroupName string `type:"string"`

	// A set of IP permissions. Can be used to specify multiple rules in a single
	// command.
	IpPermissions []IpPermission `locationNameList:"item" type:"list"`

	// The IP protocol name (tcp, udp, icmp) or number (see Protocol Numbers (http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).
	// (VPC only) Use -1 to specify all.
	IpProtocol string `type:"string"`

	// [EC2-Classic, default VPC] The name of the source security group. You can't
	// specify this parameter in combination with the following parameters: the
	// CIDR IP address range, the start of the port range, the IP protocol, and
	// the end of the port range. Creates rules that grant full ICMP, UDP, and TCP
	// access. To create a rule with a specific IP protocol and port range, use
	// a set of IP permissions instead. For EC2-VPC, the source security group must
	// be in the same VPC.
	SourceSecurityGroupName string `type:"string"`

	// [EC2-Classic] The AWS account number for the source security group, if the
	// source security group is in a different account. You can't specify this parameter
	// in combination with the following parameters: the CIDR IP address range,
	// the IP protocol, the start of the port range, and the end of the port range.
	// Creates rules that grant full ICMP, UDP, and TCP access. To create a rule
	// with a specific IP protocol and port range, use a set of IP permissions instead.
	SourceSecurityGroupOwnerId string `type:"string"`

	// The end of port range for the TCP and UDP protocols, or an ICMP code number.
	// For the ICMP code number, use -1 to specify all ICMP codes for the ICMP type.
	ToPort int64 `type:"integer"`
}

// Describes a security group rule.
type IpPermission struct {

	// The start of port range for the TCP and UDP protocols, or an ICMP type number.
	// A value of -1 indicates all ICMP types.
	FromPort int64 `locationName:"fromPort" type:"integer"`

	// The IP protocol name (for tcp, udp, and icmp) or number (see Protocol Numbers
	// (http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).
	//
	// [EC2-VPC only] When you authorize or revoke security group rules, you can
	// use -1 to specify all.
	IpProtocol string `locationName:"ipProtocol" type:"string"`

	// One or more IP ranges.
	IpRanges []IpRange `locationName:"ipRanges" locationNameList:"item" type:"list"`

	// (Valid for AuthorizeSecurityGroupEgress, RevokeSecurityGroupEgress and DescribeSecurityGroups
	// only) One or more prefix list IDs for an AWS service. In an AuthorizeSecurityGroupEgress
	// request, this is the AWS service that you want to access through a VPC endpoint
	// from instances associated with the security group.
	PrefixListIds []PrefixListId `locationName:"prefixListIds" locationNameList:"item" type:"list"`

	// The end of port range for the TCP and UDP protocols, or an ICMP code. A value
	// of -1 indicates all ICMP codes for the specified ICMP type.
	ToPort *int64 `locationName:"toPort" type:"integer"`

	// One or more security group and AWS account ID pairs.
	UserIdGroupPairs []UserIdGroupPair `locationName:"groups" locationNameList:"item" type:"list"`
}
