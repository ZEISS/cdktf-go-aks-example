package main

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/stretchr/testify/assert"
)

// The tests below are example tests, you can find more information at
// https://cdk.tf/testing

var (
	stack = NewCluster(cdktf.Testing_App(nil), "stack")
	synth = cdktf.Testing_Synth(stack, jsii.Bool(true))
)

func TestShouldContainK8sCluster(t *testing.T) {
	assertion := cdktf.Testing_ToHaveResource(synth, azurerm.KubernetesCluster_TfResourceType())
	assert.True(t, *assertion)
}

/*
func TestShouldUseUbuntuImage(t *testing.T){
	properties := map[string]interface{}{
		"name": "ubuntu:latest",
	}
	assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, docker.Image_TfResourceType(), &properties)

	if !*assertion  {
		t.Error("Assertion Failed")
	}
}

func TestCheckValidity(t *testing.T){
	assertion := cdktf.Testing_ToBeValidTerraform(cdktf.Testing_FullSynth(stack))

	if !*assertion  {
		t.Error("Assertion Failed")
	}
}
*/
