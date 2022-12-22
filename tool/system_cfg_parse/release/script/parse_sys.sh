#! /bin/bash

cd ${1}
rm -f slot_*
svn update
cp -r  ${1} ${3}

cd ${2}
svn update
cp -r  ${2} ${3}

cd ${3}/maxbet_slot_backend/tool/system_cfg_parse/release/script
sourcePath=../../../../../config_system/${5}/
outJsonPath=../../../../assets/
outGoPath=../../../../config/
../../${4}/system_cfg_parse -sp $sourcePath -ojp $outJsonPath -ogp $outGoPath

bash ./parse_slot.sh ../../${4}/system_cfg_parse ${5}