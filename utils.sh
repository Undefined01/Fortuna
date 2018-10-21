#!/bin/bash -e

deps() {
	echo "Installing dependence for BUILD environment..."

	sudo apt-get update
	sudo apt-get install -y -qq make git nodejs zip
	npm install -g npm
	npm install -g node

	echo "Installing will be started with the environment followed:"
	node -v
	npm version
	go version
	go env

	cd frontend
	npm install
	cd ../
	cd backend
	go get ./...
	cd ../
}

buildFrontend() {
	cd frontend
	npm run build
	cd ../
}

buildBackend() {
	cd backend
	go build
	cd ../
}

build() {
    echo "Building Fortuna..."

	GOOS="windows" GOARCH="amd64" buildBackend

	buildFrontend

	cp -r frontend/dist deploy
	cp backend/backend.exe deploy/backend.exe || cp backend/backend deploy/backend
	cp shell/* deploy/

	mkdir upload
	cd deploy
	zip -r ../upload/fortuna.zip *
	cd ..
}

if [ "$1" == "deps" ] ; then
	deps
fi

if [ "$1" == "build" ] ; then
	build
fi
