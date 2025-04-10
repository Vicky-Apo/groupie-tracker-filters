document.addEventListener("DOMContentLoaded", () => {
    // Toggle section buttons
    document.querySelectorAll(".toggle-section").forEach(function (button) {
      button.addEventListener("click", function () {
        const targetId = this.getAttribute("data-target");
        const section = document.getElementById(targetId);
        if (section) {
          const isHidden = window.getComputedStyle(section).display === "none";
          section.style.display = isHidden ? "block" : "none";
          this.innerText = isHidden ? "Hide" : "Show";
        }
      });
    });
  
  


  // Back to previous page button
  const backBtn = document.getElementById("goBackBtn");
  if (backBtn) {
    backBtn.addEventListener("click", (e) => {
      e.preventDefault();
      const savedPage = localStorage.getItem("currentPage") || 1;
      window.location.href = `/home?page=${savedPage}`;
    });
  }
});