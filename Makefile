##
## Makefile for gomoku in /home/karraz_s/rendu/Gomoku
## 
## Made by stephane karraz
## Login   <karraz_s@epitech.net>
## 
## Started on  Mon Oct 10 23:14:47 2016 stephane karraz
## Last update Tue Oct 11 00:00:02 2016 stephane karraz
##

NAME	= Gomoku

SRC	= main.go

all	: $(NAME)

$(NAME)	:
	go build -o $(NAME) $(SRC)

clean	:
	rm -f $(NAME)

fclean	: clean

re	: clean all

.PHONY	: all clean re
