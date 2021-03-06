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

package main

import (
	"log"
	"net/http"

	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/ubuntu-core/smb-admin/service"
)

func main() {
	i18n.MustLoadTranslationFile("./lang/en-us.all.json")

	env := service.Env{Config: service.DefaultConfig()}
	env.Config.Interface = service.InterfaceTypeUser

	// Parse the command-line parameters
	service.ParseArgs()

	// Get the config settings from the file or environment variables
	err := service.ReadConfig(&env.Config)
	if err != nil {
		log.Fatalf("Error reading the config: %v", err)
	}

	log.Println("Server will run on:", env.Config.PortUser)
	port := ":" + env.Config.PortUser

	log.Fatal(http.ListenAndServe(port, service.UserRouter(&env)))
}
