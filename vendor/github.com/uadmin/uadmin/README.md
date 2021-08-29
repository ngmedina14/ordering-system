# uAdmin the Golang Web Framework

Easy to use, blazing fast and secure.

[![go report card](https://goreportcard.com/badge/github.com/uadmin/uadmin "go report card")](https://goreportcard.com/report/github.com/uadmin/uadmin)
[![GoDoc](https://godoc.org/github.com/uadmin/uadmin?status.svg)](https://godoc.org/github.com/uadmin/uadmin)
[![codecov](https://codecov.io/gh/uadmin/uadmin/branch/master/graph/badge.svg)](https://codecov.io/gh/uadmin/uadmin)
[![Build Status](https://travis-ci.org/uadmin/uadmin.svg?branch=master)](https://travis-ci.org/uadmin/uadmin)
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://github.com/uadmin/uadmin/blob/master/LICENSE)

Originally open source by [IntegrityNet Solutions and Services](https://www.integritynet.biz/)

For Documentation:

- [Application in 2 Minutes!](https://www.youtube.com/watch?v=1WwOOYOIQBw&t=41s)
- [Coggle](https://coggle.it/diagram/XSzwl1j7lUdVWvIl/t/uadmin-the-golang-web-framework)
- [Read the Docs](https://uadmin-docs.readthedocs.io/en/latest/)
- [gophers.slack.com #uadmin](https://gophers.slack.com/messages/uadmin/)

[join gophers.slack.com](https://join.slack.com/t/gophers/shared_invite/zt-fajz7jh3-2cpkmFU~hQb8d5LmOCnhfQ)

Social Media:

- [Facebook](https://www.facebook.com/uadminio/)
- [Medium](https://medium.com/@twistedhardware)
- [Twitter](https://twitter.com/uAdminio)

## Screenshots

### Dashboard Menu

![Dashboard](https://github.com/uadmin/uadmin-docs/raw/master/assets/uadmindashboard.png)
&nbsp;

### Log

![Log](https://github.com/uadmin/uadmin-docs/raw/master/assets/log.png)
&nbsp;

### Login Form

![Login Form](https://github.com/uadmin/uadmin-docs/raw/master/tutorial/assets/loginform.png)
&nbsp;

## Features

- AB Testing System
- API Configuration
- Approval System
- Authentication and Permissions
- Clean and sharp UI
- Dashboard customization
- Data Access API (dAPI)
- Database schema migration
- Error Handling
- Export to Excel
- Form and List customization
- Image Cropping
- IP address and port configuration
- Log feature that keeps track of many things in your app
- Metric System
- Multilingual translation
- MySQL Database Support
- Offers FREE hosting for your app while you are developing by using a single command: uadmin publish
- Pretty good security features (SSL, 2-Factor Authentication, Password Reset, Hash Salt, Database Encryption)
- Public access to media
- Self relation of foreign key/many2many
- Sending an email from your app by establishing an email configuration
- System settings which can be used system wide to keep application settings
- Tag support for fields
- Translation files preloading
- Validation for user input
- Webcam support on image and file fields

## Minimum requirements

| Operating System                   |                Architectures              |                                Notes                                                |
|------------------------------------|-------------------------------------------|-------------------------------------------------------------------------------------|
| FreeBSD 10.3 or later              |  amd64, 386                               | Debian GNU/kFreeBSD not supported                                                   |
| Linux 2.6.23 or later with glibc   |  amd64, 386, arm, arm64, s390x, ppc64le   | CentOS/RHEL 5.x not supported. Install from source for other libc.                  |
| macOS 10.10 or later               |  amd64                                    | Use the clang or gcc<sup>†</sup> that comes with Xcode<sup>‡</sup> for cgo support. |
| Windows 7, Server 2008 R2 or later |  amd64, 386                               | Use MinGW gcc<sup>†</sup>. No need for cygwin or msys.                              |

- <sup>†</sup> A C compiler is required only if you plan to use cgo.
- <sup>‡</sup> You only need to install the command line tools for Xcode. If you have already installed Xcode 4.3+, you can install it from the Components tab of the Downloads preferences panel.

### Hardware

- RAM - minimum 256MB
- CPU - minimum 2GHz

### Software

- Go Version 1.12 or later

## Installation

```bash
$ go get -u github.com/uadmin/uadmin/...
```

To test if your installation is fine, run the `uadmin` command line:

```bash
$ uadmin
Usage: uadmin COMMAND [--src]
This tools helps you prepare a folder for a new project or update static files and templates

Commands:
  prepare         Generates folders and prepares static and templates
  version         Shows the version of uAdmin

Arguments:
  --src           If you want to copy static files and templates from src folder

Get full documentation online:
https://uadmin-docs.readthedocs.io/en/latest/
```

## Your First App

Let's build your first app which is a Todo list. First, we will create a folder for your project and prepare it.

```bash
$ mkdir -p ~/go/src/github.com/your_name/todo
$ cd ~/go/src/github.com/your_name/todo
$ uadmin prepare
[   OK   ]   Created: /Users/abdullah/go/src/github.com/twistedhardware/test/models
[   OK   ]   Created: /Users/abdullah/go/src/github.com/twistedhardware/test/api
[   OK   ]   Created: /Users/abdullah/go/src/github.com/twistedhardware/test/views
[   OK   ]   Created: /Users/abdullah/go/src/github.com/twistedhardware/test/media
[  INFO  ]   Copying static/templates from: /Users/abdullah/go/pkg/mod/github.com/uadmin/uadmin@v0.6.0
[   OK   ]   Created: /Users/abdullah/go/src/github.com/twistedhardware/test/static
[   OK   ]   Created: /Users/abdullah/go/src/github.com/twistedhardware/test/templates
```

Now use your code editor to create `main.go` and put this code inside it.

```golang
package main

import (
	"github.com/uadmin/uadmin"
	"time"
)

type Todo struct {
	uadmin.Model
	Name        string
	Description string `uadmin:"html"`
	TargetDate  time.Time
	Progress    int `uadmin:"progress_bar"`
}

func main() {
	uadmin.Register(Todo{})
	uadmin.StartServer()
}
```

Prepare modules

```bash
$ go mod init
go: creating new go.mod: module github.com/twistedhardware/test
go: to add module requirements and sums:
	go mod tidy

$ go mod tidy
go: finding module for package github.com/uadmin/uadmin
go: found github.com/uadmin/uadmin in github.com/uadmin/uadmin v0.6.0
```

Run your app (Linux, Apple macOS or Windows):

```bash
$ go build; ./todo
[   OK   ]   Initializing DB: [14/14]
[   OK   ]   Initializing Languages: [185/185]
[  INFO  ]   Auto generated admin user. Username:admin, Password:admin.
[   OK   ]   Synching System Settings: [49/49]
[   OK   ]   Server Started: http://0.0.0.0:8080
         ___       __          _
  __  __/   | ____/ /___ ___  (_)___
 / / / / /| |/ __  / __ '__ \/ / __ \
/ /_/ / ___ / /_/ / / / / / / / / / /
\__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/
```

In Windows:

```bash
> go build && todo.exe
[   OK   ]   Initializing DB: [14/14]
[   OK   ]   Initializing Languages: [185/185]
[  INFO  ]   Auto generated admin user. Username:admin, Password:admin.
[   OK   ]   Synching System Settings: [49/49]
[   OK   ]   Server Started: http://0.0.0.0:8080
         ___       __          _
  __  __/   | ____/ /___ ___  (_)___
 / / / / /| |/ __  / __  __ \/ / __ \
/ /_/ / ___ / /_/ / / / / / / / / / /
\__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/
```

# Quick Reference

## Overriding Save Function

```golang
func (m *Model) Save() {
	// business logic
	uadmin.Save(m)
}
```

## Validation

```golang
func (m Model) Validate() (ret map[string]string) {
  ret = map[string]string{}
  if m.Name != "test" {
    ret["Name"] = "Error name not found"
  }
  return
}
```

## Dockerize Your App

Create `Docker