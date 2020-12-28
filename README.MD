# Breakout!

This is a simple Breakout game in Go.

It works both compiled for a desktop PC or to WebAssembly. You can play it directly in your browser [here](https://cburchert.github.io/breakout/index.html).

## Running Manually

To quickly test changes, you can run as a desktop application:

```
go run src/*.go
```

You can rebuild the WebAssembly binary by running:

```
./rebuild.sh
```

You can then directly serve the main folder, e.g. with the default python web server:

```
python -m http.server
```

The submitted binary in `generated/breakout.wasm` is automatically deployed through GitHub Pages.