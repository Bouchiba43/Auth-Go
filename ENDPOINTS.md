# Ritual Growth API

The objective of this API is to allow manipulation of various entities such as projects, growth stages, cohorts, and users. It provides the capability to perform CRUD (Create, Read, Update, Delete) operations on these entities.

## Project Structure
- `cmd/` Main applications for this project
- `config/` For Configuration files
- `controllers/` Endpoint Handlers
- `routes/` Route definitions
- `repositories/` Data access layer
- `services/` For the Business Logic
- `models/` Database models
- `middlewares/` Middleware
- `docs/` For API documentation

### API Endpoints (current)

* `/api/v1/projects`
* `/api/v1/growth_stages`
* `/api/v1/cohorts`
* `/api/v1/users`
* `/api/v1/auth`
* `/api/v1/reports`

### Possible API Endpoints

**/api/v1/projects**

* `GET /api/v1/projects`: List all projects
* `GET /api/v1/projects/:id`: Get project details by ID
* `POST /api/v1/projects`: Create a new project
* `PUT /api/v1/projects/:id`: Update an existing project
* `DELETE /api/v1/projects/:id`: Delete a project
* `GET /api/v1/projects/:id/stages`: List growth stages for a specific project
* `GET /api/v1/projects/:id/users`: Retrieve users associated with a project
* `GET /api/v1/projects/:id/cohorts`: Get cohorts related to a specific project

**/api/v1/growth_stages**

* `GET /api/v1/growth_stages`: List all growth stages
* `GET /api/v1/growth_stages/:id`: Get growth stage details by ID
* `POST /api/v1/growth_stages`: Create a new growth stage
* `PUT /api/v1/growth_stages/:id`: Update an existing growth stage
* `DELETE /api/v1/growth_stages/:id`: Delete a growth stage
* `GET /api/v1/growth_stages/:id/stages`: List sub-growth stages for a specific growth stage
* `GET /api/v1/growth_stages/:id/users`: Retrieve users associated with a growth stage

**/api/v1/cohorts**

* `GET /api/v1/cohorts`: List all cohorts
* `GET /api/v1/cohorts/:id`: Get cohort details by ID
* `POST /api/v1/cohorts`: Create a new cohort
* `PUT /api/v1/cohorts/:id`: Update an existing cohort
* `DELETE /api/v1/cohorts/:id`: Delete a cohort
* `GET /api/v1/cohorts/:id/users`: Retrieve users associated with a specific cohort
* `GET /api/v1/cohorts/:id/stages`: List growth stages for a specific cohort

**/api/v1/events**

* `GET /api/v1/events`
* `POST /api/v1/events/project/:id`

**/api/v1/flows**


**/api/v1/users**

* `GET /api/v1/users`: List all users
* `GET /api/v1/users/:id`: Get user details by ID
* `POST /api/v1/users`: Create a new user
* `PUT /api/v1/users/:id`: Update an existing user
* `DELETE /api/v1/users/:id`: Delete a user
* `GET /api/v1/users/:id/projects`: Retrieve projects associated with a user
* `GET /api/v1/users/:id/growth_stages`: List growth stages for a specific user

**/api/v1/auth**

* `POST /api/v1/auth/login`: Login endpoint
* `POST /api/v1/auth/logout`: Logout endpoint
* `POST /api/v1/auth/register`: Registration endpoint (create new user)
* `PUT /api/v1/auth/reset_password`: Change password for existing user

Additional endpoints could be:

* `/api/v1/reports`: Generate reports for projects, growth stages, or cohorts.
* `/api/v1/notifications`: Manage notifications for users, such as project updates or cohort changes.