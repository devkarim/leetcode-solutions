import random

import pytest


class Solution:
    def __init__(self, w: list[int]):
        self.w = w
        self.sum = 0
        self.prefix_sum = []

        for weight in self.w:
            self.sum += weight
            self.prefix_sum.append(self.sum)

    def pickIndex(self) -> int:
        rand = random.randrange(1, self.sum+1)

        l = 0
        r = len(self.w) - 1

        while l <= r:
            m = (l + r) // 2
            if rand > self.prefix_sum[m]:
                l = m + 1
            else:
                r = m - 1

        return l


# Your Solution object will be instantiated and called as such:
# obj = Solution(w)
# param_1 = obj.pickIndex()


def test_single_weight():
    sol = Solution([1])
    assert sol.pickIndex() == 0


def test_pick_index_in_range():
    sol = Solution([1, 3])
    for _ in range(100):
        assert 0 <= sol.pickIndex() < 2


