from collections import defaultdict

import pytest


class TimeMap:

    def __init__(self):
        self.store = defaultdict(list)

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.store[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        if key not in self.store:
            return ""

        arr = self.store[key]
        l = 0
        r = len(arr) - 1
        while l <= r:
            m = (l + r) // 2
            curr_timestamp = arr[m][0]
            if curr_timestamp == timestamp:
                return arr[m][1]
            if curr_timestamp < timestamp:
                l = m + 1
            else:
                r = m - 1

        if l >= len(arr):
            return arr[-1][1]

        if r < 0:
            return ""

        return arr[r][1]


# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)

@pytest.fixture
def sol():
    return TimeMap()


def test_time_map(sol):
    sol.set("foo", "bar", 1)
    assert sol.get("foo", 1) == "bar"
    assert sol.get("foo", 3) == "bar"
    sol.set("foo", "bar2", 4)
    assert sol.get("foo", 4) == "bar2"
    assert sol.get("foo", 5) == "bar2"
