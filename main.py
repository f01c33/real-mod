import math

umul = lambda x,w: math.floor(x/w)*w
omul = lambda x,w: math.ceil(x/w)*w

prevOp = lambda D,U: lambda x,w: U(math.floor(D(x,w)),w)
nextOp = lambda D,U: lambda x,w: U(math.ceil(D(x,w)),w)

x,w = 7,3
print(x,w,umul(x,w),omul(x,w))

Umul = prevOp(lambda a,b: a/b, lambda a,b: a*b)
Omul = nextOp(lambda a,b: a/b, lambda a,b: a*b)
print(x,w,Umul(x,w),Omul(x,w))

Uexp = prevOp(lambda a,b: math.log(a,b),lambda a,b: math.pow(a,b))
Oexp = nextOp(lambda a,b: math.log(a,b),lambda a,b: math.pow(a,b))
print(x,w,Uexp(x,w),Oexp(x,w))

Mod = lambda x,w: x+omul(x,w)-umul(x+omul(x,w),w)

inlineMod = lambda x,w: x+math.ceil(x/w)*w-math.floor((x+math.ceil(x/w)*w)/w)*w

for i in range(-10,10,1):
    print(Mod(i,w))
    # print(inlineMod(i,w))