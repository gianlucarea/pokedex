# pokedex

> _A lovingly crafted, terminal-based Pokédex CLI for the curious, the nostalgic, and the adventurous._

![Pokedex Banner](https://raw.githubusercontent.com/PokeAPI/media/master/logo/pokeapi_256.png)

---

## ✨ Features

- Explore the Pokémon world from your terminal
- Catch, inspect, and list Pokémon
- Navigate locations and areas using the PokeAPI
- ASCII art and playful CLI experience
- Written entirely in Go for speed and simplicity

---

## 🚀 Getting Started

### From Source

```sh
git clone https://github.com/gianlucarea/pokedex.git
cd pokedex
go build -o pokedex
./pokedex
```

---

## 🕹️ Usage

Once running, use the following commands:

| Command    | Description                        |
|------------|------------------------------------|
| `help`     | Show help menu                     |
| `exit`     | Exit the Pokédex                   |
| `map`      | Show next map locations            |
| `mapb`     | Show previous map locations        |
| `explore`  | Explore a location                 |
| `catch`    | Attempt to catch a Pokémon         |
| `inspect`  | Inspect a caught Pokémon           |
| `pokedex`  | List all caught Pokémon            |

Example:

```sh
> map
> explore Pallet-Town
> catch Pikachu
> pokedex
> inspect Pikachu
```

---

## 💡 Next Development Ideas

- Support the "up" arrow to cycle through previous commands
- Simulate battles between Pokémon
- Add more unit tests
- Refactor code for better organization and testability
- Keep Pokémon in a "party" and allow them to level up
- Allow caught Pokémon to evolve after a set amount of time
- Persist a user's Pokédex to disk for saved progress
- Use the PokeAPI to make exploration more interactive (e.g., area choices, directional commands)
- Random encounters with wild Pokémon
- Add support for different Poké Balls (Poké Ball, Great Ball, Ultra Ball, etc.)

---

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/awesome-feature`)
3. Commit your changes (`git commit -am 'Add awesome feature'`)
4. Push to the branch (`git push origin feature/awesome-feature`)
5. Open a Pull Request

All trainers welcome!

---

pokedex is a terminal application written in Go that brings the world of Pokémon to your command line. Powered by [PokeAPI](https://pokeapi.co/)

## 🧑‍💻 About

pokedex is a terminal application written in Go that brings the world of Pokémon to your command line. Powered by [PokeAPI](https://pokeapi.co/).

> **Note:** This project was created as a learning exercise while following the Go Bootcamp on [boot.dev](https://boot.dev/). If you're interested in learning Go through hands-on projects, check them out!
