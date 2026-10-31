import pytest


class Solution:
    def removeKdigits(self, num: str, k: int) -> str:
        stack = []
        for n in num:
            while k != 0 and stack and int(stack[-1]) > int(n):
                stack.pop()
                k -= 1
            if n != "0" or stack:
                stack.append(n)
        while stack and k > 0:
            stack.pop()
            k -= 1
        if not stack:
            return "0"
        return "".join(stack)


@pytest.fixture
def sol():
    return Solution()


@pytest.mark.parametrize(
    "args, expected",
    [
        (("1432219", 3), "1219"),
        (("1432239", 3), "1223"),
        (("10200", 1), "200"),
        (("10", 2), "0"),
        (("9", 1), "0"),
    ],
)
def test_solution(sol, args, expected):
    assert sol.removeKdigits(*args) == expected

# k = 1
# 1[4]32219 -> 132219

# k = 2
# 1[43]2219 -> 12219

# k = 3
# 1[432]219 -> 1219 | 1221
# 1000000 + 400000 + 30000 + 2000 + 200 + 10 + 9
#             ^        ^       ^

# ------

# k = 1
# 1[4]32269 -> 132269

# k = 2
# 1[43]2269 -> 12269

# k = 3
# 1[43]226[9] -> 1226 | 12[6]9
# 1000000 + 400000 + 30000 + 2000 + 200 + 60 + 9
#              ^       ^                       ^
