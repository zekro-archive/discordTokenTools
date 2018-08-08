# CHANGELOG

---

### 0.1.0

- init

---

### 0.2.0

- Renamed the projekt from "Token Checker" to "Discord Token Tools"
- Fixed missing guild information (guild name and owner ID)
- Added asynchronous collecting and displaying of guild list data next to general data
- Added changelog link
- Added version display below title
- Splitted JavaScript stuff into external files for better overview
- Moved library JavaScript files to assets/lib

---

### 0.2.1

- Removed DiscordLogin function and replaced it with GetTokenData function  
  *So the socket-io library is nor pulled out of the Discord functions which makes it a bit more abstract for other usages which could need this data maybe laster.*
- Added GetTokenValidity function to simply check if a token is valid or not
- Added a lot of comments for better understanding and easier contribution

---