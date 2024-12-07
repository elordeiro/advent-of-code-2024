mat = []

with open("../input.txt") as f:
    for line in f:
        mat.append(line.strip())

n, m = len(mat), len(mat[0])

total = 0

for r in range(1, n-1):
    for c in range(1, m-1):
        if mat[r][c] != 'A':
            continue
        str = mat[r-1][c-1] + 'A' + mat[r+1][c+1]
        if str != 'MAS' and str != 'SAM':
            continue
        str = mat[r+1][c-1] + 'A' + mat[r-1][c+1]
        if str != 'MAS' and str != 'SAM':
            continue
        total += 1

print(total)