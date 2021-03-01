import random

def getreq():
    for i in range(len(rq)):
        if(rq[i]!=0):
            rq[i]=rq[i]-1
            return i


def do(param):
    if (rq[param] == 0):
        return colors[getreq()]
    else:
        rq[param] = rq[param] - 1
        return colors[param]


def color():
    v=random.randint(0,level-2)
    v=colors[v]
    if(v=='red'):
        return do(0)
    elif(v=='yellow'):
        return do(1)
    elif(v=='green'):
        return do(2)
    else:
        return do(3)


def checkwin():
    c=None
    for i in s:
        try:
            c=i[0]
        except:
            continue
        for j in i:
            if(j!=c):
                return False
    return True



def get(level):
    k=level-1
    for i in range(k):
        rq.append(k)
    for i in range(k):
        t=[]
        for j in range(k):
            t.append(color())
        s.append(t)
    s.append([])
    print(rq)


def getlevel():
    if(level==3):
        return 'Basic'
    if(level==4):
        return 'Medium'
    if(level==5):
        return 'Hard'



play=True
level=2
colors = ['red', 'yellow', 'green', 'blue']
while(play):
    s=[]
    rq=[]
    level=level+1
    if(level>5):
        print('Max Level Reached')
        break
    print('level :-', getlevel())
    get(level)
    print(s)
    while(True):
        print('\n Current State Of Stacks')
        print(s)
        f=int(input('Enter between :- '+str(level)))
        t=int(input('To Stack must except stack :- '+str(f)))
        e=s[f-1].pop()
        s[t-1].append(e)
        if(checkwin()):
            print(s)
            break
        