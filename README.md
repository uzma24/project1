# project1
Service that accepts input as text and provides Json Output as Top ten most used words and times of occurrence in the text

How to run the project:
Clone the repo in your local 
Run main.go file(use command from terminal: go run main.go)
use this curl request:``` curl --location --request GET 'http://localhost:8080/text' \
--header 'Content-Type: application/json' \
--data-raw '{"text": "the uzma uzma uzma rafe the a an ae empty english hindi urdu urdu urdu english"}' ```

the API will return 10 most frequent words with frequency in JSON format


