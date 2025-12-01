## Day 01

sample_data = [
"L68",
"L30",
"R48",
"L5",
"R60",
"L55",
"L1",
"L99",
"R14",
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
  return(0)

#print("Solution 1 test: ",solution_01(sample_data))
print("Solution 1: ",solution_01(data))  ## Solution 1:  xxx

#print("Solution 2 test: ",solution_02(sample_data))
#print("Solution 2: ",solution_02(data)) ## Solution 2:  xxx
