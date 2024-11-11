

---

# Groupie Tracker

Groupie Tracker is a web application that tracks music artists and provides detailed information about their concerts, locations, and relations with other artists. The app allows users to browse a list of artists, view details about a specific artist, and explore data related to their tours.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Search Functionality](#search-functionality)
- [Visualization Features](#visualization-features)
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
- Search for artists, members, locations, and more
- View specific artist details
- See event locations and dates
- Visualize artist relations and concert geolocation

The application provides seamless interaction through client-server communication and ensures error handling for a smooth user experience.

## Features

- **Artist Listings**: View a list of artists fetched from the API.
- **Artist Details**: Detailed information on artists including:
  - Locations of their concerts
  - Dates of upcoming events
  - Relations with other artists
  - **Concert Geolocation**: Geographical visualization of concert locations on an interactive map.
- **Search Functionality**: A search bar to search by artist, member, location, or other attributes.
- **Pagination**: Artist listing page with pagination, loading 20 artist cards per page.
- **Media Responsiveness**: Optimized for various screen sizes, with mobile-friendly layouts and dynamic components.
- **Error Handling**: Custom error pages for common HTTP errors like 404 and 500.
- **Data Visualization**: Present artist data in a clear and structured format with enhanced visual features for user navigation.

## Search Functionality

The search bar enables users to quickly find artists, members, concert locations, and more. The search feature handles the following cases:

- **Artist/Band Name**: Find an artist by their name.
- **Members**: Search for members of a band.
- **Locations**: Search by concert locations.
- **First Album Date**: Search based on the release date of an artist’s first album.
- **Creation Date**: Search by the artist’s creation date.

### Key Features of the Search Bar:

- **Case-Insensitive**: Searches are case-insensitive, making it easier to find results regardless of input case.
- **Typing Suggestions**: As you type, suggestions will appear, showing possible matches from multiple categories (artist, member, location, etc.).
- **Category Display**: The suggestions clearly identify the type of match (e.g., member or artist). For example, typing "phil" could show `Phil Collins - member` and `Phil Collins - artist/band`.

This search feature enhances the user experience by allowing quick access to detailed artist information.

## Visualization Features

The application provides rich visual elements to improve user interaction, including:

- **Arrow Navigation**:
  - **Arrow Up/Down**: Navigate vertically between artist cards or search results.
  - **Arrow Left/Right**: Navigate horizontally between artist cards within the grid layout.
- **Pagination**: Users can navigate through multiple pages of artists, with each page displaying 20 artist cards. Pagination controls include "Next" and "Previous" buttons and page indicators.
- **Geolocation**: For artists with concert locations, an interactive map is displayed with pinpointed locations of upcoming concerts. Each location marker shows the venue name and other relevant details.
- **Media Responsiveness**: The web app dynamically adjusts for different screen sizes (mobile, tablet, desktop) with optimized layouts. Navigation menus, search bars, and artist cards all adjust smoothly to fit various devices.

## Project Structure

```
.
├── internal/
│   ├── api/
│   │   └── handlers.go       # Handles web request and responses
│   ├── fetch/
│   │   └── fetch.go          # Fetches data from the API
│   ├── models/
│   |   └── models.go         # Structs for Artists, Locations, Dates, and Relations
|   |__ utiliies
|       |_geocode.go     
├── static/
│   ├── css/                  # Stylesheets for the UI
│   ├── js/                   # JavaScript files including search logic and pagination
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
- **Search Bar**: JavaScript-driven search bar with typing suggestions and categories.
- **Pagination**: Manage and load pages of artists dynamically with pagination controls.
- **Geolocation**: Integration with Mapbox for concert location visualization.
- **Error Handling**: Custom error pages (404, 500).

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/MauriceOmbewa/visualizations-groupie-trackers
   cd groupie-tracker-visualization
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
- `/artists`: Lists all the artists retrieved from the API with pagination (20 artists per page).
- `/artist/{id}`: Detailed information about a specific artist, including concert locations, dates, and relations with other artists.
- **Search Bar**: Use the search bar at the top of the site to quickly find artists, members, and more.

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

### Geolocation Integration

For artists' concert locations, the app integrates with Mapbox to provide an interactive map where users can see pinpointed concert locations. Each marker on the map displays the concert location and venue information, allowing users to visually explore where artists are performing.

### Fetch Data

The `fetch.FetchAllData()` function in the `fetch.go` file retrieves data from the external API and processes it into Go structs for further use.

## Error Handling

The application implements custom error handling for common HTTP errors:

- **404 Not Found**: If the user tries to access an invalid artist ID or a non-existent page.
- **500 Internal Server Error**: If there is a problem processing the request or rendering the templates.

The custom error handler dynamically loads the appropriate error message and status code to provide better user feedback.

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

### **Developer**

- [Maurice Ombewa](www.github.com/MauriceOmbewa) - Software Developer

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

