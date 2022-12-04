from functools import wraps
import time

def timeit(func):
    """Decorator to calculate time taken by a function to run"""
    @wraps(func)
    def timed(*args, **kw):
        start = time.time()
        result = func(*args, **kw)
        end = time.time()

        format_str = 'function "{f}" took {took:.3f} seconds'
        message = format_str.format(f=func.__name__, took=end - start)

        # Use logger if provided        
        if kw.get('logger') is not None:
            kw.get('logger').info(message)
        else:
            print(message)

        return result
    return timed

@timeit
def part1():
    with open("input.txt") as fin:
        lines = fin.read().strip().split()

    ans = 0
    for line in lines:
        elves = line.split(",")
        ranges = [list(map(int, elf.split("-"))) for elf in elves]

        a, b = ranges[0]
        c, d = ranges[1]

        if a <= c and b >= d or a >= c and b <= d:
            ans += 1
    return ans  

@timeit
def part2():
    with open("input.txt") as fin:
        lines = fin.read().strip().split()
    
    ans = 0
    for line in lines:
        elves = line.split(",")
        ranges = [list(map(int, elf.split("-"))) for elf in elves]

        a, b = ranges[0]
        c, d = ranges[1]

        if not (b < c or a > d):
            ans += 1
    return ans

def main():
    print(part1())
    print(part2())  

if __name__ == '__main__':
    main()
