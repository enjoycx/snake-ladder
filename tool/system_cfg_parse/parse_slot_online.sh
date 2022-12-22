#!/bin/bash

outJsonPath=../../assets/
outGoPath=../../config/
slotOutGoPath=../../controller/spin/config/
slotOutJsonPath=../../assets/slot/
for i in `ls ../../../config_slots/Online_Config/ | grep '^slot*'`;
    do ${1} -s=false -f ../../../config_slots/Online_Config/${i} -ogp ${slotOutGoPath} -ojp ${slotOutJsonPath}
done

for i in `ls ../../../config_slots/Online_Config/ | grep '^progressive'`;
    do ${1} -f ../../../config_slots/Online_Config/${i} -ojp ${outJsonPath} -ogp ${outGoPath}
done