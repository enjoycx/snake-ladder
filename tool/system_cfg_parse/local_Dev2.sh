#! /bin/bash

sourcePath=../../../config_system/${2}/
outJsonPath=../../assets/
outGoPath=../../config/
./mac/system_cfg_parse -sp $sourcePath -ojp $outJsonPath -ogp $outGoPath

sourcePath2=../../../config_system/${2}/System_Dev_Config/
./mac/system_cfg_parse -sp $sourcePath2 -ojp $outJsonPath -ogp $outGoPath


slotOutGoPath=../../controller/spin/config/
slotOutJsonPath=../../assets/slot/
for i in `ls ../../../config_slots/${1}/ | grep '^slot*'`;
    do ./mac/system_cfg_parse -s=false -f ../../../config_slots/${1}/${i} -ogp ${slotOutGoPath} -ojp ${slotOutJsonPath}
done

for i in `ls ../../../config_slots/${1}/ | grep '^progressive'`;
    do ./mac/system_cfg_parse -f ../../../config_slots/${1}/${i} -ojp ${outJsonPath} -ogp ${outGoPath}
done