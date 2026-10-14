import pytest


class Solution:
    def fourSum(self, nums: list[int], target: int) -> list[list[int]]:
        nums = sorted(nums)
        res = []

        a = 0
        while a < len(nums) - 3:
            if a > 0 and nums[a] == nums[a - 1]:
                a += 1
                continue
            b = a + 1
            while b < len(nums) - 2:
                c = b + 1
                d = len(nums) - 1
                while c < d:
                    total = nums[a] + nums[b] + nums[c] + nums[d]
                    if total == target:
                        res.append([nums[a], nums[b], nums[c], nums[d]])
                        c += 1
                        d -= 1
                        while c < d and nums[c] == nums[c - 1]:
                            c += 1
                    elif total < target:
                        c += 1
                    else:
                        d -= 1
                b += 1
                while b < len(nums) and nums[b] == nums[b - 1]:
                    b += 1

            a += 1

        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([1, 0, -1, 0, -2, 2], 0), [[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]]),
])
def test_solution(sol, args, expected):
    assert sol.fourSum(*args) == expected
