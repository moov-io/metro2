---
layout: page
title: Overview
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Overview

![Moov Metro2 Logo](https://repository-images.githubusercontent.com/268563554/600eb180-c6d9-11ea-85bf-8d8d41d2cac1)

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

Metro2 implements a reader, writer, and validator for consumer credit history reports in an HTTP server and Go library. The HTTP server is available in a [Docker image](https://moov-io.github.io/metro2/usage-docker.html) and the Go package `github.com/moov-io/metro2` is available.

**Note:** By design, Metro2 does not persist (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart Metro2 will have no files or data saved. Also, no in-memory encryption of the data is performed.


