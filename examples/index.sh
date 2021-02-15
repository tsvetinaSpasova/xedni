#!/bin/bash

# Example of inverted index api /index

curl -X POST -H 'Content-Type application/json' -d @./example-index.json localhost:8888/api/index
