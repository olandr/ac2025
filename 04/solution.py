## Day xx


sample_data = [
"..@@.@@@@.",
"@@@.@.@.@@",
"@@@@@.@.@@",
"@.@@@@..@.",
"@@.@@@@.@@",
".@@@@@@@.@",
".@.@.@.@@@",
"@.@@@.@@@@",
".@@@@@@@@.",
"@.@.@@@.@."
]

data = []
with open("./input.txt", "r+") as f:
  #data = [list(map(int, line.rstrip('\n'))) for line in f]
  data = [line.rstrip('\n') for line in f]


def solution_01(data):
  mat = [[0 for _ in range(len(data[1])+2)] for _ in range(len(data)+2)]
  mat_box = [[0 for _ in range(len(data[1])+2)] for _ in range(len(data)+2)]
  for i,row in enumerate(data):
    for j,value in enumerate(row):
      if value == "@":
        mat[i][j] +=1
        mat[i+1][j] +=1
        mat[i+2][j] +=1
        mat[i][j+1] +=1
        mat[i][j+2] +=1
        mat[i+1][j+2] +=1
        mat[i+2][j+1] +=1
        mat[i+2][j+2] +=1

        mat_box[i+1][j+1] = 1
  
  count = 0
  for i,row in enumerate(mat_box):
    for j,value in enumerate(row):
      if value and mat[i][j] < 4:
        count +=1
  return(count)

def solution_02(data):
  return(0)


print("Solution 1 test: ",solution_01(sample_data))
print("Solution 1: ",solution_01(data))  ## Solution 1:  xxx

#print("Solution 2 test: ",solution_02(sample_data))
#print("Solution 2: ",solution_02(data)) ## Solution 2:  xxx
