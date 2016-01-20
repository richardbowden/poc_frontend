#!/bin/bash
set -e

header() {
  echo "================================================================================"
  echo $@ | sed  -e :a -e 's/^.\{1,77\}$/ & /;ta'
  echo "================================================================================"

}
header "Starting up..."

header "Resolving the configuration..."
# TODO: Here we should resolve the configuration from the config service
export RND_BACKEND_URL="http://random-generator"
export CALC_PI_BACKEND_URL="http://calc-pi"
export FRONTEND_PORT=8080

header "Running the application"
./poc/frontend-svc