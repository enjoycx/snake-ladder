import requests
import json

address = 'http://10.0.84.28:4050/api/version/get'
data = {'versionType':'IOS_appstore_test_24'}
rsp = requests.post(url=address,json=data)
print(rsp.text)

address = 'http://10.0.84.28:4050/api/version/upload'
data = {'versionType':'IOS_appstore_test_24', 'versionInfo': '{"abPatchIndexs":{"0":7337593,"1":6958525,"2":4825246,"3":4754652},"abVersion":0,"appVersion":21,"downLoadApkUrl":"http://10.0.88.93:8086/s/Ox66","inAppMachines":{"machinejalapenofiesta":1,"machinepowerofolympus":1},"machinesIndex":0,"machinesInfos":{"machinebeautyofegypt":0,"machinecoinrush":0,"machinedancinglion":0,"machinedjgrandma":0,"machinedollarrush":0,"machinedynamiteblast":0,"machinefortuneofleprechaun":0,"machinejalapenofiesta":0,"machinepowerofolympus":0,"machines":0,"machinewildclub":0,"machineyellowbrickroad":0}}'}
print(data)
rsp = requests.post(url=address,json=data)
print(rsp.text)