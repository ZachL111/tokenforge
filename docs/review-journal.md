# Review Journal

The review surface for `tokenforge` is deliberately narrow: one fixture, one scoring rule, and one local check.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its developer tools focus without claiming live deployment or external usage.

## Cases

- `baseline`: `change width`, score 214, lane `ship`
- `stress`: `diagnostic quality`, score 123, lane `watch`
- `edge`: `review cost`, score 164, lane `ship`
- `recovery`: `safe rewrite`, score 205, lane `ship`
- `stale`: `change width`, score 199, lane `ship`

## Note

The repository should be understandable without pretending it is larger than it is.
