# Online ticket sales system


This project is a microservices-based system for selling tickets online. It allows ticket purchases through partner APIs, acting as intermediaries, and integrates the backend and frontend through an API gateway. We use robust technologies, widely adopted in large systems and organizations.


# Required Technologies

 - ![Docker](https://img.shields.io/badge/Docker-blue)
 - ![Go](https://img.shields.io/badge/Go-blue)
 - ![Nest.ja](https://img.shields.io/badge/Nest.js-red)
 - ![Next.ja](https://img.shields.io/badge/Next.js-gray)
 - ![Kong](https://img.shields.io/badge/Kong-green)

# Project overview

```mermaid

graph TD
    subgraph KONG API GATEWAY
        A[Next.js \nTicket sales frontend] -->|http| B[Golang Backend\nSales API]
        B -->|http| C[Nest.js Backend 1\nPartners API 1]
        B -->|http| D[Nest.js Backend 2\nPartners API 2]
        
    end
        C --> E[(MySQL 1)]
        D --> F[(MySQL 2)]
        B --> G[(MySQL 3)]

```

# Operating flow

```mermaid
graph LR;
    APIGateway[API Gateway] --> Interface[System Interface];
    Interface --> Management[Event Management];
    Management --> PartnersAPI[Partners API];

```

```mermaid

graph LR
    A[User] -->|open/checkout<br>click the pay button| B[Frontend]
    B -->|http| C[Golang]
    C -->|http| D[Nest.js]
    
    B -->|Returns success or failure response| A
    C -->|Returns success or failure response| B
    D -->|Return whether you booked or not| C

```

# Project funcionalities

developing... 




## Credits

This project was created based on Full Cycle immersion



