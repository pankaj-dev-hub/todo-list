#!/bin/bash

# Run user service in background
./user &

# Run todo service in foreground
./todo
