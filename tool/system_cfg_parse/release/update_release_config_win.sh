#! /bin/bash

#absolute path
slot_svn_path="/Users/xmh/Documents/svn/config_slots"
sys_svn_path="/Users/xmh/Documents/svn/config_system"
maxbet_path="/Users/xmh/Documents/project"
platform="windows"
mode="Online_Config"

./script/parse_sys.sh ${slot_svn_path} ${sys_svn_path} ${maxbet_path} ${platform} ${mode}