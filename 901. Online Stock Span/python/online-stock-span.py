import pytest


class StockSpanner:
    def __init__(self):
        self.history = [] # (span, price)

    def next(self, price: int) -> int:
        span = 1
        while self.history and self.history[-1][1] <= price:
            span += self.history[-1][0]
            self.history.pop()

        self.history.append((span, price))
        return span


# Your StockSpanner object will be instantiated and called as such:
# obj = StockSpanner()
# param_1 = obj.next(price)

@pytest.fixture
def sol():
    return StockSpanner()


@pytest.mark.parametrize(
    "prices, expected",
    [
        ([100, 80, 60, 70, 60, 75, 85], [1, 1, 1, 2, 1, 4, 6]),
    ],
)
def test_solution(sol, prices, expected):
    assert [sol.next(p) for p in prices] == expected
