package main
import "fmt"
func main() {
    fmt.Print("\x1b[2J\x1b[H\x1b[1;32m");
    fmt.Println(" ! x Q @ m 9 X    # 8 W v z    [ 7 k M p    $ 2 L K r    a o 1 4 9 ");
    fmt.Println(" 4   7   $   #    9   2   X    v   z   k    M   p   L    K   r   a   o ");
    fmt.Println(" \x1b[1;37m9\x1b[1;32m   x   Q   @    m   9   X    #   8   W    v   z   [    7   k   M    p   $   2 ");
    fmt.Println(" L   K   r   a    o   1   4    9   4   7    $   #   9    2   X   v    z   [ ");
    fmt.Println(" 7   \x1b[1;37mk\x1b[1;32m   M   p    $   2   L    K   r   a    o   1   4    9   4   7    $   #   9   2 ");
    fmt.Println(" X   v   z   [    7   k   M    p   $   2    L   K   r    a   o   1    4   9   4 ");
    fmt.Println(" 7   $   #   9    2   X   v    z   [   7    k   M   p    $   \x1b[1;37m2\x1b[1;32m   L    K   r   a ");
    fmt.Println(" o   1   4   9    4   7   $    #   9   2    X   v   z    [   7   k    M   p   $ ");
    fmt.Println(" 2   L   K   r    a   o   1    4   9   4    7   $   #    9   2   X    v   z   [ ");
    fmt.Println(" 7   k   M   p    $   2   L    K   \x1b[1;37mr\x1b[1;32m   a    o   1   4    9   4   7    $   #   9   2 ");
    fmt.Println(" X   v   z   [    7   k   M    p   $   2    L   K   r    a   o   1    4   9   4 ");
    fmt.Print("\x1b[0m");
}
