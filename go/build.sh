#!/usr/bin/env bash
docker run -v /root/ndp/NDP-task-courier/go/bin/:/ntc/bin/ -v /root/ndp/NDP-task-courier/go/src/:/ntc/src/ -e "CGO_ENABLED=0" ndp-task-courier-builder
