# Introduction

This is an improved version of neagix's improved verison of 0xe2-0x9a-0x9b's [Go-SDL](https://github.com/0xe2-0x9a-0x9b/Go-SDL).

The improvements over/differences to neagix's version are:

* working event loop under Windows
* support for RWops

# Installation

Make sure you have SDL, SDL-image, SDL-mixer and SDL-ttf (all in -dev version).

Installing libraries and examples:

    go get -v github.com/asig/Go-SDL/sdl
    go get -v github.com/asig/Go-SDL/ttf
    go get -v github.com/asig/Go-SDL/sdl/audio


# Credits

Music to test SDL-mixer is by Kevin MacLeod.
