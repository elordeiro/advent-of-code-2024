from collections import namedtuple
from enum import Enum

Dir = Enum('Dir', ['Up', 'Right', 'Down', 'Left'])
Step = Enum('Step', ['OutOfBounds', 'NewCell', 'Visited', 'Blocked'])
Cell = namedtuple('Cell', ['row', 'col'])

room = []
curr = Cell(0, 0)

def main():
    with open('../input.txt') as f:
        while (line := f.readline().strip()):
            room.append(list(line))
            row = len(room) - 1
            if (col := line.find('^')) != -1:
                curr = Cell(row, col)
                room[row][col] = 'X'
        
        visited_count = 1
        n, m = len(room), len(room[0])
        dir = Dir.Up

        def next_step(row, col):
            nonlocal curr
            if row < 0 or row >= n or col < 0 or col >= m:
                return Step.OutOfBounds
            if room[row][col] == '#':
                return Step.Blocked
            curr = Cell(row, col) 
            if room[row][col] == 'X':
                return Step.Visited
            room[row][col] = 'X'
            return Step.NewCell

        while True:
            step = 0
            match dir:
                case Dir.Up:
                    step = next_step(curr.row - 1, curr.col)
                    if step == Step.Blocked:
                        dir = Dir.Right
                case Dir.Right:
                    step = next_step(curr.row, curr.col + 1)
                    if step == Step.Blocked:
                        dir = Dir.Down
                case Dir.Down:
                    step = next_step(curr.row + 1, curr.col)
                    if step == Step.Blocked:
                        dir = Dir.Left
                case Dir.Left:
                    step = next_step(curr.row, curr.col - 1)
                    if step == Step.Blocked:
                        dir = Dir.Up

            if step == Step.OutOfBounds:
                break
            if step == step.NewCell:
                visited_count += 1  
    
    print(visited_count)


if __name__ == '__main__':
    main()