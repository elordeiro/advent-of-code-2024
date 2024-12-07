from collections import defaultdict

graph = defaultdict(list) 
updates = []

with open('../input.txt') as f:
    while (line := f.readline().strip()):
        if not line:
            break
        src, dst = list(map(int, line.split('|')))
        graph[src].append(dst)
    
    while (line := f.readline().strip()):
        parts = list(map(int, line.split(',')))
        updates.append(parts)

total = 0
for update in updates:
    path = {}
    for page in update:
        path[page] = True if all(neighbor not in path for neighbor in graph[page]) else False
    
    if all(path[page] for page in update):
        total += update[len(update)//2]

print(total)