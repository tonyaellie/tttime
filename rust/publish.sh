#!/bin/bash

python publish.py

cargo package
cargo publish --allow-dirty # as this is running in CI any dirty files should be ignored