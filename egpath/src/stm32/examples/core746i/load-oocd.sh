#!/bin/sh

INTERFACE=stlink-v2
TARGET=stm32f7x
TRACECLKIN=192000000

. ../../../../../scripts/load-oocd.sh $@
