# Generator of Contracts - MVP
Service that generates PDF documents from a set of templates filled with the data of specific users. User data is expected to be generated through other services.

 - When a contract is requested, the service first fetches the user data (by user_id) that includes contract template_id.
- The template is requested on an internal endpoint using this id.
- Templates are stored as HTML and we can update this HTML with the values of the previously fetched user. The resulting HTML is fed to a library that generates PDF and returns the file.
- New templates can be added by creating GoogleDocs saved to HTML which then manually converted into template the `html/template` package can process.

## Entities
 - Users
 - Templates

## Routes
### External
 - GET /users/contracts/:user_id - returns PDF file with the filled-in contract

### Internal
 - GET /templates/:template_id - returns template
 - GET /users/:user_id - returns stored user data

