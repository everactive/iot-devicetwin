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

package web

import (
	"github.com/CanonicalLtd/iot-devicetwin/config"
	"testing"
)

func TestService_DeviceGet(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		code   int
		result string
	}{
		{"valid", "/v1/device/abc/a111", 200, ""},
		{"invalid", "/v1/device/abc/invalid", 400, "DeviceGet"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wb := NewService(config.TestConfig(), testController())

			w := sendRequest("GET", tt.url, nil, wb)
			if w.Code != tt.code {
				t.Errorf("Web.DeviceGet() got = %v, want %v", w.Code, tt.code)
			}
			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Web.DeviceGet() got = %v", err)
			}
			if resp.Code != tt.result {
				t.Errorf("Web.DeviceGet() got = %v, want %v", resp.Code, tt.result)
			}
		})
	}
}

func TestService_DeviceList(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		code   int
		result string
	}{
		{"valid", "/v1/device/abc", 200, ""},
		{"invalid", "/v1/device/invalid", 400, "DeviceList"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wb := NewService(config.TestConfig(), testController())

			w := sendRequest("GET", tt.url, nil, wb)
			if w.Code != tt.code {
				t.Errorf("Web.DeviceList() got = %v, want %v", w.Code, tt.code)
			}
			resp, err := parseDevicesResponse(w.Body)
			if err != nil {
				t.Errorf("Web.DeviceList() got = %v", err)
			}
			if resp.Code != tt.result {
				t.Errorf("Web.DeviceList() got = %v, want %v", resp.Code, tt.result)
			}
		})
	}
}
