import pytest


class Solution:
    def bagOfTokensScore(self, tokens: list[int], power: int) -> int:
        tokens.sort()

        curr_score = max_score = 0
        l, r = 0, len(tokens) - 1

        while l <= r:
            if power >= tokens[l]:
                curr_score += 1
                power -= tokens[l]
                l += 1
                max_score = max(max_score, curr_score)
            elif curr_score > 0:
                curr_score -= 1
                power += tokens[r]
                r -= 1
            else:
                break
        return max_score


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([100,200,300,400],200), 2),
])
def test_solution(sol, args, expected):
    assert sol.bagOfTokensScore(*args) == expected
