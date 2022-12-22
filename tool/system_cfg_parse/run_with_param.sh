#! /bin/bash

sourcePath=../../../config_system/${2}/
outJsonPath=../../assets/
outGoPath=../../config/
./${3}/system_cfg_parse -sp $sourcePath -ojp $outJsonPath -ogp $outGoPath

sourcePath2=../../../config_system/${2}/System_Dev_Config/
./${3}/system_cfg_parse -sp $sourcePath2 -ojp $outJsonPath -ogp $outGoPath


slotOutGoPath=../../controller/spin/config/
slotOutJsonPath=../../assets/slot/
for i in `ls ../../../config_slots/${1}/ | grep '^slot*'`;
    do ./${3}/system_cfg_parse -s=false -f ../../../config_slots/${1}/${i} -ogp ${slotOutGoPath} -ojp ${slotOutJsonPath}
done

for i in `ls ../../../config_slots/${1}/ | grep '^progressive'`;
    do ./${3}/system_cfg_parse -f ../../../config_slots/${1}/${i} -ojp ${outJsonPath} -ogp ${outGoPath}
done

for i in `ls ../../../config_slots/${1}/ | grep '^bg_transition_mapping'`;
    do ./${3}/system_cfg_parse -f ../../../config_slots/${1}/${i} -ojp ${outJsonPath} -ogp ${outGoPath}
done

for i in `ls ../../../config_slots/${1}/BG/`;
    do ./${3}/system_cfg_parse -f ../../../config_slots/${1}/BG/${i} -ojp ${outJsonPath} -ogp ${outGoPath} -onlyJson true
done