#!/bin/bash

scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
cd "$scriptDir/.."

./bin/job && ./scripts/publish.sh