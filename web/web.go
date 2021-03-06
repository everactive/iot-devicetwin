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
	"fmt"
	"github.com/CanonicalLtd/iot-devicetwin/service/controller"
	"net/http"

	"github.com/CanonicalLtd/iot-devicetwin/config"
)

// Service is the implementation of the web API
type Service struct {
	Settings   *config.Settings
	Controller controller.Controller
}

// NewService returns a new web controller
func NewService(settings *config.Settings, ctrl controller.Controller) *Service {
	return &Service{
		Settings:   settings,
		Controller: ctrl,
	}
}

// Run starts the web service
func (wb Service) Run() error {
	fmt.Printf("Starting service on port :%s\n", wb.Settings.Port)
	return http.ListenAndServe(":"+wb.Settings.Port, wb.Router())
}
