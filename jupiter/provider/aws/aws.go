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



package aws

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/provider"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"weibo.com/opendcp/jupiter/conf"
)

func init() {
	provider.RegisterProviderDriver("aws", new)
}

type awsProvider struct {
	client *ec2.EC2
	lock   sync.Mutex
}

var instanceTypesInAws = map[string]string{
	"1Core-0.5Gib":  "t2.nano",
	"1Core-1Gib":	 "t2.micro",
	"1Core-2GiB":    "t2.small",
	"2Cores-8GiB":   "t2.large",
	"8Cores-15GiB":  "c4.2xlarge",
	"16Cores-30GiB": "c4.4xlarge",
}

var dataCategoryInAws = []string{
	"gp2",
	"io1",
	"st1",
	"sc1",
	"standard",
}

var internetChargeTypeInAws = []string{}

func (driver awsProvider) GetInstanceType(key string) string {
	return instanceTypesInAws[key]
}

func (driver awsProvider) ListDiskCategory() []string {
	return dataCategoryInAws
}

func (driver awsProvider) ListInternetChargeType() []string {
	return internetChargeTypeInAws
}

func (driver awsProvider) Create(input *models.Cluster, number int) ([]string, []error) {
	driver.client.Config.Credentials = credentials.NewStaticCredentials(conf.Config.AwsKeyId, conf.Config.AwsKeySecret, "")

	runResult, err := driver.client.RunInstances(&ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for imageId (such as t2.micro instances) in the cn-north-1 region
		ImageId:      aws.String(input.ImageId),
		InstanceType: aws.String(input.InstanceType),
		MinCount:     aws.Int64(int64(number)),
		MaxCount:     aws.Int64(int64(number)),
		KeyName:      aws.String(input.KeyName),
		Monitoring: &ec2.RunInstancesMonitoringEnabled{
			Enabled: aws.Bool(true),
		},
		//SecurityGroupIds: []*string{
		//	aws.String(input.Network.SecurityGroup),
		//},
		SubnetId: aws.String(input.Network.SubnetId),
		//BlockDeviceMappings: []*ec2.BlockDeviceMapping{
		//	{
		//		DeviceName: aws.String("/dev/sdh"),
		//		Ebs: &ec2.EbsBlockDevice{
		//			VolumeSize: aws.Int64(int64(input.DataDiskSize)),
		//			VolumeType: aws.String(input.DataDiskCategory),
		//		},
		//	},
		//},
	})
	if err != nil {
		beego.Error("Could not create instance", err)
		return nil, []error{err}
	}
	var instanceIds []string

	time.Sleep(180 * time.Second)

	for i := 0; i < len(runResult.Instances); i++ {
		beego.Debug("Created instance", *runResult.Instances[i].InstanceId)
		instanceIds = append(instanceIds, *(runResult.Instances[i].InstanceId))

		allocRes, err := driver.client.AllocateAddress(&ec2.AllocateAddressInput{
			Domain: aws.String("vpc"),
		})
		if allocRes.PublicIp == nil {
			beego.Error("Unable to allow prublic IP", err)
		}

		if err != nil {
			beego.Error("Unable to allocate IP address, %v", err)
		}

		_, errAssociate := driver.client.AssociateAddress(&ec2.AssociateAddressInput{
			AllocationId: allocRes.AllocationId,
			InstanceId:   runResult.Instances[i].InstanceId,
		})
		if errAssociate != nil {
			beego.Error("Unable to associate IP address with %s, %v",
				runResult.Instances[i].InstanceId, err)
		}
	}
	return instanceIds, nil
}

func (driver awsProvider) Delete(instanceId string) (time.Time, error) {
	var instanceIds []*string
	instanceIds = append(instanceIds, aws.String(instanceId))
	params := &ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
		DryRun:      aws.Bool(false),
	}
	_, err := driver.client.TerminateInstances(params)
	if err != nil {
		beego.Error(err.Error())
		return time.Now(), err
	}
	return time.Now(), nil
}

func (driver awsProvider) Start(instanceId string) (bool, error) {
	params := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
		DryRun: aws.Bool(false),
	}
	_, err := driver.client.StartInstances(params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (driver awsProvider) Stop(instanceId string) (bool, error) {
	params := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
		DryRun: aws.Bool(false),
		Force:  aws.Bool(true),
	}
	_, err := driver.client.StopInstances(params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (driver awsProvider) List(regionId string, pageNumber int, pageSize int) (*models.ListInstancesResponse, error) {
	//params := &ec2.DescribeInstancesInput{
	//	DryRun: aws.Bool(false),
	//	Filters: []*ec2.Filter{
	//		{
	//			Name: aws.String("instance-state-code"),
	//			Values: []*string{
	//				aws.String("16"),
	//			},
	//		},
	//	},
	//}
	//listResult, err := driver.client.DescribeInstances(params)
	//if err != nil {
	//	beego.Error(err.Error())
	//	return nil, err
	//}
	//var resp models.ListInstancesResp
	//respJson, err := json.Marshal(listResult)
	//if err != nil {
	//	beego.Error(err.Error())
	//	return nil, err
	//}
	//err = json.Unmarshal(respJson, &resp)
	//if err != nil {
	//	beego.Error(err.Error())
	//	return nil, err
	//}
	//beego.Info(resp)
	//return nil, nil
	return nil, nil
}

func (driver awsProvider) ListInstanceTypes() ([]string, error) {
	instanceTypes := make([]string, 0, len(instanceTypesInAws))
	for instanceType := range instanceTypesInAws {
		instanceTypes = append(instanceTypes, instanceType)
	}
	return instanceTypes, nil
}

func (driver awsProvider) GetInstance(instanceId string) (*models.Instance, error) {

	resp, err := driver.client.DescribeInstances(&ec2.DescribeInstancesInput{
		DryRun: aws.Bool(false),
		InstanceIds: [] *string {&instanceId},
	})
	if err != nil {
		beego.Error("Unable to describe instance", err)
		return nil, err
	}

	res := resp.Reservations[0].Instances[0]

	var instance models.Instance
	instance.InstanceId = *res.InstanceId
	instance.Provider = "aws"
	instance.CreateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", res.LaunchTime.String(), time.Local)
	instance.ImageId = *res.ImageId
	instance.InstanceType = *res.InstanceType
	instance.VpcId = *res.VpcId
	instance.SubnetId = *res.SubnetId
	instance.SecurityGroupId = *res.SecurityGroups[0].GroupId
	instance.PrivateIpAddress = *res.PrivateIpAddress
	instance.PublicIpAddress = *res.PublicIpAddress
	instance.RegionId = "aws region "
	instance.ZoneId = "aws zone"
	instance.NatIpAddress = "aws nia"
	instance.CostWay = "aws costway"

	return &instance, err
}

func (driver awsProvider) CreateSecurityGroup(input *models.CreateSecurityGroupParam) (*models.SecurityGroupResp, error) {
	params := &ec2.CreateSecurityGroupInput{
		Description: aws.String(input.Description),
		GroupName:   aws.String(input.GroupName),
		DryRun:      aws.Bool(false),
		VpcId:       aws.String(input.VpcId),
	}
	result, err := driver.client.CreateSecurityGroup(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.SecurityGroupResp
	respJson, err := json.Marshal(result)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) ListSecurityGroup(regionId string, vpcId string) (*models.SecurityGroupsResp, error) {
	params := &ec2.DescribeSecurityGroupsInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String(vpcId),
				},
			},
		},
	}
	listResult, err := driver.client.DescribeSecurityGroups(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.SecurityGroupsResp
	respJson, err := json.Marshal(listResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) AuthorizeSecurityGroupIngress(input *models.AuthorizeSecurityGroupIngress) (bool, error) {
	params := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:    aws.String(input.GroupId),
		IpProtocol: aws.String(input.IpProtocol),
		FromPort:   aws.Int64(input.FromPort),
		ToPort:     aws.Int64(input.ToPort),
		CidrIp:     aws.String(input.CidrIp),
	}
	_, err := driver.client.AuthorizeSecurityGroupIngress(params)
	if err != nil {
		beego.Error(err.Error())
		return false, err
	}
	return true, nil
}

func (driver awsProvider) ListAvailabilityZones(regionId string) (*models.AvailabilityZonesResp, error) {
	params := &ec2.DescribeAvailabilityZonesInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name: aws.String("state"),
				Values: []*string{
					aws.String("available"),
				},
			},
			{
				Name: aws.String("region-name"),
				Values: []*string{
					aws.String(regionId),
				},
			},
		},
	}
	listResult, err := driver.client.DescribeAvailabilityZones(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.AvailabilityZonesResp
	respJson, err := json.Marshal(listResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) ListRegions() (*models.RegionsResp, error) {
	params := &ec2.DescribeRegionsInput{
		DryRun: aws.Bool(false),
	}
	listResult, err := driver.client.DescribeRegions(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.RegionsResp
	respJson, err := json.Marshal(listResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) ListVpcs(regionId string, pageNumber int, pageSize int) (*models.VpcsResp, error) {
	params := &ec2.DescribeVpcsInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name: aws.String("state"),
				Values: []*string{
					aws.String("available"),
				},
			},
		},
	}
	listResult, err := driver.client.DescribeVpcs(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.VpcsResp
	respJson, err := json.Marshal(listResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) ListImages(regionId string, snapshotId string, pageSize int, pageNumber int) (*models.ImagesResp, error) {
	params := &ec2.DescribeImagesInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name: aws.String("state"),
				Values: []*string{
					aws.String("available"),
				},
			},
		},
	}
	listResult, err := driver.client.DescribeImages(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.ImagesResp
	respJson, err := json.Marshal(listResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) ListSubnets(zoneId string, vpcId string) (*models.SubnetsResp, error) {
	params := &ec2.DescribeSubnetsInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name: aws.String("state"),
				Values: []*string{
					aws.String("available"),
				},
			},
		},
	}
	listResult, err := driver.client.DescribeSubnets(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.SubnetsResp
	respJson, err := json.Marshal(listResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) CreateVpc(input *models.CreateVpcParam) (*models.VpcResp, error) {
	params := &ec2.CreateVpcInput{
		CidrBlock: aws.String(input.CidrBlock),
	}
	createResult, err := driver.client.CreateVpc(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.VpcResp
	respJson, err := json.Marshal(createResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) CreateGateway() (*models.GatewayResp, error) {
	params := &ec2.CreateInternetGatewayInput{
		DryRun: aws.Bool(false),
	}
	createResult, err := driver.client.CreateInternetGateway(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.GatewayResp
	respJson, err := json.Marshal(createResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) CreateSubnet(input *models.CreateSubnetParam) (*models.SubnetResp, error) {
	params := &ec2.CreateSubnetInput{
		CidrBlock: aws.String(input.CidrBlock),
		VpcId:     aws.String(input.VpcId),
		DryRun:    aws.Bool(false),
	}
	createResult, err := driver.client.CreateSubnet(params)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	var resp models.SubnetResp
	respJson, err := json.Marshal(createResult)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	err = json.Unmarshal(respJson, &resp)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	return &resp, nil
}

func (driver awsProvider) AttachGateway(input *models.AttachGateway) (bool, error) {
	params := &ec2.AttachInternetGatewayInput{
		InternetGatewayId: aws.String(input.InternetGatewayId),
		VpcId:             aws.String(input.VpcId),
		DryRun:            aws.Bool(false),
	}
	_, err := driver.client.AttachInternetGateway(params)
	if err != nil {
		beego.Error(err.Error())
		return false, err
	}
	return true, nil
}


func (driver awsProvider) AllocatePublicIpAddress(instanceId string) (string, error) {
	input := &ec2.DescribeAddressesInput{}
	result, err := driver.client.DescribeAddresses(input)
	if err != nil {
		beego.Error("Fail to get public Ip", err)
	}

	return *result.Addresses[0].PrivateIpAddress, nil
}

func (driver awsProvider) WaitForInstanceToStop(instanceId string) bool {
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
		DryRun: aws.Bool(false),
	}

	err := driver.client.WaitUntilInstanceStopped(input)
	if err != nil {
		beego.Error("the wait instance err:", err)
		return false
	}

	return true
	}

func (driver awsProvider) WaitToStartInstance(instanceId string) bool {
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
		DryRun: aws.Bool(false),
	}

	err := driver.client.WaitUntilInstanceRunning(input)
	if err != nil {
		beego.Error("the wait instance err:", err)
		return false
	}

	return true
}

func new() (provider.ProviderDriver, error) {
	return  newProvider()
}

func newProvider() (provider.ProviderDriver, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("cn-north-1"),
	}))
	client := ec2.New(sess)
	client.Config.Credentials = credentials.NewStaticCredentials(conf.Config.KeyId, conf.Config.KeySecret, "")
	ret := awsProvider{
		client: client,
	}
	return ret, nil
}
