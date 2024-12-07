test_vals = []
num_Lists = []

with open("../input.txt") as file:
    for line in file:
        parts = line.strip().split(":")
        test_vals.append(int(parts[0]))
        num_Lists.append(list(map(int, parts[1].strip().split(" "))))

test_val = 0
num_List = []

def makes_true(acum, i):
    if acum == test_val and i == len(num_List):
        return True
    if acum > test_val or i == len(num_List):
        return False
    return makes_true(acum * num_List[i], i + 1) or \
            makes_true(acum + num_List[i], i + 1) or \
            makes_true(int(str(acum) + str(num_List[i])), i + 1)

total = 0
for i in range(len(test_vals)):
    test_val = test_vals[i]
    num_List = num_Lists[i]
    if makes_true(num_List[0], 1):
        total += test_val

print(total)