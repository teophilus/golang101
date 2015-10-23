# Problem Statement

The Utopian Tree goes through 2 cycles of growth every year. The first growth cycle occurs during the spring, when it doubles in height. The second growth cycle occurs during the summer, when its height increases by 1 meter.

Now, a new Utopian Tree sapling is planted at the onset of spring. Its height is 1 meter. Can you find the height of the tree after N growth cycles?

### Scope

Write a program outputs the height of the tree after 0 cycles, 1 cycle and 4 cycles.

### Sample results
```
Cycles -> Result
0 -> 1
1 -> 2
4 -> 7
```

### Explanation

In the first sample (N=0), the initial height (1) of the tree remains unchanged.

In the second sample (when N = 1, i.e. after the 1st cycle), the tree doubles its height as it's planted at the onset of spring.

In the third sample (N=4), the tree first doubles its height (2), then grows a meter (3), then doubles again (6), before growing another meter; at the end of the 4th cycle, its height is 7 meters.