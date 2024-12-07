from collections import namedtuple
from enum import Enum

Dir = Enum('Dir', ['Up', 'Right', 'Down', 'Left'])
Step = Enum('Step', ['OutOfBounds', 'NewCell', 'Visited', 'Blocked', 'Loop'])
Cell = namedtuple('Cell', ['row', 'col'])


def main():
    room = []
    start = Cell(0, 0)
    
    with open('../input.txt') as f:
        while (line := f.readline().strip()):
            room.append(list(line))
            row = len(room) - 1
            if (col := line.find('^')) != -1:
                start = Cell(row, col)
                room[row][col] = Dir.Up
        
        room_copy = []
        cur = start
        n, m = len(room), len(room[0])
        dir = Dir.Up

        def next_step(row, col):
            nonlocal cur
            if row < 0 or row >= n or col < 0 or col >= m:
                return Step.OutOfBounds
            if room_copy[row][col] == '#':
                return Step.Blocked
            cur = Cell(row, col) 
            if room_copy[row][col] == '.':
                room_copy[row][col] = dir 
                return Step.NewCell
            if room_copy[row][col] == dir:
                return Step.Loop
            return Step.Visited

        loop_count = 0
        for row in range(n):
            for col in range(m):
                if room[row][col] != '.':
                    continue
                print(row, col)
                room_copy = [row.copy() for row in room]
                room_copy[row][col] = '#'
                dir = Dir.Up
                cur = start
                while True:
                    step = Step.Visited
                    match dir:
                        case Dir.Up:
                            step = next_step(cur.row - 1, cur.col)
                            if step == Step.Blocked:
                                dir = Dir.Right
                        case Dir.Right:
                            step = next_step(cur.row, cur.col + 1)
                            if step == Step.Blocked:
                                dir = Dir.Down
                        case Dir.Down:
                            step = next_step(cur.row + 1, cur.col)
                            if step == Step.Blocked:
                                dir = Dir.Left
                        case Dir.Left:
                            step = next_step(cur.row, cur.col - 1)
                            if step == Step.Blocked:
                                dir = Dir.Up

                    if step == Step.OutOfBounds:
                        break
                    if step == step.Loop:
                        loop_count += 1 
                        break
    
    print(loop_count)


if __name__ == '__main__':
    main()