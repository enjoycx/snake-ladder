#!/bin/bash


slotOutGoPath=../../../../controller/spin/controller/config/
slotOutJsonPath=../../../../assets/slot/
for i in `ls ../../../../../config_slots/${2}/ | grep '^slot*'`;
    do ${1} -s=false -f ../../../../../config_slots/${2}/${i} -ogp ${slotOutGoPath} -ojp ${slotOutJsonPath}
done