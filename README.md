# Generator of Contracts - MVP
Service that generates PDF documents from a set of templates filled with the data of specific users. User data is expected to be generated through other services.

  - When a contract is requested, the service first finds the user data (by user_id) that includes template_id of the contract they signed.
  - Templates are stored as HTML files in the local forder and populated with the values of the previously fetched user. The resulting HTML is fed to a library that generates PDF. After this we return generated file.
- New templates can be added by creating GoogleDocs saved to HTML which are then manually converted into templates that the `html/template` package can process.

## Entities
 - Applicants
 - Templates

## Routes
 - GET `/contracts/:applicant_id` - returns PDF file that contains contract filled-in with user data

