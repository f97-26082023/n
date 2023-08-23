# n.f97.xyz
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/f97/n/blob/master/LICENSE)
[![Latest Build](https://img.shields.io/github/actions/workflow/status/f97/n/docker-snapshot.yml?branch=main)](https://github.com/f97/n/actions)
[![Go Report](https://goreportcard.com/badge/github.com/f97/n)](https://goreportcard.com/report/github.com/f97/n)
[![Latest Docker Image Size](https://img.shields.io/docker/image-size/f97/n.svg?style=flat)](https://hub.docker.com/r/f97/n)
[![Latest Release](https://img.shields.io/github/release/f97/n.svg?style=flat)](https://github.com/f97/n/releases)

## Introduction
n.f97.xyz is a lightweight personal bookkeeping app hosted by yourself. It can be deployed on almost all platforms, including Windows, macOS and Linux on x86, amd64 and ARM architectures. You can even deploy it on an raspberry device. It also supports many different databases, including sqlite and mysql. With docker, you can just deploy it via one command without complicated configuration.

## Features
1. Open source & Self-hosted
2. Lightweight & Fast
3. Easy to install
    * Docker support
    * Multiple database support (SQLite, MySQL, PostgreSQL, etc.)
    * Multiple operation system & hardware support (Windows, macOS, Linux & x86, amd64, ARM)
4. User-friendly interface
    * Close to native app experience (for mobile device)
    * Two-level account & two-level category support
    * Plentiful preset categories
    * Geographic location and map support
    * Searching & filtering history records
    * Data statistics
    * Dark theme
5. Multiple currency support & automatically updating exchange rates
6. Multiple timezone support
7. Multi-language support
8. Two-factor authentication
9. Application lock (PIN code / WebAuthn)
10. Data export

## Screenshots
### Mobile Device
[![n.f97.xyz](https://raw.githubusercontent.com/wiki/f97/n/img/en.png)](https://raw.githubusercontent.com/wiki/f97/n/img/en.png)

## Installation
### Ship with docker
Visit [Docker Hub](https://hub.docker.com/r/f97/n) to see all images and tags.

Latest Release:

    $ docker run -p8080:8080 f97/n

Latest Daily Build:

    $ docker run -p8080:8080 f97/n:latest-snapshot

### Install from binary
Latest release: [https://github.com/f97/n/releases](https://github.com/f97/n/releases)

    $ ./ezbookkeeping server run

n.f97.xyz will listen at port 8080 as default. Then you can visit http://{YOUR_HOST_ADDRESS}:8080/ .

### Build from source
Make sure you have [Golang](https://golang.org/), [GCC](http://gcc.gnu.org/), [Node.js](https://nodejs.org/) and [NPM](https://www.npmjs.com/) installed. Then download the source code, and follow these steps:

    $ ./build.sh package -o ezbookkeeping.tar.gz

All the files will be packaged in `ezbookkeeping.tar.gz`.

You can also build docker image, make sure you have [docker](https://www.docker.com/) installed, then follow these steps:

    $ ./build.sh docker

## Documents
1. [English](http://ezbookkeeping.mayswind.net)
1. [简体中文 (Simplified Chinese)](http://ezbookkeeping.mayswind.net/zh_Hans)

## License
[MIT](https://github.com/f97/n/blob/master/LICENSE)
