import pytest


class Solution:
    def search(self, nums: list[int], target: int) -> int:
        l = 0
        r = len(nums) - 1
        while l <= r:
            m = (l + r) // 2
            if nums[m] == target:
                return m
            if nums[m] >= nums[l]:
                if target > nums[m] or target < nums[l]:
                    l = m + 1
                else:
                    r = m - 1
            else:
                if target < nums[m] or target > nums[r]:
                    r = m - 1
                else:
                    l = m + 1
        return -1


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([4,5,6,7,0,1,2], 0), 4),
])
def test_solution(sol, args, expected):
    assert sol.search(*args) == expected
