#!/bin/bash

sourcePath=../../../config_system/Online_Config/
outJsonPath=../../assets/
outGoPath=../../config/
./linux/system_cfg_parse -sp $sourcePath -ojp $outJsonPath -ogp $outGoPath

bash ./parse_slot_online.sh ./linux/system_cfg_parse