// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017-2018 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package service

import "testing"

import "os"

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.DocRootAdmin != defaultDocRootAdmin {
		t.Errorf("Expected default docroot of %s, got %s", defaultDocRootAdmin, config.DocRootAdmin)
	}
	if config.PortAdmin != defaultPortAdmin {
		t.Errorf("Expected default port of %s, got %s", defaultPortAdmin, config.PortAdmin)
	}
}

func TestReadConfigFromFile(t *testing.T) {
	settingsFile = "../settings.example.yaml"
	config := ConfigSettings{}
	err := ReadConfig(&config)
	if err != nil {
		t.Errorf("Error reading config file: %v", err)
	}
}

func TestReadConfigInvalidPath(t *testing.T) {
	settingsFile = "not a good path"
	config := ConfigSettings{}
	err := ReadConfig(&config)
	if err == nil {
		t.Error("Expected an error with an invalid config file.")
	}
}

func TestReadConfigInvalidFile(t *testing.T) {
	settingsFile = "../README.md"
	config := ConfigSettings{}
	err := ReadConfig(&config)
	if err == nil {
		t.Error("Expected an error with an invalid config file.")
	}
}

func TestReadConfigEnv(t *testing.T) {
	settingsFile = ""
	config := ConfigSettings{}
	os.Setenv(envTitle, "SuperMegaBiz Admin")
	os.Setenv(envLogo, "supermegabiz.jpg")
	os.Setenv(envPortAdmin, "9000")
	os.Setenv(envDocRootAdmin, "../")
	os.Setenv(envPortUser, "9001")
	err := ReadConfig(&config)
	if err != nil {
		t.Errorf("Error setting config from env: %v", err)
	}

	if config.Title != "SuperMegaBiz Admin" {
		t.Errorf("Error setting title from env, expected 'SuperMegaBiz Admin', got: %v", config.Title)
	}
	if config.Logo != "supermegabiz.jpg" {
		t.Errorf("Error setting title from env, expected 'supermegabiz.jpg', got: %v", config.Logo)
	}
	if config.PortAdmin != "9000" {
		t.Errorf("Error setting port from env, expected 9000, got: %v", config.PortAdmin)
	}
	if config.DocRootAdmin != "../" {
		t.Errorf("Error setting docroot from env, expected '../', got: %v", config.DocRootAdmin)
	}
}

func TestReadConfigEnvInvalidPort(t *testing.T) {
	settingsFile = ""
	config := ConfigSettings{}

	os.Setenv(envPortAdmin, "INVALID")

	err := ReadConfig(&config)
	if err == nil {
		t.Errorf("Expected invalid port error for %v", config.PortAdmin)
	}
}
