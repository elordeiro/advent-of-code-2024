mat = []

with open("../input.txt") as file:
    while (line := file.readline()):
        mat.append(line.strip())

n, m = len(mat), len(mat[0])

xmas_count = 0
for row in range(n):
    for col in range(m):
        if col + 3 < m and mat[row][col:col+4] == 'XMAS':
            xmas_count += 1
        if col - 3 >= 0 and mat[row][col-3:col+1] == 'SAMX':
            xmas_count += 1
        if row + 3 < n and [mat[row+i][col] for i in range(4)] == ['X', 'M', 'A', 'S']:
            xmas_count += 1
        if row - 3 >= 0 and [mat[row-i][col] for i in range(4)] == ['X', 'M', 'A', 'S']:
            xmas_count += 1
        if row + 3 < n and col + 3 < m and [mat[row+i][col+i] for i in range(4)] == ['X', 'M', 'A', 'S']:
            xmas_count += 1
        if row + 3 < n and col - 3 >= 0 and [mat[row+i][col-i] for i in range(4)] == ['X', 'M', 'A', 'S']:
            xmas_count += 1
        if row - 3 >= 0 and col + 3 < m and [mat[row-i][col+i] for i in range(4)] == ['X', 'M', 'A', 'S']:
            xmas_count += 1
        if row - 3 >= 0 and col - 3 >= 0 and [mat[row-i][col-i] for i in range(4)] == ['X', 'M', 'A', 'S']:
            xmas_count += 1

print(xmas_count)