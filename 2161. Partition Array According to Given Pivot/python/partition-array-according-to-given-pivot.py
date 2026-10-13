import pytest


class Solution:
    def pivotArray(self, nums: list[int], pivot: int) -> list[int]:
        leftSide = []
        midSide = []
        rightSide = []

        for n in nums:
            if n < pivot:
                leftSide.append(n)
            elif n > pivot:
                rightSide.append(n)
            else:
                midSide.append(n)

        return leftSide + midSide + rightSide


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([9,12,5,10,14,3,10],10), [9,5,3,10,10,12,14]),
])
def test_solution(sol, args, expected):
    assert sol.pivotArray(*args) == expected
