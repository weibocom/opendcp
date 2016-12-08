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
