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

type ImagesResp struct {
	Images []Image `locationName:"imagesSet" locationNameList:"item" type:"list"`
}

// Describes an image.
type Image struct {

	// The architecture of the image.
	Architecture string `locationName:"architecture" type:"string" enum:"ArchitectureValues"`

	// Any block device mapping entries.
	BlockDeviceMappings []BlockDeviceMapping `locationName:"blockDeviceMapping" locationNameList:"item" type:"list"`

	// The date and time the image was created.
	CreationDate string `locationName:"creationDate" type:"string"`

	// The description of the AMI that was provided during image creation.
	Description string `locationName:"description" type:"string"`

	// Specifies whether enhanced networking with ENA is enabled.
	EnaSupport bool `locationName:"enaSupport" type:"boolean"`

	// The hypervisor type of the image.
	Hypervisor string `locationName:"hypervisor" type:"string" enum:"HypervisorType"`

	// The ID of the AMI.
	ImageId string `locationName:"imageId" type:"string"`

	// The location of the AMI.
	ImageLocation string `locationName:"imageLocation" type:"string"`

	// The AWS account alias (for example, amazon, self) or the AWS account ID of
	// the AMI owner.
	ImageOwnerAlias string `locationName:"imageOwnerAlias" type:"string"`

	// The type of image.
	ImageType string `locationName:"imageType" type:"string" enum:"ImageTypeValues"`

	// The kernel associated with the image, if any. Only applicable for machine
	// images.
	KernelId string `locationName:"kernelId" type:"string"`

	// The name of the AMI that was provided during image creation.
	Name string `locationName:"name" type:"string"`

	// The AWS account ID of the image owner.
	OwnerId string `locationName:"imageOwnerId" type:"string"`

	// The value is Windows for Windows AMIs; otherwise blank.
	Platform string `locationName:"platform" type:"string" enum:"PlatformValues"`

	// Any product codes associated with the AMI.
	ProductCodes []ProductCode `locationName:"productCodes" locationNameList:"item" type:"list"`

	// Indicates whether the image has public launch permissions. The value is true
	// if this image has public launch permissions or false if it has only implicit
	// and explicit launch permissions.
	Public bool `locationName:"isPublic" type:"boolean"`

	// The RAM disk associated with the image, if any. Only applicable for machine
	// images.
	RamdiskId string `locationName:"ramdiskId" type:"string"`

	// The device name of the root device (for example, /dev/sda1 or /dev/xvda).
	RootDeviceName string `locationName:"rootDeviceName" type:"string"`

	// The type of root device used by the AMI. The AMI can use an EBS volume or
	// an instance store volume.
	RootDeviceType string `locationName:"rootDeviceType" type:"string" enum:"DeviceType"`

	// Specifies whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport string `locationName:"sriovNetSupport" type:"string"`

	// The current state of the AMI. If the state is available, the image is successfully
	// registered and can be used to launch an instance.
	State string `locationName:"imageState" type:"string" enum:"ImageState"`

	// The reason for the state change.
	StateReason StateReason `locationName:"stateReason" type:"structure"`

	// Any tags assigned to the image.
	Tags []Tag `locationName:"tagSet" locationNameList:"item" type:"list"`

	// The type of virtualization of the AMI.
	VirtualizationType string `locationName:"virtualizationType" type:"string" enum:"VirtualizationType"`
}

// Describes a block device mapping.
type BlockDeviceMapping struct {

	// The device name exposed to the instance (for example, /dev/sdh or xvdh).
	DeviceName string `locationName:"deviceName" type:"string"`

	// Parameters used to automatically set up EBS volumes when the instance is
	// launched.
	Ebs EbsBlockDevice `locationName:"ebs" type:"structure"`

	// Suppresses the specified device included in the block device mapping of the
	// AMI.
	NoDevice string `locationName:"noDevice" type:"string"`

	// The virtual device name (ephemeralN). Instance store volumes are numbered
	// starting from 0. An instance type with 2 available instance store volumes
	// can specify mappings for ephemeral0 and ephemeral1.The number of available
	// instance store volumes depends on the instance type. After you connect to
	// the instance, you must mount the volume.
	//
	// Constraints: For M3 instances, you must specify instance store volumes in
	// the block device mapping for the instance. When you launch an M3 instance,
	// we ignore any instance store volumes specified in the block device mapping
	// for the AMI.
	VirtualName string `locationName:"virtualName" type:"string"`
}

// Describes a block device for an EBS volume.
type EbsBlockDevice struct {

	// Indicates whether the EBS volume is deleted on instance termination.
	DeleteOnTermination bool `locationName:"deleteOnTermination" type:"boolean"`

	// Indicates whether the EBS volume is encrypted. Encrypted Amazon EBS volumes
	// may only be attached to instances that support Amazon EBS encryption.
	Encrypted bool `locationName:"encrypted" type:"boolean"`

	// The number of I/O operations per second (IOPS) that the volume supports.
	// For io1, this represents the number of IOPS that are provisioned for the
	// volume. For gp2, this represents the baseline performance of the volume and
	// the rate at which the volume accumulates I/O credits for bursting. For more
	// information about General Purpose SSD baseline performance, I/O credits,
	// and bursting, see Amazon EBS Volume Types (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Constraint: Range is 100-20000 IOPS for io1 volumes and 100-10000 IOPS for
	// gp2 volumes.
	//
	// Condition: This parameter is required for requests to create io1 volumes;
	// it is not used in requests to create gp2, st1, sc1, or standard volumes.
	Iops int64 `locationName:"iops" type:"integer"`

	// The ID of the snapshot.
	SnapshotId string `locationName:"snapshotId" type:"string"`

	// The size of the volume, in GiB.
	//
	// Constraints: 1-16384 for General Purpose SSD (gp2), 4-16384 for Provisioned
	// IOPS SSD (io1), 500-16384 for Throughput Optimized HDD (st1), 500-16384 for
	// Cold HDD (sc1), and 1-1024 for Magnetic (standard) volumes. If you specify
	// a snapshot, the volume size must be equal to or larger than the snapshot
	// size.
	//
	// Default: If you're creating the volume from a snapshot and don't specify
	// a volume size, the default is the snapshot size.
	VolumeSize int64 `locationName:"volumeSize" type:"integer"`

	// The volume type: gp2, io1, st1, sc1, or standard.
	//
	// Default: standard
	VolumeType string `locationName:"volumeType" type:"string" enum:"VolumeType"`
}
