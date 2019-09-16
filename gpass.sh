#!/bin/sh

export LD_LIBRARY_PATH=$ORACLE_HOME/lib:/lib:/usr/lib
export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig

gpass -conf ./conf/gpass.conf
