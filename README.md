# Ordering System

> Neil Medina
> 
> What I want to achieve in this system.
> - Well documented system.
> 	- Changelog
>	- Badges
>	- Readthedocs
>	- Explantion in every function(Description in Comments)
>	- Unit Testing
> - Using stardard coding to all languages with references.
>   - CSS (SASS)
>   - JS (jquery)
>   - GO (MVC)
> - The features of this system has government policy in mind.
>   - Tax( policy here )

------------


[![Go Report Card](https://goreportcard.com/badge/github.com/ngmedina14/ordering-system)](https://goreportcard.com/report/github.com/ngmedina14/ordering-system)
[![Go Reference](https://pkg.go.dev/badge/github.com/ngmedina14/ordering-system.svg)](https://pkg.go.dev/github.com/ngmedina14/ordering-system)
[![codecov](https://codecov.io/gh/ngmedina14/ordering-system/branch/master/graph/badge.svg?token=KJ5LKK67EF)](https://codecov.io/gh/ngmedina14/ordering-system)
[![Build Status](https://travis-ci.org/ngmedina14/ordering-system.svg?branch=master)](https://travis-ci.org/ngmedina14/ordering-system)


Personal Project of - Neil Medina [ngmedina14@gmail.com](https://github.com/ngmedina14 "ngmedina14@gmail.com") (*Credit Creator*)

<!-- For Documentation: 

- [Link of Documentation](https://github.com/ngmedina14/ordering-system "Link of Documentation")

Social Media:

- [Facebook](https://www.facebook.com)

## Screenshots

### GUI OF THE SYSTEM

![Dashboard](https://cdn.corporatefinanceinstitute.com/assets/systems-thinking.jpeg)
&nbsp;
-->
## Features

- Online Ordering
- Cashier System
- E-cash payment
- Order Monitoring
- Sales Analysis
- Printing Receipt

## Minimum requirements

| Operating System                   |                Architectures              |                                Notes                                                |
|------------------------------------|-------------------------------------------|-------------------------------------------------------------------------------------|
| FreeBSD 10.3 or later              |  amd64, 386                               | Debian GNU/kFreeBSD not supported                                                   |
| Linux 2.6.23 or later with glibc   |  amd64, 386, arm, arm64, s390x, ppc64le   | CentOS/RHEL 5.x not supported. Install from source for other libc.                  |
| macOS 10.10 or later               |  amd64                                    | Use the clang or gcc<sup>†</sup> that comes with Xcode<sup>‡</sup> for cgo support. |
| Windows 7, Server 2008 R2 or later |  amd64, 386                               | Use MinGW gcc<sup>†</sup>. No need for cygwin or msys.                              |


### Hardware

- RAM - minimum 256MB
- CPU - minimum 2GHz

### Software

- Go Version 1.15 or later

## Installation

```bash
$ go get -u github.com/ngmedina14/ordering-system/...
```

<!-- some description of the installation -->

Prepare modules

```bash
$ go mod init

$ go mod tidy
```

Run your app (Linux, Apple macOS or Windows):

In Linux:

```bash
$ go build; ./program
```

In Windows:

```bash
> go build && program.exe
```

<!--# Quick Reference

Other Things Here -->
