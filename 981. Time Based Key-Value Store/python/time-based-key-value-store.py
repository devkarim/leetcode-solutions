from collections import defaultdict

import pytest


class TimeMap:
    def __init__(self):
        self.store = defaultdict(list)

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.store[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        res = ""
        values = self.store.get(key, [])
        l, r = 0, len(values) - 1
        while l <= r:
            m = (l + r) // 2
            curr_timestamp = values[m][0]
            if curr_timestamp == timestamp:
                return values[m][1]
            if curr_timestamp < timestamp:
                res = values[m][1]
                l = m + 1
            else:
                r = m - 1

        return res


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
