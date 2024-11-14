# GrinGram

GrinGram is a social media platform built with Go, MySQL, and Docker. It allows users to create posts, share images, comment, follow others, and manage profile pictures.

## Features

- **User Authentication**: Users can register, log in, and manage their profiles.
- **Profile Images**: Users can upload and manage their profile images.
- **Posts**: Users can create posts with text and images.
- **Comments**: Users can comment on posts and like comments.
- **Followers**: Users can follow other users and view posts from those they follow.
  
## Tech Stack

- **Backend**: Go (Golang)
- **Database**: MySQL
- **Containerization**: Docker
- **Migration**: `migrate` tool for database schema management
- **Authentication**: JWT for secure API authentication

## Installation

### Prerequisites

1. **Go**: Ensure you have Go installed. You can download it from [here](https://golang.org/dl/).
2. **MySQL**: Make sure MySQL is installed or use Docker to run MySQL.
3. **Docker** (Optional): For easier management of MySQL and the application.

### Steps

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/gringram.git
   cd gringram
