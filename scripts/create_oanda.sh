#!/bin/bash

#this will exit the script if any none 0 values are returned
set -euo pipefail
IFS=$'\n\t'

psql -U postgres -f create_oanda.sql
