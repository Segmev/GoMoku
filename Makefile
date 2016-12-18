##
## Makefile for gomoku in /home/karraz_s/rendu/Gomoku
##
## Made by stephane karraz
## Login   <karraz_s@epitech.net>
##
## Started on  Mon Oct 10 23:14:47 2016 stephane karraz
## Last update Fri Oct 14 01:04:29 2016 stephane karraz
##

export GOPATH=$(shell pwd)

PSRC	= src/gomoku/

NAME	= Gomoku

SRC	= $(PSRC)/main.go

all	: $(NAME)

lib	:
	go get github.com/gtalent/starfish/gfx
	go get github.com/gtalent/starfish/input


$(NAME)	: lib build

build:
	go build -o $(NAME) $(SRC)

clean	:
	go clean

fclean	:
	rm -f $(NAME)

re	: clean all

.PHONY	: all clean re
