#!/bin/bash

read name

ending=pb

folderName="$name$ending"

protoc $name/$folderName/$name.proto --go_out=plugins=grpc:.
