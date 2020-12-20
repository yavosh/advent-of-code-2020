from collections import deque


# Function to generate power set of given set S
def findPowerSet(S, set, n):

    # if we have considered all elements
    if n == 0:
        print(set)
        return

    # consider nth element
    set.append(S[n - 1])
    print("consider", set)
    findPowerSet(S, set, n - 1)

    # or don't consider nth element
    set.pop()
    print("don't consider", set)
    findPowerSet(S, set, n - 1)


if __name__ == '__main__':

    S = [1, 2, 3]

    set = []
    findPowerSet(S, set, len(S))