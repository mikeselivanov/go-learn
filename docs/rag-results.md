# RAG Query Log

## how to wrap an error with context in Go
_2026-08-06 09:26 UTC — collection: go_learn_docs_

**1.** `docs/errors.md` (distance 0.626)
> Package errors implements functions to manipulate errors. The New function creates errors whose only content is a text message. An error e wraps another error if e's type has one of the methods: Unwrap() error or Unwrap() []error. An easy way to create wrapped errors is to call fmt.Errorf and apply the %w verb to the error argument. Successive unwrapping of an error creates a tree. The Is and As functions inspect an error's tree by examining first the error itself followed by the tree of each of its children in turn (pre-order, depth-first traversal).

**2.** `docs/fmt.md` (distance 1.242)
> func Errorf(format string, a ...any) (err error) - Formats according to a format specifier and returns the string as an error value. Supports %w to wrap an error.

**3.** `docs/error_improvements.md` (distance 1.675)
> main.go:31-34 — body, urls, err := fetcher.Fetch(url); if err != nil { fmt.Printf(...); return } — Ошибка не оборачивается и не возвращается — она теряется внутри Crawl. Предложение: изменить сигнатуру Crawl на func Crawl(url string, depth int, fetcher Fetcher, cache *UrlCache) error и оборачивать ошибку контекстом: return fmt.Errorf("crawl %s: %w", url, err).

---
