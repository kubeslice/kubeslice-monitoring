#!/bin/bash

header="//This file is autogenerated, do not edit;"
license=$(
  cat <<EOF
/*
 *  Copyright (c) 2022 Avesha, Inc. All rights reserved.
 *
 *  SPDX-License-Identifier: Apache-2.0
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */
EOF
)
input1="config/events/controller.yaml"
input2="config/events/worker.yaml"
output="pkg/schema/event_names.go"

if [ ! -f $input1 ] || [ ! -f $input2 ]; then
  printf "Error: file does not exist"
  exit 1
fi
printf "%s\n\n%s\n\npackage schema\n\n" "$license" "$header" >$output
printf "var (\n" >>$output
while IFS=$'\t' read -r; do
  x=$(grep -o "name: [A-Za-z0-9]*" | awk '{print $2}')
  IFS=$'\n' read -ra ADDR -d $'\0' <<<"$x"
  for i in "${ADDR[@]}"; do
    printf "\tEvent%s = \"%s\"\n" "${i}" "${i}" >>$output
  done
done < <(paste $input1 $input2)
printf ")" >>$output
