# Bartender-UI-v2

Magic Bartender is a touchscreen-controlled, Raspberry Pi-powered automated bartender. Built with a React frontend and Go backend, it intelligently detects available ingredients across multiple mixer stations and dynamically displays a list of cocktails that can be made in real time.

---

## 📚 Table of Contents

1. [Features](#features)
2. [Tech Stack](#tech-stack)
3. [Project Structure](#project-structure)
4. [Getting Started](#getting-started)
5. [Hardware Setup](#hardware-setup)
6. [Future Enhancements](#future-enhancements)
7. [License](#license)
8. [Feedback / Contributing](#feedback--contributing)
9. [Screenshots](#screenshots)

---

## <a id="features"></a> 🚀 Features

- 🧠 Smart cocktail detection based on selected mixers
- 🍸 Drink availability updates in real time
- 🧊 Interactive touchscreen UI built with React
- ⚙️ Go-based backend optimized for Raspberry Pi
- 🧃 Supports up to 8 mixer stations
- 💾 SQLite database for fast, low-memory operation
- 📦 Dockerized for easy deployment
- 🔌 GPIO control for pump activation (Raspberry Pi)
- 🧠 Redis caching for optimized performance

---

## <a id="tech-stack"></a> 🧱 Tech Stack

| Frontend       | Backend                | Infra / Other                   |
| -------------- | ---------------------- | ------------------------------- |
| React (Vite)   | Go                     | Raspberry Pi GPIO + Touchscreen |
| TypeScript     | SQLite                 | Docker, Redis (optional)        |
| SCSS           | REST API (custom)      |                                 |
| TanStack Query | Cocktail Recipe Engine |                                 |

---

## <a id="project-structure"></a> 📁 Project Structure

```
bartender-ui-v2/
├── frontend/ # React UI (Vite + TS)
├── backend/ # Go cocktail engine + GPIO controller
├── shared/ # Shared cocktail recipes, types
├── hardware/ # Scripts for GPIO and calibration
├── .env # Environment variables
├── Dockerfile # Containerization
├── docker-compose.yml # Multi-container orchestration
└── README.md
```

---

## <a id="getting-started"></a> 🌱 Getting Started

### Clone the Repo

```bash
git clone https://github.com/hifromkevin/bartender-ui-v2.git
cd bartender-ui-v2
```

### Install Dependencies and Run

#### Run Makefile Install (Recommended)

_From parent directory_

```bash
make install
```

#### Run Makefile Run (Recommended)

_From parent directory_

```bash
make start
```

#### Frontend Only

_From parent directory_
_(Not required if running Makefile)_

```bash
cd frontend
npm install
npm run build
```

#### Backend Only

_From parent directory_
_(Not required if running Makefile)_

```bash
cd backend
go run main.go
```

#### Run Docker

_COMING SOON_
_From parent directory_

```bash
docker-compose up --build
```

---

## Requirements

_Tested on the following versions_
**Node Version**: `v23.9.0`
**NPM and NPX Version**: `v11.3.0`
**Go Version**: `go1.24.2`

---

## <a id="hardware-setup"></a> 🛠️ Hardware Setup

| Component                     | Notes                                         |
| ----------------------------- | --------------------------------------------- |
| Raspberry Pi 5                | Required with GPIO + Touchscreen              |
| Peristaltic Pumps with tubing | 8x, controlled via GPIO pins                  |
| Power Supply                  | 12v, with 5v stepper for delicate electronics |
| Relays                        | 8 channel relay board for pump switching      |
| Touchscreen                   | Official Pi Touchscreen recommended           |

For a full wiring diagram, see /hardware/wiring.png (COMING SOON).

---

## <a id="future-enhancements"></a> 🔮 Future Enhancements

- Drink queue for multiple selections
- AI suggested cocktails based on user preference (available ingredients, flavor profile, mood, weather, etc.)
- Voice-controlled drink ordering (wake word + OpenAI)
- LED status lights for each station
- Liquid level detection per bottle
- Web-based remote ordering from phone/tablet

---

## <a id="license"></a> 📝 License

MIT License. See the [LICENSE](LICENSE) file for more details.

---

## <a id="feedback--contributing"></a> 💬 Feedback / Contributing

Feel free to open issues or PRs. More features and optimizations are on the horizon!

---

## <a id="screenshots"></a> 📷 Screenshots

![Bartender V2](homescreen.png?raw=true)
