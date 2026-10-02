import pytest


class Solution:
    def suggestedProducts(
        self, products: list[str], searchWord: str
    ) -> list[list[str]]:
        products.sort()
        ans, prefix, initial_l = [], "", 0
        for w in searchWord:
            prefix += w
            l = initial_l
            r = len(products) - 1
            while l <= r:
                m = (l + r) // 2
                p = products[m]
                if prefix > p:
                    l = initial_l = m + 1
                else:
                    r = m - 1
            matches = products[l : l + 3]
            ans.append([p for p in matches if p.startswith(prefix)])
        return ans


@pytest.fixture
def sol():
    return Solution()


@pytest.mark.parametrize(
    "args, expected",
    [
        (
            (["mobile", "mouse", "moneypot", "monitor", "mousepad"], "mouse"),
            [
                ["mobile", "moneypot", "monitor"],
                ["mobile", "moneypot", "monitor"],
                ["mouse", "mousepad"],
                ["mouse", "mousepad"],
                ["mouse", "mousepad"],
            ],
        ),
    ],
)
def test_solution(sol, args, expected):
    assert sol.suggestedProducts(*args) == expected
