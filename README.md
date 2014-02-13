bloxorz
=======

A bloxorz solver written in Go.

[Original game](http://www.coolmath-games.com/0-bloxorz/)

![](https://raw.github.com/YouriAckx/bloxorz/master/bloxorz.jpg)

### Demo

```
youri@foobar bloxorz> go test
terrain1 first solution in 2 moves: [(0, 0); (1, 0)-(2, 0); (3, 0)]
terrain2 first solution in 4 moves: [(0, 3); (0, 1)-(0, 2); (0, 0); (1, 0)-(2, 0); (3, 0)]
(...)
infinite first solution in 7 moves: [(0, 0); (0, -2)-(0, -1); (1, -2)-(1, -1); (2, -2)-(2, -1);
(2, 0); (3, 0)-(4, 0); (3, 1)-(4, 1); (5, 1)]
terrain1 detailed solution ok
terrain1 only one solution ok
unsolvable terrain has no solution
PASS
```

### License

[GNU General Public License version 3](http://www.gnu.org/licenses/gpl.html)