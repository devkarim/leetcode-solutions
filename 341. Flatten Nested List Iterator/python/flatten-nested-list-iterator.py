import pytest


class NestedInteger:
    def __init__(self, value):
        if isinstance(value, int):
            self._integer = value
            self._list = None
        else:
            self._integer = None
            self._list = [NestedInteger(v) for v in value]

    def isInteger(self) -> bool:
        return self._integer is not None

    def getInteger(self) -> int:
        return self._integer

    def getList(self) -> list:
        return self._list


class NestedIterator:
    def __init__(self, nestedList: list):
        self.stack = [[nestedList, 0]]

    def next(self) -> int:
        (l, idx) = self.stack[-1]
        self.stack[-1][1] += 1
        return l[idx].getInteger()

    def hasNext(self) -> bool:
        if not self.stack:
            return False

        (l, idx) = self.stack[-1]

        while len(l) > idx and not l[idx].isInteger():
            l = l[idx].getList()
            idx = 0
            self.stack.append([l, idx])

        while self.stack and len(self.stack[-1][0]) <= self.stack[-1][1]:
            self.stack.pop()
            if self.stack:
                self.stack[-1][1] += 1
                (l, idx) = self.stack[-1]

                while len(l) > idx and not l[idx].isInteger():
                    l = l[idx].getList()
                    idx = 0
                    self.stack.append([l, idx])

        return len(self.stack) > 0


# Your NestedIterator object will be instantiated and called as such:
# i, v = NestedIterator(nestedList), []
# while i.hasNext(): v.append(i.next())


def _build(nested):
    return [NestedInteger(v) for v in nested]


def _flatten(nested):
    it = NestedIterator(_build(nested))
    result = []
    while it.hasNext():
        result.append(it.next())
    return result


@pytest.mark.parametrize(
    "nested, expected",
    [
        ([[1, 1], 2, [1, 1]], [1, 1, 2, 1, 1]),
    ],
)
def test_nested_iterator(nested, expected):
    assert _flatten(nested) == expected
