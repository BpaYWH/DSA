# Math Properties

### Modulus

- (a + b) % c = ((a % c) + (b % c)) % c | [CF682A]

### Right Angled Triangle

- a^2 + b^2 = c^2 | [CF707C]
  - let a be the shortest side, b < c
  - if a is odd, then b = (a^2 - 1) / 2, c = (a^2 + 1) / 2
  - if a is even, then b = a^2 / 4 - 1, c = a^2 / 4 + 1
  - note that there is no solution for a <= 2
