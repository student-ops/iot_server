#!/bin/sh
set -e

influx bucket create -n vuoy_monitor -o iot -r 72h