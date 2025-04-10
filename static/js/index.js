document.addEventListener("DOMContentLoaded", () => {
  const homeLink = document.querySelector('a[href="/home"]');
  if (homeLink) {
    homeLink.addEventListener("click", () => {
      localStorage.removeItem("currentPage"); // reset page number
    });
  }
});

  