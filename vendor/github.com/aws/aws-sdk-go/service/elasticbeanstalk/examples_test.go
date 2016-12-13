// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticbeanstalk_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleElasticBeanstalk_AbortEnvironmentUpdate() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.AbortEnvironmentUpdateInput{
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.AbortEnvironmentUpdate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_CheckDNSAvailability() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.CheckDNSAvailabilityInput{
		CNAMEPrefix: aws.String("DNSCnamePrefix"), // Required
	}
	resp, err := svc.CheckDNSAvailability(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_ComposeEnvironments() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.ComposeEnvironmentsInput{
		ApplicationName: aws.String("ApplicationName"),
		GroupName:       aws.String("GroupName"),
		VersionLabels: []*string{
			aws.String("VersionLabel"), // Required
			// More values...
		},
	}
	resp, err := svc.ComposeEnvironments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_CreateApplication() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.CreateApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		Description:     aws.String("Description"),
	}
	resp, err := svc.CreateApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_CreateApplicationVersion() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.CreateApplicationVersionInput{
		ApplicationName:       aws.String("ApplicationName"), // Required
		VersionLabel:          aws.String("VersionLabel"),    // Required
		AutoCreateApplication: aws.Bool(true),
		Description:           aws.String("Description"),
		Process:               aws.Bool(true),
		SourceBundle: &elasticbeanstalk.S3Location{
			S3Bucket: aws.String("S3Bucket"),
			S3Key:    aws.String("S3Key"),
		},
	}
	resp, err := svc.CreateApplicationVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_CreateConfigurationTemplate() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.CreateConfigurationTemplateInput{
		ApplicationName: aws.String("ApplicationName"),           // Required
		TemplateName:    aws.String("ConfigurationTemplateName"), // Required
		Description:     aws.String("Description"),
		EnvironmentId:   aws.String("EnvironmentId"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
				Value:        aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		SolutionStackName: aws.String("SolutionStackName"),
		SourceConfiguration: &elasticbeanstalk.SourceConfiguration{
			ApplicationName: aws.String("ApplicationName"),
			TemplateName:    aws.String("ConfigurationTemplateName"),
		},
	}
	resp, err := svc.CreateConfigurationTemplate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_CreateEnvironment() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.CreateEnvironmentInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		CNAMEPrefix:     aws.String("DNSCnamePrefix"),
		Description:     aws.String("Description"),
		EnvironmentName: aws.String("EnvironmentName"),
		GroupName:       aws.String("GroupName"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
				Value:        aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		OptionsToRemove: []*elasticbeanstalk.OptionSpecification{
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
			},
			// More values...
		},
		SolutionStackName: aws.String("SolutionStackName"),
		Tags: []*elasticbeanstalk.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		TemplateName: aws.String("ConfigurationTemplateName"),
		Tier: &elasticbeanstalk.EnvironmentTier{
			Name:    aws.String("String"),
			Type:    aws.String("String"),
			Version: aws.String("String"),
		},
		VersionLabel: aws.String("VersionLabel"),
	}
	resp, err := svc.CreateEnvironment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_CreateStorageLocation() {
	svc := elasticbeanstalk.New(session.New())

	var params *elasticbeanstalk.CreateStorageLocationInput
	resp, err := svc.CreateStorageLocation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DeleteApplication() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DeleteApplicationInput{
		ApplicationName:     aws.String("ApplicationName"), // Required
		TerminateEnvByForce: aws.Bool(true),
	}
	resp, err := svc.DeleteApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DeleteApplicationVersion() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DeleteApplicationVersionInput{
		ApplicationName:    aws.String("ApplicationName"), // Required
		VersionLabel:       aws.String("VersionLabel"),    // Required
		DeleteSourceBundle: aws.Bool(true),
	}
	resp, err := svc.DeleteApplicationVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DeleteConfigurationTemplate() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DeleteConfigurationTemplateInput{
		ApplicationName: aws.String("ApplicationName"),           // Required
		TemplateName:    aws.String("ConfigurationTemplateName"), // Required
	}
	resp, err := svc.DeleteConfigurationTemplate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DeleteEnvironmentConfiguration() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DeleteEnvironmentConfigurationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		EnvironmentName: aws.String("EnvironmentName"), // Required
	}
	resp, err := svc.DeleteEnvironmentConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeApplicationVersions() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeApplicationVersionsInput{
		ApplicationName: aws.String("ApplicationName"),
		VersionLabels: []*string{
			aws.String("VersionLabel"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeApplicationVersions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeApplications() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeApplicationsInput{
		ApplicationNames: []*string{
			aws.String("ApplicationName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeApplications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeConfigurationOptions() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeConfigurationOptionsInput{
		ApplicationName: aws.String("ApplicationName"),
		EnvironmentName: aws.String("EnvironmentName"),
		Options: []*elasticbeanstalk.OptionSpecification{
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
			},
			// More values...
		},
		SolutionStackName: aws.String("SolutionStackName"),
		TemplateName:      aws.String("ConfigurationTemplateName"),
	}
	resp, err := svc.DescribeConfigurationOptions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeConfigurationSettings() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeConfigurationSettingsInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		EnvironmentName: aws.String("EnvironmentName"),
		TemplateName:    aws.String("ConfigurationTemplateName"),
	}
	resp, err := svc.DescribeConfigurationSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeEnvironmentHealth() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeEnvironmentHealthInput{
		AttributeNames: []*string{
			aws.String("EnvironmentHealthAttribute"), // Required
			// More values...
		},
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.DescribeEnvironmentHealth(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeEnvironmentResources() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeEnvironmentResourcesInput{
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.DescribeEnvironmentResources(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeEnvironments() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeEnvironmentsInput{
		ApplicationName: aws.String("ApplicationName"),
		EnvironmentIds: []*string{
			aws.String("EnvironmentId"), // Required
			// More values...
		},
		EnvironmentNames: []*string{
			aws.String("EnvironmentName"), // Required
			// More values...
		},
		IncludeDeleted:        aws.Bool(true),
		IncludedDeletedBackTo: aws.Time(time.Now()),
		VersionLabel:          aws.String("VersionLabel"),
	}
	resp, err := svc.DescribeEnvironments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeEvents() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeEventsInput{
		ApplicationName: aws.String("ApplicationName"),
		EndTime:         aws.Time(time.Now()),
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
		MaxRecords:      aws.Int64(1),
		NextToken:       aws.String("Token"),
		RequestId:       aws.String("RequestId"),
		Severity:        aws.String("EventSeverity"),
		StartTime:       aws.Time(time.Now()),
		TemplateName:    aws.String("ConfigurationTemplateName"),
		VersionLabel:    aws.String("VersionLabel"),
	}
	resp, err := svc.DescribeEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_DescribeInstancesHealth() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.DescribeInstancesHealthInput{
		AttributeNames: []*string{
			aws.String("InstancesHealthAttribute"), // Required
			// More values...
		},
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
		NextToken:       aws.String("NextToken"),
	}
	resp, err := svc.DescribeInstancesHealth(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_ListAvailableSolutionStacks() {
	svc := elasticbeanstalk.New(session.New())

	var params *elasticbeanstalk.ListAvailableSolutionStacksInput
	resp, err := svc.ListAvailableSolutionStacks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_RebuildEnvironment() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.RebuildEnvironmentInput{
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.RebuildEnvironment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_RequestEnvironmentInfo() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.RequestEnvironmentInfoInput{
		InfoType:        aws.String("EnvironmentInfoType"), // Required
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.RequestEnvironmentInfo(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_RestartAppServer() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.RestartAppServerInput{
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.RestartAppServer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_RetrieveEnvironmentInfo() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.RetrieveEnvironmentInfoInput{
		InfoType:        aws.String("EnvironmentInfoType"), // Required
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.RetrieveEnvironmentInfo(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_SwapEnvironmentCNAMEs() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.SwapEnvironmentCNAMEsInput{
		DestinationEnvironmentId:   aws.String("EnvironmentId"),
		DestinationEnvironmentName: aws.String("EnvironmentName"),
		SourceEnvironmentId:        aws.String("EnvironmentId"),
		SourceEnvironmentName:      aws.String("EnvironmentName"),
	}
	resp, err := svc.SwapEnvironmentCNAMEs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_TerminateEnvironment() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.TerminateEnvironmentInput{
		EnvironmentId:      aws.String("EnvironmentId"),
		EnvironmentName:    aws.String("EnvironmentName"),
		ForceTerminate:     aws.Bool(true),
		TerminateResources: aws.Bool(true),
	}
	resp, err := svc.TerminateEnvironment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_UpdateApplication() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.UpdateApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		Description:     aws.String("Description"),
	}
	resp, err := svc.UpdateApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_UpdateApplicationVersion() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.UpdateApplicationVersionInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		VersionLabel:    aws.String("VersionLabel"),    // Required
		Description:     aws.String("Description"),
	}
	resp, err := svc.UpdateApplicationVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_UpdateConfigurationTemplate() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.UpdateConfigurationTemplateInput{
		ApplicationName: aws.String("ApplicationName"),           // Required
		TemplateName:    aws.String("ConfigurationTemplateName"), // Required
		Description:     aws.String("Description"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
				Value:        aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		OptionsToRemove: []*elasticbeanstalk.OptionSpecification{
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateConfigurationTemplate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_UpdateEnvironment() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.UpdateEnvironmentInput{
		ApplicationName: aws.String("ApplicationName"),
		Description:     aws.String("Description"),
		EnvironmentId:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
		GroupName:       aws.String("GroupName"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
				Value:        aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		OptionsToRemove: []*elasticbeanstalk.OptionSpecification{
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
			},
			// More values...
		},
		SolutionStackName: aws.String("SolutionStackName"),
		TemplateName:      aws.String("ConfigurationTemplateName"),
		Tier: &elasticbeanstalk.EnvironmentTier{
			Name:    aws.String("String"),
			Type:    aws.String("String"),
			Version: aws.String("String"),
		},
		VersionLabel: aws.String("VersionLabel"),
	}
	resp, err := svc.UpdateEnvironment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticBeanstalk_ValidateConfigurationSettings() {
	svc := elasticbeanstalk.New(session.New())

	params := &elasticbeanstalk.ValidateConfigurationSettingsInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{ // Required
			{ // Required
				Namespace:    aws.String("OptionNamespace"),
				OptionName:   aws.String("ConfigurationOptionName"),
				ResourceName: aws.String("ResourceName"),
				Value:        aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		EnvironmentName: aws.String("EnvironmentName"),
		TemplateName:    aws.String("ConfigurationTemplateName"),
	}
	resp, err := svc.ValidateConfigurationSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
