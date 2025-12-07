## Day 07

sample_data = [
".......S.......",
"...............",
".......^.......",
"...............",
"......^.^......",
"...............",
".....^.^.^.....",
"...............",
"....^.^...^....",
"...............",
"...^.^...^.^...",
"...............",
"..^...^.....^..",
"...............",
".^.^.^.^.^...^.",
"..............."
]

data = []
with open("./input.txt", "r+") as f:
  #data = [list(map(int, line.rstrip('\n'))) for line in f]
  data = [line.rstrip('\n') for line in f]


def solution_01(data):
  count=0
  idx_lst=set()
  for row in data:
    for idx,value in enumerate(row):
      if value == "S":
        idx_lst.add(idx)
      elif value == "^" and idx in idx_lst:
        count+=1
        idx_lst.remove(idx)
        idx_lst.add(idx-1)
        idx_lst.add(idx+1)

  return(count)

def solution_02(data):
  count=0
  idx_lst=set()
  count_lst={}
  for row in data:
    for idx,value in enumerate(row):
      if value == "S":
        idx_lst.add(idx)
        count_lst[str(idx)]=1
      elif value == "^" and count_lst.get(str(idx)):

        if count_lst.get(str(idx-1)):
          count_lst[str(idx-1)]=count_lst[str(idx-1)]+count_lst[str(idx)]
        else:
          count_lst[str(idx-1)]=count_lst[str(idx)]

        if count_lst.get(str(idx+1)):
          count_lst[str(idx+1)]=count_lst[str(idx+1)]+count_lst[str(idx)]
        else:
          count_lst[str(idx+1)]=count_lst[str(idx)]

        count_lst.pop(str(idx))
  
  return(sum(count_lst.values()))


print("Solution 1 test: ",solution_01(sample_data))
#print("Solution 1: ",solution_01(data))  ## Solution 1:  1613

#print("Solution 2 test: ",solution_02(sample_data))
print("Solution 2: ",solution_02(data)) ## Solution 2:  48021610271997
