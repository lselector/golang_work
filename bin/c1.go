package main
import "fmt"


a b c

(a+b)

(a+b) + (a+b+c)




func solveMeFirst(a uint32,b uint32) uint32{
  // Hint: Type return (a+b) below
    return (a+b)
}

func main() {
    var a, b, res uint32
    fmt.Print("Enter 1st number:")
    fmt.Scanf("%v", &a)
    fmt.Print("Enter 2nd number:")
    fmt.Scanf("%v", &b)
    res = solveMeFirst(a,b)
    fmt.Println(res)
}

def reductionCost(num):
l = len(num)
lmax = l-1
num2 = sorted(num)
res = num2[0]
res_total = 0
n = 1
while n<=lmax:
    res += num2[n]
    res_total += res
    n+=1
    return res_total

=====================================

kDifference
   array of "n" positive distinct  integers "a"
   an integer "k"

return - integer - total number of pair elements (ai,aj)
where ai+k=aj

n   n>5
a1  a<2mln
a2
a3
...
k  <2mln

def kd(a,k):
    aa=sorted(set(a))
    N=len(aa)
    res = 0
    for i1 in range(N):
        for i2 in range(i1,N):
            dif = aa[i2]-aa[i1]
            if dif == k:
                res += 1
            elif dif >k:
                break
    return res
    
=============================================
moves(a)
  i1 = 0
  i2 = l-1
  

