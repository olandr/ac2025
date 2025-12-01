## Day 01

sample_data = [
"L68",
"L130",
"R148",
"L5",
"R260",
"L55",
"L100",
"L99",
"L100",
"R100",
"R314",
"L82",
]

data = []
with open("./input.txt", "r+") as f:
  data = [line.rstrip('\n') for line in f]


def solution_01(data):
  count = 0
  val = 50
  for x in data:
    num = int(x[1:]) if x[0] == "R" else 100-int(x[1:])
    val = (val+num)%100
    if val == 0:
      count += 1
  return(count)

def solution_02(data):
  count = 0
  val = 50
  for x in data:
    count += int(x[1:])//100
    num = int(x[1:])%100

    if x[0] == "R":
      if val+num >= 100 and val != 0:
        count += 1
    else:
      if val-num <= 0 and val != 0:
        count += 1

    val = (val+num)%100 if x[0] == "R" else (val+100-num)%100

   # print(x,val,count) 
  return(count)

#print("Solution 1 test: ",solution_01(sample_data))
#print("Solution 1: ",solution_01(data))  ## Solution 1:  1105

#print("Solution 2 test: ",solution_02(sample_data))
print("Solution 2: ",solution_02(data)) ## Solution 2:  6599
