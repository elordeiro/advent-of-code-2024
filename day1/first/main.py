llist = [0] * 1000
rlist = [0] * 1000

with open("../input.txt", "r") as file:
    i = 0
    while (line := file.readline()) and line:
        llist[i], rlist[i] = map(int, line.split())
        i += 1

llist.sort()
rlist.sort()

diff = [abs(left - right) for left, right in zip(llist, rlist)]
print(sum(diff))