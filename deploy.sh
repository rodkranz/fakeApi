#!/usr/bin/env bash

main="main.go"
bin="server"

folderWWW="$HOME/www/fake-api"
folderBKP="$HOME/www/fake-api-bkp-`date +%Y%m%d-%H%M%S`"


# Update repository
git pull origin master

# bind data info
go-bindata -o "./modules/bindata/bindata.go" -pkg "bindata" conf/*
# build new version of project
go build -o "$bin" "$main"

# if folder exists rename to keep backup
if [ -d "$folderWWW" ]; then
  mv "$folderWWW" "$folderBKP"
fi

# recreate folder
mkdir -p "$folderWWW"

# copy local files to destiny
cp -r "public" "$folderWWW"
cp -r "templates" "$folderWWW"
cp -r "fakes" "$folderWWW"

if [ -d "$folderBKP" ]; then
    # recover old configuration
    cp -rf "$folderBKP/custom" "$folderWWW/"
    cp -rf "$folderBKP/fakes/*" "$folderWWW/fakes/."
else
    # create a custom configuration
    mkdir -p "$folderWWW/custom"
    cp "conf" "$folderWWW/custom"
fi

# copy binary
cp "$bin" "$folderWWW"

# reload supervisor
sudo supervisorctl stop all
sudo supervisorctl start all

echo "The deploy is done!"
