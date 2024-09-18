// Function to search for artists based on user query
function searchArtists(query) {
    console.log("Search query:", query);
    const resultsContainer = document.getElementById('search-results');

    // Clear previous results if query is empty
    if (query.length === 0) {
        resultsContainer.innerHTML = '';
        return;
    }

    // Fetch matching artists from the server
    fetch(`/search?q=${encodeURIComponent(query)}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            resultsContainer.innerHTML = '';

            // Display each search result
            data.forEach(result => {
                const resultItem = document.createElement('div');
                resultItem.className = 'search-result-item';
                resultItem.textContent = result.name;

                // Set up click event to fetch and display artist details
                resultItem.onclick = function() {
                    const artistId = result.id;

                    fetch(`/artist/${artistId}`)
                        .then(response => {
                            if (!response.ok) {
                                throw new Error('Failed to fetch artist details');
                            }
                            return response.text();
                        })
                        .then(html => {
                            const popup = document.getElementById("artist-popup");
                            const popupContent = document.getElementById("popup-artist-content");
                            popupContent.innerHTML = html;
                            popup.classList.remove("hidden");
                        })
                        .catch(error => {
                            console.error('Error fetching artist details:', error);
                            // Handle error display if needed
                        });
                };
                resultsContainer.appendChild(resultItem);
            });
        })
        .catch(error => console.error('Error fetching search results:', error));
}

// Initialize existing popup functionality for artist detail pages
document.addEventListener("DOMContentLoaded", function() {
    const artistLinks = document.querySelectorAll("#artists-container a");
    const popup = document.getElementById("artist-popup");
    const closeBtn = document.querySelector(".close-button");
    const popupContent = document.getElementById("popup-artist-content");

    // Set up click events for artist links
    artistLinks.forEach(link => {
        link.addEventListener("click", function(event) {
            event.preventDefault();
            const artistId = this.getAttribute("href").split("/").pop();

            fetch(`/artist/${artistId}`)
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Failed to fetch artist details');
                    }
                    return response.text();
                })
                .then(html => {
                    popupContent.innerHTML = html;
                    popup.classList.remove("hidden");
                })
                .catch(error => {
                    console.error('Error fetching artist details:', error);
                    // Handle error display if needed
                });
        });
    });

    // Close popup when close button is clicked
    closeBtn.addEventListener("click", function() {
        popup.classList.add("hidden");
    });

    // Close popup when clicking outside of it
    window.addEventListener("click", function(event) {
        if (event.target == popup) {
            popup.classList.add("hidden");
        }
    });
});

// Toggle functionality for search bar and navigation links
document.addEventListener("DOMContentLoaded", function() {
    const hamburger = document.querySelector(".hamburger");
    const navLinks = document.querySelector(".nav-links");
    const searchIcon = document.querySelector(".search-icon");
    const search = document.querySelector(".search");

    hamburger.addEventListener("click", () => {
        navLinks.classList.toggle("active");
    });

    searchIcon.addEventListener("click", () => {
        search.classList.toggle("active");
    });
});


// Optional: Additional configuration if you want to control the slider behavior.
document.addEventListener('DOMContentLoaded', function() {
    const slideTrack = document.querySelector('.slide-track');
    const slideItems = document.querySelectorAll('.slide-item');

    // Clone the slide items to make the slider continuous
    slideItems.forEach(item => {
        slideTrack.appendChild(item.cloneNode(true));
    });
});