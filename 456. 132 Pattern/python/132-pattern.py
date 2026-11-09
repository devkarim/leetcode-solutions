import pytest


class Solution:
    def find132pattern(self, nums: list[int]) -> bool:
        stack = [] #[n, min_left]
        cur_min = nums[0]
        for j in range(1, len(nums)):
            while stack and nums[j] >= stack[-1][0]:
                stack.pop()
            if stack and stack[-1][1] < nums[j]:
                return True
            stack.append((nums[j], cur_min))
            cur_min = min(cur_min, nums[j])
        return False

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([3,1,4,2],), True),
    (([1,0,1,-4,-3],), False),
])
def test_solution(sol, args, expected):
    assert sol.find132pattern(*args) == expected
