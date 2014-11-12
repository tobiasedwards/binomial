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

![Probability Formula](/img/probability_formula.png?raw=true)

To calculate the probability that *a <= X <= b*:

![img](/img/probability_cumulative_formula.png?raw=true)
