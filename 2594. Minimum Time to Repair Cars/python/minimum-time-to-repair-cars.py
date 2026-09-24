import pytest


class Solution:
    def repairCars(self, ranks: list[int], cars: int) -> int:
        left = 1
        right = min(ranks) * cars ** 2

        while left <= right:
            mid = (left + right) // 2
            total_cars_needed = 0
            for r in ranks:
                total_cars_needed += int((mid/r) ** 0.5)
            if total_cars_needed >= cars:
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
