Feature: Book Info Service

Scenario: List all books
  Given the following books exist in the system:
    | title                | author            | genre  | 
    | War and Peace        | Leo Tolstoy       | Novel  | 
    | Crime and Punishment | Fyodor Dostoevsky | Novel  | 
    | Foundation           | Isaac Asimov      | Sci-Fi | 
    | Dune                 | Frank Herbert     | Sci-Fi | 
   When I retrieve the list of all books
   Then I should receive the following books:
    | title                | author            | genre  | 
    | War and Peace        | Leo Tolstoy       | Novel  | 
    | Crime and Punishment | Fyodor Dostoevsky | Novel  | 
    | Foundation           | Isaac Asimov      | Sci-Fi | 
    | Dune                 | Frank Herbert     | Sci-Fi | 
    And the response status code should be 200

Scenario: List books by genre
  Given the following books exist in the system:
    | title                | author            | genre  | 
    | War and Peace        | Leo Tolstoy       | Novel  | 
    | Crime and Punishment | Fyodor Dostoevsky | Novel  | 
    | Foundation           | Isaac Asimov      | Sci-Fi | 
    | Dune                 | Frank Herbert     | Sci-Fi | 
   When I retrieve the list of books by genre "Novel"
   Then I should receive the following books:
    | title                | author            | genre | 
    | War and Peace        | Leo Tolstoy       | Novel | 
    | Crime and Punishment | Fyodor Dostoevsky | Novel | 
    And the response status code should be 200

Scenario: Add a new book
  Given the following books exist in the system:
    | title         | author       | genre  | 
    | War and Peace | Leo Tolstoy  | Novel  | 
    | Foundation    | Isaac Asimov | Sci-Fi | 
   When I add a book with the following details:
    | title | author        | genre  | 
    | Dune  | Frank Herbert | Sci-Fi | 
   Then the response status code should be 201
    And the book list should include:
    | title         | author        | genre  | 
    | War and Peace | Leo Tolstoy   | Novel  | 
    | Foundation    | Isaac Asimov  | Sci-Fi | 
    | Dune          | Frank Herbert | Sci-Fi | 

Scenario: Add a new book with a missing field
  Given the following books exists in the system:
    | title         | author       | genre  | 
    | War and Peace | Leo Tolstoy  | Novel  | 
    | Foundation    | Isaac Asimov | Sci-Fi | 
   When I add a book with missing fields:
    | title         | author       | genre  | 
    | War and Peace | Leo Tolstoy  | Novel  | 
    |               | Isaac Asimov | Sci-Fi | 
   Then the response status code should be 400




