import requests
import json
import time
import datetime 

address = ['http://localhost:4000']
openIds = ['552']

def yellowPrint(data):
    print("\033[33m"+data+"\033[0m")

def redPrint(data):
    print("\033[35m"+data+"\033[0m")

def greenPrint(data):
    print("\033[32m"+data+"\033[0m")




def login(openID, address):
    yellowPrint("remote server:"+address)
    data = {'openID': openID}
    yellowPrint("login openID:"+data['openID'])
    rsp = requests.post(url=address+'/api/login',json=data)
    user = rsp.json()['data']['user']
    userID = user['userID']
    headers = {'uid': userID,"openID":openID}
    print(rsp.text)
    return headers

reqStateNil=[]
reqStateFail={}
reqStateRetCode={}
reqStateSuccess=[]

def req(address, headers, router, data):
    print("\n")
    yellowPrint(router+" start:")
    t0=time.time()
    rsp = requests.post(url=address+'/api/'+router,json=data,headers=headers)
    t1=time.time()
    duration=str(((t1-t0)))
    if rsp.status_code!=200:
        reqStateFail[router]=rsp.status_code
        redPrint(router+" http error:"+rsp)
        return
    if rsp.json() is None:
        reqStateNil.append(router)
        redPrint(router +" no implemented")
        return
    if rsp.json()['retCode']!=0:
        reqStateRetCode[router] = rsp.json()
        redPrint(router +" respose error")
        return
    reqStateSuccess.append(router)
    print(rsp.text)
    yellowPrint(router+" end"+" duration: "+duration+" sec")
    return rsp


for adr in address:
    for id in openIds:
        headers = login(id,adr)
        req(adr,headers,'batman/gm',{'option': 0,'param': 'ChangeTime,20220101 150000'})
        # req(adr,headers,'quest/getQuestInfo',{'questID':2})
        # req(adr,headers,'quest/selectQuestDiff',{'difficulty':1})
        # req(adr,headers,'user/cdKeyReward',{'cdKey': 'uvz05r','index':-1})

print("\n")
greenPrint("rsp is success:")
print(reqStateSuccess)

if len(reqStateNil)>0:
    print("\n")
    redPrint("Fail reason:controller not implemented rsp is nil:")
    print(reqStateNil)

if len(reqStateFail)>0:
    print("\n")
    redPrint("Fail http error:")
    for k,v in reqStateFail.items():
        print (k)
        print (v)

if len(reqStateRetCode)>0:
    print("\n")
    redPrint("Fail response retCode error:")
    for k,v in reqStateRetCode.items():
        print (k)
        print (v)