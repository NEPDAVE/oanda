#!/bin/bash
set -euo pipefail
IFS=$'\n\t

#connect to Postgres with our postgres role
psql -U postgres

CREATE DATABASE oanda;
