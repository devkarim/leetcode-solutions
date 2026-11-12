import pytest
from collections import Counter


class FreqStack:
    def __init__(self):
        self.max_freq = 0
        self.freq = Counter()
        self.stacks = {}

    def push(self, val: int) -> None:
        self.freq[val] += 1
        if self.freq[val] not in self.stacks:
            self.stacks[self.freq[val]] = []
        self.stacks[self.freq[val]].append(val)
        self.max_freq = max(self.max_freq, self.freq[val])

    def pop(self) -> int:
        res = self.stacks[self.max_freq].pop()
        if not self.stacks[self.max_freq]:
            self.max_freq -= 1
        self.freq[res] -= 1
        return res


# Your FreqStack object will be instantiated and called as such:
# obj = FreqStack()
# obj.push(val)
# param_2 = obj.pop()


@pytest.fixture
def freq_stack():
    return FreqStack()


def test_freq_stack(freq_stack):
    for val in [5, 7, 5, 7, 4, 5]:
        freq_stack.push(val)
    assert freq_stack.pop() == 5
    assert freq_stack.pop() == 7
    assert freq_stack.pop() == 5
    assert freq_stack.pop() == 4


def test_freq_stack_interleaved(freq_stack):
    for val in [1, 1, 1, 2]:
        freq_stack.push(val)
    assert freq_stack.pop() == 1
    assert freq_stack.pop() == 1
    # [1,2,2,2,1]
    for val in [2, 2, 1]:
        freq_stack.push(val)
    assert freq_stack.pop() == 2
    assert freq_stack.pop() == 1
    assert freq_stack.pop() == 2
