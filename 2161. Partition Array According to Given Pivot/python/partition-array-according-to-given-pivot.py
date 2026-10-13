import pytest


class Solution:
    def pivotArray(self, nums: list[int], pivot: int) -> list[int]:
        lessNum = 0
        equalNum = 0

        for n in nums:
            if n < pivot:
                lessNum += 1
            elif n == pivot:
                equalNum += 1

        ans = [0] * len(nums)
        lessIdx = 0
        equalIdx = lessNum
        greaterIdx = lessNum + equalNum

        for n in nums:
            if n < pivot:
                ans[lessIdx] = n
                lessIdx += 1
            elif n > pivot:
                ans[greaterIdx] = n
                greaterIdx += 1
            else:
                ans[equalIdx] = n
                equalIdx += 1

        return ans

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([9,12,5,10,14,3,10],10), [9,5,3,10,10,12,14]),
])
def test_solution(sol, args, expected):
    assert sol.pivotArray(*args) == expected
