#!/usr/bin/env bash

e="$@"

if [[ -z $e ]]; then
	read -p "Entity name(s): " e
fi

go run -mod=mod entgo.io/ent/cmd/ent new $e
