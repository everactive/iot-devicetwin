// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * This file is part of the IoT Device Twin Service
 * Copyright 2019 Canonical Ltd.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU Affero General Public License version 3, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranties of MERCHANTABILITY,
 * SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package devicetwin

import (
	"fmt"
	"github.com/CanonicalLtd/iot-devicetwin/domain"
)

// MockDeviceTwin mocks a device twin service
type MockDeviceTwin struct {
	Actions []string
}

// HealthHandler mocks the health handler
func (twin *MockDeviceTwin) HealthHandler(payload domain.Health) error {
	if payload.DeviceID == "invalid" || payload.DeviceID == "new-device" {
		return fmt.Errorf("MOCK error in health handler")
	}
	return nil
}

// ActionResponse mocks the action handler
func (twin *MockDeviceTwin) ActionResponse(clientID, actionID, action string, payload []byte) error {
	if action == "invalid" {
		return fmt.Errorf("MOCK error in action")
	}
	return nil
}

// DeviceSnaps mocks the snap list
func (twin *MockDeviceTwin) DeviceSnaps(clientID string) ([]domain.DeviceSnap, error) {
	if clientID == "invalid" {
		return nil, fmt.Errorf("MOCK snaps list")
	}
	return []domain.DeviceSnap{
		{Name: "example-snap", InstalledSize: 2000, Status: "active"},
	}, nil
}

// ActionCreate mocks the action log creation
func (twin *MockDeviceTwin) ActionCreate(orgID, deviceID string, act domain.SubscribeAction) error {
	if deviceID == "invalid" {
		return fmt.Errorf("MOCK action log create")
	}
	if twin.Actions == nil {
		twin.Actions = []string{}
	}
	twin.Actions = append(twin.Actions, act.ID)
	return nil
}

// ActionUpdate mocks the action log update
func (twin *MockDeviceTwin) ActionUpdate(actionID, status, message string) error {
	if actionID == "invalid" {
		return fmt.Errorf("MOCK action log update")
	}
	return nil
}