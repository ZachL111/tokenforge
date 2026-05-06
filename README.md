# tokenforge

`tokenforge` is a compact Go repository for developer tools, centered on this goal: Generate a table-driven lexer from ordered token rules.

## Purpose

The point is to make a small domain rule concrete enough that a reader can change it and immediately see what broke.

## Tokenforge Review Notes

The first comparison I would make is `change width` against `diagnostic quality` because it shows where the rule is most opinionated.

## What Is Covered

- `fixtures/domain_review.csv` adds cases for change width and diagnostic quality.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/tokenforge-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `change width` and `diagnostic quality`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Implementation Notes

The core code exposes a scoring path and the added review layer uses `signal`, `slack`, `drag`, and `confidence`. The domain terms are `change width`, `diagnostic quality`, `review cost`, and `safe rewrite`.

The Go addition stays small enough to inspect in one sitting.

## Command

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Audit Path

The check exercises the source code and the review fixture. `baseline` is the high score at 214; `stress` is the low score at 123.

## Limits

The fixture set is small enough to audit by hand. The next useful expansion is malformed input coverage, not extra surface area.
