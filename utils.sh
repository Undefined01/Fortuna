#!/bin/bash

deps() {
	echo "Installing dependence for ${CI_JOB}..."

	if [ "${CI_JOB}" == "build" ] ; then
		sudo apt-get update
		sudo apt-get install -y -qq make git nodejs zip
		npm i -g npm
		npm i -g node

		node -v
		npm version
		go version
		go env
	else
		go version
		go env
	fi
}

build() {
    echo "Building Fortuna..."

	GOOS="windows" GOARCH="amd64" make

	cp -r frontend/dist deploy
	cp backend/backend.exe deploy/backend.exe || cp backend/backend deploy/backend
	cp backend/config.json deploy/config.json
	cp shell/* deploy/

	mkdir upload
	cd deploy
	zip -r ../upload/fortuna.zip *
	cd ..
}

upload() {
	if [ "${CI_JOB}" == "build" ] ; then
		echo "Uploading deploy package..."

		mkdir -p upload_tmp/${TRAVIS_BRANCH}
		cd upload_tmp

		git init
		git remote add origin https://${Repo_Path}
		git fetch origin ${Deploy_Branch}
		git checkout ${Deploy_Branch}

		rm -rf .git
		git init

		\cp -rf ../upload/* ./${TRAVIS_BRANCH}/
		du -h -d 2 .

		git config user.name "Undefined01-Travis"
		git config user.email "amoscr@163.com"
		git add .
		git commit -m "Auto Built by Travis-CI    ——Undefined01"
		git push --force "https://${Github_Token}@${Repo_Path}" master:${Deploy_Branch}
	fi
}

if [ "$1" == "" ] ; then
	build
fi

if [ "$1" == "deps" ] ; then
	deps
fi

if [ "$1" == "build" ] ; then
	build
fi

if [ "$1" == "test" ] ; then
	test
fi

if [ "$1" == "upload" ] ; then
	upload
fi
