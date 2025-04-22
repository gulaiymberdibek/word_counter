##  Project Overview:

This beginner project is a simple word counter written in the Go programming language. The program reads a `.txt` file provided by the user, counts how often each word appears, and displays the **top 5 most frequent words** in the file.

---

## ⚙️ Technologies & Concepts Used

- **bufio** – for reading files line by line
- **os** – to open and handle files
- **strings** – for splitting lines into words
- **map[string]int** – to store word counts
- **sort** – to sort words by their frequency
- **Basic CLI** – to interact with the user via the terminal

---

##  How It Works:

1. The program asks the user to enter the path to a `.txt` file.
2. It reads the file line by line using a scanner.
3. Each line is split into words using `strings.Fields`.
4. Words are converted to lowercase to ensure case-insensitive counting.
5. The program stores and updates each word's count in a map.
6. After scanning, it sorts the words by frequency and displays the top 5.
