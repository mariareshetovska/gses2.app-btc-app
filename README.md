# GSES2 BTC application

This is a GSES2 BTC application that provides an API for retrieving the current exchange rate of Bitcoin (BTC) in Ukrainian Hryvnia (UAH) and managing subscriptions for rate updates via email

# How to Use

1. Clone this repository

2. Install Golang if it is not already installed.

3. Create a .env file in the root directory of the project and add the following variables. For example:
    
        EMAIL_HOST="smtp.gmail.com"
        EMAIL_PORT= <port>
        EMAIL_ADDRESS= <email>
        EMAIL_PASSWORD= <password>
    
Replace  **EMAIL_PORT**,  **EMAIL_EMAIL**, and  **EMAIL_PASSWORD** with your own email server configuration.

4. Start aplication with Docker:
     
	   docker-compose up
	 
This will build and start the Docker containers required for the application.
     
## API Endpoints

#### Get Current Exchange Rate
- **URL**: http://localhost:8080/api/rate
- **Method:** GET
- **Description:** Retrieve the current exchange rate of BTC to UAH using a third-party public API
- **Response:** Returns the current exchange rate in JSON format

#### Subscribe to Rate Updates
- **URL**: http://localhost:8080/api/subscribe
- **Method:** POST
- **Description:** Subscribe an email address to receive rate update notifications.
- **Request Body:** Pass the email address in the request body as email.
- **Response:** Returns a success message if the email address is successfully subscribed.

#### Send Rate Update Emails
- **URL:** http://localhost:8080/api/sendEmails
- **Method:** POST
- **Description:** Send rate update emails to all subscribed email addresses.
- **Response:** Returns a success message once the emails are sent.

### Notes
- The application uses the specified third-party API to retrieve the exchange rate. Make sure the API is accessible and functioning correctly.
- The email server configuration should be provided in the .env file for the email subscription and notification functionality to work properly.
- For development, the application utilizes the CoinMarketCap API for retrieving the exchange rate of Bitcoin (BTC) in Ukrainian Hryvnia (UAH). For more information about the CoinMarketCap API [here](https://coinmarketcap.com/api/ "here")
