#!/bin/bash
go build client.go protocol.go
go build server.go protocol.go