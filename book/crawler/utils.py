import re


async def GetEvaluateNum(string):
    s = re.search("(\d+)人评价", string)
    num = s.group(1)
    return int(num)


async def GetScore(string):
    s = re.search("\d+(\.\d+)", string)
    if s is None:
        return 0.0
    score = s.group()
    if score is None:
        return 0.0
    return float(score)



