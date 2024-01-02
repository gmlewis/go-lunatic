#!/bin/bash -e
# -*- compile-command: "./run-all.sh"; -*-

EXAMPLES=$(echo */main.go)
for i in ${EXAMPLES}; do
    dir=${i%"/main.go"}
    echo && echo && echo "examples/$dir"
    pushd $dir
    ./build-go.sh
    ./run.sh
    popd
done
