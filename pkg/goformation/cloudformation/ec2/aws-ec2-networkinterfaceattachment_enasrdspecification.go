package ec2

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// NetworkInterfaceAttachment_EnaSrdSpecification AWS CloudFormation Resource (AWS::EC2::NetworkInterfaceAttachment.EnaSrdSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterfaceattachment-enasrdspecification.html
type NetworkInterfaceAttachment_EnaSrdSpecification struct {

	// EnaSrdEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterfaceattachment-enasrdspecification.html#cfn-ec2-networkinterfaceattachment-enasrdspecification-enasrdenabled
	EnaSrdEnabled *types.Value `json:"EnaSrdEnabled,omitempty"`

	// EnaSrdUdpSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterfaceattachment-enasrdspecification.html#cfn-ec2-networkinterfaceattachment-enasrdspecification-enasrdudpspecification
	EnaSrdUdpSpecification *NetworkInterfaceAttachment_EnaSrdUdpSpecification `json:"EnaSrdUdpSpecification,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *NetworkInterfaceAttachment_EnaSrdSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInterfaceAttachment.EnaSrdSpecification"
}
