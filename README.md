# BS-618

## Problem

Our application allows the submition of questions, which can be answered by other users. It can happen that the same question is submitted multiple times, but we want to keep only a single instance of each question.

Your task is to filter out duplicate questions from an input list, and output a list of distinct questions ordered by their `id` (ascending). Uniqueness of questions is based on an exact match of the `content` property (case-sensitive, same whitespace, etc). If the question content matches, the question should be chosen based on the following rules:

* The question with the highest rated answer has priority.
* If both have the same answer rating, the older (`createTimestamp`) question has priority.

How would you make sure that your code works?


### Sample input
```json
[{
    "id": 123,
    "content": "Test content",
    "createTimestamp": 123213, 
    "answers": [{
        "id": 142,
        "rating": 10,
        "content": "Test answer"
    }, {
        "id": 242,
        "rating": 2,
        "content": "Test answer 2"
    }]
}, {
    "id": 1024,
    "content": "Test content",
    "createTimestamp": 54343, 
    "answers": [{
        "id": 454,
        "rating": 9,
        "content": "Test answer 2"
    }, {
        "id": 342,
        "rating": 4,
        "content": "Test answer 3"
    }]
},
{
    "id": 250,
    "content": "Different test content",
    "createTimestamp": 543431, 
    "answers": [{
        "id": 854,
        "rating": 10,
        "content": "Test answer 4"
    }, {
        "id": 346,
        "rating": 3,
        "content": "Test answer 5"
    }]
}]

```

### Sample ouptut
```json
[{
    "id": 123,
    "content": "Test content",
    "createTimestamp": 123213,
    "answers": [{
        "id": 142,
        "rating": 10,
        "content": "Test answer"
    }, {
        "id": 242,
        "rating": 2,
        "content": "Test answer 2"
    }]
},
{
    "id": 250,
    "content": "Different test content",
    "createTimestamp": 543431, 
    "answers": [{
        "id": 854,
        "rating": 10,
        "content": "Test answer 4"
    }, {
        "id": 346,
        "rating": 3,
        "content": "Test answer 5"
    }]
}]

```