# Hotel Ticket Booking App

A simple command-line application for booking hotel tickets, built with Go.

## Description

This application simulates a hotel ticket booking system. It allows users to book tickets by providing their first name, last name, email address, and the number of tickets they wish to purchase. The app tracks remaining tickets and maintains a list of all bookings.

### Features

- Interactive command-line interface
- User input validation for booking details
- Real-time tracking of remaining tickets
- Booking history display
- Continuous booking loop until manually stopped

## Requirements

- Go 1.22.5 or later

## How to Run

1. Ensure you have Go installed on your system.
2. Clone or download this repository.
3. Navigate to the project directory.
4. Run the application using:

   ```bash
   go run main.go
   ```

   Or build and run:

   ```bash
   go build -o booking-app main.go
   ./booking-app
   ```

## Usage

The application will prompt you for:
- First name
- Last name
- Email address
- Number of tickets

After booking, it will display the remaining tickets and all current bookings.

## Credits

Inspired by tutorials from [TechWorld with Nana](https://www.techworld-with-nana.com/)
