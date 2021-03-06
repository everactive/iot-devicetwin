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

package controller

import "github.com/CanonicalLtd/iot-devicetwin/domain"

// DeviceGet gets the device from the database cache
func (srv *Service) DeviceGet(orgID, clientID string) (domain.Device, error) {
	return srv.DeviceTwin.DeviceGet(orgID, clientID)
}

// DeviceList gets the devices from the database cache
func (srv *Service) DeviceList(orgID string) ([]domain.Device, error) {
	return srv.DeviceTwin.DeviceList(orgID)
}
