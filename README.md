https://chenyitian.gitbooks.io/gin-tutorials/content/


Application Functionality

The application we'll build is a simple article manager. This application should:

1. Let users register with a username and a password (non-logged in users only)
2. Let users log in with a username and a password (non-logged in users only)
3. Let users log out (logged in users only)
4. Let users create new articles (logged in users only)
5. Display the list of all articles on the home page (for all users)
6. Display a single article on its own page (for all users)

* In addition to this functionality, the list of articles and a single article should be accessible in the HTML, JSON and XML formats.

In our application, we will:

1. Serve the index page at route / (HTTP GET request),

2. Group user-related routes under the /u route,
    - Serve the login page at /u/login (HTTP GET request),
    - Process the login credentials at /u/login (HTTP POST request),
    - Log out at /u/logout (HTTP GET request),
    - Serve the registration page at /u/register (HTTP GET request),
    - Process the registration information at /u/register (HTTP POST request) ,

3. Group article related routes under the /article route,
    - Serve the article creation page at /article/create (HTTP GET request),
    - Process the submitted article at /article/create (HTTP POST request), and
    - Serve the article page at /article/view/:article_id (HTTP GET request). Take note of the :article_id part in this route. The : at the beginning indicates that this is a dynamic route. This means that :article_id can contain any value and Gin will make this value available in the route handler.
