import pytest


class Solution:
    def search(self, nums: list[int], target: int) -> bool:
        l = 0
        r = len(nums) - 1
        while l <= r:
            while l <= r and nums[l] == nums[r] and r != l:
                l += 1
            m = (l + r) // 2
            if nums[m] == target:
                return True
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
        return False



@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([2,5,6,0,0,1,2], 0), True),
])
def test_solution(sol, args, expected):
    assert sol.search(*args) == expected
