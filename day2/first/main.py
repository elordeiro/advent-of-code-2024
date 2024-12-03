safe = 0

with open("../input.txt") as file:
    while (line := file.readline()):
        levels = line.split()
        levels = [int(level) for level in levels]

        is_decreasing = False
        if levels[0] > levels[1]:
            is_decreasing = True
        
        for i in range(1, len(levels)):
            if is_decreasing:
                if levels[i] >= levels[i-1] or levels[i] < levels[i-1] - 3:
                    break
            else:
                if levels[i] <= levels[i-1] or levels[i] > levels[i-1] + 3:
                    break
        else:
            safe += 1
    
print(safe)