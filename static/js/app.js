const searchFilter = document.getElementById("searchFilter");
const artistList = document.getElementById("artistList");
const paginationContainer = document.getElementById("pagination");
const backToTopBtn = document.getElementById("backToTopBtn");

let currentPage = 1;
const limit = 10;

function fetchArtists(page) {
    const offset = (page - 1) * limit;
    const url = `/api/artists?offset=${offset}&limit=${limit}`;
    console.log("Fetching from:", url);
  
    fetch(url)
      .then((res) => res.json())
      .then((data) => {
        const artists = data.artists;
        const total = data.total;
        
        currentPage = page; // ✅ This syncs currentPage to keep memory
        localStorage.setItem("currentPage", currentPage);
  
        displayArtists(artists);
        updatePagination(page, total); // pass total to updatePagination
      })
      .catch((error) => {
        console.error("Fetch failed:", error);
      });
  }
  

function displayArtists(artists) {
  artistList.innerHTML = "";
  artists.forEach((artist) => {
    const card = document.createElement("div");
    card.className = "artist-card";
    card.innerHTML = `
    <a href="/artist/${artist.name.replace(/\s+/g, "-")}" onclick="localStorage.setItem('currentPage', ${currentPage})">
      <img src="${artist.image}" alt="${artist.name}" class="artist-img" />
      <h3>${artist.name}</h3>
    </a>
  `;  
    artistList.appendChild(card);
  });
}

function updatePagination(currentPage, totalArtists) {
    paginationContainer.innerHTML = "";
    const totalPages = Math.ceil(totalArtists / limit);
  
    for (let i = 1; i <= totalPages; i++) {
      const btn = document.createElement("button");
      btn.className = "page-btn";
      btn.textContent = i;
      if (i === currentPage) btn.classList.add("active");
  
      btn.onclick = () => {
        localStorage.setItem("currentPage", i); // remember page
        fetchArtists(i);
        window.scrollTo({ top: 0, behavior: "smooth" });
      };
  
      paginationContainer.appendChild(btn);
    }
  }

 
// Show/hide back-to-top button
window.addEventListener("scroll", () => {
  if (backToTopBtn) {
    backToTopBtn.style.display = window.scrollY > 300 ? "block" : "none";
  }
});

if (backToTopBtn) {
  backToTopBtn.addEventListener("click", () => {
    window.scrollTo({ top: 0, behavior: "smooth" });
  });
}

document.addEventListener("DOMContentLoaded", () => {
  const params = new URLSearchParams(window.location.search);
  const shouldReset = params.get("reset");
  const urlPage = parseInt(params.get("page"));

  if (shouldReset) {
    localStorage.removeItem("currentPage");
    currentPage = 1;
  } else if (!isNaN(urlPage)) {
    currentPage = urlPage;
    localStorage.setItem("currentPage", currentPage); 
  } else {
    const savedPage = localStorage.getItem("currentPage");
    currentPage = savedPage ? parseInt(savedPage) : 1;
  }

  if (document.getElementById("artistList")) {
    fetchArtists(currentPage);
  }

  const homeLink = document.querySelector('a.nav-link[href="/"]');
  if (homeLink) {
    homeLink.setAttribute("href", "/?reset=true");
  }
});

