document.addEventListener("DOMContentLoaded", function() {
    let status = false; // Track whether focus is on search results or not

    const searchBar = document.getElementById('search-bar');
    const resultsContainer = document.getElementById('search-results');
    let currentIndex = -1; // Track the currently highlighted index

    // Function to search for artists based on user query
    function searchArtists(query) {
        console.log("Search query:", query);

        if (query.length === 0) {
            resultsContainer.innerHTML = ''; // Clear results if search query is empty
            return;
        }

        fetch(`/search?q=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(data => {
                resultsContainer.innerHTML = '';
                currentIndex = -1; // Reset the index when new search results are loaded

                // Display each search result
                data.forEach((result, index) => {
                    const resultItem = document.createElement('div');
                    resultItem.className = 'search-result-item';
                    resultItem.textContent = result.name;

                    // Direct the user to the artist detail page on click
                    resultItem.onclick = function() {
                        window.location.href = `/artist/${result.id}`;
                    };

                    resultsContainer.appendChild(resultItem);
                });

                resultsContainer.style.display = 'block'; // Show search results
            })
            .catch(error => console.error('Error fetching search results:', error));
    }

    // Handle input events in the search bar
    searchBar.addEventListener('input', function() {
        const query = searchBar.value;
        searchArtists(query);
        status = true; // Focus is on search results when typing in the search bar
    });

    // Hide the search results when clicking outside the search bar and results container
    document.addEventListener('click', function(event) {
        if (!searchBar.contains(event.target) && !resultsContainer.contains(event.target)) {
            resultsContainer.style.display = 'none'; // Hide search results
            status = false; // Focus is now on the main page content
        }
    });

    // Show the search results again if the user clicks on the search bar and there is still a query
    searchBar.addEventListener('focus', function() {
        if (searchBar.value.length > 0) {
            resultsContainer.style.display = 'block'; // Show results again if there's still a search query
            status = true; // Focus is back on the search results
        }
    });

    // Update the blur event listener for the search bar
    searchBar.addEventListener('blur', function() {
        // Use setTimeout to allow time for potential clicks on search results
        setTimeout(() => {
            if (!resultsContainer.contains(document.activeElement)) {
                resultsContainer.style.display = 'none';
                status = false; // Set focus back to page content
            }
        }, 100);
    });

    // Handle keydown events for both search results and page content
    document.addEventListener('keydown', function(event) {
        if (status) {
            // Handle arrow keys for search results
            handleSearchResultsNavigation(event);
        } else {
            // Handle arrow keys for page content
            handlePageContentNavigation(event);
        }
    });

    function handleSearchResultsNavigation(event) {
        const resultItems = resultsContainer.querySelectorAll('.search-result-item');
        if (resultItems.length === 0) return;

        if (event.key === 'ArrowDown') {
            event.preventDefault();
            if (currentIndex < resultItems.length - 1) {
                currentIndex++;
                highlightResult(resultItems);
            }
        } else if (event.key === 'ArrowUp') {
            event.preventDefault();
            if (currentIndex > 0) {
                currentIndex--;
                highlightResult(resultItems);
            }
        } else if (event.key === 'Enter' && currentIndex >= 0) {
            event.preventDefault();
            resultItems[currentIndex].click();
        }
    }

    function highlightResult(resultItems) {
        resultItems.forEach((item, index) => {
            if (index === currentIndex) {
                item.classList.add('highlighted');
                item.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
            } else {
                item.classList.remove('highlighted');
            }
        });
    }

    function handlePageContentNavigation(event) {
        const cards = document.querySelectorAll('.card');
        let currentCardIndex = -1;
        let numColumns = calculateColumns();

        // Find the currently selected card
        cards.forEach((card, index) => {
            if (card.classList.contains('selected')) {
                currentCardIndex = index;
            }
        });

        switch (event.key) {
            case 'ArrowRight':
                if (currentCardIndex === -1) {
                    currentCardIndex = 0;
                } else if (currentCardIndex < cards.length - 1) {
                    currentCardIndex++;
                }
                break;
            case 'ArrowLeft':
                if (currentCardIndex === -1) {
                    currentCardIndex = cards.length - 1;
                } else if (currentCardIndex > 0) {
                    currentCardIndex--;
                }
                break;
            case 'ArrowDown':
                if (currentCardIndex === -1) {
                    currentCardIndex = 0;
                } else if (currentCardIndex + numColumns < cards.length) {
                    currentCardIndex += numColumns;
                }
                break;
            case 'ArrowUp':
                if (currentCardIndex === -1) {
                    currentCardIndex = 0;
                } else if (currentCardIndex - numColumns >= 0) {
                    currentCardIndex -= numColumns;
                }
                break;
            case 'Enter':
                openSelectedCard();
                return;
            default:
                return;
        }

        event.preventDefault();
        selectCard(currentCardIndex);
    }

    function calculateColumns() {
        const container = document.getElementById('artists-container');
        const card = document.querySelector('.card');
        if (!container || !card) return 1;
        const cardWidth = card.offsetWidth;
        const containerWidth = container.offsetWidth;
        return Math.floor(containerWidth / cardWidth);
    }

    function selectCard(index) {
        const cards = document.querySelectorAll('.card');
        cards.forEach((card, i) => {
            card.classList.toggle('selected', i === index);
        });
        if (index >= 0 && index < cards.length) {
            cards[index].focus();
            cards[index].scrollIntoView({ behavior: 'smooth', block: 'nearest' });
        }
    }

    function openSelectedCard() {
        const selectedCard = document.querySelector('.card.selected');
        if (selectedCard) {
            const cardLink = selectedCard.closest('a');
            if (cardLink) {
                cardLink.click();
            }
        }
    }

    // Initialize map for artist details page
    const mapContainer = document.getElementById('map');
    if (mapContainer) {
        initializeMapbox();
    }

    function initializeMapbox() {
        mapboxgl.accessToken = 'your-mapbox-access-token';

        const concertLocations = JSON.parse(document.getElementById('ConcertLocationsJSON').innerText);

        const map = new mapboxgl.Map({
            container: 'map',
            style: 'mapbox://styles/mapbox/streets-v11',
            center: [0, 0],
            zoom: 2 // Default zoom level
        });

        concertLocations.forEach(function(location) {
            const popup = new mapboxgl.Popup({ offset: 25 }).setText(location.locationName);

            new mapboxgl.Marker()
                .setLngLat(location.coordinates)
                .setPopup(popup)
                .addTo(map);
        });

        // Ensure that the map resizes correctly
        map.resize();
    }

    // // Toggle functionality for search bar and navigation links
    // const hamburger = document.querySelector(".hamburger");
    // const navLinks = document.querySelector(".nav-links");
    // const searchIcon = document.querySelector(".search-icon");
    // const search = document.querySelector(".search");

    // hamburger.addEventListener("click", () => {
    //     navLinks.classList.toggle("active");
    // });

    // searchIcon.addEventListener("click", () => {
    //     search.classList.toggle("active");
    // });

    // Functionality for slider and output value in artist ratings
    const slider = document.getElementById("rating-range");
    const output = document.getElementById("rating-value");

    if (slider && output) {
        output.innerHTML = slider.value; // Set default value

        slider.oninput = function() {
            output.innerHTML = this.value; // Update value dynamically
        };
    }

    // Recalculate the number of columns when the window is resized
    window.addEventListener('resize', function() {
        calculateColumns();
    });
});

// Initialize map for artist details page
document.addEventListener("DOMContentLoaded", function() {
    const mapContainer = document.getElementById('map');
    if (mapContainer) {
        initializeMapbox();
    }
});

function initializeMapbox() {
    mapboxgl.accessToken = 'pk.eyJ1IjoiYXRob29oIiwiYSI6ImNtMWY2N3prZjJsN3MybHNjMWd3bThzOXcifQ.HNgAHQBkzGdrnuS1MtwYlQ';

    const concertLocations = JSON.parse(document.getElementById('ConcertLocationsJSON').innerText);

    const map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/streets-v11',
        center: [0, 0],
        zoom: 1 // Default zoom level
    });

    concertLocations.forEach(function(location) {
        const popup = new mapboxgl.Popup({ offset: 25 }).setText(location.locationName);

        new mapboxgl.Marker()
            .setLngLat(location.coordinates)
            .setPopup(popup)
            .addTo(map);
    });

    // Ensure that the map resizes correctly
    map.resize();
}

// Optional: Additional configuration if you want to control the slider behavior.
document.addEventListener('DOMContentLoaded', function() {
    const slideTrack = document.querySelector('.slide-track');
    const slideItems = document.querySelectorAll('.slide-item');

    // Clone the slide items to make the slider continuous
    slideItems.forEach(item => {
        slideTrack.appendChild(item.cloneNode(true));
    });
});

// header scrolls

document.addEventListener('DOMContentLoaded', function () {
    const header = document.querySelector('header');
    const hamburger = document.querySelector('.hamburger');
    const navLinks = document.querySelector('.nav-links');
    let lastScrollY = window.scrollY;

    // Hide/Show Header on Scroll
    window.addEventListener('scroll', () => {
        if (window.scrollY > lastScrollY) {
            header.style.top = '-100px';  // Hide header on scroll down
        } else {
            header.style.top = '0';       // Show header on scroll up
        }
        lastScrollY = window.scrollY;
    });
});

document.getElementById('hamburger').addEventListener('click', function() {
    const navMenu = document.getElementById('nav-menu');
    navMenu.classList.toggle('active');

    // Toggle no-scroll class to prevent body from scrolling when menu is open
    document.body.classList.toggle('no-scroll');
});z