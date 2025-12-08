project plan.md

Here is your **Project Manifest**. You can save this as `README.md` or `project_plan.md` in your project folder. It contains everything a future LLM (or you) needs to pick up exactly where we left off.

````markdown
# Project: Texas Hold'em Poker Trainer (CLI)

## 1. Executive Summary
**Goal:** Build a terminal-based Texas Hold'em simulator focused on *training* rather than just playing.
**The Edge:** A real-time "Advisor Engine" that observes the game state and provides strategic advice, math checks (pot odds), and rule reminders based on poker course material.
**UI Philosophy:** Function over form. Clean Terminal UI (TUI) using Go.

---

## 2. Technology Stack & Rationale
* **Language:** **Go (Golang)**
    * *Why:* High performance for simulations, excellent concurrency for managing multiple bots, strict typing prevents state bugs, easy deployment.
* **UI Library:** **Bubble Tea** (by Charm)
    * *Why:* strict Model-View-Update (Elm architecture) allows for a robust, interactive terminal dashboard split into "Table" and "Advisor" views.
* **Math/Logic:** Custom Go implementation (no heavy dependencies).

---

## 3. Architectural Design
We use **Composition over Inheritance**. We strictly separate the **Mechanics** (Game Engine) from the **Intelligence** (Advisor).

### The Core Pattern: The Observer / Snapshot
The Game Engine never talks directly to the Advisor. Instead:
1.  **Engine** runs the mechanics.
2.  **Engine** creates a read-only `GameSnapshot` (a frozen moment in time).
3.  **Agent/Advisor** accepts the `GameSnapshot` to make decisions or generate tips.

### The Modules
1.  **Engine:** Handles deck, dealing, pot calculation, and turn order. Blind to advice.
2.  **Advisor:** A pure function. Input: `GameSnapshot`. Output: `Advice` (Text + Alert Level).
3.  **Agents:** Interface for Players. A `HumanAgent` calls the `Advisor` before prompting the user.

---

## 4. Suggested File Structure

```text
poker-guru/
├── go.mod                # Go module definition
├── main.go               # Entry point (initializes Game and UI)
├── model/                # SHARED DATA STRUCTURES (The "Language")
│   ├── types.go          # Card, Action, Player structs
│   └── snapshot.go       # The GameSnapshot struct (Crucial!)
├── engine/               # THE DEALER (Mechanics)
│   ├── game.go           # Main game loop, state management
│   ├── deck.go           # Shuffling, dealing logic
│   ├── evaluator.go      # Hand strength logic (Flush vs Straight)
│   └── pot.go            # Side-pot and split-pot math
├── advisor/              # THE COACH (Intelligence)
│   ├── brain.go          # Main Analyze() function
│   ├── math.go           # Pot odds, equity calculations
│   └── rules.go          # Specific course heuristics (e.g., "Open Ranges")
├── agents/               # THE PLAYERS
│   ├── interface.go      # The 'Agent' interface definition
│   ├── bot_random.go     # A stupid bot (for testing)
│   ├── bot_tight.go      # A tight-aggressive bot
│   └── human.go          # Connects UI input to Game Engine + Calls Advisor
└── ui/                   # THE VISUALS (Bubble Tea)
    ├── model.go          # The UI State
    ├── view.go           # How to render the table (ASCII art)
    └── update.go         # Handling keyboard events
````

-----

## 5\. Core Interface Definitions (The "Skeleton")

Use these definitions to jump-start the coding process.

### `/model/snapshot.go`

```go
package model

// The "Single Source of Truth" for the Advisor
type GameSnapshot struct {
    Round          string   // "PreFlop", "Flop", "Turn", "River"
    PotSize        int
    CommunityCards []Card
    
    // Hero Context
    HeroHand       []Card
    HeroPosition   string   // "BTN", "SB", "BB"
    HeroStack      int
    
    // Table Context
    Players        []PlayerSnapshot // Status of all players
    LastAction     Action           // What happened right before this? (e.g., "Villain Raised")
}

type PlayerSnapshot struct {
    Name       string
    Stack      int
    IsActive   bool
    CurrentBet int      // Amount in front of player this street
}
```

### `/agents/interface.go`

```go
package agents

import "poker-guru/model"

// Composition: Any struct with this method can play the game
type Agent interface {
    GetMove(snap model.GameSnapshot) model.Action
    GetName() string
}
```

### `/advisor/brain.go`

```go
package advisor

import "poker-guru/model"

type Advice struct {
    Suggestion string // "FOLD", "CALL", "RAISE"
    Reasoning  string // "Pot odds (20%) < Draw odds (35%)"
    Severity   string // "INFO", "WARNING", "CRITICAL"
}

// Pure function: Same input always results in same advice
func Analyze(snap model.GameSnapshot) Advice {
    // Logic goes here...
    return Advice{}
}
```

-----

## 6\. Implementation Roadmap

1.  **Phase 1: The Skeleton:** `go mod init`, create folder structure, define structs.
2.  **Phase 2: The Loop:** Build `engine` to deal cards and ask players for moves (CLI text only, no UI lib yet).
3.  **Phase 3: The Advisor:** Implement basic Pot Odds math in `advisor` and print it before the Human acts.
4.  **Phase 4: The TUI:** Implement `Bubble Tea` to make the split-screen UI.

<!-- end list -->

```

```
