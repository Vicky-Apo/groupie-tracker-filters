# 🎸 Groupie Tracker

Groupie Tracker is a web application written in Go that allows users to explore and search for legendary music artists and bands. It loads and displays artist data (including name, members, creation date, album date, and locations) in a stylish, user-friendly interface.

---

## 🌐 Live Features

- Dynamic homepage listing artists (5 per row, with Load More functionality)
- Detailed artist pages with toggle buttons (Members, Locations, Dates)
- Search bar with live suggestions and advanced filters
- About page showcasing the development team
- Stylish, dark-themed interface with deep red glow effects

---

## 🔍 Advanced Search Bar

The search bar supports:

- ✅ Case-insensitive search
- ✅ Real-time suggestions as you type
- ✅ Filtering by:
  - Artist/Band Name
  - Member Name
  - Creation Year
  - First Album Date
  - Location

Example: Typing `phil` suggests:
- *Phil Collins – Artist*
- *Phil Collins – Member*

---

## 🗂️ Project Structure

```
.
├── internal
│   ├── data
│   │   ├── fetch.go         # Fetches data from external API
│   │   └── types.go         # Artist structs and types
│   ├── handlers
│   │   ├── detail.go        # Artist details page
│   │   ├── event.go         # (Optional) Event logic
│   │   ├── home.go          # Home page logic
│   │   └── search.go        # Search handler and logic
│   ├── routes
│   │   └── routes.go        # All routes defined here
│   └── utils
│       └── string_utils.go  # Custom utility functions (e.g., replace spaces)
│
├── static
│   ├── css
│   │   └── style.css        # Custom CSS styling
│   └── js
│       └── app.js           # Frontend JavaScript (load more, search bar etc.)
│
├── templates
│   ├── 404.html             # Not Found page
│   ├── about.html           # Team information
│   ├── artist.html          # Artist detail view
│   └── home.html            # Main landing page
│
├── go.mod
└── main.go                  # Entry point of the app
```

---

## ⚙️ How to Run

1. **Clone the repo**  

2. **Run the server**  
   ```bash
   go run main.go
   ```

3. **Visit in your browser**  
   Go to `http://localhost:8080`

---

## 🧠 Concepts & Skills Learned

- Golang (structs, HTTP, JSON, net/http, templates)
- Frontend integration (HTML, CSS, JavaScript)
- Dynamic DOM updates with JS
- Clean project structure & routing
- JSON data parsing
- Search algorithm implementation
- UI/UX styling (dark theme + glow effects)

---

## 👩‍💻 About the Team

We are a team of passionate developers at Zone01, working together to create a polished and functional user experience:

- **Kostas Apostolou** 
- **Yana Kopilova** 
- **Vicky Apostolou**  

---

## 📁 Resources

- [Groupie Tracker Subject Instructions](https://github.com/01-edu/public/blob/master/subjects/groupie-tracker/README.md)
- [Standard Go Documentation](https://pkg.go.dev/std)

---

## ❤️ Feedback

Feel free to open issues or contribute with ideas or pull requests!

---

## © License

Made with 💙 for music lovers at [Zone01](https://01.al).

