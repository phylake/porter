/*
 * (c) 2016-2017 Adobe. All rights reserved.
 * This file is licensed to you under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
 * OF ANY KIND, either express or implied. See the License for the specific language
 * governing permissions and limitations under the License.
 */
package cfn

const (
	AutoScaling_AutoScalingGroup           = "AWS::AutoScaling::AutoScalingGroup"
	AutoScaling_LaunchConfiguration        = "AWS::AutoScaling::LaunchConfiguration"
	AutoScaling_LifecycleHook              = "AWS::AutoScaling::LifecycleHook"
	AutoScaling_ScalingPolicy              = "AWS::AutoScaling::ScalingPolicy"
	AutoScaling_ScheduledAction            = "AWS::AutoScaling::ScheduledAction"
	CloudFormation_Authentication          = "AWS::CloudFormation::Authentication"
	CloudFormation_CustomResource          = "AWS::CloudFormation::CustomResource"
	CloudFormation_Init                    = "AWS::CloudFormation::Init"
	CloudFormation_Interface               = "AWS::CloudFormation::Interface"
	CloudFormation_Stack                   = "AWS::CloudFormation::Stack"
	CloudFormation_WaitCondition           = "AWS::CloudFormation::WaitCondition"
	CloudFormation_WaitConditionHandle     = "AWS::CloudFormation::WaitConditionHandle"
	CloudFront_Distribution                = "AWS::CloudFront::Distribution"
	CloudTrail_Trail                       = "AWS::CloudTrail::Trail"
	CloudWatch_Alarm                       = "AWS::CloudWatch::Alarm"
	CodeDeploy_Application                 = "AWS::CodeDeploy::Application"
	CodeDeploy_DeploymentConfig            = "AWS::CodeDeploy::DeploymentConfig"
	CodeDeploy_DeploymentGroup             = "AWS::CodeDeploy::DeploymentGroup"
	CodePipeline_CustomActionType          = "AWS::CodePipeline::CustomActionType"
	CodePipeline_Pipeline                  = "AWS::CodePipeline::Pipeline"
	Config_ConfigRule                      = "AWS::Config::ConfigRule"
	Config_ConfigurationRecorder           = "AWS::Config::ConfigurationRecorder"
	Config_DeliveryChannel                 = "AWS::Config::DeliveryChannel"
	DataPipeline_Pipeline                  = "AWS::DataPipeline::Pipeline"
	DirectoryService_MicrosoftAD           = "AWS::DirectoryService::MicrosoftAD"
	DirectoryService_SimpleAD              = "AWS::DirectoryService::SimpleAD"
	DynamoDB_Table                         = "AWS::DynamoDB::Table"
	EC2_CustomerGateway                    = "AWS::EC2::CustomerGateway"
	EC2_DHCPOptions                        = "AWS::EC2::DHCPOptions"
	EC2_EIP                                = "AWS::EC2::EIP"
	EC2_EIPAssociation                     = "AWS::EC2::EIPAssociation"
	EC2_Instance                           = "AWS::EC2::Instance"
	EC2_InternetGateway                    = "AWS::EC2::InternetGateway"
	EC2_NetworkAcl                         = "AWS::EC2::NetworkAcl"
	EC2_NetworkAclEntry                    = "AWS::EC2::NetworkAclEntry"
	EC2_NetworkInterface                   = "AWS::EC2::NetworkInterface"
	EC2_NetworkInterfaceAttachment         = "AWS::EC2::NetworkInterfaceAttachment"
	EC2_PlacementGroup                     = "AWS::EC2::PlacementGroup"
	EC2_Route                              = "AWS::EC2::Route"
	EC2_RouteTable                         = "AWS::EC2::RouteTable"
	EC2_SecurityGroup                      = "AWS::EC2::SecurityGroup"
	EC2_SecurityGroupEgress                = "AWS::EC2::SecurityGroupEgress"
	EC2_SecurityGroupIngress               = "AWS::EC2::SecurityGroupIngress"
	EC2_SpotFleet                          = "AWS::EC2::SpotFleet"
	EC2_Subnet                             = "AWS::EC2::Subnet"
	EC2_SubnetNetworkAclAssociation        = "AWS::EC2::SubnetNetworkAclAssociation"
	EC2_SubnetRouteTableAssociation        = "AWS::EC2::SubnetRouteTableAssociation"
	EC2_Volume                             = "AWS::EC2::Volume"
	EC2_VolumeAttachment                   = "AWS::EC2::VolumeAttachment"
	EC2_VPC                                = "AWS::EC2::VPC"
	EC2_VPCDHCPOptionsAssociation          = "AWS::EC2::VPCDHCPOptionsAssociation"
	EC2_VPCEndpoint                        = "AWS::EC2::VPCEndpoint"
	EC2_VPCGatewayAttachment               = "AWS::EC2::VPCGatewayAttachment"
	EC2_VPCPeeringConnection               = "AWS::EC2::VPCPeeringConnection"
	EC2_VPNConnection                      = "AWS::EC2::VPNConnection"
	EC2_VPNConnectionRoute                 = "AWS::EC2::VPNConnectionRoute"
	EC2_VPNGateway                         = "AWS::EC2::VPNGateway"
	EC2_VPNGatewayRoutePropagation         = "AWS::EC2::VPNGatewayRoutePropagation"
	ECS_Cluster                            = "AWS::ECS::Cluster"
	ECS_Service                            = "AWS::ECS::Service"
	ECS_TaskDefinition                     = "AWS::ECS::TaskDefinition"
	EFS_FileSystem                         = "AWS::EFS::FileSystem"
	EFS_MountTarget                        = "AWS::EFS::MountTarget"
	ElastiCache_CacheCluster               = "AWS::ElastiCache::CacheCluster"
	ElastiCache_ParameterGroup             = "AWS::ElastiCache::ParameterGroup"
	ElastiCache_ReplicationGroup           = "AWS::ElastiCache::ReplicationGroup"
	ElastiCache_SecurityGroup              = "AWS::ElastiCache::SecurityGroup"
	ElastiCache_SecurityGroupIngress       = "AWS::ElastiCache::SecurityGroupIngress"
	ElastiCache_SubnetGroup                = "AWS::ElastiCache::SubnetGroup"
	ElasticBeanstalk_Application           = "AWS::ElasticBeanstalk::Application"
	ElasticBeanstalk_ApplicationVersion    = "AWS::ElasticBeanstalk::ApplicationVersion"
	ElasticBeanstalk_ConfigurationTemplate = "AWS::ElasticBeanstalk::ConfigurationTemplate"
	ElasticBeanstalk_Environment           = "AWS::ElasticBeanstalk::Environment"
	ElasticLoadBalancing_LoadBalancer      = "AWS::ElasticLoadBalancing::LoadBalancer"
	IAM_AccessKey                          = "AWS::IAM::AccessKey"
	IAM_Group                              = "AWS::IAM::Group"
	IAM_InstanceProfile                    = "AWS::IAM::InstanceProfile"
	IAM_ManagedPolicy                      = "AWS::IAM::ManagedPolicy"
	IAM_Policy                             = "AWS::IAM::Policy"
	IAM_Role                               = "AWS::IAM::Role"
	IAM_User                               = "AWS::IAM::User"
	IAM_UserToGroupAddition                = "AWS::IAM::UserToGroupAddition"
	Kinesis_Stream                         = "AWS::Kinesis::Stream"
	KMS_Key                                = "AWS::KMS::Key"
	Lambda_EventSourceMapping              = "AWS::Lambda::EventSourceMapping"
	Lambda_Function                        = "AWS::Lambda::Function"
	Lambda_Permission                      = "AWS::Lambda::Permission"
	Logs_Destination                       = "AWS::Logs::Destination"
	Logs_LogGroup                          = "AWS::Logs::LogGroup"
	Logs_LogStream                         = "AWS::Logs::LogStream"
	Logs_MetricFilter                      = "AWS::Logs::MetricFilter"
	Logs_SubscriptionFilter                = "AWS::Logs::SubscriptionFilter"
	OpsWorks_App                           = "AWS::OpsWorks::App"
	OpsWorks_ElasticLoadBalancerAttachment = "AWS::OpsWorks::ElasticLoadBalancerAttachment"
	OpsWorks_Instance                      = "AWS::OpsWorks::Instance"
	OpsWorks_Layer                         = "AWS::OpsWorks::Layer"
	OpsWorks_Stack                         = "AWS::OpsWorks::Stack"
	RDS_DBCluster                          = "AWS::RDS::DBCluster"
	RDS_DBClusterParameterGroup            = "AWS::RDS::DBClusterParameterGroup"
	RDS_DBInstance                         = "AWS::RDS::DBInstance"
	RDS_DBParameterGroup                   = "AWS::RDS::DBParameterGroup"
	RDS_DBSecurityGroup                    = "AWS::RDS::DBSecurityGroup"
	RDS_DBSecurityGroupIngress             = "AWS::RDS::DBSecurityGroupIngress"
	RDS_DBSubnetGroup                      = "AWS::RDS::DBSubnetGroup"
	RDS_EventSubscription                  = "AWS::RDS::EventSubscription"
	RDS_OptionGroup                        = "AWS::RDS::OptionGroup"
	Redshift_Cluster                       = "AWS::Redshift::Cluster"
	Redshift_ClusterParameterGroup         = "AWS::Redshift::ClusterParameterGroup"
	Redshift_ClusterSecurityGroup          = "AWS::Redshift::ClusterSecurityGroup"
	Redshift_ClusterSecurityGroupIngress   = "AWS::Redshift::ClusterSecurityGroupIngress"
	Redshift_ClusterSubnetGroup            = "AWS::Redshift::ClusterSubnetGroup"
	Route53_HealthCheck                    = "AWS::Route53::HealthCheck"
	Route53_HostedZone                     = "AWS::Route53::HostedZone"
	Route53_RecordSet                      = "AWS::Route53::RecordSet"
	Route53_RecordSetGroup                 = "AWS::Route53::RecordSetGroup"
	S3_Bucket                              = "AWS::S3::Bucket"
	S3_BucketPolicy                        = "AWS::S3::BucketPolicy"
	SDB_Domain                             = "AWS::SDB::Domain"
	SNS_Topic                              = "AWS::SNS::Topic"
	SNS_TopicPolicy                        = "AWS::SNS::TopicPolicy"
	SQS_Queue                              = "AWS::SQS::Queue"
	SQS_QueuePolicy                        = "AWS::SQS::QueuePolicy"
	SSM_Document                           = "AWS::SSM::Document"
	WAF_ByteMatchSet                       = "AWS::WAF::ByteMatchSet"
	WAF_IPSet                              = "AWS::WAF::IPSet"
	WAF_Rule                               = "AWS::WAF::Rule"
	WAF_SqlInjectionMatchSet               = "AWS::WAF::SqlInjectionMatchSet"
	WAF_WebACL                             = "AWS::WAF::WebACL"
	WorkSpaces_Workspace                   = "AWS::WorkSpaces::Workspace"
)

var allTypes map[string]interface{}

func ValidType(resourceType string) bool {
	_, exists := allTypes[resourceType]
	return exists
}

func init() {
	allTypes = make(map[string]interface{})
	allTypes[AutoScaling_AutoScalingGroup] = nil
	allTypes[AutoScaling_LaunchConfiguration] = nil
	allTypes[AutoScaling_LifecycleHook] = nil
	allTypes[AutoScaling_ScalingPolicy] = nil
	allTypes[AutoScaling_ScheduledAction] = nil
	allTypes[CloudFormation_Authentication] = nil
	allTypes[CloudFormation_CustomResource] = nil
	allTypes[CloudFormation_Init] = nil
	allTypes[CloudFormation_Interface] = nil
	allTypes[CloudFormation_Stack] = nil
	allTypes[CloudFormation_WaitCondition] = nil
	allTypes[CloudFormation_WaitConditionHandle] = nil
	allTypes[CloudFront_Distribution] = nil
	allTypes[CloudTrail_Trail] = nil
	allTypes[CloudWatch_Alarm] = nil
	allTypes[CodeDeploy_Application] = nil
	allTypes[CodeDeploy_DeploymentConfig] = nil
	allTypes[CodeDeploy_DeploymentGroup] = nil
	allTypes[CodePipeline_CustomActionType] = nil
	allTypes[CodePipeline_Pipeline] = nil
	allTypes[Config_ConfigRule] = nil
	allTypes[Config_ConfigurationRecorder] = nil
	allTypes[Config_DeliveryChannel] = nil
	allTypes[DataPipeline_Pipeline] = nil
	allTypes[DirectoryService_MicrosoftAD] = nil
	allTypes[DirectoryService_SimpleAD] = nil
	allTypes[DynamoDB_Table] = nil
	allTypes[EC2_CustomerGateway] = nil
	allTypes[EC2_DHCPOptions] = nil
	allTypes[EC2_EIP] = nil
	allTypes[EC2_EIPAssociation] = nil
	allTypes[EC2_Instance] = nil
	allTypes[EC2_InternetGateway] = nil
	allTypes[EC2_NetworkAcl] = nil
	allTypes[EC2_NetworkAclEntry] = nil
	allTypes[EC2_NetworkInterface] = nil
	allTypes[EC2_NetworkInterfaceAttachment] = nil
	allTypes[EC2_PlacementGroup] = nil
	allTypes[EC2_Route] = nil
	allTypes[EC2_RouteTable] = nil
	allTypes[EC2_SecurityGroup] = nil
	allTypes[EC2_SecurityGroupEgress] = nil
	allTypes[EC2_SecurityGroupIngress] = nil
	allTypes[EC2_SpotFleet] = nil
	allTypes[EC2_Subnet] = nil
	allTypes[EC2_SubnetNetworkAclAssociation] = nil
	allTypes[EC2_SubnetRouteTableAssociation] = nil
	allTypes[EC2_Volume] = nil
	allTypes[EC2_VolumeAttachment] = nil
	allTypes[EC2_VPC] = nil
	allTypes[EC2_VPCDHCPOptionsAssociation] = nil
	allTypes[EC2_VPCEndpoint] = nil
	allTypes[EC2_VPCGatewayAttachment] = nil
	allTypes[EC2_VPCPeeringConnection] = nil
	allTypes[EC2_VPNConnection] = nil
	allTypes[EC2_VPNConnectionRoute] = nil
	allTypes[EC2_VPNGateway] = nil
	allTypes[EC2_VPNGatewayRoutePropagation] = nil
	allTypes[ECS_Cluster] = nil
	allTypes[ECS_Service] = nil
	allTypes[ECS_TaskDefinition] = nil
	allTypes[EFS_FileSystem] = nil
	allTypes[EFS_MountTarget] = nil
	allTypes[ElastiCache_CacheCluster] = nil
	allTypes[ElastiCache_ParameterGroup] = nil
	allTypes[ElastiCache_ReplicationGroup] = nil
	allTypes[ElastiCache_SecurityGroup] = nil
	allTypes[ElastiCache_SecurityGroupIngress] = nil
	allTypes[ElastiCache_SubnetGroup] = nil
	allTypes[ElasticBeanstalk_Application] = nil
	allTypes[ElasticBeanstalk_ApplicationVersion] = nil
	allTypes[ElasticBeanstalk_ConfigurationTemplate] = nil
	allTypes[ElasticBeanstalk_Environment] = nil
	allTypes[ElasticLoadBalancing_LoadBalancer] = nil
	allTypes[IAM_AccessKey] = nil
	allTypes[IAM_Group] = nil
	allTypes[IAM_InstanceProfile] = nil
	allTypes[IAM_ManagedPolicy] = nil
	allTypes[IAM_Policy] = nil
	allTypes[IAM_Role] = nil
	allTypes[IAM_User] = nil
	allTypes[IAM_UserToGroupAddition] = nil
	allTypes[Kinesis_Stream] = nil
	allTypes[KMS_Key] = nil
	allTypes[Lambda_EventSourceMapping] = nil
	allTypes[Lambda_Function] = nil
	allTypes[Lambda_Permission] = nil
	allTypes[Logs_Destination] = nil
	allTypes[Logs_LogGroup] = nil
	allTypes[Logs_LogStream] = nil
	allTypes[Logs_MetricFilter] = nil
	allTypes[Logs_SubscriptionFilter] = nil
	allTypes[OpsWorks_App] = nil
	allTypes[OpsWorks_ElasticLoadBalancerAttachment] = nil
	allTypes[OpsWorks_Instance] = nil
	allTypes[OpsWorks_Layer] = nil
	allTypes[OpsWorks_Stack] = nil
	allTypes[RDS_DBCluster] = nil
	allTypes[RDS_DBClusterParameterGroup] = nil
	allTypes[RDS_DBInstance] = nil
	allTypes[RDS_DBParameterGroup] = nil
	allTypes[RDS_DBSecurityGroup] = nil
	allTypes[RDS_DBSecurityGroupIngress] = nil
	allTypes[RDS_DBSubnetGroup] = nil
	allTypes[RDS_EventSubscription] = nil
	allTypes[RDS_OptionGroup] = nil
	allTypes[Redshift_Cluster] = nil
	allTypes[Redshift_ClusterParameterGroup] = nil
	allTypes[Redshift_ClusterSecurityGroup] = nil
	allTypes[Redshift_ClusterSecurityGroupIngress] = nil
	allTypes[Redshift_ClusterSubnetGroup] = nil
	allTypes[Route53_HealthCheck] = nil
	allTypes[Route53_HostedZone] = nil
	allTypes[Route53_RecordSet] = nil
	allTypes[Route53_RecordSetGroup] = nil
	allTypes[S3_Bucket] = nil
	allTypes[S3_BucketPolicy] = nil
	allTypes[SDB_Domain] = nil
	allTypes[SNS_Topic] = nil
	allTypes[SNS_TopicPolicy] = nil
	allTypes[SQS_Queue] = nil
	allTypes[SQS_QueuePolicy] = nil
	allTypes[SSM_Document] = nil
	allTypes[WAF_ByteMatchSet] = nil
	allTypes[WAF_IPSet] = nil
	allTypes[WAF_Rule] = nil
	allTypes[WAF_SqlInjectionMatchSet] = nil
	allTypes[WAF_WebACL] = nil
	allTypes[WorkSpaces_Workspace] = nil
}
