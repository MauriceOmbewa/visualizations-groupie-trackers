
---

# Groupie Tracker

Groupie Tracker is a web application that tracks music artists and provides detailed information about their concerts, locations, and relations with other artists. The app allows users to browse a list of artists, view details about a specific artist, and explore data related to their tours.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Usage](#usage)
- [API Integration](#api-integration)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## Overview

Groupie Tracker aims to provide a user-friendly platform for users to explore information about music artists, their concert locations, and their relationships with other artists. The project integrates an API that serves data related to artists and their events. The primary focus is to allow users to:
- Browse artists
- View specific artist details
- See event locations and dates
- Visualize artist relations

The application provides seamless interaction through client-server communication and ensures error handling for a smooth user experience.

## Features

- **Artist Listings**: View a list of artists fetched from the API.
- **Artist Details**: Detailed information on artists including:
  - Locations of their concerts
  - Dates of upcoming events
  - Relations with other artists
- **Error Handling**: Custom error pages for common HTTP errors like 404 and 500.
- **Data Visualization**: Present artist data in a clear and structured format.
  
## Project Structure

```
.
├── internal/
│   ├── api/
│   │   └── handlers.go       # Handles web request and responses
│   ├── fetch/
│   │   └── fetch.go          # Fetches data from the API
│   ├── models/
│       └── models.go         # Structs for Artists, Locations, Dates, and Relations
├── static/
│   ├── css/                  # Stylesheets for the UI
│   ├── js/                   # JavaScript files
├── templates/
│   ├── index.html            # Home page template
│   ├── artists.html          # Artists listing page
│   ├── artist_detail.html    # Artist detail page
│   └── error.html            # Error page template
├── main.go                   # Entry point of the application
├── go.mod                    # Go module file
└── README.md                 # Project documentation
```

## Tech Stack

- **Go (Golang)**: The backend server is written in Go.
- **HTML/CSS/JavaScript**: For the frontend to create the user interface.
- **API Integration**: Fetches data about artists, locations, dates, and relations from a provided API.
- **Text Templates**: Go's built-in HTML templating for rendering dynamic content.
- **Error Handling**: Custom error pages (404, 500).

## Installation

1. **Clone the repository**
   ```bash
   git clone https://learn.zone01kisumu.ke/git/seodhiambo/groupie-tracker
   cd groupie-tracker
   ```

2. **Install Go Dependencies**
   Make sure you have Go installed on your machine. Then run:
   ```bash
   go mod tidy
   ```

3. **Run the Application**
   ```bash
   go run .
   ```

4. **Access the Application**
   Open your browser and navigate to `http://localhost:8080`.

## Usage

### Endpoints

- `/`: Home page that provides a general overview of the project.
- `/artists`: Lists all the artists retrieved from the API.
- `/artist/{id}`: Detailed information about a specific artist, including concert locations, dates, and relations with other artists.

### Custom Error Handling
The application includes custom error handling for the following scenarios:
- **404 - Not Found**: When a user attempts to access a page that doesn't exist.
- **500 - Internal Server Error**: In case of any server-side issues.

These errors are handled gracefully using a custom `error.html` page that displays relevant messages to the user.

## API Integration

The project integrates with an external API that provides information about:
- **Artists**: Basic information about music artists.
- **Locations**: Where the artists are performing.
- **Dates**: When the artists are performing.
- **Relations**: Connections between different artists.

Data is fetched and stored in the following structs:
- `Artist`
- `LocationsData`
- `DatesData`
- `RelationsData`

### Fetch Data

The `fetch.FetchAllData()` function in the `fetch.go` file retrieves data from the external API and processes it into Go structs for further use.

## Error Handling

The application implements custom error handling for common HTTP errors:

- **404 Not Found**: If the user tries to access an invalid artist ID or a non-existent page.
- **500 Internal Server Error**: If there is a problem processing the request or rendering the templates.

The custom error handler dynamically loads the appropriate error message and status code to provide better user feedback.

```go
func ErrorHandler(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    errDetail := models.ErrorDetail{
        Title:   http.StatusText(statusCode),
        Message: message,
    }
    templates.ExecuteTemplate(w, "error.html", errDetail)
}
```

### Example of Dynamic Error Handling:
- `http://localhost:8080/artists/9999` will trigger a `404 Not Found` error if the artist ID doesn't exist.

## Contributing

1. **Fork the repository**
2. **Create a new branch**
   ```bash
   git checkout -b features
   ```
3. **Make your changes**
4. **Commit your changes**
   ```bash
   git commit -m "Add new feature"
   ```
5. **Push to the branch**
   ```bash
   git push origin features
   ```
6. **Open a pull request**

### **Contributors**

- **Seth Athooh** - Software Developer
- **Maurice Ombewa** - Software Developer

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---