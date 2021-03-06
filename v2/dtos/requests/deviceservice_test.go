//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testAddDeviceService = AddDeviceServiceRequest{
	BaseRequest: common.BaseRequest{
		RequestID: ExampleUUID,
	},
	Service: dtos.DeviceService{
		Name:           TestDeviceServiceName,
		BaseAddress:    TestBaseAddress,
		OperatingState: models.Enabled,
		Labels:         []string{"MODBUS", "TEMP"},
		AdminState:     models.Locked,
	},
}

var testUpdateDeviceService = UpdateDeviceServiceRequest{
	BaseRequest: common.BaseRequest{
		RequestID: ExampleUUID,
	},
	Service: mockDeviceServiceDTO(),
}

func mockDeviceServiceDTO() dtos.UpdateDeviceService {
	testUUID := ExampleUUID
	testName := TestDeviceServiceName
	testBaseAddress := TestBaseAddress
	testOperatingState := models.Enabled
	testAdminState := models.Locked
	ds := dtos.UpdateDeviceService{}
	ds.Id = &testUUID
	ds.Name = &testName
	ds.BaseAddress = &testBaseAddress
	ds.OperatingState = &testOperatingState
	ds.AdminState = &testAdminState
	ds.Labels = testLabels
	return ds
}

func TestAddDeviceServiceRequest_Validate(t *testing.T) {
	valid := testAddDeviceService
	noReqID := testAddDeviceService
	noReqID.RequestID = ""
	invalidReqID := testAddDeviceService
	invalidReqID.RequestID = "jfdw324"
	noName := testAddDeviceService
	noName.Service.Name = ""
	noOperatingState := testAddDeviceService
	noOperatingState.Service.OperatingState = ""
	invalidOperatingState := testAddDeviceService
	invalidOperatingState.Service.OperatingState = "invalid"
	noAdminState := testAddDeviceService
	noAdminState.Service.OperatingState = ""
	invalidAdminState := testAddDeviceService
	invalidAdminState.Service.OperatingState = "invalid"
	noBaseAddress := testAddDeviceService
	noBaseAddress.Service.BaseAddress = ""
	invalidBaseAddress := testAddDeviceService
	invalidBaseAddress.Service.BaseAddress = "invalid"
	tests := []struct {
		name          string
		DeviceService AddDeviceServiceRequest
		expectError   bool
	}{
		{"valid AddDeviceServiceRequest", valid, false},
		{"valid AddDeviceServiceRequest, no Request Id", noReqID, false},
		{"invalid AddDeviceServiceRequest, Request Id is not an uuid", invalidReqID, true},
		{"invalid AddDeviceServiceRequest, no Name", noName, true},
		{"invalid AddDeviceServiceRequest, no OperatingState", noOperatingState, true},
		{"invalid AddDeviceServiceRequest, invalid OperatingState", invalidOperatingState, true},
		{"invalid AddDeviceServiceRequest, no AdminState", noAdminState, true},
		{"invalid AddDeviceServiceRequest, invalid AdminState", invalidAdminState, true},
		{"invalid AddDeviceServiceRequest, no BaseAddress", noBaseAddress, true},
		{"invalid AddDeviceServiceRequest, no BaseAddress", invalidBaseAddress, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.DeviceService.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addDeviceServiceRequest validation result.", err)
		})
	}
}

func TestAddDeviceService_UnmarshalJSON(t *testing.T) {
	valid := testAddDeviceService
	resultTestBytes, _ := json.Marshal(testAddDeviceService)
	type args struct {
		data []byte
	}
	tests := []struct {
		name             string
		addDeviceService AddDeviceServiceRequest
		args             args
		wantErr          bool
	}{
		{"unmarshal AddDeviceServiceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddDeviceServiceRequest, empty data", AddDeviceServiceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddDeviceServiceRequest, string data", AddDeviceServiceRequest{}, args{[]byte("Invalid AddDeviceServiceRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addDeviceService
			err := tt.addDeviceService.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.addDeviceService, "Unmarshal did not result in expected AddDeviceServiceRequest.")
			}
		})
	}
}

func TestAddDeviceServiceReqToDeviceServiceModels(t *testing.T) {
	requests := []AddDeviceServiceRequest{testAddDeviceService}
	expectedDeviceServiceModel := []models.DeviceService{{
		Name:           TestDeviceServiceName,
		BaseAddress:    TestBaseAddress,
		OperatingState: models.Enabled,
		Labels:         []string{"MODBUS", "TEMP"},
		AdminState:     models.Locked,
	}}
	resultModels := AddDeviceServiceReqToDeviceServiceModels(requests)
	assert.Equal(t, expectedDeviceServiceModel, resultModels, "AddDeviceServiceReqToDeviceServiceModels did not result in expected DeviceService model.")
}

func TestUpdateDeviceService_UnmarshalJSON(t *testing.T) {
	valid := testUpdateDeviceService
	resultTestBytes, _ := json.Marshal(testUpdateDeviceService)
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		req     UpdateDeviceServiceRequest
		args    args
		wantErr bool
	}{
		{"unmarshal UpdateDeviceServiceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid UpdateDeviceServiceRequest, empty data", UpdateDeviceServiceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid UpdateDeviceServiceRequest, string data", UpdateDeviceServiceRequest{}, args{[]byte("Invalid UpdateDeviceServiceRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.req
			err := tt.req.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.req, "Unmarshal did not result in expected UpdateDeviceServiceRequest.", err)
			}
		})
	}
}

func TestUpdateDeviceServiceRequest_Validate(t *testing.T) {
	valid := testUpdateDeviceService
	validWithoutId := valid
	validWithoutId.Service.Id = nil
	validWithoutName := valid
	validWithoutName.Service.Name = nil
	noReqID := valid
	noReqID.RequestID = ""
	invalidReqID := valid
	invalidReqID.RequestID = "2h022mc"
	noIdAndName := valid
	noIdAndName.Service.Id = nil
	noIdAndName.Service.Name = nil
	invalidOperatingState := valid
	invalid := "invalid"
	invalidOperatingState.Service.OperatingState = &invalid
	invalidAdminState := valid
	invalidAdminState.Service.OperatingState = &invalid
	tests := []struct {
		name        string
		req         UpdateDeviceServiceRequest
		expectError bool
	}{
		{"valid UpdateDeviceServiceRequest", valid, false},
		{"valid UpdateDeviceServiceRequest without Id", validWithoutId, false},
		{"valid UpdateDeviceServiceRequest without name", validWithoutName, false},
		{"valid UpdateDeviceServiceRequest, no Request Id", noReqID, false},
		{"invalid UpdateDeviceServiceRequest, no Request Id", invalidReqID, true},
		{"invalid UpdateDeviceServiceRequest, no Id and Name", noIdAndName, true},
		{"invalid UpdateDeviceServiceRequest, invalid OperatingState", invalidOperatingState, true},
		{"invalid UpdateDeviceServiceRequest, invalid AdminState", invalidAdminState, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateDeviceServiceRequest validation result.", err)
		})
	}
}

func TestUpdateDeviceServiceRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"service":{
			"name":"test device service"
		}
	}`
	var req UpdateDeviceServiceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, nil, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Service.BaseAddress)
	assert.Nil(t, req.Service.AdminState)
	assert.Nil(t, req.Service.OperatingState)
	assert.Nil(t, req.Service.Labels)
}

func TestUpdateDeviceServiceRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"service":{
			"name":"test device",
			"labels":[]
		}
	}`
	var req UpdateDeviceServiceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.Service.Labels)
}

func TestReplaceDeviceServiceModelFieldsWithDTO(t *testing.T) {
	ds := models.DeviceService{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: "test device service",
	}
	patch := mockDeviceServiceDTO()

	ReplaceDeviceServiceModelFieldsWithDTO(&ds, patch)

	assert.Equal(t, TestBaseAddress, ds.BaseAddress)
	assert.Equal(t, models.Enabled, string(ds.OperatingState))
	assert.Equal(t, models.Locked, string(ds.AdminState))
	assert.Equal(t, testLabels, ds.Labels)
}
