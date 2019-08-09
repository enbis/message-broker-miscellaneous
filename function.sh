#!/bin/bash

_get_config_value() {
    IFS=':' 
    value=$(cat config.yml | grep $1)

    IFS=':' read -r -a array <<< "$value"

    res="$(echo -e "${array[1]}" | tr -d '[:space:]')"

	echo ${res}
}

$*