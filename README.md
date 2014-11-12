binomial
========

Command line tool to calculate probabilities using a binomial distribution

## Maths Principles

A binomial distribution is a discrete probability distribution where
*X* represents the number off successes in a sample of size *n* with a
probability if success *p*.

The probability of success *p* must be constant for all trials.

A binomial distribution could be used to model flipping a coin 10 times (Where
*X* represents the amount of heads and thus *p* = 0.5 and *n* = 10)

But would not be suitable to model 10 cards drawn from a deck (Where *X*
represents the red cards drawn). This would not be suitable since the
probability of drawing a red card changes depending on how many red cards have
already been dawn.

To calculate the probability that *X = x*:

![equation](http://www.sciweavers.org/tex2img.php?eq=P%28X%3Dx%29%20%3D%20%5Cbegin%7Bpmatrix%7Dn%5C%5Cx%5Cend%7Bpmatrix%7D%20p%5Ex%20%281-p%29%5Ex&bc=White&fc=Black&im=jpg&fs=12&ff=arev&edit=0)

To calculate the probability that *a <= X <= b*:

![equation](http://www.sciweavers.org/tex2img.php?eq=P%28a%20%5Cle%20X%20%5Cle%20b%29%20%3D%20P%28X%3Da%29%20%2B%20P%28X%3Da%20%2B%201%29%20%2B%20...%20%2B%20P%28X%3Db%29&bc=White&fc=Black&im=jpg&fs=12&ff=arev&edit=0)
