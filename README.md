# STRYK

This is a silly game made with Web Assembly (wasm), written in go. Its only real purpose was to write something with
wasm and see how effective it was.

# How do I run it?

Easiest way is to first run it by first building the whole shebang:

```
./build.sh
```

And then open up a local service, perhaps using python, like this:

```python
# version 3.X and above
python3 -m http.server

# version 2.X and below
python -m SimpleHTTPServer
```

Then visit http://localhost:8000