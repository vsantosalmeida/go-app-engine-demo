#!/bin/bash
gcloud beta emulators datastore start --project gcp-app-engine --no-store-on-disk --host-port 0.0.0.0:8000 --quiet