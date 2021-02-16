
# Inverted Index API


## Introduction

An application that implements an efficient search functionality using inverted index. Rather than map a document object to an array of the terms it contains, inverted index maps a term to an array of the documents that contains it. In other words the inverted index contains the frequencies of each words in each document.


## Getting started
### Build and run locally 

```bash
git clone https://github.com/tsvetinaSpasova/xedni.git 
```

```bash
cd xedni
```

```bash
make go-run
```

### To create index from a document 
```bash
cd examples
```

```bash
./index.sh
```

### To search index for terms
```bash
cd examples
```

```bash
./search.sh
```

### Run tests
```bash
make go-test
```

### Run benchmarks
```bash
make go-bench
```

## Advantage of Inverted Index
 - Inverted index is to allow fast full text searches, at a cost of increased processing when a document is added to the database.

## Useful links
- [A first take at building an inverted index](https://nlp.stanford.edu/IR-book/html/htmledition/a-first-take-at-building-an-inverted-index-1.html)
- [18 3 The Inverted Index Stanford NLP Professor Dan Jurafsky & Chris Manning YouTube](https://www.youtube.com/watch?v=bnP6TsqyF30&ab_channel=AdityaAmbasth)
- [18 4 Query Processing with the Inverted Index Stanford NLP Dan Jurafsky & Chris Manning YouTub](https://www.youtube.com/watch?v=B-e297yK50U&ab_channel=AdityaAmbasth)

## References
Special thanks

- [https://github.com/enoti-bg/go-template-edge](https://github.com/enoti-bg/go-template-edge)