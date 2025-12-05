## Day 05

sample_data = [
  "3-5",
  "10-14",
  "16-20",
  "12-18"
]

sample_ids=[1,5,8,11,17,32]

data = []
with open("./input.txt", "r+") as f:
  data = [line.rstrip('\n') for line in f]

ids = []
with open("./input2.txt", "r+") as f:
  ids = [int(line.rstrip('\n')) for line in f]

def prep_data(data):
  ranges=[]
  for line in data:
    nums=line.split("-")
    ranges.append((int(nums[0]),int(nums[1])))
  ranges = sorted(ranges,key=lambda x: x[0])
  min_ids=[x[0] for x in ranges]
  max_ids=[x[1] for x in ranges]
  return(min_ids,max_ids)

def solution_01(data,ids):
  
  count=0
  min_ids,max_ids=prep_data(data)

  for i in ids:
    idx=0
    while(idx<len(min_ids) and min_ids[idx]<=i):
      if max_ids[idx]>=i:
        count+=1
        break
      idx+=1

  return(count)

def solution_02(data):
  count=0
  min_ids,max_ids=prep_data(data)

  idx=0
  min_id=min_ids[0]
  max_id=max_ids[0]
  while True:
    if min_ids[idx+1] > max_id:
      count+=max_id-min_id+1
      min_id=min_ids[idx+1]
      max_id=max_ids[idx+1]
    else:
      max_id=max(max_id,max_ids[idx+1])
    idx+=1
    if idx == len(min_ids)-1:
      #final round
      count+=max_id-min_id+1
      break
  return(count)


#print("Solution 1 test: ",solution_01(sample_data,sample_ids))
print("Solution 1: ",solution_01(data,ids))  ## Solution 1:  789

#print("Solution 2 test: ",solution_02(sample_data))
print("Solution 2: ",solution_02(data)) ## Solution 2:  343329651880509
