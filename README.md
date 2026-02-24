## CS Journal – Learning Journey

This repository is a personal learning journal where I practice core programming concepts across multiple languages and paradigms. It contains small exercises, experiments, and notes rather than a single cohesive app.

### Goals

- **Practice fundamentals**: Data structures, algorithms, I/O, and basic problem-solving.
- **Explore multiple languages**: Python, Go, C++, JavaScript, and React.
- **Build intuition**: Focus on understanding over polishing production-ready code.

### Project Structure

- **`python/`**: Small Python exercises and scripts.
  - `ordering_system.py`: Simple console ordering system that calculates subtotals, tax, and totals from a menu.
- **`golang/`**: Go programs to practice concurrency, networking, and algorithms.
  - `goSort.go`: Concurrent sorting of chunks of an array using goroutines and `sync.WaitGroup`.
  - `tcp.go`, `netp.go`, `findian.go`, etc.: Networking and string-processing exercises.
- **`cpp/`**: C++ programs focused on classes, objects, and basic OOP.
  - Files like `friends.cpp`, `member.cpp`, `employe.cpp`, `copy.cpp` explore constructors, copying, and relationships between classes.
- **`javascript/`**: JavaScript practice files.
  - Examples like `solution.js`, `finalProjetc.js`, `jesting.js` for basic logic and testing ideas.
- **`react/`**: React component experiments.
  - `event.jsx`: A simple component to practice state and event handling (currently intentionally broken as an exercise).
- **`sniffer/`**: Small tooling experiments.
  - `main.py`, `sniff.sh`, and a local `README.md` related to packet sniffing / scripting.
- **`uml/`**: UML diagrams and images used for planning or understanding designs.

### How to Run Examples

- **Python**
  - Make sure you have Python 3 installed.
  - From the project root:
    - `cd python`
    - Run a file, for example: `python ordering_system.py`
- **Go**
  - Install Go, then from the project root:
    - `cd golang`
    - Run a program, for example: `go run goSort.go`
- **C++**
  - Use `g++` (or another compiler). For example:
    - `cd cpp`
    - `g++ friends.cpp -o friends && ./friends`
- **JavaScript**
  - Use Node.js for simple scripts:
    - `cd javascript`
    - `node solution.js`
- **React**
  - `react/event.jsx` is a standalone component used for practice; it is not wired into a full React app yet. You can copy it into a React project (e.g. Vite or Create React App) and fix/experiment with it there.

### Notes

- **Work-in-progress**: Many files are intentionally rough or incomplete as part of learning exercises.
- **Experiments over perfection**: Code here may not follow strict style guides or best practices yet; the focus is on trying ideas and learning from mistakes.

