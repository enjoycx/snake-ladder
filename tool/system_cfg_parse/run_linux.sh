#!/bin/bash

sourcePath=../../../config_system/Dev_Config/
outJsonPath=../../assets/
outGoPath=../../config/
./linux/system_cfg_parse -sp $sourcePath -ojp $outJsonPath -ogp $outGoPath

sourcePath2=../../../config_system/Dev_Config/System_Dev_Config/
./linux/system_cfg_parse -sp $sourcePath2 -ojp $outJsonPath -ogp $outGoPath

bash ./parse_slot.sh ./linux/system_cfg_parse