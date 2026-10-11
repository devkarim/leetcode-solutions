import pytest


class Solution:
    def compress(self, chars: list[str]) -> int:
        i, res = 0, 0
        while i < len(chars):
            group_length = 1
            # compute how many the group contains
            while (i + group_length < len(chars) and chars[i] == chars[i + group_length]):
                group_length += 1
            chars[res] = chars[i]
            res += 1
            # if there are repeated characters, append the group length str
            if group_length > 1:
                grp_length_str = str(group_length)
                chars[res:res+len(grp_length_str)] = list(grp_length_str)
                res += len(grp_length_str)
            i += group_length
        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    ([["a", "a", "b", "b", "c", "c", "c"]], 6),
])
def test_solution(sol, args, expected):
    assert sol.compress(*args) == expected

