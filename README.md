# Web Forum Project README

Welcome to the Web Forum Project! This project aims to create a web forum where users can communicate, interact with posts and comments, and filter content. Below you'll find all the necessary information to understand, set up, and contribute to this project.
Objectives

## This project's main objectives are as follows:

    Communication: Users can create posts and comments, associating categories with posts.
    Authentication: Users can register, log in, and have their own sessions using cookies.
    Likes and Dislikes: Registered users can like or dislike posts and comments.
    Filter: Users can filter posts by categories, created posts, and liked posts.
    SQLite: The project uses SQLite for data storage, including users, posts, comments, etc.

## Getting Started

To get started with this project, follow these steps:
* Prerequisites

    Install Go (Go programming language).
    Install Docker for containerization.

* Installation

    Clone the project repository:
    
    >git clone https://learn.zone01dakar.sn/git/mouhametadiouf/forum.git 
    
    
    >cd forum

* Build and run the Docker container:

    -bash

    >docker build -t forum-app .

    >docker run -p 8080:8080 forum-app

    -use the script to run the Docker container

    >chmod +x docker_run.sh

     >./docker_run.sh

    Access the forum in your web browser at http://localhost:8080.

* User Registration and Authentication

    Users can register with their email, username, and password.
    Email addresses must be unique; duplicates are not allowed.
    Passwords are securely encrypted before storage.

* Communication

    Registered users can create posts and comments.
    Posts can be associated with one or more categories.
    Posts and comments are visible to all users, registered or not.

* Likes and Dislikes

    Only registered users can like or dislike posts and comments.
    The number of likes and dislikes is visible to all users.

* Filtering

    Users can filter posts by categories, created posts, and liked posts.
    Filtering by categories functions as subforums, focusing on specific topics.

* Dependencies

This project utilizes the following Go packages and libraries:

    All standard Go packages
    sqlite3: SQLite database management
    bcrypt: Password encryption
    UUID: Universally unique identifier (for a bonus task)

# License

This project is licensed under the Zone01 License.
Acknowledgments

If you have any questions or need assistance, please feel free to reach out to us. Happy coding!