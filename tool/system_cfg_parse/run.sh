#!/bin/bash

sourcePath=../../../config_system/Dev_Config/
outJsonPath=../../assets/
outGoPath=../../config/
./mac/system_cfg_parse -sp $sourcePath -ojp $outJsonPath -ogp $outGoPath


sourcePath2=../../../config_system/Dev_Config/System_Dev_Config/
./mac/system_cfg_parse -sp $sourcePath2 -ojp $outJsonPath -ogp $outGoPath

bash ./parse_slot.sh ./mac/system_cfg_parse