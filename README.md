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

![equation](http%3A%2F%2Fwww.sciweavers.org%2Ftex2img.php%3Feq%3DP%2528X%253Dx%2529%2520%253D%2520%255Cbegin%257Bpmatrix%257Dn%255C%255Cx%255Cend%257Bpmatrix%257D%2520p%255Ex%2520%25281-p%2529%255Ex%26bc%3DWhite%26fc%3DBlack%26im%3Djpg%26fs%3D12%26ff%3Darev%26edit%3D0)

To calculate the probability that *a <= X <= b*:

![equation](http%3A%2F%2Fwww.sciweavers.org%2Ftex2img.php%3Feq%3DP%2528a%2520%255Cle%2520X%2520%255Cle%2520b%2529%2520%253D%2520P%2528X%253Da%2529%2520%252B%2520P%2528X%253Da%2520%252B%25201%2529%2520%252B%2520...%2520%252B%2520P%2528X%253Db%2529%26bc%3DWhite%26fc%3DBlack%26im%3Djpg%26fs%3D12%26ff%3Darev%26edit%3D0)
