document.addEventListener("DOMContentLoaded", () => {
  // Grab references
  const toggleFiltersBtn = document.getElementById("toggleFiltersBtn");
  const filtersOverlay = document.getElementById("filtersOverlay");
  const locationSelect = document.getElementById("locationSelect");

  // We'll store filtered results here, plus track the current page for filtered items.
  let filteredArtistsData = [];
  let filteredCurrentPage = 1;
  const limit = 10; // show 10 at a time

  let allLocations = []; // Holds all fetched locations (unfiltered)

  function renderLocationCheckboxes(locations) {
    const container = document.getElementById("locationCheckboxes");
    if (!container) return;
  
    // Step 1: Track currently selected checkboxes (before re-render)
    const selected = new Set();
    document.querySelectorAll(".locationChk:checked").forEach((cb) =>
      selected.add(cb.value)
    );
  
    // Step 2: Clear existing checkboxes
    container.innerHTML = "";
  
    // Step 3: Rebuild checkboxes and restore checked state
    locations.forEach((loc, index) => {
      const id = `loc-${index}`;
      const label = document.createElement("label");
      label.setAttribute("for", id);
  
      // Check if this location was previously selected
      const isChecked = selected.has(loc) ? "checked" : "";
  
      label.innerHTML = `
        <input type="checkbox" class="locationChk" id="${id}" value="${loc}" ${isChecked} />
        ${loc}
      `;
      container.appendChild(label);
    });
  }

  function populateLocationCheckboxes(locations) {
    allLocations = locations; // Save for future filtering
    renderLocationCheckboxes(locations);
  }

  // 1) Fetch all possible concert locations and populate the select box.
  fetch("/api/all-locations")
  .then((res) => res.json())
  .then((locations) => {
    populateLocationCheckboxes(locations);
    if (locationSearchInput) {
      locationSearchInput.addEventListener("input", () => {
        const query = locationSearchInput.value.toLowerCase();
        const filtered = allLocations.filter((loc) =>
          loc.toLowerCase().includes(query)
        );
        renderLocationCheckboxes(filtered);
      });
    }
  })
  .catch((err) => {
    console.error("Error fetching all locations:", err);
  });

  // 2) Show/hide filter overlay logic
  if (toggleFiltersBtn && filtersOverlay) {
    toggleFiltersBtn.textContent = "Show Filters";
    toggleFiltersBtn.addEventListener("click", () => {
      if (!filtersOverlay.classList.contains("visible")) {
        // Show
        filtersOverlay.classList.add("visible");
        toggleFiltersBtn.textContent = "Hide Filters";
      } else {
        // Hide
        filtersOverlay.classList.remove("visible");
        toggleFiltersBtn.textContent = "Show Filters";
      }
    });
  }

  // 3) Grab references to the rest of our filter elements
  const minCreation = document.getElementById("minCreation");
  const maxCreation = document.getElementById("maxCreation");
  const minCreationValue = document.getElementById("minCreationValue");
  const maxCreationValue = document.getElementById("maxCreationValue");

  const minAlbum = document.getElementById("minAlbum");
  const maxAlbum = document.getElementById("maxAlbum");
  const minAlbumValue = document.getElementById("minAlbumValue");
  const maxAlbumValue = document.getElementById("maxAlbumValue");

  const membersCheckboxes = document.querySelectorAll(".membersChk");

  const applyFiltersBtn = document.getElementById("applyFiltersBtn");
  const clearFiltersBtn = document.getElementById("clearFiltersBtn");
  const filteredList = document.getElementById("filteredList");

  // 4) Update the range slider labels in real time
  if (minCreation && minCreationValue) {
    minCreation.addEventListener("input", () => {
      minCreationValue.textContent = minCreation.value;
    });
  }
  if (maxCreation && maxCreationValue) {
    maxCreation.addEventListener("input", () => {
      maxCreationValue.textContent = maxCreation.value;
    });
  }
  if (minAlbum && minAlbumValue) {
    minAlbum.addEventListener("input", () => {
      minAlbumValue.textContent = minAlbum.value;
    });
  }
  if (maxAlbum && maxAlbumValue) {
    maxAlbum.addEventListener("input", () => {
      maxAlbumValue.textContent = maxAlbum.value;
    });
  }

  // 5) "Apply Filters" button
  if (applyFiltersBtn) {
    applyFiltersBtn.addEventListener("click", () => {
      applyFilters();
    });
  }

  // 6) "Clear Filters" button
  if (clearFiltersBtn) {
    clearFiltersBtn.addEventListener("click", () => {
      clearFilters();
    });
  }

  // 7) Build and send query, then display results
  function applyFilters() {
    const mc = minCreation ? minCreation.value : "";
    const xc = maxCreation ? maxCreation.value : "";
    const ma = minAlbum ? minAlbum.value : "";
    const xa = maxAlbum ? maxAlbum.value : "";

    let selectedMembers = [];
    membersCheckboxes.forEach((cb) => {
      if (cb.checked) {
        selectedMembers.push(cb.value);
      }
    });

    let selectedLocations = [];
    document.querySelectorAll(".locationChk:checked").forEach((cb) => {
      selectedLocations.push(cb.value);
    });

    // Build the query string
    const params = new URLSearchParams();
    if (mc) params.append("minCreation", mc);
    if (xc) params.append("maxCreation", xc);
    if (ma) params.append("minAlbum", ma);
    if (xa) params.append("maxAlbum", xa);
    selectedMembers.forEach((m) => params.append("members", m));
    selectedLocations.forEach((loc) => params.append("location", loc));

    const url = "/api/filters?" + params.toString();
    fetch(url)
      .then((res) => res.json())
      .then((filteredArtists) => {
        displayFilteredArtists(filteredArtists);
      })
      .catch((err) => {
        console.error("Error fetching filtered artists:", err);
      });
  }

  // 8) Clear filters => revert to unfiltered home list
  function clearFilters() {
    // Reset sliders
    if (minCreation) {
      minCreation.value = 1950;
      if (minCreationValue) minCreationValue.textContent = "1950";
    }
    if (maxCreation) {
      maxCreation.value = 2025;
      if (maxCreationValue) maxCreationValue.textContent = "2025";
    }
    if (minAlbum) {
      minAlbum.value = 1950;
      if (minAlbumValue) minAlbumValue.textContent = "1950";
    }
    if (maxAlbum) {
      maxAlbum.value = 2025;
      if (maxAlbumValue) maxAlbumValue.textContent = "2025";
    }

    // Uncheck members
    membersCheckboxes.forEach((cb) => {
      cb.checked = false;
    });

    // Clear location boxes
    document.querySelectorAll(".locationChk").forEach((cb) => {
      cb.checked = false;
    });

    // Clear filtered results
    filteredList.innerHTML = "";
    filteredArtistsData = [];
    filteredCurrentPage = 1;

    // Re-fetch the default unfiltered list (if your home page uses fetchArtists)
    if (window.fetchArtists) {
      const currentPage = localStorage.getItem("currentPage") || 1;
      fetchArtists(parseInt(currentPage));
    }
  }

  // 9) Show the filtered results (or "No matching" if none)
  function displayFilteredArtists(artists) {
    const homepageList = document.getElementById("artistList");
    const pagination = document.getElementById("pagination");

    // Hide the default home list & pagination
    if (homepageList) homepageList.innerHTML = "";
    if (pagination) pagination.innerHTML = "";

    // If there's no container for filtered results, do nothing
    if (!filteredList) return;
    filteredList.innerHTML = "";

    // If empty
    if (!artists || artists.length === 0) {
      filteredList.innerHTML = "<p>No matching artists found.</p>";
      return;
    }

    // Store results in our global array
    filteredArtistsData = artists;
    filteredCurrentPage = 1;
    renderFilteredPage(filteredCurrentPage);
  }

  // 10) "Pagination" for the filtered results, 10 at a time
  function renderFilteredPage(page) {
    const pagination = document.getElementById("pagination");
    if (!filteredArtistsData || filteredArtistsData.length === 0) {
      filteredList.innerHTML = "<p>No matching artists found.</p>";
      return;
    }

    // Calculate page slice
    const start = (page - 1) * limit;
    const end = start + limit;
    const chunk = filteredArtistsData.slice(start, end);

    // Clear old results
    filteredList.innerHTML = "";

    // Render chunk
    chunk.forEach((artist) => {
      const card = document.createElement("div");
      card.className = "artist-card";
      card.innerHTML = `
        <a href="/artist/${artist.name.replace(/\s+/g, "-")}">
          <img src="${artist.image}" alt="${artist.name}" class="artist-img" />
          <h3>${artist.name}</h3>
        </a>
      `;
      filteredList.appendChild(card);
    });

    // Update pagination
    pagination.innerHTML = "";
    const totalPages = Math.ceil(filteredArtistsData.length / limit);
    for (let i = 1; i <= totalPages; i++) {
      const btn = document.createElement("button");
      btn.className = "page-btn";
      btn.textContent = i;
      if (i === page) {
        btn.classList.add("active");
      }
      btn.onclick = () => {
        filteredCurrentPage = i;
        renderFilteredPage(i);
        window.scrollTo({ top: 0, behavior: "smooth" });
      };
      pagination.appendChild(btn);
    }
  }
});
