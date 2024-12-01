llist = [0] * 1000
rlist = [0] * 1000
counts = [] 

with open("../input.txt") as file:
    i = 0
    while (line := file.readline()) and line:
        llist[i], rlist[i] = map(int, line.split())
        i += 1

for num in llist:
    counts.append(rlist.count(num) * num)

print(sum(counts))