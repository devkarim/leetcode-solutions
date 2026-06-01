import pytest


class Solution(object):
    def minAddToMakeValid(self, s):
        """
        :type s: str
        :rtype: int
        """
        min_add = 0
        open_brackets = 0

        for c in s:
            if c == '(':
                open_brackets += 1
            else:
                if open_brackets > 0:
                    open_brackets -= 1
                else:
                    min_add += 1

        return min_add + open_brackets


@pytest.fixture
def sol():
    return Solution()


def test_minAddToMakeValid(sol: Solution):
    assert sol.minAddToMakeValid("())") == 1
    assert sol.minAddToMakeValid("(((") == 3
