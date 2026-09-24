import pytest

from collections import Counter
from math import sqrt


class Solution:
    def repairCars(self, ranks: list[int], cars: int) -> int:
        freq = Counter(ranks)
        left = 1
        right = min(ranks) * cars ** 2

        def canRepairAllCars(t: int):
            total_cars_needed = 0
            for r, f in freq.items():
                total_cars_needed += f*(sqrt(t/r)//1)
                if total_cars_needed >= cars:
                    return True
            return False

        while left <= right:
            mid = (left + right) // 2
            if canRepairAllCars(mid):
                right = mid - 1
            else:
                left = mid + 1
        return left

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([5,1,6], 6), 16),
])
def test_solution(sol, args, expected):
    assert sol.repairCars(*args) == expected
