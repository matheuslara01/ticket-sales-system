# Application of partners for product testing

This repository contains partner applications that will be used for testing. This application uses the concept of NestJS libraries, where we have a core with the main functions that can be reused. In this case, we created two partners that reuse this core.

# Required Technologies

 - ![Docker](https://img.shields.io/badge/Docker-blue)
 - ![Nest.ja](https://img.shields.io/badge/Nest.js-red)
 - ![Prisma ORM](https://img.shields.io/badge/Prisma-blue)

# Project structure

This project consists of two central modules:

**1. apps:** Contains APIs for partners, specifically designed for testing purposes.
**2. libs:** Provides a reusable library intended for use by all partner APIs.

```plaintext

├── nestjs-partners-api
|   └── apps
|       ├── partner1
|       └── partner2
|   └── libs
|       └── core

````

# Funcionalities

### Event Management

  - **Event Creation**: Allows you to create new events with detailed information.
  - **Event Update**: Facilitates the editing of existing events.
  - **Event Deletion**: Removes events that are no longer necessary.
  - **Event List**: Displays all registered events.
    
### Spot Management
  - **Registration of Spot**: Allows you to register new places where events will be held.
  - **Spot Editing**: Makes it easier to update place information.
  - **Spot Removal**: Excludes seats that will no longer be used.
  - **List of Spots**: Displays all registered places.

### Reservations
  - **Reservation Creation**: Allows users to make reservations for events.

### Access via Token
  - **Authorization**: Controls access to specific resources based on defined permissions.
  - **Route Protection**: Ensures that only authenticated users access certain routes.

The two partner APIs created here are designed to simulate ticket searching, providing a realistic experience for testing.


## Running the Project

This project is configured to run using Docker and Dev Containers. Choose the best option to run the development environment.

### Using Docker

1. **Run with docker:**
   ```sh
     docker compose up 
   ```
### Using Dev Containers

1. **Open the project in Visual Studio Code.**
2. **Reopen the project in the Dev Container:**
 - Press `F1` and select `Dev Containers: Reopen in Container`.

## Commands

Here are the main commands to manage the project:

- **Run migrations with Prisma:**
   ```sh
      npm run migrate:partner1
      npm run migrate:partner2
   ```

- **Start the partner1 application:**
    ```sh
      npm run start:dev
    ```

- **Start the partner2 application:**
    ```sh
      npm run start:dev -- partner2
    ```

## Credits

This project was created based on Full Cycle immersion

