# Clean Arch Mini Project

This service for testing purposes, to make it easier to connect to other libraries will be using clean architecture

## Documentation
### Clean Architecture Guide
https://github.com/bxcodec/go-clean-arch

this service will have
1. Repository
2. Usecase
3. Domain
4. Controller

every detail that we need in the process will be added later

### Sample Case
#### Booking System
This system will help to provide booking online

data sample will be provided by
https://openlibrary.org/dev/docs/api/subjects


the data from API will be breakdown like
```mermaid
    erDiagram
        book }o--o{ author : book_author
        book {
            string id
            string title
            int edition_count
        }
        author {
            string id
            string name
        }
        book }o--o{ subject : book_subject
        subject {
            int id
            string name
        }

        book }o--o{ booking : booking_book
        booking {
            int id
            string book_id
            int user_id
            int quantity
            timestamp book_at
            timestamp return_at
        }

        user }o--o{ booking : user_booking
        user {
            int id
            string name
        }
```

Booking process
```mermaid
    flowchart LR
        Start-->B{book exist?}
        subgraph ide1 [is available]
        B-- Yes-->C{quantity valid?}
        C-- Yes-->D{stock available}
        end
        D -- Yes --> CC[Booking Created ]
        B -- No --> DD[Return Error]
        C -- No --> DD
        D -- No --> DD
        CC -->Stop
        DD -->Stop

```