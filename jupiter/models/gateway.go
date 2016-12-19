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

// Contains the output of CreateInternetGateway.
type GatewayResp struct {

	// Information about the Internet gateway.
	InternetGateway InternetGateway `locationName:"internetGateway" type:"structure"`
}

// Describes an Internet gateway.
type InternetGateway struct {

	// Any VPCs attached to the Internet gateway.
	Attachments []InternetGatewayAttachment `locationName:"attachmentSet" locationNameList:"item" type:"list"`

	// The ID of the Internet gateway.
	InternetGatewayId string `locationName:"internetGatewayId" type:"string"`

	// Any tags assigned to the Internet gateway.
	Tags []Tag `locationName:"tagSet" locationNameList:"item" type:"list"`
}

// Describes the attachment of a VPC to an Internet gateway.
type InternetGatewayAttachment struct {
	// The current state of the attachment.
	State string `locationName:"state" type:"string" enum:"AttachmentStatus"`

	// The ID of the VPC.
	VpcId *string `locationName:"vpcId" type:"string"`
}

// Contains the parameters for AttachGateway.
type AttachGateway struct {

	// The name of provider
	Provider string `locationName:"provider" type:"string" required:"true"`

	// The ID of the Internet gateway.
	InternetGatewayId string `locationName:"internetGatewayId" type:"string" required:"true"`

	// The ID of the VPC.
	VpcId string `locationName:"vpcId" type:"string" required:"true"`
}
