# Inverted Index API


## Introduction

An application that implements an efficient search functionality using inverted index. Rather than map a document object to an array of the terms it contains, inverted index maps a term to an array of the documents that contains it. In other words the inverted index contains the frequencies of each words in each document.
It has the following features:
 - Allow users to create index from a JSON array of document objects
 - Allow users to search index for terms

## Steps to build an inverted index

1. Collect the documents to be indexed
2. Tokenize the text, turning each document into a list of tokens
3. Do linguistic preprocessing, producing a list of normalized tokens, which are the indexing terms
4. Index the documents that each term occurs in by creating an inverted index, consisting of a dictionary and postings. 

## Advantage of Inverted Index
 - Inverted index is to allow fast full text searches, at a cost of increased processing when a document is added to the database.

## Useful links
- [A first take at building an inverted index](https://nlp.stanford.edu/IR-book/html/htmledition/a-first-take-at-building-an-inverted-index-1.html)
- [18 3 The Inverted Index Stanford NLP Professor Dan Jurafsky & Chris Manning YouTube](https://www.youtube.com/watch?v=bnP6TsqyF30&ab_channel=AdityaAmbasth)
- [18 4 Query Processing with the Inverted Index Stanford NLP Dan Jurafsky & Chris Manning YouTub](https://www.youtube.com/watch?v=B-e297yK50U&ab_channel=AdityaAmbasth)
