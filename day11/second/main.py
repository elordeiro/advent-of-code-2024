def main():
    rocks = open("../input.txt").read().split()
    
    dp = {}
    
    def blink(rock, depth):
        if (rock, depth) in dp:
            return dp[(rock, depth)]
        if not rock:
            return 0
        if not depth:
            return 1

        left, right = "", ""
        if rock == "0":
            left = "1"
        elif (n := len(rock)) % 2 == 0:
            left = rock[:n//2]
            right = str(int(rock[n//2:]))
        else:
            left = str(int(rock) * 2024)
        
        dp[(rock, depth)] = blink(left, depth-1) + blink(right, depth-1)
        return dp[(rock, depth)]

    count = 0
    for rock in rocks:
        count += blink(rock, 75)
    
    print(count)


if __name__ == "__main__":
    main()