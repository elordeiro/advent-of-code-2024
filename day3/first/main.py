import re

string = ""
with open("../input.txt") as file:
    string = file.read()

matches = re.findall(r'mul\(\d{1,3},\d{1,3}\)', string)
matches = [int(a) * int(b) for a, b in [re.findall(r'\d{1,3}', match) for match in matches]]

print(sum(matches))