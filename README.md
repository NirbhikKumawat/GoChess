# Knightmare ♟️

A blazingly fast, from-scratch chess engine written entirely in Go.

**Knightmare** is built on a highly optimized 64-bit **Bitboard** architecture. It is designed for maximum node throughput and strict mathematical accuracy, clocking in at over **5.5 Million Nodes Per Second (NPS)** on a single thread during move generation and legality verification.

## 🧠 Artificial Intelligence & Search

The engine (dubbed "Knightmare") features a custom-built, classical AI search stack capable of deep calculation and real-time time management:

* **Alpha-Beta Pruning:** The core search utilizes Minimax with Alpha-Beta cutoffs to exponentially reduce the size of the game tree.
* **Iterative Deepening:** Searches progressively deeper layer-by-layer, allowing the engine to flawlessly manage the clock and return the best possible move the millisecond time runs out.
* **Transposition Tables (TT):** Powered by custom Incremental Zobrist Hashing, the engine memorizes previously evaluated positions and exact hash moves to instantly prune massive branches of the search tree.
* **Quiescence Search:** Dynamically extends the search depth for capture sequences to ensure tactical stability and completely eliminate the "Horizon Effect" (blundering pieces at the end of a fixed depth).
* **Advanced Move Ordering:** Implements MVV-LVA (Most Valuable Victim - Least Valuable Attacker) alongside TT Hash Move prioritization to search the strongest moves first, maximizing Alpha-Beta efficiency.
* **Static Evaluation:** Understands positional advantages using Piece-Square Tables (PSTs) and standard material weighting.

## ⚡ Performance: The Perft Benchmark

The engine's move generator has been rigorously tested against the standard Perft (Performance Test) benchmarks from the starting position. It successfully traverses all **119,060,324 leaf nodes at Depth 6 in ~21 seconds** (single-threaded), proving 100% compliance with the rules of chess (including complex edge cases like en passant discovered checks, castling rights destruction, and promotions).

| Depth | Nodes       | Accuracy |
|:------|:------------|:---------|
| 1     | 20          | ✅ Pass   |
| 2     | 400         | ✅ Pass   |
| 3     | 8,902       | ✅ Pass   |
| 4     | 197,281     | ✅ Pass   |
| 5     | 4,865,609   | ✅ Pass   |
| 6     | 119,060,324 | ✅ Pass   |

## 🧠 Core Architecture & Technical Highlights

This engine avoids slow arrays and loops in favor of raw bitwise arithmetic and pre-calculated lookup tables.

* **Bitboards:** The board is represented by arrays of `uint64` integers. Piece movement, captures, and raycasting are resolved in fractions of a nanosecond using bitwise `AND`, `OR`, `XOR`, and bit-shifting.
* **Custom Magic Bitboards:** Sliding pieces (Rooks, Bishops, Queens) use generated Magic Bitboards to instantly look up attack rays. The engine includes a custom sparse-random brute-force generator to perfectly map blocker permutations to array indices, entirely bypassing expensive on-the-fly ray calculations.
* **Pseudo-Legal Move Generation:** Moves are generated in bulk (e.g., shifting entire pawn bitboards at once) rather than square-by-square.
* **Reverse Attack Legality Checking:** To verify King safety, the engine utilizes the "Super Piece" concept—projecting attacks outward from the King's square to detect overlapping enemy pieces, ensuring blistering fast legality checks.
* **Copy-Make Paradigm:** Taking advantage of Go's highly efficient struct copying, the engine uses a `MakeMove` approach on copied board states rather than a cumbersome `UnmakeMove` function, keeping the state mutation clean and fast.

## 📦 Features

Beyond the core AI, the engine is fully equipped to interact with the standard chess ecosystem:
* **UCI Protocol:** Fully compliant with the Universal Chess Interface. You can plug Knightmare into GUIs like Arena, CuteChess, or Lichess to play against humans or other engines.
* **FEN Support:** Flawless parsing and generation of Forsyth–Edwards Notation strings.
* **SAN Disambiguation:** Complete implementation of Standard Algebraic Notation (e.g., `Nbd2`, `exd8=Q#`).
* **PGN Parsing:** A custom, regex-free PGN parser capable of ingesting entire historical grandmaster games.

## 🚀 Getting Started

Ensure you have Go installed (1.18+ recommended).

### Installation
`git clone https://github.com/NirbhikKumawat/Knightmare.git`  
`cd Knightmare`

### Running the UCI Engine
Build and run the engine to start the Universal Chess Interface listener:  
`go build -o knightmare main.go`  
`./knightmare`  

### Running the Tests
To verify the move generator's accuracy and benchmark the speed on your hardware:  
`go test -v ./chess`  

## 🛣️ Roadmap
- [x] Bitboard Representation
- [x] Magic Bitboard Generation
- [x] Pseudo-Legal / Legal Move Generation
- [x] FEN / SAN / PGN I/O Layer
- [x] Static Evaluation (Piece-Square Tables, Material Weights)
- [x] Minimax Search with Alpha-Beta Pruning
- [x] Zobrist Hashing & Transposition Tables
- [x] UCI (Universal Chess Interface) Protocol Support


---
Made by [NirbhikTheNice](https://github.com/NirbhikKumawat/GoChess.git)