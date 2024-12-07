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
    fixed = False
    i = 0
    while i < len(update): 
        page = update[i]
        path[page] = i
        for neighbor in graph[page]:
            if neighbor in path and (j := path[neighbor]) < path[page]:
                update[i], update[j] = update[j], update[i]
                path[neighbor], path[page] = path[page], path[neighbor]
                fixed = True
                i = path[page]
        i += 1
    
    if fixed:
        total += update[len(update)//2]

print(total)