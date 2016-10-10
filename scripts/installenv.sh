#!/bin/sh

mkdir -p ~/go/{bin,src}

sudo pacman -S gcc sdl2 sdl2_gfx sdl2_image sdl2_ttf

export GOPATH=~/go
echo "faut ajouter au zshrc ou bashrc : export GOPATH=~/go"

go get github.com/gtalent/starfish/gfx
go get github.com/gtalent/starfish/input
