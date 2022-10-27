#!/bin/bash

curl -s elasticsearch:9200/_cat/indices |
    awk '{print $3}' |
    xargs -i curl elasticsearch:9200/{}
